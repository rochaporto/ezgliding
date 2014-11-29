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
	commander "code.google.com/p/go-commander"
	"fmt"
	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/config"
	"time"
)

var airspaceGetFlags = commonFlags

// CmdGetAirspace command gets airspace information.
var CmdAirspaceGet = &commander.Command{
	UsageLine: "airspace-get [options]",
	Short:     "gets airspace information",
	Long: `
Gets latest airspace data corresponding to the given parameters.
` + "\n" + helpFlags(airspaceGetFlags),
	Run:  runAirspaceGet,
	Flag: *airspaceGetFlags,
}

// runAirspaceGet invokes the configured plugin and outputs airspace data.
func runAirspaceGet(cmd *commander.Command, args []string) {
	var err error
	ctx := config.Ctx
	airspace := ctx.Airspace
	airspaces, err := airspace.(common.Airspacer).GetAirspace([]string{"FR"}, time.Time{})
	if err != nil {
		fmt.Printf("Failed to get airspace :: %v\n", err)
	}
	for i := range airspaces {
		fmt.Printf("\n%v\n", airspaces[i])
	}

}
