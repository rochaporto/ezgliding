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

	"github.com/scalingdata/gcfg"
)

// FusionTables holds all config information for the google fusiontables plugin.
type FusionTables struct {
	AirfieldTableID string
	AirspaceTableID string
	WaypointTableID string
	BaseURL         string
	UploadURL       string
	APIKey          string
	OAuthEmail      string
	OAuthKey        string
}

// SoaringWeb holds all config information for the soaringweb plugin.
type SoaringWeb struct {
	Baseurl string
}

// Welt2000 holds all config information for the welt2000 plugin.
type Welt2000 struct {
	Rssurl     string
	Releaseurl string
}

// Global holds all common information for all ezgliding plugins and apps.
type Global struct {
	Airspacer  string
	Airfielder string
	Waypointer string
}

// Config holds all the config information for ezgliding plugins and apps.
type Config struct {
	Global
	FusionTables
	SoaringWeb
	Welt2000
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
