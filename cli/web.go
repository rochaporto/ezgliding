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

	"github.com/rochaporto/ezgliding/web"

	commander "code.google.com/p/go-commander"
)

// CmdWeb command gets airspace information.
var CmdWeb = &commander.Command{
	UsageLine: "web [options]",
	Short:     "launches a local web server",
	Long: `
Launches a local web server, serving ezgliding data.

Example:
  ezgliding web -port 80
` + "\n" + helpFlags(flag.CommandLine),
	Run:  runWeb,
	Flag: *flag.CommandLine,
}

// runWeb invokes the configured plugin and outputs airspace data.
func runWeb(cmd *commander.Command, args []string) {
	srv, _ := web.NewServer(web.Config{})
	srv.Start()
}
