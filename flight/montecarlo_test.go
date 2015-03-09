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

package flight

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type MontecarloTest struct {
	t   string
	loc string
	res Result
}

var montecarloTests = []MontecarloTest{
	MontecarloTest{t: "sample flight", loc: "./t/sample-flight.igc", res: Result{Distance: 710000.0}},
	MontecarloTest{t: "sample flight2", loc: "./t/sample-flight2.igc", res: Result{Distance: 710000.0}},
}

func TestMontecarlo(t *testing.T) {
	mc := NewMontecarlo()
	for _, test := range montecarloTests {
		content, err := fetch(test.loc)
		if err != nil {
			t.Errorf("failed to load flight :: %v", err)
			continue
		}
		f, err := ParseIGC(content)
		if err != nil {
			t.Errorf("failed to parse content :: %v", err)
			continue
		}
		result, err := mc.Optimize(f.Points)
		if err != nil {
			t.Errorf("failed montecarlo optimize :: %v", err)
			continue
		}
		fmt.Printf("%v\n%v\n", test.t, result.Distance)
		for i, p := range result.TurnPoints {
			fmt.Printf("data.setCell(%v, 0, %v);\n", i, p.Latitude)
			fmt.Printf("data.setCell(%v, 1, %v);\n", i, p.Longitude)
			fmt.Printf("data.setCell(%v, 2, 'WP%v');\n", i, i)
		}
	}
}

// FIXME: should be a common function in another package
func fetch(location string) (string, error) {
	var content []byte
	// case http
	resp, err := http.Get(location)
	if err == nil {
		defer resp.Body.Close()
		content, err = ioutil.ReadAll(resp.Body)
	} else { // case file
		resp, err := ioutil.ReadFile(location)
		if err != nil {
			return "", err
		}
		content = resp
	}
	return string(content), nil
}
