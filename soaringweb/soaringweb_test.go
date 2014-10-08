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

package soaringweb

import (
	"testing"
)

func TestList(t *testing.T) {
	//FIXME:
}

func TestListEmpty(t *testing.T) {
	_, err := List("")
	if err == nil {
		t.Errorf("List empty string should give error")
	}
}

func TestListMissing(t *testing.T) {
	_, err := List("./nonexisting.file")
	if err == nil {
		t.Errorf("List non existing should give error")
	}
}

func TestFetchEmpty(t *testing.T) {
	_, err := Fetch("")
	if err == nil {
		t.Errorf("Fetch empty string should give error")
	}
}

func TestFetchMissing(t *testing.T) {
	_, err := Fetch("./nonexisting.file")
	if err == nil {
		t.Errorf("Fetch non existing should give error")
	}
}
