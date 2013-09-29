package com.ezgliding.crawl;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.IOException;
import java.net.URL;
import java.net.MalformedURLException;
import java.util.ArrayList;
import java.util.Calendar;
import java.util.TimeZone;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import com.ezgliding.igc.Task;
import com.ezgliding.igc.WayPoint;

public class CrawlerNetcoupe extends Crawler {

	private static Logger logger = Logger.getLogger(CrawlerNetcoupe.class.getName());

	private static String FLIGHT_DETAIL_SUBURL = "Results/FlightDetail.aspx?FlightID=";

	private static Calendar cal = Calendar.getInstance(TimeZone.getTimeZone("UTC"));

	static { cal.set(Calendar.MILLISECOND, 0); }

	private static String FLIGHT_DETAIL_REGEX = ".*<b>Nom&nbsp;:</b>.*4113.*>(.*)</a>"
		+ ".*<b>Cat.+gorie&nbsp;:</b></div></td><td><div align=\"left\">(.*)</div>"
		+ ".*<b>Club&nbsp;\\(r.+gion, pays\\) :</b>.*43.*>(.*)</a>&nbsp;\\((.*), (.*)\\)</div>"
		+ ".*<b>Date&nbsp;:</b></div></td><td><div align=\"left\">(.*)</div>"
		+ ".*<b>A.+rodrome de d.+part&nbsp;:</b></div></td><td><div align=\"left\">(.*)</div>"
		+ ".*<b>R.+gion&nbsp;:</b></div></td><td><div align=\"left\">(.*)</div>"
		+ ".*<b>Pays&nbsp;:</b></div></td><td><div align=\"left\">(.*)</div>"
		+ ".*<b>Distance&nbsp;:</b></div></td><td><div align=\"left\">(.*)&nbsp;kms</div>"
		+ ".*<b>Points&nbsp;:</b></div></td><td><div align=\"left\">(.*)</div>"
		+ ".*<b>Planeur.*\"middle\">(.*)&nbsp;&nbsp;</td>"
		+ ".*<b>Type de circuit&nbsp;:</b></div></td><td><div align=\"left\">(.*)</div>"
		+ ".*<b>Fichier \\.IGC&nbsp;:.*align=\"left\"><a href=\"(.*)\">T"
		+ ".*<b>Vitesse moyenne du circuit&nbsp;:</b></div></td><td><div align=\"left\">(.*)&nbsp;km/h</div>"
		+ ".*<b>Point de d.+part&nbsp;:</b></div></td><td><div align=\"left\">(.*)</div>"
		+ ".*<b>Point de virage n.+1&nbsp;:</b></div></td><td><div align=\"left\">(.*)</div>"
		+ ".*<b>Point de virage n.+2&nbsp;:</b></div></td><td><div align=\"left\">(.*)</div>"
		+ ".*<b>Point de virage n.+3&nbsp;:</b></div></td><td><div align=\"left\">(.*)</div>"
		+ ".*<b>Point d'arriv.+e&nbsp;:</b></div></td><td><div align=\"left\">(.*)</div>"
		+ ".*<b>Commentaires&nbsp;:</b></div></td><td><div align=\"left\">(.*)</div></td></tr></tbody>"
		+ ".*Fermer.*";

	private String baseUrl;

	private Pattern regex;

	public CrawlerNetcoupe(String baseUrl) {
		if (baseUrl == null) throw new IllegalArgumentException("A baseUrl must be provided");
		this.baseUrl = baseUrl;
		regex = Pattern.compile(FLIGHT_DETAIL_REGEX, Pattern.MULTILINE | Pattern.DOTALL);
	}

	@Override
	public int getLastId() { return 0; }

	@Override
	public void setLastId(int lastId) { }

	@Override
	public FlightEntry getFlight(int id) {
		return getFlight(FLIGHT_DETAIL_SUBURL + id);
	}

	protected FlightEntry getFlight(String urlStr) {
		if (urlStr == null) return null;

		StringBuffer buff = new StringBuffer();
		try {
			URL url = new URL(urlStr);
			BufferedReader reader = new BufferedReader(new InputStreamReader(url.openStream()));
			String line;
			while ((line = reader.readLine()) != null) {
			    buff.append(line.trim());
			}
			reader.close();
		} catch (MalformedURLException e) {
			logger.log(Level.SEVERE, "A malformed url was made: [{0}] : [{1}]", new Object[] { urlStr, e });  
			return null;
		} catch (IOException e) {
			logger.log(Level.WARNING, "Failed to read url contents: [{0}] : [{1}]", new Object[] { urlStr, e });
			return null;
		}

		return parseFlightEntry(buff.toString());
	}

	protected FlightEntry parseFlightEntry(String data) {
		if (data == null) return null;

		FlightEntry flight = new FlightEntry();

		Matcher matcher = regex.matcher(data);
		matcher.matches();
		flight.setPilot(matcher.group(1));
		flight.setCategory(matcher.group(2));
		flight.setClub(matcher.group(3));
		flight.setRegion(matcher.group(4));
		flight.setCountry(matcher.group(5));
		String dateStr = matcher.group(6);
		cal.set(Integer.parseInt(dateStr.substring(6)),
			Integer.parseInt(dateStr.substring(3,5))-1,
			Integer.parseInt(dateStr.substring(0,2)), 0, 0, 0);
		flight.setDate(cal.getTime());
		flight.setAirfield(matcher.group(7));
		flight.setRegion(matcher.group(8));
		flight.setCountry(matcher.group(9));
		flight.setDistance(Float.parseFloat(matcher.group(10).replace(",",".")));
		flight.setPoints(Float.parseFloat(matcher.group(11).replace(",",".")));
		flight.setGliderType(matcher.group(12));
		FlightEntry.CircuitType cType = FlightEntry.CircuitType.FREE;
		String type = matcher.group(13);
		if (type.equals("Libre")) cType = FlightEntry.CircuitType.FREE;
		flight.setCircuitType(cType);
		flight.setFileLocation(matcher.group(14));
		flight.setSpeed(Float.parseFloat(matcher.group(15).replace(",",".")));
		Task task = new Task();
		task.setStart(new WayPoint(null, matcher.group(16)));
		task.addTurnPoint(new WayPoint(null, matcher.group(17)));
		task.addTurnPoint(new WayPoint(null, matcher.group(18)));
		task.addTurnPoint(new WayPoint(null, matcher.group(19)));
		task.setFinish(new WayPoint(null, matcher.group(20)));
		flight.setTask(task);
		flight.setComment(matcher.group(21));

		return flight;
	}

}
