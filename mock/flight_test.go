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

	"github.com/rochaporto/ezgliding/common"
)

func TestGetFlight(t *testing.T) {
	flights := []common.Flight{
		common.Flight{Header: common.Header{UniqueID: "MockUniqueID"}},
	}
	mock := Mock{
		GetFlightF: func(regions []string, updatedSince time.Time) ([]common.Flight, error) {
			return flights, nil
		},
	}
	result, err := mock.GetFlight(nil, time.Time{})
	if err != nil {
		t.Errorf("failed to query mock flights")
	}
	if len(result) != len(flights) {
		t.Errorf("got %v flights but expected %v", len(result), len(flights))
	}
}

func TestGetFlightNotImplemented(t *testing.T) {
	mock := Mock{}
	result, err := mock.GetFlight(nil, time.Time{})
	if err != nil {
		t.Errorf("failed to get flight :: %v", err)
	}
	if result == nil || len(result) != 0 {
		t.Errorf("expected empty flight list but got %v", result)
	}
}

func TestGetFlightFromID(t *testing.T) {
	flights := []common.Flight{
		common.Flight{Header: common.Header{UniqueID: "MockUniqueID"}},
	}
	mock := Mock{
		GetFlightFromIDF: func(startID int, max int) ([]common.Flight, error) {
			return flights, nil
		},
	}
	result, err := mock.GetFlightFromID(10, -1)
	if err != nil {
		t.Errorf("failed to query mock flights")
	}
	if len(result) != len(flights) {
		t.Errorf("got %v flights but expected %v", len(result), len(flights))
	}
}

func TestGetFlightFromIDNotImplemented(t *testing.T) {
	mock := Mock{}
	result, err := mock.GetFlightFromID(10, -1)
	if err != nil {
		t.Errorf("failed to get flight :: %v", err)
	}
	if result == nil || len(result) != 0 {
		t.Errorf("expected empty flight list but got %v", result)
	}
}

func TestGetFlightByID(t *testing.T) {
	flight := common.Flight{Header: common.Header{UniqueID: "MockUniqueID"}}
	mock := Mock{
		GetFlightByIDF: func(id int) (common.Flight, error) {
			return flight, nil
		},
	}
	result, err := mock.GetFlightByID(10)
	if err != nil {
		t.Errorf("failed to query mock flights")
	}
	if !reflect.DeepEqual(result, flight) {
		t.Errorf("got %v but expected %v", result, flight)
	}
}

func TestGetFlightByIDNotImplemented(t *testing.T) {
	mock := Mock{}
	result, err := mock.GetFlightByID(10)
	if err != nil {
		t.Errorf("failed to get flight :: %v", err)
	}
	flight := common.Flight{}
	if !reflect.DeepEqual(result, flight) {
		t.Errorf("expected empty flight but got %v", result)
	}
}

func TestPutFlight(t *testing.T) {
	flights := []common.Flight{
		common.Flight{Header: common.Header{UniqueID: "MockUniqueID"}},
	}
	var result []common.Flight
	mock := Mock{
		PutFlightF: func(a []common.Flight) error {
			result = a
			return nil
		},
	}
	err := mock.PutFlight(flights)
	if err != nil {
		t.Errorf("failed to put mock flights")
	}
	if len(result) != len(flights) {
		t.Errorf("got %v flights but expected %v", len(result), len(flights))
	}
	for i := range result {
		if !reflect.DeepEqual(result[i], flights[i]) {
			t.Errorf("expected %v got %v", flights[i], result[i])
		}
	}
}

func TestPutFlightNotImplemented(t *testing.T) {
	mock := Mock{}
	err := mock.PutFlight(nil)
	if err != nil {
		t.Errorf("failed to put flight :: %v", err)
	}
}
