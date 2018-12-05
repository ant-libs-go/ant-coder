/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-06-06 09:31:40
# File Name: tag.go
# Description:
####################################################################### */

package util

import (
	"reflect"
	"strings"
)

// TODO
/*
type S sturct {
	Name stirng `tagname:range(1,10);require`
	Age  int32  `tagname:range(1,10);require`
}
*/

type FieldTag struct {
	Name    string
	Range   string
	Require bool
}

func ParseTags(obj interface{}, tagname string) (r map[string]*FieldTag) {
	r = map[string]*FieldTag{}

	elem := reflect.TypeOf(obj).Elem()
	for i := 0; i < elem.NumField(); i++ {
		tags := elem.Field(i).Tag.Get(tagname)
		if tags == "-" {
			continue
		}
		ft := &FieldTag{}
		ft.Name = elem.Field(i).Name
		ftElem := reflect.ValueOf(ft).Elem()
		for _, tag := range strings.Split(tags, ";") {
			tag = strings.ToLower(strings.TrimSpace(tag))
			if tag == "" {
				continue
			}
			if i := strings.Index(tag, "("); i > 0 && strings.Index(tag, ")") == len(tag)-1 {
				val := tag[i+1 : len(tag)-1]
				ftElem.FieldByName(CamelString(tag[:i])).SetString(val)
			} else {
				ftElem.FieldByName(CamelString(tag)).SetBool(true)
			}
		}
		r[ft.Name] = ft
	}
	return
}
