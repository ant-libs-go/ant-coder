/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-11-30 11:25:26
# File Name: string.go
# Description:
####################################################################### */

package util

import (
	"errors"
	"regexp"
	"strings"
)

// camel string, xx_yy to XxYy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// snake string, XxYy to xx_yy, XxYY to xx_yy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// regexp.Compile(`\[(?P<node>[\d_]+)\]$`)
// return {"node":val}
func FindStringSubmatch(re *regexp.Regexp, s string) (r map[string]string, err error) {
	r = make(map[string]string)
	match := re.FindStringSubmatch(s)
	if match == nil {
		return nil, errors.New("no match")
	}
	for i, name := range re.SubexpNames() {
		if i == 0 || name == "" {
			continue
		}
		r[name] = match[i]
	}
	return
}
