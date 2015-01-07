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
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/rochaporto/ezgliding/common"
	"github.com/rochaporto/ezgliding/context"
	"github.com/rochaporto/ezgliding/util"
)

const (
	// Port is the default port the server listens to
	Port = 80
	// Static is the fs location of the static files (html, css, js, ...)
	Static = "web/static"
)

// gZipResponseWriter is a response wrapper if gzip is supported.
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// Server is a web server implementation, serving ezgliding data.
type Server struct {
	ctx    context.Context
	mux    *http.ServeMux
	Port   int
	Static string
}

// Init prepares the web server to be started.
func (srv *Server) Init(ctx context.Context) error {
	var z context.Context
	if ctx == z {
		return errors.New("got a zero value Context, cannot handle this")
	}
	srv.ctx = ctx
	// set config
	cfg := ctx.Config
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
	// set handlers
	srv.mux = http.NewServeMux()
	srv.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(srv.Static))))
	srv.mux.HandleFunc("/airfield/", makeHandler(srv.airspaceHandler))
	srv.mux.HandleFunc("/waypoint/", makeHandler(srv.waypointHandler))
	return nil
}

// Start starts the http server instance.
func (srv *Server) Start() {
	glog.V(0).Infof("starting web server :: %v", *srv)
	http.ListenAndServe(fmt.Sprintf(":%v", srv.Port), srv.mux)
}

// makeHandler is a common wrapper for all handlers.
// Allows having common bits of code for all handlers (logging, err handling, ...).
func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		glog.V(2).Infof("%v", *r)
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			fn(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		fn(gzr, r)
	}
}

// airspaceHandler handles /airspace/.
func (srv *Server) airspaceHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	updated := time.Time{}
	if t, ok := params["updated"]; ok {
		var err error
		updated, err = time.Parse("2006-01-02", t[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	airfield := srv.ctx.Airfield
	airfields, err := airfield.(common.Airfielder).GetAirfield(params["region"], updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	wrapper := []interface{}{}
	for _, a := range airfields {
		wrapper = append(wrapper, a)
	}
	result, err := srv.toOutput(r.Header.Get("Accept"), wrapper)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%v", result)
}

// waypointHandler handles /waypoint/.
func (srv *Server) waypointHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	updated := time.Time{}
	if t, ok := params["updated"]; ok {
		var err error
		updated, err = time.Parse("2006-01-02", t[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	waypoint := srv.ctx.Waypoint
	waypoints, err := waypoint.(common.Waypointer).GetWaypoint(params["region"], updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	wrapper := []interface{}{}
	for _, a := range waypoints {
		wrapper = append(wrapper, a)
	}
	result, err := srv.toOutput(r.Header.Get("Accept"), wrapper)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%v", result)
}

// toOutput returns a string representation (in the requested format) of the given content.
// format should be as in the Accept: header (application/json, ...), content is an array
// of Struct which can be Airfield, Waypoint, etc.
func (srv *Server) toOutput(format string, content []interface{}) (string, error) {
	var output string
	if strings.Contains(format, "application/json") {
		collection, err := util.Struct2GeoJSON(content)
		if err != nil {
			return "", err
		}
		bytes, _ := collection.MarshalJSON()
		output = string(bytes)
		//} else if strings.Contains(format, "application/csv") { FIXME: enable csv output
		//	output = util.Struct2CSV(content)
	} else {
		return "", fmt.Errorf("format %v not supported", format)
	}
	return output, nil
}
