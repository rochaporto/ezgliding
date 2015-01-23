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
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/golang/glog"
	"github.com/rochaporto/ezgliding/config"
	"github.com/rochaporto/ezgliding/context"
)

const (
	// Port is the default port the server listens to
	Port = 80
	// Static is the fs location of the static files (html, css, js, ...)
	Static = "web/static"
)

// Server is a web server implementation, serving ezgliding data.
type Server struct {
	ctx      context.Context
	mux      *http.ServeMux
	memcache *memcache.Client
	Port     int
	Static   string
	Memcache string
}

// Init prepares the web server to be started.
func (srv *Server) Init(ctx context.Context) error {
	var z context.Context
	if ctx == z {
		return errors.New("got a zero value Context, cannot handle this")
	}
	srv.ctx = ctx
	// handle config
	err := srv.setConfig(srv.ctx.Config)
	if err != nil {
		return fmt.Errorf("failed to init with config :: %v", err)
	}
	// init the memcache client if appropriate
	if srv.Memcache != "" {
		srv.memcache = memcache.New(srv.Memcache)
		_, err := srv.memcache.Get("nonexistingitem")
		if err != nil && err != memcache.ErrCacheMiss {
			return err
		}
	}
	// set handlers
	srv.mux = http.NewServeMux()
	srv.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(srv.Static))))
	srv.mux.HandleFunc("/airfield/", srv.makeHandler(srv.airspaceHandler))
	srv.mux.HandleFunc("/waypoint/", srv.makeHandler(srv.waypointHandler))
	return nil
}

// Start starts the http server instance.
func (srv *Server) Start() {
	glog.V(0).Infof("starting web server :: %v", *srv)
	http.ListenAndServe(fmt.Sprintf(":%v", srv.Port), srv.mux)
}

func (srv *Server) setConfig(cfg config.Config) error {
	if cfg.Web.Port == 0 {
		srv.Port = Port
	} else {
		srv.Port = cfg.Web.Port
	}
	if cfg.Web.Static == "" {
		srv.Static = Static
	} else {
		srv.Static = cfg.Web.Static
	}
	if _, err := os.Stat(srv.Static); os.IsNotExist(err) {
		return fmt.Errorf("static location %v does not exist", srv.Static)
	}
	srv.Memcache = cfg.Web.Memcache
	return nil
}
