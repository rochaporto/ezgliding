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

package fusiontables

import (
	"fmt"
	"reflect"
	"time"

	"github.com/golang/glog"
	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/util"
)

// GetWaypoint follows common.GetWaypoint().
func (ft *FusionTables) GetWaypoint(regions []string, updatedSince time.Time) ([]common.Waypoint, error) {
	glog.V(10).Infof("GetWaypoint with regions %v and updatedSince %v", regions, updatedSince)

	var qry string
	qry = fmt.Sprintf("SELECT ID,Name,Description,Region,Flags,Elevation,Latitude,Longitude FROM %s", ft.WaypointTableID)
	if updatedSince != *new(time.Time) { // FIXME: shouldn't allocate Time each time
		qry = fmt.Sprintf("%v WHERE lastUpdate > %v", qry, updatedSince)
	}
	if len(regions) > 0 {
		qry = fmt.Sprintf("%v WHERE Region = '%v'", qry, regions[0])
	}
	resp, err := ft.doGet(qry)
	if err != nil {
		return nil, fmt.Errorf("%v :: %v", err, resp)
	}
	glog.V(20).Infof("unparsed response :: %v", resp)

	r, err := util.CSV2Struct(resp, reflect.ValueOf([]common.Waypoint{}).Type(),
		reflect.ValueOf(common.Waypoint{}).Type())
	if err != nil {
		return nil, err
	}
	result := r.Interface().([]common.Waypoint)
	glog.V(5).Infof("request %v returned %v results", qry, len(result))

	return result, nil
}

// PutWaypoint follows common.PutWaypoint().
func (ft *FusionTables) PutWaypoint(waypoints []common.Waypoint) error {
	csv := util.Struct2CSV(waypoints)
	resp, err := ft.doImport(csv, ft.WaypointTableID)
	if err != nil {
		return fmt.Errorf("failed to put waypoint :: %v %v", resp, err)
	}
	return nil
}
