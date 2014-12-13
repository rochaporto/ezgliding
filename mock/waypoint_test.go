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
	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/config"
	"github.com/rochaporto/ezgliding/plugin"
	"testing"
	"time"
)

func TestWaypointInit(t *testing.T) {
	mockI := &Waypoint{
		InitF: func(cfg config.Config) error {
			return nil
		},
	}
	x := plugin.Pluginer(mockI)
	err := x.Init(config.Config{})
	if err != nil {
		t.Errorf("Failed to call init on mock waypoint")
	}
}

func TestGetWaypoint(t *testing.T) {
	waypoints := []common.Waypoint{
		common.Waypoint{Name: "TestMockWaypoint"},
	}
	mockI := &Waypoint{
		GetF: func(regions []string, updatedSince time.Time) ([]common.Waypoint, error) {
			return waypoints, nil
		},
	}
	x := common.Waypointer(mockI)
	result, err := x.GetWaypoint(nil, time.Time{})
	if err != nil {
		t.Errorf("Failed to query mock waypoints")
	}
	if len(result) != len(waypoints) {
		t.Errorf("Got %v waypoints but expected %v", len(result), len(waypoints))
	}
}
func TestPutWaypoint(t *testing.T) {
	mockI := &Waypoint{
		PutF: func([]common.Waypoint) error {
			return nil // FIXME: implement
		},
	}
	x := common.Waypointer(mockI)
	err := x.PutWaypoint(nil) // FIXME: implement
	if err != nil {
		t.Errorf("Failed to put mock waypoints")
	}
}
