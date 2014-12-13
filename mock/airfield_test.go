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
	"github.com/rochaporto/ezgliding/plugin"
	"testing"
	"time"
)

func TestAirfieldInit(t *testing.T) {
	mockI := &Airfield{
		InitF: func(cfg config.Config) error {
			return nil
		},
	}
	x := plugin.Pluginer(mockI)
	err := x.Init(config.Config{})
	if err != nil {
		t.Errorf("Failed to call init on mock airfield")
	}
}

func TestGetAirfield(t *testing.T) {
	airfields := []common.Airfield{
		common.Airfield{Name: "TestMockAirfield"},
	}
	mockI := &Airfield{
		GetF: func(regions []string, updatedSince time.Time) ([]common.Airfield, error) {
			return airfields, nil
		},
	}
	x := common.Airfielder(mockI)
	result, err := x.GetAirfield(nil, time.Time{})
	if err != nil {
		t.Errorf("Failed to query mock airfields")
	}
	if len(result) != len(airfields) {
		t.Errorf("Got %v airfields but expected %v", len(result), len(airfields))
	}
}
func TestPutAirfield(t *testing.T) {
	mockI := &Airfield{
		PutF: func([]common.Airfield) error {
			return nil // FIXME: implement
		},
	}
	x := common.Airfielder(mockI)
	err := x.PutAirfield(nil) // FIXME: implement
	if err != nil {
		t.Errorf("Failed to put mock airfields")
	}
}
