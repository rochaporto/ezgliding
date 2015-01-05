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

// ExampleAirfieldGet uses the mock airfield implementation to query data and
// verify airfield-get works. First, no region is passed. Second, a region but
// no updatedAfter is passed. Finally, both region and updatedAfter are given.
func ExampleAirfieldGet() {
	ctx := context.Context{
		Airfield: &mock.Mock{
			GetAirfieldF: func(regions []string, updatedSince time.Time) ([]common.Airfield, error) {
				return []common.Airfield{
					common.Airfield{ID: "MockID", ShortName: "MockShortName", Name: "MockName",
						Region: "FR", ICAO: "AAAA", Flags: 0, Catalog: 11, Length: 1000, Elevation: 2000,
						Runway: "32R", Frequency: 123.45, Latitude: 32.533, Longitude: 100.376},
				}, nil
			},
		},
	}
	setupContext(ctx)
	runAirfieldGet(CmdAirfieldGet, []string{})
	// Output:
	// ID,ShortName,Name,Region,ICAO,Flags,Catalog,Length,Elevation,Runway,Frequency,Latitude,Longitude
	// MockID,MockShortName,MockName,FR,AAAA,0,11,1000,2000,32R,123.45,32.533,100.376
}

func TestAirfieldGetFailed(t *testing.T) {
	ctx := context.Context{
		Airfield: &mock.Mock{
			GetAirfieldF: func(regions []string, updatedSince time.Time) ([]common.Airfield, error) {
				return nil, errors.New("mock testing get airfield failed")
			},
		},
	}
	setupContext(ctx)
	runAirfieldGet(CmdAirfieldGet, []string{})
}

// ExampleAirfieldPut uses the mock airfield implementation to push data and
// verify airfield-put works.
func ExampleAirfieldPut() {
	ctx := context.Context{
		Airfield: &mock.Mock{
			GetAirfieldF: func(regions []string, updatedSince time.Time) ([]common.Airfield, error) {
				return []common.Airfield{
					common.Airfield{ID: "MockID", ShortName: "MockShortName", Name: "MockName",
						Region: "FR", ICAO: "AAAA", Flags: 0, Catalog: 11, Length: 1000, Elevation: 2000,
						Runway: "32R", Frequency: 123.45, Latitude: 32.533, Longitude: 100.376},
				}, nil
			},
		},
	}
	setupContext(ctx)
	runAirfieldPut(CmdAirfieldPut, []string{"mock"})
	// Output:
	// pushed 1 airfields into mock
}

func TestAirfieldPutFailed(t *testing.T) {
	err := plugin.Register(plugin.ID("afputfailed"), plugin.Pluginer(
		&mock.Mock{
			PutAirfieldF: func(airfields []common.Airfield) error {
				return errors.New("mock testing put airfield failed")
			},
		}))
	if err != nil {
		t.Errorf("failed to register plugin :: %v", err)
	}
	ctx := context.Context{
		Airfield: &mock.Mock{
			GetAirfieldF: func(regions []string, updatedSince time.Time) ([]common.Airfield, error) {
				return []common.Airfield{common.Airfield{}}, nil
			},
		},
	}
	setupContext(ctx)
	runAirfieldPut(CmdAirfieldPut, []string{"afputfailed"})
}

func TestAirfieldPutBadGet(t *testing.T) {
	ctx := context.Context{
		Airfield: &mock.Mock{
			GetAirfieldF: func(regions []string, updatedSince time.Time) ([]common.Airfield, error) {
				return nil, errors.New("mock testing get airfield failed")
			},
		},
	}
	setupContext(ctx)
	runAirfieldPut(CmdAirfieldPut, []string{"mock"})
}

func TestAirfieldPutBadPluginID(t *testing.T) {
	ctx := context.Context{
		Airfield: &mock.Mock{},
	}
	setupContext(ctx)
	runAirfieldPut(CmdAirfieldPut, []string{"afnonexisting"})
}

func TestAirfieldPutFailInit(t *testing.T) {
	err := plugin.Register(plugin.ID("affailinit"), plugin.Pluginer(
		&mock.Mock{
			InitF: func(config.Config) error {
				return errors.New("failed to init plugin")
			},
		}))
	if err != nil {
		t.Errorf("failed to register plugin :: %v", err)
	}
	ctx := context.Context{
		Airfield: &mock.Mock{
			GetAirfieldF: func(regions []string, updatedSince time.Time) ([]common.Airfield, error) {
				return []common.Airfield{common.Airfield{}}, nil
			},
		},
	}
	setupContext(ctx)
	runAirfieldPut(CmdAirfieldPut, []string{"affailinit"})
}

func TestAirfieldPutBadArgNumber(t *testing.T) {
	ctx := context.Context{
		Airfield: &mock.Mock{},
	}
	setupContext(ctx)
	runAirfieldPut(CmdAirfieldPut, []string{})
}
