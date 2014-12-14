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

// Package fusiontables provides the Airspace, Airfield and Waypoint plugin
// implementation for a fusion tables backend.
//
// It includes functionality for both pushing and retrieving information,
// allowing fusion tables to be used as a backend for the frontend applications.
//
// All requests use the Fusion Tables REST API, as defined in:
// 	https://developers.google.com/fusiontables/docs/v1/using
//
package fusiontables

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/golang/glog"
	"github.com/rochaporto/ezgliding/config"
)

const (
	// ID for this plugin implementation.
	ID string = "fusiontables"
	// BaseURL is the default base path for fusion tables REST queries
	BaseURL string = "https://www.googleapis.com/fusiontables/v2"
)

// FusionTables is the plugin implementation for a google fusion
// tables based backend.
type FusionTables struct {
	BaseURL         string
	AirspaceTableID string
	AirfieldTableID string
	WaypointTableID string
	APIKey          string
}

// Init follows the plugin.Plugin interface (see plugin.Pluginer).
func (ft *FusionTables) Init(cfg config.Config) error {
	glog.V(10).Infof("Init with config %+v", cfg.FusionTables)
	if cfg.FusionTables.Baseurl != "" {
		ft.BaseURL = cfg.FusionTables.Baseurl
	} else {
		ft.BaseURL = BaseURL
	}
	ft.AirfieldTableID = cfg.FusionTables.AirfieldTableID
	ft.AirspaceTableID = cfg.FusionTables.AirspaceTableID
	ft.WaypointTableID = cfg.FusionTables.WaypointTableID
	ft.APIKey = cfg.FusionTables.APIKey
	glog.V(20).Infof("Plugin fusion tables initialized :: %+v", ft)
	return nil
}

// get gives an utility function to wrap fusion tables queries with common
// required information (api key, base url, etc).
func (ft *FusionTables) get(sql string) (string, error) {
	var content []byte
	u, err := url.Parse(ft.BaseURL)
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("sql", sql)
	q.Set("key", ft.APIKey)
	q.Set("alt", "csv")
	u.RawQuery = q.Encode()
	r := u.String()
	glog.V(10).Infof("request %v", r)
	resp, err := http.Get(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, _ = ioutil.ReadAll(resp.Body)
	return string(content), nil
}
