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

type DBTaskSearch struct {
	DBTaskQuery
	Ids    []int32
	LastId int32
}

func NewDBTaskSearch() *DBTaskSearch {
	o := &DBTaskSearch{}
	return o
}

func (this *DBTaskSearch) load(inp interface{}) {
	util.Assign(inp, this)
}

func (this *DBTaskSearch) Search(params interface{}) (r []*DBTask, r2 map[int32]*DBTask, err error) {
	query := NewDBTaskQuery().Active().OrderBy("id DESC").Limit(10)
	this.load(params)

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
