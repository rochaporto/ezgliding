package com.ezgliding.igc;

import java.nio.file.FileSystems;
import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;

public class BrokenLineOptimizer extends Optimizer {

	private int numPoints;

	private TreeMap<Candidate,Double> maxTree;

	public BrokenLineOptimizer(Flight flight, int numPoints) {
		super(flight);

		if (numPoints <= 0) 
			throw new IllegalArgumentException("numPoints should be 1 or greater");
		
		this.numPoints = numPoints;
		maxTree = new TreeMap<Candidate,Double>();
	}

	public Result optimize() {
		if (flight == null || flight.fixes() == null) return null;

		// We start with a candidate containing only one rectangle (with all points)
		ArrayList<RectangleSet> initialSet = new ArrayList<RectangleSet>();
		initialSet.add(new RectangleSet(flight.fixes()));
		Candidate first = new Candidate(initialSet);
		maxTree.put(first, first.max());

		// From here we trigger the logic of fetching the max from tree
		Map.Entry<Candidate,Double> maxEntry = null;
		while (maxTree.size() != 0) {
			maxEntry = maxTree.lastEntry();
			bound(maxEntry.getKey());
			maxTree.remove(maxEntry.getKey()); //TODO: remove this
		}
	
		return null;	
	}

	private void bound(Candidate candate) {

	}

	private Iterator<Candidate> branch(Candidate candate) {
		return null;
	}

	protected List<Candidate> permutations(List<RectangleSet> availableSets) {
		ArrayList<Candidate> finalCandidates = new ArrayList<Candidate>();

		permutations(availableSets, numPoints, 0, new Candidate(), finalCandidates);
		return finalCandidates; 
	}

	private void permutations(List<RectangleSet> availableSets, int size, int from, 
		Candidate current, List<Candidate> candidates) {
		
		if (current.getRectangles().size() == size) { // Final condition, add and return
			candidates.add(current);
			return;
		}

		for (int i=from; i<availableSets.size(); i++) {
			Candidate newCurrent = current.clone();
			newCurrent.add(availableSets.get(i));
			permutations(availableSets, size, i, newCurrent, candidates);
		}
	}

	public Flight getFlight() {
		return flight;
	}

	public int getNumPoints() {
		return numPoints;
	}
}
