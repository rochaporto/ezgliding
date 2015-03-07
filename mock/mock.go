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

	"github.com/rochaporto/ezgliding/airfield"
	"github.com/rochaporto/ezgliding/airspace"
	"github.com/rochaporto/ezgliding/flight"
	"github.com/rochaporto/ezgliding/waypoint"
)

const (
	// ID for this plugin implementation.
	ID string = "mock"
)

// Config holds the Mock configuration.
type Config struct{}

// Mock implements all interfaces (airfield, airspace, waypoint).
// Use the struct fields to provide the function implementations.
type Mock struct {
	Config
	GetAirfieldF     func(regions []string, updatedSince time.Time) ([]airfield.Airfield, error)
	PutAirfieldF     func(afield []airfield.Airfield) error
	GetAirspaceF     func(regions []string, updatedSince time.Time) ([]airspace.Airspace, error)
	PutAirspaceF     func(aspace []airspace.Airspace) error
	GetFlightF       func(regions []string, updatedSince time.Time) ([]flight.Flight, error)
	GetFlightFromIDF func(startID int, max int) ([]flight.Flight, error)
	GetFlightByIDF   func(id int) (flight.Flight, error)
	PutFlightF       func(flights []flight.Flight) error
	GetWaypointF     func(regions []string, updatedSince time.Time) ([]waypoint.Waypoint, error)
	PutWaypointF     func(waypoint []waypoint.Waypoint) error
}

// New returns a new instance of Mock.
func New(cfg Config) (*Mock, error) {
	mk := Mock{}
	return &mk, nil
}
