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
	"errors"

	"github.com/rochaporto/ezgliding/config"
	"github.com/rochaporto/ezgliding/soaringweb"
	"github.com/rochaporto/ezgliding/welt2000"
)

// Pluginer is to be implemented by every plugin implementation
type Pluginer interface {
	Init(cfg config.Config) error
}

// pluginRegistry holds instances of available pluginRegistry mapped by IDs (for discovery).
var pluginRegistry = map[ID]Pluginer{
	ID(soaringweb.ID): Pluginer(&soaringweb.SoaringWeb{}),
	ID(welt2000.ID):   Pluginer(&welt2000.Welt2000{}),
}

// ID is the specific id of a given plugin.
type ID string

// Register adds a new Pluginer implementation to the available list.
// It fails if a plugin with the given ID already exists.
func Register(id ID, plugin Pluginer) error {
	_, present := pluginRegistry[id]
	if present {
		return errors.New("A plugin with ID " + string(id) + " already exists")
	}
	pluginRegistry[id] = plugin
	return nil
}

// NewPlugin returns a new Pluginer instance for the given plugin ID.
func NewPlugin(id ID) (Pluginer, error) {
	var result Pluginer
	result, present := pluginRegistry[id]
	if !present {
		return result, errors.New("No plugin available with id " + string(id))
	}
	return result, nil
}
