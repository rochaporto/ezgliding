// Copyright 2015 The ezgliding Authors.
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

package netcoupe

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/rochaporto/ezgliding/flight"
)

func TestNew(t *testing.T) {
	cfg := Config{}
	cfg.BaseURL = "random/base/url"
	cfg.FlightDetailURL = "random/flight/detail/url"
	cfg.MaxIDGap = 5
	nc, err := New(cfg)
	if err != nil {
		t.Errorf("failed to get new nc :: %v", err)
		return
	}
	if nc.BaseURL != cfg.BaseURL {
		t.Errorf("expected baseurl %v but got %v", cfg.BaseURL, nc.BaseURL)
	}
	if nc.FlightDetailURL != cfg.FlightDetailURL {
		t.Errorf("expected flightdetailurl %v but got %v",
			cfg.FlightDetailURL, nc.FlightDetailURL)
	}
	if nc.MaxIDGap != cfg.MaxIDGap {
		t.Errorf("expected maxidgap %v but got %v", cfg.MaxIDGap, nc.MaxIDGap)
	}
}

func TestNewDefault(t *testing.T) {
	nc, err := New(Config{})
	if err != nil {
		t.Errorf("failed to get new nc:: %v", err)
		return
	}
	if nc.BaseURL != baseURL {
		t.Errorf("expected baseurl %v but got %v", baseURL, nc.BaseURL)
	}
	if nc.FlightDetailURL != flightDetailURL {
		t.Errorf("expected baseurl %v but got %v", flightDetailURL, nc.FlightDetailURL)
	}
	if nc.MaxIDGap != maxIDGap {
		t.Errorf("expected baseurl %v but got %v", maxIDGap, nc.MaxIDGap)
	}
}

type GetFlightByIDTest struct {
	t   string
	id  int
	r   flight.Source
	err bool
}

var getFlightByIDTests = []GetFlightByIDTest{
	GetFlightByIDTest{t: "basic get flight by id", id: 1, r: getSource(1), err: false},
	GetFlightByIDTest{t: "non existing get flight by id", id: 999, r: flight.Source{}, err: true},
	GetFlightByIDTest{t: "bad date get flight by id", id: 300, r: flight.Source{}, err: true},
	GetFlightByIDTest{t: "bad distance get flight by id", id: 301, r: flight.Source{}, err: true},
	GetFlightByIDTest{t: "bad points get flight by id", id: 302, r: flight.Source{}, err: true},
	GetFlightByIDTest{t: "bad speed get flight by id", id: 303, r: flight.Source{}, err: true},
	GetFlightByIDTest{t: "malformed flight get flight by id", id: 305, r: flight.Source{}, err: true},
	GetFlightByIDTest{t: "bad location get flight by id", id: 306, r: flight.Source{}, err: true},
}

func TestGetFlightByID(t *testing.T) {
	cfg := Config{}
	cfg.BaseURL = "./t/"
	cfg.FlightDetailURL = "Results/FlightDetail.aspx?FlightID="
	nc, err := New(cfg)
	if err != nil {
		t.Errorf("failed to get new nc :: %v", err)
		return
	}
	for _, test := range getFlightByIDTests {
		var flight flight.Flight
		var err error
		flight, err = nc.GetFlightByID(test.id)
		if err != nil && test.err {
			continue
		} else if err != nil {
			t.Errorf("failed to get flight by id :: %v :: %v", test.t, err)
			continue
		}
		result := flight.Sources[ID]
		if !reflect.DeepEqual(result, test.r) {
			t.Errorf("%v :: expected\n%v but got\n%v", test.t, test.r, result)
			continue
		}
	}
}

func TestGetFlightByIDBadRegexp(t *testing.T) {
	cfg := Config{}
	cfg.BaseURL = "./t/"
	cfg.FlightDetailURL = "Results/FlightDetail.aspx?FlightID="
	nc, err := New(cfg)
	if err != nil {
		t.Errorf("failed to get new nc :: %v", err)
		return
	}
	flight, err := nc.GetFlightByID(304)
	if err != nil {
		t.Errorf("failed to get flight by id bad regexp :: %v", err)
		return
	}
	expected := getSource(304)
	expected.Comment = "UNKNOWN"
	result := flight.Sources[ID]
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected\n%v but got\n%v", expected, result)
		return
	}
}

func TestGetFlightByIDHTTP(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadFile("./t/Results/FlightDetail.aspx?FlightID=1")
		io.WriteString(w, string(resp))
	}))
	defer ts.Close()

	nc, err := New(Config{})
	if err != nil {
		t.Errorf("failed to get new nc :: %v", err)
		return
	}
	content, err := nc.fetch(ts.URL)
	if err != nil {
		t.Errorf("failed to get flight by id http :: %v", err)
		return
	}
	source, err := nc.parseDetails(content)
	if err != nil {
		t.Errorf("failed to parse details :: %v", err)
		return
	}
	expected := getSource(1)
	if !reflect.DeepEqual(source, expected) {
		t.Errorf("expected %v but got %v", expected, source)
		return
	}
}

type GetFlightFromIDTest struct {
	t   string
	sid int
	max int
	r   []flight.Source
	err bool
}

var getFlightFromIDTests = []GetFlightFromIDTest{
	GetFlightFromIDTest{
		t: "basic get flight from id", sid: 2, max: -1, r: []flight.Source{getSource(2), getSource(5)}, err: false,
	},
	GetFlightFromIDTest{
		t: "get flight from id with max", sid: 2, max: 1, r: []flight.Source{getSource(2)}, err: false,
	},
	GetFlightFromIDTest{
		t: "get flight from id with max 0", sid: 2, max: 0, r: []flight.Source{}, err: false,
	},
}

func TestGetFlightFromID(t *testing.T) {
	cfg := Config{}
	cfg.BaseURL = "./t/"
	cfg.FlightDetailURL = "Results/FlightDetail.aspx?FlightID="
	nc, err := New(cfg)
	if err != nil {
		t.Errorf("failed to get new nc :: %v", err)
		return
	}
	for _, test := range getFlightFromIDTests {
		flights, err := nc.GetFlightFromID(test.sid, test.max)
		if err != nil && test.err {
			continue
		} else if err != nil {
			t.Errorf("failed to get flight from id :: %v", err)
			return
		}
		sources := []flight.Source{}
		for _, flight := range flights {
			sources = append(sources, flight.Sources[ID])
		}
		if !reflect.DeepEqual(sources, test.r) {
			t.Errorf("expected\n%v but got\n%v", test.r, sources)
			return
		}
	}
}
func TestGetFlightNotImplemented(t *testing.T) {
	nc, err := New(Config{})
	if err != nil {
		t.Errorf("init failed :: %v", err)
		return
	}
	if _, err := nc.GetFlight(nil, time.Time{}); err == nil {
		t.Errorf("expected error but got success")
		return
	}

}

func TestPutFlightNotImplemented(t *testing.T) {
	nc, err := New(Config{})
	if err != nil {
		t.Errorf("init failed :: %v", err)
		return
	}
	if err := nc.PutFlight([]flight.Flight{}); err == nil {
		t.Errorf("expected error but got success")
		return
	}

}

func getSource(id int) flight.Source {
	sid := strconv.Itoa(id)
	return flight.Source{
		Name: "PILOT " + sid, Category: "CATEGORY " + sid, Club: "CLUB " + sid,
		Region: "REGION " + sid, Country: "COUNTRY " + sid,
		Date:    time.Date(2015, time.Month(id), id, 0, 0, 0, 0, time.UTC),
		Takeoff: "TAKEOFF " + sid, Distance: 100.10 + float64(id), Points: 100.20 + float64(id),
		Type: "TYPE " + sid, CircuitType: "CIRCUIT TYPE " + sid,
		Speed: 100.30 + float64(id), Start: "START " + sid, Finish: "FINISH " + sid,
		Turnpoints: []flight.Point{
			flight.Point{Description: "POINT1 " + sid},
			flight.Point{Description: "POINT2 " + sid},
			flight.Point{Description: "POINT3 " + sid},
		},
		Comment:     "COMMENTS " + sid,
		DownloadURL: "/Results/DownloadIGC.aspx?FileID=" + sid,
	}
}
