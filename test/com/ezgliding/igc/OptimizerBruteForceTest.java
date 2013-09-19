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
import java.util.List;

public class OptimizerBruteForceTest {

	@Before
	public void setUp() {
	}

	@Test
	public void testOptimize1TP3Points() throws IOException, ParseException {
		Parser parser = new Parser();
		Flight flight = parser.parse(
			FileSystems.getDefault().getPath("test/data/optimize-with-3-points.igc"));
		OptimizerBruteForce opt = new OptimizerBruteForce(flight, 3);
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
		OptimizerBruteForce opt = new OptimizerBruteForce(flight, 3);
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
		OptimizerBruteForce opt = new OptimizerBruteForce(flight, 4);
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
		OptimizerBruteForce opt = new OptimizerBruteForce(flight, 5);
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
		OptimizerBruteForce opt = new OptimizerBruteForce(flight, 3);
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
	public void testOptimize2TP10Points() throws IOException, ParseException {
		Parser parser = new Parser();
		Flight flight = parser.parse(
			FileSystems.getDefault().getPath("test/data/optimize-with-10-points.igc"));
		OptimizerBruteForce opt = new OptimizerBruteForce(flight, 4);
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
		OptimizerBruteForce opt = new OptimizerBruteForce(flight, 5);
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

	private int getTime(int hour, int min, int second) {
		return (hour*3600)+(min*60)+second;
	}

}
