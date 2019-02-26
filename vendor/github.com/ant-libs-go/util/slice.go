/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-06-06 07:47:40
# File Name: slice.go
# Description:
####################################################################### */

package util

import (
	"reflect"
)

func InSlice(val interface{}, slice interface{}) (exists bool, index int) {
	exists = false
	index = -1

	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return
	}
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(val, s.Index(i).Interface()) == false {
			continue
		}
		index = i
		exists = true
		return
	}
	return
}

func SliceDiff(slice1, slice2 interface{}) (r []interface{}) {
	if reflect.TypeOf(slice1).Kind() != reflect.Slice || reflect.TypeOf(slice2).Kind() != reflect.Slice {
		return
	}
	s := reflect.ValueOf(slice1)
	for i := 0; i < s.Len(); i++ {
		if exists, _ := InSlice(s.Index(i).Interface(), slice2); exists {
			continue
		}
		r = append(r, s.Index(i).Interface())
	}
	s = reflect.ValueOf(slice2)
	for i := 0; i < s.Len(); i++ {
		if exists, _ := InSlice(s.Index(i).Interface(), slice1); exists {
			continue
		}
		r = append(r, s.Index(i).Interface())
	}
	return
}

func SliceUnique(slice interface{}) (r []interface{}) {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return
	}
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		if exists, _ := InSlice(s.Index(i).Interface(), r); exists {
			continue
		}
		r = append(r, s.Index(i).Interface())
	}
	return
}

func SliceColumn(slice interface{}, col string) (r interface{}) {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return
	}
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		f := s.Index(i).Elem().FieldByName(col)
		if f.IsValid() != true {
			continue
		}
		switch f.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
			if r == nil {
				r = []int32{}
			}
			r = append(r.([]int32), int32(f.Int()))
		case reflect.Int64:
			if r == nil {
				r = []int64{}
			}
			r = append(r.([]int64), f.Int())
		case reflect.String:
			if r == nil {
				r = []string{}
			}
			r = append(r.([]string), f.String())
		}
	}
	return
}
