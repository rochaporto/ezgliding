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
	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/config"
	"time"
)

// Waypoint is the mock implementation of Waypoint.
// It provides a variation where you can pass the actual function implementation.
// Especially useful for testing.
type Waypoint struct {
	GetF  func(regions []string, updatedSince time.Time) ([]common.Waypoint, error)
	PutF  func(waypoint []common.Waypoint) error
	InitF func(cfg config.Config) error
}

// Init is the mock implementation of common.Pluginer.Init.
func (wp *Waypoint) Init(cfg config.Config) error {
	return wp.InitF(cfg)
}

// GetWaypoint is the mock implementation of common.Waypointer.GetWaypoint.
func (wp *Waypoint) GetWaypoint(regions []string, updatedSince time.Time) ([]common.Waypoint, error) {
	return wp.GetF(regions, updatedSince)
}

// PutWaypoint is the mock implementation of common.Waypointr.PutWaypoint.
func (wp *Waypoint) PutWaypoint(waypoint []common.Waypoint) error {
	return wp.PutF(waypoint)
}
