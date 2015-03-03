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

// Package spatial provides functions to calculate spatial values or convert
// between formats.
package spatial

import (
	"math"
	"strconv"

	"github.com/rochaporto/ezgliding/common"
)

const (
	// EarthRadius is the defined earth radius (in meters)
	EarthRadius = 6371000
	// Deg2Rad defines a constant to easily convert from degrees to radians
	Deg2Rad = math.Pi / 180
	// Rad2Deg defines a constant to easily convert from radians to degrees
	Rad2Deg = 180 / math.Pi
)

// DMS2Decimal converts the given coordinates from DMS to decimal format.
func DMS2Decimal(dms string) float64 {
	var degrees, minutes, seconds float64
	if len(dms) == 7 {
		degrees, _ = strconv.ParseFloat(dms[1:3], 64)
		minutes, _ = strconv.ParseFloat(dms[3:5], 64)
		seconds, _ = strconv.ParseFloat(dms[5:], 64)
	} else {
		degrees, _ = strconv.ParseFloat(dms[1:4], 64)
		minutes, _ = strconv.ParseFloat(dms[4:6], 64)
		seconds, _ = strconv.ParseFloat(dms[6:], 64)
	}
	var r float64
	r = degrees + (minutes / 60.0) + (seconds / 3600.0)
	if dms[0] == 'S' || dms[0] == 'W' {
		r = r * -1
	}
	return r
}

// GCDistance returns the great circle distance between the two points (in meters).
// It uses this formula:
//   d=2*asin(sqrt((sin((lat1-lat2)/2))^2 + cos(lat1)*cos(lat2)*(sin((lon1-lon2)/2))^2))
//
// Check EarthRadius in this pkg for the assumed earth radius.
func GCDistance(p1 common.Point, p2 common.Point) float64 {
	lat1 := p1.Latitude * Deg2Rad
	lon1 := p1.Longitude * Deg2Rad
	lat2 := p2.Latitude * Deg2Rad
	lon2 := p2.Longitude * Deg2Rad
	return (2 * math.Asin(
		math.Sqrt(math.Pow(math.Sin((lat1-lat2)/2), 2)+
			math.Cos(lat1)*math.Cos(lat2)*
				math.Pow(math.Sin((lon1-lon2)/2), 2)))) * EarthRadius
}

// Bearing returns the true course from point1 to point2 (in degrees).
// It uses the formula:
//   mod(atan2(sin(lon1-lon2)*cos(lat2), cos(lat1)*sin(lat2)-sin(lat1)*cos(lat2)*cos(lon1-lon2)), 2*pi)
func Bearing(p1 common.Point, p2 common.Point) float64 {
	lat1 := p1.Latitude * Deg2Rad
	lon1 := p1.Longitude * Deg2Rad
	lat2 := p2.Latitude * Deg2Rad
	lon2 := p2.Longitude * Deg2Rad
	return math.Mod(
		math.Atan2(
			math.Sin(lon1-lon2)*math.Cos(lat2),
			math.Cos(lat1)*math.Sin(lat2)-math.Sin(lat1)*math.Cos(lat2)*math.Cos(lon1-lon2)),
		2*math.Pi) * Rad2Deg
}
