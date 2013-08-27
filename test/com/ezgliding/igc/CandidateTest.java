package com.ezgliding.igc;

import org.junit.Before;
import org.junit.Test;
import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertNotNull;
import static org.junit.Assert.assertTrue;

import java.util.ArrayList;
import java.util.Date;


public class CandidateTest {

	private RectangleSet set;

	@Before
	public void setUp() {
		ArrayList<Fix> fixes = new ArrayList<Fix>();
		fixes.add(new Fix(new Date(), 45.888, 108.999, 0, 0, 'V'));
		fixes.add(new Fix(new Date(), 44.223, 109.112, 0, 0, 'V'));
		fixes.add(new Fix(new Date(), 43.123, 109.998, 0, 0, 'V'));

		set = new RectangleSet(fixes);
	}

	@Test
	public void testCreationEmpty() {
		Candidate candate = new Candidate();
		assertEquals(0, candate.getRectangles().size());
	}

	@Test
	public void testCreationNull() {
		Candidate candate = new Candidate(null);
	}

	@Test
	public void testCreation() {
		ArrayList<RectangleSet> rects = new ArrayList<RectangleSet>();
		rects.add(set);
		Candidate candate = new Candidate(rects);
		assertEquals(1, candate.getRectangles().size());
		assertEquals(set, candate.getRectangles().get(0));
	}

	@Test
	public void testMax() {
		assertTrue(false);
	}

	@Test
	public void testMin() {
		assertTrue(false);
	}

	@Test
	public void testAdd() {
		Candidate candate = new Candidate();
		candate.add(set);
		assertEquals(1, candate.getRectangles().size());

		ArrayList<Fix> fixes = new ArrayList<Fix>();
		fixes.add(new Fix(new Date(), 42.111, 107.333, 0, 0, 'V'));
		fixes.add(new Fix(new Date(), 44.411, 103.333, 0, 0, 'A'));
		RectangleSet newSet = new RectangleSet(fixes);
		candate.add(newSet);
		
		assertEquals(2, candate.getRectangles().size());
		assertEquals(newSet, candate.getRectangles().get(1));
	}

	@Test
	public void testAddNull() {
		Candidate candate = new Candidate();
		candate.add(set);
		assertEquals(1, candate.getRectangles().size());

		candate.add(null);
		assertEquals(1, candate.getRectangles().size());
	}

	@Test
	public void testGetRectangles() {
		Candidate candate = new Candidate();
		candate.add(set);
		assertNotNull(candate.getRectangles());
		assertEquals(1, candate.getRectangles().size());
		assertEquals(set, candate.getRectangles().get(0));
	}

	@Test
	public void testClone() {
		Candidate initial = new Candidate();
		initial.add(set);

		Candidate clone = initial.clone();
		assertFalse(initial == clone);
		assertTrue(initial.equals(clone));
		assertEquals(initial.getRectangles().size(), clone.getRectangles().size());
	}
}
