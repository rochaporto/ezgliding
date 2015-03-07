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
	"reflect"
	"testing"

	"github.com/rochaporto/ezgliding/config"
	"github.com/rochaporto/ezgliding/fusiontables"
	"github.com/rochaporto/ezgliding/mock"
	"github.com/rochaporto/ezgliding/netcoupe"
	"github.com/rochaporto/ezgliding/soaringweb"
	"github.com/rochaporto/ezgliding/welt2000"
)

func TestGetFlighter(t *testing.T) {
	m := mock.Mock{}
	Register("mockflighter", &m)
	f, err := GetFlighter("mockflighter", config.Config{Global: config.Global{Flighter: "default"}})
	if err != nil {
		t.Errorf("failed to get fligher :: %v", err)
	}
	if !reflect.DeepEqual(f, &m) {
		t.Errorf("expected %v got %v", &m, f)
	}
}

func TestGetFlighterDefault(t *testing.T) {
	m := mock.Mock{}
	Register("default", &m)
	f, err := GetFlighter("", config.Config{Global: config.Global{Flighter: "default"}})
	if err != nil {
		t.Errorf("failed to get fligher :: %v", err)
	}
	if !reflect.DeepEqual(f, &m) {
		t.Errorf("expected %v got %v", &m, f)
	}
}

func TestGetFlighterMissing(t *testing.T) {
	_, err := GetFlighter("", config.Config{Global: config.Global{Flighter: "nonexisting"}})
	if err == nil {
		t.Errorf("expected error but got success")
	}
}

func TestGetAirfielder(t *testing.T) {
	m := mock.Mock{}
	Register("mockairfielder", &m)
	a, err := GetAirfielder("mockairfielder", config.Config{Global: config.Global{Airfielder: "default"}})
	if err != nil {
		t.Errorf("failed to get airfielder :: %v", err)
	}
	if !reflect.DeepEqual(a, &m) {
		t.Errorf("expected %v got %v", &m, a)
	}
}

func TestGetAirfielderDefault(t *testing.T) {
	m := mock.Mock{}
	Register("default", &m)
	a, err := GetAirfielder("", config.Config{Global: config.Global{Airfielder: "default"}})
	if err != nil {
		t.Errorf("failed to get airfielder :: %v", err)
	}
	if !reflect.DeepEqual(a, &m) {
		t.Errorf("expected %v got %v", &m, a)
	}
}

func TestGetAirfielderMissing(t *testing.T) {
	_, err := GetAirfielder("", config.Config{Global: config.Global{Airfielder: "nonexisting"}})
	if err == nil {
		t.Errorf("expected error but got success")
	}
}

func TestGetAirspacer(t *testing.T) {
	m := mock.Mock{}
	Register("mockairspacer", &m)
	a, err := GetAirspacer("mockairspacer", config.Config{Global: config.Global{Airspacer: "default"}})
	if err != nil {
		t.Errorf("failed to get airfielder :: %v", err)
	}
	if !reflect.DeepEqual(a, &m) {
		t.Errorf("expected %v got %v", &m, a)
	}
}

func TestGetAirspacerDefault(t *testing.T) {
	m := mock.Mock{}
	Register("default", &m)
	a, err := GetAirspacer("", config.Config{Global: config.Global{Airspacer: "default"}})
	if err != nil {
		t.Errorf("failed to get airfielder :: %v", err)
	}
	if !reflect.DeepEqual(a, &m) {
		t.Errorf("expected %v got %v", &m, a)
	}
}

func TestGetAirspacerMissing(t *testing.T) {
	_, err := GetAirspacer("", config.Config{Global: config.Global{Airspacer: "nonexisting"}})
	if err == nil {
		t.Errorf("expected error but got success")
	}
}

func TestGetWaypointer(t *testing.T) {
	m := mock.Mock{}
	Register("mockwaypointer", &m)
	a, err := GetWaypointer("mockwaypointer", config.Config{Global: config.Global{Waypointer: "default"}})
	if err != nil {
		t.Errorf("failed to get airfielder :: %v", err)
	}
	if !reflect.DeepEqual(a, &m) {
		t.Errorf("expected %v got %v", &m, a)
	}
}

func TestGetWaypointerDefault(t *testing.T) {
	m := mock.Mock{}
	Register("default", &m)
	a, err := GetWaypointer("", config.Config{Global: config.Global{Waypointer: "default"}})
	if err != nil {
		t.Errorf("failed to get airfielder :: %v", err)
	}
	if !reflect.DeepEqual(a, &m) {
		t.Errorf("expected %v got %v", &m, a)
	}
}

func TestGetWaypointerMissing(t *testing.T) {
	_, err := GetWaypointer("", config.Config{Global: config.Global{Waypointer: "nonexisting"}})
	if err == nil {
		t.Errorf("expected error but got success")
	}
}

func TestGetInstanceFusionTables(t *testing.T) {
	e, _ := fusiontables.New(fusiontables.Config{})
	r, err := GetInstance("fusiontables", config.Config{})
	if err != nil {
		t.Errorf("failed to get instance :: %v", err)
		return
	}
	if !reflect.DeepEqual(r, e) {
		t.Errorf("expected %v but got %v", e, r)
	}
}

func TestGetInstanceNetcoupe(t *testing.T) {
	e, _ := netcoupe.New(netcoupe.Config{})
	r, err := GetInstance("netcoupe", config.Config{})
	if err != nil {
		t.Errorf("failed to get instance :: %v", err)
		return
	}
	if !reflect.DeepEqual(r, e) {
		t.Errorf("expected %v but got %v", e, r)
	}
}

func TestGetInstanceSoaringWeb(t *testing.T) {
	e, _ := soaringweb.New(soaringweb.Config{})
	r, err := GetInstance("soaringweb", config.Config{})
	if err != nil {
		t.Errorf("failed to get instance :: %v", err)
		return
	}
	if !reflect.DeepEqual(r, e) {
		t.Errorf("expected %v but got %v", e, r)
	}
}

func TestGetInstanceWelt2000(t *testing.T) {
	e, _ := welt2000.New(welt2000.Config{})
	r, err := GetInstance("welt2000", config.Config{})
	if err != nil {
		t.Errorf("failed to get instance :: %v", err)
		return
	}
	if !reflect.DeepEqual(r, e) {
		t.Errorf("expected %v but got %v", e, r)
	}
}
