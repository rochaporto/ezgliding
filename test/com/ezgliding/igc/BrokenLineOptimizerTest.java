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

	ArrayList<Fix> fixes, fixes2;

	private static Calendar calendar;

	@BeforeClass
	public static void setUpClass() {
		calendar = Calendar.getInstance();
	}

	@Before
	public void setUp() {
		fixes = new ArrayList<Fix>();
		fixes.add(new Fix(1000, 41.000, 101.000, 0, 0, 'V'));
		fixes.add(new Fix(2000, 42.000, 102.000, 0, 0, 'V'));
		fixes.add(new Fix(3000, 43.000, 103.000, 0, 0, 'V'));

		fixes2 = new ArrayList<Fix>();
		fixes2.add(new Fix(4000, 44.000, 104.000, 0, 0, 'V'));
		fixes2.add(new Fix(5000, 45.000, 105.000, 0, 0, 'V'));
		fixes2.add(new Fix(6000, 48.000, 108.000, 0, 0, 'V'));
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
	public void testOptimize1TP3Points() throws IOException, ParseException {
		Parser parser = new Parser();
		Flight flight = parser.parse(
			FileSystems.getDefault().getPath("test/data/optimize-with-3-points.igc"));
		BrokenLineOptimizer opt = new BrokenLineOptimizer(flight, 3);
		Optimizer.Result result = opt.optimize();
		assertNotNull(result);
		assertEquals(845.129, result.distance(), 0.001);
		Fix[] points = new Fix[] {
			new Fix(getTime(9,00,00), Util.minDec2decimal("4500000N"), 
				Util.minDec2decimal("00500000E"), 0, 0, 'A'),
			new Fix(getTime(9,20,00), Util.minDec2decimal("4700000N"), 
				Util.minDec2decimal("00700000E"), 0, 0, 'A'),
			new Fix(getTime(9,40,00), Util.minDec2decimal("5200000N"), 
				Util.minDec2decimal("00900000E"), 0, 0, 'A'),
		};
		assertEquals(points.length, result.points.length);
		for (int i=0; i<points.length; i++)
			assertEquals(points[i], result.points[i]);
	}

	@Test
	public void testOptimize1TP5Points() throws IOException, ParseException {
		Parser parser = new Parser();
		Flight flight = parser.parse(
			FileSystems.getDefault().getPath("test/data/optimize-with-5-points.igc"));
		BrokenLineOptimizer opt = new BrokenLineOptimizer(flight, 3);
		Optimizer.Result result = opt.optimize();
		assertNotNull(result);
		assertEquals(855.378, result.distance(), 0.001);
		Fix[] points = new Fix[] {
			new Fix(getTime(9,00,00), Util.minDec2decimal("4500000N"), 
				Util.minDec2decimal("00500000E"), 0, 0, 'A'),
			new Fix(getTime(9,30,00), Util.minDec2decimal("4800000N"), 
				Util.minDec2decimal("00800000E"), 0, 0, 'A'),
			new Fix(getTime(9,40,00), Util.minDec2decimal("5200000N"), 
				Util.minDec2decimal("00900000E"), 0, 0, 'A'),
		};
		assertEquals(points.length, result.points.length);
		for (int i=0; i<points.length; i++)
			assertEquals(points[i], result.points[i]);
	}

	@Test
	public void testOptimize2TP5Points() throws IOException, ParseException {
		Parser parser = new Parser();
		Flight flight = parser.parse(
			FileSystems.getDefault().getPath("test/data/optimize-with-5-points.igc"));
		BrokenLineOptimizer opt = new BrokenLineOptimizer(flight, 4);
		Optimizer.Result result = opt.optimize();
		assertNotNull(result);
		assertEquals(855.424, result.distance(), 0.001);
		Fix[] points = new Fix[] {
			new Fix(getTime(9,00,00), Util.minDec2decimal("4500000N"), 
				Util.minDec2decimal("00500000E"), 0, 0, 'A'),
			new Fix(getTime(9,20,00), Util.minDec2decimal("4700000N"), 
				Util.minDec2decimal("00700000E"), 0, 0, 'A'),
			new Fix(getTime(9,30,00), Util.minDec2decimal("4800000N"), 
				Util.minDec2decimal("00800000E"), 0, 0, 'A'),
			new Fix(getTime(9,40,00), Util.minDec2decimal("5200000N"), 
				Util.minDec2decimal("00900000E"), 0, 0, 'A'),
		};
		assertEquals(points.length, result.points.length);
		for (int i=0; i<points.length; i++)
			assertEquals(points[i], result.points[i]);
	}

	@Test
	public void testOptimize3TP5Points() throws IOException, ParseException {
		Parser parser = new Parser();
		Flight flight = parser.parse(
			FileSystems.getDefault().getPath("test/data/optimize-with-5-points.igc"));
		BrokenLineOptimizer opt = new BrokenLineOptimizer(flight, 5);
		Optimizer.Result result = opt.optimize();
		assertNotNull(result);
		assertEquals(855.439, result.distance(), 0.001);
		Fix[] points = new Fix[] {
			new Fix(getTime(9,00,00), Util.minDec2decimal("4500000N"), 
				Util.minDec2decimal("00500000E"), 0, 0, 'A'),
			new Fix(getTime(9,10,00), Util.minDec2decimal("4600000N"), 
				Util.minDec2decimal("00600000E"), 0, 0, 'A'),
			new Fix(getTime(9,20,00), Util.minDec2decimal("4700000N"), 
				Util.minDec2decimal("00700000E"), 0, 0, 'A'),
			new Fix(getTime(9,30,00), Util.minDec2decimal("4800000N"), 
				Util.minDec2decimal("00800000E"), 0, 0, 'A'),
			new Fix(getTime(9,40,00), Util.minDec2decimal("5200000N"), 
				Util.minDec2decimal("00900000E"), 0, 0, 'A'),
		};
		assertEquals(points.length, result.points.length);
		for (int i=0; i<points.length; i++)
			assertEquals(points[i], result.points[i]);
	}

	@Test
	public void testOptimize1TP10Points() throws IOException, ParseException {
		Parser parser = new Parser();
		Flight flight = parser.parse(
			FileSystems.getDefault().getPath("test/data/optimize-with-10-points.igc"));
		BrokenLineOptimizer opt = new BrokenLineOptimizer(flight, 3);
		Optimizer.Result result = opt.optimize();
		assertNotNull(result);
		assertEquals(315.589, result.distance(), 0.001);
		Fix[] points = new Fix[] {
			new Fix(getTime(9,00,00), Util.minDec2decimal("4600000N"), 
				Util.minDec2decimal("00600000E"), 0, 0, 'A'),
			new Fix(getTime(9,50,00), Util.minDec2decimal("4650000N"), 
				Util.minDec2decimal("00650000E"), 0, 0, 'A'),
			new Fix(getTime(10,40,00), Util.minDec2decimal("4505000N"), 
				Util.minDec2decimal("00605000E"), 0, 0, 'A'),
		};
		assertEquals(points.length, result.points.length);
		for (int i=0; i<points.length; i++)
			assertEquals(points[i], result.points[i]);
	}

	@Test
	public void testOptimize2TP10101010101010101010Points() throws IOException, ParseException {
		Parser parser = new Parser();
		Flight flight = parser.parse(
			FileSystems.getDefault().getPath("test/data/optimize-with-10-points.igc"));
		BrokenLineOptimizer opt = new BrokenLineOptimizer(flight, 4);
		Optimizer.Result result = opt.optimize();
		assertNotNull(result);
		assertEquals(417.724, result.distance(), 0.001);
		Fix[] points = new Fix[] {
			new Fix(getTime(9,00,00), Util.minDec2decimal("4600000N"), 
				Util.minDec2decimal("00600000E"), 0, 0, 'A'),
			new Fix(getTime(9,50,00), Util.minDec2decimal("4650000N"), 
				Util.minDec2decimal("00650000E"), 0, 0, 'A'),
			new Fix(getTime(10,40,00), Util.minDec2decimal("4505000N"), 
				Util.minDec2decimal("00605000E"), 0, 0, 'A'),
			new Fix(getTime(10,50,00), Util.minDec2decimal("4600000N"), 
				Util.minDec2decimal("00600000E"), 0, 0, 'A'),
		};
		assertEquals(points.length, result.points.length);
		for (int i=0; i<points.length; i++)
			assertEquals(points[i], result.points[i]);
	}

	@Test
	public void testOptimize3TP10Points() throws IOException, ParseException {
		Parser parser = new Parser();
		Flight flight = parser.parse(
			FileSystems.getDefault().getPath("test/data/optimize-with-10-points.igc"));
		BrokenLineOptimizer opt = new BrokenLineOptimizer(flight, 5);
		Optimizer.Result result = opt.optimize();
		assertNotNull(result);
		assertEquals(425.883, result.distance(), 0.001);
		Fix[] points = new Fix[] {
			new Fix(getTime(9,00,00), Util.minDec2decimal("4600000N"), 
				Util.minDec2decimal("00600000E"), 0, 0, 'A'),
			new Fix(getTime(9,50,00), Util.minDec2decimal("4650000N"), 
				Util.minDec2decimal("00650000E"), 0, 0, 'A'),
			new Fix(getTime(10,00,00), Util.minDec2decimal("4545000N"), 
				Util.minDec2decimal("00645000E"), 0, 0, 'A'),
			new Fix(getTime(10,40,00), Util.minDec2decimal("4505000N"), 
				Util.minDec2decimal("00605000E"), 0, 0, 'A'),
			new Fix(getTime(10,50,00), Util.minDec2decimal("4600000N"), 
				Util.minDec2decimal("00600000E"), 0, 0, 'A'),
		};
		assertEquals(points.length, result.points.length);
		for (int i=0; i<points.length; i++)
			assertEquals(points[i], result.points[i]);
	}
	@Test(expected=IllegalArgumentException.class)
	public void testBranchNull() {
		BrokenLineOptimizer opt = new BrokenLineOptimizer(new Flight(), 2);
		List<Candidate> result = opt.branch(null);
	}

	@Test(expected=IllegalArgumentException.class)
	public void testBranch0Rects() {
		BrokenLineOptimizer opt = new BrokenLineOptimizer(new Flight(), 2);
		List<Candidate> result = opt.branch(new Candidate());
	}

	@Test
	public void testBranch() {
		ArrayList<RectangleSet> sets = new ArrayList<RectangleSet>();
		sets.add(new RectangleSet(fixes));
		sets.add(new RectangleSet(fixes2)); // fixes2 has larger diagonal

		Candidate candate = new Candidate(sets);
		BrokenLineOptimizer opt = new BrokenLineOptimizer(new Flight(), 2);
		List<Candidate> result = opt.branch(candate);
		assertEquals(6, result.size()); 
		// TODO: Check each element
	}

	@Test
	public void testPermutations2Sets2Points() {
		ArrayList<RectangleSet> tmp;

		ArrayList<RectangleSet> availableSets = new ArrayList<RectangleSet>();
		availableSets.add(new RectangleSet(fixes));
		availableSets.add(new RectangleSet(fixes2));

		BrokenLineOptimizer opt = new BrokenLineOptimizer(new Flight(), 2);
		List<Candidate> result = opt.permutations(availableSets);
		assertNotNull(result);
		assertEquals(3, result.size());

		tmp = new ArrayList<RectangleSet>();
		tmp.add(availableSets.get(0));
		tmp.add(availableSets.get(0));
		assertEquals(tmp.size(), result.get(0).getRectangles().size());
		assertEquals(tmp, result.get(0).getRectangles());
			
		tmp = new ArrayList<RectangleSet>();
		tmp.add(availableSets.get(0));
		tmp.add(availableSets.get(1));
		assertEquals(tmp.size(), result.get(1).getRectangles().size());
		assertEquals(tmp, result.get(1).getRectangles());

		tmp = new ArrayList<RectangleSet>();
		tmp.add(availableSets.get(1));
		tmp.add(availableSets.get(1));
		assertEquals(tmp.size(), result.get(2).getRectangles().size());
		assertEquals(tmp, result.get(2).getRectangles());
	}

	@Test
	public void testPermutations2Sets3Points() {
		ArrayList<RectangleSet> tmp;

		ArrayList<RectangleSet> availableSets = new ArrayList<RectangleSet>();
		availableSets.add(new RectangleSet(fixes,0,1));
		availableSets.add(new RectangleSet(fixes,1,3));

		BrokenLineOptimizer opt = new BrokenLineOptimizer(new Flight(), 3);
		List<Candidate> result = opt.permutations(availableSets);
		assertNotNull(result);
		assertEquals(4, result.size());

		tmp = new ArrayList<RectangleSet>();
		tmp.add(availableSets.get(0));
		tmp.add(availableSets.get(0));
		tmp.add(availableSets.get(0));
		assertEquals(tmp.size(), result.get(0).getRectangles().size());
		assertEquals(tmp, result.get(0).getRectangles());
			
		tmp = new ArrayList<RectangleSet>();
		tmp.add(availableSets.get(0));
		tmp.add(availableSets.get(0));
		tmp.add(availableSets.get(1));
		assertEquals(tmp.size(), result.get(1).getRectangles().size());
		assertEquals(tmp, result.get(1).getRectangles());

		tmp = new ArrayList<RectangleSet>();
		tmp.add(availableSets.get(0));
		tmp.add(availableSets.get(1));
		tmp.add(availableSets.get(1));
		assertEquals(tmp.size(), result.get(2).getRectangles().size());
		assertEquals(tmp, result.get(2).getRectangles());

		tmp = new ArrayList<RectangleSet>();
		tmp.add(availableSets.get(1));
		tmp.add(availableSets.get(1));
		tmp.add(availableSets.get(1));
		assertEquals(tmp.size(), result.get(3).getRectangles().size());
		assertEquals(tmp, result.get(3).getRectangles());
	}
	@Test

	public void testPermutations3Sets3Points() {
		ArrayList<RectangleSet> tmp;

		ArrayList<RectangleSet> availableSets = new ArrayList<RectangleSet>();
		availableSets.add(new RectangleSet(fixes,0,1));
		availableSets.add(new RectangleSet(fixes,1,2));
		availableSets.add(new RectangleSet(fixes,2,3));

		BrokenLineOptimizer opt = new BrokenLineOptimizer(new Flight(), 3);
		List<Candidate> result = opt.permutations(availableSets);
		assertNotNull(result);
		assertEquals(10, result.size());
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
