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

package soaringweb

import (
	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/config"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

type ParseTest struct {
	t  string
	rg string
	c  string
	r  []Release
}

var parseTests = []ParseTest{
	{"single entry",
		"FR",
		`
<html>
<body>
<small>[ Version 2014-05e (140708): Effective 01 May 2014 ]</small>
<small>[ 13 July 2014 ]</small>
<li>
<a href="http://soaringweb.org/Airspace/FR/140708__AIRSPACE_France_1405e.txt">OpenAir format</a>
</li>
</body>
</html>
`,
		[]Release{
			Release{Location: "http://soaringweb.org/Airspace/FR/140708__AIRSPACE_France_1405e.txt",
				Region: "FR", Date: time.Date(2014, time.July, 13, 0, 0, 0, 0, time.UTC)},
		},
	},
	{"multiple entries",
		"FR",
		`
<html>
<body>
<small>[ Version 2014-05e (140708): Effective 01 May 2014 ]</small>
<small>[ 13 July 2014 ]</small>
<li>
<a href="http://soaringweb.org/Airspace/FR/140708__AIRSPACE_France_1405e.txt">OpenAir format</a>
</li>
<small>[ 25 January 2012 ]</small>
<li>
<a href="http://soaringweb.org/Airspace/FR/250112__AIRSPACE_France_2501e.txt">OpenAir format</a>
</li>
</body>
</html>
`,
		[]Release{
			Release{Location: "http://soaringweb.org/Airspace/FR/140708__AIRSPACE_France_1405e.txt",
				Region: "FR", Date: time.Date(2014, time.July, 13, 0, 0, 0, 0, time.UTC)},
			Release{Location: "http://soaringweb.org/Airspace/FR/250112__AIRSPACE_France_2501e.txt",
				Region: "FR", Date: time.Date(2012, time.January, 25, 0, 0, 0, 0, time.UTC)},
		},
	},
	{"single entry inverted date",
		"FR",
		`
<html>
<body>
<small>[ Version 2014-05e (140708): Effective 01 May 2014 ]</small>
<li>
<a href="http://soaringweb.org/Airspace/FR/140708__AIRSPACE_France_1405e.txt">OpenAir format</a>
<small>[ 13 July 2014 ]</small>
</li>
</body>
</html>
`,
		[]Release{
			Release{Location: "http://soaringweb.org/Airspace/FR/140708__AIRSPACE_France_1405e.txt",
				Region: "FR", Date: time.Date(2014, time.July, 13, 0, 0, 0, 0, time.UTC)},
		},
	},
}

func TestListLocal(t *testing.T) {
	plugin := SoaringWeb{}
	releases, err := plugin.list("./t", []string{"FR"})
	if err != nil {
		t.Errorf("Failed to list releases :: %v", err)
	}
	if len(releases) < 1 {
		t.Errorf("Got wrong number of releases :: %v", len(releases))
	}
}

func TestListHTTP(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadFile("./t/FR")
		io.WriteString(w, string(resp))
	}))
	defer ts.Close()

	plugin := SoaringWeb{}
	releases, err := plugin.list(ts.URL, []string{"FR"})
	if err != nil {
		t.Errorf("Failed to list releases :: %v", err)
	}
	if len(releases) < 1 {
		t.Errorf("Got wrong number of releases :: %v", len(releases))
	}
}

func TestListEmpty(t *testing.T) {
	plugin := SoaringWeb{}
	_, err := plugin.list("", nil)
	if err == nil {
		t.Errorf("List empty string should give error")
	}
}

func TestListMissing(t *testing.T) {
	plugin := SoaringWeb{}
	_, err := plugin.list("./nonexisting.file", nil)
	if err == nil {
		t.Errorf("List non existing should give error")
	}
}

func TestParse(t *testing.T) {
	for i := range parseTests {
		test := parseTests[i]

		plugin := SoaringWeb{}
		releases, err := plugin.parse("./", test.rg, []byte(test.c))
		if err != nil {
			t.Errorf("Failed to parse '%v' :: %v", test.t, err)
		}
		if len(releases) != len(test.r) {
			t.Errorf("Wrong num of releases in '%v' :: %v but expected %v",
				test.t, len(releases), len(test.r))
		}
		for r := range releases {
			var release = releases[r]
			var expected = test.r[r]
			if release.Date != expected.Date {
				t.Errorf("Wrong date in release :: got %v expected %v", release, expected)
			}
			if release.Location != expected.Location {
				t.Errorf("Wrong location in release :: got %v expected %v", release, expected)
			}
			if release.Region != expected.Region {
				t.Errorf("Wrong region in release :: got %v expected %v", release, expected)
			}
		}
	}
}

func TestInit(t *testing.T) {
	cfg := config.Config{}
	cfg.SoaringWeb.Baseurl = "some.random/location"
	plugin := SoaringWeb{}
	err := plugin.Init(cfg)
	if err != nil {
		t.Errorf("Failed to initialize plugin :: %v", err)
	}

	if plugin.BaseURL != cfg.SoaringWeb.Baseurl {
		t.Errorf("Expected baseurl '%v' but got '%v'", cfg.SoaringWeb.Baseurl, plugin.BaseURL)
	}
}

func TestInitDefault(t *testing.T) {
	plugin := SoaringWeb{}
	err := plugin.Init(config.Config{})
	if err != nil {
		t.Errorf("Failed to initialize plugin :: %v", err)
	}

	if plugin.BaseURL != baseURL {
		t.Errorf("Expected baseurl '%v' but got '%v'", baseURL, plugin.BaseURL)
	}
}

type GetAirspaceTest struct {
	t  string
	b  string
	rg string
	d  time.Time
	r  []common.Airspace
}

var getAirspaceTests = []GetAirspaceTest{
	{"basic get airspace",
		"./t",
		"FR",
		time.Time{},
		[]common.Airspace{
			common.Airspace{
				Class: 'C', Name: "CTR Annecy 118.2",
				Floor: "SFC", Ceiling: "3500FT AMSL",
				Segments: []common.AirspaceSegment{
					common.AirspaceSegment{
						Type: common.Polygon, Coordinate1: "46:02:56 N 006:09:33 E",
					},
					common.AirspaceSegment{
						Type: common.Polygon, Coordinate1: "45:59:06 N 006:14:32 E",
					},
					common.AirspaceSegment{
						Type: common.Polygon, Coordinate1: "45:48:36 N 006:02:30 E",
					},
					common.AirspaceSegment{
						Type: common.Arc, X: "45:55:40 N 006:05:41 E", Clockwise: false,
						Coordinate1: "45:48:36 N 006:02:30 E", Coordinate2: "45:55:57 N 005:55:05 E",
					},
					common.AirspaceSegment{
						Type: common.Polygon, X: "45:55:40 N 006:05:41 E",
						Coordinate1: "45:55:57 N 005:55:05 E",
					},
				},
			},
			common.Airspace{
				Class: 'C', Name: "Geneve9 C 126.35",
				Floor: "FL115", Ceiling: "FL195",
				Segments: []common.AirspaceSegment{
					common.AirspaceSegment{
						Type: common.Polygon, Clockwise: false,
						Coordinate1: "45:52:25 N 006:07:45 E",
					},
					common.AirspaceSegment{
						Type: common.Polygon, Clockwise: false,
						Coordinate1: "45:50:38 N 006:06:05 E",
					},
				},
			},
		},
	},
	{"multiple airspaces with updated since",
		"./t",
		"MP",
		time.Date(2014, time.August, 25, 0, 0, 0, 0, time.UTC),
		[]common.Airspace{
			common.Airspace{
				Class: 'C', Name: "CTR Chambery2 118.3",
				Floor: "1160FT AMSL", Ceiling: "3500FT AMSL",
				Segments: []common.AirspaceSegment{
					common.AirspaceSegment{
						Type: common.Polygon, Coordinate1: "45:39:35 N 005:55:48 E",
					},
					common.AirspaceSegment{
						Type: common.Polygon, Coordinate1: "45:36:28 N 005:56:03 E",
					},
				},
			},
		},
	},
	{"with base url dependent location",
		"./t/wb",
		"PT",
		time.Time{},
		[]common.Airspace{
			common.Airspace{
				Class: 'C', Name: "CTR Chambery2 118.3",
				Floor: "1160FT AMSL", Ceiling: "3500FT AMSL",
				Segments: []common.AirspaceSegment{
					common.AirspaceSegment{
						Type: common.Polygon, Coordinate1: "45:39:35 N 005:55:48 E",
					},
					common.AirspaceSegment{
						Type: common.Polygon, Coordinate1: "45:36:28 N 005:56:03 E",
					},
				},
			},
		},
	},
}

func TestGetAirspace(t *testing.T) {
	for i := range getAirspaceTests {
		test := getAirspaceTests[i]

		plugin := SoaringWeb{}
		cfg := config.Config{}
		cfg.SoaringWeb.Baseurl = test.b
		err := plugin.Init(cfg)
		if err != nil {
			t.Errorf("Failed to initialize plugin :: %v", err)
		}

		var airspaces []common.Airspace
		airspaces, err = plugin.GetAirspace([]string{test.rg}, test.d)
		if err != nil {
			t.Errorf("Failed to get airspace :: %v", err)
		}

		if len(airspaces) != len(test.r) {
			t.Errorf("Got %v airspaces but expected %v in test '%v'", len(airspaces), len(test.r), test.t)
		}

		for i := range airspaces {
			var airspace = airspaces[i]
			var expected = test.r[i]
			airspace.Pen = common.Pen{}
			if !reflect.DeepEqual(airspace, expected) {
				t.Errorf("Got wrong airspace. %v instead of %v", airspace, expected)
			}
		}
	}
}

func TestGetAirspaceEmptyRegion(t *testing.T) {
	plugin := SoaringWeb{}
	cfg := config.Config{}
	cfg.SoaringWeb.Baseurl = "./t"
	err := plugin.Init(cfg)
	if err != nil {
		t.Errorf("Failed to initialize plugin :: %v", err)
	}

	var airspaces []common.Airspace
	airspaces, err = plugin.GetAirspace([]string{}, time.Time{})
	if err != nil {
		t.Errorf("Got error when retrieving airspace with empty regions :: %v", err)
	}

	if len(airspaces) != 0 {
		t.Errorf("Passing empty regions should return 0 airspaces, got %v", len(airspaces))
	}
}

func TestGetAirspaceMissingRegion(t *testing.T) {
	plugin := SoaringWeb{}
	cfg := config.Config{}
	cfg.SoaringWeb.Baseurl = "./t"
	err := plugin.Init(cfg)
	if err != nil {
		t.Errorf("Failed to initialize plugin :: %v", err)
	}

	_, err = plugin.GetAirspace([]string{"II"}, time.Time{})
	if err == nil {
		t.Errorf("Get airspace with missing region did not return error")
	}
}

func TestGetAirspaceMissingLocation(t *testing.T) {
	plugin := SoaringWeb{}
	cfg := config.Config{}
	cfg.SoaringWeb.Baseurl = "./t"
	err := plugin.Init(cfg)
	if err != nil {
		t.Errorf("Failed to initialize plugin :: %v", err)
	}

	_, err = plugin.GetAirspace([]string{"MS"}, time.Time{})
	if err == nil {
		t.Errorf("Get airspace with missing/bad location did not return error")
	}
}

func TestGetAirspaceMissingLocationWithBaseURL(t *testing.T) {
	plugin := SoaringWeb{}
	cfg := config.Config{}
	cfg.SoaringWeb.Baseurl = "./t/wb"
	err := plugin.Init(cfg)
	if err != nil {
		t.Errorf("Failed to initialize plugin :: %v", err)
	}

	_, err = plugin.GetAirspace([]string{"MS"}, time.Time{})
	if err == nil {
		t.Errorf("Get airspace with missing/bad location and base url did not return error")
	}
}

func TestPutAirspace(t *testing.T) {
	plugin := SoaringWeb{}
	cfg := config.Config{}
	cfg.SoaringWeb.Baseurl = "./t/wb"
	err := plugin.Init(cfg)
	if err != nil {
		t.Errorf("Failed to initialize plugin :: %v", err)
	}

	err = plugin.PutAirspace(nil) // FIXME: implement
	if err != nil {
		t.Errorf("Failed to put airspace")
	}
}
