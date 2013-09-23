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

public class OptimizerBruteForceTestPerf {

	@Before
	public void setUp() {
	}

	@Test
	public void testOptimize1TPSampleFlight() throws IOException, ParseException {
		Parser parser = new Parser();
		Flight flight = parser.parse(
			FileSystems.getDefault().getPath("test/data/SampleFlight.igc"));
		OptimizerBruteForce opt = new OptimizerBruteForce(flight, 3);
		Optimizer.Result result = opt.optimize();
		assertNotNull(result);
		assertEquals(492.646, result.distance(), 0.001);
	}

	@Test
	public void testOptimize2TPSampleFlight() throws IOException, ParseException {
		Parser parser = new Parser();
		Flight flight = parser.parse(
			FileSystems.getDefault().getPath("test/data/SampleFlight.igc"));
		OptimizerBruteForce opt = new OptimizerBruteForce(flight, 4);
		Optimizer.Result result = opt.optimize();
		assertNotNull(result);
		assertEquals(646.424, result.distance(), 0.001);
	}

	@Test
	public void testOptimize3TPSampleFlight() throws IOException, ParseException {
		Parser parser = new Parser();
		Flight flight = parser.parse(
			FileSystems.getDefault().getPath("test/data/SampleFlight.igc"));
		OptimizerBruteForce opt = new OptimizerBruteForce(flight, 5);
		Optimizer.Result result = opt.optimize();
		assertNotNull(result);
		assertEquals(742.052, result.distance(), 0.001);
	}

	private int getTime(int hour, int min, int second) {
		return (hour*3600)+(min*60)+second;
	}

}
