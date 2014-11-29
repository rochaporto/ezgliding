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

// Package soaringweb provides functionality to fetch and parse airspace
// information, taking the international soaringweb db as input.
//
// Check the soaringweb website for more information on the data:
// http://soaringweb.org/Airspace/HomePage.html
//
// Airspace data is handled in OpenAir format.
//
package soaringweb

import (
	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/openair"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// ID for this plugin implementation.
const (
	ID string = "soaringweb"
)

// Regions supported for queries.
// FIXME: Missing finland and portugal (non standard page references)
var Regions = []string{"AF", "AU", "AT", "BE", "HR", "CZ", "DK", "FR",
	"DE", "HU", "EI", "IT", "LV", "LT", "MK", "NL", "NO", "PL", "SE",
	"SI", "SK", "ES", "CH", "UK", "CO", "BR"}

// Release contains information regarding a soaringweb airspace release.
// Location is the URL of the file containing the data,
// Region is one of 'Regions', Date is the date of the release.
type Release struct {
	Location string
	Region   string
	Date     time.Time
}

// baseURL is the default base URL to parse airspace releases from.
var baseURL = "http://soaringweb.org/Airspace"

// timeFormats are the supported formats for dates in soaringweb page.
var timeFormats = []string{"02 January 2006", "02 January, 2006"}

// reDate is a regexp to detect date values in soaringweb pages.
var reDate = regexp.MustCompile(`.*\[\s*\w*\s*(\d\d\s\w+,?\s\d\d\d\d)\s*].*`)

// reLocation is a regexp to detect OpenAir file URLs in soaringweb pages.
var reLocation = regexp.MustCompile(`.*txt.*`)

// testDate is a global value used in parseNode to detect update dates.
// FIXME: we should not need this, it's mainly a replacement for nil compare
var testDate time.Time

// SoaringWeb is the plugin implementation to collect soaringweb.org info.
//
// BaseURL is the prefix to add 'region' to get the URL of the releases page
// ( eg. soaringweb.org/Airspace, which is the default ).
type SoaringWeb struct {
	BaseURL string
}

// Init follows the plugin.Plugin interface (see plugin.Pluginer).
func (sw *SoaringWeb) Init(Params map[string]string) error {
	if Params["BaseURL"] != "" {
		sw.BaseURL = Params["BaseURL"]
	} else {
		sw.BaseURL = baseURL
	}

	return nil
}

// GetAirspace follows Airspace.GetAirspace().
func (sw *SoaringWeb) GetAirspace(regions []string, updatedSince time.Time) ([]common.Airspace, error) {
	var result []common.Airspace

	releases, err := sw.list(sw.BaseURL, regions)
	if err != nil {
		return nil, err
	}
	for r := range releases {
		release := releases[r]
		if release.Date.After(updatedSince) {
			var airspaces []common.Airspace
			airspaces, err = openair.Fetch(release.Location)
			if err != nil {
				// retry by prefixing the base url
				airspaces, err = openair.Fetch(sw.BaseURL + "/" + release.Location)
				if err != nil {
					return nil, err
				}
			}
			result = append(result, airspaces...)
		}
	}

	return result, nil
}

// PutAirspace follows Airspacer.PutAirspace().
func (sw *SoaringWeb) PutAirspace(airspace []common.Airspace) error {
	return nil
}

// list returns latest available airspace information (releases) for the given regions.
//
// basepath is the base for the url ( eg. soaringweb.org/Airspace ), regions are the
// regions to be retrieved ( eg. ['AT', 'FR'] ).
func (sw *SoaringWeb) list(basepath string, regions []string) ([]Release, error) {
	var releases []Release

	if regions == nil {
		regions = Regions
	}
	var content []byte
	for i := range regions {
		location := strings.Join([]string{basepath, regions[i]}, "/")
		// case http
		resp, err := http.Get(location)
		if err == nil {
			defer resp.Body.Close()
			content, err = ioutil.ReadAll(resp.Body)
		} else { // case file
			resp, err := ioutil.ReadFile(location)
			if err != nil {
				return nil, err
			}
			content = resp
		}
		var items []Release
		items, _ = sw.parse(basepath, regions[i], content)
		releases = append(releases, items...)
	}
	return releases, nil
}

// parse finds the openair releases location in the given soaringweb region page.
//
// region is one of 'AT', 'FR', etc and is used to build the Release object.
// content is the html content of the soaringweb region page.
//
// Returns an array of Release objects correspoding to the given region.
func (sw *SoaringWeb) parse(basepath string, region string, content []byte) ([]Release, error) {
	var releases []Release
	var location = ""
	date := testDate
	// No check for err as html.Parse does a very good job parsing broken docs
	z, _ := html.Parse(strings.NewReader(string(content)))
	sw.parseNode(z, &releases, basepath, region, &location, &date)

	return releases, nil
}

// parseNode is used to recursively parse the soaringweb html document.
//
// a bit complex but couldn't make it work with regexp. parses the document
// depth-first and picks dates and locations as it founds them. every time a
// location is found and there is a date already picked a new release object
// is added to the list.
func (sw *SoaringWeb) parseNode(n *html.Node, releases *[]Release, basepath string, region string, location *string, date *time.Time) {
	var err error
	if n.Type == html.ElementNode && n.Data == "a" {
		for attr := range n.Attr {
			if n.Attr[attr].Key == "href" && reLocation.MatchString(n.Attr[attr].Val) {
				*location = n.Attr[attr].Val
			}
		}

	} else if n.Type == html.TextNode {
		match := reDate.FindStringSubmatch(n.Data)
		if match != nil {
			for f := range timeFormats {
				*date, err = time.Parse(timeFormats[f], match[1])
				if err == nil {
					break
				}
			}
		}
	}
	if *date != testDate && *location != "" {
		*releases = append(*releases, Release{Location: *location, Date: *date, Region: region})
		*location = "" // so that we don't keep adding releases in the next child
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		sw.parseNode(c, releases, basepath, region, location, date)
	}
}
