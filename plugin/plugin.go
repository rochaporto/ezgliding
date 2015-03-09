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

// Package plugin holds plugin registry functionality.
//
// All plugins must be statically defined here as there is no dynamic
// loading functionality available.
//
// Plugins should all implement Pluginer, and should then implement one
// or several of the interfaces defined in package common. They should
// also set their ID, which must be unique and should look like:
// 	const (
//		ID string = "welt2000"
//	)
//
// Once written plugins should be added to the the pluginRegistry,
// along with a zero value wrapped in Pluginer.
//
// In the future there might be a Plugin implementation which would allow
// running plugins as external processes, with rpc calls.
//
package plugin

import (
	"fmt"

	"github.com/rochaporto/ezgliding/airfield"
	"github.com/rochaporto/ezgliding/airspace"
	"github.com/rochaporto/ezgliding/config"
	"github.com/rochaporto/ezgliding/flight"
	"github.com/rochaporto/ezgliding/fusiontables"
	"github.com/rochaporto/ezgliding/netcoupe"
	"github.com/rochaporto/ezgliding/soaringweb"
	"github.com/rochaporto/ezgliding/waypoint"
	"github.com/rochaporto/ezgliding/welt2000"
)

// GetFlighter returns an instance of the request Flighter plugin.
// Passing an empty string will try to load an instance of the plugin
// currently set in the configuration (if any).
func GetFlighter(id string, cfg config.Config) (flight.Flighter, error) {
	fid := id
	if fid == "" {
		fid = cfg.Global.Flighter
	}
	r, err := GetInstance(fid, cfg)
	if err != nil {
		return nil, err
	}
	return r.(flight.Flighter), err
}

// GetAirfielder returns an instance of the request Airfielder plugin.
// Passing an empty string will try to load an instance of the plugin
// currently set in the configuration (if any).
func GetAirfielder(id string, cfg config.Config) (airfield.Airfielder, error) {
	fid := id
	if fid == "" {
		fid = cfg.Global.Airfielder
	}
	r, err := GetInstance(fid, cfg)
	if err != nil {
		return nil, err
	}
	return r.(airfield.Airfielder), err
}

// GetAirspacer returns an instance of the request Airspacer plugin.
// Passing an empty string will try to load an instance of the plugin
// currently set in the configuration (if any).
func GetAirspacer(id string, cfg config.Config) (airspace.Airspacer, error) {
	fid := id
	if fid == "" {
		fid = cfg.Global.Airspacer
	}
	r, err := GetInstance(fid, cfg)
	if err != nil {
		return nil, err
	}
	return r.(airspace.Airspacer), err
}

// GetWaypointer returns an instance of the request Waypointer plugin.
// Passing an empty string will try to load an instance of the plugin
// currently set in the configuration (if any).
func GetWaypointer(id string, cfg config.Config) (waypoint.Waypointer, error) {
	fid := id
	if fid == "" {
		fid = cfg.Global.Waypointer
	}
	r, err := GetInstance(fid, cfg)
	if err != nil {
		return nil, err
	}
	return r.(waypoint.Waypointer), err
}

// GetScorer returns an instance of the requested Scorer plugin.
// Passing an empty string will try to load an instance of the plugin
// currently set in the configuration (if any).
func GetScorer(id string, cfg config.Config) (flight.Scorer, error) {
	fid := id
	if fid == "" {
		fid = cfg.Global.Scorer
	}
	r, err := GetInstance(fid, cfg)
	if err != nil {
		return nil, err
	}
	return r.(flight.Scorer), err
}

// registry holds additional string/plugin mappings to the default ones.
// The default ones are set in the GetInstance function.
var registry = map[string]interface{}{}

// Register adds a new plugin mapping to the available list.
// This is useful to register mock plugins (probably based on the mock.Mock
// struct) for testing.
func Register(id string, plugin interface{}) {
	registry[id] = plugin
}

// GetInstance returns a new instance of the requested plugin.
func GetInstance(id string, cfg config.Config) (interface{}, error) {
	switch id {
	case "fusiontables":
		ft, _ := fusiontables.New(cfg.FusionTables)
		return ft, nil
	case "netcoupe":
		nc, _ := netcoupe.New(cfg.Netcoupe)
		return nc, nil
	case "soaringweb":
		sw, _ := soaringweb.New(cfg.SoaringWeb)
		return sw, nil
	case "welt2000":
		wt, _ := welt2000.New(cfg.Welt2000)
		return wt, nil
	default:
		p, ok := registry[id]
		if !ok {
			return nil, fmt.Errorf("unknown plugin id :: %v", id)
		}
		return p, nil
	}
}
