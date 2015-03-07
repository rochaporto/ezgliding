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

// Package config provides parsing and loading of configuration values.
//
// Includes definitions of plugins and formats to use, as well as the
// ability to provide any additional plugin specific configuration.
//
// Sample configuration:
//
// 	[global]
// 	airspacer=soaringweb
// 	airfielder=welt2000
// 	waypointer=welt2000
//
// 	[soaringweb]
// 	baseurl=http://soaringweb.org/Airspace
//
package config

import (
	"os"
	"os/user"

	"github.com/rochaporto/ezgliding/fusiontables"
	"github.com/rochaporto/ezgliding/mock"
	"github.com/rochaporto/ezgliding/netcoupe"
	"github.com/rochaporto/ezgliding/soaringweb"
	"github.com/rochaporto/ezgliding/web"
	"github.com/rochaporto/ezgliding/welt2000"
	"github.com/scalingdata/gcfg"
)

var empty = Config{}

var singleton Config

// Get returns the current config object
// FIXME: need to have a way to pass an alternative config location
func Get() (Config, error) {
	if singleton == empty {
		singleton, _ = NewConfig("")
	}
	return singleton, nil
}

// Global holds all common information for all ezgliding plugins and apps.
type Global struct {
	Airspacer  string
	Airfielder string
	Flighter   string
	Waypointer string
}

// Config holds all the config information for ezgliding plugins and apps.
type Config struct {
	Global       Global
	FusionTables fusiontables.Config
	Mock         mock.Config
	Netcoupe     netcoupe.Config
	SoaringWeb   soaringweb.Config
	Web          web.Config
	Welt2000     welt2000.Config
}

// NewConfig returns a new Config based on given location (or default).
//
// If zero value ("") location is given, then the default locations are tried:
// 	./ezgliding.cfg, ~/.ezgliding.cfg, ~/.ezgliding
func NewConfig(location string) (Config, error) {
	cfg := Config{}
	var err error
	locations := []string{location}
	if location == "" {
		usr, _ := user.Current()
		locations = append(locations, "./ezgliding.cfg", usr.HomeDir+"/.ezgliding.cfg", usr.HomeDir+"/.ezgliding")
	}
	for _, l := range locations {
		_, err = os.Stat(l)
		if err == nil {
			err = gcfg.ReadFileInto(&cfg, l)
			break
		}
	}
	return cfg, err
}

// Set sets the current configuration to the given one.
func Set(cfg Config) {
	singleton = cfg
}
