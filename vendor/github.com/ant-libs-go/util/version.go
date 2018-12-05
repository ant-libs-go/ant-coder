/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-11-30 11:27:02
# File Name: version.go
# Description:
####################################################################### */

package util

import (
	"fmt"
	"strconv"
	"strings"
)

func VersionStringToInt(inp string) int32 {
	var s []string
	for _, v := range strings.Split(inp, ".") {
		if len(v) == 1 {
			v = fmt.Sprintf("0%s", v)
		}
		s = append(s, v)
	}
	r, _ := strconv.ParseInt(strings.Join(s, ""), 10, 64)
	return int32(r)
}

func VersionIntToString(inp int32) string {
	var s []string
	for _, v := range []int32{10000, 100, 1} {
		broken := inp / v
		inp -= broken * v
		s = append(s, strconv.Itoa(int(broken)))
	}
	return strings.Join(s, ".")
}
