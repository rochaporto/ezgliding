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

	"github.com/rochaporto/ezgliding/airfield"
	"github.com/rochaporto/ezgliding/mock"
	"github.com/rochaporto/ezgliding/spatial"
	"github.com/rochaporto/ezgliding/waypoint"
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
			airfield.Airfield{ID: "MockID", ShortName: "MockShortName", Name: "MockName",
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
			airfield.Airfield{ID: "MockID", ShortName: "MockShortName", Name: "MockName",
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
			waypoint.Waypoint{
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
			waypoint.Waypoint{
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

func serverFromTest(test ServerTest) *Server {
	airfields := []airfield.Airfield{}
	waypoints := []waypoint.Waypoint{}
	for _, a := range test.dt {
		switch a.(type) {
		case airfield.Airfield:
			airfields = append(airfields, a.(airfield.Airfield))
		case waypoint.Waypoint:
			waypoints = append(waypoints, a.(waypoint.Waypoint))
		}
	}
	mock := &mock.Mock{
		GetAirfieldF: func(regions []string, updatedSince time.Time) ([]airfield.Airfield, error) {
			return airfields, test.perr
		},
		GetWaypointF: func(regions []string, updatedSince time.Time) ([]waypoint.Waypoint, error) {
			return waypoints, test.perr
		},
	}
	cfg := Config{Static: "/tmp", Airfielder: mock, Waypointer: mock}
	srv, _ := NewServer(cfg)
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

func TestServerNewDefault(t *testing.T) {
	srv, _ := NewServer(Config{})
	if srv.Port != Port || srv.Static != Static {
		t.Errorf("unexpected web config :: %v %v", srv.Port, srv.Static)
		return
	}
}

type NewTest struct {
	t   string
	cfg Config
	err bool
}

var newTests = []NewTest{
	{
		"simple config",
		Config{Port: 8888, Static: "/tmp"},
		false,
	},
	{
		"config bad static",
		Config{Port: 8888, Static: "/does/not/exist"},
		true,
	},
}

func TestServerNewConfig(t *testing.T) {
	for _, test := range newTests {
		srv, err := NewServer(test.cfg)
		if err != nil && test.err {
			continue
		}
		if err != nil {
			t.Errorf("%v failed to init server :: %v", test.t, err)
			continue
		}
		result := Config{}
		result.Port = srv.Port
		result.Static = srv.Static
		if !reflect.DeepEqual(result, test.cfg) {
			t.Errorf("%v failed :: expected %v got %v", test.t, test.cfg, result)
			continue
		}
	}
}

func TestServerStart(t *testing.T) {
	cfg := Config{}
	cfg.Port = 7777
	cfg.Static = "/tmp"
	srv, err := NewServer(cfg)
	if err != nil {
		t.Errorf("failed to init server :: %v", err)
		return
	}
	go srv.Start()
}

func TestServerStartBadPort(t *testing.T) {
	cfg := Config{}
	cfg.Port = 1
	cfg.Static = "/tmp"
	srv, err := NewServer(cfg)
	if err != nil {
		t.Errorf("failed to init server :: %v", err)
		return
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
