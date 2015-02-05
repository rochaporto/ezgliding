// Copyright 2015 The ezgliding Authors.
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

// Package netcoupe provides code to fetch and parse Netcoupe flights.
//
// Netcoupe (www.netcoupe.net) is an online competition between glider
// pilots, mostly used by pilots in France.
//
package netcoupe

import (
	"errors"
	"time"

	"github.com/rochaporto/ezgliding/flight"
)

// ID for this plugin implementation.
const (
	ID string = "netcoupe"
)

const (
	// baseURL is the base url for the netcoupe website
	baseURL string = "http://netcoupe.net"
	// flightDetailURL is the subpath to fetch details of flight with 'ID'
	flightDetailURL string = "/Results/FlightDetail.aspx?FlightID="
	// maxIDGap is the max num of subsequent missing IDs when crawling flights
	maxIDGap int = 3
)

// Config holds the netcoupe configuration.
type Config struct {
	BaseURL         string
	FlightDetailURL string
	MaxIDGap        int
}

// Netcoupe gives functionality to fetch and parse information regarding
// flights from the netcoupe online gliding competition.
type Netcoupe struct {
	Config
}

// New returns a new instance of Netcoupe, with the given Config.
func New(config Config) (*Netcoupe, error) {
	if config.MaxIDGap == 0 {
		config.MaxIDGap = maxIDGap
	}
	if config.FlightDetailURL == "" {
		config.FlightDetailURL = flightDetailURL
	}
	if config.BaseURL == "" {
		config.BaseURL = baseURL
	}
	return &Netcoupe{config}, nil
}

// GetFlight implements flight.GetFlight().
func (nc *Netcoupe) GetFlight(regions []string, updatedSince time.Time) ([]flight.Flight, error) {
	return nil, errors.New("Not implemented")
}

// GetFlightByID implements flight.GetFlightByID().
func (nc *Netcoupe) GetFlightByID(id int) (flight.Flight, error) {
	var err error
	html, err := nc.fetch(nc.detailURL(id))
	if err != nil {
		return flight.Flight{}, err
	}
	source, err := nc.parseDetails(html)
	if err != nil {
		return flight.Flight{}, err
	}
	content, err := nc.fetch(nc.BaseURL + source.DownloadURL)
	if err != nil {
		return flight.Flight{}, err
	}
	f, err := flight.ParseIGC(content)
	if err != nil {
		return flight.Flight{}, err
	}
	f.Sources[ID] = source
	return f, nil
}

// GetFlightFromID implements flight.GetFlightFromID().
func (nc *Netcoupe) GetFlightFromID(startID int, max int) ([]flight.Flight, error) {
	var missed int
	result := []flight.Flight{}
	for cid := startID; missed < nc.MaxIDGap && (max < 0 || len(result) < max); cid++ {
		flight, err := nc.GetFlightByID(cid)
		if err != nil {
			missed++
			continue
		}
		result = append(result, flight)
	}
	return result, nil
}

// PutFlight implements flight.PutFlight().
func (nc *Netcoupe) PutFlight(flights []flight.Flight) error {
	return errors.New("Not implemented")
}
