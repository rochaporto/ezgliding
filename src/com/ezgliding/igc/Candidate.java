package com.ezgliding.igc;

public class Candidate implements Comparable<Candidate> {
	
	RectangleSet[] rectangles;

	boolean isFinal;

	Candidate(RectangleSet[] rectangles) {
		this.rectangles = rectangles;

		isFinal = true;
		for (RectangleSet set: rectangles) 
			if (set.fixes.size() > 1) {
				isFinal = false;
				break;
			}
	}

	public double max() {
		double max = 0.0;
		for (int i=0; i<rectangles.length-1; i++)
			max += rectangles[i].maxDistance(rectangles[i+1]);
		return max;
	}

	public double min() {
		double min = 0;
		for (int i=0; i<rectangles.length-1; i++)
			if (!rectangles[i].overlap(rectangles[i+1]))
				min += rectangles[i].minDistance(rectangles[i+1]);
		return min;
	}

	public boolean isFinal() {
		return isFinal;			
	}

	public RectangleSet[] getRectangles() {
		return rectangles;
	}

	public int compareTo(Candidate other) {
		double diff = max() - other.max();
		if (diff == 0) return 0;
		else if (diff > 0) return 1;
		return -1;
	}

	public boolean equals(Candidate other) {
		if (compareTo(other) == 0) return true;
		return false;
	}

	public String toString() { 
		return rectangles.length + "";
	}
}

