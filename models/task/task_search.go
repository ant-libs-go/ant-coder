/* ######################################################################
# Author: (zhengfei@fcadx.cn)
# Created Time: 2018-12-05 14:33:22
# File Name: task_search.go
# Description:
####################################################################### */

package task

import (
	"github.com/ant-libs-go/util"
	"github.com/go-xorm/builder"
)

type UserSearch struct {
	User
	Ids      []int32
	LastId   int32
	PageSize int
}

func NewSearch() *UserSearch {
	o := &UserSearch{}
	o.PageSize = 10
	return o
}

func (this *UserSearch) Load(inp interface{}, excludes ...string) *UserSearch {
	util.Assign(inp, this, excludes...)
	return this
}

func (this *UserSearch) Search() (r []*User, r2 map[int32]*User, err error) {
	query := New().Active().OrderBy("id DESC").Limit(this.PageSize)

	if this.LastId > 0 {
		query.And(builder.Lt{"id": this.LastId})
	}
	if this.Id > 0 {
		query.And(builder.Eq{"id": this.Id})
	}
	if this.Id > 0 {
		query.And(builder.Eq{"id": this.Ids})
	}
	if len(this.Name) > 0 {
		query.And(builder.Like{"name", this.Name})
	}
	return query.Find()
}
