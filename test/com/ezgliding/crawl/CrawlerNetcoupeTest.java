package com.ezgliding.crawl;

import org.junit.Before;
import org.junit.Test;
import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertNotNull;
import static org.junit.Assert.assertTrue;

import java.util.Calendar;
import java.util.TimeZone;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class CrawlerNetcoupeTest {

	private CrawlerNetcoupe crawlerNetcoupe;

	private static Calendar cal = Calendar.getInstance(TimeZone.getTimeZone("UTC"));

	static { cal.set(Calendar.MILLISECOND, 0); }

	@Before
	public void setUp() {
		crawlerNetcoupe = new CrawlerNetcoupe("file:///home/rocha/ws/ezgliding/test/data");
	}

	@Test
	public void testCreation() {
		CrawlerNetcoupe crawler = new CrawlerNetcoupe("file:///home/ezgliding/test");
	}

	@Test(expected=IllegalArgumentException.class)
	public void testCreationNull() {
		CrawlerNetcoupe crawler = new CrawlerNetcoupe(null);
	}

	
	@Test
	public void testGetFlight() {
		FlightEntry flight = crawlerNetcoupe.getFlight(
			"file:///home/rocha/ws/ezgliding/test/data/crawler-netcoupe-flight-detail.html"); //TODO: make it path generic
		assertEquals("EZ PILOT", flight.getPilot());
		assertEquals("EZ CATEGORY", flight.getCategory());
		assertEquals("EZ CLUB", flight.getClub());
		assertEquals("EZ REGION", flight.getRegion());
		assertEquals("EZ COUNTRY", flight.getCountry());
		cal.set(2003, 1, 1, 0, 0, 0);
		assertEquals(cal.getTime(), flight.getDate());
		assertEquals("EZ AIRFIELD", flight.getAirfield());
		assertEquals("EZ REGION", flight.getRegion());
		assertEquals("EZ COUNTRY", flight.getCountry());
		assertEquals(123.45, flight.getDistance(), 0.01);
		assertEquals(234.56, flight.getPoints(), 0.01);
		assertEquals("EZ GLIDER", flight.getGliderType());
		assertEquals(FlightEntry.CircuitType.FREE, flight.getCircuitType());
		assertEquals("http://netcoupe.net/Download/DownloadIGC.aspx?FileID=01234", flight.getFileLocation());
		assertEquals(12.34, flight.getSpeed(), 0.01);
		assertEquals("EZ START", flight.getTask().getStart().getDescription());
		assertEquals("EZ TP1", flight.getTask().getTurnPoints().get(0).getDescription());
		assertEquals("EZ TP2", flight.getTask().getTurnPoints().get(1).getDescription());
		assertEquals("EZ TP3", flight.getTask().getTurnPoints().get(2).getDescription());
		assertEquals("EZ FINISH", flight.getTask().getFinish().getDescription());
		assertEquals("EZ COMMENT", flight.getComment());
	}
}
