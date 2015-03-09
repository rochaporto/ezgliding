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

package spatial

import (
	"reflect"
	"testing"

	"github.com/rochaporto/ezgliding/airfield"
	"github.com/rochaporto/ezgliding/waypoint"
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

type DMD2DecimalTest struct {
	t  string
	in string
	r  float64
}

var dmd2DecimalTests = []DMD2DecimalTest{
	{
		"latitude north conversion",
		"N4616018",
		46.26696666666667,
	},
	{
		"latitude north conversion inverted",
		"4616018N",
		46.26696666666667,
	},
	{
		"latitude south conversion",
		"S4616018",
		-46.26696666666667,
	},
	{
		"latitude south conversion inverted",
		"4616018S",
		-46.26696666666667,
	},
	{
		"longitude east conversion",
		"E00627679",
		6.461316666666667,
	},
	{
		"longitude east conversion inverted",
		"00627679E",
		6.461316666666667,
	},
	{
		"longitude west conversion",
		"W00627679",
		-6.461316666666667,
	},
	{
		"longitude west conversion inverted",
		"00627679W",
		-6.461316666666667,
	},
}

func TestDMD2Decimal(t *testing.T) {
	for _, test := range dmd2DecimalTests {
		result := DMD2Decimal(test.in)
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
			airfield.Airfield{ID: "HABER", ShortName: "HABER", Name: "HABERE POC69",
				Region: "FR", ICAO: "HHHH", Flags: 1032, Catalog: 69, Length: 900, Elevation: 1113,
				Runway: "0119", Frequency: 122.5, Latitude: 46.270, Longitude: 6.463},
		},
		`{"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[6.463,46.27]},"properties":{"Catalog":69,"Elevation":1113,"Flags":1032,"Frequency":122.5,"Go":"Airfield","ICAO":"HHHH","ID":"HABER","Length":900,"Name":"HABERE POC69","Region":"FR","Runway":"0119","ShortName":"HABER"}}]}`,
	},
	Struct2GeoJSONTest{
		"simple waypoint conversion",
		[]interface{}{
			waypoint.Waypoint{
				ID: "FURKAP", Name: "FURKAP", Description: "FURKAPASS PASSHOEHE", Elevation: 2432,
				Latitude: 46.572, Longitude: 8.415, Region: "CH", Flags: 0,
			},
		},
		`{"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[8.415,46.572]},"properties":{"Description":"FURKAPASS PASSHOEHE","Elevation":2432,"Flags":0,"Go":"Waypoint","ID":"FURKAP","Name":"FURKAP","Region":"CH"}}]}`,
	},
	Struct2GeoJSONTest{
		"multiple type conversion",
		[]interface{}{
			airfield.Airfield{ID: "HABER", ShortName: "HABER", Name: "HABERE POC69",
				Region: "FR", ICAO: "HHHH", Flags: 1032, Catalog: 69, Length: 900, Elevation: 1113,
				Runway: "0119", Frequency: 122.5, Latitude: 46.270, Longitude: 6.463},
			waypoint.Waypoint{
				ID: "FURKAP", Name: "FURKAP", Description: "FURKAPASS PASSHOEHE", Elevation: 2432,
				Latitude: 46.572, Longitude: 8.415, Region: "CH", Flags: 0,
			},
		},
		`{"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[6.463,46.27]},"properties":{"Catalog":69,"Elevation":1113,"Flags":1032,"Frequency":122.5,"Go":"Airfield","ICAO":"HHHH","ID":"HABER","Length":900,"Name":"HABERE POC69","Region":"FR","Runway":"0119","ShortName":"HABER"}},{"type":"Feature","geometry":{"type":"Point","coordinates":[8.415,46.572]},"properties":{"Description":"FURKAPASS PASSHOEHE","Elevation":2432,"Flags":0,"Go":"Waypoint","ID":"FURKAP","Name":"FURKAP","Region":"CH"}}]}`,
	},
}

func TestStruct2GeoJSON(t *testing.T) {
	for _, test := range struct2GeoJSONTests {
		result, err := Struct2GeoJSON(test.in)
		if err != nil {
			t.Errorf("%v failed :: %v", test.t, err)
			continue
		}
		geo, err := result.MarshalJSON()
		if err != nil {
			t.Errorf("%v failed :: %v", test.t, err)
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
			t.Errorf("%v failed :: %v", test.t, err)
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
	_, err := GeoJSON2Struct(`{"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[8.415,46.572]},"properties":{"Go":"UnsupportedType"}}]}`)
	if err == nil {
		t.Errorf("expected error got success")
	}
}

func TestGeoJSON2StructInvalid(t *testing.T) {
	_, err := GeoJSON2Struct(`{"type":invalid"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[8.415,46.572]},"properties":{}}]}`)
	if err == nil {
		t.Errorf("expected error got success")
	}
}

type GCDistanceTest struct {
	t  string
	p1 []float64
	p2 []float64
	r  float64
}

var gcDistanceTests = []GCDistanceTest{
	GCDistanceTest{
		t:  "basic gc distance",
		p1: []float64{46.2697223, 6.4633333},
		p2: []float64{43.6111111, 6.6919444},
		r:  296170.7842520111,
	},
}

func TestGCDistance(t *testing.T) {
	var result float64
	for _, test := range gcDistanceTests {
		result = GCDistance(test.p1[0], test.p1[1], test.p2[0], test.p2[1])
		if result != test.r {
			t.Errorf("%v :: expected %v but got %v", test.t, test.r, result)
			continue
		}
	}
}

type BearingTest struct {
	t  string
	p1 []float64
	p2 []float64
	r  float64
}

var bearingTests = []BearingTest{
	BearingTest{
		t:  "basic bearing test",
		p1: []float64{46.2697223, 6.4633333},
		p2: []float64{43.6111111, 6.6919444},
		r:  -176.43582068293497,
	},
}

func TestBearing(t *testing.T) {
	var result float64
	for _, test := range bearingTests {
		result = Bearing(test.p1[0], test.p1[1], test.p2[0], test.p2[1])
		if result != test.r {
			t.Errorf("%v :: expected %v but got %v", test.t, test.r, result)
			continue
		}
	}
}

func BenchmarkDistance(b *testing.B) {
	p1 := []float64{46.2697223, 6.4633333}
	p2 := []float64{43.6111111, 6.6919444}
	for i := 0; i < b.N; i++ {
		_ = GCDistance(p1[0], p1[1], p2[0], p2[1])
	}
}

func BenchmarkBearing(b *testing.B) {
	p1 := []float64{46.2697223, 6.4633333}
	p2 := []float64{43.6111111, 6.6919444}
	for i := 0; i < b.N; i++ {
		_ = Bearing(p1[0], p1[1], p2[0], p2[1])
	}
}
