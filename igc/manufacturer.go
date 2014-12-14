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
