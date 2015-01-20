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

// Package util provides multiple functions for encoding/decoding data,
// converting between commonly used units, etc.
//
package util

import (
	"encoding/csv"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Struct2CSV converts a list of structs to CSV.
func Struct2CSV(v interface{}) string {
	// FIXME: better validation of input value
	rv := reflect.ValueOf(v)
	if rv.Len() == 0 {
		return ""
	}
	var fields []string
	val := rv.Index(0)
	c := val.Type().Field(0).Name
	fields = append(fields, c)
	for i := 1; i < val.NumField(); i++ {
		f := val.Type().Field(i).Name
		c = c + "," + f
		fields = append(fields, f)
	}
	c = c + "\n"
	for i := 0; i < rv.Len(); i++ {
		e := rv.Index(i)
		for i := range fields {
			f := e.Field(i)
			if i != 0 {
				c = c + ","
			}
			c = c + fmt.Sprintf("%v", f.Interface())
		}
		c = c + "\n"
	}
	return c
}

// CSV2Struct converts the given CSV content to a slice of struct objects.
func CSV2Struct(content string, tp reflect.Type, si reflect.Type) (reflect.Value, error) {
	parser := csv.NewReader(strings.NewReader(content))
	parsed, err := parser.ReadAll()
	if err != nil {
		return reflect.Value{}, err
	}
	if len(parsed) == 0 {
		return reflect.Value{}, errors.New("No data in content given")
	}
	var result = reflect.MakeSlice(tp, 0, len(parsed)-1)
	fields := []string{}
	for _, field := range parsed[0] {
		fields = append(fields, field)
	}
	for i := 1; i < len(parsed); i++ {
		n := reflect.New(si)
		item := n.Elem()
		for j, field := range fields {
			structF := item.FieldByName(field)
			k := structF.Kind()
			if k == reflect.String {
				structF.SetString(parsed[i][j])
			} else if k == reflect.Int {
				v, err := strconv.ParseInt(parsed[i][j], 10, 64)
				if err != nil {
					return reflect.Value{}, err
				}
				structF.SetInt(v)
			} else if k == reflect.Float64 {
				v, err := strconv.ParseFloat(parsed[i][j], 64)
				if err != nil {
					return reflect.Value{}, err
				}
				structF.SetFloat(v)
			}
			// FIXME: handle slice and struct
		}
		result = reflect.Append(result, item)
	}
	return result, nil
}
