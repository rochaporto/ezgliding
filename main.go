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
	"flag"
	"fmt"
	"github.com/rochaporto/ezgliding/welt2000"
	"os"
)

var commonFlags = flag.CommandLine
var _ = commonFlags.String("after", "", "consider only items updated after this date")
var _ = commonFlags.String("region", "", "region(s) to retrieve items for, comma separated ( default is all )")

func main() {
	c := commander.Commander{
		Name: "ezgliding",
		Commands: []*commander.Command{
			CmdListAirfields,
			CmdListAirspace,
			CmdListWaypoints,
		},
		Flag: commonFlags,
	}

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "help")
	}
	if err := c.Run(os.Args[1:]); err != nil {
		fmt.Printf("Failed running command %q: %v\n", os.Args[1:], err)
	}
}

func helpFlags(flags *flag.FlagSet) string {
	result := ""
	flags.VisitAll(func(f *flag.Flag) {
		result = result + "\t" + f.Name + "\t" + f.Usage
		if f.DefValue != "" {
			result = result + " ( default is " + f.DefValue + " ) "
		}
		result = result + "\n"
	})
	if result != "" {
		result = "Options:\n" + result
	}
	return result
}

// A CmdListAirfields command lists all available airfields
var listAirfieldFlags = commonFlags
var CmdListAirfields = &commander.Command{
	UsageLine: "airfield-ls [options]",
	Short:     "lists all available airfields",
	Long: `
Lists all available releases for the different kinds of information - airfields,
waypoints, airspace, etc. It includes all releases matching the current configuration.
` + "\n" + helpFlags(listAirfieldFlags),
	Run: func(cmd *commander.Command, args []string) {
		releases, _ := welt2000.List("./welt2000/test-releases-list.xml")
		for i := range releases {
			fmt.Printf("%v\n", releases[i])
		}
	},
	Flag: *flag.CommandLine,
}

// A CmdListAirspace command lists all available airspaces
var listAirspaceFlags = commonFlags
var CmdListAirspace = &commander.Command{
	UsageLine: "airspace-ls [options]",
	Short:     "lists all latest airspace information",
	Long: `
Lists all latest airspace available, for all countries along with their
correspondent latest update time.
` + "\n" + helpFlags(listAirspaceFlags),
	Run: func(cmd *commander.Command, args []string) {
		fmt.Printf("%v", args)
		fmt.Printf("%v", listAirspaceFlags.Lookup("after").Value)
	},
	Flag: *listAirspaceFlags,
}

// A CmdListWaypoints command lists all available waypoints
var listWaypointFlags = commonFlags
var CmdListWaypoints = &commander.Command{
	UsageLine: "waypoint-ls [options]",
	Short:     "lists all waypoints",
	Long: `
Lists all waypoints available.
` + "\n" + helpFlags(listWaypointFlags),
	Run: func(cmd *commander.Command, args []string) {
	},
	Flag: *flag.CommandLine,
}
