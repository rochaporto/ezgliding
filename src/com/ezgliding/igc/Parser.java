package com.ezgliding.igc;

import java.io.BufferedReader;
import java.io.File;
import java.io.IOException;
import java.nio.charset.Charset;
import java.nio.file.Files;
import java.nio.file.FileSystems;
import java.nio.file.Path;
import java.text.ParseException;
import java.util.Calendar;
import java.util.TimeZone;

public class Parser {

	private enum HSubType { 
		DTE, FXA, PLT, CM2, GTY, GID, DTM, RFW, RHW, FTY, GPS, PRS, CID, CCL
	}

	private static Calendar cal = Calendar.getInstance(TimeZone.getTimeZone("UTC"));

	static { cal.set(Calendar.MILLISECOND, 0); }

	public Parser() {

	}

	public Flight parse(Path location) throws IOException, ParseException {
		BufferedReader reader;
		try {
			reader = Files.newBufferedReader(location, Charset.defaultCharset());
		} catch (IOException e) { 
			throw new IOException("Failed to open location: " + location + "\n" + e);

		}

		String content = "", line;
		try {
			while ((line = reader.readLine()) != null)
				content += line + "\n";
		} catch (IOException e) { 
			throw new IOException("Error reading file: " + location + "\n" + e);
		}

		return parse(content);
	}

	public Flight parse(String content) throws ParseException {
		if (content == null) return null;

		Flight flight = new Flight();

		String[] lines = content.split("\n");
		for (String line: lines)
			parseLine(line, flight);

		return flight;
	}	

	private void parseLine(String line, Flight flight) throws ParseException {
		if (line == null || flight == null) return;

		char type = line.charAt(0);
		switch(type) {
			case 'A':
				parseA(line, flight);
				break;
			case 'B': 
				parseB(line, flight);
				break;
			case 'H': 
				parseH(line, flight);
				break;
			default: break;
		}
	}

	private void parseA(String line, Flight flight) throws ParseException {
		if (line == null || flight == null) return;

		flight.setManufacturer(line.substring(1,4));
		flight.setUniqueId(line.substring(4,7));
		flight.setAdditionalData(line.substring(7));
	}

	private void parseB(String line, Flight flight) throws ParseException {
		if (line == null || flight == null) return;

		Fix fix = new Fix(
			Integer.parseInt(line.substring(1,3))*3600 
				+ Integer.parseInt(line.substring(3,5))*60
				+ Integer.parseInt(line.substring(5,7)),
			Util.minDec2decimal(line.substring(7,15)),
			Util.minDec2decimal(line.substring(15,24)),
			Integer.parseInt(line.substring(25,30)),
			Integer.parseInt(line.substring(30,35)),
			line.charAt(24)
		);
		flight.addFix(fix);
	}

	private void parseH(String line, Flight flight) throws ParseException {
		if (line == null || flight == null) return;

		String subType = line.substring(2,5);

		switch (HSubType.valueOf(subType)) {
			case DTE:
				cal.set(
					Integer.parseInt(line.substring(9,11)), 
					Integer.parseInt(line.substring(7,9))-1, 
					Integer.parseInt(line.substring(5,7)),
					0, 0, 0);
				flight.setDate(cal.getTime());
				break;
			case FXA:
				flight.setFixAccuracy(Integer.parseInt(line.substring(5,8)));
				break;
			case PLT:
				flight.setPilot(line.substring(line.indexOf(":")+1));
				break;
			case CM2:
				flight.setCrew2(line.substring(line.indexOf(":")+1));
				break;
			case GTY:
				flight.setGliderType(line.substring(line.indexOf(":")+1));
				break;
			case GID:
				flight.setGliderId(line.substring(line.indexOf(":")+1));
				break;
			case DTM:
				flight.setGpsDatum(line.substring(line.indexOf(":")+1));
				break;
			case RFW:
				flight.setFirmwareVersion(line.substring(line.indexOf(":")+1));
				break;
			case RHW:
				flight.setHardwareVersion(line.substring(line.indexOf(":")+1));
				break;
			case FTY:
				flight.setFrType(line.substring(line.indexOf(":")+1));
				break;
			case GPS:
				flight.setGpsManufacturer(line.substring(line.indexOf(":")+1));
				break;
			case PRS:
				flight.setPressAltSensor(line.substring(line.indexOf(":")+1));
				break;
			case CID:
				flight.setCompetitionId(line.substring(line.indexOf(":")+1));
				break;
			case CCL:
				flight.setCompetitionClass(line.substring(line.indexOf(":")+1));
				break;
			default:
				break;
		}
	}

}
