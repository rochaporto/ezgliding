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

package igc

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rochaporto/ezgliding/util"
)

// Parse returns a Flight object corresponding to the given content.
// content should be a text string in the IGC format.
func Parse(content string) (Flight, error) {
	flight := NewFlight()
	var err error
	p := Parser{}
	lines := strings.Split(content, "\n")
	for i := range lines {
		line := lines[i]
		// ignore empty lines
		if len(strings.Trim(line, " ")) < 1 {
			continue
		}
		switch line[0] {
		case 'A':
			err = p.parseA(line, &flight)
		case 'B':
			err = p.parseB(line, &flight)
		case 'C':
			if !p.taskDone {
				err = p.parseC(lines[i:], &flight)
			}
		case 'D':
			err = p.parseD(line, &flight)
		case 'E':
			err = p.parseE(line, &flight)
		case 'F':
			err = p.parseF(line, &flight)
		case 'G':
			err = p.parseG(line, &flight)
		case 'H':
			err = p.parseH(line, &flight)
		case 'I':
			err = p.parseI(line)
		case 'J':
			err = p.parseJ(line)
		case 'K':
			err = p.parseK(line, &flight)
		case 'L':
			err = p.parseL(line, &flight)
		}
		if err != nil {
			return flight, err
		}
	}

	return flight, nil
}

type field struct {
	start int64
	end   int64
	tlc   string
}

// Parser gives functionality to parse IGC flight files.
type Parser struct {
	IFields  []field
	JFields  []field
	taskDone bool
	numSat   int
}

func (p *Parser) parseA(line string, flight *Flight) error {
	if len(line) < 7 {
		return fmt.Errorf("line too short :: %v", line)
	}
	flight.Header.Manufacturer = line[1:4]
	flight.Header.UniqueID = line[4:7]
	flight.Header.AdditionalData = line[7:]
	return nil
}

func (p *Parser) parseB(line string, flight *Flight) error {
	if len(line) < 37 {
		return fmt.Errorf("line too short :: %v", line)
	}
	pt := NewPoint()
	var err error
	pt.Time, err = time.Parse(TimeFormat, line[1:7])
	if err != nil {
		return err
	}
	pt.Latitude = util.DMS2Decimal(line[7:15])
	pt.Longitude = util.DMS2Decimal(line[15:24])
	if line[24] == 'A' || line[24] == 'V' {
		pt.FixValidity = line[24]
	} else {
		return fmt.Errorf("invalid fix validity :: %v", line[24])
	}
	pt.PressureAltitude, err = strconv.ParseInt(line[25:30], 10, 64)
	if err != nil {
		return err
	}
	pt.GNSSAltitude, err = strconv.ParseInt(line[30:35], 10, 64)
	if err != nil {
		return err
	}
	for _, f := range p.IFields {
		pt.IData[f.tlc] = line[f.start-1 : f.end]
	}
	pt.NumSatellites = p.numSat
	flight.Points = append(flight.Points, pt)
	return nil
}

func (p *Parser) parseC(lines []string, flight *Flight) error {
	line := lines[0]
	if len(line) < 25 {
		return fmt.Errorf("wrong line size :: %v", line)
	}
	var err error
	var nTP int
	if nTP, err = strconv.Atoi(line[23:25]); err != nil {
		return fmt.Errorf("invalid number of turnpoints :: %v", line)
	}
	if len(lines) < 5+nTP {
		return fmt.Errorf("invalid number of C record lines :: %v", lines)
	}
	if flight.Task.DeclarationDate, err = time.Parse(DateFormat+TimeFormat, lines[0][1:13]); err != nil {
		return err
	}
	if flight.Task.FlightDate, err = time.Parse(DateFormat, lines[0][13:19]); err != nil {
		return err
	}
	if flight.Task.Number, err = strconv.Atoi(line[19:23]); err != nil {
		return err
	}
	flight.Task.Description = line[25:]
	if flight.Task.Takeoff, err = p.taskPoint(lines[1]); err != nil {
		return err
	}
	if flight.Task.Start, err = p.taskPoint(lines[2]); err != nil {
		return err
	}
	for i := 0; i < nTP; i++ {
		var tp Point
		if tp, err = p.taskPoint(lines[3+i]); err != nil {
			return err
		}
		flight.Task.Turnpoints = append(flight.Task.Turnpoints, tp)
	}
	if flight.Task.Finish, err = p.taskPoint(lines[3+nTP]); err != nil {
		return err
	}
	if flight.Task.Landing, err = p.taskPoint(lines[4+nTP]); err != nil {
		return err
	}
	p.taskDone = true
	return nil
}

func (p *Parser) taskPoint(line string) (Point, error) {
	if len(line) < 18 {
		return Point{}, fmt.Errorf("line too short :: %v", line)
	}
	return Point{
		Latitude:    util.DMS2Decimal(line[1:9]),
		Longitude:   util.DMS2Decimal(line[9:18]),
		Description: line[18:],
	}, nil
}

func (p *Parser) parseD(line string, flight *Flight) error {
	if len(line) < 6 {
		return fmt.Errorf("line too short :: %v", line)
	}
	if line[1] == '2' {
		flight.DGPSStationID = line[2:6]
	}
	return nil
}

func (p *Parser) parseE(line string, flight *Flight) error {
	if len(line) < 10 {
		return fmt.Errorf("line too short :: %v", line)
	}
	t, err := time.Parse(TimeFormat, line[1:7])
	if err != nil {
		return err
	}
	if flight.Events[t] == nil {
		flight.Events[t] = make(map[string]string)
	}
	flight.Events[t][line[7:10]] = line[10:]
	return nil
}

func (p *Parser) parseF(line string, flight *Flight) error {
	if len(line) < 7 {
		return fmt.Errorf("line too short :: %v", line)
	}
	t, err := time.Parse(TimeFormat, line[1:7])
	if err != nil {
		return err
	}
	if flight.Satellites[t] == nil {
		flight.Satellites[t] = []int{}
	}
	for i := 7; i < len(line)-1; i = i + 2 {
		var n int
		if n, err = strconv.Atoi(line[i : i+2]); err != nil {
			return err
		}
		flight.Satellites[t] = append(flight.Satellites[t], n)
	}
	p.numSat = len(flight.Satellites[t])
	return nil
}

func (p *Parser) parseG(line string, flight *Flight) error {
	flight.Signature = flight.Signature + line[1:]
	return nil
}

func (p *Parser) parseH(line string, flight *Flight) error {
	var err error
	if len(line) < 5 {
		return fmt.Errorf("line too short :: %v", line)
	}

	switch line[2:5] {
	case "DTE":
		if len(line) < 11 {
			return fmt.Errorf("line too short :: %v", line)
		}
		flight.Header.Date, err = time.Parse(DateFormat, line[5:11])
	case "FXA":
		if len(line) < 8 {
			return fmt.Errorf("line too short :: %v", line)
		}
		flight.Header.FixAccuracy, err = strconv.ParseInt(line[5:8], 10, 64)
	case "PLT":
		flight.Header.Pilot = line[5:]
	case "CM2":
		flight.Header.Crew = line[5:]
	case "GTY":
		flight.Header.GliderType = line[5:]
	case "GID":
		flight.Header.GliderID = line[5:]
	case "DTM":
		if len(line) < 8 {
			return fmt.Errorf("line too short :: %v", line)
		}
		flight.Header.GPSDatum = line[5:8]
	case "RFW":
		flight.Header.FirmwareVersion = line[5:]
	case "RHW":
		flight.Header.HardwareVersion = line[5:]
	case "FTY":
		flight.Header.FlightRecorder = line[5:]
	case "GPS":
		flight.Header.GPS = line[5:]
	case "PRS":
		flight.Header.PressureSensor = line[5:]
	case "CID":
		flight.Header.CompetitionID = line[5:]
	case "CCL":
		flight.Header.CompetitionClass = line[5:]
	default:
		err = fmt.Errorf("unknown error record :: %v", line)
	}

	return err
}

func (p *Parser) parseI(line string) error {
	if len(line) < 3 {
		return fmt.Errorf("line too short :: %v", line)
	}
	n, err := strconv.ParseInt(line[1:3], 10, 0)
	if err != nil {
		return fmt.Errorf("invalid number of I fields :: %v", line)
	}
	if len(line) != int(n*7+3) {
		return fmt.Errorf("wrong line size :: %v", line)
	}
	for i := 0; i < int(n); i++ {
		s := i*7 + 3
		start, _ := strconv.ParseInt(line[s:s+2], 10, 0)
		end, _ := strconv.ParseInt(line[s+2:s+4], 10, 0)
		tlc := line[s+4 : s+7]
		p.IFields = append(p.IFields, field{start: start, end: end, tlc: tlc})
	}
	return nil
}

func (p *Parser) parseJ(line string) error {
	if len(line) < 3 {
		return fmt.Errorf("line too short :: %v", line)
	}
	n, err := strconv.ParseInt(line[1:3], 10, 0)
	if err != nil {
		return fmt.Errorf("invalid number of J fields :: %v", line)
	}
	if len(line) != int(n*7+3) {
		return fmt.Errorf("wrong line size :: %v", line)
	}
	for i := 0; i < int(n); i++ {
		s := i*7 + 3
		start, _ := strconv.ParseInt(line[s:s+2], 10, 0)
		end, _ := strconv.ParseInt(line[s+2:s+4], 10, 0)
		tlc := line[s+4 : s+7]
		p.JFields = append(p.JFields, field{start: start, end: end, tlc: tlc})
	}
	return nil
}

func (p *Parser) parseK(line string, flight *Flight) error {
	if len(line) < 7 {
		return fmt.Errorf("line too short :: %v", line)
	}
	t, err := time.Parse(TimeFormat, line[1:7])
	if err != nil {
		return err
	}
	fields := make(map[string]string)
	for _, f := range p.JFields {
		fields[f.tlc] = line[f.start-1 : f.end]
	}
	flight.K[t] = fields
	return nil
}

func (p *Parser) parseL(line string, flight *Flight) error {
	if len(line) < 4 {
		return fmt.Errorf("line too short :: %v", line)
	}
	flight.Logbook = append(flight.Logbook, LogEntry{line[1:4], line[4:]})
	return nil
}
