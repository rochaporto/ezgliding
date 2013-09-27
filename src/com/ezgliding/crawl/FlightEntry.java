package com.ezgliding.crawl;

import com.ezgliding.igc.Flight;
import com.ezgliding.igc.Task;

public class FlightEntry extends Flight {

	public enum CircuitType {
		FREE, SET, COMPETITION
	}

	private String id;

	private String category;

	private String club;

	private String region;

	private String country;

	private String airfield;

	private float distance;

	private float points;

	private CircuitType circuitType;

	private String fileLocation;

	private float speed;

	private String comment;

	private String endpoint;
	
	public FlightEntry() {

	}

	public void setId(String id) { this.id = id; }

	public String getId() { return id; }

	public void setCategory(String category) { this.category = category; }

	public String getCategory() { return category; }

	public void setClub(String club) { this.club = club; }

	public String getClub() { return club; }

	public void setRegion(String region) { this.region = region; }

	public String getRegion() { return region; }

	public void setCountry(String country) { this.country = country; }

	public String getCountry() { return country; }

	public void setAirfield(String airfield) { this.airfield = airfield; }

	public String getAirfield() { return airfield; }

	public void setDistance(float distance) { this.distance = distance; }

	public float getDistance() { return distance; }

	public void setPoints(float points) { this.points = points; }

	public float getPoints() { return points; }

	public void setCircuitType(CircuitType circuitType) { this.circuitType = circuitType; }

	public CircuitType getCircuitType() { return circuitType; }

	public void setFileLocation(String fileLocation) { this.fileLocation = fileLocation; }

	public String getFileLocation() { return fileLocation; }

	public void setSpeed(float speed) { this.speed = speed; }

	public float getSpeed() { return speed; }

	public void setComment(String comment) { this.comment = comment; }

	public String getComment() { return comment; }

}
