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

func TestGetWaypoint(t *testing.T) {
	waypoints := []common.Waypoint{
		common.Waypoint{Name: "TestMockWaypoint"},
	}
	mock := Mock{
		GetWaypointF: func(regions []string, updatedSince time.Time) ([]common.Waypoint, error) {
			return waypoints, nil
		},
	}
	result, err := mock.GetWaypoint(nil, time.Time{})
	if err != nil {
		t.Errorf("Failed to query mock waypoints")
	}
	if len(result) != len(waypoints) {
		t.Errorf("Got %v waypoints but expected %v", len(result), len(waypoints))
	}
}

func TestGetWaypointNotImplemented(t *testing.T) {
	mock := Mock{}
	result, err := mock.GetWaypoint(nil, time.Time{})
	if err != nil {
		t.Errorf("failed to get waypoint :: %v", err)
	}
	if result == nil || len(result) != 0 {
		t.Errorf("expected empty list but got %v", result)
	}
}

func TestPutWaypoint(t *testing.T) {
	waypoints := []common.Waypoint{
		common.Waypoint{Name: "TestMockWaypoint"},
	}
	var result []common.Waypoint
	mock := Mock{
		PutWaypointF: func(w []common.Waypoint) error {
			result = w
			return nil
		},
	}
	err := mock.PutWaypoint(waypoints)
	if err != nil {
		t.Errorf("Failed to put mock waypoints")
	}
	if len(result) != len(waypoints) {
		t.Errorf("got %v waypoints but expected %v", len(result), len(waypoints))
	}
	for i := range result {
		if !reflect.DeepEqual(result[i], waypoints[i]) {
			t.Errorf("expected %v got %v", waypoints[i], result[i])
		}
	}
}

func TestPutWaypointNotImplemented(t *testing.T) {
	mock := Mock{}
	err := mock.PutWaypoint(nil)
	if err != nil {
		t.Errorf("failed to put waypoint :: %v", err)
	}
}
