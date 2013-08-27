package com.ezgliding.igc;

import org.junit.Before;
import org.junit.Test;
import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNotNull;

import java.io.IOException;
import java.nio.file.FileSystems;
import java.text.ParseException;

import com.ezgliding.igc.Parser;

public class ParserTest {

	@Before
	public void setUp() {

	}

	@Test(expected=IOException.class)
	public void testPathNotExists() throws IOException {
		Parser parser = new Parser();
		try {
			Flight flight = parser.parse(
				FileSystems.getDefault().getPath("file.does.not.exist.at.all"));
		} catch(ParseException e) { }
	}	

	@Test
	public void testParseSampleFromPath() throws IOException, ParseException {
		Parser parser = new Parser();
		Flight flight = parser.parse(
			FileSystems.getDefault().getPath("test/com/ezgliding/igc/SampleFlight.igc"));
		assertNotNull(flight);	
		assertNotNull(flight.fixes());
		assertEquals(2426, flight.fixes().size());
	}

}
