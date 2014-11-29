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
//
// Package mock provides mock implementations of all interfaces.

package mock

import (
	"github.com/rochaporto/ezgliding/common"
	"time"
)

// Airspace is the mock implementation of Airspace.
// It provides a variation where you can pass the actual function implementation.
// Especially useful for testing.
type Airspace struct {
	GetF  func(regions []string, updatedSince time.Time) ([]common.Airspace, error)
	PutF  func(airspace []common.Airspace) error
	InitF func(params map[string]string) error
}

// Init is the mock implementation of common.Pluginer.Init.
func (ma *Airspace) Init(params map[string]string) error {
	return ma.InitF(params)
}

// GetAirspace is the mock implementation of common.Airspacer.GetAirspace.
func (ma *Airspace) GetAirspace(regions []string, updatedSince time.Time) ([]common.Airspace, error) {
	return ma.GetF(regions, updatedSince)
}

// PutAirspace is the mock implementation of common.Airspacer.PutAirspace.
func (ma *Airspace) PutAirspace(airspace []common.Airspace) error {
	return ma.PutF(airspace)
}
