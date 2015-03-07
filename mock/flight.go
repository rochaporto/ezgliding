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

package mock

import (
	"time"

	"github.com/rochaporto/ezgliding/flight"
)

// GetFlight is the mock implementation of flight.Flighter.GetFlight.
func (mk Mock) GetFlight(regions []string, updatedSince time.Time) ([]flight.Flight, error) {
	if mk.GetFlightF != nil {
		return mk.GetFlightF(regions, updatedSince)
	}
	return []flight.Flight{}, nil
}

// GetFlightFromID is the mock implementation of flight.Flighter.GetFlightFromID.
func (mk Mock) GetFlightFromID(startID int, max int) ([]flight.Flight, error) {
	if mk.GetFlightFromIDF != nil {
		return mk.GetFlightFromIDF(startID, max)
	}
	return []flight.Flight{}, nil
}

// GetFlightByID is the mock implementation of flight.Flighter.GetFlightByID.
func (mk Mock) GetFlightByID(id int) (flight.Flight, error) {
	if mk.GetFlightByIDF != nil {
		return mk.GetFlightByIDF(id)
	}
	return flight.Flight{}, nil
}

// PutFlight is the mock implementation of flight.Flighter.PutFlight.
func (mk Mock) PutFlight(f []flight.Flight) error {
	if mk.PutFlightF != nil {
		return mk.PutFlightF(f)
	}
	return nil
}
