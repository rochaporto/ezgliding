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

package welt2000

import (
	"github.com/rochaporto/ezgliding/common"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListLocal(t *testing.T) {
	releases, err := List("./test-releases-list.xml")
	if err != nil {
		t.Errorf("Failed to list releases :: %v", err)
	}
	if len(releases) < 1 {
		t.Errorf("Got wrong number of releases :: %v", len(releases))
	}
}

func TestListHTTP(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadFile("./test-releases-list.xml")
		io.WriteString(w, string(resp))
	}))
	defer ts.Close()

	releases, err := List(ts.URL)
	if err != nil {
		t.Errorf("Failed to list releases from http endpoint :: %v", err)
	}
	if len(releases) < 1 {
		t.Errorf("Got wrong number of releases :: %v", len(releases))
	}
}

func TestListEmpty(t *testing.T) {
	_, err := List("")
	if err == nil {
		t.Errorf("List empty string should give error")
	}
}

func TestListMissing(t *testing.T) {
	_, err := List("./nonexisting.file")
	if err == nil {
		t.Errorf("List non existing file should give error")
	}
}

func TestListBrokenFeed(t *testing.T) {
	_, err := List("./test-brokenfeed.xml")
	if err == nil {
		t.Errorf("Parsing a broken rss feed should have failed")
	}
}

func TestFetchLocal(t *testing.T) {
	release, err := Fetch("./test-release-basic.txt")
	if err != nil {
		t.Errorf("Failed to fetch release from local :: %v", err)
	}
	if len(release.Airfields) < 1 || len(release.Waypoints) < 1 {
		t.Errorf("Got wrong number of airfields (%v) or waypoints (%v)", len(release.Airfields), len(release.Waypoints))
	}
}

func TestFetchHTTP(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadFile("./test-release-basic.txt")
		io.WriteString(w, string(resp))
	}))
	defer ts.Close()

	release, err := Fetch(ts.URL)
	if err != nil {
		t.Errorf("Failed to fetch release from http endpoint :: %v", err)
	}
	if len(release.Airfields) < 1 || len(release.Waypoints) < 1 {
		t.Errorf("Got wrong number of airfields or waypoints :: %v :: %v", len(release.Airfields), len(release.Waypoints))
	}
}

func TestFetchEmpty(t *testing.T) {
	_, err := Fetch("")
	if err == nil {
		t.Errorf("Fetching an empty string should return error")
	}
}

func TestFetchMissing(t *testing.T) {
	_, err := Fetch("nonexisting.release")
	if err == nil {
		t.Errorf("Fetching a non existing release should return error")
	}
}

func TestParseNil(t *testing.T) {
	r := Release{}
	err := r.Parse(nil)
	if err == nil {
		t.Errorf("Parsing a nil value should return error")
	}
}

func TestParseEmpty(t *testing.T) {
	r := Release{}
	err := r.Parse([]byte{})
	if err == nil {
		t.Errorf("Parsing an empty value should return an error")
	}
}

func TestParseComment(t *testing.T) {
	r := Release{}
	r.Parse([]byte("$ this is a comment line"))
	if len(r.Airfields) > 0 || len(r.Waypoints) > 0 {
		t.Errorf("Parsing a comment line should fill up airfields or waypoints")
	}
}

func TestParseAirfield(t *testing.T) {
	r := Release{}
	r.Parse([]byte("ANNEM1 ANNEMASSE       #LFLIA129123012587 494N461131E0061606FRQ0"))
	airfield := r.Airfields[0]

	expected := common.Airfield{ID: "LFLI", ShortName: "ANNEM",
		Name: "ANNEMASSE", ICAO: "LFLI", Flags: 0 | common.Asphalt,
		Length: 1290, Runway: "1230", Frequency: 125.87, Elevation: 494,
		Latitude: "N461131", Longitude: "E0061606"}
	if airfield != expected {
		t.Errorf("Failed to parse airfield :: %v :: %v", expected, airfield)
	}
}

func TestParseUnclearAirstrip(t *testing.T) {
	r := Release{}
	r.Parse([]byte("AMBL21 AMBLETEUSE AERO #   ?G       1      32N504901E0013658FRQ0"))
	airfield := r.Airfields[0]
	if airfield.Flags&common.UnclearAirstrip != 1 {
		t.Errorf("Parse failed for unclear airstrip")
	}
}

func TestParseGliderSite(t *testing.T) {
	r := Release{}
	// case GLD#
	r.Parse([]byte("CHALA1 CHALAIS      GLD#LFIHG 83072512350  88N451605E0000058FRQ0"))
	airfield := r.Airfields[0]
	if airfield.Flags&common.GliderSite == 0 {
		t.Errorf("Parse failed for glider site")
	}
	// case GLD#GLD
	r.Parse([]byte("HABER1 HABERE POC69 GLD#GLD!G 980119122501113N461611E0062748FRP3"))
	airfield = r.Airfields[0]
	if airfield.Flags&common.GliderSite == 0 {
		t.Errorf("Parse failed for glider site")
	}
}

func TestParseULMSite(t *testing.T) {
	r := Release{}
	r.Parse([]byte("CERVE2 CERVENS UL      *ULM!G 28052312350 619N461713E0062638FRQ0"))
	airfield := r.Airfields[0]
	if airfield.Flags&common.ULMSite == 0 {
		t.Errorf("Parse failed for ulm site")
	}
}
func TestParseAsphalt(t *testing.T) {
	r := Release{}
	r.Parse([]byte("ANNEM1 ANNEMASSE       #LFLIA129123012587 494N461131E0061606FRQ0"))
	airfield := r.Airfields[0]
	if airfield.Flags&common.Asphalt == 0 {
		t.Errorf("Parse failed for asphalt airstrip")
	}
}

func TestParseConcrete(t *testing.T) {
	r := Release{}
	r.Parse([]byte("ANNEM1 ANNEMASSE       #LFLIC129123012587 494N461131E0061606FRQ0"))
	airfield := r.Airfields[0]
	if airfield.Flags&common.Concrete == 0 {
		t.Errorf("Parse failed for concrete airstrip")
	}
}

func TestParseLoam(t *testing.T) {
	r := Release{}
	r.Parse([]byte("ANNEM1 ANNEMASSE       #LFLIL129123012587 494N461131E0061606FRQ0"))
	airfield := r.Airfields[0]
	if airfield.Flags&common.Loam == 0 {
		t.Errorf("Parse failed for loam airstrip")
	}
}

func TestParseSand(t *testing.T) {
	r := Release{}
	r.Parse([]byte("ANNEM1 ANNEMASSE       #LFLIS129123012587 494N461131E0061606FRQ0"))
	airfield := r.Airfields[0]
	if airfield.Flags&common.Sand == 0 {
		t.Errorf("Parse failed for sand airstrip")
	}
}

func TestParseClay(t *testing.T) {
	r := Release{}
	r.Parse([]byte("ANNEM1 ANNEMASSE       #LFLIY129123012587 494N461131E0061606FRQ0"))
	airfield := r.Airfields[0]
	if airfield.Flags&common.Clay == 0 {
		t.Errorf("Parse failed for asphalt airstrip")
	}
}

func TestParseGrass(t *testing.T) {
	r := Release{}
	r.Parse([]byte("ANNEM1 ANNEMASSE       #LFLIG129123012587 494N461131E0061606FRQ0"))
	airfield := r.Airfields[0]
	if airfield.Flags&common.Grass == 0 {
		t.Errorf("Parse failed for grass airstrip")
	}
}

func TestParseGravel(t *testing.T) {
	r := Release{}
	r.Parse([]byte("ANNEM1 ANNEMASSE       #LFLIV129123012587 494N461131E0061606FRQ0"))
	airfield := r.Airfields[0]
	if airfield.Flags&common.Gravel == 0 {
		t.Errorf("Parse failed for gravel airstrip")
	}
}

func TestParseDirt(t *testing.T) {
	r := Release{}
	r.Parse([]byte("ANNEM1 ANNEMASSE       #LFLID129123012587 494N461131E0061606FRQ0"))
	airfield := r.Airfields[0]
	if airfield.Flags&common.Dirt == 0 {
		t.Errorf("Parse failed for dirt airstrip")
	}
}

func TestParseUnknownRunwayType(t *testing.T) {
	r := Release{}
	r.Parse([]byte("ANNEM1 ANNEMASSE       #LFLIO129123012587 494N461131E0061606FRQ0"))
	airfield := r.Airfields[0]
	if airfield.Flags != 0 {
		t.Errorf("Parse failed for invalid runway type")
	}
}

func TestParseCatalogNumber(t *testing.T) {
	r := Release{}
	r.Parse([]byte("BONVI2 BONNEVILLE      *FL53S 400523      450N460441E0062310FRP0"))
	airfield := r.Airfields[0]
	if airfield.Catalog != 53 || airfield.Flags&common.Outlanding == 0 {
		t.Errorf("Parse failed for outlanding catalog number")
	}
}

func TestParseWaypoint(t *testing.T) {
	r := Release{}
	r.Parse([]byte("FURKAP FURKAPASS PASSHOEHE               2432N463422E0082455CHQ6"))
	waypoint := r.Waypoints[0]
	expected := common.Waypoint{
		Name: "FURKAP", ID: "FURKAP", Description: "FURKAPASS PASSHOEHE",
		Latitude: "N463422", Longitude: "E0082455", Elevation: 2432,
	}
	if waypoint != expected {
		t.Errorf("Parse failed for waypoint: got %v instead of %v", waypoint, expected)
	}
}

func BenchmarkFetch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Fetch("./test-release-bench.txt")
		if err != nil {
			b.Errorf("Failed to fetch release :: %v", err)
		}
	}
}
