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
package fusiontables

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	cfg := Config{
		BaseURL: "some.random/location", AirfieldTableID: "1234",
		AirspaceTableID: "5678", WaypointTableID: "4321",
		APIKey: "myapikey", OAuthEmail: "oauthemail", OAuthKey: "oauthkey.test",
	}
	plugin, err := New(cfg)
	if err != nil {
		t.Errorf("failed to get new plugin :: %v", err)
		return
	}

	result := Config{BaseURL: plugin.BaseURL, AirspaceTableID: plugin.AirspaceTableID,
		AirfieldTableID: plugin.AirfieldTableID, WaypointTableID: plugin.WaypointTableID,
		APIKey: plugin.APIKey, OAuthEmail: plugin.OAuthEmail, OAuthKey: plugin.OAuthKey}
	if !reflect.DeepEqual(result, cfg) {
		t.Errorf("expected\n%v\ngot\n%v", cfg, result)
		return
	}

	if string(plugin.oAuthKeyContent) != "oauthkey test\n" {
		t.Errorf("expected\n%v\ngot\n%v", "oauthkey test", string(plugin.oAuthKeyContent))
		return
	}
}

func TestNewDefault(t *testing.T) {
	plugin, err := New(Config{})
	if err != nil {
		t.Errorf("failed to get new plugin :: %v", err)
	}

	if plugin.BaseURL != BaseURL {
		t.Errorf("expected baseurl '%v' but got '%v'", BaseURL, plugin.BaseURL)
	}
}

func TestNewBadOAuthKeyLocation(t *testing.T) {
	cfg := Config{OAuthEmail: "myemail", OAuthKey: "nonexisting.file"}
	_, err := New(cfg)
	if err == nil {
		t.Errorf("expected error got success")
	}
}

func TestDoGetError(t *testing.T) {
	cfg := Config{
		OAuthEmail: "myemail", OAuthKey: "nonexisting.file", BaseURL: "%%%.."}
	plugin, err := New(cfg)
	_, err = plugin.doGet("")
	if err == nil {
		t.Errorf("expected error got success")
	}
}

func TestDoImportError(t *testing.T) {
	cfg := Config{
		OAuthEmail: "myemail", OAuthKey: "nonexisting.file", UploadURL: "%%%.."}
	plugin, err := New(cfg)
	_, err = plugin.doImport("", "")
	if err == nil {
		t.Errorf("expected error got success")
	}
}

func TestDoOAuth(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer ts.Close()

	cfg := Config{
		AirfieldTableID: "testairfieldid", UploadURL: ts.URL,
		OAuthEmail: "myemail", OAuthKey: "oauthkey.test"}
	plugin, err := New(cfg)
	if err != nil {
		t.Errorf("Failed to initialize plugin :: %v", err)
		return
	}
	_, err = plugin.doImport("", cfg.AirfieldTableID)
	// FIXME: figure how to test oauth here
}
