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

package collection

import "testing"

func TestEcho(t *testing.T) {
  res := Echo("Hello")
  if res != "Hello" {
    t.Errorf("Failed")
  }
}

func TestEchoEmpty(t *testing.T) {
  res := Echo("")
  if res != "" {
    t.Errorf("Failed to Echo empty string :: got '%v'", res)
  }
}
