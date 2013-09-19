package com.ezgliding.igc;

import java.util.ArrayList;
import java.util.List;

public class Flight {

	private String manufacturer;

	private String uniqueID;

	private String additionalData;

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

	public void setManufacturer(String manuf) { manufacturer = manuf; }

	public String getManufacturer() { return manufacturer; }

	public void setUniqueID(String uniID) { uniqueID = uniID; }

	public String getUniqueID() { return uniqueID; }

	public void setAdditionalData(String addData) { additionalData = addData; }

	public String getAdditionalData() { return additionalData; }

	@Override
	public String toString() {
		return "Fixes: " + fixes.size();
	}
}
