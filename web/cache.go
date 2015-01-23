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
	"bytes"
	"net/http"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
)

const (
	// DefaultExpiration is the default item expiration time (in seconds)
	DefaultExpiration = 120
)

// CacheItem holds an Item to be stored in the cache.
type CacheItem struct {
	s  *bytes.Buffer
	hs []http.Header
}

// CacheResponseWriter holds a bytes buffer (writer) and a http ResponseWriter.
type CacheResponseWriter struct {
	http.ResponseWriter
	b *bytes.Buffer
	m *memcache.Client
	r *http.Request
}

// NewCacheResponseWriter instantiates a new instance.
func NewCacheResponseWriter(w http.ResponseWriter, m *memcache.Client,
	r *http.Request) *CacheResponseWriter {
	b := new(bytes.Buffer)
	return &CacheResponseWriter{
		ResponseWriter: w, b: b, m: m, r: r,
	}
}

// Writer is the io.Writer implementation.
func (cw CacheResponseWriter) Write(b []byte) (int, error) {
	cw.b.Write(b)
	return cw.ResponseWriter.Write(b)
}

// Close flushes the written buffer to memcache.
func (cw CacheResponseWriter) Close() error {
	// TODO(ricardo): as Close() is called in defer, might be a good idea to log here
	// but need to find a way to test failure
	return cw.m.Set(&memcache.Item{
		Key: cw.requestKey(cw.r), Value: cw.b.Bytes(),
		Flags: 0, Expiration: DefaultExpiration,
	})
}

// requestKey returns a string with the memcache key for the given request.
func (cw CacheResponseWriter) requestKey(r *http.Request) string {
	return r.URL.String() + strings.ToLower(r.Header.Get("Accept")) + strings.ToLower(r.Header.Get("Accept-Encoding"))

}
