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

import (
	"image/color"
	"time"
)

// Airspacer is implemented by any data source which can manage
// airspace information. Only one of Get() or Put() or both can be implemented
// by the source.
// FIXME: add boundbox filter to Get
type Airspacer interface {
	GetAirspace(regions []string, updatedSince time.Time) ([]Airspace, error)
	PutAirspace(airspaces []Airspace) error
}

// Airspace keeps details about a specific airspace area
//
// Date is of airspace definition or update.
//
// Label is a list of Lat/Lon coordinates where the airspace label
// (usually the name) should be placed.
type Airspace struct {
	ID       string
	Date     time.Time
	Class    byte
	Name     string
	Ceiling  string
	Floor    string
	Label    []string
	Segments []AirspaceSegment
	Pen      Pen
}

// AirspaceSegment is one of polygon, arc, circle.
//
// Clockwise indicates direction for building arcs.
//
// X is the center for arcs and circles, W is the width for airways (unused).
//
// Data interpretation depends on record type:
//   Polygon: coordinate point (to be added)
//   Arc: radius, start, end || coordinate1, coordinate2 (center in X)
//   Circle: radius (from X)
//
type AirspaceSegment struct {
	Type        AirspaceSegmentType
	Clockwise   bool
	X           string
	W           int
	Radius      float64
	AngleStart  float64
	AngleEnd    float64
	Coordinate1 string
	Coordinate2 string
}

// AirspaceSegmentType is an int for an AirspaceSegment.
type AirspaceSegmentType int

// Constants for airspace record types
const (
	Polygon AirspaceSegmentType = iota
	Arc
	Circle
)

// Pen has drawing info for an Airspace.
type Pen struct {
	Style       PenStyle
	Width       int
	Color       color.Color
	InsideColor color.Color
}

// PenStyle is one of Solid, Dash, None.
type PenStyle int

const (
	// Solid PenStyle.
	Solid PenStyle = iota
	// Dash (ed) PenStyle.
	Dash
	// None is no PenStyle
	None
)
