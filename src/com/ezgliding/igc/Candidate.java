package com.ezgliding.igc;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Iterator;

public class Candidate implements Comparable<Candidate> {
	
	ArrayList<RectangleSet> rectangles;

	private double max;

	private double min;

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
		if (max == 0.0)
			for (int i=0; i<rectangles.size()-1; i++)
				max += rectangles.get(i).maxDistance(rectangles.get(i+1));
		return max;
	}

	public double min() {
		if (min == 0.0)
			for (int i=0; i<rectangles.size()-1; i++)
				min += rectangles.get(i).minDistance(rectangles.get(i+1));
		return min;
	}

	public boolean isFinal() {
		for (RectangleSet set: getRectangles())
			if (set.numFixes() != 1) return false;
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

	public void replace(RectangleSet oldSet, RectangleSet newSet) {
		replace(oldSet, new RectangleSet[] { newSet });
	}

	public void replace(RectangleSet oldSet, RectangleSet[] newSets) {
		List<RectangleSet> sets = getRectangles();

		if (newSets == null || newSets.length == 0)
			throw new IllegalArgumentException("No new sets provided, cannot replace");

		while (sets.remove(oldSet)) continue;
		for (RectangleSet newSet: newSets)
			sets.add(newSet);	
		reset();
	}

	public RectangleSet largestDiagonal() {
		List<RectangleSet> sets = getRectangles();
		if (sets.size() == 0) return null;

		RectangleSet result = sets.get(0);
		for (int i=1; i<sets.size(); i++)
			if (sets.get(i).diagonal() > result.diagonal())
				result = sets.get(i);

		return result;
	}

	private void reset() {
		max = 0.0; min = 0.0;
		Collections.sort(rectangles);
	}

	@Override
	public int compareTo(Candidate other) {
		double diff = max() - other.max();
		if (this.equals(other) && diff == 0)
			return 0;
		else if (diff > 0) return 1;
		else return -1;
	}

	@Override
	public boolean equals(Object other) {
		Candidate otherC = (Candidate)other;
		if (otherC.getRectangles().size() != getRectangles().size())
			return false;
		for (int i=0; i<getRectangles().size(); i++)
			if (!getRectangles().get(i).equals(otherC.getRectangles().get(i)))
				return false;
		return true;
	}

	@Override
	public int hashCode() {
		int result = 0;
		for (RectangleSet set: getRectangles())
			result += set.start;
		return result;
	}	

	@Override
	public Candidate clone() {
		return new Candidate(this.rectangles);
	}

	@Override
	public String toString() { 
		String str = "Candidate (" 
			+ String.format("%1$,.2f", min()) + " " + String.format("%1$,.2f", max()) 
			+ " " + isFinal() + " " + largestDiagonal() + ")\n";
		str += "{";
		for (int i=0; i<getRectangles().size(); i++)
			str += "\n\tRectangle " + i + " {\n" + getRectangles().get(i) + "\n\t},";
		return str + "\n}";
	}
}

