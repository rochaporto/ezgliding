package com.ezgliding.igc;

import java.util.Date;

public class Fix {
	
	public Date date;
	private double lat, latrd;
	private double lon, lonrd;
	public int pressureAlt;
	public int gnssAlt;
	public char validity;

	public Fix(Date date, double lat, double lon, int pressureAlt, int gnssAlt, char validity) {
		this.date = date;
		setLat(lat);
		setLon(lon);
		this.pressureAlt = pressureAlt;
		this.gnssAlt = gnssAlt;
		this.validity = validity;
	}

	/**
	 * Like equal(), but ignoring the date, and usage of pressureAlt is optional.
	 */
	public boolean equivalent(Fix other, boolean withAlt) {
		if (latrd != other.latrd || lonrd != other.lonrd)
			return false;
		if (withAlt && pressureAlt != other.pressureAlt)
			return false;
		return true;
	}

	public boolean equals(Fix other) {
		if (!date.equals(other.date) || latrd != other.latrd || lonrd != other.lonrd
			|| pressureAlt != other.pressureAlt || gnssAlt != other.gnssAlt
			|| validity != other.validity) 
			return false;
		
		return true;
	}

	public void setLat(double degrees) {
		lat = degrees;
		latrd = Math.toRadians(lat);
	}

	public void setLon(double degrees) {
		lon = degrees;
		lonrd = Math.toRadians(lon);
	}

	public double lat() { return lat; }
	
	public double latrd() { return latrd; }

	public double lon() { return lon; }

	public double lonrd() { return lonrd; }

	public Fix clone() {
		return new Fix((Date)date.clone(), lat, lon, pressureAlt, gnssAlt, validity);
	}

	public String toString() {
		return "Lat(rd): " + lat + "(" + latrd + ") Lon(rd): " + lon + "(" + lonrd
			+ ") Alt(gnss): " + pressureAlt + "(" + gnssAlt + ")"; 
	}
}
