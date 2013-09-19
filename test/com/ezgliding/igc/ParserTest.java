package com.ezgliding.igc;

import org.junit.Before;
import org.junit.Test;
import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNotNull;
import static org.junit.Assert.assertTrue;

import java.io.IOException;
import java.nio.file.FileSystems;
import java.text.ParseException;
import java.util.Calendar;
import java.util.TimeZone;

import com.ezgliding.igc.Parser;

public class ParserTest {

	private Flight flight;

	private Fix[] fixes;

	private static Calendar cal = Calendar.getInstance(TimeZone.getTimeZone("UTC"));

	static { cal.set(Calendar.MILLISECOND, 0); }

	@Before
	public void setUp() throws IOException, ParseException {
		Parser parser = new Parser();
		flight = parser.parse(
			FileSystems.getDefault().getPath("test/data/parse-basic-records.igc"));
		fixes = new Fix[] {
			new Fix(getTime(9,9,9), Util.minDec2decimal("4505005N"), 
				Util.minDec2decimal("00505005E"), 1111, 1112, 'A'),
			new Fix(getTime(10,10,10), Util.minDec2decimal("4606006N"), 
				Util.minDec2decimal("00606006E"), 12222, 12223, 'V'),
			new Fix(getTime(11,11,11), Util.minDec2decimal("4707007N"), 
				Util.minDec2decimal("00707007E"), 23333, 23334, 'A'),
		};
	}

	@Test(expected=IOException.class)
	public void testPathNotExists() throws IOException {
		Parser parser = new Parser();
		try {
			Flight flight = parser.parse(
				FileSystems.getDefault().getPath("file.does.not.exist.at.all"));
		} catch(ParseException e) { }
	}	

	@Test
	public void testParseA() {
		assertEquals("MAN", flight.getManufacturer());
		assertEquals("UID", flight.getUniqueId());
		assertEquals("MOREDATA", flight.getAdditionalData());
	}

	@Test
	public void testParseB() {
		assertEquals(fixes.length, flight.fixes().size());
		for (int i=0; i<fixes.length; i++)
			assertEquals(fixes[i], flight.fixes().get(i));
	}

	@Test 
	public void testParseC() {
		Task task = flight.getTask();
		WayPoint[] taskPoints = new WayPoint[] {
			new WayPoint(new Fix(getTime(0,0,0), Util.minDec2decimal("4505005N"), 
				Util.minDec2decimal("00505005E"), 0, 0, 'A'), "EZTAKEOFF"),
			new WayPoint(new Fix(getTime(0,0,0), Util.minDec2decimal("4606006N"), 
				Util.minDec2decimal("00606006E"), 0, 0, 'A'), "EZSTART"),
			new WayPoint(new Fix(getTime(0,0,0), Util.minDec2decimal("4707007N"), 
				Util.minDec2decimal("00707007E"), 0, 0, 'A'), "EZTP1"),
			new WayPoint(new Fix(getTime(0,0,0), Util.minDec2decimal("4808008N"), 
				Util.minDec2decimal("00808008E"), 0, 1112, 'A'), "EZTP2"),
			new WayPoint(new Fix(getTime(0,0,0), Util.minDec2decimal("4909009N"), 
				Util.minDec2decimal("00909009E"), 0, 0, 'A'), "EZFINISH"),
			new WayPoint(new Fix(getTime(0,0,0), Util.minDec2decimal("5050050N"), 
				Util.minDec2decimal("01010010E"), 0, 0, 'A'), "EZLANDING"),
		};
		cal.set(3,1,1,10,11,12);
		assertEquals(cal.getTime(), task.getDate());
		cal.set(4,2,2,0,0,0);
		assertEquals(cal.getTime(), task.getFlightDate());
		assertEquals(1234, task.getTaskId());
		assertEquals(2, task.getTurnPoints().size());
		assertEquals("EZTASK", task.getDescription());
		assertTrue(taskPoints[0].getPoint().equivalent(task.getTakeoff().getPoint(), false));
		assertTrue(taskPoints[1].getPoint().equivalent(task.getStart().getPoint(), false));
		assertEquals(2, task.getTurnPoints().size());
		assertTrue(taskPoints[2].getPoint().equivalent(task.getTurnPoints().get(0).getPoint(), false));
		assertTrue(taskPoints[3].getPoint().equivalent(task.getTurnPoints().get(1).getPoint(), false));
		assertTrue(taskPoints[4].getPoint().equivalent(task.getFinish().getPoint(), false));
		assertTrue(taskPoints[5].getPoint().equivalent(task.getLanding().getPoint(), false));
	}

	@Test
	public void testParseH() {
		cal.set(3, 1, 1, 0, 0, 0);
		assertEquals(cal.getTime(), flight.getDate());
		assertEquals(123, flight.getFixAccuracy());
		assertEquals("EZPILOT", flight.getPilot());
		assertEquals("EZCREW2", flight.getCrew2());
		assertEquals("EZGLIDER", flight.getGliderType());
		assertEquals("EZGLIDERID", flight.getGliderId());
		assertEquals("WGS-1984", flight.getGpsDatum());
		assertEquals("1.1", flight.getFirmwareVersion());
		assertEquals("2.2", flight.getHardwareVersion());
		assertEquals("EZFRTYPE", flight.getFrType());
		assertEquals("EZGPS", flight.getGpsManufacturer());
		assertEquals("EZPRESSALTSENSOR", flight.getPressAltSensor());
		assertEquals("EZCOMPID", flight.getCompetitionId());
		assertEquals("EZCOMPCLASS", flight.getCompetitionClass());
	}

	private int getTime(int hour, int min, int second) {
		return (hour*3600)+(min*60)+second;
	}

}
