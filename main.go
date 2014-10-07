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

func main() {
	c := commander.Commander{
		Name: "ezgliding",
		Commands: []*commander.Command{
			&commander.Command{
				UsageLine: "list-releases [options]",
				Short:     "lists all available releases (airfields, waypoints, etc)",
				Long: `
Lists all available releases for the different kinds of information - airfields,
waypoints, etc. It includes all releases matching the current configuration.
`,
				Run: func(cmd *commander.Command, args []string) {
					releases, _ := welt2000.List("./welt2000/updates.xml")
					for i := range releases {
						fmt.Println("%v", releases[i])
					}
				},
				Flag: *flag.CommandLine,
			},
		},
	}

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "help")
	}
	if err := c.Run(os.Args[1:]); err != nil {
		fmt.Println("Failed running command %q: %v", os.Args[1:], err)
	}
}
