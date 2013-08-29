package com.ezgliding.igc;

public abstract class Optimizer {

	protected Flight flight;

	public Optimizer(Flight flight) {
		if (flight == null) throw new IllegalArgumentException("Flight cannot be null");
		this.flight = flight;
	}

	public abstract Result optimize();

	public class Result {
		
		private double distance;
		
		public Fix[] points;

		public Result(Fix[] points) {
			this.points = points;
			for (int i=0; i<points.length-1; i++)
				distance += Util.distance(points[i], points[i+1]);
		}

		public double distance() {
			return distance;
		}

		public boolean equals(Result other) {
			if (points.length != other.points.length) return false;
			for (int i=0; i<points.length; i++)
				if (!points[i].equals(other.points[i])) return false;
			return true;
		}

		public String toString() {
			String str = "{" + String.format("%1$,.2f", distance) + ",";
			for (Fix point: points)
				str += point + ",";
			return str + "}";
		}
	}
}
