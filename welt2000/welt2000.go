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

// welt2000 provides functionality to fetch and parse airfield and
// waypoint information, taking the welt release as input.
//
// Check the welt2000 website for more information on the data:
// http://www.segelflug.de/vereine/welt2000/
//
package welt2000

import (
	"errors"
	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/rss"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Release contains info about a specific release
type Release struct {
	Date      time.Time
	Source    string
	Airfields []common.Airfield
	Waypoints []common.Waypoint
}

// List checks the welt2000 rss feed and lists the releases found
func List(location string) ([]Release, error) {
	var content []byte
	// case http
	resp, err := http.Get(location)
	if err == nil {
		defer resp.Body.Close()
		content, err = ioutil.ReadAll(resp.Body)
	} else { // case file
		resp, err := ioutil.ReadFile(location)
		if err != nil {
			return nil, err
		}
		content = resp
	}
	feed, err := rss.Parse(content)
	if err != nil {
		return nil, err
	}

	res := make([]Release, 10)
	for i, item := range feed.Items {
		res[i].Date = item.Date
		res[i].Source = item.Link
	}
	return res, nil
}

// Fetch gets and returns the Release at the given location
func Fetch(location string) (*Release, error) {
	r := Release{Source: location}
	err := r.Fetch()
	return &r, err
}

// Fetch fills up the Release object with data after parsing the content at Release.Source
func (r *Release) Fetch() error {
	var content []byte

	resp, err := http.Get(r.Source)
	// case http
	if err == nil {
		defer resp.Body.Close()
		content, err = ioutil.ReadAll(resp.Body)
	} else { // case file
		content, err = ioutil.ReadFile(r.Source)
	}
	if err != nil {
		return err
	}
	return r.Parse(content)
}

// Parse fills in the Release object by parsing r.Data
func (r *Release) Parse(content []byte) error {
	if content == nil {
		return errors.New("No data available to parse")
	}

	lines := strings.Split(string(content), "\n")
	for i := range lines {
		switch {
		case len(lines[i]) < 1: // empty line
			continue
		case lines[i][0] == '$': // comment
			continue
		case lines[i][5] == '1' || lines[i][5] == '2': // airfield
			r.parseAirfield(lines[i])
		default: // waypoint
			r.parseWaypoint(lines[i])
		}
	}
	return nil
}

func (r *Release) parseAirfield(line string) error {
	airfield := common.Airfield{}
	if line[4] == '2' { // unclear airstrip
		airfield.Flags |= common.UnclearAirstrip
		airfield.ShortName = strings.Trim(line[0:4], " ")
	} else { // regular airstrip
		airfield.ShortName = strings.Trim(line[0:5], " ")
	}
	airfield.Name = strings.Trim(line[7:20], " ")
	if line[23] == '#' && line[24] != ' ' { // ICAO available
		airfield.ICAO = line[24:28]
		airfield.ID = airfield.ICAO
	} else {
		airfield.ID = airfield.ShortName
	}
	if line[23:27] == "*ULM" { // ultralight site
		airfield.Flags |= common.ULMSite
	} else if line[5] == '2' { // outlanding
		airfield.Flags |= common.Outlanding
		airfield.Catalog, _ = strconv.Atoi(line[26:28])
	} else if line[20:24] == "GLD#" || line[23:28] == "#GLD" { // glider site
		airfield.Flags |= common.GliderSite
	}
	airfield.Flags |= r.runwayType2Bit(line[28])
	airfield.Length, _ = strconv.Atoi(line[29:32])
	airfield.Length *= 10
	airfield.Runway = line[32:36]
	decimal, _ := strconv.ParseFloat(line[39:41], 64)
	airfield.Frequency, _ = strconv.ParseFloat(line[36:39], 64)
	airfield.Frequency += decimal * 0.01
	elevation := strings.Trim(line[41:45], " ")
	airfield.Elevation, _ = strconv.Atoi(elevation)
	airfield.Latitude = line[45:52]
	airfield.Longitude = line[52:60]
	r.Airfields = append(r.Airfields, airfield)
	return nil
}

func (r *Release) runwayType2Bit(t uint8) int {
	switch t {
	case 'A':
		return common.Asphalt
	case 'C':
		return common.Concrete
	case 'L':
		return common.Loam
	case 'S':
		return common.Sand
	case 'Y':
		return common.Clay
	case 'G':
		return common.Grass
	case 'V':
		return common.Gravel
	case 'D':
		return common.Dirt
	}
	return 0
}

func (r *Release) parseWaypoint(line string) error {
	waypoint := common.Waypoint{
		Name: strings.Trim(line[0:6], " "), ID: strings.Trim(line[0:6], " "),
		Description: strings.Trim(line[7:41], " "),
		Latitude:    line[45:52], Longitude: line[52:60],
	}
	elevation := strings.Trim(line[41:45], " ")
	waypoint.Elevation, _ = strconv.Atoi(elevation)
	r.Waypoints = append(r.Waypoints, waypoint)
	return nil
}
