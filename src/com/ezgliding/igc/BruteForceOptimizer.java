package com.ezgliding.igc;

import java.util.HashMap;
import java.util.List;

public class BruteForceOptimizer extends Optimizer {

	private HashMap<String,Double> distCache = new HashMap<String,Double>();

	private List<Fix> fixes;

	public BruteForceOptimizer(Flight flight, int numPoints) {
		super(flight, numPoints);
	
		if (flight != null) fixes = flight.fixes();
	}

	public Result optimize() {
		switch (numPoints) {
			case 3: return optimize1TP();
			case 4: return optimize2TP();
			case 5: return optimize3TP();
		}
		return null;
	}

	private double distance(int i, int j) {
		Double distance = distCache.get(i + ":" + j);
		if (distance == null) {
			distance = Util.distance(fixes.get(i), fixes.get(j));
			distCache.put(i + ":" + j, distance);
		}
		return distance;
	}

	private Result optimize1TP() {

		double max = 0.0;
		Fix[] result = new Fix[3];

		double distance = 0.0;
		List<Fix> fixes = flight.fixes().subList(flightStart(), flightEnd());
		for (int i=0; i<fixes.size()-2; i++) {
			for (int j=i+1; j<fixes.size()-1; j++) {
				for (int z=j+1; z<fixes.size(); z++) {
					distance = distance(i, j) + distance(j, z);
					if (distance > max) {
						max = distance;
						result[0] = fixes.get(i);
						result[1] = fixes.get(j);
						result[2] = fixes.get(z);
					}
				}
			}
		}

		return new Result(result);
	}

	private Result optimize2TP() {

		double max = 0.0;
		Fix[] result = new Fix[4];

		double distance = 0.0;
		List<Fix> fixes = flight.fixes().subList(flightStart(), flightEnd());
		for (int i=0; i<fixes.size()-3; i++) {
			for (int j=i+1; j<fixes.size()-2; j++) {
				for (int w=j+1; w<fixes.size()-1; w++) {
					for (int z=w+1; z<fixes.size(); z++) {
						distance =
							distance(i, j) + distance(j, w) + distance(w, z);
						if (distance > max) {
							max = distance;
							result[0] = fixes.get(i);
							result[1] = fixes.get(j);
							result[2] = fixes.get(w);
							result[3] = fixes.get(z);
						}
					}
				}
			}
		}

		return new Result(result);
	}

	private Result optimize3TP() {

		double max = 0.0;
		Fix[] result = new Fix[5];

		double distance = 0.0;
		List<Fix> fixes = flight.fixes().subList(flightStart(), flightEnd());
		for (int i=0; i<fixes.size()-4; i++) {
			for (int j=i+1; j<fixes.size()-3; j++) {
				for (int w=j+1; w<fixes.size()-2; w++) {
					for (int y=w+1; y<fixes.size()-1; y++) {
						for (int z=y+1; z<fixes.size(); z++) {
							distance = 
								distance(i, j) + distance(j, w)
								+ distance(w, y) + distance(y, z);
							if (distance > max) {
								max = distance;
								result[0] = fixes.get(i);
								result[1] = fixes.get(j);
								result[2] = fixes.get(w);
								result[3] = fixes.get(y);
								result[4] = fixes.get(z);
							}
						}
					}
				}
			}
		}

		return new Result(result);
	}

}
