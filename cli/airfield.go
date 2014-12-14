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
	"strings"
	"time"

	commander "code.google.com/p/go-commander"
	"github.com/golang/glog"
	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/context"
	"github.com/rochaporto/ezgliding/util"
)

// CmdAirfieldGet command gets airfield information and outputs the result.
var CmdAirfieldGet = &commander.Command{
	UsageLine: "airfield-get [options]",
	Short:     "gets airfield information",
	Long: `
Gets available airfield information according to the given parameters
` + "\n" + helpFlags(flag.CommandLine),
	Run:  runAirfieldGet,
	Flag: *flag.CommandLine,
}

// runAirfieldGet invokes the configured plugin and outputs airfield data.
func runAirfieldGet(cmd *commander.Command, args []string) {
	var err error
	ctx := context.Ctx
	airfield := ctx.Airfield
	airfields, err := airfield.(common.Airfielder).GetAirfield(strings.Split(*region, ","), time.Time{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get airfield :: %v\n", err)
		// FIXME: must return -1, but no way now to check this in test
	}
	glog.V(5).Infof("airfield get with args '%v' got %d results", args, len(airfields))
	glog.V(20).Infof("%+v", airfields)
	fmt.Printf("ID,ShortName,Name,Region,ICAO,Flags,Catalog,Length,Elevation,Runway,Frequency,Location\n")
	for _, a := range airfields {
		fmt.Printf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v %v\n", a.ID, a.ShortName, a.Name, a.Region, a.ICAO,
			a.Flags, a.Catalog, a.Length, a.Elevation, a.Runway, a.Frequency,
			strconv.FormatFloat(util.DMS2Decimal(a.Latitude), 'f', 3, 64),
			strconv.FormatFloat(util.DMS2Decimal(a.Longitude), 'f', 3, 64))
	}
}
