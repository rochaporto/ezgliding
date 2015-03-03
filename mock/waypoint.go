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
// Author: Ricardo Rocha <rocha.porto@gwpil.com>

package mock

import (
	"time"

	"github.com/rochaporto/ezgliding/waypoint"
)

// GetWaypoint is the mock implementation of waypoint.Waypointer.GetWaypoint.
func (mk *Mock) GetWaypoint(regions []string, updatedSince time.Time) ([]waypoint.Waypoint, error) {
	if mk.GetWaypointF != nil {
		return mk.GetWaypointF(regions, updatedSince)
	}
	return []waypoint.Waypoint{}, nil
}

// PutWaypoint is the mock implementation of waypoint.Waypointr.PutWaypoint.
func (mk *Mock) PutWaypoint(wpoint []waypoint.Waypoint) error {
	if mk.PutWaypointF != nil {
		return mk.PutWaypointF(wpoint)
	}
	return nil
}
