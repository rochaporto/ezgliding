package com.ezgliding.igc;

import java.util.ArrayList;
import java.util.Date;

public class Task {

	private Date date;

	private Date flightDate;

	private int taskId;

	private String description;

	private WayPoint takeoff;
	
	private WayPoint start;

	private ArrayList<WayPoint> turnPoints;

	private WayPoint finish;

	private WayPoint landing;

	public Task() {
		turnPoints = new ArrayList<WayPoint>();
	}

	public void setDate(Date date) { this.date = date; }

	public Date getDate() { return date; }

	public void setFlightDate(Date flightDate) { this.flightDate = flightDate; }

	public Date getFlightDate() { return flightDate; }

	public void setTaskId(int taskId) { this.taskId = taskId; }

	public int getTaskId() { return taskId; }

	public void setDescription(String description) { this.description = description; }

	public String getDescription() { return description; }

	public void setTakeoff(WayPoint takeoff) { this.takeoff = takeoff; }

	public WayPoint getTakeoff() { return takeoff; }

	public void setStart(WayPoint start) { this.start = start; }

	public WayPoint getStart() { return start; }

	public void setFinish(WayPoint finish) { this.finish = finish; }

	public WayPoint getFinish() { return finish; }

	public void setLanding(WayPoint landing) { this.landing = landing; }

	public WayPoint getLanding() { return landing; }

	public void addTurnPoint(WayPoint turnPoint) { turnPoints.add(turnPoint); }

	public void removeTurnPoint(WayPoint turnPoint) { turnPoints.remove(turnPoint); }

	public ArrayList<WayPoint> getTurnPoints() { return turnPoints; }

	@Override
	public String toString() {
		String str = getDate() + "," + getFlightDate() + ","
			+ getTaskId() + "," + getDescription() + ","
			+ getTakeoff() + "," + getStart();
		for (WayPoint wp: getTurnPoints())
			str += "," + wp;
		str += "," + getFinish() + "," + getLanding();
		return str;
	}

}
