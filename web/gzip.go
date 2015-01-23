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
	"compress/gzip"
	"io"
	"net/http"
)

// GzipResponseWriter gzips the data passed to Writer.
type GzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
	gz *gzip.Writer
}

// NewGzipResponseWriter creates a new Writer, which gzips data passed to h.
func NewGzipResponseWriter(w http.ResponseWriter) *GzipResponseWriter {
	w.Header().Set("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	return &GzipResponseWriter{
		Writer: gz, ResponseWriter: w, gz: gz,
	}
}

// Write writes the given bytes (io.Writer).
func (gw GzipResponseWriter) Write(b []byte) (int, error) {
	return gw.Writer.Write(b)
}

// Close closes the gzip writer.
func (gw GzipResponseWriter) Close() error {
	return gw.gz.Close()
}
