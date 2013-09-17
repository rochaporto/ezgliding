package com.ezgliding.igc;

import java.util.List;

public class RectangleSet implements Comparable<RectangleSet> {

	List<Fix> fixes;

	List<Fix> fixesSublist;

	Fix[] vertices; // set by setBound()

	int start;
	
	int end;

	private double diagonal = -1;

	public RectangleSet(List<Fix> fixes) {
		this(fixes, 0, fixes == null ? 0 : fixes.size());
	}

	public RectangleSet(List<Fix> fixes, int start, int end) {
		this.fixes = fixes;
		this.start = start;
		this.end = end;
		if (fixes != null)
			this.fixesSublist = fixes.subList(start, end);
		setBound();
	}
	
	public boolean overlap(RectangleSet other) { 
		if (other == null) return false;

		for (Fix vertice: other.getVertices())
			if (contains(vertice)) return true;
		return false;
	}

	public boolean contains(Fix fix) {
		if (fix == null) return false;

		if (fix.latrd() <= nw().latrd() && fix.latrd() >= sw().latrd()
			&& fix.lonrd() >= nw().lonrd() && fix.lonrd() <= ne().lonrd())
			return true;
		return false;
	}

	public RectangleSet[] split() {
		int mid = start + ((end-start) / 2);
		RectangleSet[] sets = new RectangleSet[] {
			new RectangleSet(fixes, start, mid),
			new RectangleSet(fixes, mid, end) };
		return sets;
	}

	public double diagonal() {
		if (diagonal == -1)
			diagonal = Util.distance(nw(), se());
		return diagonal;
	}

	public double minDistance(RectangleSet other) { 
		if (overlap(other)) return 0.0;

		double min = Double.MAX_VALUE;
		Fix[] sVertices = getVertices();
		Fix[] dVertices = other.getVertices();

		double dist;
		for (int i=0; i<sVertices.length; i++)
			for (int j=0; j<dVertices.length; j++) {
				dist = Util.distance(sVertices[i], dVertices[j]);
				if (dist < min) min = dist;	
			}
		return min;
	}

	public double maxDistance(RectangleSet other) { 
		double max = -1;
		Fix[] sVertices = getVertices();
		Fix[] dVertices = other.getVertices();
		
		double dist;
		for (int i=0; i<sVertices.length; i++) {
			for (int j=0; j<dVertices.length; j++) {
				dist = Util.distance(sVertices[i], dVertices[j]);
				if (dist > max) max = dist; 
			}
		}
		return max;
	}

	public Fix[] getVertices() {
		return vertices; 
	}

	public List<Fix> getFixes() {
		return fixesSublist;
	}

	@Override
	public int compareTo(RectangleSet other) {
		if (start == other.start()) return 0;
		else if (start > other.start()) return 1;
		else return -1;
	}

	@Override
	public boolean equals(Object otherO) {
		if (otherO == null) return false;
		Fix[] vertices = getVertices();
		Fix[] otherVertices = ((RectangleSet)otherO).getVertices();
		for (int i=0; i<vertices.length; i++)
			if (!vertices[i].equivalent(otherVertices[i], false))
				return false;
		return true;	
	}

	public Fix ne() { return getVertices()[0]; }

	public Fix se() { return getVertices()[1]; }

	public Fix nw() { return getVertices()[2]; }

	public Fix sw() { return getVertices()[3]; }

	public int start() { return start; }
	
	public int end() { return end; }

	public int numFixes() { return end() - start(); }

	private void setBound() {
		if (fixes == null || fixes.size() < 1) return;

		Fix f = fixes.get(start);
		f.pressureAlt = f.gnssAlt = 0;
		this.vertices = new Fix[] { f.clone(), f.clone(), f.clone(), f.clone() };

		Fix fix;
		for (int i=start; i<end; i++) {
			fix = fixes.get(i);
			if (fix.lat() < se().lat()) {
				se().setLat(fix.lat());
				sw().setLat(fix.lat());
			}
			if (fix.lat() > ne().lat()) {
				ne().setLat(fix.lat());
				nw().setLat(fix.lat());
			}
			if (fix.lon() < nw().lon()) {
				nw().setLon(fix.lon());
				sw().setLon(fix.lon());
			}
			if (fix.lon() > ne().lon()) {
				ne().setLon(fix.lon());
				se().setLon(fix.lon());
			}
		}
	}

	@Override
	public String toString() {
		String str = "\t(" + start + "," + end + ")\n";
		Fix fix;
		for (int i=start; i<end; i++) {
			fix = fixes.get(i);
			str += "\t" + fix + ",\n";
		}
		return str;
	}
}
