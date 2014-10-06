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

package common

// Airfield keeps details about a specific airfield
type Airfield struct {
	ID        string
	ShortName string
	Name      string
	Text      string
	ICAO      string
	Flags     int
	Catalog   int
	Length    int
	Elevation int
	Runway    string
	Frequency float64
	Latitude  string
	Longitude string
}

// Enum for Airfield flags
const (
	UnclearAirstrip = 1 << iota
	Outlanding      = 1 << iota
	ULMSite         = 1 << iota
	GliderSite      = 1 << iota
	ElevationProved = 1 << iota
	Asphalt         = 1 << iota
	Concrete        = 1 << iota
	Loam            = 1 << iota
	Sand            = 1 << iota
	Clay            = 1 << iota
	Grass           = 1 << iota
	Gravel          = 1 << iota
	Dirt            = 1 << iota
)
