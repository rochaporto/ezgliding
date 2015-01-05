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

package plugin

import (
	"testing"
)

type TestPluginType Pluginer

func TestRegister(t *testing.T) {
	id := ID("testplugin")
	var x TestPluginType
	err := Register(id, x)
	if err != nil {
		t.Errorf("Failed to register new plugin with ID '%v' :: %v", id, err)
	}
	z, err := NewPlugin(id)
	if err != nil {
		t.Errorf("Failed to get new plugin instance of ID '%v' :: %v", id, err)
	}
	_, ok := z.(TestPluginType)
	if ok {
		t.Errorf("Got wrong plugin type after registering")
	}
}

func TestUnregister(t *testing.T) {
	id := ID("testplugin")
	err := Unregister(id)
	if err != nil {
		t.Errorf("failed to unregister plugin :: %v", err)
	}
}

func TestUnregisterMissing(t *testing.T) {
	id := ID("nonexisting")
	err := Unregister(id)
	if err == nil {
		t.Errorf("got success but should have failed")
	}
}

func TestRegisterExisting(t *testing.T) {
	id := ID("testpluginexisting")
	var TestPlugin Pluginer
	err := Register(id, TestPlugin)
	if err != nil {
		t.Errorf("Failed to register new plugin with ID '%v' :: %v", id, err)
	}
	err = Register(id, TestPlugin)
	if err == nil {
		t.Errorf("Succeded in registering plugin ID '%v' twice", id)
	}
}

func TestNewPluginMissingID(t *testing.T) {
	_, err := NewPlugin("nonexisting")
	if err == nil {
		t.Errorf("Got instance for missing ID")
	}
}
