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
	"testing"
	"time"

	"github.com/rochaporto/ezgliding/airfield"
	"github.com/rochaporto/ezgliding/config"
	"github.com/rochaporto/ezgliding/context"
	"github.com/rochaporto/ezgliding/mock"
)

// ExampleWeb .
func ExampleWeb() {
	cfg := config.Config{}
	cfg.Web.Port = 7777
	ctx := context.Context{
		Airfield: &mock.Mock{
			GetAirfieldF: func(regions []string, updatedSince time.Time) ([]airfield.Airfield, error) {
				return []airfield.Airfield{
					airfield.Airfield{ID: "MockID", ShortName: "MockShortName", Name: "MockName",
						Region: "FR", ICAO: "AAAA", Flags: 0, Catalog: 11, Length: 1000, Elevation: 2000,
						Runway: "32R", Frequency: 123.45, Latitude: 32.533, Longitude: 100.376},
				}, nil
			},
		},
		Config: cfg,
	}
	setupContext(ctx)
	go runWeb(CmdWeb, []string{})
	// Output:
}

func TestWebFailInit(t *testing.T) {
	ctx := context.Context{}
	setupContext(ctx)
	go runWeb(CmdWeb, []string{})
	// Output:
	// failed to init web server :: got a zero value Context, cannot handle this
}
func TestWebFailStart(t *testing.T) {
	cfg := config.Config{}
	cfg.Web.Port = 80
	ctx := context.Context{
		Config: cfg,
	}
	setupContext(ctx)
	go runWeb(CmdWeb, []string{})
	// Output:
	// failed to start web server :: listen tcp:80: bind: permission denied
}
