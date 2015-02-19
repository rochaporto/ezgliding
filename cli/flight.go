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
	"strconv"

	"github.com/golang/glog"
	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/context"

	commander "code.google.com/p/go-commander"
)

var (
	startID = flag.String("startID", "", "return only flights with ID higher than this")
	id      = flag.String("id", "", "return flight with this ID")
	max     = flag.String("max", "", "max number of flights to return (used with startID)")
)

// CmdFlightGet command gets flight information.
var CmdFlightGet = &commander.Command{
	UsageLine: "flight-get [options]",
	Short:     "gets flight information",
	Long: `
	Retrieves flight information from the configured plugin.
` + "\n" + helpFlags(flag.CommandLine),
	Run:  runFlightGet,
	Flag: *flag.CommandLine,
}

// runFlightGet invokes the configured plugin and outputs flight data.
func runFlightGet(cmd *commander.Command, args []string) {
	ctx := context.Ctx
	f := ctx.Flight
	glog.V(20).Infof("flight plugin instance ::  %v", f)

	var flights []common.Flight
	if *startID != "" {
		vmax := -1
		// query for flights with ID higher than startID
		sid, err := strconv.Atoi(*startID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get flight :: %v\n", err)
			return
		}
		if *max != "" {
			vmax, err = strconv.Atoi(*max)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to get flight :: %v\n", err)
				return
			}
		}
		glog.V(10).Infof("querying from start id %v, max %v", sid, vmax)
		flights, err = f.(common.Flighter).GetFlightFromID(sid, vmax)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get flight :: %v\n", err)
			return
		}
	} else if *id != "" {
		// query for flights with specific ID
		sid, err := strconv.Atoi(*id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get flight :: %v\n", err)
			return
		}
		glog.V(10).Infof("querying id %v", sid)
		flight, err := f.(common.Flighter).GetFlightByID(sid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get flight :: %v\n", err)
			return
			// FIXME: must return -1, but no way now to check this in test
		}
		flights = append(flights, flight)
	}
	glog.V(5).Infof("flight get with args '%v' got %d results", args, len(flights))
	glog.V(20).Infof("%+v", flights)
	for _, flight := range flights {
		fmt.Printf("%v,%v,%v,%v\n", flight.Header.Date.Format("02/01/2006"),
			flight.Header.Pilot, flight.Header.GliderType, flight.Header.GliderID)
		for sourceID, source := range flight.Sources {
			fmt.Printf("\t%v,%v,%v,%v,%v,%v,%v,%v\n", sourceID, source.Name,
				source.Category, source.Club, source.Country, source.Region,
				source.Distance, source.Points)
		}
	}
}
