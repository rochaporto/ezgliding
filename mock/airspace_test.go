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

package mock

import (
	"reflect"
	"testing"
	"time"

	"github.com/rochaporto/ezgliding/airspace"
)

func TestGetAirspace(t *testing.T) {
	airspaces := []airspace.Airspace{
		airspace.Airspace{Name: "TestMockAirspace"},
	}
	mock := Mock{
		GetAirspaceF: func(regions []string, updatedSince time.Time) ([]airspace.Airspace, error) {
			return airspaces, nil
		},
	}
	result, err := mock.GetAirspace(nil, time.Time{})
	if err != nil {
		t.Errorf("Failed to query mock airspaces")
	}
	if len(result) != len(airspaces) {
		t.Errorf("Got %v airspaces but expected %v", len(result), len(airspaces))
	}
}

func TestGetAirspaceNotImplemented(t *testing.T) {
	mock := Mock{}
	result, err := mock.GetAirspace(nil, time.Time{})
	if err != nil {
		t.Errorf("failed to get airspace :: %v", err)
	}
	if result == nil || len(result) != 0 {
		t.Errorf("expected empty list but got %v", result)
	}
}

func TestPutAirspace(t *testing.T) {
	airspaces := []airspace.Airspace{
		airspace.Airspace{Name: "TestMockAirspace"},
	}
	var result []airspace.Airspace
	mock := Mock{
		PutAirspaceF: func(a []airspace.Airspace) error {
			result = a
			return nil
		},
	}
	err := mock.PutAirspace(airspaces)
	if err != nil {
		t.Errorf("Failed to put mock airspaces")
	}
	if len(result) != len(airspaces) {
		t.Errorf("got %v airspaces but expected %v", len(result), len(airspaces))
	}
	for i := range result {
		if !reflect.DeepEqual(result[i], airspaces[i]) {
			t.Errorf("expected %v got %v", airspaces[i], result[i])
		}
	}
}

func TestPutAirspaceNotImplemented(t *testing.T) {
	mock := Mock{}
	err := mock.PutAirspace(nil)
	if err != nil {
		t.Errorf("failed to put airspace :: %v", err)
	}
}
