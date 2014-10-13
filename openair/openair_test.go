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

package openair

import (
	"github.com/rochaporto/ezgliding/common"
	"image/color"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type ParseTest struct {
	t string
	c string
	r []common.Airspace
}

var parseTests = []ParseTest{
	{"single airspace",
		`
AC A
AN TMA GENEVE partie  2 
AH FL 185
AL 5500FT AMSL
V D=+
V X=46:03:03 N 005:47:12 E
* comment
** comment

DB 45:55:41 N 005:54:39 E,46:10:24 N 005:39:42 E
DP 46:10:24 N 005:39:42 E
DC 1.35
DA 10,270,290`,
		[]common.Airspace{
			common.Airspace{
				Class: 'A', Name: "TMA GENEVE partie  2",
				Floor: "5500FT AMSL", Ceiling: "FL 185",
				Segments: []common.AirspaceSegment{
					common.AirspaceSegment{
						Type: common.Arc, X: "46:03:03 N 005:47:12 E", Clockwise: true,
						Coordinate1: "45:55:41 N 005:54:39 E",
						Coordinate2: "46:10:24 N 005:39:42 E",
					},
					common.AirspaceSegment{
						Type: common.Polygon, X: "46:03:03 N 005:47:12 E", Clockwise: true,
						Coordinate1: "46:10:24 N 005:39:42 E",
					},
					common.AirspaceSegment{
						Type: common.Circle, X: "46:03:03 N 005:47:12 E", Clockwise: true,
						Radius: 1.35,
					},
					common.AirspaceSegment{
						Type: common.Arc, X: "46:03:03 N 005:47:12 E", Clockwise: true,
						Radius: 10, AngleStart: 270, AngleEnd: 290,
					},
				},
			}},
	},
	{"multiple with empty airspace",
		`
AC A
AN TMA GENEVE partie  2 
AH FL 185
AL 5500FT AMSL
V D=+
V X=46:03:03 N 005:47:12 E
DB 45:55:41 N 005:54:39 E,46:10:24 N 005:39:42 E
DP 46:10:24 N 005:39:42 E
*`,
		[]common.Airspace{
			common.Airspace{
				Class: 'A', Name: "TMA GENEVE partie  2",
				Floor: "5500FT AMSL", Ceiling: "FL 185",
				Segments: []common.AirspaceSegment{
					common.AirspaceSegment{
						Type: common.Arc, X: "46:03:03 N 005:47:12 E", Clockwise: true,
						Coordinate1: "45:55:41 N 005:54:39 E",
						Coordinate2: "46:10:24 N 005:39:42 E",
					},
					common.AirspaceSegment{
						Type: common.Polygon, X: "46:03:03 N 005:47:12 E", Clockwise: true,
						Coordinate1: "46:10:24 N 005:39:42 E",
					},
				},
			}},
	},
	{"multiple airspaces",
		`
AC C
AN TMA GENEVE partie  1
AH FL 195
AL 3500FT AMSL
DP 46:22:03 N 006:33:04 E
*
** (04/04/2014) TMA GENEVE partie  2
AC A
AN TMA GENEVE partie  2
AH FL 185
AL 5500FT AMSL
V D=-
V X=46:03:03 N 005:47:12 E
DB 45:55:41 N 005:54:39 E,46:10:24 N 005:39:42 E
DP 46:10:24 N 005:39:42 E`,
		[]common.Airspace{
			common.Airspace{
				Class: 'C', Name: "TMA GENEVE partie  1",
				Floor: "3500FT AMSL", Ceiling: "FL 195",
				Segments: []common.AirspaceSegment{
					common.AirspaceSegment{
						Type: common.Polygon, Clockwise: false,
						Coordinate1: "46:22:03 N 006:33:04 E",
					},
				},
			},
			common.Airspace{
				Class: 'A', Name: "TMA GENEVE partie  2",
				Floor: "5500FT AMSL", Ceiling: "FL 185",
				Segments: []common.AirspaceSegment{
					common.AirspaceSegment{
						Type: common.Arc, X: "46:03:03 N 005:47:12 E", Clockwise: false,
						Coordinate1: "45:55:41 N 005:54:39 E",
						Coordinate2: "46:10:24 N 005:39:42 E",
					},
					common.AirspaceSegment{
						Type: common.Polygon, X: "46:03:03 N 005:47:12 E", Clockwise: false,
						Coordinate1: "46:10:24 N 005:39:42 E",
					},
				},
			}},
	},
	{"airspace with colors",
		`
AC C
SP 0,2,0,0,255
SB -1,-1,-1
*
AC C
AN TMA GENEVE partie  1
AH FL 195
AL 3500FT AMSL
DP 46:22:03 N 006:33:04 E
`,
		[]common.Airspace{
			common.Airspace{
				Class: 'C',
			},
			common.Airspace{
				Class: 'C', Name: "TMA GENEVE partie  1",
				Floor: "3500FT AMSL", Ceiling: "FL 195",
				Segments: []common.AirspaceSegment{
					common.AirspaceSegment{
						Type: common.Polygon, Clockwise: false,
						Coordinate1: "46:22:03 N 006:33:04 E",
					},
				},
				Pen: common.Pen{
					Style: common.Solid, Width: 2,
					Color:       color.RGBA64{R: 0, G: 0, B: 255, A: 1.0},
					InsideColor: color.RGBA64{},
				},
			},
		},
	},
}

func TestParse(t *testing.T) {
	for i := range parseTests {
		test := parseTests[i]
		airspace, err := Parse([]byte(test.c))
		if err != nil || len(airspace) != len(test.r) {
			t.Errorf("Failed to parse :: %v :: %v :: got %v vs %v expected",
				test.t, err, len(airspace), len(test.r))
		} else {
			for i := range airspace {
				if !reflect.DeepEqual(airspace[i], test.r[i]) {
					t.Errorf("Failed to parse :: %v\nr: %v\ne: %v",
						test.t, airspace[i], test.r[i])
				}
			}
		}
	}
}

func TestParseUnknownSegment(t *testing.T) {
	_, err := Parse([]byte("AC A\nNN 1.0"))
	if err == nil {
		t.Errorf("Parsing unknown segment should fail")
	}
}

func TestFetchLocal(t *testing.T) {
	airspace, err := Fetch("./test-airspace-basic.txt")
	if err != nil {
		t.Errorf("Failed to fetch airspace :: %v", err)
	}
	if len(airspace) < 1 {
		t.Errorf("Got wrong number of airspaces :: %v", len(airspace))
	}
}

func TestFetchHTTP(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadFile("./test-airspace-basic.txt")
		io.WriteString(w, string(resp))
	}))
	defer ts.Close()

	airspace, err := Fetch(ts.URL)
	if err != nil {
		t.Errorf("Failed to fetch airspaces from http endpoint :: %v", err)
	}
	if len(airspace) < 1 {
		t.Errorf("Got wrong number of airspaces :: %v", len(airspace))
	}
}

func TestFetchMissing(t *testing.T) {
	_, err := Fetch("./nonexisting.file")
	if err == nil {
		t.Errorf("Fetching non existing file should fail")
	}
}

func TestFetchEmpty(t *testing.T) {
	_, err := Fetch("")
	if err == nil {
		t.Errorf("Fetching empty location should fail")
	}
}

func TestStyleToAirspace(t *testing.T) {
	if styleToAirspace(0) != common.Solid {
		t.Errorf("Failed to return proper Solid pen style")
	}
	if styleToAirspace(1) != common.Dash {
		t.Errorf("Failed to return proper Dash pen style")
	}
	if styleToAirspace(5) != common.None {
		t.Errorf("Failed to return proper None pen style")
	}
	if styleToAirspace(-1) != common.None {
		t.Errorf("Failed to return proper unknown pen style")
	}
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Fetch("./test-airspace.txt")
		if err != nil {
			b.Errorf("Failed to fetch airspace :: %v", err)
		}
	}
}
