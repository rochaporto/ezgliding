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
//
// Common values for ezgliding related data, and some utility functions.
//

// Set of bit fields mapping to the available ezgliding flags.
var UnclearAirstrip = 1 << 0
var Outlanding      = 1 << 1
var ULMSite         = 1 << 2
var GliderSite      = 1 << 3
var ElevationProved = 1 << 4
var Asphalt         = 1 << 5
var Concrete        = 1 << 6
var Loam            = 1 << 7
var Sand            = 1 << 8
var Clay            = 1 << 9
var Grass           = 1 << 10
var Gravel          = 1 << 11
var Dirt            = 1 << 12

// feature2LatLng returns a maps LatLng object corresponding to feature's coords.
function feature2LatLng(feature) {
	return new google.maps.LatLng(
		parseFloat(feature.geometry.coordinates[1]),
		parseFloat(feature.geometry.coordinates[0]));
}

