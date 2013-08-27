package com.ezgliding.igc;

import java.util.List;

public class RectangleSet {

	List<Fix> fixes;

	Fix nw, ne, se, sw;

	Fix[] vertices; // set by setBound()

	private double diagonal = -1;

	public RectangleSet(List<Fix> fixes) {
		this.fixes = fixes;
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

		if (fix.latrd() <= nw.latrd() && fix.latrd() >= sw.latrd()
			&& fix.lonrd() >= nw.lonrd() && fix.lonrd() <= ne.lonrd())
			return true;
		return false;
	}

	public RectangleSet[] split() {
		int mid = fixes.size() / 2;
		return new RectangleSet[] {
			new RectangleSet(fixes.subList(0, mid)),
			new RectangleSet(fixes.subList(mid+1, fixes.size()-1)) };
	}

	public double diagonal() {
		if (diagonal == -1)
			diagonal = Util.distance(nw, se);
		return diagonal;
	}

	public double minDistance(RectangleSet other) { 
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

	public boolean equals(RectangleSet other) {
		Fix[] vertices = getVertices();
		Fix[] otherVertices = other.getVertices();
		for (int i=0; i<vertices.length; i++)
			if (!vertices[i].equals(otherVertices[i]))
				return false;
		return true;	
	}

	private void setBound() {
		if (fixes == null || fixes.size() < 1) return;

		Fix f = fixes.get(0);
		f.pressureAlt = f.gnssAlt = 0;
		nw = f.clone();
		ne = f.clone();
		se = f.clone();
		sw = f.clone();
		for (Fix fix: fixes) {
			if (fix.lat() < se.lat()) {
				se.setLat(fix.lat());
				sw.setLat(fix.lat());
			}
			if (fix.lat() > ne.lat()) {
				ne.setLat(fix.lat());
				nw.setLat(fix.lat());
			}
			if (fix.lon() < nw.lon()) {
				nw.setLon(fix.lon());
				sw.setLon(fix.lon());
			}
			if (fix.lon() > ne.lon()) {
				ne.setLon(fix.lon());
				se.setLon(fix.lon());
			}
		}

		this.vertices = new Fix[] { nw, sw, ne, se };
	}
}
