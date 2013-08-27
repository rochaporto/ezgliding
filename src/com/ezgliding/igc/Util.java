package com.ezgliding.igc;

public class Util {

	public static int earthRadius = 6371;

	public static double minDec2decimal(String minDec) {
		if (minDec == null || !(minDec.length() == 8 || minDec.length() == 9)) 
			throw new IllegalArgumentException(
				"Expected format is 'DDMMmmm' for lat or 'DDDMMmmm' for lon, for '" + minDec + "'");

		double decimal;
		char cardinal = minDec.charAt(minDec.length()-1);

		if (cardinal == 'N' || cardinal == 'S')
			decimal = Integer.parseInt(minDec.substring(0,2)) 
				+ (( Integer.parseInt(minDec.substring(2,4)) + (Integer.parseInt(minDec.substring(4,7)) / 1000.0)) / 60.0);
		else
			decimal = Integer.parseInt(minDec.substring(0,3)) 
				+ (( Integer.parseInt(minDec.substring(3,5)) + (Integer.parseInt(minDec.substring(5,8)) / 1000.0)) / 60.0);
		if (cardinal == 'S' || cardinal == 'W')
			decimal = decimal * -1;

		return decimal;
	}

	public static double distance(Fix fix1, Fix fix2) {
		//d = 2*asin(sqrt((sin((lat1-lat2)/2))^2 + cos(lat1)*cos(lat2)*(sin((lon1-lon2)/2))^2))
		return 2 * Math.asin(
			Math.sqrt( Math.pow((Math.sin( (fix1.latrd() - fix2.latrd()) / 2) ), 2)
				+ Math.cos(fix1.latrd()) * Math.cos(fix2.latrd()) * Math.pow( (Math.sin( (fix1.lonrd()-fix2.lonrd()) / 2) ), 2))
			) * earthRadius;
	}
}
