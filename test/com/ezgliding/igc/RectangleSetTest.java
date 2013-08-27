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
		fixes1.add(new Fix(new Date(), 45.888, 108.999, 0, 0, 'V'));
		fixes1.add(new Fix(new Date(), 44.223, 109.112, 0, 0, 'V'));
		fixes1.add(new Fix(new Date(), 43.123, 109.998, 0, 0, 'V'));
		fixesOverlap1 = new ArrayList<Fix>();
		fixesOverlap1.add(new Fix(new Date(), 44.888, 109.555, 0, 0, 'V'));
		fixesOverlap1.add(new Fix(new Date(), 38.227, 110.234, 0, 0, 'V'));
		fixesOverlap1.add(new Fix(new Date(), 39.132, 111.255, 0, 0, 'V'));
		fixesNoOverlap1 = new ArrayList<Fix>();
		fixesNoOverlap1.add(new Fix(new Date(), 44.888, 110.999, 0, 0, 'V'));
		fixesNoOverlap1.add(new Fix(new Date(), 44.223, 112.112, 0, 0, 'V'));
		fixesNoOverlap1.add(new Fix(new Date(), 43.123, 113.998, 0, 0, 'V'));
	}

	@Test
	public void testCreation() {
		RectangleSet set = new RectangleSet(fixes1);
		Fix[] manualVertices = new Fix[] { 
			new Fix(null, 45.888, 108.999, 0, 0, 'V'),
			new Fix(null, 43.123, 108.999, 0, 0, 'V'),
			new Fix(null, 45.888, 109.998, 0, 0, 'V'),
			new Fix(null, 43.123, 109.998, 0, 0, 'V')
		};

		Fix[] vertices = set.getVertices();
		assertNotNull(vertices);
		for (int i=0; i<vertices.length; i++) {
			assertTrue("Unexpected vertice (" + i + ")\n" 
				+ vertices[i] + "\n" + manualVertices[i],
				vertices[i].equivalent(manualVertices[i], false));
		}
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
		assertTrue(set1.contains(new Fix(new Date(), 43.765, 109.231, 0, 0, 'V')));
	}

	@Test
	public void testNotContains() {
		RectangleSet set1 = new RectangleSet(fixes1);
		assertFalse(set1.contains(new Fix(new Date(), 47.765, 109.231, 0, 0, 'V')));
		assertFalse(set1.contains(new Fix(new Date(), 43.765, 113.231, 0, 0, 'V')));
	}


	@Test
	public void testSplit() { 
		assertTrue(false);
	}

	@Test
	public void testDiagonal() {
		RectangleSet set1 = new RectangleSet(fixes1);

		Fix v1 = new Fix(new Date(), 45.888, 108.999, 0, 0, 'V');
		Fix v2 = new Fix(new Date(), 43.123, 109.998, 0, 0, 'V');
		
		double expected = Util.distance(v1, v2);
		assertEquals(expected, set1.diagonal(), 0.0);
	}

	@Test
	public void testMaxDistance() {
		Fix v11 = new Fix(new Date(), 47.900, 106.700, 0, 0, 'V');
		Fix v12 = new Fix(new Date(), 45.900, 108.700, 0, 0, 'V');
		ArrayList<Fix> fixes1 = new ArrayList<Fix>();
		fixes1.add(v11);
		fixes1.add(v12);
		RectangleSet set1 = new RectangleSet(fixes1);

		Fix v21 = new Fix(new Date(), 43.900, 110.700, 0, 0, 'V');
		Fix v22 = new Fix(new Date(), 41.900, 112.700, 0, 0, 'V');
		ArrayList<Fix> fixes2 = new ArrayList<Fix>();
		fixes2.add(v21);
		fixes2.add(v22);
		RectangleSet set2 = new RectangleSet(fixes2);

		double expected = Util.distance(set1.nw, set2.se);
		assertEquals(expected, set1.maxDistance(set2), 0.0);
		assertEquals(expected, set2.maxDistance(set1), 0.0);
	}

	@Test
	public void testMinDistance() { 
		Fix v11 = new Fix(new Date(), 47.900, 106.700, 0, 0, 'V');
		Fix v12 = new Fix(new Date(), 45.900, 108.700, 0, 0, 'V');
		ArrayList<Fix> fixes1 = new ArrayList<Fix>();
		fixes1.add(v11);
		fixes1.add(v12);
		RectangleSet set1 = new RectangleSet(fixes1);

		Fix v21 = new Fix(new Date(), 43.900, 110.700, 0, 0, 'V');
		Fix v22 = new Fix(new Date(), 41.900, 112.700, 0, 0, 'V');
		ArrayList<Fix> fixes2 = new ArrayList<Fix>();
		fixes2.add(v21);
		fixes2.add(v22);
		RectangleSet set2 = new RectangleSet(fixes2);

		double expected = Util.distance(set1.se, set2.nw);
		assertEquals(expected, set1.minDistance(set2), 0.0);
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
	public void testEquals() {
		RectangleSet set1 = new RectangleSet(fixes1);
		RectangleSet set2 = new RectangleSet(fixes1);
		assertTrue(set1.equals(set2));
	}
}
