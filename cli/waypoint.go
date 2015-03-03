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
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	commander "code.google.com/p/go-commander"
	"github.com/golang/glog"
	"github.com/rochaporto/ezgliding/context"
	"github.com/rochaporto/ezgliding/plugin"
	"github.com/rochaporto/ezgliding/util"
	"github.com/rochaporto/ezgliding/waypoint"
)

// CmdWaypointGet command gets waypoint information and outputs the result.
var CmdWaypointGet = &commander.Command{
	UsageLine: "waypoint-get [options]",
	Short:     "gets waypoint information",
	Long: `
Gets waypoint information according to the given parameters.
` + "\n" + helpFlags(flag.CommandLine),
	Run:  runWaypointGet,
	Flag: *flag.CommandLine,
}

// runWaypointGet invokes the configured plugin and outputs waypoint data.
func runWaypointGet(cmd *commander.Command, args []string) {
	var err error
	ctx := context.Ctx
	wpoint := ctx.Waypoint

	tafter := time.Time{}
	if *after != "" {
		tafter, err = time.Parse("2006-01-02", *after)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get waypoint :: %v\n", err)
			return
		}
	}
	waypoints, err := wpoint.(waypoint.Waypointer).GetWaypoint(strings.Split(*region, ","), tafter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get waypoint :: %v", err)
		// FIXME: must return -1, but no way now to check this in test
	}
	glog.V(5).Infof("waypoint get with args '%v' got %d results", args, len(waypoints))
	glog.V(20).Infof("%+v", waypoints)
	fmt.Printf("%v", util.Struct2CSV(waypoints))
}

// CmdWaypointPut command puts waypoint information from a source to a destination.
var CmdWaypointPut = &commander.Command{
	UsageLine: "waypoint-put [options] destination",
	Short:     "puts waypoint information",
	Long: `
Puts waypoint information according to the given parameters
` + "\n" + helpFlags(flag.CommandLine),
	Run:  runWaypointPut,
	Flag: *flag.CommandLine,
}

// runWaypointPut invokes the configured plugins to put waypoint data from source to dest.
func runWaypointPut(cmd *commander.Command, args []string) {
	var err error
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "failed to put waypoint data :: no destination given\n")
		return
	}
	pluginID := args[0]
	ctx := context.Ctx
	destPlugin, err := plugin.NewPlugin(plugin.ID(pluginID))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get plugin '%v' :: %v\n", pluginID, err)
		return
	}
	err = destPlugin.Init(ctx.Config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to init plugin '%v' :: %v\n", pluginID, err)
		return
	}
	wpoint := ctx.Waypoint
	waypoints, err := wpoint.(waypoint.Waypointer).GetWaypoint(strings.Split(*region, ","), time.Time{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get waypoint :: %v\n", err)
		return
	}
	glog.V(5).Infof("putting %v waypoints", len(waypoints))
	glog.V(20).Infof("%v", waypoints)
	if len(waypoints) > 0 {
		err = destPlugin.(waypoint.Waypointer).PutWaypoint(waypoints)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to put waypoints :: %v\n", err)
			return
		}
	}
	fmt.Printf("pushed %v waypoints into %v\n", len(waypoints), pluginID)
}
