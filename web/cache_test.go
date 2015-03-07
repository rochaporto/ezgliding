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
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/rochaporto/ezgliding/airfield"
	"github.com/rochaporto/ezgliding/mock"
	"github.com/rochaporto/ezgliding/spatial"
)

func TestMemcacheBadServer(t *testing.T) {
	cfg := Config{Memcache: "nonexisting:10000", Static: "/tmp"}
	_, err := NewServer(cfg)
	if err == nil {
		t.Errorf("expected error got success")
		return
	} else if err != memcache.ErrNoServers {
		t.Errorf("unexpected error type :: %v", err)
		return
	}
}

func TestMemcache(t *testing.T) {
	airfields := []airfield.Airfield{
		airfield.Airfield{ID: "MockID", ShortName: "MockShortName", Name: "MockName",
			Region: "FR", ICAO: "AAAA", Flags: 0, Catalog: 11, Length: 1000, Elevation: 2000,
			Runway: "32R", Frequency: 123.45, Latitude: 32.533, Longitude: 100.376},
	}
	// get the expected geojson to compare at the end
	geojson, _ := spatial.Struct2GeoJSON([]interface{}{
		airfields[0],
	})
	expected, _ := json.MarshalIndent(geojson, "", "\t")

	// using a mock object to return airfields
	mock := &mock.Mock{
		GetAirfieldF: func(regions []string, updatedSince time.Time) ([]airfield.Airfield, error) {
			return airfields, nil
		},
	}

	// use a local memcached server
	cfg := Config{Airfielder: mock, Waypointer: mock, Static: "/tmp", Memcache: "localhost:11211"}
	srv, err := NewServer(cfg)
	if err != nil {
		t.Errorf("failed to init server with memcache :: %v", err)
		return
	}

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/airfield/", nil)
	req.Header.Set("Accept", "application/json")
	srv.mux.ServeHTTP(resp, req)
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read response :: %v", err)
		return
	}

	// first we compare the reply with the expected json
	if string(r) != string(expected) {
		t.Errorf("reply expected %v but got %v", string(expected), string(r))
		return
	}

	// then we compare the value in memcache with the expected json
	mclient := memcache.New(cfg.Memcache)
	a, err := mclient.Get(NewCacheResponseWriter(nil, nil, nil).requestKey(req))
	if err != nil {
		t.Errorf("failed to contact memcache :: %v", err)
		return
	}
	if string(a.Value) != string(expected) {
		t.Errorf("memcache expected %v but got %v", string(expected), string(a.Value))
		return
	}
}

type RequestKeyTest struct {
	h map[string][]string
	u string
	r string
}

var requestKeyTests = []RequestKeyTest{
	{
		map[string][]string{},
		"/airfield/",
		"/airfield/",
	},
	{
		map[string][]string{
			"Accept": []string{"application/json"},
		},
		"/airfield/",
		"/airfield/application/json",
	},
}

func TestRequestKey(t *testing.T) {
	c := NewCacheResponseWriter(nil, nil, nil)
	for _, test := range requestKeyTests {
		req, _ := http.NewRequest("GET", test.u, nil)
		req.Header = http.Header(test.h)
		r := c.requestKey(req)
		if r != test.r {
			t.Errorf("expected %v but got %v", test.r, r)
			continue
		}
	}
}
