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
	"flag"
	"testing"
	"time"

	"github.com/rochaporto/ezgliding/config"
	"github.com/rochaporto/ezgliding/context"
	"github.com/rochaporto/ezgliding/mock"
	"github.com/rochaporto/ezgliding/plugin"
	"github.com/rochaporto/ezgliding/waypoint"
)

// ExampleWaypointGet uses the mock waypoint implementation to query data and
// verify waypoint-get works. First, no region is passed. Second, a region but
// no updatedAfter is passed. Finally, both region and updatedAfter are given.
func ExampleWaypointGet() {
	ctx := context.Context{
		Waypoint: &mock.Mock{
			GetWaypointF: func(regions []string, updatedSince time.Time) ([]waypoint.Waypoint, error) {
				waypoints := []waypoint.Waypoint{
					waypoint.Waypoint{
						ID: "MockID1", Name: "MockName",
						Description: "MockDescription", Region: "FR", Flags: 0,
						Elevation: 2000, Latitude: 32.533, Longitude: 100.376,
						Update: time.Date(2014, 02, 01, 0, 0, 0, 0, time.UTC),
					},
					waypoint.Waypoint{
						ID: "MockID2", Name: "MockName",
						Description: "MockDescription", Region: "CH", Flags: 0,
						Elevation: 2000, Latitude: 32.533, Longitude: 100.376,
						Update: time.Date(2014, 02, 02, 0, 0, 0, 0, time.UTC),
					},
					waypoint.Waypoint{
						ID: "MockID3", Name: "MockName",
						Description: "MockDescription", Region: "CH", Flags: 0,
						Elevation: 2000, Latitude: 32.533, Longitude: 100.376,
						Update: time.Date(2014, 02, 03, 0, 0, 0, 0, time.UTC),
					},
				}
				result := []waypoint.Waypoint{}
				for _, waypoint := range waypoints {
					b := false
					for _, r := range regions {
						if waypoint.Region == r {
							b = true
						}
					}
					if waypoint.Update.After(updatedSince) && b {
						result = append(result, waypoint)
					}
				}
				return result, nil
			},
		},
	}
	setupContext(ctx)
	flag.Set("region", "CH")
	flag.Set("after", "2014-02-02")
	runWaypointGet(CmdWaypointGet, []string{})
	// Output:
	// ID,Name,Description,Region,Flags,Elevation,Latitude,Longitude,Update
	// MockID3,MockName,MockDescription,CH,0,2000,32.533,100.376,2014-02-03 00:00:00 +0000 UTC
}

func TestWaypointGetFailed(t *testing.T) {
	ctx := context.Context{
		Waypoint: &mock.Mock{
			GetWaypointF: func(regions []string, updatedSince time.Time) ([]waypoint.Waypoint, error) {
				return nil, errors.New("mock testing get waypoint failed")
			},
		},
	}
	setupContext(ctx)
	flag.Set("region", "")
	flag.Set("after", "")
	runWaypointGet(CmdWaypointGet, []string{})
}

func TestWaypointGetBadAfter(t *testing.T) {
	ctx := context.Context{
		Waypoint: &mock.Mock{
			GetWaypointF: func(regions []string, updatedSince time.Time) ([]waypoint.Waypoint, error) {
				return nil, nil
			},
		},
	}
	setupContext(ctx)
	flag.Set("after", "22-00-11")
	flag.Set("region", "")
	runWaypointGet(CmdWaypointGet, []string{})
}

// ExampleWaypointPut uses the mock waypoint implementation to push data and
// verify waypoint-put works.
func ExampleWaypointPut() {
	ctx := context.Context{
		Waypoint: &mock.Mock{
			GetWaypointF: func(regions []string, updatedSince time.Time) ([]waypoint.Waypoint, error) {
				return []waypoint.Waypoint{
					waypoint.Waypoint{ID: "MockID", Name: "MockName", Description: "MockDescription",
						Region: "FR", Flags: 0, Elevation: 2000, Latitude: 32.533, Longitude: 100.576},
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
			PutWaypointF: func(waypoints []waypoint.Waypoint) error {
				return errors.New("mock testing put waypoint failed")
			},
		}))
	if err != nil {
		t.Errorf("failed to register plugin :: %v", err)
	}
	ctx := context.Context{
		Waypoint: &mock.Mock{
			GetWaypointF: func(regions []string, updatedSince time.Time) ([]waypoint.Waypoint, error) {
				return []waypoint.Waypoint{waypoint.Waypoint{}}, nil
			},
		},
	}
	setupContext(ctx)
	runWaypointPut(CmdWaypointPut, []string{"wpputfailed"})
}

func TestWaypointPutBadGet(t *testing.T) {
	ctx := context.Context{
		Waypoint: &mock.Mock{
			GetWaypointF: func(regions []string, updatedSince time.Time) ([]waypoint.Waypoint, error) {
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
			GetWaypointF: func(regions []string, updatedSince time.Time) ([]waypoint.Waypoint, error) {
				return []waypoint.Waypoint{waypoint.Waypoint{}}, nil
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
