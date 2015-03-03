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

type GCDistanceTest struct {
	t  string
	p1 common.Point
	p2 common.Point
	r  float64
}

var gcDistanceTests = []GCDistanceTest{
	GCDistanceTest{
		t:  "basic gc distance",
		p1: common.Point{Latitude: 46.2697223, Longitude: 6.4633333},
		p2: common.Point{Latitude: 43.6111111, Longitude: 6.6919444},
		r:  296170.7842520111,
	},
}

func TestGCDistance(t *testing.T) {
	var result float64
	for _, test := range gcDistanceTests {
		result = GCDistance(test.p1, test.p2)
		if result != test.r {
			t.Errorf("%v :: expected %v but got %v", test.t, test.r, result)
			continue
		}
	}
}

type BearingTest struct {
	t  string
	p1 common.Point
	p2 common.Point
	r  float64
}

var bearingTests = []BearingTest{
	BearingTest{
		t:  "basic bearing test",
		p1: common.Point{Latitude: 46.2697223, Longitude: 6.4633333},
		p2: common.Point{Latitude: 43.6111111, Longitude: 6.6919444},
		r:  -176.43582068293497,
	},
}

func TestBearing(t *testing.T) {
	var result float64
	for _, test := range bearingTests {
		result = Bearing(test.p1, test.p2)
		if result != test.r {
			t.Errorf("%v :: expected %v but got %v", test.t, test.r, result)
			continue
		}
	}
}

func BenchmarkDistance(b *testing.B) {
	p1 := common.Point{Latitude: 46.2697223, Longitude: 6.4633333}
	p2 := common.Point{Latitude: 43.6111111, Longitude: 6.6919444}
	for i := 0; i < b.N; i++ {
		_ = GCDistance(p1, p2)
	}
}

func BenchmarkBearing(b *testing.B) {
	p1 := common.Point{Latitude: 46.2697223, Longitude: 6.4633333}
	p2 := common.Point{Latitude: 43.6111111, Longitude: 6.6919444}
	for i := 0; i < b.N; i++ {
		_ = Bearing(p1, p2)
	}
}
