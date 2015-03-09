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

package netcoupe

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/rochaporto/ezgliding/flight"
)

var reNom = regexp.MustCompile("(?s)DisplayContactDetail\\(.\\d*.,\\s.hasPrevious...>([\\w\\s]*)</a>")
var reCategorie = regexp.MustCompile("(?s)Cat.gorie&nbsp;:</b>\\s*</div>\\s*</td>\\s*<td>\\s*<div align=\"left\">([\\w\\s]*)</div>")
var reClub = regexp.MustCompile("(?s)DisplayClubDetail\\(.\\d*.,\\s.hasPrevious...>([\\w\\s]*)</a>")
var reDate = regexp.MustCompile("(?s)Date&nbsp;:</b>\\s*</div>\\s*</td>\\s*<td>\\s*<div align=\"left\">\\s*(\\d+/\\d+/\\d+)\\s*</div>")
var reDepart = regexp.MustCompile("(?s)A.rodrome de d.part&nbsp;:</b>\\s*</div>\\s*</td>\\s*<td>\\s*<div align=\"left\">([\\w\\s]*)</div>")
var reRegion = regexp.MustCompile("(?s)R.gion&nbsp;:</b>\\s*</div>\\s*</td>\\s*<td>\\s*<div align=\"left\">([\\w\\s]*)</div>")
var rePays = regexp.MustCompile("(?s)Pays&nbsp;:</b>\\s*</div>\\s*</td>\\s*<td>\\s*<div align=\"left\">([\\w\\s]*)</div>")
var reDistance = regexp.MustCompile("(?s)Distance&nbsp;:</b>\\s*</div>\\s*</td>\\s*<td>\\s*<div align=\"left\">(.*)\\s*&nbsp;kms\\s*</div>")
var rePoints = regexp.MustCompile("(?s)Points&nbsp;:</b>\\s*</div>\\s*</td>\\s*<td>\\s*<div align=\"left\">\\s*(\\S*)\\s*</div>")
var rePlaneur = regexp.MustCompile("(?s)Planeur&nbsp;:</b>\\s*</div>\\s*</td>\\s*<td>\\s*<div align=\"left\">\\s*<table border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\\s*<tbody>\\s*<tr>\\s*<td valign=\"middle\">([\\w\\s]*)&nbsp;&nbsp;\\s*</td>")
var reTypeCircuit = regexp.MustCompile("(?s)Type de circuit&nbsp;:</b>\\s*</div>\\s*</td>\\s*<td>\\s*<div align=\"left\">([\\w\\s]*)</div>")
var reVitesse = regexp.MustCompile("(?s)Vitesse moyenne du circuit&nbsp;:</b>\\s*</div>\\s*</td>\\s*<td>\\s*<div align=\"left\">\\s*(.*)&nbsp;km/h\\s*</div>")
var rePointDepart = regexp.MustCompile("(?s)Point de d.part&nbsp;:</b>\\s*</div>\\s*</td>\\s*<td>\\s*<div align=\"left\">([\\w\\s]*)</div>")
var rePointVirage1 = regexp.MustCompile("(?s)Point de virage n.1&nbsp;:</b>\\s*</div>\\s*</td>\\s*<td>\\s*<div align=\"left\">([\\w\\s]*)</div>")
var rePointVirage2 = regexp.MustCompile("(?s)Point de virage n.2&nbsp;:</b>\\s*</div>\\s*</td>\\s*<td>\\s*<div align=\"left\">([\\w\\s]*)</div>")
var rePointVirage3 = regexp.MustCompile("(?s)Point de virage n.3&nbsp;:</b>\\s*</div>\\s*</td>\\s*<td>\\s*<div align=\"left\">([\\w\\s]*)</div>")
var rePointArrivee = regexp.MustCompile("(?s)Point d.arriv.e&nbsp;:</b>\\s*</div>\\s*</td>\\s*<td>\\s*<div align=\"left\">([\\w\\s]*)</div>")
var reCommentaires = regexp.MustCompile("(?s)Commentaires&nbsp;:</b>\\s*</div>\\s*</td>\\s*<td>\\s*<div align=\"left\">([\\w\\s]*)</div>")
var reFichierIGC = regexp.MustCompile("(?s)Fichier .IGC&nbsp;:</b>\\s*</div>\\s*</td>\\s*<td>\\s*<div align=\"left\">\\s*<a href=\"([\\S]*)\">")

func (nc *Netcoupe) parseDetails(html string) (flight.Source, error) {
	var err error
	sourceData := flight.Source{}
	sourceData.Name = nc.getRegexpField(reNom, html)
	sourceData.Category = nc.getRegexpField(reCategorie, html)
	sourceData.Club = nc.getRegexpField(reClub, html)
	dateStr := nc.getRegexpField(reDate, html)
	if sourceData.Date, err = time.Parse("02/01/2006", dateStr); err != nil {
		return flight.Source{}, err
	}
	sourceData.Takeoff = nc.getRegexpField(reDepart, html)
	sourceData.Region = nc.getRegexpField(reRegion, html)
	sourceData.Country = nc.getRegexpField(rePays, html)
	distance := nc.getRegexpField(reDistance, html)
	if sourceData.Distance, err = strconv.ParseFloat(strings.Replace(distance, ",", ".", 1), 64); err != nil {
		return flight.Source{}, err
	}
	points := nc.getRegexpField(rePoints, html)
	if sourceData.Points, err = strconv.ParseFloat(strings.Replace(points, ",", ".", 1), 64); err != nil {
		return flight.Source{}, err
	}
	sourceData.Type = nc.getRegexpField(rePlaneur, html)
	sourceData.CircuitType = nc.getRegexpField(reTypeCircuit, html)
	speed := nc.getRegexpField(reVitesse, html)
	if sourceData.Speed, err = strconv.ParseFloat(strings.Replace(speed, ",", ".", 1), 64); err != nil {
		return flight.Source{}, err
	}
	sourceData.Start = nc.getRegexpField(rePointDepart, html)
	sourceData.Turnpoints = make([]flight.Point, 3)
	sourceData.Turnpoints[0] = flight.Point{Description: nc.getRegexpField(rePointVirage1, html)}
	sourceData.Turnpoints[1] = flight.Point{Description: nc.getRegexpField(rePointVirage2, html)}
	sourceData.Turnpoints[2] = flight.Point{Description: nc.getRegexpField(rePointVirage3, html)}
	sourceData.Finish = nc.getRegexpField(rePointArrivee, html)
	sourceData.Comment = nc.getRegexpField(reCommentaires, html)
	sourceData.DownloadURL = nc.getRegexpField(reFichierIGC, html)
	return sourceData, nil
}

// FIXME: should be a common function in another package
func (nc *Netcoupe) fetch(location string) (string, error) {
	var content []byte
	// case http
	resp, err := http.Get(location)
	if err == nil {
		defer resp.Body.Close()
		content, err = ioutil.ReadAll(resp.Body)
	} else { // case file
		resp, err := ioutil.ReadFile(location)
		if err != nil {
			return "", err
		}
		content = resp
	}
	return string(content), nil
}

func (nc *Netcoupe) detailURL(id int) string {
	return nc.BaseURL + nc.FlightDetailURL + strconv.Itoa(id)
}

func (nc *Netcoupe) getRegexpField(re *regexp.Regexp, content string) string {
	result := re.FindStringSubmatch(content)
	if len(result) < 2 {
		return "UNKNOWN"
	}
	return strings.TrimSpace(result[1])
}
