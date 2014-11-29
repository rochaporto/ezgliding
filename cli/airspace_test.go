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
	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/config"
	"github.com/rochaporto/ezgliding/mock"
	"time"
)

// FIXME: should go away when we start passing the context explicitly
// to the runAirspace* functions.
func setupContext(ctx config.Context) {
	config.Ctx = ctx
}

// ExampleAirspaceGet uses the mock airspace implementation to query data and
// verify airspace-get works. First, no region is passed. Second, a region but
// no updatedAfter is passed. Finally, both region and updatedAfter are given.
func ExampleAirspaceGet() {
	ctx := config.Context{
		Airspace: &mock.Airspace{
			GetF: func(regions []string, updatedSince time.Time) ([]common.Airspace, error) {
				return []common.Airspace{
					common.Airspace{ID: "MockID", Date: time.Time{}, Class: 'C', Name: "MockName", Ceiling: "1000FT AMSL",
						Floor: "500FT AMSL"},
				}, nil
			},
		},
	}
	setupContext(ctx)
	runAirspaceGet(CmdAirspaceGet, []string{})
	// Output: {MockID 0001-01-01 00:00:00 +0000 UTC 67 MockName 1000FT AMSL 500FT AMSL [] [] {0 0 <nil> <nil>}}
}

func ExampleAirspaceGetFailed() {
	ctx := config.Context{
		Airspace: &mock.Airspace{
			GetF: func(regions []string, updatedSince time.Time) ([]common.Airspace, error) {
				return nil, errors.New("mock testing get airspace failed")
			},
		},
	}
	setupContext(ctx)
	runAirspaceGet(CmdAirspaceGet, []string{})
	// Output: Failed to get airspace :: mock testing get airspace failed
}
