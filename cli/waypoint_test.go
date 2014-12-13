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

package cli

import (
	"errors"
	"testing"
	"time"

	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/context"
	"github.com/rochaporto/ezgliding/mock"
)

// ExampleWaypointGet uses the mock waypoint implementation to query data and
// verify waypoint-get works. First, no region is passed. Second, a region but
// no updatedAfter is passed. Finally, both region and updatedAfter are given.
func ExampleWaypointGet() {
	ctx := context.Context{
		Waypoint: &mock.Waypoint{
			GetF: func(regions []string, updatedSince time.Time) ([]common.Waypoint, error) {
				return []common.Waypoint{
					common.Waypoint{ID: "MockID", Name: "MockName", Description: "MockDescription",
						Region: "FR", Flags: 0, Elevation: 2000, Latitude: "N323200", Longitude: "E1002233"},
				}, nil
			},
		},
	}
	setupContext(ctx)
	runWaypointGet(CmdWaypointGet, []string{})
	// Output: {ID:MockID Name:MockName Description:MockDescription Region:FR Flags:0 Elevation:2000 Latitude:N323200 Longitude:E1002233}
}

func TestWaypointGetFailed(t *testing.T) {
	ctx := context.Context{
		Waypoint: &mock.Waypoint{
			GetF: func(regions []string, updatedSince time.Time) ([]common.Waypoint, error) {
				return nil, errors.New("mock testing get waypoint failed")
			},
		},
	}
	setupContext(ctx)
	runWaypointGet(CmdWaypointGet, []string{})
}
