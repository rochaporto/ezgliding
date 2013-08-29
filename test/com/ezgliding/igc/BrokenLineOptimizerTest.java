package com.ezgliding.igc;

import org.junit.Before;
import org.junit.Test;
import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNotNull;
import static org.junit.Assert.assertTrue;

import java.io.IOException;
import java.nio.file.FileSystems;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.logging.Logger;
import java.text.ParseException;

public class BrokenLineOptimizerTest {

	private static Logger logger = Logger.getLogger(BrokenLineOptimizerTest.class.getName());

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
	public void testOptimize() throws IOException, ParseException {
		Parser parser = new Parser();
		Flight flight = parser.parse(
			FileSystems.getDefault().getPath("test/com/ezgliding/igc/SampleFlight.igc"));
		BrokenLineOptimizer opt = new BrokenLineOptimizer(flight, 5);
		Optimizer.Result result = opt.optimize();
		assertNotNull(result);
		assertEquals(111, result.distance(), 0.0);
		logger.finest("" + result);
	}

	@Test
	public void testBranch() {
		ArrayList<RectangleSet> sets = new ArrayList<RectangleSet>();

		ArrayList<Fix> fixes1 = new ArrayList<Fix>();
		fixes1.add(new Fix(new Date(), 45.888, 108.999, 0, 0, 'V'));
		fixes1.add(new Fix(new Date(), 44.223, 109.112, 0, 0, 'V'));
		fixes1.add(new Fix(new Date(), 43.123, 110.998, 0, 0, 'V'));
		fixes1.add(new Fix(new Date(), 42.077, 111.877, 0, 0, 'V'));
		sets.add(new RectangleSet(fixes1));

		ArrayList<Fix> fixes2 = new ArrayList<Fix>();
		fixes2.add(new Fix(new Date(), 42.888, 111.999, 0, 0, 'V'));
		fixes2.add(new Fix(new Date(), 41.223, 112.112, 0, 0, 'V'));
		fixes2.add(new Fix(new Date(), 41.123, 113.998, 0, 0, 'V'));
		sets.add(new RectangleSet(fixes2));

		Candidate candate = new Candidate(sets);

		ArrayList<Fix> expected1F = new ArrayList<Fix>();
		expected1F.add(fixes1.get(0));
		expected1F.add(fixes1.get(1));
		RectangleSet expected1 = new RectangleSet(expected1F);
		ArrayList<Fix> expected2F = new ArrayList<Fix>();
		expected2F.add(fixes1.get(2));
		expected2F.add(fixes1.get(3));
		RectangleSet expected2 = new RectangleSet(expected2F);
	
		BrokenLineOptimizer opt = new BrokenLineOptimizer(new Flight(), 2);
		List<Candidate> result = opt.branch(candate);
		assertEquals(2, result.size()); // As 2 is == sets.size(), we get 2 candidates
		assertEquals(result.get(0).getRectangles().get(0), expected1);
		assertEquals(result.get(1).getRectangles().get(0), expected2);
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

	@Test
	public void testGetFlight() {
		Flight flight = new Flight();
		BrokenLineOptimizer opt = new BrokenLineOptimizer(flight, 4);
		assertEquals(flight, opt.getFlight());
	}

	@Test
	public void testGetNumPoints() {
		BrokenLineOptimizer opt = new BrokenLineOptimizer(new Flight(), 4);
		assertEquals(4, opt.getNumPoints());
	}

}
