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

	"github.com/rochaporto/ezgliding/common"
)

// GetFlight is the mock implementation of common.Flighter.GetFlight.
func (mk *Mock) GetFlight(regions []string, updatedSince time.Time) ([]common.Flight, error) {
	if mk.GetFlightF != nil {
		return mk.GetFlightF(regions, updatedSince)
	}
	return []common.Flight{}, nil
}

// GetFlightFromID is the mock implementation of common.Flighter.GetFlightFromID.
func (mk *Mock) GetFlightFromID(startID int, max int) ([]common.Flight, error) {
	if mk.GetFlightFromIDF != nil {
		return mk.GetFlightFromIDF(startID, max)
	}
	return []common.Flight{}, nil
}

// GetFlightByID is the mock implementation of common.Flighter.GetFlightByID.
func (mk *Mock) GetFlightByID(id int) (common.Flight, error) {
	if mk.GetFlightByIDF != nil {
		return mk.GetFlightByIDF(id)
	}
	return common.Flight{}, nil
}

// PutFlight is the mock implementation of common.Flighter.PutFlight.
func (mk *Mock) PutFlight(flight []common.Flight) error {
	if mk.PutFlightF != nil {
		return mk.PutFlightF(flight)
	}
	return nil
}
