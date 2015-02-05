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

// Optimizer returns the optimized distance and score for the given track.
// nTP is the number of turnpoints to optimize for, and includes start and
// finish. This means 3 for out and return, 4 for triangle, ...
// The optimizer implementation will evaluation most or all combinations
// of track points building a set of valid TPs, and pass that to the scorer.
// The set of TPs with the higher score is returned, along with the actual
// distance between TPs.
type Optimizer interface {
	Optimize(track []Point, nTP int, scorer Scorer) (OptResult, error)
}

// OptResult holds information about a given flight optimization.
type OptResult struct {
	// TurnPoints is the set of turn points in this optimization.
	TurnPoints []Point
	// Distance is the total distance between turnpoints (direct line).
	Distance float64
	// Score is the number of points given by the Scorer implementation.
	Score float64
	// Description is a given task description by the Scorer implementation.
	ScorerID string
}

// Optimizers holds a map of all available Optimizers, keyed on ID.
var Optimizers = map[string]Optimizer{
	"montecarlo": Optimizer(NewMontecarlo()),
}
