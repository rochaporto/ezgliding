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
	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/context"
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

// runWaypointGet invokes the configured plugin and outputs airfield data.
func runWaypointGet(cmd *commander.Command, args []string) {
	var err error
	ctx := context.Ctx
	waypoint := ctx.Waypoint
	waypoints, err := waypoint.(common.Waypointer).GetWaypoint(strings.Split(*region, ","), time.Time{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get waypoint :: %v", err)
		// FIXME: must return -1, but no way now to check this in test
	}
	glog.V(5).Infof("waypoint get with args '%v' got %d results", args, len(waypoints))
	glog.V(20).Infof("%+v", waypoints)
	for i := range waypoints {
		fmt.Printf("%+v\n", waypoints[i])
	}
}
