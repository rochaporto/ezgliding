// Copyright 2014 The ezgliding Authors.
//
// This file is part of ezgliding.
//
// ezgliding is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ezgliding is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ezgliding.  If not, see <http://www.gnu.org/licenses/>.
//
// Author: Ricardo Rocha <rocha.porto@gmail.com>

package web

import (
	"compress/gzip"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/config"
	"github.com/rochaporto/ezgliding/context"
	"github.com/rochaporto/ezgliding/mock"
	"github.com/rochaporto/ezgliding/spatial"
)

type ServerTest struct {
	t    string
	dt   []interface{}
	rq   string
	fmt  string
	zip  bool
	err  bool
	perr error
}

var serverTests = []ServerTest{
	{
		"query airfield json",
		[]interface{}{
			common.Airfield{ID: "MockID", ShortName: "MockShortName", Name: "MockName",
				Region: "FR", ICAO: "AAAA", Flags: 0, Catalog: 11, Length: 1000, Elevation: 2000,
				Runway: "32R", Frequency: 123.45, Latitude: 32.533, Longitude: 100.376},
		},
		"/airfield/?region=FR&updated=2012-12-12",
		"application/json",
		false,
		false,
		nil,
	},
	{
		"query airfield json accept in querystring",
		[]interface{}{
			common.Airfield{ID: "MockID", ShortName: "MockShortName", Name: "MockName",
				Region: "FR", ICAO: "AAAA", Flags: 0, Catalog: 11, Length: 1000, Elevation: 2000,
				Runway: "32R", Frequency: 123.45, Latitude: 32.533, Longitude: 100.376},
		},
		"/airfield/?region=FR&updated=2012-12-12&accept=application/json",
		"",
		false,
		false,
		nil,
	},
	{
		"query waypoint json",
		[]interface{}{
			common.Waypoint{
				ID: "FURKAP", Name: "FURKAP", Description: "FURKAPASS PASSHOEHE", Elevation: 2432,
				Latitude: 46.572, Longitude: 8.415, Region: "CH", Flags: 0,
			},
		},
		"/waypoint/?region=CH&updated=2012-12-12",
		"application/json",
		false,
		false,
		nil,
	},
	{
		"query waypoint json zipped",
		[]interface{}{
			common.Waypoint{
				ID: "FURKAP", Name: "FURKAP", Description: "FURKAPASS PASSHOEHE", Elevation: 2432,
				Latitude: 46.572, Longitude: 8.415, Region: "CH", Flags: 0,
			},
		},
		"/waypoint/?region=CH&updated=2012-12-12",
		"application/json",
		true,
		false,
		nil,
	},
	{
		"invalid update in airfield",
		[]interface{}{},
		"/airfield/?region=FR&updated=20a12-12-12",
		"application/json",
		false,
		true,
		nil,
	},
	{
		"invalid update in waypoint",
		[]interface{}{},
		"/waypoint/?region=FR&updated=20a12-12-12",
		"application/json",
		false,
		true,
		nil,
	},
	{
		"invalid output format in airfield",
		[]interface{}{},
		"/airfield/",
		"unknown/format",
		false,
		true,
		nil,
	},
	{
		"invalid update in waypoint",
		[]interface{}{},
		"/waypoint/",
		"unknown/format",
		false,
		true,
		nil,
	},
	{
		"plugin err in airfield",
		[]interface{}{},
		"/airfield/",
		"application/json",
		false,
		true,
		errors.New("plugin err in airfield"),
	},
	{
		"plugin err in waypoint",
		[]interface{}{},
		"/waypoint/",
		"application/json",
		false,
		true,
		errors.New("plugin err in waypoint"),
	},
}

func serverFromTest(test ServerTest) Server {
	airfields := []common.Airfield{}
	waypoints := []common.Waypoint{}
	for _, a := range test.dt {
		switch a.(type) {
		case common.Airfield:
			airfields = append(airfields, a.(common.Airfield))
		case common.Waypoint:
			waypoints = append(waypoints, a.(common.Waypoint))
		}
	}
	mock := &mock.Mock{
		GetAirfieldF: func(regions []string, updatedSince time.Time) ([]common.Airfield, error) {
			return airfields, test.perr
		},
		GetWaypointF: func(regions []string, updatedSince time.Time) ([]common.Waypoint, error) {
			return waypoints, test.perr
		},
	}
	ctx := context.Context{Airfield: mock, Waypoint: mock}
	ctx.Config.Web.Static = "/tmp"
	srv := Server{}
	srv.Init(ctx)
	return srv
}

func TestServer(t *testing.T) {
	for _, test := range serverTests {
		var err error
		resp := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.rq, nil)
		req.Header.Set("Accept", test.fmt)
		if test.zip {
			req.Header.Set("Accept-Encoding", "gzip")
		}

		srv := serverFromTest(test)
		srv.mux.ServeHTTP(resp, req)
		var r []byte
		if test.zip {
			zipr, err := gzip.NewReader(resp.Body)
			if err != nil {
				t.Errorf("%v :: failed to create zip reader :: %v", test.t, err)
				continue
			}
			defer zipr.Close()
			r, err = ioutil.ReadAll(zipr)
		} else {
			r, err = ioutil.ReadAll(resp.Body)
		}
		if err != nil {
			t.Errorf("%v :: failed to read response :: %v", test.t, err)
			continue
		}
		if resp.Code != 200 && test.err {
			continue
		} else if resp.Code != 200 {
			t.Errorf("%v failed :: %v", test.t, resp)
			continue
		}
		result, err := spatial.GeoJSON2Struct(string(r))
		if err != nil {
			t.Errorf("%v :: failed to convert response :: %v :: %v", test.t, err, result)
			continue
		}
		if len(result) != len(test.dt) {
			t.Errorf("expected %v got %v", len(result), len(test.dt))
		}
	}
}

func TestServerInitDefault(t *testing.T) {
	srv := Server{}
	ctx := context.Context{}
	ctx.Config.Global.Airfielder = "randomairfield" // just to have non zero value ctx
	_ = srv.Init(ctx)
	if srv.Port != Port || srv.Static != Static {
		t.Errorf("unexpected web config :: %v %v", srv.Port, srv.Static)
		return
	}
}

func TestServerInitBadContext(t *testing.T) {
	srv := Server{Port: 7777, Static: "/tmp"}
	err := srv.Init(context.Context{})
	if err == nil {
		t.Errorf("expected error got success")
	}
}

type InitTest struct {
	t   string
	cfg config.Web
	err bool
}

var initTests = []InitTest{
	{
		"simple config",
		config.Web{Port: 8888, Static: "/tmp"},
		false,
	},
	{
		"config bad static",
		config.Web{Port: 8888, Static: "/does/not/exist"},
		true,
	},
}

func TestServerInitConfig(t *testing.T) {
	for _, test := range initTests {
		srv := Server{}
		ctx := context.Context{}
		ctx.Config.Web = test.cfg
		err := srv.Init(ctx)
		if err != nil && test.err {
			continue
		}
		if err != nil {
			t.Errorf("%v failed to init server :: %v", test.t, err)
			continue
		}
		result := context.Context{}
		result.Config.Web.Port = srv.Port
		result.Config.Web.Static = srv.Static
		if !reflect.DeepEqual(result.Config.Web, ctx.Config.Web) {
			t.Errorf("%v failed :: expected %v got %v", test.t, test.cfg, result.Config.Web)
			continue
		}
	}
}

func TestServerStart(t *testing.T) {
	srv := Server{}
	ctx := context.Context{}
	ctx.Config.Web.Port = 7777
	ctx.Config.Web.Static = "/tmp"
	err := srv.Init(ctx)
	if err != nil {
		t.Errorf("failed to init server :: %v", err)
	}
	go srv.Start()
}

func TestServerStartBadPort(t *testing.T) {
	srv := Server{}
	ctx := context.Context{}
	ctx.Config.Web.Port = 1
	ctx.Config.Web.Static = "/tmp"
	err := srv.Init(ctx)
	if err != nil {
		t.Errorf("failed to init server :: %v", err)
	}
	srv.Start()
}

func TestToOutputJSONBadType(t *testing.T) {
	srv := Server{}
	// pass an unsupported type
	_, err := srv.toOutput("application/json", []interface{}{1})
	if err == nil {
		t.Errorf("expected error got success")
		return
	}
}

func TestAccept(t *testing.T) {
	srv := Server{}

	r := srv.accept("[text/html,application/xhtml+xml,application/json;q=0.9,image/webp,*/*;q=0.8]")
	if r != "application/json" {
		t.Errorf("bad result for application/json :: %v", r)
	}

	r = srv.accept("[text/html,application/xhtml+xml,application/csv;q=0.9,image/webp,*/*;q=0.8]")
	if r != "application/csv" {
		t.Errorf("bad result for application/csv :: %v", r)
	}

	r = srv.accept("[text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8]")
	if r != "" {
		t.Errorf("bad result for unknown accept :: %v", r)
	}
}
