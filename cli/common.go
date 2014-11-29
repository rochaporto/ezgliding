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

// Package cli provides the command line tool implementations for the full
// ezgliding functionality (airspace, waypoint, airfield, ...).
//
// Most tools share a common set of flags (commonFlags, such as region,
// update time, plugins to use, ...) and will have their own additions.
package cli

import (
	"flag"
)

// commonFlags are to be shared by all (or most) commands.
var commonFlags = flag.CommandLine
var _ = commonFlags.String("after", "", "consider only items updated after this date")
var _ = commonFlags.String("region", "", "region(s) to retrieve items for, comma separated ( default is all )")

// helpFlags builds the text in 'help' regarding available command flags.
func helpFlags(fp *flag.FlagSet) string {
	result := ""
	fp.VisitAll(func(f *flag.Flag) {
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
