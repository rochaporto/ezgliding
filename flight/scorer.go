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

package flight

// Scorer returns the score for the given set of turn points.
// The number of turn points given is variable and the resulting score will
// depend on it. As an example, the netcoupe scorer will take any of 1, 2 or
// 3 turn points (but not more), and for 2TP (triangle) will give a higher
// score if it is an FAI triangle (for the same total distance between TPs).
type Scorer interface {
	Score(turnPoints []Point) (float64, error)
}
