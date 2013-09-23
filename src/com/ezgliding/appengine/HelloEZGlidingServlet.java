package com.ezgliding.appengine;

import java.io.IOException;
import java.util.logging.Logger;
import javax.servlet.http.*;

public class HelloEZGlidingServlet extends HttpServlet {
	
	private static final Logger logger = Logger.getLogger(HelloEZGlidingServlet.class.getName());

	@Override
	public void doGet(HttpServletRequest req, HttpServletResponse resp)
			throws IOException {
		logger.info("HelloEZGliding");

		resp.setContentType("text/plain");
		resp.getWriter().println("Hello world from ezgliding");
	}
}
