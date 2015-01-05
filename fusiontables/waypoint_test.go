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

	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/config"
)

type GetWaypointTest struct {
	t   string
	rg  []string
	tm  time.Time
	rp  string
	rs  []common.Waypoint
	err bool
}

var getGetWaypointTests = []GetWaypointTest{
	{
		"simple query",
		[]string{"CH"},
		time.Time{},
		`
ID,Name,Description,Region,Flags,Elevation,Latitude,Longitude
FURKAP,FURKAP,FURKAPASS PASSHOEHE,CH,0,2432,46.270,6.463
`,
		[]common.Waypoint{
			common.Waypoint{ID: "FURKAP", Name: "FURKAP", Description: "FURKAPASS PASSHOEHE",
				Region: "CH", Flags: 0, Elevation: 2432,
				Latitude: "46.270", Longitude: "6.463"},
		},
		false,
	},
	{
		"query with invalid response",
		[]string{"CH"},
		time.Time{},
		`
ID,Name,Description,Region,Flags,Elevation,Latitude,Longitude
FURKAP,FURKAP,FURKAPASS PASSHOEHE,CH,0,2432,46.270,6.463,a
`,
		[]common.Waypoint{common.Waypoint{}},
		true,
	},
	{
		"query with updated time",
		[]string{"CH"},
		time.Date(2012, time.January, 25, 0, 0, 0, 0, time.UTC), // FIXME: figure out how to actually filter here
		`
ID,Name,Description,Region,Flags,Elevation,Latitude,Longitude
FURKAP,FURKAP,FURKAPASS PASSHOEHE,CH,0,2432,46.270,6.463
`,
		[]common.Waypoint{
			common.Waypoint{ID: "FURKAP", Name: "FURKAP", Description: "FURKAPASS PASSHOEHE",
				Region: "CH", Flags: 0, Elevation: 2432,
				Latitude: "46.270", Longitude: "6.463"},
		},
		false,
	},
}

func TestGetWaypoint(t *testing.T) {
	for _, at := range getGetWaypointTests {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, at.rp)
		}))
		defer ts.Close()

		cfg := config.Config{}
		cfg.FusionTables.BaseURL = ts.URL
		cfg.FusionTables.WaypointTableID = "testwaypointid"
		plugin := FusionTables{}
		err := plugin.Init(cfg)
		if err != nil {
			t.Errorf("Failed to initialize plugin :: %v", err)
		}
		waypoints, err := plugin.GetWaypoint(at.rg, at.tm)
		if err != nil && at.err {
			continue
		} else if err != nil {
			t.Errorf("Failed to get waypoints in test %v :: %v", at.t, err)
			continue
		}

		if len(waypoints) != len(at.rs) {
			t.Errorf("Expected %v but got %v waypoints", len(at.rs), len(waypoints))
		}
		for i, waypoint := range waypoints {
			if waypoint != at.rs[i] {
				t.Errorf("Expected %v but got %v", at.rs[i], waypoint)
			}
		}
	}
}

func TestGetWaypointWithMissingLocation(t *testing.T) {
	cfg := config.Config{}
	cfg.FusionTables.BaseURL = "http://doesnotexist"
	cfg.FusionTables.WaypointTableID = "testwaypointid"
	plugin := FusionTables{}
	err := plugin.Init(cfg)
	if err != nil {
		t.Errorf("Failed to initialize plugin :: %v", err)
	}
	_, err = plugin.GetWaypoint([]string{"CH"}, time.Time{})
	if err == nil {
		t.Errorf("expected error but was successful")
	}
}

func TestGetWaypointWithMalformedURL(t *testing.T) {
	cfg := config.Config{}
	cfg.FusionTables.BaseURL = "wrong%url"
	cfg.FusionTables.WaypointTableID = "testwaypointid"
	plugin := FusionTables{}
	err := plugin.Init(cfg)
	if err != nil {
		t.Errorf("Failed to initialize plugin :: %v", err)
	}
	_, err = plugin.GetWaypoint([]string{"CH"}, time.Time{})
	if err == nil {
		t.Errorf("expected error but was successful")
	}
}

type PutWaypointTest struct {
	t   string
	in  []common.Waypoint
	csv string
	err bool
}

var putWaypointTests = []PutWaypointTest{
	{
		"simple update",
		[]common.Waypoint{
			common.Waypoint{ID: "FURKAP", Name: "FURKAP", Description: "FURKAPASS PASSHOEHE",
				Region: "CH", Flags: 0, Elevation: 2432,
				Latitude: "46.270", Longitude: "6.463"},
		},
		`ID,Name,Description,Region,Flags,Elevation,Latitude,Longitude
FURKAP,FURKAP,FURKAPASS PASSHOEHE,CH,0,2432,46.270,6.463
`,
		false,
	},
	{
		"simple failure",
		[]common.Waypoint{
			common.Waypoint{ID: "FURKAP", Name: "FURKAP", Description: "FURKAPASS PASSHOEHE",
				Region: "CH", Flags: 0, Elevation: 2432,
				Latitude: "46.270", Longitude: "6.463"},
		},
		`ID,Name,Description,Region,Flags,Elevation,Latitude,Longitude
aFURKAP,FURKAP,FURKAPASS PASSHOEHE,CH,0,2432,46.270,6.463
`,
		true,
	},
}

func TestPutWaypoint(t *testing.T) {
	for _, test := range putWaypointTests {
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

		cfg := config.Config{}
		cfg.FusionTables.UploadURL = ts.URL
		cfg.FusionTables.WaypointTableID = "testwaypointid"
		plugin := FusionTables{}
		err := plugin.Init(cfg)
		if err != nil {
			t.Errorf("%v failed to initialize plugin :: %v", test.t, err)
		}
		err = plugin.PutWaypoint(test.in)
		if err != nil && test.err {
			continue
		} else if err != nil {
			t.Errorf("%v failed :: %v", test.t, err)
		}
	}
}

func TestPutWaypointWithMissingLocation(t *testing.T) {
	cfg := config.Config{}
	cfg.FusionTables.BaseURL = "http://thisurlreallydoesnotexist.pt"
	cfg.FusionTables.WaypointTableID = "testwaypointid"
	plugin := FusionTables{}
	err := plugin.Init(cfg)
	if err != nil {
		t.Errorf("Failed to initialize plugin :: %v", err)
	}
	err = plugin.PutWaypoint([]common.Waypoint{})
	if err == nil {
		t.Errorf("expected error but was successful")
	}
}

func TestPutWaypointWithMalformedURL(t *testing.T) {
	cfg := config.Config{}
	cfg.FusionTables.BaseURL = "wrong%url"
	cfg.FusionTables.WaypointTableID = "testwaypointid"
	plugin := FusionTables{}
	err := plugin.Init(cfg)
	if err != nil {
		t.Errorf("Failed to initialize plugin :: %v", err)
	}
	err = plugin.PutWaypoint([]common.Waypoint{})
	if err == nil {
		t.Errorf("expected error but was successful")
	}
}

func TestPutWaypointWithBadStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	cfg := config.Config{}
	cfg.FusionTables.BaseURL = ts.URL
	cfg.FusionTables.WaypointTableID = "testwaypointid"
	plugin := FusionTables{}
	err := plugin.Init(cfg)
	if err != nil {
		t.Errorf("failed to initialize plugin :: %v", err)
	}
	err = plugin.PutWaypoint([]common.Waypoint{
		common.Waypoint{ID: "aFURKAP", Name: "FURKAP", Description: "FURKAPASS PASSHOEHE",
			Region: "CH", Flags: 0, Elevation: 2432,
			Latitude: "46.270", Longitude: "6.463"},
	})
	if err == nil {
		t.Errorf("expected error but got success")
	}
}
