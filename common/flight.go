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

package common

import (
	"time"

	"github.com/rochaporto/ezgliding/config"
)

// Flighter is implemented by any data source which can provide or
// receive flight information.
type Flighter interface {
	// GetFlight returns all flights in the given regions, which have been
	// added or updated since the given time.
	GetFlight(regions []string, updatedSince time.Time) ([]Flight, error)
	// GetFlightFromID returns all flights starting from the given ID (inclusive), up to a max number of flights.
	// max can be a negative number if unlimited flights should be retrieved.
	GetFlightFromID(startID int, max int) ([]Flight, error)
	// GetFlightByID returns the flight corresponding to the given ID.
	GetFlightByID(id int) (Flight, error)
	// PutFlight adds the given flights.
	PutFlight(flights []Flight) error
}

// Flight represents all the flight data (header and track).
// Apart from Sources, all the information matches the IGC specification.
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
	// Sources is a map keyed on the plugin ID containing flight info
	// directly taken from the online site (no flight log parsing). There can
	// be multiple sources for the same flight.
	Sources map[string]Source
}

// NewFlight returns a new instance of Flight.
// It initializes all the structures with zero values.
func NewFlight() Flight {
	flight := Flight{}
	flight.K = make(map[time.Time]map[string]string)
	flight.Events = make(map[time.Time]map[string]string)
	flight.Satellites = make(map[time.Time][]int)
	flight.Sources = make(map[string]Source)
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

// Manufacturer holds the char identifier, the short id and the full name of
// an IGC Manufacturer, as defined in Appendix A (Codes for Manufacturers)
// of the IGC spec.
type Manufacturer struct {
	char  byte
	short string
	name  string
}

// Manufacturers holds the list of available manufacturers, as defined in
// Appendix A (Codes for Manufacturers) of the IGC spec.
var Manufacturers = map[string]Manufacturer{
	"GCS": Manufacturer{'A', "GCS", "Garrecht"},
	"LGS": Manufacturer{'B', "LGS", "Logstream"},
	"CAM": Manufacturer{'C', "CAM", "Cambridge Aero Instruments"},
	"DSX": Manufacturer{'D', "DSX", "Data Swan/DSX"},
	"EWA": Manufacturer{'E', "EWA", "EW Avionics"},
	"FIL": Manufacturer{'F', "FIL", "Filser"},
	"FLA": Manufacturer{'G', "FLA", "Flarm (Flight Alarm)"},
	"SCH": Manufacturer{'H', "SCH", "Scheffel"},
	"ACT": Manufacturer{'I', "ACT", "Aircotec"},
	"CNI": Manufacturer{'K', "CNI", "ClearNav Instruments"},
	"NKL": Manufacturer{'K', "NKL", "NKL"},
	"LXN": Manufacturer{'L', "LXN", "LX Navigation"},
	"IMI": Manufacturer{'M', "IMI", "IMI Gliding Equipment"},
	"NTE": Manufacturer{'N', "NTE", "New Technologies s.r.l."},
	"NAV": Manufacturer{'O', "NAV", "Naviter"},
	"PES": Manufacturer{'P', "PES", "Peschges"},
	"PRT": Manufacturer{'R', "PRT", "Print Technik"},
	"SDI": Manufacturer{'S', "SDI", "Streamline Data Instruments"},
	"TRI": Manufacturer{'T', "TRI", "Triadis Engineering GmbH"},
	"LXV": Manufacturer{'V', "LXV", "LXNAV d.o.o."},
	"WES": Manufacturer{'W', "WES", "Westerboer"},
	"XYY": Manufacturer{'X', "XYY", "Other manufacturer"},
	"ZAN": Manufacturer{'Z', "ZAN", "Zander"},
}

// Source holds information taken directly from the online source
// holding the flight.
type Source struct {
	SourceID    string
	Name        string
	Category    string
	Club        string
	Region      string
	Country     string
	Date        time.Time
	Takeoff     string
	Distance    float64
	Points      float64
	Type        string
	CircuitType string
	Speed       float64
	Start       string
	Turnpoints  []Point
	Finish      string
	Comment     string
	DownloadURL string
}

// OptimizationType identifies a type of Optimization (out and return, 3 TPs, 4 TPs, ...)
type OptimizationType int

const (
	// OutReturn is equivalent to 1 turnpoint
	OutReturn OptimizationType = iota
	// TP3 means 3 turnopoint optimization
	TP3
	// TP4 means 4 turnpoint optimization
	TP4
	// Triangle means 2 turnpoint
	Triangle
	// FAITriangle means a triangle with special conditions
	FAITriangle
)

// Contest defines the interface to calculate points for a given contest.
type Contest interface {
	// Points returns the number of points for the given set of TPs for a contest.
	Points(tps []Point, flags int) (Result, error)
}

// Contests holds all available online contests, keyed by their IDs.
var Contests = map[string]Contest{
	"netcoupe": Contest(NewNetcoupe(config.GetConfig())),
}

// Optimizer is the interface for flight optimization.
type Optimizer interface {
	// Optimize returns the number of points and optimization type for TPs of a contest.
	Optimize(flight Flight, contest Contest) (Result, error)
}

// Optimizers holds a map of all available Optimizers, keyed on ID.
var Optimizers = map[string]Optimizer{
	"montecarlo": Optimizer(NewMontecarlo(config.GetConfig())),
}

// Result is a track/distance/points result for a given Optimizer run.
type Result struct {
	Type     OptimizationType
	TPs      []Point
	Distance float64
	Points   float64
}
