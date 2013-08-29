package com.ezgliding.igc;

import org.junit.Before;
import org.junit.Test;
import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertTrue;

import com.ezgliding.igc.Fix;

import java.util.Calendar;
import java.util.Date;

public class FixTest {

	@Before
	public void setUp() {

	}

	@Test
	public void testEquivalentNull() {
		Fix fix1 = new Fix(new Date(), 45.020, 108.000, 0, 0, 'V');
		assertFalse(fix1.equivalent(null, true));
	}
	
	@Test
	public void testEquivalent() {
		Date d = new Date();
		Fix fix1 = new Fix(d, -45.666, 108.345, 267, 288, 'V');
		Fix fix2 = new Fix(new Date(), -45.666, 108.345, 267, 288, 'V');
		
		assertTrue(fix1.equivalent(fix2,false));
		assertTrue(fix1.equivalent(fix2,true));

		fix2.pressureAlt = 1267;
		fix2.gnssAlt = 1288;
		assertTrue(fix1.equivalent(fix2,false));
		assertFalse(fix1.equivalent(fix2,true));
	}

	@Test
	public void testEqualsWrongType() {
		Fix fix1 = new Fix(new Date(), 45.020, 108.000, 0, 0, 'V');
		assertFalse(fix1.equals(new Flight()));
	}

	@Test
	public void testEquals() {
		Date d = new Date();
		Fix fix1 = new Fix(d, -45.666, 108.345, 267, 288, 'V');

		Fix fix2 = new Fix(d, -45.666, 108.345, 267, 288, 'V');
		assertTrue(fix1.equals(fix2));
		assertTrue(fix2.equals(fix1));

		Calendar cal = Calendar.getInstance();
		cal.setTime(d);
		cal.add(Calendar.DATE, 1);
		Date nd = cal.getTime();
		Fix fix3 = new Fix(nd, -45.666, 108.345, 267, 288, 'V');
		assertFalse(fix1.equals(fix3));

		Fix fix4 = new Fix(d, 45.666, 108.345, 267, 288, 'V');
		assertFalse(fix1.equals(fix4));

		Fix fix5 = new Fix(d, -45.666, -108.345, 267, 288, 'V');
		assertFalse(fix1.equals(fix5));

		Fix fix6 = new Fix(d, -45.666, 108.345, 1267, 288, 'V');
		assertFalse(fix1.equals(fix6));

		Fix fix7 = new Fix(d, -45.666, 108.345, 267, 1288, 'V');
		assertFalse(fix1.equals(fix7));

		Fix fix8 = new Fix(d, -45.666, 108.345, 267, 288, 'A');
		assertFalse(fix1.equals(fix8));
	}

	@Test
	public void testClone() {
		Fix fix1 = new Fix(new Date(), -45.666, 108.345, 267, 288, 'V');
		Fix fix2 = fix1.clone();

		assertFalse(fix1 == fix2);
		assertTrue(fix1.equals(fix2));
		assertTrue(fix2.equals(fix1));
	}

	@Test
	public void testSetLat() {
		Fix fix1 = new Fix(new Date(), -45.666, 108.345, 267, 288, 'V');
		double lat = 66.788;
		fix1.setLat(lat);
		assertEquals(lat, fix1.lat(), 0.0);
		assertEquals(Math.toRadians(lat), fix1.latrd(), 0.0);
	}

	@Test
	public void testSetLon() {
		Fix fix1 = new Fix(new Date(), -45.666, 108.345, 267, 288, 'V');
		double lon = 88.344;
		fix1.setLon(lon);
		assertEquals(lon, fix1.lon(), 0.0);
		assertEquals(Math.toRadians(lon), fix1.lonrd(), 0.0);
	}

}
