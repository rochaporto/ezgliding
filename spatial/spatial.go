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

// Package spatial provides functionality for handling spatial data.
//
// This includes conversion for lat/lon between different formats (dms,
// decimal, ...) and other functions coming in the future.
package spatial

import (
	"errors"
	"math"
	"strconv"

	"github.com/paulmach/go.geojson"
	"github.com/rochaporto/ezgliding/airfield"
	"github.com/rochaporto/ezgliding/waypoint"
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

// DMD2Decimal converts the given coordinates from DMD (deg,min,decimalmin) to decimal format.
func DMD2Decimal(dmd string) float64 {
	var degrees, minutes, dminutes float64
	if dmd[0] == 'S' || dmd[0] == 'N' {
		degrees, _ = strconv.ParseFloat(dmd[1:3], 64)
		minutes, _ = strconv.ParseFloat(dmd[3:5], 64)
		dminutes, _ = strconv.ParseFloat(dmd[5:], 64)
	} else if dmd[len(dmd)-1] == 'S' || dmd[len(dmd)-1] == 'N' {
		degrees, _ = strconv.ParseFloat(dmd[0:2], 64)
		minutes, _ = strconv.ParseFloat(dmd[2:4], 64)
		dminutes, _ = strconv.ParseFloat(dmd[4:7], 64)
	} else if dmd[0] == 'W' || dmd[0] == 'E' {
		degrees, _ = strconv.ParseFloat(dmd[1:4], 64)
		minutes, _ = strconv.ParseFloat(dmd[4:6], 64)
		dminutes, _ = strconv.ParseFloat(dmd[6:], 64)
	} else if dmd[len(dmd)-1] == 'W' || dmd[len(dmd)-1] == 'E' {
		degrees, _ = strconv.ParseFloat(dmd[0:3], 64)
		minutes, _ = strconv.ParseFloat(dmd[3:5], 64)
		dminutes, _ = strconv.ParseFloat(dmd[5:8], 64)
	}
	var r float64
	r = degrees + ((minutes + (dminutes / 1000.0)) / 60.0)
	if dmd[0] == 'S' || dmd[0] == 'W' || dmd[len(dmd)-1] == 'S' || dmd[len(dmd)-1] == 'W' {
		r = r * -1
	}
	return r
}

// Struct2GeoJSON returns a collection of GeoJSON objects from the given structs.
// The given array can have distinct types (Airfield, Waypoint, Airspace) and the
// resulting GeoJSON will contain all fields as properties, and an additional one
// specifying the Go type (ex: "Go": "Airfield"), used later to unmarshal.
// Airfield and Waypoint result in Point geometries, Airspace in LineString.
func Struct2GeoJSON(features []interface{}) (*geojson.FeatureCollection, error) {
	result := geojson.NewFeatureCollection()
	for _, e := range features {
		var f []*geojson.Feature
		switch e.(type) {
		default:
			return nil, errors.New("geojson convertion not supported")
		case airfield.Airfield:
			f = airfield2GeoJSON([]airfield.Airfield{e.(airfield.Airfield)})
		case waypoint.Waypoint:
			f = waypoint2GeoJSON([]waypoint.Waypoint{e.(waypoint.Waypoint)})
		}
		result.AddFeature(f[0])
	}
	return result, nil
}

// airfield2GeoJSON converts the given airfield to GeoJSON format.
func airfield2GeoJSON(airfields []airfield.Airfield) []*geojson.Feature {
	result := []*geojson.Feature{}
	for _, airfield := range airfields {
		g := geojson.NewPointFeature([]float64{airfield.Longitude, airfield.Latitude})
		g.SetProperty("ID", airfield.ID)
		g.SetProperty("ShortName", airfield.ShortName)
		g.SetProperty("Name", airfield.Name)
		g.SetProperty("Region", airfield.Region)
		g.SetProperty("ICAO", airfield.ICAO)
		g.SetProperty("Flags", airfield.Flags)
		g.SetProperty("Catalog", airfield.Catalog)
		g.SetProperty("Length", airfield.Length)
		g.SetProperty("Elevation", airfield.Elevation)
		g.SetProperty("Runway", airfield.Runway)
		g.SetProperty("Frequency", airfield.Frequency)
		g.SetProperty("Go", "Airfield")
		result = append(result, g)
	}
	return result
}

// waypoint2GeoJSON converts the given waypoint to GeoJSON format.
func waypoint2GeoJSON(waypoints []waypoint.Waypoint) []*geojson.Feature {
	result := []*geojson.Feature{}
	for _, waypoint := range waypoints {
		g := geojson.NewPointFeature([]float64{waypoint.Longitude, waypoint.Latitude})
		g.SetProperty("ID", waypoint.ID)
		g.SetProperty("Name", waypoint.Name)
		g.SetProperty("Description", waypoint.Description)
		g.SetProperty("Region", waypoint.Region)
		g.SetProperty("Flags", waypoint.Flags)
		g.SetProperty("Elevation", waypoint.Elevation)
		g.SetProperty("Go", "Waypoint")
		result = append(result, g)
	}
	return result
}

// GeoJSON2Struct returns airfield, waypoint, etc objects from the given GeoJSON.
// The resulting array contains distinct types (airfield.Airfield, waypoint.Waypoint,
// airspace.Airspace, ...) and the unmarshaling is done with the same rules as
// described in Struct2GeoJSON.
func GeoJSON2Struct(json string) ([]interface{}, error) {
	result := []interface{}{}
	collection, err := geojson.UnmarshalFeatureCollection([]byte(json))
	if err != nil {
		return nil, err
	}
	for _, f := range collection.Features {
		var o interface{}
		goType := f.PropertyMustString("Go")
		switch goType {
		case "Airfield":
			o = feature2Airfield(f)
		case "Waypoint":
			o = feature2Waypoint(f)
		default:
			return result, errors.New("geojson feature given not supported")
		}
		result = append(result, o)
	}
	return result, nil
}

func feature2Airfield(f *geojson.Feature) airfield.Airfield {
	a := airfield.Airfield{
		ID: f.PropertyMustString("ID"), ShortName: f.PropertyMustString("ShortName"),
		Name: f.PropertyMustString("Name"), Region: f.PropertyMustString("Region"),
		ICAO: f.PropertyMustString("ICAO"), Flags: f.PropertyMustInt("Flags"),
		Catalog: f.PropertyMustInt("Catalog"), Length: f.PropertyMustInt("Length"),
		Elevation: f.PropertyMustInt("Elevation"), Runway: f.PropertyMustString("Runway"),
		Frequency: f.PropertyMustFloat64("Frequency"),
	}
	a.Longitude = f.Geometry.Point[0]
	a.Latitude = f.Geometry.Point[1]
	return a
}
func feature2Waypoint(f *geojson.Feature) waypoint.Waypoint {
	w := waypoint.Waypoint{
		ID: f.PropertyMustString("ID"), Description: f.PropertyMustString("Description"),
		Name: f.PropertyMustString("Name"), Region: f.PropertyMustString("Region"),
		Flags: f.PropertyMustInt("Flags"), Elevation: f.PropertyMustInt("Elevation"),
	}
	w.Longitude = f.Geometry.Point[0]
	w.Latitude = f.Geometry.Point[1]
	return w
}

// GCDistance returns the great circle distance between the two points.
// Points latitude and longitude are in decimal degrees, result in meters.
// It uses this formula:
//   d=2*asin(sqrt((sin((lat1-lat2)/2))^2 + cos(lat1)*cos(lat2)*(sin((lon1-lon2)/2))^2))
//
// Check EarthRadius in this pkg for the assumed earth radius.
func GCDistance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	lat1 = lat1 * Deg2Rad
	lon1 = lon1 * Deg2Rad
	lat2 = lat2 * Deg2Rad
	lon2 = lon2 * Deg2Rad
	return GCDistanceRadians(lat1, lon1, lat2, lon2)
}

// GCDistanceRadians returns the great circle distance between the two points.
// Points latitude and longitude are in radians, result in meters.
func GCDistanceRadians(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	return (2 * math.Asin(
		math.Sqrt(math.Pow(math.Sin((lat1-lat2)/2), 2)+
			math.Cos(lat1)*math.Cos(lat2)*
				math.Pow(math.Sin((lon1-lon2)/2), 2)))) * EarthRadius
}

// Bearing returns the true course from point1 to point2.
// Points latitude and longitude are decimal degrees, result is in degrees.
// It uses the formula:
//   mod(atan2(sin(lon1-lon2)*cos(lat2), cos(lat1)*sin(lat2)-sin(lat1)*cos(lat2)*cos(lon1-lon2)), 2*pi)
func Bearing(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	lat1 = lat1 * Deg2Rad
	lon1 = lon1 * Deg2Rad
	lat2 = lat2 * Deg2Rad
	lon2 = lon2 * Deg2Rad
	return BearingRadians(lat1, lon1, lat2, lon2)
}

// BearingRadians returns the true course from point1 to point2.
// Points latitude and longitude are in radians, result is in degrees.
// It uses the formula:
//   mod(atan2(sin(lon1-lon2)*cos(lat2), cos(lat1)*sin(lat2)-sin(lat1)*cos(lat2)*cos(lon1-lon2)), 2*pi)
func BearingRadians(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	return math.Mod(
		math.Atan2(
			math.Sin(lon1-lon2)*math.Cos(lat2),
			math.Cos(lat1)*math.Sin(lat2)-math.Sin(lat1)*math.Cos(lat2)*math.Cos(lon1-lon2)),
		2*math.Pi) * Rad2Deg
}
