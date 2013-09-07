package com.ezgliding.igc;

import java.nio.file.FileSystems;
import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.SortedMap;
import java.util.TreeMap;
import java.util.logging.Logger;

public class BrokenLineOptimizer extends Optimizer {

	private static Logger logger = Logger.getLogger(BrokenLineOptimizer.class.getName());

	private TreeMap<Double,Candidate> maxTree;

	private Candidate current;

	private Iterator<Candidate> candIter;

	public BrokenLineOptimizer(Flight flight, int numPoints) {
		super(flight, numPoints);

		reset();
	}

	public void reset() {

		maxTree = new TreeMap<Double,Candidate>();

		// We start with a candidate containing only one rectangle (with all points)
		ArrayList<RectangleSet> initialSet = new ArrayList<RectangleSet>();
		initialSet.add(new RectangleSet(flight.fixes()));
		Candidate first = new Candidate(initialSet);
		maxTree.put(first.max(), first);

		candIter = iterator();
	}

	@Override
	public Result optimize() {
		if (flight == null || flight.fixes() == null) return null;

		Candidate result = null, current;

		while (candIter.hasNext()) {
			// We get the max entry and remove it from the tree
			current = candIter.next();

			// If final and better than current max, update result
			if (current.isFinal() && (result == null || current.max() > result.max()))
				result = current;

		}
		
		ArrayList<Fix> points = new ArrayList<Fix>();
		for (RectangleSet set: result.getRectangles())
			points.add(set.getFixes().get(0));
		return new Result(points.toArray(new Fix[] {}));
	}

	protected List<Candidate> branch(Candidate candate) {
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

	protected void prune(double min) {
		Set<Double> pruneKeys = maxTree.headMap(min).keySet();
		Double[] prune = pruneKeys.toArray(new Double[] {});
		pruneKeys = null;
		for (Double d: prune)
			maxTree.remove(d);
	}

	public Iterator<Candidate> iterator() {
		return new CandidateIterator();
	}

	public class CandidateIterator implements Iterator {

		private Map.Entry<Double,Candidate> entry;

		private Candidate current;

		public CandidateIterator() {

		}

		public boolean hasNext() { 
			if (maxTree != null && maxTree.size() != 0)
				return true; 
			return false;
		}

		public Candidate next() { 
			if (!hasNext()) return null;
	
			// Get maximum and remove it from tree
			entry = maxTree.lastEntry();
			maxTree.remove(entry.getKey());
			current = entry.getValue();

			// If not final, branch and add to treemap
			if (!current.isFinal()) {
				List<Candidate> branchCandidates = branch(current);
				for (Candidate candate: branchCandidates)
					maxTree.put(candate.max(), candate);
			}

			// Prune the tree
			prune(current.min());

			return entry.getValue(); 
		}

		public void remove() { }
	}

	public Flight getFlight() {
		return flight;
	}

	public int getNumPoints() {
		return numPoints;
	}
}
