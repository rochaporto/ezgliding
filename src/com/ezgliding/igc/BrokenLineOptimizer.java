package com.ezgliding.igc;

import java.nio.file.FileSystems;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;

public class BrokenLineOptimizer extends Optimizer {

	private int numPoints;

	private TreeMap<Double,Candidate> maxTree;

	public BrokenLineOptimizer(Flight flight, int numPoints) {
		super(flight);

		this.numPoints = numPoints;
		maxTree = new TreeMap<Double,Candidate>();
	}

	public Result optimize() {
		if (flight == null || flight.fixes() == null) return null;

		RectangleSet singleSet = new RectangleSet(flight.fixes());
		RectangleSet[] initialSet = new RectangleSet[numPoints];
		for (int i=0; i<initialSet.length; i++)
			initialSet[i] = singleSet;

		Candidate first = new Candidate(initialSet);
		maxTree.put(first.max(), first);

		Map.Entry<Double,Candidate> maxEntry = null;
		while (maxTree.size() != 0) {
			maxEntry = maxTree.lastEntry();
			evaluate(maxEntry.getValue());
			maxTree.remove(maxEntry.getKey()); //TODO: remove this
		}
		
		RectangleSet[] finalSet = maxEntry.getValue().getRectangles();
		Fix[] points = new Fix[finalSet.length];
		for (int i=0; i<finalSet.length; i++)
			points[i] = finalSet[i].fixes.get(0);
		return new Result(points);
	}

	private void evaluate(Candidate candate) {
		
	}

	public static void main(String[] args) {
		Flight flight = null;
		Parser parser = new Parser();
		try {
			flight = parser.parse(FileSystems.getDefault().getPath("sample.igc"));
		} catch(Exception e) { e.printStackTrace(); }
		System.out.println(flight);
		double distance = 0.0;
		List<Fix> fixes = flight.fixes();
		for (int i=1; i<fixes.size(); i++)
			distance += Util.distance(fixes.get(i-1), fixes.get(i));
		System.out.println("distance = " + distance);
		System.out.println(new BrokenLineOptimizer(flight, 5).optimize());
	}
}
