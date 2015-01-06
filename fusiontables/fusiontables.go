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
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/golang/glog"
	"github.com/rochaporto/ezgliding/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

const (
	// ID for this plugin implementation.
	ID string = "fusiontables"
	// BaseURL is the default base path for fusion tables REST queries
	BaseURL string = "https://www.googleapis.com/fusiontables/v2"
	// UploadURL is the default base path for fusion tables data imports
	UploadURL string = "https://www.googleapis.com/upload/fusiontables/v2"
)

// FusionTables is the plugin implementation for a google fusion
// tables based backend.
type FusionTables struct {
	BaseURL         string
	UploadURL       string
	AirspaceTableID string
	AirfieldTableID string
	WaypointTableID string
	APIKey          string
	OAuthEmail      string
	OAuthKey        string
	oAuthKeyContent []byte
}

// Init follows the plugin.Plugin interface (see plugin.Pluginer).
func (ft *FusionTables) Init(cfg config.Config) error {
	glog.V(10).Infof("Init with config %+v", cfg.FusionTables)
	if cfg.FusionTables.BaseURL != "" {
		ft.BaseURL = cfg.FusionTables.BaseURL
	} else {
		ft.BaseURL = BaseURL
	}
	if cfg.FusionTables.UploadURL != "" {
		ft.UploadURL = cfg.FusionTables.UploadURL
	} else {
		ft.UploadURL = UploadURL
	}
	ft.AirfieldTableID = cfg.FusionTables.AirfieldTableID
	ft.AirspaceTableID = cfg.FusionTables.AirspaceTableID
	ft.WaypointTableID = cfg.FusionTables.WaypointTableID
	ft.OAuthEmail = cfg.FusionTables.OAuthEmail
	ft.OAuthKey = cfg.FusionTables.OAuthKey
	ft.APIKey = cfg.FusionTables.APIKey

	// Load the private key contents if oauthkey location was given
	glog.V(10).Infof("plugin fusiontables initialized :: %v", ft)
	if ft.OAuthKey != "" {
		var err error
		ft.oAuthKeyContent, err = ioutil.ReadFile(ft.OAuthKey)
		if err != nil {
			return err
		}
	}
	return nil
}

// doGet wraps the given sql query into a REST call to fusion tables.
func (ft *FusionTables) doGet(sql string) (string, error) {
	u, err := url.Parse(ft.BaseURL + "/query")
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("sql", sql)
	q.Set("key", ft.APIKey)
	q.Set("alt", "csv")
	u.RawQuery = q.Encode()
	r := u.String()
	req, _ := http.NewRequest("GET", r, nil) // no err check (already above)
	return ft.do(req)
}

// doImport pushes the given content via a fusion tables REST import call.
func (ft *FusionTables) doImport(content string, tableID string) (string, error) {
	req, err := http.NewRequest("POST",
		ft.UploadURL+"/tables/"+tableID+"/import", strings.NewReader(content))
	if err != nil {
		return "", err
	}
	return ft.do(req)
}

// do performs the given request to fusion tables.
func (ft *FusionTables) do(req *http.Request) (string, error) {
	glog.V(10).Infof("http request %v", req)
	var resp *http.Response
	var err error

	req.Header.Add("Authorization", ft.APIKey)
	req.Header.Add("Content-Type", "application/octet-stream")
	if ft.OAuthEmail != "" && ft.OAuthKey != "" {
		glog.V(5).Infof("oauth for authentication :: id=%v, key=%v",
			ft.OAuthEmail, ft.OAuthKey)
		conf := &jwt.Config{
			Email:      ft.OAuthEmail,
			PrivateKey: ft.oAuthKeyContent,
			Scopes: []string{
				"https://www.googleapis.com/auth/fusiontables",
			},
			TokenURL: google.JWTTokenURL,
		}
		client := conf.Client(oauth2.NoContext)
		resp, err = client.Do(req)
	} else {
		client := http.Client{}
		resp, err = client.Do(req)
	}
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http %v: %v", resp.StatusCode, string(content))
	}
	return string(content), nil
}
