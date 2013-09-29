package com.ezgliding.crawl;

import java.util.Iterator;

public abstract class Crawler implements Iterable {

	public FlightEntry getFlight() { return getFlight(getLastId()); }

	public FlightEntry getFlight(int id) { return null; }

	public int getLastId() { return 0; }

	public void setLastId(int lastId) { }

	@Override
	public Iterator<FlightEntry> iterator() {
		return new FlightEntryIterator();
	}

	public class FlightEntryIterator implements Iterator {

		public FlightEntryIterator() {

		}

		@Override
		public FlightEntry next() { 
			return getFlight();
		}

		@Override
		public boolean hasNext() { 
			if (getLastId() != -1) return true;
			return false;
 		}

		@Override
		public void remove() { }
	}
}
