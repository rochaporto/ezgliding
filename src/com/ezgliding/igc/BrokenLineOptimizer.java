package com.ezgliding.igc;

import java.nio.file.FileSystems;
import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;

public class BrokenLineOptimizer extends Optimizer {

	private int numPoints;

	private TreeMap<Double,Candidate> maxTree;

	public BrokenLineOptimizer(Flight flight, int numPoints) {
		super(flight);

		if (numPoints <= 0) 
			throw new IllegalArgumentException("numPoints should be 1 or greater");
		
		this.numPoints = numPoints;
		maxTree = new TreeMap<Double,Candidate>();
	}

	@Override
	public Result optimize() {
		if (flight == null || flight.fixes() == null) return null;

		Candidate result = null;

		// We start with a candidate containing only one rectangle (with all points)
		ArrayList<RectangleSet> initialSet = new ArrayList<RectangleSet>();
		initialSet.add(new RectangleSet(flight.fixes()));
		Candidate first = new Candidate(initialSet);
		maxTree.put(first.max(), first);

		// From here we start the branch / bound procedure 
		Map.Entry<Double,Candidate> maxEntry = null;
		Candidate current;
		while (maxTree.size() != 0) {
			maxEntry = maxTree.lastEntry();
			current = maxEntry.getValue();
			maxTree.remove(maxEntry.getKey());
			// If final and better than current max, update result
			if (current.isFinal() && current.max() > result.max())
				result = current;
			else // If not final, branch and add to treemap
				for (Candidate candate: branch(current))
					maxTree.put(candate.max(), candate);

			// Prune the tree (remove keys < current.min())
			for (Double d: maxTree.headMap(current.min()).keySet())
				maxTree.remove(d);
		}
	
		ArrayList<Fix> points = new ArrayList<Fix>();
		for (RectangleSet set: result.getRectangles())
			points.add(set.getFixes().get(0));
		return new Result(points.toArray(new Fix[] {}));
	}

	private List<Candidate> branch(Candidate candate) {
		if (candate.getRectangles().size() <= 0) 
			throw new IllegalArgumentException("Cannot branch empty candidate");

		ArrayList<Candidate> result = new ArrayList<Candidate>();

		// Get the index of the rectangle with the largest diagonal
		int largerDiagonal = candate.largestDiagonal();

		RectangleSet[] newSets = candate.getRectangles().get(largerDiagonal).split();
		// If we already had enough rectangles, then create new candidate per rect 
		if (candate.getRectangles().size() == numPoints) {
			for (RectangleSet set: newSets) {
				Candidate newCandidate = candate.clone();
				newCandidate.replace(largerDiagonal, set);
				result.add(newCandidate);
			}
		} else { // Else replace with both new rects in a single candidate
			Candidate newCandidate = candate.clone();
			newCandidate.replace(largerDiagonal, newSets);
			result.add(newCandidate);
		}

		return result;
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
