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

import "testing"

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
		}
	}
}
