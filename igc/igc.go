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

// Package igc provides IGC format parsing and handling.
// Check ... for full information on the IGC format.
package igc

import "time"

const (
	// TimeFormat is the golang time.Parse format for IGC time
	TimeFormat = "150405"
	// DateFormat is the golang time.Parse format for IGC time
	DateFormat = "020106"
)

// Flight represents all the flight data (header and track).
type Flight struct {
	Header        Header
	Points        []Point // FIXME: use a map keyed by time.Time instead?
	K             map[time.Time]map[string]string
	Events        map[time.Time]map[string]string
	Satellites    map[time.Time][]int
	Logbook       []LogEntry
	Task          Task
	DGPSStationID string
	Signature     string
}

// NewFlight returns a new instance of Flight.
// It initializes all the structures with zero values.
func NewFlight() Flight {
	flight := Flight{}
	flight.K = make(map[time.Time]map[string]string)
	flight.Events = make(map[time.Time]map[string]string)
	flight.Satellites = make(map[time.Time][]int)
	return flight
}

// Header holds the meta information of a flight.
type Header struct {
	Manufacturer     string
	UniqueID         string
	AdditionalData   string
	Date             time.Time
	FixAccuracy      int64
	Pilot            string
	Crew             string
	GliderType       string
	GliderID         string
	GPSDatum         string
	FirmwareVersion  string
	HardwareVersion  string
	FlightRecorder   string
	GPS              string
	PressureSensor   string
	CompetitionID    string
	CompetitionClass string
}

// Point represents a gps read (single point in the flight track).
type Point struct {
	Time             time.Time
	Latitude         float64
	Longitude        float64
	FixValidity      byte
	PressureAltitude int64
	GNSSAltitude     int64
	IData            map[string]string
	NumSatellites    int
	Description      string
}

// NewPoint creates a new Point struct and returns it.
// It initializes all structures to zero values.
func NewPoint() Point {
	var pt Point
	pt.IData = make(map[string]string)
	return pt
}

// Task is a pre-declared flight task to be performed in this flight.
type Task struct {
	DeclarationDate time.Time
	FlightDate      time.Time
	Number          int
	Takeoff         Point
	Start           Point
	Turnpoints      []Point
	Finish          Point
	Landing         Point
	Description     string
}

// LogEntry holds a logbook/comment entry in the IGC file.
type LogEntry struct {
	Type string
	Text string
}
