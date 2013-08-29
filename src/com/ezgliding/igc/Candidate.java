package com.ezgliding.igc;

import java.util.ArrayList;
import java.util.List;
import java.util.Iterator;

public class Candidate implements Comparable<Candidate> {
	
	ArrayList<RectangleSet> rectangles;

	private double max = 0.0;

	private double min = Double.MAX_VALUE;

	public Candidate() {
		this(null);
	}

	public Candidate(ArrayList<RectangleSet> inputRectangles) {
		this.rectangles = new ArrayList<RectangleSet>();

		if (inputRectangles != null)
			this.rectangles.addAll(inputRectangles);
		reset();
	}

	public double max() {
		max = 0.0;
		for (int i=0; i<rectangles.size()-1; i++)
			max += rectangles.get(i).maxDistance(rectangles.get(i+1));
		return max;
	}

	public double min() {
		min = 0.0;
		for (int i=0; i<rectangles.size()-1; i++)
			min += rectangles.get(i).minDistance(rectangles.get(i+1));
		return min;
	}

	public boolean isFinal() {
		for (RectangleSet set: getRectangles())
			if (set.getFixes().size() != 1) return false;
		return true;
	}

	public List<RectangleSet> getRectangles() {
		return rectangles;
	}

	public void add(RectangleSet rSet) {
		if (rSet == null) 
			throw new IllegalArgumentException("Cannot add empty set");
		rectangles.add(rSet);
		reset();
	}

	public void replace(int i, RectangleSet newSet) {
		getRectangles().set(i, newSet);
		reset();
	}

	public void replace(int i, RectangleSet[] newSets) {
		List<RectangleSet> sets = getRectangles();

		if (i<0 || i>sets.size()-1)
			throw new IllegalArgumentException("Invalid index given");
		if (newSets == null || newSets.length == 0)
			throw new IllegalArgumentException("No new sets provided, cannot replace");

		sets.remove(i);
		for (RectangleSet newSet: newSets)
			sets.add(i, newSet);	
		reset();
	}

	public int largestDiagonal() {
		List<RectangleSet> sets = getRectangles();
		if (sets.size() == 0) return -1;

		int larger = 0;
		for (int i=1; i<sets.size(); i++)
			if (sets.get(i).diagonal() > sets.get(larger).diagonal())
				larger = i;

		return larger;
	}

	private void reset() {
		max = 0.0; min = Double.MAX_VALUE;
	}

	@Override
	public int compareTo(Candidate other) {
		double diff = max() - other.max();
		if (diff == 0) return 0;
		else if (diff > 0) return 1;
		return -1;
	}

	@Override
	public boolean equals(Object other) {
		if (compareTo((Candidate)other) == 0) return true;
		return false;
	}

	@Override
	public Candidate clone() {
		return new Candidate(this.rectangles);
	}

	@Override
	public String toString() { 
		String str = String.format("%1$,.2f", min()) + " " + String.format("%1$,.2f", max()) 
			+ " " + isFinal() + " " + largestDiagonal() + " {";
		for (RectangleSet set: getRectangles())
			str += set + ",";
		return str + "}";
	}
}

