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
	"testing"

	"github.com/rochaporto/ezgliding/config"
)

func TestInit(t *testing.T) {
	cfg := config.Config{}
	cfg.FusionTables = config.FusionTables{
		BaseURL: "some.random/location", AirfieldTableID: "1234",
		AirspaceTableID: "5678", WaypointTableID: "4321",
		APIKey: "myapikey", OAuthEmail: "oauthemail", OAuthKey: "oauthkey.test",
	}
	plugin := FusionTables{}
	err := plugin.Init(cfg)
	if err != nil {
		t.Errorf("Failed to initialize plugin :: %v", err)
	}

	result := config.FusionTables{BaseURL: plugin.BaseURL, AirspaceTableID: plugin.AirspaceTableID,
		AirfieldTableID: plugin.AirfieldTableID, WaypointTableID: plugin.WaypointTableID,
		APIKey: plugin.APIKey, OAuthEmail: plugin.OAuthEmail, OAuthKey: plugin.OAuthKey}
	if result != cfg.FusionTables {
		t.Errorf("Expected cfg %v but got %v", cfg.FusionTables, result)
	}

	if string(plugin.oAuthKeyContent) != "oauthkey test\n" {
		t.Errorf("expected %v got %v", "oauthkey test", string(plugin.oAuthKeyContent))
	}
}

func TestInitDefault(t *testing.T) {
	plugin := FusionTables{}
	err := plugin.Init(config.Config{})
	if err != nil {
		t.Errorf("Failed to initialize plugin :: %v", err)
	}

	if plugin.BaseURL != BaseURL {
		t.Errorf("Expected baseurl '%v' but got '%v'", BaseURL, plugin.BaseURL)
	}
}

func TestInitBadOAuthKeyLocation(t *testing.T) {
	plugin := FusionTables{}
	cfg := config.Config{}
	cfg.FusionTables = config.FusionTables{OAuthEmail: "myemail", OAuthKey: "nonexisting.file"}
	err := plugin.Init(cfg)
	if err == nil {
		t.Errorf("expected error got success")
	}
}

func TestDoGetError(t *testing.T) {
	plugin := FusionTables{}
	cfg := config.Config{}
	cfg.FusionTables = config.FusionTables{
		OAuthEmail: "myemail", OAuthKey: "nonexisting.file", BaseURL: "%%%.."}
	err := plugin.Init(cfg)
	_, err = plugin.doGet("")
	if err == nil {
		t.Errorf("expected error got success")
	}
}

func TestDoImportError(t *testing.T) {
	plugin := FusionTables{}
	cfg := config.Config{}
	cfg.FusionTables = config.FusionTables{
		OAuthEmail: "myemail", OAuthKey: "nonexisting.file", UploadURL: "%%%.."}
	err := plugin.Init(cfg)
	_, err = plugin.doImport("", "")
	if err == nil {
		t.Errorf("expected error got success")
	}
}

func TestDoOAuth(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer ts.Close()

	plugin := FusionTables{}
	cfg := config.Config{}
	cfg.FusionTables = config.FusionTables{
		AirfieldTableID: "testairfieldid", UploadURL: ts.URL,
		OAuthEmail: "myemail", OAuthKey: "oauthkey.test"}
	err := plugin.Init(cfg)
	if err != nil {
		t.Errorf("Failed to initialize plugin :: %v", err)
	}
	_, err = plugin.doImport("", cfg.FusionTables.AirfieldTableID)
	// FIXME: figure how to test oauth here
}
