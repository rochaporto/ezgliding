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

	"github.com/rochaporto/ezgliding/airfield"
)

func TestGetAirfield(t *testing.T) {
	airfields := []airfield.Airfield{
		airfield.Airfield{Name: "TestMockAirfield"},
	}
	mock := Mock{
		GetAirfieldF: func(regions []string, updatedSince time.Time) ([]airfield.Airfield, error) {
			return airfields, nil
		},
	}
	result, err := mock.GetAirfield(nil, time.Time{})
	if err != nil {
		t.Errorf("failed to query mock airfields")
	}
	if len(result) != len(airfields) {
		t.Errorf("got %v airfields but expected %v", len(result), len(airfields))
	}
}

func TestGetAirfieldNotImplemented(t *testing.T) {
	mock := Mock{}
	result, err := mock.GetAirfield(nil, time.Time{})
	if err != nil {
		t.Errorf("failed to get airfield :: %v", err)
	}
	if result == nil || len(result) != 0 {
		t.Errorf("expected empty airfield list but got %v", result)
	}
}

func TestPutAirfield(t *testing.T) {
	airfields := []airfield.Airfield{
		airfield.Airfield{Name: "TestMockAirfield"},
	}
	var result []airfield.Airfield
	mock := Mock{
		PutAirfieldF: func(a []airfield.Airfield) error {
			result = a
			return nil
		},
	}
	err := mock.PutAirfield(airfields)
	if err != nil {
		t.Errorf("failed to put mock airfields")
	}
	if len(result) != len(airfields) {
		t.Errorf("got %v airfields but expected %v", len(result), len(airfields))
	}
	for i := range result {
		if !reflect.DeepEqual(result[i], airfields[i]) {
			t.Errorf("expected %v got %v", airfields[i], result[i])
		}
	}
}

func TestPutAirfieldNotImplemented(t *testing.T) {
	mock := Mock{}
	err := mock.PutAirfield(nil)
	if err != nil {
		t.Errorf("failed to put airfield :: %v", err)
	}
}
