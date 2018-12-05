/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-06-06 09:23:03
# File Name: map.go
# Description:
####################################################################### */

package util

import (
	"fmt"
	"strings"
)

func JoinMap(arr map[string]string, glue string, glue2 string) (r string) {
	var t []string
	for n, v := range arr {
		t = append(t, fmt.Sprintf("%s%s%s", n, glue, v))
	}
	r = strings.Join(t, glue2)
	return
}
