package com.ezgliding.igc;

import java.awt.Color;
import java.awt.Dimension;
import java.awt.Graphics;
import java.awt.Graphics2D;
import java.awt.Insets;
import java.nio.file.FileSystems;
import java.util.Random;
import javax.swing.JFrame;
import javax.swing.JPanel;
import javax.swing.SwingUtilities;

public class BrokenLineOptimizerAnim extends JFrame {

	public class CandidatePanel extends JPanel {

		private void draw(Graphics g) {
			Graphics2D g2d = (Graphics2D)g;

			g2d.setColor(Color.blue);

			Dimension size = getSize();
			Insets insets = getInsets();

			int w = size.width - insets.left - insets.right;
			int h = size.height - insets.top - insets.bottom;

			g2d.drawRect(0,0,100,100);
		}


		@Override
		public void paintComponent(Graphics g) {
			super.paintComponent(g);
			draw(g);
		}
	}

	private BrokenLineOptimizer opt;

	public BrokenLineOptimizerAnim(BrokenLineOptimizer opt) {
		this.opt = opt;

		initUI();
	}

	private void initUI() {
		setTitle("Candidate animation");
		setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);

		add(new CandidatePanel());

		setSize(800,600);
		setLocationRelativeTo(null);
	}

	public static void main(String[] args) {
		SwingUtilities.invokeLater(new Runnable() {
			@Override
			public void run() {
				try {
					Parser parser = new Parser();
					Flight flight = parser.parse(
						FileSystems.getDefault().getPath("test/data/optimize-with-5-points-only.igc"));
					BrokenLineOptimizer opt = new BrokenLineOptimizer(flight, 5);
	
					BrokenLineOptimizerAnim optAnim = new BrokenLineOptimizerAnim(opt);
					optAnim.setVisible(true);
				} catch(Exception e) { }
			}
		});
	}

}


