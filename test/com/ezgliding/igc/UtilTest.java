package com.ezgliding.igc;

import java.util.HashMap;
import java.util.Map.Entry;

import org.junit.Before;
import org.junit.Test;
import static org.junit.Assert.assertEquals;

import com.ezgliding.igc.Util;

public class UtilTest {

	@Before
	public void setUp() {

	}

	@Test(expected=IllegalArgumentException.class)	
	public void testMinDecNullThrowsIllegalArgument() {
		double d = Util.minDec2decimal(null);
	}

	@Test
	public void testLatMinDec2Decimal() {
		HashMap<String,Double> examples = new HashMap<String,Double>();
		examples.put("4026771N", 40.44618);
		examples.put("3522750S", -35.37917);

		for (Entry<String,Double> ex: examples.entrySet())
			assertEquals(ex.getValue(), Util.minDec2decimal(ex.getKey()), 0.00001);
	}

	@Test
	public void testLonMinDec2Decimal() {
		HashMap<String,Double> examples = new HashMap<String,Double>();
		examples.put("07956931W", -79.94885);
		examples.put("10837900E", 108.63167);

		for (Entry<String,Double> ex: examples.entrySet())
			assertEquals(ex.getValue(), Util.minDec2decimal(ex.getKey()), 0.00001);
	}

	@Test
	public void testDistance() {
		HashMap<Double,Fix[]> examples = new HashMap<Double,Fix[]>();
		examples.put(968.85348, new Fix[] {
			new Fix(null, Util.minDec2decimal("5003983N"), Util.minDec2decimal("00542883W"), 0, 0, 'V'),
			new Fix(null, Util.minDec2decimal("5838633N"), Util.minDec2decimal("00304200W"), 0, 0, 'V') });

		for (Entry<Double,Fix[]> ex: examples.entrySet()) {
			assertEquals(ex.getKey(), Util.distance(ex.getValue()[0], ex.getValue()[1]), 0.00001);
			assertEquals(ex.getKey(), Util.distance(ex.getValue()[1], ex.getValue()[0]), 0.00001);
		}
	}

}
