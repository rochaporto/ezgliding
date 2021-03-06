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

package web

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rochaporto/ezgliding/spatial"
)

// toOutput returns a string representation (in the requested format) of the given content.
// format should be as in the Accept: header (application/json, ...), content is an array
// of Struct which can be Airfield, Waypoint, etc.
func (srv *Server) toOutput(format string, content []interface{}) (string, error) {
	var output string
	switch format {
	case "application/json":
		collection, err := spatial.Struct2GeoJSON(content)
		if err != nil {
			return "", err
		}
		bytes, _ := json.MarshalIndent(collection, "", "\t")
		output = string(bytes)
	//case "application/csv": FIXME: enable csv output
	//output = util.Struct2CSV(content)
	default:
		return "", fmt.Errorf("format %v not supported", format)
	}
	return output, nil
}

func (srv *Server) accept(accept string) string {
	if strings.Contains(accept, "application/json") {
		return "application/json"
	} else if strings.Contains(accept, "application/csv") {
		return "application/csv"
	}
	return ""
}
