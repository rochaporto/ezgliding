package com.ezgliding.igc;

import org.junit.Before;
import org.junit.Test;
import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertTrue;

import java.util.Date;

public class FlightTest {

	@Before
	public void setUp() {

	}

	@Test
	public void testAddFix() {
		Flight flight = new Flight();
		assertEquals(0, flight.fixes().size());

		Fix fix = new Fix(0, -45.666, 108.345, 1089, 1200, 'V');
		flight.addFix(fix);
		assertEquals(1, flight.fixes().size());

		assertTrue(fix.equals(flight.fixes().get(0)));
	}
}
