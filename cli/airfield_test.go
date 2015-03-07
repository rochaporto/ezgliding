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

package cli

import (
	"errors"
	"flag"
	"testing"
	"time"

	"github.com/rochaporto/ezgliding/airfield"
	"github.com/rochaporto/ezgliding/config"
	"github.com/rochaporto/ezgliding/mock"
	"github.com/rochaporto/ezgliding/plugin"
)

// ExampleAirfieldGet uses the mock airfield implementation to query data and
// verify airfield-get works. First, no region is passed. Second, a region but
// no updatedAfter is passed. Finally, both region and updatedAfter are given.
func ExampleAirfieldGet() {
	plugin.Register("mockairfieldget", &mock.Mock{
		GetAirfieldF: func(regions []string, updatedSince time.Time) ([]airfield.Airfield, error) {
			airfields := []airfield.Airfield{
				airfield.Airfield{
					ID: "MockID1", ShortName: "MockShortName",
					Name: "MockName", Region: "FR", ICAO: "AAAA", Flags: 0,
					Catalog: 11, Length: 1000, Elevation: 2000, Runway: "32R",
					Frequency: 123.45, Latitude: 32.533, Longitude: 100.376,
					Update: time.Date(2014, 02, 02, 0, 0, 0, 0, time.UTC)},
				airfield.Airfield{
					ID: "MockID2", ShortName: "MockShortName",
					Name: "MockName", Region: "CH", ICAO: "AAAA", Flags: 0,
					Catalog: 11, Length: 1000, Elevation: 2000, Runway: "32R",
					Frequency: 123.45, Latitude: 32.533, Longitude: 100.376,
					Update: time.Date(2014, 02, 02, 0, 0, 0, 0, time.UTC)},
				airfield.Airfield{
					ID: "MockID3", ShortName: "MockShortName",
					Name: "MockName", Region: "CH", ICAO: "AAAA", Flags: 0,
					Catalog: 11, Length: 1000, Elevation: 2000, Runway: "32R",
					Frequency: 123.45, Latitude: 32.533, Longitude: 100.376,
					Update: time.Date(2014, 02, 03, 0, 0, 0, 0, time.UTC)},
			}
			result := []airfield.Airfield{}
			for _, airfield := range airfields {
				b := false
				for _, r := range regions {
					if airfield.Region == r {
						b = true
					}
				}
				if airfield.Update.After(updatedSince) && b {
					result = append(result, airfield)
				}
			}
			return result, nil
		},
	},
	)
	config.Set(config.Config{Global: config.Global{Airfielder: "mockairfieldget"}})
	_ = flag.Set("after", "2014-02-02")
	_ = flag.Set("region", "CH")
	runAirfieldGet(CmdAirfieldGet, []string{})
	// Output:
	// ID,ShortName,Name,Region,ICAO,Flags,Catalog,Length,Elevation,Runway,Frequency,Latitude,Longitude,Update
	// MockID3,MockShortName,MockName,CH,AAAA,0,11,1000,2000,32R,123.45,32.533,100.376,2014-02-03 00:00:00 +0000 UTC
}

func TestAirfieldGetFailed(t *testing.T) {
	plugin.Register("mockairfieldgetfailed", &mock.Mock{
		GetAirfieldF: func(regions []string, updatedSince time.Time) ([]airfield.Airfield, error) {
			return nil, errors.New("mock testing get airfield failed")
		},
	},
	)
	config.Set(config.Config{Global: config.Global{Airfielder: "mockairfieldgetfailed"}})
	flag.Set("after", "")
	flag.Set("region", "")
	runAirfieldGet(CmdAirfieldGet, []string{})
	// Output:
	// failed to get airfield :: mock testing get airfield failed
}

func TestAirfieldGetBadAfter(t *testing.T) {
	plugin.Register("mockairfieldbadafter", &mock.Mock{
		GetAirfieldF: func(regions []string, updatedSince time.Time) ([]airfield.Airfield, error) {
			return nil, nil
		},
	},
	)
	config.Set(config.Config{Global: config.Global{Airfielder: "mockairfieldbadafter"}})
	flag.Set("after", "22-00-11")
	flag.Set("region", "")
	runAirfieldGet(CmdAirfieldGet, []string{})
}

func TestAirfieldGetBadPluginID(t *testing.T) {
	config.Set(config.Config{Global: config.Global{Airfielder: "mockairfieldnonexisting"}})
	runAirfieldGet(CmdAirfieldGet, []string{})
}

// ExampleAirfieldPut uses the mock airfield implementation to push data and
// verify airfield-put works.
func ExampleAirfieldPut() {
	plugin.Register("mockairfield", &mock.Mock{
		GetAirfieldF: func(regions []string, updatedSince time.Time) ([]airfield.Airfield, error) {
			return []airfield.Airfield{
				airfield.Airfield{ID: "MockID", ShortName: "MockShortName", Name: "MockName",
					Region: "FR", ICAO: "AAAA", Flags: 0, Catalog: 11, Length: 1000, Elevation: 2000,
					Runway: "32R", Frequency: 123.45, Latitude: 32.533, Longitude: 100.376,
					Update: time.Time{}},
			}, nil
		},
	},
	)
	plugin.Register("mockairfieldput", &mock.Mock{
		PutAirfieldF: func(airfields []airfield.Airfield) error {
			return nil
		},
	},
	)
	config.Set(config.Config{Global: config.Global{Airfielder: "mockairfield"}})
	runAirfieldPut(CmdAirfieldPut, []string{"mockairfieldput"})
	// Output:
	// pushed 1 airfields into mockairfieldput
}

func TestAirfieldPutMissingArg(t *testing.T) {
	runAirfieldPut(CmdAirfieldPut, []string{})
}

func TestAirfieldPutBadGet(t *testing.T) {
	plugin.Register("mockairfieldput", &mock.Mock{})
	plugin.Register("mockairfieldputbadget", &mock.Mock{
		GetAirfieldF: func(regions []string, updatedSince time.Time) ([]airfield.Airfield, error) {
			return nil, errors.New("mock testing get airfield failed")
		},
	},
	)
	config.Set(config.Config{Global: config.Global{Airfielder: "mockairfieldputbadget"}})
	runAirfieldPut(CmdAirfieldPut, []string{"mockairfieldput"})
	// Output:
	// failed to get airfield :: mock testing get airfield failed
}

func TestAirfieldPutFailed(t *testing.T) {
	plugin.Register("mockairfieldget", &mock.Mock{
		GetAirfieldF: func(regions []string, updatedSince time.Time) ([]airfield.Airfield, error) {
			return []airfield.Airfield{
				airfield.Airfield{ID: "MockID", ShortName: "MockShortName", Name: "MockName",
					Region: "FR", ICAO: "AAAA", Flags: 0, Catalog: 11, Length: 1000, Elevation: 2000,
					Runway: "32R", Frequency: 123.45, Latitude: 32.533, Longitude: 100.376,
					Update: time.Time{}},
			}, nil
		},
	})
	plugin.Register("mockairfieldputfailed", &mock.Mock{
		PutAirfieldF: func(airfields []airfield.Airfield) error {
			return errors.New("mock testing put airfield failed")
		},
	})
	config.Set(config.Config{Global: config.Global{Airfielder: "mockairfieldget"}})
	runAirfieldPut(CmdAirfieldPut, []string{"mockairfieldputfailed"})
	// Output:
	// failed to put airfield :: mock testing put airfield failed
}

func TestAirfieldPutBadPluginID(t *testing.T) {
	runAirfieldPut(CmdAirfieldPut, []string{"afnonexisting"})
}
