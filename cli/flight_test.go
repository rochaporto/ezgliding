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

	"github.com/rochaporto/ezgliding/config"
	"github.com/rochaporto/ezgliding/flight"
	"github.com/rochaporto/ezgliding/mock"
	"github.com/rochaporto/ezgliding/plugin"
)

// ExampleFlightGetByID uses the mock flight implementation to query data and
// verify flight-get works. First, a specific id is passed. Second, a start id
// is passed.
func ExampleFlightGetByID() {
	plugin.Register("mockflightgetbyid", &mock.Mock{
		GetFlightByIDF: func(id int) (flight.Flight, error) {
			f := flight.Flight{
				Header: flight.Header{
					Date:       time.Date(2015, 1, 10, 0, 0, 0, 0, time.UTC),
					Pilot:      "MOCK PILOT 1",
					GliderType: "MOCK GLIDER 1", GliderID: "MOCK ID 1",
				},
				Sources: map[string]flight.Source{
					"netcoupe": flight.Source{
						Name:     "MOCK NAME 1",
						Category: "MOCK CATEGORY 1",
						Club:     "MOCK CLUB 1",
						Region:   "MOCK REGION 1",
						Country:  "MOCK COUNTRY 1",
						Distance: 100.01, Points: 100.02,
					},
				},
			}
			return f, nil
		},
	},
	)
	config.Set(config.Config{Global: config.Global{Flighter: "mockflightgetbyid"}})
	_ = flag.Set("id", "1")
	_ = flag.Set("startID", "")
	_ = flag.Set("max", "")
	runFlightGet(CmdFlightGet, []string{})
	// Output:
	// 10/01/2015,MOCK PILOT 1,MOCK GLIDER 1,MOCK ID 1
	//	netcoupe,MOCK NAME 1,MOCK CATEGORY 1,MOCK CLUB 1,MOCK COUNTRY 1,MOCK REGION 1,100.01,100.02
}

// ExampleFlightGetBadID tests giving a non integer as ID value, with null output
func ExampleFlightGetBadID() {
	_ = flag.Set("id", "i")
	_ = flag.Set("startID", "")
	_ = flag.Set("max", "")
	runFlightGet(CmdFlightGet, []string{})
	// Output:
}

// ExampleFlightGetMissingID tests giving a non existing ID value, with null output
func ExampleFlightGetMissingID() {
	plugin.Register("mockflightmissingid", &mock.Mock{
		GetFlightByIDF: func(id int) (flight.Flight, error) {
			return flight.Flight{}, errors.New("given id does not exist")
		},
	},
	)
	config.Set(config.Config{Global: config.Global{Flighter: "mockflightmissingid"}})
	_ = flag.Set("id", "9")
	_ = flag.Set("startID", "")
	_ = flag.Set("max", "")
	runFlightGet(CmdFlightGet, []string{})
	// Output:
}

// ExampleFlightGetFromID queries using flight-get. First no max is given, so all records
// starting at 'startID' are returned. Second, a max is given so the number of results is limited.
func ExampleFlightGetFromID() {
	plugin.Register("mockflightgetfromid", &mock.Mock{
		GetFlightFromIDF: func(startID int, max int) ([]flight.Flight, error) {
			flights := []flight.Flight{
				flight.Flight{
					Header: flight.Header{
						Date:       time.Date(2015, 1, 10, 0, 0, 0, 0, time.UTC),
						Pilot:      "MOCK PILOT 1",
						GliderType: "MOCK GLIDER 1", GliderID: "MOCK ID 1",
					},
					Sources: map[string]flight.Source{
						"netcoupe": flight.Source{
							Name:     "MOCK NAME 1",
							Category: "MOCK CATEGORY 1",
							Club:     "MOCK CLUB 1",
							Region:   "MOCK REGION 1",
							Country:  "MOCK COUNTRY 1",
							Distance: 101.01, Points: 101.02,
						},
					},
				},
				flight.Flight{
					Header: flight.Header{
						Date:       time.Date(2015, 1, 11, 0, 0, 0, 0, time.UTC),
						Pilot:      "MOCK PILOT 2",
						GliderType: "MOCK GLIDER 2", GliderID: "MOCK ID 2",
					},
					Sources: map[string]flight.Source{
						"netcoupe": flight.Source{
							Name:     "MOCK NAME 2",
							Category: "MOCK CATEGORY 2",
							Club:     "MOCK CLUB 2",
							Region:   "MOCK REGION 2",
							Country:  "MOCK COUNTRY 2",
							Distance: 102.01, Points: 102.02,
						},
					},
				},
			}
			if max > 0 {
				flights = flights[0:max]
			}
			return flights, nil
		},
	},
	)
	config.Set(config.Config{Global: config.Global{Flighter: "mockflightgetfromid"}})
	_ = flag.Set("startID", "1")
	_ = flag.Set("max", "")
	runFlightGet(CmdFlightGet, []string{})
	_ = flag.Set("max", "1")
	runFlightGet(CmdFlightGet, []string{})
	// Output:
	// 10/01/2015,MOCK PILOT 1,MOCK GLIDER 1,MOCK ID 1
	//	netcoupe,MOCK NAME 1,MOCK CATEGORY 1,MOCK CLUB 1,MOCK COUNTRY 1,MOCK REGION 1,101.01,101.02
	// 11/01/2015,MOCK PILOT 2,MOCK GLIDER 2,MOCK ID 2
	//	netcoupe,MOCK NAME 2,MOCK CATEGORY 2,MOCK CLUB 2,MOCK COUNTRY 2,MOCK REGION 2,102.01,102.02
	// 10/01/2015,MOCK PILOT 1,MOCK GLIDER 1,MOCK ID 1
	//	netcoupe,MOCK NAME 1,MOCK CATEGORY 1,MOCK CLUB 1,MOCK COUNTRY 1,MOCK REGION 1,101.01,101.02
}

// ExampleFlightGetBadStartID tests giving a non integer as startID value, with null output
func ExampleFlightGetBadStartID() {
	plugin.Register("mockflight", &mock.Mock{})
	config.Set(config.Config{Global: config.Global{Flighter: "mockflight"}})
	_ = flag.Set("startID", "s")
	runFlightGet(CmdFlightGet, []string{})
	// Output:
}

// ExampleFlightGetMissingStartID tests giving a non existing startID value, with null output
func ExampleFlightGetMissingStartID() {
	plugin.Register("mockflight", &mock.Mock{})
	config.Set(config.Config{Global: config.Global{Flighter: "mockflight"}})
	_ = flag.Set("startID", "9")
	runFlightGet(CmdFlightGet, []string{})
	// Output:
}

// ExampleFlightGetBadMax tests giving a non integer as max value, with null output
func ExampleFlightGetBadMax() {
	_ = flag.Set("startID", "9")
	_ = flag.Set("max", "a")
	runFlightGet(CmdFlightGet, []string{})
	// Output:
}

func TestFlightGetFailed(t *testing.T) {
	plugin.Register("mockflightgetfailed", &mock.Mock{
		GetFlightByIDF: func(id int) (flight.Flight, error) {
			return flight.Flight{}, errors.New("mock testing get flight failed")
		},
	},
	)
	config.Set(config.Config{Global: config.Global{Flighter: "mockflightgetfailed"}})
	flag.Set("id", "")
	flag.Set("startID", "")
	flag.Set("max", "")
	runFlightGet(CmdFlightGet, []string{})
}

func TestFlightGetFromIDFailed(t *testing.T) {
	plugin.Register("mockflightgetfromidfailed", &mock.Mock{
		GetFlightFromIDF: func(startID int, max int) ([]flight.Flight, error) {
			return []flight.Flight{}, errors.New("mock testing get flight from id failed")
		},
	},
	)
	config.Set(config.Config{Global: config.Global{Flighter: "mockflightgetfromidfailed"}})
	flag.Set("startID", "1")
	runFlightGet(CmdFlightGet, []string{})
}
