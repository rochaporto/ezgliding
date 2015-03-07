// Copyright 2014 The ezgliding Authors.
//
// This file is part of ezgliding.
//
// ezgliding is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ezgliding is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ezgliding.  If not, see <http://www.gnu.org/licenses/>.
//
// Author: Ricardo Rocha <rocha.porto@gmail.com>

// Package web provides a web server implementation, serving ezgliding data.
package web

import (
	"fmt"
	"net/http"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/golang/glog"
	"github.com/rochaporto/ezgliding/airfield"
	"github.com/rochaporto/ezgliding/waypoint"
)

const (
	// Port is the default port the server listens to
	Port = 80
	// Static is the fs location of the static files (html, css, js, ...)
	Static = "web/static"
)

// Config holds all the web server configuration.
type Config struct {
	Port       int
	Static     string
	Memcache   string
	Airfielder airfield.Airfielder
	Waypointer waypoint.Waypointer
}

// Server is a web server implementation, serving ezgliding data.
type Server struct {
	Config
	mux      *http.ServeMux
	memcache *memcache.Client
}

// NewServer returns a new instance of a web.Server.
func NewServer(cfg Config) (*Server, error) {
	// set defaults if appropriate
	if cfg.Port == 0 {
		cfg.Port = Port
	}
	if cfg.Static == "" {
		cfg.Static = Static
	}
	srv := Server{Config: cfg}
	// init the memcache client if appropriate
	if srv.Memcache != "" {
		srv.memcache = memcache.New(srv.Memcache)
		_, err := srv.memcache.Get("nonexistingitem")
		if err != nil && err != memcache.ErrCacheMiss {
			return &srv, err
		}
	}
	// set handlers
	srv.mux = http.NewServeMux()
	srv.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(srv.Static))))
	srv.mux.HandleFunc("/airfield/", srv.makeHandler(srv.airspaceHandler))
	srv.mux.HandleFunc("/waypoint/", srv.makeHandler(srv.waypointHandler))
	return &srv, nil
}

// Start starts the http server instance.
func (srv *Server) Start() {
	glog.V(0).Infof("starting web server :: %v", *srv)
	http.ListenAndServe(fmt.Sprintf(":%v", srv.Port), srv.mux)
}
