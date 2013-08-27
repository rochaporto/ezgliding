package com.ezgliding.igc;

import java.util.ArrayList;
import java.util.List;
import java.util.Iterator;

public class Candidate implements Comparable<Candidate> {
	
	ArrayList<RectangleSet> rectangles;

	boolean isFinal;

	public Candidate() {
		this(null);
	}

	public Candidate(ArrayList<RectangleSet> inputRectangles) {
		this.rectangles = new ArrayList<RectangleSet>();

		isFinal = true;
		if (inputRectangles != null)
			this.rectangles.addAll(inputRectangles);
	}

	public double max() {
		double max = 0.0;
		for (int i=0; i<rectangles.size()-1; i++)
			max += rectangles.get(i).maxDistance(rectangles.get(i+1));
		return max;
	}

	public double min() {
		double min = 0;
		RectangleSet rSet;
		for (int i=0; i<rectangles.size()-1; i++) {
			rSet = rectangles.get(i);
			if (!rSet.overlap(rectangles.get(i+1)))
				min += rSet.minDistance(rectangles.get(i+1));
		}
		return min;
	}

	public boolean isFinal() {
		return isFinal;			
	}

	public List<RectangleSet> getRectangles() {
		return rectangles;
	}

	public void add(RectangleSet rSet) {
		if (rSet == null) return;
		rectangles.add(rSet);
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

	public Candidate clone() {
		return new Candidate(this.rectangles);
	}

	public String toString() { 
		return rectangles.size() + "";
	}
}

