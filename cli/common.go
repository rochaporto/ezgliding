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

// helpFlags builds the text in 'help' regarding available command flags.
func helpFlags(fp *flag.FlagSet) string {
	result := ""
	fp.VisitAll(func(f *flag.Flag) {
		if f.DefValue != "" {
			result = result + "  -" + f.Name + "=" + f.DefValue + "\n    " + f.Usage
		} else {
			result = result + "  -" + f.Name + "\n    " + f.Usage
		}
		result = result + "\n"
	})
	if result != "" {
		result = "Full Options:\n" + result
	}
	return result
}
