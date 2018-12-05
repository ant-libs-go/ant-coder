/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-06-06 09:12:31
# File Name: json.go
# Description:
####################################################################### */

package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func JsonEncode(origin interface{}) (target string, err error) {
	var b []byte

	switch m := origin.(type) {
	case map[int32]int32:
		t := make(map[string]int32)
		for k, v := range m {
			t[strconv.Itoa(int(k))] = v
		}
		b, err = json.Marshal(t)
	case map[int32][]int32:
		t := make(map[string][]int32)
		for k, v := range m {
			t[strconv.Itoa(int(k))] = v
		}
		b, err = json.Marshal(t)
	case []int64, []int32, []int, []string, map[string]string, map[string][]int, [][]int32, [][]int64, [][]string, []map[string]string:
		b, err = json.Marshal(m)
		if string(b) == "null" {
			b = []byte("[]")
		}
	default:
		return "", fmt.Errorf("not support type: %s", reflect.TypeOf(origin))
	}

	if err == nil {
		target = string(b)
	}
	return
}

func JsonDecode(origin string, target interface{}) (err error) {

	switch r := target.(type) {
	case *map[int32]int32:
		m := make(map[int32]int32)
		t := make(map[string]int32)
		err = json.Unmarshal([]byte(origin), &t)
		for k, v := range t {
			ki, _ := strconv.Atoi(k)
			m[int32(ki)] = v
		}
		*r = m
	case *map[int32][]int32:
		m := make(map[int32][]int32)
		t := make(map[string][]int32)
		err = json.Unmarshal([]byte(origin), &t)
		for k, v := range t {
			ki, _ := strconv.Atoi(k)
			m[int32(ki)] = v
		}
		*r = m
	case *[]int64, *[]int32, *[]string, *map[string]string, *[][]string, *[][]int64, *[][]int32, *[]map[string]string:
		err = json.Unmarshal([]byte(origin), &r)
	default:
		return fmt.Errorf("not support type: %s", reflect.TypeOf(target))
	}
	return
}
