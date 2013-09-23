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

	private ArrayList<Fix> fixes;

	private RectangleSet set1, set2;

	@Before
	public void setUp() {
		fixes = new ArrayList<Fix>();
		fixes.add(new Fix(1000, 41.000, 101.000, 100, 110, 'V'));
		fixes.add(new Fix(2000, 42.000, 102.000, 200, 210, 'V'));
		fixes.add(new Fix(3000, 43.000, 103.000, 300, 310, 'V'));
		fixes.add(new Fix(4000, 44.000, 104.000, 400, 410, 'V'));
		fixes.add(new Fix(5000, 45.000, 105.000, 500, 510, 'V'));
		fixes.add(new Fix(6000, 48.000, 108.000, 800, 810, 'V'));
		set1 = new RectangleSet(fixes, 0, 3);
		set2 = new RectangleSet(fixes, 3, 6);
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
		ArrayList<RectangleSet> sets = new ArrayList<RectangleSet>();
		sets.add(set1);
		Candidate candate = new Candidate(sets);
		assertEquals(1, candate.getRectangles().size());
		assertEquals(set1, candate.getRectangles().get(0));
	}

	@Test
	public void testIsFinal() {
		ArrayList<RectangleSet> sets1 = new ArrayList<RectangleSet>();
		ArrayList<RectangleSet> sets2 = new ArrayList<RectangleSet>();

		sets1.add(new RectangleSet(fixes,0,1));
		sets2.add(new RectangleSet(fixes,0,1));

		sets1.add(new RectangleSet(fixes,1,2));
		sets2.add(new RectangleSet(fixes,1,2));

		sets1.add(new RectangleSet(fixes,2,4));
		sets2.add(new RectangleSet(fixes,3,4));

		Candidate candate1 = new Candidate(sets1);
		Candidate candate2 = new Candidate(sets2);

		assertEquals(3, candate1.getRectangles().size());
		assertEquals(3, candate2.getRectangles().size());
		assertFalse(candate1.isFinal());
		assertTrue(candate2.isFinal());
	}
	
	@Test
	public void testMax() {
		ArrayList<RectangleSet> sets = new ArrayList<RectangleSet>();

		sets.add(new RectangleSet(fixes,0,2));
		sets.add(new RectangleSet(fixes,2,4));

		Candidate candate = new Candidate(sets);
		double expected = Util.distance(sets.get(0).sw(), sets.get(1).ne());
		assertEquals(expected, candate.max(), 0.0);
	}
	@Test
	public void testMax3Sets() {
		ArrayList<RectangleSet> sets = new ArrayList<RectangleSet>();

		sets.add(new RectangleSet(fixes,0,2));
		sets.add(new RectangleSet(fixes,2,4));
		sets.add(new RectangleSet(fixes,4,6));

		Candidate candate = new Candidate(sets);
		double expected = Util.distance(sets.get(0).sw(), sets.get(1).se()) 
			+ Util.distance(sets.get(1).se(), sets.get(2).ne());
		assertEquals(expected, candate.max(), 0.0);
	}

	@Test
	public void testMin() {
		ArrayList<RectangleSet> sets = new ArrayList<RectangleSet>();

		sets.add(new RectangleSet(fixes,0,2));
		sets.add(new RectangleSet(fixes,2,4));
		sets.add(new RectangleSet(fixes,4,6));

		Candidate candate = new Candidate(sets);
		double expected = Util.distance(sets.get(0).ne(), sets.get(1).sw()) 
			+ Util.distance(sets.get(1).ne(), sets.get(2).sw());
		assertEquals(expected, candate.min(), 0.0);
	}

	@Test
	public void testAdd() {
		Candidate candate = new Candidate();
		candate.add(set1);
		assertEquals(1, candate.getRectangles().size());

		candate.add(set2);
		assertEquals(2, candate.getRectangles().size());
		assertEquals(set2, candate.getRectangles().get(1));
	}

	@Test(expected=IllegalArgumentException.class)
	public void testAddNull() {
		Candidate candate = new Candidate();
		candate.add(null);
	}

	@Test
	public void testGetRectangles() {
		Candidate candate = new Candidate();
		candate.add(set1);
		assertNotNull(candate.getRectangles());
		assertEquals(1, candate.getRectangles().size());
		assertEquals(set1, candate.getRectangles().get(0));
	}

	@Test
	public void testReplaceSingle() {
		ArrayList<RectangleSet> sets = new ArrayList<RectangleSet>();
		sets.add(set1);
		sets.add(set2);
		Candidate candate = new Candidate(sets);

		RectangleSet replacement = new RectangleSet(fixes, 5, 6);
		
		assertEquals(2, candate.getRectangles().size());
		assertFalse(replacement.equals(candate.getRectangles().get(1)));
		candate.replace(sets.get(1), replacement);
		assertEquals(2, candate.getRectangles().size());
		assertEquals(replacement, candate.getRectangles().get(1));
	}

	@Test
	public void testReplaceMultiple() {
		ArrayList<RectangleSet> sets = new ArrayList<RectangleSet>();
		sets.add(set1);
		sets.add(set2);
	
		Candidate candate = new Candidate(sets);

		RectangleSet replacement1 = new RectangleSet(fixes, 3, 4);
		RectangleSet replacement2 = new RectangleSet(fixes, 4, 5);
		
		assertEquals(2, candate.getRectangles().size());
		assertFalse(replacement1.equals(candate.getRectangles().get(1)));
		assertFalse(replacement2.equals(candate.getRectangles().get(1)));
		candate.replace(sets.get(1), new RectangleSet[] { replacement1, replacement2 });
		assertEquals(3, candate.getRectangles().size());
		assertEquals(replacement1, candate.getRectangles().get(1));
		assertEquals(replacement2, candate.getRectangles().get(2));
	}

	@Test
	public void testLargestDiagonal() {
		ArrayList<RectangleSet> sets = new ArrayList<RectangleSet>();

		sets.add(set1);
		sets.add(set2);

		Candidate candate = new Candidate(sets);
		RectangleSet largerDiagonal = candate.largestDiagonal();
		assertEquals(sets.get(1), largerDiagonal);
	}

	@Test
	public void testClone() {
		Candidate initial = new Candidate();
		initial.add(set1);

		Candidate clone = initial.clone();
		assertFalse(initial == clone);
		assertTrue(initial.equals(clone));
		assertEquals(initial.getRectangles().size(), clone.getRectangles().size());
	}

	@Test
	public void testEquals() {
		Candidate c1 = new Candidate();
		c1.add(set1);
		Candidate c2 = new Candidate();
		c2.add(set1);
		assertTrue(c1.equals(c2));
	}

	@Test
	public void testNotEquals() {
		Candidate c1 = new Candidate();
		c1.add(set1);
		Candidate c2 = new Candidate();
		c2.add(set2);
		assertFalse(c1.equals(c2));
	}
}
