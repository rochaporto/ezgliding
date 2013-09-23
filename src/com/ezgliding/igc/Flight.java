package com.ezgliding.igc;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;

public class Flight {

	private String manufacturer;

	private String uniqueId;

	private String additionalData;

	private Date date;

	private int fixAccuracy;

	private String pilot;

	private String crew2;

	private String gliderType;

	private String gliderId;

	private String gpsDatum;

	private String firmwareVersion;

	private String hardwareVersion;

	private String frType;

	private String gpsManufacturer;

	private String pressAltSensor;

	private String competitionId;

	private String competitionClass; 

	private ArrayList<Fix> fixes;

	private Task task;

	public Flight() {
		this.fixes = new ArrayList<Fix>();
	}

	public List<Fix> fixes() {
		return fixes;
	}

	public void addFix(Fix fix) {
		fixes.add(fix);
	}

	public void setManufacturer(String manufacturer) { this.manufacturer = manufacturer; }

	public String getManufacturer() { return manufacturer; }

	public void setUniqueId(String uniqueId) { this.uniqueId = uniqueId; }

	public String getUniqueId() { return uniqueId; }

	public void setAdditionalData(String additionalData) { this.additionalData = additionalData; }

	public String getAdditionalData() { return additionalData; }

	public void setDate(Date date) { this.date = date; }

	public Date getDate() { return date; }

	public void setFixAccuracy(int fixAccuracy) { this.fixAccuracy = fixAccuracy; }

	public int getFixAccuracy() { return fixAccuracy; }

	public void setPilot(String pilot) { this.pilot = pilot; }

	public String getPilot() { return pilot; }

	public void setCrew2(String crew2) { this.crew2 = crew2; }

	public String getCrew2() { return crew2; }

	public void setGliderType(String gliderType) { this.gliderType = gliderType; }

	public String getGliderType() { return gliderType; }

	public void setGliderId(String gliderId) { this.gliderId = gliderId; }

	public String getGliderId() { return gliderId; }

	public void setGpsDatum(String gpsDatum) { this.gpsDatum = gpsDatum; }

	public String getGpsDatum() { return gpsDatum; }

	public void setFirmwareVersion(String firmwareVersion) { this.firmwareVersion = firmwareVersion; } 

	public String getFirmwareVersion() { return firmwareVersion; }

	public void setHardwareVersion(String hardwareVersion) { this.hardwareVersion = hardwareVersion; }

	public String getHardwareVersion() { return hardwareVersion; }

	public void setFrType(String frType) { this.frType = frType; }

	public String getFrType() { return frType; }

	public void setGpsManufacturer(String gpsManufacturer) { this.gpsManufacturer = gpsManufacturer; }

	public String getGpsManufacturer() { return gpsManufacturer; }

	public void setPressAltSensor(String pressAltSensor) { this.pressAltSensor = pressAltSensor; }

	public String getPressAltSensor() { return pressAltSensor; }

	public void setCompetitionId(String competitionId) { this.competitionId = competitionId; }

	public String getCompetitionId() { return competitionId; }

	public void setCompetitionClass(String competitionClass) { this.competitionClass = competitionClass; }

	public String getCompetitionClass() { return competitionClass; }

	public void setTask(Task task) { this.task = task; }

	public Task getTask() { return task; }

	@Override
	public String toString() {
		return "Fixes: " + fixes.size();
	}
}
