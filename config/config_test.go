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

package config

import (
	"reflect"
	"testing"
)

func TestConfigGetDefault(t *testing.T) {
	e := Config{Global: Global{
		Airspacer: "airspacerdefault", Airfielder: "airfielderdefault", Waypointer: "waypointerdefault",
	}}
	result, err := Get()
	if err != nil {
		t.Errorf("failed to get config :: %v", err)
		return
	}
	if !reflect.DeepEqual(result, e) {
		t.Errorf("expected\n%v\ngot\n%v", e, result)
		return
	}
}

func TestConfigGet(t *testing.T) {
	cfg := Config{Global: Global{Airspacer: "airspacerconfigget"}}
	Set(cfg)
	result, err := Get()
	if err != nil {
		t.Errorf("failed to get config :: %v", err)
		return
	}
	if !reflect.DeepEqual(result, cfg) {
		t.Errorf("expected\n%v\ngot\n%v", cfg, result)
		return
	}
}
func TestNewConfig(t *testing.T) {
	e := Config{}
	e.Global.Airspacer = "airspacer"
	e.Global.Airfielder = "airfielder"
	e.Global.Waypointer = "waypointer"
	cfg, err := NewConfig("ezgliding-simpleconfig")
	if err != nil {
		t.Errorf("Failed to get new config :: %v", err)
	}
	if cfg != e {
		t.Errorf("Expected %v but got %v", e, cfg)
	}
}

func TestNewConfigDefault(t *testing.T) {
	e := Config{}
	e.Global.Airspacer = "airspacerdefault"
	e.Global.Airfielder = "airfielderdefault"
	e.Global.Waypointer = "waypointerdefault"
	cfg, err := NewConfig("")
	if err != nil {
		t.Errorf("Failed to get new config :: %v", err)
	}
	if cfg != e {
		t.Errorf("Expected %v but got %v", e, cfg)
	}
}

func TestNewConfigMissing(t *testing.T) {
	_, err := NewConfig("nonexistingconfig")
	if err == nil {
		t.Errorf("Got no error when loading non existing config file")
	}
}
