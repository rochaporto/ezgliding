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

	"github.com/rochaporto/ezgliding/common"
)

type DMS2DecimalTest struct {
	t  string
	in string
	r  float64
}

var dms2DecimalTests = []DMS2DecimalTest{
	{
		"latitude north conversion",
		"N323200",
		32.53333333333333,
	},
	{
		"latitude south conversion",
		"S323200",
		-32.53333333333333,
	},
	{
		"longitude east conversion",
		"E1002233",
		100.37583333333333,
	},
	{
		"longitude west conversion",
		"W1002233",
		-100.37583333333333,
	},
}

func TestDMS2Decimal(t *testing.T) {
	for _, test := range dms2DecimalTests {
		result := DMS2Decimal(test.in)
		if result != test.r {
			t.Errorf("test %v failed, expected %v got %v", test.t, test.r, result)
			continue
		}
	}
}

type Struct2GeoJSONTest struct {
	t  string
	in []interface{}
	r  string
}

var struct2GeoJSONTests = []Struct2GeoJSONTest{
	Struct2GeoJSONTest{
		"simple airfield conversion",
		[]interface{}{
			common.Airfield{ID: "HABER", ShortName: "HABER", Name: "HABERE POC69",
				Region: "FR", ICAO: "HHHH", Flags: 1032, Catalog: 69, Length: 900, Elevation: 1113,
				Runway: "0119", Frequency: 122.5, Latitude: 46.270, Longitude: 6.463},
		},
		`{"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[46.27,6.463]},"properties":{"Catalog":69,"Elevation":1113,"Flags":1032,"Frequency":122.5,"Go":"Airfield","ICAO":"HHHH","ID":"HABER","Length":900,"Name":"HABERE POC69","Region":"FR","Runway":"0119","ShortName":"HABER"}}]}`,
	},
	Struct2GeoJSONTest{
		"simple waypoint conversion",
		[]interface{}{
			common.Waypoint{
				ID: "FURKAP", Name: "FURKAP", Description: "FURKAPASS PASSHOEHE", Elevation: 2432,
				Latitude: 46.572, Longitude: 8.415, Region: "CH", Flags: 0,
			},
		},
		`{"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[46.572,8.415]},"properties":{"Description":"FURKAPASS PASSHOEHE","Elevation":2432,"Flags":0,"Go":"Waypoint","ID":"FURKAP","Name":"FURKAP","Region":"CH"}}]}`,
	},
	Struct2GeoJSONTest{
		"multiple type conversion",
		[]interface{}{
			common.Airfield{ID: "HABER", ShortName: "HABER", Name: "HABERE POC69",
				Region: "FR", ICAO: "HHHH", Flags: 1032, Catalog: 69, Length: 900, Elevation: 1113,
				Runway: "0119", Frequency: 122.5, Latitude: 46.270, Longitude: 6.463},
			common.Waypoint{
				ID: "FURKAP", Name: "FURKAP", Description: "FURKAPASS PASSHOEHE", Elevation: 2432,
				Latitude: 46.572, Longitude: 8.415, Region: "CH", Flags: 0,
			},
		},
		`{"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[46.27,6.463]},"properties":{"Catalog":69,"Elevation":1113,"Flags":1032,"Frequency":122.5,"Go":"Airfield","ICAO":"HHHH","ID":"HABER","Length":900,"Name":"HABERE POC69","Region":"FR","Runway":"0119","ShortName":"HABER"}},{"type":"Feature","geometry":{"type":"Point","coordinates":[46.572,8.415]},"properties":{"Description":"FURKAPASS PASSHOEHE","Elevation":2432,"Flags":0,"Go":"Waypoint","ID":"FURKAP","Name":"FURKAP","Region":"CH"}}]}`,
	},
}

func TestStruct2GeoJSON(t *testing.T) {
	for _, test := range struct2GeoJSONTests {
		result, err := Struct2GeoJSON(test.in)
		if err != nil {
			t.Errorf("%v failed :: %v", test.t)
			continue
		}
		geo, err := result.MarshalJSON()
		if err != nil {
			t.Errorf("%v failed :: %v", test.t)
			continue
		}
		if string(geo) != test.r {
			t.Errorf("%v failed :: expected\n%v\ngot\n%v", test.t, test.r, string(geo))
			continue
		}
	}
}

func TestStruct2GeoJSOUnsupported(t *testing.T) {
	_, err := Struct2GeoJSON([]interface{}{
		string("random type"),
	})
	if err == nil {
		t.Errorf("expected error got success")
	}
}

func TestGeoJSON2Struct(t *testing.T) {
	for _, test := range struct2GeoJSONTests {
		result, err := GeoJSON2Struct(test.r)
		if err != nil {
			t.Errorf("%v failed :: %v", test.t)
			continue
		}
		if len(result) != len(test.in) {
			t.Errorf("%v failed :: expected %v got %v airfields", test.t, len(test.in), len(result))
			continue
		}
		if !reflect.DeepEqual(result, test.in) {
			t.Errorf("%v failed :: expected %v got %v", test.t, test.in, result)
			continue
		}
	}
}

func TestGeoJSON2StructUnsupported(t *testing.T) {
	_, err := GeoJSON2Struct(`{"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[46.572,8.415]},"properties":{"Go":"UnsupportedType"}}]}`)
	if err == nil {
		t.Errorf("expected error got success")
	}
}

func TestGeoJSON2StructInvalid(t *testing.T) {
	_, err := GeoJSON2Struct(`{"type":invalid"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[46.572,8.415]},"properties":{}}]}`)
	if err == nil {
		t.Errorf("expected error got success")
	}
}
