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
package util

import (
	"reflect"
	"testing"
	"time"

	"github.com/rochaporto/ezgliding/airfield"
)

type CSV2AirfieldTest struct {
	t   string
	in  string
	r   []airfield.Airfield
	err bool
}

var csv2AirfieldTests = []CSV2AirfieldTest{
	{
		"simple parse",
		`
ID,ShortName,Name,Region,ICAO,Flags,Catalog,Length,Elevation,Runway,Frequency,Latitude,Longitude,Update
HABER,HABER,HABERE POC69,FR,,1032,0,0,1113,0119,122.5,46.270,6.463,0001-01-01 00:00:00 +0000 UTC
`,
		[]airfield.Airfield{
			airfield.Airfield{ID: "HABER", ShortName: "HABER", Name: "HABERE POC69",
				Region: "FR", ICAO: "", Flags: 1032, Catalog: 0, Length: 0, Elevation: 1113,
				Runway: "0119", Frequency: 122.5, Latitude: 46.270, Longitude: 6.463,
				Update: time.Time{}},
		},
		false,
	},
	{
		"multiline parse",
		`
ID,ShortName,Name,Region,ICAO,Flags,Catalog,Length,Elevation,Runway,Frequency,Latitude,Longitude,Update
LSGG,GENEV,GENEVE COINTR,CH,LSGG,64,0,3880,430,0523,118.7,46.238,6.109,0001-01-01 00:00:00 +0000 UTC
LSGB,BEX,BEX,CH,LSGB,1024,0,0,399,1533,122.15,46.258,6.986,0001-01-01 00:00:00 +0000 UTC
`,
		[]airfield.Airfield{
			airfield.Airfield{ID: "LSGG", ShortName: "GENEV", Name: "GENEVE COINTR",
				Region: "CH", ICAO: "LSGG", Flags: 64, Catalog: 0, Length: 3880, Elevation: 430,
				Runway: "0523", Frequency: 118.7, Latitude: 46.238, Longitude: 6.109,
				Update: time.Time{}},
			airfield.Airfield{ID: "LSGB", ShortName: "BEX", Name: "BEX",
				Region: "CH", ICAO: "LSGB", Flags: 1024, Catalog: 0, Length: 0, Elevation: 399,
				Runway: "1533", Frequency: 122.15, Latitude: 46.258, Longitude: 6.986,
				Update: time.Time{}},
		},
		false,
	},
	{
		"parse with invalid csv format",
		`
ID,ShortName,Name,Region,ICAO,Flags,Catalog,Length,Elevation,Runway,Frequency,Latitude,Longitude,Update
HABER,HABER,HABERE POC69,FR,,1032,0,0,1113,0119,122.5,46.270,6.463,0001-01-01 00:00:00 +0000 UTC,a
`,
		[]airfield.Airfield{airfield.Airfield{}},
		true,
	},
	{
		"parse with invalid flags",
		`
ID,ShortName,Name,Region,ICAO,Flags,Catalog,Length,Elevation,Runway,Frequency,Latitude,Longitude,Update
HABER,HABER,HABERE POC69,FR,,badflags,0,0,1113,0119,122.5,46.270,6.463,0001-01-01 00:00:00 +0000 UTC
`,
		[]airfield.Airfield{airfield.Airfield{}},
		true,
	},
	{
		"parse with invalid catalog",
		`
ID,ShortName,Name,Region,ICAO,Flags,Catalog,Length,Elevation,Runway,Frequency,Latitude,Longitude,Update
HABER,HABER,HABERE POC69,FR,,1032,badcatalog,0,1113,0119,122.5,46.270,6.463,0001-01-01 00:00:00 +0000 UTC
`,
		[]airfield.Airfield{airfield.Airfield{}},
		true,
	},
	{
		"parse with invalid length",
		`
ID,ShortName,Name,Region,ICAO,Flags,Catalog,Length,Elevation,Runway,Frequency,Latitude,Longitude,Update
HABER,HABER,HABERE POC69,FR,,1032,0,badlength,1113,0119,122.5,46.270,6.463,0001-01-01 00:00:00 +0000 UTC
`,
		[]airfield.Airfield{airfield.Airfield{}},
		true,
	},
	{
		"parse with invalid elevation",
		`
ID,ShortName,Name,Region,ICAO,Flags,Catalog,Length,Elevation,Runway,Frequency,Latitude,Longitude,Update
HABER,HABER,HABERE POC69,FR,,1032,0,0,badelevation,0119,122.5,46.270,6.463,0001-01-01 00:00:00 +0000 UTC
`,
		[]airfield.Airfield{airfield.Airfield{}},
		true,
	},
	{
		"parse with invalid frequency",
		`
ID,ShortName,Name,Region,ICAO,Flags,Catalog,Length,Elevation,Runway,Frequency,Latitude,Longitude,Update
HABER,HABER,HABERE POC69,FR,,1032,0,0,1113,0119,badfrequency,46.270,6.463,0001-01-01 00:00:00 +0000 UTC
`,
		[]airfield.Airfield{airfield.Airfield{}},
		true,
	},
	{
		"parse with no records",
		"",
		[]airfield.Airfield{},
		true,
	},
}

func TestCSV2Airfield(t *testing.T) {
	for _, test := range csv2AirfieldTests {
		result, err := CSV2Struct(test.in, reflect.ValueOf([]airfield.Airfield{}).Type(),
			reflect.ValueOf(airfield.Airfield{}).Type())
		if err != nil && test.err {
			continue
		} else if err != nil {
			t.Errorf("failed to parse csv in test %v :: %v", test.t, err)
			continue
		}

		resultv := result.Interface().([]airfield.Airfield)
		if len(resultv) != len(test.r) {
			t.Errorf("%v :: expected %v but got %v airfields", test.t, len(test.r), len(resultv))
			continue
		}
		for i, airfield := range resultv {
			if airfield != test.r[i] {
				t.Errorf("%v :: expected %v but got %v", test.t, test.r[i], airfield)
				continue
			}
		}
	}
}

type Airfield2CSVTest struct {
	t   string
	in  []airfield.Airfield
	csv string
	err bool
}

var airfield2CSVTests = []Airfield2CSVTest{
	{
		"simple conversion",
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
		"conversion of empty array",
		[]airfield.Airfield{},
		``,
		false,
	},
	{
		"conversion with all empty",
		[]airfield.Airfield{
			airfield.Airfield{ID: "", ShortName: "", Name: "",
				Region: "", ICAO: "", Flags: 0, Catalog: 0, Length: 0, Elevation: 0,
				Runway: "", Frequency: 0, Latitude: 0.0, Longitude: 0.0, Update: time.Time{}},
		},
		`ID,ShortName,Name,Region,ICAO,Flags,Catalog,Length,Elevation,Runway,Frequency,Latitude,Longitude,Update
,,,,,0,0,0,0,,0,0,0,0001-01-01 00:00:00 +0000 UTC
`,
		false,
	},
}

func TestAirfield2CSV(t *testing.T) {
	for _, test := range airfield2CSVTests {
		result := Struct2CSV(test.in)
		if result != test.csv {
			t.Errorf("expected\n%v\ngot\n%v", test.csv, result)
			continue
		}
	}
}
