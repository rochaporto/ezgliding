package com.ezgliding.igc;

import java.awt.BasicStroke;
import java.awt.Color;
import java.awt.Dimension;
import java.awt.Graphics;
import java.awt.Graphics2D;
import java.awt.Insets;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.nio.file.FileSystems;
import java.util.Iterator;
import java.util.List;
import java.util.Random;
import javax.swing.AbstractAction;
import javax.swing.KeyStroke;
import javax.swing.JFrame;
import javax.swing.JPanel;
import javax.swing.SwingUtilities;

public class OptimizerBrokenLineAnim extends JFrame {

	public class CandidatePanel extends JPanel {

		private void draw(Graphics g) {
			Graphics2D g2d = (Graphics2D)g;

			g2d.setColor(Color.blue);
			g2d.setStroke(new BasicStroke(1));

			Dimension size = getSize();

			int[] xy;
			List<Fix> fixes = opt.getFlight().fixes();
			for (int i=0; i<fixes.size(); i++) {
				xy = latLon2XY(fixes.get(i), size.width-50, size.height-50);
				g2d.fillOval(xy[0]+20, xy[1]+20, 5, 5);
				g2d.drawString(""+i, xy[0]+20, xy[1]+20);
			}

			if (current != null) {
				int[] xyS, xyE;
				List<RectangleSet> sets = current.getRectangles();
				for (int i=0; i<sets.size(); i++) {
					xyS = latLon2XY(sets.get(i).nw(), size.width-50, size.height-50);
					xyE = latLon2XY(sets.get(i).se(), size.width-50, size.height-50);
					if (sets.get(i).getFixes().size() == 1) {
						g2d.setColor(Color.red);
						g2d.fillOval(xyS[0]+20, xyS[1]+20, 8, 8);
					} else {
						g2d.setColor(Color.blue);
						g2d.drawRect(xyS[0]+20, xyS[1]+20, xyE[0]-xyS[0], xyE[1]-xyS[1]);
					}
				}
			}
		}

		@Override
		public void paintComponent(Graphics g) {
			super.paintComponent(g);
			draw(g);
		}
	}

	private OptimizerBrokenLine opt;

	private Candidate current;

	private RectangleSet boundSet;

	private double xDiff, yDiff;

	private Iterator<Candidate> candIter;

	public OptimizerBrokenLineAnim(OptimizerBrokenLine opt) {
		this.opt = opt;

		initUI();

		candIter = opt.iterator();
		boundSet = candIter.next().getRectangles().get(0);
		xDiff = boundSet.ne().lon() - boundSet.nw().lon();
		yDiff = boundSet.nw().lat() - boundSet.sw().lat();
	}

	private int[] latLon2XY(Fix p, int w, int h) {

		int x = (int)(((p.lon()-boundSet.nw().lon()) / xDiff) * w);
		int y = (int)(((boundSet.nw().lat()-p.lat()) / yDiff) * h);


		return new int[] {x,y};
	}

	private void initUI() {
		setTitle("Candidate animation");
		setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);

		setSize(800,600);
		setLocationRelativeTo(null);

		CandidatePanel candPanel = new CandidatePanel();
		candPanel.getInputMap().put(KeyStroke.getKeyStroke("SPACE"), "repaint");
		candPanel.getActionMap().put("repaint", new AbstractAction() {
				public void actionPerformed(ActionEvent e) {
					if (!candIter.hasNext())
						return;

					current = candIter.next();
					repaint();
				}
			}
		);
		add(candPanel);
	}

	public static void main(final String[] args) {
		SwingUtilities.invokeLater(new Runnable() {
			@Override
			public void run() {
				try {
					Parser parser = new Parser();
					Flight flight = parser.parse(
						FileSystems.getDefault().getPath(args[0]));
					OptimizerBrokenLine opt = new OptimizerBrokenLine(flight, Integer.parseInt(args[1]));
					OptimizerBrokenLineAnim optAnim = new OptimizerBrokenLineAnim(opt);
					optAnim.setVisible(true);
				} catch(Exception e) { System.out.println("EXCEPTION!!"); e.printStackTrace(); }
			}
		});
	}

}


