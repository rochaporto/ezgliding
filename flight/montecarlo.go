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

import (
	"fmt"
	"math/rand"

	"github.com/rochaporto/ezgliding/spatial"
)

// Montecarlo is the montecarlo based implementation of the flight Optimizer.
type Montecarlo struct {
	Cycles    int
	MCCycles  int
	MCGuesses int
}

// NewMontecarlo returns a new Montecarlo optimizer instance.
func NewMontecarlo() *Montecarlo {
	return &Montecarlo{Cycles: 10, MCCycles: 100000, MCGuesses: 0}
}

// Optimize implements Optimizer().
func (mc *Montecarlo) Optimize(track []Point, nTP int, scorer Scorer) (OptResult, error) {
	maxDistChan := make(chan float64, mc.Cycles)
	maxChan := make(chan []int, mc.Cycles)

	for c := 0; c < mc.Cycles; c++ {
		go func() {
			var candidate = make([]int, nTP)
			var max = make([]int, nTP)
			var cdistance float64
			var maxDistance = 0.0

			// start with uniform distribution (equal distance)
			for i := 0; i < nTP; i++ {
				candidate[i] = (len(track) / (nTP - 1)) * i
			}
			cdistance = mc.distance(track, candidate)

			// run montecarlo cycles
			var index, bwp, twp, nwp int
			for i := 0; i < mc.MCCycles; i++ {
				index = rand.Intn(nTP-2) + 1
				bwp = candidate[0]
				if index > 0 {
					bwp = candidate[index-1]
				}
				twp = len(track) - 1
				if index < nTP-1 {
					twp = candidate[index+1]
				}
				nwp = rand.Intn(twp-bwp) + bwp
				candidate[index] = nwp
				//sort.Sort(sort.IntSlice(candidate))
				cdistance = mc.distance(track, candidate)
				if cdistance > maxDistance {
					maxDistance = cdistance
					max = candidate
				}
			}

			maxDistChan <- maxDistance
			maxChan <- max
		}()
	}

	gMax := make([]int, nTP)
	gMaxDistance := 0.0
	var dist float64
	var pts = make([]int, nTP)
	for i := 0; i < mc.Cycles; i++ {
		dist = <-maxDistChan
		pts = <-maxChan
		if dist > gMaxDistance {
			gMaxDistance = dist
			gMax = pts
		}
	}
	fmt.Printf("%v :: %v\n", gMax, len(track))
	result := OptResult{TurnPoints: make([]Point, nTP), Distance: gMaxDistance}
	for i := 0; i < nTP; i++ {
		result.TurnPoints[i] = track[gMax[i]]
	}
	return result, nil
}

func (mc *Montecarlo) distance(track []Point, tps []int) float64 {
	distance := 0.0
	for i := 0; i < len(tps)-1; i++ {
		v := spatial.GCDistance(track[tps[i]].Latitude, track[tps[i]].Longitude,
			track[tps[i+1]].Latitude, track[tps[i+1]].Longitude)
		distance += v
	}
	return distance
}
