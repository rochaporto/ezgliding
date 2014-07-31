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

package welt2000

import (
	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/rss"
	"time"
)

// Release contains info about a specific release
type Release struct {
	Date      time.Time
	Source    string
	Airfields []common.Airfield
}

// List checks the welt2000 rss feed and lists the releases found
func List() ([]Release, error) {
	feed, err := rss.Fetch("http://www.segelflug.de/vereine/welt2000/content/en/news/updates.xml")
	if err != nil {
		return nil, err
	}

	res := make([]Release, 10)
	for i, item := range feed.Items {
		res[i].Date = item.Date
		res[i].Source = item.Link
	}
	return res, nil
}

// Fetch gets and returns the Release at the given location
func Fetch(location string) (*Release, error) {
	r := Release{Source: location}
	r.Fetch()
	return &r, nil
}

// Fetch fills up the Release object with data after parsing the content at Release.Source
func (r *Release) Fetch() error {
	return nil
}
