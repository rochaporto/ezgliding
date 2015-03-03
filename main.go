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

package main

import (
	"os"

	commander "code.google.com/p/go-commander"
	"github.com/golang/glog"
	"github.com/rochaporto/ezgliding/airfield"
	"github.com/rochaporto/ezgliding/airspace"
	"github.com/rochaporto/ezgliding/cli"
	"github.com/rochaporto/ezgliding/config"
	"github.com/rochaporto/ezgliding/context"
	"github.com/rochaporto/ezgliding/flight"
	"github.com/rochaporto/ezgliding/plugin"
	"github.com/rochaporto/ezgliding/waypoint"
)

func exit(c int) {
	glog.Flush()
	os.Exit(c)
}

func main() {
	defer glog.Flush()

	cfg, err := config.NewConfig("")
	if err != nil {
		glog.Errorf("Failed to load config :: %v", err)
		exit(-1)
	}
	aspace, err := plugin.NewPlugin(plugin.ID(cfg.Global.Airspacer))
	afield, err := plugin.NewPlugin(plugin.ID(cfg.Global.Airfielder))
	fght, err := plugin.NewPlugin(plugin.ID(cfg.Global.Flighter))
	wpoint, err := plugin.NewPlugin(plugin.ID(cfg.Global.Waypointer))
	aspace.Init(cfg)
	afield.Init(cfg)
	fght.Init(cfg)
	wpoint.Init(cfg)
	ctx, err := context.NewContext(cfg, aspace.(airspace.Airspacer),
		afield.(airfield.Airfielder), fght.(flight.Flighter),
		wpoint.(waypoint.Waypointer))
	if err != nil {
		glog.Errorf("Failed to create context object :: %v", err)
		exit(-1)
	}
	context.Ctx = ctx
	c := commander.Commander{
		Name: "ezgliding",
		Commands: []*commander.Command{
			cli.CmdAirfieldGet,
			cli.CmdAirfieldPut,
			cli.CmdAirspaceGet,
			cli.CmdFlightGet,
			cli.CmdWaypointGet,
			cli.CmdWaypointPut,
			cli.CmdWeb,
		},
	}

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "help")
	}
	if err := c.Run(os.Args[1:]); err != nil {
		glog.Errorf("Failed running command %q: %v", os.Args[1:], err)
		exit(-1)
	}
}
