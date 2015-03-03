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

package igc

import (
	"io/ioutil"
	"reflect"
	"testing"
	"time"

	"github.com/rochaporto/ezgliding/flight"
)

type ParseTest struct {
	t string
	c string
	r flight.Flight
	e bool
}

var parseTests = []ParseTest{
	{
		"basic header test",
		`
AFLA001Some Additional Data
HFDTE010203
HFFXA500
HFPLTEZ PILOT
HFCM2EZ CREW
HFGTYEZ TYPE
HFGIDEZ ID
HFDTM100
HFRFWv 0.1
HFRHWv 0.2
HFFTYEZ RECORDER,001
HFGPSEZ GPS,002,12,5000
HFPRSEZ PRESSURE
HFCIDEZ COMPID
HFCCLEZ COMPCLASS
`,
		flight.Flight{
			Header: flight.Header{
				Manufacturer: "FLA", UniqueID: "001", AdditionalData: "Some Additional Data",
				Date:        time.Date(2003, time.February, 01, 0, 0, 0, 0, time.UTC),
				FixAccuracy: 500, Pilot: "EZ PILOT", Crew: "EZ CREW",
				GliderType: "EZ TYPE", GliderID: "EZ ID", GPSDatum: "100",
				FirmwareVersion: "v 0.1", HardwareVersion: "v 0.2",
				FlightRecorder: "EZ RECORDER,001", GPS: "EZ GPS,002,12,5000",
				PressureSensor: "EZ PRESSURE", CompetitionID: "EZ COMPID",
				CompetitionClass: "EZ COMPCLASS",
			},
			K:          map[time.Time]map[string]string{},
			Events:     map[time.Time]map[string]string{},
			Satellites: map[time.Time][]int{},
			Sources:    make(map[string]flight.Source),
		},
		false,
	},
	{"A record failure too short",
		"AFLA0", flight.Flight{}, true},
	{"H record failure too short",
		"HFFX", flight.Flight{}, true},
	{"H record failure bad date",
		"HFDTE330203", flight.Flight{}, true},
	{"H record failure date too short",
		"HFDTE33", flight.Flight{}, true},
	{"H record failure bad fix accuracy",
		"HFFXAAAA", flight.Flight{}, true},
	{"H record failure fix accuracy too short",
		"HFFXA20", flight.Flight{}, true},
	{"H record failure gps datum too short",
		"HFDTM20", flight.Flight{}, true},
	{"H record failure unknown field",
		"HFZZZaaa", flight.Flight{}, true},
	{
		"basic flight test",
		`
I033638FXA3940SIU4143ENL
J010812HDT
C150701213841160701000102500KTri
C5111359N00101899WEZ TAKEOFF
C5110179N00102644WEZ START
C5209092N00255227WEZ TP1
C5230147N00017612WEZ TP2
C5110179N00102644WEZ FINISH
C5111359N00101899WEZ LANDING
F160240040609123624
D20331
E160245ATS102312
B1602455107126N00149300WA002880042919509020
K16024800090
B1603105107212N00149174WV002930043519608024
LPLTLOG TEXT
GREJNGJERJKNJKRE31895478537H43982FJN9248F942389T433T
GJNJK2489IERGNV3089IVJE9GO398535J3894N358954983O0934
`,
		flight.Flight{
			Points: []flight.Point{
				flight.Point{
					Time:     time.Date(0, 1, 1, 16, 2, 45, 0, time.UTC),
					Latitude: 107.2, Longitude: 15.55, FixValidity: 'A',
					PressureAltitude: 288, GNSSAltitude: 429,
					IData: map[string]string{
						"FXA": "195", "SIU": "09", "ENL": "020",
					},
					NumSatellites: 6,
				},
				flight.Point{
					Time:     time.Date(0, 1, 1, 16, 3, 10, 0, time.UTC),
					Latitude: 107.35, Longitude: 15.516666666666666,
					FixValidity: 'V', PressureAltitude: 293, GNSSAltitude: 435,
					IData: map[string]string{
						"FXA": "196", "SIU": "08", "ENL": "024",
					},
					NumSatellites: 6,
				},
			},
			K: map[time.Time]map[string]string{
				time.Date(0, 1, 1, 16, 2, 48, 0, time.UTC): map[string]string{
					"HDT": "00090",
				},
			},
			Events: map[time.Time]map[string]string{
				time.Date(0, 1, 1, 16, 2, 45, 0, time.UTC): map[string]string{
					"ATS": "102312",
				},
			},
			Satellites: map[time.Time][]int{
				time.Date(0, 1, 1, 16, 02, 40, 0, time.UTC): []int{4, 6, 9, 12, 36, 24},
			},
			Logbook: []flight.LogEntry{
				flight.LogEntry{Type: "PLT", Text: "LOG TEXT"},
			},
			Task: flight.Task{
				DeclarationDate: time.Date(2001, time.July, 15, 21, 38, 41, 0, time.UTC),
				FlightDate:      time.Date(2001, time.July, 16, 0, 0, 0, 0, time.UTC),
				Number:          1,
				Takeoff: flight.Point{
					Latitude: 111.58333333333333, Longitude: 10.3,
					Description: "EZ TAKEOFF"},
				Start: flight.Point{
					Latitude: 110.28333333333333, Longitude: 10.433333333333334,
					Description: "EZ START"},
				Turnpoints: []flight.Point{
					flight.Point{
						Latitude: 209.15, Longitude: 25.866666666666667,
						Description: "EZ TP1"},
					flight.Point{
						Latitude: 230.23333333333332, Longitude: 2.2666666666666666,
						Description: "EZ TP2"},
				},
				Finish: flight.Point{
					Latitude: 110.28333333333333, Longitude: 10.433333333333334,
					Description: "EZ FINISH"},
				Landing: flight.Point{
					Latitude: 111.58333333333333, Longitude: 10.3,
					Description: "EZ LANDING"},
				Description: "500KTri",
			},
			DGPSStationID: "0331",
			Signature:     "REJNGJERJKNJKRE31895478537H43982FJN9248F942389T433TJNJK2489IERGNV3089IVJE9GO398535J3894N358954983O0934",
			Sources:       make(map[string]flight.Source),
		},
		false,
	},
	{"point/fix wrong size",
		"B110001", flight.Flight{}, true},
	{"point/fix bad time",
		"B3103105107212N00149174WV002930043519608024", flight.Flight{}, true},
	{"point/fix bad fix validity",
		"B1603105107212N00149174WX002930043519608024", flight.Flight{}, true},
	{"point/fix bad pressure altitude",
		"B1603105107212N00149174WV0029a0043519608024", flight.Flight{}, true},
	{"point/fix bad gnss altitude",
		"B1603105107212N00149174WV002930043a19608024", flight.Flight{}, true},
	{"irecord wrong size",
		"I0", flight.Flight{}, true},
	{"irecord invalid value for field number",
		"I0a", flight.Flight{}, true},
	{"irecord wrong size with fields",
		"I02AAA0102BBB030", flight.Flight{}, true},
	{"jrecord wrong size",
		"J0", flight.Flight{}, true},
	{"jrecord invalid value for field number",
		"J0a", flight.Flight{}, true},
	{"jrecord wrong size with fields",
		"J02AAA0102BBB030", flight.Flight{}, true},
	{"k wrong size",
		"K16024", flight.Flight{}, true},
	{"k invalid date",
		"K160271", flight.Flight{}, true},
	{"k wrong size",
		"K16027000090", flight.Flight{}, true},
	{"e wrong size",
		"E16024", flight.Flight{}, true},
	{"e invalid date",
		"E160271ATS", flight.Flight{}, true},
	{"f wrong size",
		"F16024", flight.Flight{}, true},
	{"f invalid date",
		"F1602710102", flight.Flight{}, true},
	{"f invalid num satellites",
		"F1602310a02", flight.Flight{}, true},
	{"l wrong size",
		"LPL", flight.Flight{}, true},
	{"c bad num lines",
		"C150701213841160701000102500KTri", flight.Flight{}, true},
	{"c wrong size first line",
		"C15070121384116070100010", flight.Flight{}, true},
	{"c invalid num of tps",
		"C15070121384116070100010a", flight.Flight{}, true},
	{"c invalid declaration date",
		"C350701213841160701000102500KTri\nC5111359N00101899WEZ TAKEOFF\nC5110179N00102644WEZ START\nC5209092N00255227WEZ TP1\nC5230147N00017612WEZ TP2\nC5110179N00102644WEZ FINISH\nC5111359N00101899WEZ LANDING", getFlight(flight.Task{
			DeclarationDate: time.Time{},
			FlightDate:      time.Date(2001, time.July, 16, 0, 0, 0, 0, time.UTC),
			Number:          1,
			Takeoff: flight.Point{
				Latitude: 111.58333333333333, Longitude: 10.3,
				Description: "EZ TAKEOFF"},
			Start: flight.Point{
				Latitude: 110.28333333333333, Longitude: 10.433333333333334,
				Description: "EZ START"},
			Turnpoints: []flight.Point{
				flight.Point{
					Latitude: 209.15, Longitude: 25.866666666666667,
					Description: "EZ TP1"},
				flight.Point{
					Latitude: 230.23333333333332, Longitude: 2.2666666666666666,
					Description: "EZ TP2"},
			},
			Finish: flight.Point{
				Latitude: 110.28333333333333, Longitude: 10.433333333333334,
				Description: "EZ FINISH"},
			Landing: flight.Point{
				Latitude: 111.58333333333333, Longitude: 10.3,
				Description: "EZ LANDING"},
			Description: "500KTri",
		}), false},
	{"c invalid flight date",
		"C150701213841360701000102500KTri\nC5111359N00101899WEZ TAKEOFF\nC5110179N00102644WEZ START\nC5209092N00255227WEZ TP1\nC5230147N00017612WEZ TP2\nC5110179N00102644WEZ FINISH\nC5111359N00101899WEZ LANDING", getFlight(flight.Task{
			DeclarationDate: time.Date(2001, time.July, 15, 21, 38, 41, 0, time.UTC),
			FlightDate:      time.Time{},
			Number:          1,
			Takeoff: flight.Point{
				Latitude: 111.58333333333333, Longitude: 10.3,
				Description: "EZ TAKEOFF"},
			Start: flight.Point{
				Latitude: 110.28333333333333, Longitude: 10.433333333333334,
				Description: "EZ START"},
			Turnpoints: []flight.Point{
				flight.Point{
					Latitude: 209.15, Longitude: 25.866666666666667,
					Description: "EZ TP1"},
				flight.Point{
					Latitude: 230.23333333333332, Longitude: 2.2666666666666666,
					Description: "EZ TP2"},
			},
			Finish: flight.Point{
				Latitude: 110.28333333333333, Longitude: 10.433333333333334,
				Description: "EZ FINISH"},
			Landing: flight.Point{
				Latitude: 111.58333333333333, Longitude: 10.3,
				Description: "EZ LANDING"},
			Description: "500KTri",
		}), false},
	{"c invalid task number",
		"C150701213841160701000a01500KTri\nC5111359N00101899WEZ TAKEOFF\nC5110179N00102644WEZ START\nC5209092N00255227WEZ TP1\nC5110179N00102644WEZ FINISH\nC5111359N00101899WEZ LANDING", flight.Flight{}, true},
	{"c invalid takeoff",
		"C150701213841160701000101500KTri\nC5111359N00101899\nC5110179N00102644WEZ START\nC5209092N00255227WEZ TP1\nC5110179N00102644WEZ FINISH\nC5111359N00101899WEZ LANDING", flight.Flight{}, true},
	{"c invalid start",
		"C150701213841160701000101500KTri\nC5111359N00101899WEZ TAKEOFF\nC5110179N00102644\nC5209092N00255227WEZ TP1\nC5110179N00102644WEZ FINISH\nC5111359N00101899WEZ LANDING", flight.Flight{}, true},
	{"c invalid tp",
		"C150701213841160701000101500KTri\nC5111359N00101899WEZ TAKEOFF\nC5110179N00102644WEZ START\nC5209092N00255227\nC5110179N00102644WEZ FINISH\nC5111359N00101899WEZ LANDING", flight.Flight{}, true},
	{"c invalid finish",
		"C150701213841160701000101500KTri\nC5111359N00101899WEZ TAKEOFF\nC5110179N00102644WEZ START\nC5209092N00255227WEZ TP1\nC5110179N00102644\nC5111359N00101899WEZ LANDING", flight.Flight{}, true},
	{"c invalid landing",
		"C150701213841160701000101500KTri\nC5111359N00101899WEZ TAKEOFF\nC5110179N00102644WEZ START\nC5209092N00255227WEZ TP1\nC5110179N00102644WEZ FINISH\nC5111359N00101899", flight.Flight{}, true},
	{"d wrong size",
		"D2033", flight.Flight{}, true},
	{"invalid record",
		"RANDOM GARBAGE", flight.Flight{}, true},
}

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		result, err := Parse(test.c)
		if err != nil && test.e {
			continue
		} else if err != nil {
			t.Errorf("%v failed :: %v", test.t, err)
			continue
		}
		if !reflect.DeepEqual(result, test.r) {
			t.Errorf("%v failed :: expected\n%+v\ngot\n%+v", test.t, test.r, result)
			continue
		}
	}
}

func BenchmarkParse(b *testing.B) {
	c, err := ioutil.ReadFile("t/sample-flight.igc")
	if err != nil {
		b.Errorf("failed to load sample flight :: %v", err)
	}
	content := string(c)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(content)
	}
}

func getFlight(task flight.Task) flight.Flight {
	f := flight.NewFlight()
	f.Task = task
	return f
}
