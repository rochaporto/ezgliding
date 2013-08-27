package com.ezgliding.igc;

import java.util.ArrayList;
import java.util.List;

public class Flight {

	private ArrayList<Fix> fixes;

	public Flight() {
		this.fixes = new ArrayList<Fix>();
	}

	public List<Fix> fixes() {
		return fixes;
	}

	public void addFix(Fix fix) {
		fixes.add(fix);
	}

	public String toString() {
		return "Fixes: " + fixes.size();
	}
}
