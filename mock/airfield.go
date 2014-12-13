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
	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/config"
	"time"
)

// Airfield is the mock implementation of Airfield.
// It provides a variation where you can pass the actual function implementation.
// Especially useful for testing.
type Airfield struct {
	GetF  func(regions []string, updatedSince time.Time) ([]common.Airfield, error)
	PutF  func(airfield []common.Airfield) error
	InitF func(cfg config.Config) error
}

// Init is the mock implementation of common.Pluginer.Init.
func (ma *Airfield) Init(cfg config.Config) error {
	return ma.InitF(cfg)
}

// GetAirfield is the mock implementation of common.Airfielder.GetAirfield.
func (ma *Airfield) GetAirfield(regions []string, updatedSince time.Time) ([]common.Airfield, error) {
	return ma.GetF(regions, updatedSince)
}

// PutAirfield is the mock implementation of common.Airfielder.PutAirfield.
func (ma *Airfield) PutAirfield(airfield []common.Airfield) error {
	return ma.PutF(airfield)
}
