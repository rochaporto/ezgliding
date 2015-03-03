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

// Package context provides request context info management.
//
// It includes a common structure to be shared among the different plugin
// implementations, and which allows injecting information regarding
// configuration, global parameters or any kind of more stateful resource
// pointers (db pools, etc).
package context

import (
	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/config"
	"github.com/rochaporto/ezgliding/flight"
)

// FIXME(rocha): This should really go away.
// Only required right now for the CLI, as couldn't get to extend
// commander.Command to pass a context in addition. Once that's achieved
// we should pass the context obj explicitly to the Command objects or
// to the Run function.
var Ctx Context

// Context holds all information required between multiple calls.
type Context struct {
	Config   config.Config
	Airspace common.Airspacer
	Airfield common.Airfielder
	Flight   flight.Flighter
	Waypoint common.Waypointer
}

// NewContext returns a new Context object.
func NewContext(cfg config.Config, airspace common.Airspacer, airfield common.Airfielder,
	flight flight.Flighter, waypoint common.Waypointer) (Context, error) {
	return Context{Config: cfg, Airspace: airspace, Airfield: airfield,
		Flight: flight, Waypoint: waypoint}, nil
}
