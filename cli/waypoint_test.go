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
	"github.com/rochaporto/ezgliding/config"
	"github.com/rochaporto/ezgliding/context"
	"github.com/rochaporto/ezgliding/mock"
	"github.com/rochaporto/ezgliding/plugin"
)

// ExampleWaypointGet uses the mock waypoint implementation to query data and
// verify waypoint-get works. First, no region is passed. Second, a region but
// no updatedAfter is passed. Finally, both region and updatedAfter are given.
func ExampleWaypointGet() {
	ctx := context.Context{
		Waypoint: &mock.Mock{
			GetWaypointF: func(regions []string, updatedSince time.Time) ([]common.Waypoint, error) {
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
		Waypoint: &mock.Mock{
			GetWaypointF: func(regions []string, updatedSince time.Time) ([]common.Waypoint, error) {
				return nil, errors.New("mock testing get waypoint failed")
			},
		},
	}
	setupContext(ctx)
	runWaypointGet(CmdWaypointGet, []string{})
}

// ExampleWaypointPut uses the mock waypoint implementation to push data and
// verify waypoint-put works.
func ExampleWaypointPut() {
	ctx := context.Context{
		Waypoint: &mock.Mock{
			GetWaypointF: func(regions []string, updatedSince time.Time) ([]common.Waypoint, error) {
				return []common.Waypoint{
					common.Waypoint{ID: "MockID", Name: "MockName", Description: "MockDescription",
						Region: "FR", Flags: 0, Elevation: 2000, Latitude: "N323200", Longitude: "E1002233"},
				}, nil
			},
		},
	}
	setupContext(ctx)
	runWaypointPut(CmdWaypointPut, []string{"mock"})
	// Output:
	// pushed 1 waypoints into mock
}

func TestWaypointPutFailed(t *testing.T) {
	err := plugin.Register(plugin.ID("wpputfailed"), plugin.Pluginer(
		&mock.Mock{
			PutWaypointF: func(waypoints []common.Waypoint) error {
				return errors.New("mock testing put waypoint failed")
			},
		}))
	if err != nil {
		t.Errorf("failed to register plugin :: %v", err)
	}
	ctx := context.Context{
		Waypoint: &mock.Mock{
			GetWaypointF: func(regions []string, updatedSince time.Time) ([]common.Waypoint, error) {
				return []common.Waypoint{common.Waypoint{}}, nil
			},
		},
	}
	setupContext(ctx)
	runWaypointPut(CmdWaypointPut, []string{"wpputfailed"})
}

func TestWaypointPutBadGet(t *testing.T) {
	ctx := context.Context{
		Waypoint: &mock.Mock{
			GetWaypointF: func(regions []string, updatedSince time.Time) ([]common.Waypoint, error) {
				return nil, errors.New("mock testing get waypoint failed")
			},
		},
	}
	setupContext(ctx)
	runWaypointPut(CmdWaypointPut, []string{"mock"})
}

func TestWaypointPutBadPluginID(t *testing.T) {
	ctx := context.Context{
		Waypoint: &mock.Mock{},
	}
	setupContext(ctx)
	runWaypointPut(CmdWaypointPut, []string{"wpnonexisting"})
}

func TestWaypointPutFailInit(t *testing.T) {
	err := plugin.Register(plugin.ID("wpfailinit"), plugin.Pluginer(
		&mock.Mock{
			InitF: func(config.Config) error {
				return errors.New("failed to init plugin")
			},
		}))
	if err != nil {
		t.Errorf("failed to register plugin :: %v", err)
	}
	ctx := context.Context{
		Waypoint: &mock.Mock{
			GetWaypointF: func(regions []string, updatedSince time.Time) ([]common.Waypoint, error) {
				return []common.Waypoint{common.Waypoint{}}, nil
			},
		},
	}
	setupContext(ctx)
	runWaypointPut(CmdWaypointPut, []string{"wpfailinit"})
}

func TestWaypointPutBadArgNumber(t *testing.T) {
	ctx := context.Context{
		Waypoint: &mock.Mock{},
	}
	setupContext(ctx)
	runWaypointPut(CmdWaypointPut, []string{})
}
