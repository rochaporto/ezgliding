package com.ezgliding.igc;

public class WayPoint {

	private Fix point;
	
	private String description;

	public WayPoint(Fix point, String description) {
		this.point = point;
		this.description = description;
	}

	public Fix getPoint() { return point; }

	public String getDescription() { return description; }

	@Override
	public String toString() { 
		return "{" + point + ",\"" + description + "\"}";
	}

}
