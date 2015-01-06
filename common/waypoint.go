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

package common

import (
	"time"
)

// Waypoint keeps details about a specific waypoint
type Waypoint struct {
	ID          string
	Name        string
	Description string
	Region      string
	Flags       int
	Elevation   int
	Latitude    float64
	Longitude   float64
}

// Waypointer is implemented in any data source which can provide or
// receive waypoint information.
type Waypointer interface {
	GetWaypoint(regions []string, updatedSince time.Time) ([]Waypoint, error)
	PutWaypoint(waypoints []Waypoint) error
}

// Enum for Waypoint flags
const ()
