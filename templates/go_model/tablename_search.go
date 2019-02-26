/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: __TABLE_NAME___search.go
# Description:
####################################################################### */

package __TABLE_NAME__

import (
	"github.com/ant-libs-go/util"
	"github.com/go-xorm/builder"
)

type Pager struct {
	Page     int32 `form:"page"`
	LastId   int32 `form:"last_id"`
	PageSize int32 `form:"page_size"`
}

type __TABLE_NAME_CAMEL__Search struct {
	__TABLE_NAME_CAMEL__
	Pager
	Ids []int32
}

func NewSearch() *__TABLE_NAME_CAMEL__Search {
	o := &__TABLE_NAME_CAMEL__Search{}
	o.PageSize = 100
	return o
}

func (this *__TABLE_NAME_CAMEL__Search) Load(inp interface{}, excludes ...string) *__TABLE_NAME_CAMEL__Search {
	util.Assign(inp, this, excludes...)
	return this
}

func (this *__TABLE_NAME_CAMEL__Search) FilterIds(ids []int32) *__TABLE_NAME_CAMEL__Search {
	this.Ids = ids
	return this
}

func (this *__TABLE_NAME_CAMEL__Search) Search() (r []*__TABLE_NAME_CAMEL__, r2 map[int32]*__TABLE_NAME_CAMEL__, err error) {
	query := New().Active().OrderBy("id DESC").Limit(this.PageSize)

	if this.Id > 0 {
		query.And(builder.Eq{"id": this.Id})
	}
	if this.Ids > 0 {
		query.And(builder.Eq{"id": this.Ids})
	}
	if this.LastId > 0 {
		query.And(builder.Lt{"id": this.LastId})
	}
	/*
		if len(this.Name) > 0 {
			query.And(builder.Like{"name", this.Name})
		}
	*/
	if this.Page > 0 {
		if this.Page == 0 {
			this.Page = 1
		}
		query.Limit(this.PageSize, (this.Page-1)*this.PageSize)
	}
	return query.Find()
}
