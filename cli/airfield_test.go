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

// ExampleAirfieldGet uses the mock airfield implementation to query data and
// verify airfield-get works. First, no region is passed. Second, a region but
// no updatedAfter is passed. Finally, both region and updatedAfter are given.
func ExampleAirfieldGet() {
	ctx := context.Context{
		Airfield: &mock.Airfield{
			GetF: func(regions []string, updatedSince time.Time) ([]common.Airfield, error) {
				return []common.Airfield{
					common.Airfield{ID: "MockID", ShortName: "MockShortName", Name: "MockName",
						Region: "FR", ICAO: "AAAA", Flags: 0, Catalog: 11, Length: 1000, Elevation: 2000,
						Runway: "32R", Frequency: 123.45, Latitude: "N323200", Longitude: "E1002233"},
				}, nil
			},
		},
	}
	setupContext(ctx)
	runAirfieldGet(CmdAirfieldGet, []string{})
	// Output: {ID:MockID ShortName:MockShortName Name:MockName Region:FR ICAO:AAAA Flags:0 Catalog:11 Length:1000 Elevation:2000 Runway:32R Frequency:123.45 Latitude:N323200 Longitude:E1002233}
}

func TestAirfieldGetFailed(t *testing.T) {
	ctx := context.Context{
		Airfield: &mock.Airfield{
			GetF: func(regions []string, updatedSince time.Time) ([]common.Airfield, error) {
				return nil, errors.New("mock testing get airfield failed")
			},
		},
	}
	setupContext(ctx)
	runAirfieldGet(CmdAirfieldGet, []string{})
}
