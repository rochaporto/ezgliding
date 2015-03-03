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

import "github.com/rochaporto/ezgliding/config"

const (
	// DECLARED is set when a flight was pre declared
	DECLARED = iota << 1
)

// Netcoupe is the www.netcoupe.net Contest implementation.
type Netcoupe struct {
}

// NewNetcoupe returns a new Netcoupe instance.
func NewNetcoupe(cfg config.Config) *Netcoupe {
	return &Netcoupe{}
}

// Points implements Contest.Points().
func (nc *Netcoupe) Points(tps []Point, flags int) (Result, error) {
	//FIXME: implement
	return Result{}, nil
}
