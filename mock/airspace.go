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

// GetAirspace is the mock implementation of common.Airspacer.GetAirspace.
func (mk *Mock) GetAirspace(regions []string, updatedSince time.Time) ([]common.Airspace, error) {
	if mk.GetAirspaceF != nil {
		return mk.GetAirspaceF(regions, updatedSince)
	}
	return []common.Airspace{}, nil
}

// PutAirspace is the mock implementation of common.Airspacer.PutAirspace.
func (mk *Mock) PutAirspace(airspace []common.Airspace) error {
	if mk.PutAirspaceF != nil {
		return mk.PutAirspaceF(airspace)
	}
	return nil
}
