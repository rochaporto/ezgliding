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

	"github.com/rochaporto/ezgliding/airfield"
)

// GetAirfield is the mock implementation of airfield.Airfielder.GetAirfield.
func (mk Mock) GetAirfield(regions []string, updatedSince time.Time) ([]airfield.Airfield, error) {
	if mk.GetAirfieldF != nil {
		return mk.GetAirfieldF(regions, updatedSince)
	}
	return []airfield.Airfield{}, nil
}

// PutAirfield is the mock implementation of airfield.Airfielder.PutAirfield.
func (mk Mock) PutAirfield(airfield []airfield.Airfield) error {
	if mk.PutAirfieldF != nil {
		return mk.PutAirfieldF(airfield)
	}
	return nil
}
