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

package fusiontables

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rochaporto/ezgliding/airfield"
)

type GetAirfieldTest struct {
	t   string
	rg  []string
	tm  time.Time
	rp  string
	rs  []airfield.Airfield
	err bool
}

var getGetAirfieldTests = []GetAirfieldTest{
	{
		"simple query",
		[]string{"FR"},
		time.Time{},
		`
ID,ShortName,Name,Region,RLICAO,Flags,Catalog,Length,Elevation,Runway,Frequency,Latitude,Longitude
HABER,HABER,HABERE POC69,FR,,1032,0,0,1113,0119,122.5,46.270,6.463
`,
		[]airfield.Airfield{
			airfield.Airfield{ID: "HABER", ShortName: "HABER", Name: "HABERE POC69",
				Region: "FR", ICAO: "", Flags: 1032, Catalog: 0, Length: 0, Elevation: 1113,
				Runway: "0119", Frequency: 122.5, Latitude: 46.270, Longitude: 6.463},
		},
		false,
	},
	{
		"query with invalid response",
		[]string{"FR"},
		time.Time{},
		`
ID,ShortName,Name,Region,ICAO,Flags,Catalog,Length,Elevation,Runway,Frequency,Latitude,Longitude
HABER,HABER,HABERE POC69,FR,,1032,0,0,1113,0119,122.5,46.270,6.463,a
`,
		[]airfield.Airfield{airfield.Airfield{}},
		true,
	},
	{
		"query with updated time",
		[]string{"FR"},
		time.Date(2012, time.January, 25, 0, 0, 0, 0, time.UTC), // FIXME: figure out how to actually filter here
		`
ID,ShortName,Name,Region,ICAO,Flags,Catalog,Length,Elevation,Runway,Frequency,Latitude,Longitude
HABER,HABER,HABERE POC69,FR,,1032,0,0,1113,0119,122.5,46.270,6.463
`,
		[]airfield.Airfield{
			airfield.Airfield{ID: "HABER", ShortName: "HABER", Name: "HABERE POC69",
				Region: "FR", ICAO: "", Flags: 1032, Catalog: 0, Length: 0, Elevation: 1113,
				Runway: "0119", Frequency: 122.5, Latitude: 46.270, Longitude: 6.463},
		},
		false,
	},
}

func TestGetAirfield(t *testing.T) {
	for _, at := range getGetAirfieldTests {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, at.rp)
		}))
		defer ts.Close()

		cfg := Config{}
		cfg.BaseURL = ts.URL
		cfg.AirfieldTableID = "testairfieldid"
		plugin, err := New(cfg)
		if err != nil {
			t.Errorf("failed to get fusiontables :: %v", err)
			continue
		}
		airfields, err := plugin.GetAirfield(at.rg, at.tm)
		if err != nil && at.err {
			continue
		} else if err != nil {
			t.Errorf("Failed to get airfields in test %v :: %v", at.t, err)
			continue
		}

		if len(airfields) != len(at.rs) {
			t.Errorf("Expected %v but got %v airfields", len(at.rs), len(airfields))
		}
		for i, airfield := range airfields {
			if airfield != at.rs[i] {
				t.Errorf("Expected %v but got %v", at.rs[i], airfield)
			}
		}
	}
}

func TestGetAirfieldWithMissingLocation(t *testing.T) {
	cfg := Config{}
	cfg.BaseURL = "http://doesnotexist"
	cfg.AirfieldTableID = "testairfieldid"
	plugin, err := New(cfg)
	if err != nil {
		t.Errorf("failed to get fusiontables :: %v", err)
		return
	}
	_, err = plugin.GetAirfield([]string{"FR"}, time.Time{})
	if err == nil {
		t.Errorf("expected error but was successful")
	}
}

func TestGetAirfieldWithMalformedURL(t *testing.T) {
	cfg := Config{}
	cfg.BaseURL = "wrong%url"
	cfg.AirfieldTableID = "testairfieldid"
	plugin, err := New(cfg)
	if err != nil {
		t.Errorf("failed to get fusiontables :: %v", err)
		return
	}
	_, err = plugin.GetAirfield([]string{"FR"}, time.Time{})
	if err == nil {
		t.Errorf("expected error but was successful")
	}
}

type PutAirfieldTest struct {
	t   string
	in  []airfield.Airfield
	csv string
	err bool
}

var putAirfieldTests = []PutAirfieldTest{
	{
		"simple update",
		[]airfield.Airfield{
			airfield.Airfield{ID: "HABER", ShortName: "HABER", Name: "HABERE POC69",
				Region: "FR", ICAO: "", Flags: 1032, Catalog: 0, Length: 0, Elevation: 1113,
				Runway: "0119", Frequency: 122.5, Latitude: 46.270, Longitude: 6.463,
				Update: time.Time{}},
		},
		`ID,ShortName,Name,Region,ICAO,Flags,Catalog,Length,Elevation,Runway,Frequency,Latitude,Longitude,Update
HABER,HABER,HABERE POC69,FR,,1032,0,0,1113,0119,122.5,46.27,6.463,0001-01-01 00:00:00 +0000 UTC
`,
		false,
	},
	{
		"simple failure",
		[]airfield.Airfield{
			airfield.Airfield{ID: "HABER", ShortName: "HABER", Name: "HABERE POC69",
				Region: "FR", ICAO: "", Flags: 1032, Catalog: 0, Length: 0, Elevation: 1113,
				Runway: "0119", Frequency: 122.5, Latitude: 46.270, Longitude: 6.463,
				Update: time.Time{}},
		},
		`ID,ShortName,Name,Region,ICAO,Flags,Catalog,Length,Elevation,Runway,Frequency,Latitude,Longitude,Update
aHABER,HABER,HABERE POC69,FR,,1032,0,0,1113,0119,122.5,46.27,6.463,0001-01-01 00:00:00 +0000 UTC
`,
		true,
	},
}

func TestPutAirfield(t *testing.T) {
	for _, test := range putAirfieldTests {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			content, _ := ioutil.ReadAll(r.Body)
			str := string(content)
			if str != test.csv {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("expected\n" + test.csv + "\ngot\n" + string(content)))
			} else {
				w.WriteHeader(http.StatusOK)
			}
		}))
		defer ts.Close()

		cfg := Config{}
		cfg.UploadURL = ts.URL
		cfg.AirfieldTableID = "testairfieldid"
		plugin, err := New(cfg)
		if err != nil {
			t.Errorf("failed to get fusiontables :: %v", err)
			continue
		}
		err = plugin.PutAirfield(test.in)
		if err != nil && test.err {
			continue
		} else if err != nil {
			t.Errorf("%v failed :: %v", test.t, err)
		}
	}
}

func TestPutAirfieldWithMissingLocation(t *testing.T) {
	cfg := Config{}
	cfg.BaseURL = "http://thisurlreallydoesnotexist.pt"
	cfg.AirfieldTableID = "testairfieldid"
	plugin, err := New(cfg)
	if err != nil {
		t.Errorf("failed to get fusiontables :: %v", err)
		return
	}
	err = plugin.PutAirfield([]airfield.Airfield{})
	if err == nil {
		t.Errorf("expected error but was successful")
	}
}

func TestPutAirfieldWithMalformedURL(t *testing.T) {
	cfg := Config{}
	cfg.BaseURL = "wrong%url"
	cfg.AirfieldTableID = "testairfieldid"
	plugin, err := New(cfg)
	if err != nil {
		t.Errorf("failed to get fusiontables :: %v", err)
		return
	}
	err = plugin.PutAirfield([]airfield.Airfield{})
	if err == nil {
		t.Errorf("expected error but was successful")
	}
}

func TestPutAirfieldWithBadStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	cfg := Config{}
	cfg.BaseURL = ts.URL
	cfg.AirfieldTableID = "testairfieldid"
	plugin, err := New(cfg)
	if err != nil {
		t.Errorf("failed to get fusiontables :: %v", err)
		return
	}
	err = plugin.PutAirfield([]airfield.Airfield{
		airfield.Airfield{ID: "aHABER", ShortName: "HABER", Name: "HABERE POC69",
			Region: "FR", ICAO: "", Flags: 1032, Catalog: 0, Length: 0, Elevation: 1113,
			Runway: "0119", Frequency: 122.5, Latitude: 46.270, Longitude: 6.463},
	})
	if err == nil {
		t.Errorf("expected error but got success")
	}
}
