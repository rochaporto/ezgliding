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

package web

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/rochaporto/ezgliding/common"
)

// makeHandler is a common wrapper for all handlers.
// Allows having common bits of code for all handlers (logging, err handling, ...).
func (srv *Server) makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		glog.V(2).Infof("%v", *r)
		params := r.URL.Query()
		if v, ok := params["accept"]; ok {
			glog.V(10).Infof("adding %v accept from querystring", v)
			r.Header.Set("Accept", fmt.Sprintf("%v,%v", r.Header.Get("accept"), v))
		}
		// check for cache and return immediately if present (and valid)
		/**if cached := CheckCache(w, r); cached {
			return
		}*/

		// no cache, so prepare a response writer for the handler
		writer := w
		// gzip if requested
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			writer = NewGzipResponseWriter(writer)
			defer writer.(io.Closer).Close()
		}
		// cache the generated response if requested
		if srv.memcache != nil {
			writer = NewCacheResponseWriter(w, srv.memcache, r)
			defer writer.(io.Closer).Close()
		}
		fn(writer, r)
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
	format := srv.accept(r.Header.Get("Accept"))
	result, err := srv.toOutput(format, wrapper)
	w.Header().Set("Content-Type", format)
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
	format := srv.accept(r.Header.Get("Accept"))
	result, err := srv.toOutput(format, wrapper)
	w.Header().Set("Content-Type", format)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%v", result)
}
