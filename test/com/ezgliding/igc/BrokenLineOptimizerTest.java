package com.ezgliding.igc;

import org.junit.Before;
import org.junit.Test;
import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNotNull;
import static org.junit.Assert.assertTrue;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;

public class BrokenLineOptimizerTest {

	ArrayList<Fix> fixes;

	@Before
	public void setUp() {
		fixes = new ArrayList<Fix>();
		fixes.add(new Fix(new Date(), 45.888, 108.999, 0, 0, 'V'));
		fixes.add(new Fix(new Date(), 44.223, 109.112, 0, 0, 'V'));
		fixes.add(new Fix(new Date(), 43.123, 109.998, 0, 0, 'V'));
	}

	@Test
	public void testCreation() {
		Flight flight = new Flight();
		BrokenLineOptimizer opt = new BrokenLineOptimizer(flight, 5);
		assertEquals(5, opt.getNumPoints());
		assertEquals(flight, opt.getFlight());
	}

	@Test(expected=IllegalArgumentException.class)
	public void testCreationNull() {
		BrokenLineOptimizer opt = new BrokenLineOptimizer(null, 5);
	}

	@Test(expected=IllegalArgumentException.class)
	public void testCreationNegativePoints() {
		BrokenLineOptimizer opt = new BrokenLineOptimizer(new Flight(), -1);
	}

	@Test
	public void testOptimize() {
		assertTrue(false);
	}

	@Test
	public void testBound() {
		assertTrue(false);
	}

	@Test
	public void testBranch() {
		assertTrue(false);
	}

	@Test
	public void testPermutations() {
		ArrayList<RectangleSet> availableSets = new ArrayList<RectangleSet>();
		for (int i=0; i<3; i++)
			availableSets.add(new RectangleSet(fixes));

		BrokenLineOptimizer opt = new BrokenLineOptimizer(new Flight(), 4);
		List<Candidate> result = opt.permutations(availableSets);
		assertNotNull(result);
		assertEquals(15, result.size());
	}

}
