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
	commander "code.google.com/p/go-commander"
	"fmt"
	"github.com/rochaporto/ezgliding/cli"
	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/config"
	"github.com/rochaporto/ezgliding/context"
	"github.com/rochaporto/ezgliding/plugin"
	"os"
)

func main() {
	cfg, err := config.NewConfig("")
	if err != nil {
		fmt.Printf("Failed to load config :: %v\n", err)
		os.Exit(-1)
	}
	if err != nil {
		fmt.Printf("Failed to find airspacer plugin :: %v\n", err)
		os.Exit(-1)
	}
	airspace, _ := plugin.NewPlugin(plugin.ID(cfg.Global.Airspacer))
	airspace.Init(cfg)
	ctx, err := context.NewContext(cfg, airspace.(common.Airspacer))
	if err != nil {
		fmt.Printf("Failed to create context object :: %v\n", err)
		os.Exit(-1)
	}
	context.Ctx = ctx
	c := commander.Commander{
		Name: "ezgliding",
		Commands: []*commander.Command{
			cli.CmdAirfieldGet,
			cli.CmdAirspaceGet,
			cli.CmdWaypointGet,
		},
	}

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "help")
	}
	if err := c.Run(os.Args[1:]); err != nil {
		fmt.Printf("Failed running command %q: %v\n", os.Args[1:], err)
		os.Exit(-1)
	}
}
