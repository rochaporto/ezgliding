package com.ezgliding.igc;

import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;
import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNotNull;
import static org.junit.Assert.assertTrue;

import java.io.IOException;
import java.nio.file.FileSystems;
import java.sql.Time;
import java.text.ParseException;
import java.util.ArrayList;
import java.util.Calendar;
import java.util.Date;
import java.util.List;
import java.util.logging.Logger;

public class BrokenLineOptimizerTest {

	private static Logger logger = Logger.getLogger(BrokenLineOptimizerTest.class.getName());

	ArrayList<Fix> fixes;

	private static Calendar calendar;

	@BeforeClass
	public static void setUpClass() {
		calendar = Calendar.getInstance();
	}

	@Before
	public void setUp() {
		fixes = new ArrayList<Fix>();
		fixes.add(new Fix(0, 45.888, 108.999, 0, 0, 'V'));
		fixes.add(new Fix(0, 44.223, 109.112, 0, 0, 'V'));
		fixes.add(new Fix(0, 43.123, 109.998, 0, 0, 'V'));
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
	public void testOptimize5TP5Points() throws IOException, ParseException {
		Parser parser = new Parser();
		Flight flight = parser.parse(
			FileSystems.getDefault().getPath("test/data/optimize-with-5-points-only.igc"));
		BrokenLineOptimizer opt = new BrokenLineOptimizer(flight, 5);
		Optimizer.Result result = opt.optimize();
		assertNotNull(result);
		assertEquals(494.785, result.distance(), 0.001);
		Fix[] points = new Fix[] {
			new Fix(getTime(17,21,15), Util.minDec2decimal("4533504N"), 
				Util.minDec2decimal("00558638E"), 0, 0, 'A'),
			new Fix(getTime(16,39,44), Util.minDec2decimal("4547266N"), 
				Util.minDec2decimal("00644677E"), 0, 0, 'A'),
			new Fix(getTime(15,05,57), Util.minDec2decimal("4451913N"), 
				Util.minDec2decimal("00641047E"), 0, 0, 'A'),
			new Fix(getTime(12,52,56), Util.minDec2decimal("4621959N"), 
				Util.minDec2decimal("00730485E"), 0, 0, 'A'),
			new Fix(getTime(9,31,18), Util.minDec2decimal("4533432N"), 
				Util.minDec2decimal("00558760E"), 0, 0, 'A'),
		};
		for (int i=0; i<points.length; i++)
			assertEquals(points[1], result.points[1]);
	}

	@Test
	public void testBranch() {
		ArrayList<RectangleSet> sets = new ArrayList<RectangleSet>();

		ArrayList<Fix> fixes1 = new ArrayList<Fix>();
		fixes1.add(new Fix(0, 45.888, 108.999, 0, 0, 'V'));
		fixes1.add(new Fix(0, 44.223, 109.112, 0, 0, 'V'));
		fixes1.add(new Fix(0, 43.123, 110.998, 0, 0, 'V'));
		fixes1.add(new Fix(0, 42.077, 111.877, 0, 0, 'V'));
		sets.add(new RectangleSet(fixes1));

		ArrayList<Fix> fixes2 = new ArrayList<Fix>();
		fixes2.add(new Fix(0, 42.888, 111.999, 0, 0, 'V'));
		fixes2.add(new Fix(0, 41.223, 112.112, 0, 0, 'V'));
		fixes2.add(new Fix(0, 41.123, 113.998, 0, 0, 'V'));
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

	private int getTime(int hour, int min, int second) {
		return (hour*3600)+(min*60)+second;
	}

}
