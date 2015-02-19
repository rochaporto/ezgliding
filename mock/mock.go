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

// Package mock provides mock implementation of all interfaces.
//
package mock

import (
	"time"

	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/config"
)

const (
	// ID for this plugin implementation.
	ID string = "mock"
)

// Mock implements all interfaces (airfield, airspace, waypoint).
// Use the struct fields to provide the function implementations.
type Mock struct {
	InitF            func(cfg config.Config) error
	GetAirfieldF     func(regions []string, updatedSince time.Time) ([]common.Airfield, error)
	PutAirfieldF     func(airfield []common.Airfield) error
	GetAirspaceF     func(regions []string, updatedSince time.Time) ([]common.Airspace, error)
	PutAirspaceF     func(airspace []common.Airspace) error
	GetFlightF       func(regions []string, updatedSince time.Time) ([]common.Flight, error)
	GetFlightFromIDF func(startID int, max int) ([]common.Flight, error)
	GetFlightByIDF   func(id int) (common.Flight, error)
	PutFlightF       func(flights []common.Flight) error
	GetWaypointF     func(regions []string, updatedSince time.Time) ([]common.Waypoint, error)
	PutWaypointF     func(waypoint []common.Waypoint) error
}

// Init is the mock implementation of common.Pluginer.Init.
func (mk *Mock) Init(cfg config.Config) error {
	if mk.InitF != nil {
		return mk.InitF(cfg)
	}
	return nil
}
