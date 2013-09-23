package com.ezgliding.igc;

import org.junit.Before;
import org.junit.Test;
import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertNull;
import static org.junit.Assert.assertNotNull;
import static org.junit.Assert.assertTrue;

import java.util.ArrayList;
import java.util.Date;

public class RectangleSetTest {

	private ArrayList<Fix> fixes1, fixesOverlap1, fixesNoOverlap1;

	@Before
	public void setUpClass() {
		fixes1 = new ArrayList<Fix>();
		fixes1.add(new Fix(0, 45.888, 108.999, 0, 0, 'V'));
		fixes1.add(new Fix(0, 44.223, 109.112, 0, 0, 'V'));
		fixes1.add(new Fix(0, 43.123, 109.998, 0, 0, 'V'));
		fixesOverlap1 = new ArrayList<Fix>();
		fixesOverlap1.add(new Fix(0, 44.888, 109.555, 0, 0, 'V'));
		fixesOverlap1.add(new Fix(0, 38.227, 110.234, 0, 0, 'V'));
		fixesOverlap1.add(new Fix(0, 39.132, 111.255, 0, 0, 'V'));
		fixesNoOverlap1 = new ArrayList<Fix>();
		fixesNoOverlap1.add(new Fix(0, 44.888, 110.999, 0, 0, 'V'));
		fixesNoOverlap1.add(new Fix(0, 44.223, 112.112, 0, 0, 'V'));
		fixesNoOverlap1.add(new Fix(0, 43.123, 113.998, 0, 0, 'V'));
	}

	@Test
	public void testCreation() {
		RectangleSet set = new RectangleSet(fixes1);
		assertNotNull(set.getVertices());
		Fix nw = new Fix(0, 45.888, 108.999, 0, 0, 'V');
		Fix sw = new Fix(0, 43.123, 108.999, 0, 0, 'V');
		Fix ne = new Fix(0, 45.888, 109.998, 0, 0, 'V');
		Fix se = new Fix(0, 43.123, 109.998, 0, 0, 'V');
		assertTrue(set.nw().equivalent(nw, false));
		assertTrue(set.sw().equivalent(sw, false));
		assertTrue(set.ne().equivalent(ne, false));
		assertTrue(set.se().equivalent(se, false));
	}

	@Test
	public void testCreationNull() {
		RectangleSet set = new RectangleSet(null);
		assertNull(set.getVertices());
	}

	@Test
	public void testOverlap() {
		RectangleSet set1 = new RectangleSet(fixes1);
		RectangleSet set2 = new RectangleSet(fixesOverlap1);
		RectangleSet set3 = new RectangleSet(fixesNoOverlap1);

		assertTrue(set1.overlap(set1));
		assertTrue(set1.overlap(set2));
		assertTrue(set2.overlap(set1));
		assertFalse(set1.overlap(set3));
	}

	@Test
	public void testOverlapNull() {
		RectangleSet set1 = new RectangleSet(fixes1);
		assertFalse(set1.overlap(null));
	}

	@Test
	public void testContains() {
		RectangleSet set1 = new RectangleSet(fixes1);
		assertTrue(set1.contains(new Fix(0, 43.765, 109.231, 0, 0, 'V')));
	}

	@Test
	public void testNotContains() {
		RectangleSet set1 = new RectangleSet(fixes1);
		assertFalse(set1.contains(new Fix(0, 47.765, 109.231, 0, 0, 'V')));
		assertFalse(set1.contains(new Fix(0, 43.765, 113.231, 0, 0, 'V')));
	}


	@Test
	public void testSplit() { 
		ArrayList<Fix> all = new ArrayList<Fix>();
		all.add(new Fix(0, 40.300, 106.999, 0, 0, 'V'));
		all.add(new Fix(0, 42.700, 108.008, 0, 0, 'V'));
		all.add(new Fix(0, 43.500, 110.999, 0, 0, 'V'));
		all.add(new Fix(0, 44.234, 111.877, 0, 0, 'V'));
		RectangleSet allSet = new RectangleSet(all);

		RectangleSet firstHalfSet = new RectangleSet(all, 0, 2);
		RectangleSet secondHalfSet = new RectangleSet(all, 2, 4);

		RectangleSet[] result = allSet.split();
		assertNotNull(result);
		assertEquals(2, result.length);
		assertEquals(firstHalfSet, result[0]);
		assertEquals(secondHalfSet, result[1]);
	}

	@Test
	public void testSplitTwo() { 
		ArrayList<Fix> all = new ArrayList<Fix>();
		all.add(new Fix(0, 40.300, 106.999, 0, 0, 'V'));
		all.add(new Fix(0, 44.234, 111.877, 0, 0, 'V'));
		RectangleSet allSet = new RectangleSet(all);

		RectangleSet firstHalfSet = new RectangleSet(all, 0, 1);
		RectangleSet secondHalfSet = new RectangleSet(all, 1, 2);

		RectangleSet[] result = allSet.split();
		assertNotNull(result);
		assertEquals(2, result.length);
		assertNotNull(result[0]);
		assertNotNull(result[1]);
		assertEquals(1, firstHalfSet.getFixes().size());
		assertEquals(1, secondHalfSet.getFixes().size());
		assertEquals(1, result[0].getFixes().size());
		assertEquals(1, result[1].getFixes().size());
		assertEquals(firstHalfSet, result[0]);
		assertEquals(secondHalfSet, result[1]);
	}

	@Test
	public void testSplitSingle() { //TODO: what should we expect here?
		ArrayList<Fix> all = new ArrayList<Fix>();
		all.add(new Fix(0, 40.300, 106.999, 0, 0, 'V'));
		RectangleSet allSet = new RectangleSet(all);

		RectangleSet[] result = allSet.split();
		assertNotNull(result);
	}

	@Test
	public void testDiagonal() {
		RectangleSet set1 = new RectangleSet(fixes1);

		Fix v1 = new Fix(0, 45.888, 108.999, 0, 0, 'V');
		Fix v2 = new Fix(0, 43.123, 109.998, 0, 0, 'V');
		
		double expected = Util.distance(v1, v2);
		assertEquals(expected, set1.diagonal(), 0.0);
	}

	@Test
	public void testMaxDistance() {
		ArrayList<Fix> fixes1 = new ArrayList<Fix>();
		fixes1.add(new Fix(0, 47.900, 106.700, 0, 0, 'V'));
		fixes1.add(new Fix(0, 45.900, 108.700, 0, 0, 'V'));
		RectangleSet set1 = new RectangleSet(fixes1);

		ArrayList<Fix> fixes2 = new ArrayList<Fix>();
		fixes2.add(new Fix(0, 43.900, 110.700, 0, 0, 'V'));
		fixes2.add(new Fix(0, 41.900, 112.700, 0, 0, 'V'));
		RectangleSet set2 = new RectangleSet(fixes2);

		double expected = Util.distance(set1.nw(), set2.se());
		assertEquals(expected, set1.maxDistance(set2), 0.0);
		assertEquals(expected, set2.maxDistance(set1), 0.0);
	}

	@Test
	public void testMinDistance() { 
		ArrayList<Fix> fixes1 = new ArrayList<Fix>();
		fixes1.add(new Fix(0, 47.900, 106.700, 0, 0, 'V'));
		fixes1.add(new Fix(0, 45.900, 108.700, 0, 0, 'V'));
		RectangleSet set1 = new RectangleSet(fixes1);

		ArrayList<Fix> fixes2 = new ArrayList<Fix>();
		fixes2.add(new Fix(0, 43.900, 110.700, 0, 0, 'V'));
		fixes2.add(new Fix(0, 41.900, 112.700, 0, 0, 'V'));
		RectangleSet set2 = new RectangleSet(fixes2);

		double expected = Util.distance(set1.se(), set2.nw());
		assertEquals(expected, set1.minDistance(set2), 0.0);
	}

	@Test
	public void testMinDistanceOverlap() { 
		ArrayList<Fix> fixes1 = new ArrayList<Fix>();
		fixes1.add(new Fix(0, 47.900, 106.700, 0, 0, 'V'));
		fixes1.add(new Fix(0, 45.900, 108.700, 0, 0, 'V'));
		RectangleSet set1 = new RectangleSet(fixes1);

		ArrayList<Fix> fixes2 = new ArrayList<Fix>();
		fixes2.add(new Fix(0, 46.900, 107.700, 0, 0, 'V'));
		fixes2.add(new Fix(0, 48.900, 112.700, 0, 0, 'V'));
		RectangleSet set2 = new RectangleSet(fixes2);

		assertEquals(0.0, set1.minDistance(set2), 0.0);
	}

	@Test
	public void testStart() {
		//TODO:
	}

	@Test
	public void testEnd() {
		//TODO:
	}

	@Test
	public void testNumFixes() {
		//TODO:
	}

	@Test
	public void testGetVertices() {
		RectangleSet set1 = new RectangleSet(fixes1);
		assertNotNull(set1.getVertices());
		assertEquals(4, set1.getVertices().length);
		for (Fix v: set1.getVertices())
			assertNotNull(v);
	}

	@Test
	public void testCompareTo() {
		RectangleSet set1 = new RectangleSet(fixes1, 0, 1);
		RectangleSet set2 = new RectangleSet(fixes1, 1, 3);
		assertEquals(-1, set1.compareTo(set2));
		assertEquals(1, set2.compareTo(set1));

		RectangleSet set3 = new RectangleSet(fixes1, 1, 3);
		assertEquals(0, set2.compareTo(set3));
		assertEquals(0, set3.compareTo(set2));
	}

	@Test
	public void testNotEquals() {
		RectangleSet set1 = new RectangleSet(fixes1);
		ArrayList<Fix> fixes = new ArrayList<Fix>();
		fixes.add(new Fix(0, 45.888, 108.999, 0, 0, 'V'));
		RectangleSet set2 = new RectangleSet(fixes);
		assertFalse(set1.equals(set2));
	}

	@Test
	public void testEquals() {
		RectangleSet set1 = new RectangleSet(fixes1);
		RectangleSet set2 = new RectangleSet(fixes1);
		assertTrue(set1.equals(set2));
	}
}
