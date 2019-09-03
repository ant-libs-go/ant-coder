/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: __TABLE_NAME___search.go
# Description:
####################################################################### */

package __TABLE_NAME__

import (
	"__PROJECT_NAME__/models"

	"github.com/ant-libs-go/util"
	"github.com/go-xorm/builder"
)

type __TABLE_NAME_CAMEL__Search struct {
	__TABLE_NAME_CAMEL__
	models.Pager
	Sort  []*models.SortParams
	query *__TABLE_NAME_CAMEL__Query
}

func NewSearch() *__TABLE_NAME_CAMEL__Search {
	o := &__TABLE_NAME_CAMEL__Search{}
	o.query = New(nil)
	o.Limit = 100
	return o
}

func (this *__TABLE_NAME_CAMEL__Search) Load(inp interface{}, excludes ...string) *__TABLE_NAME_CAMEL__Search {
	util.Assign(inp, this, excludes...)
	return this
}

func (this *__TABLE_NAME_CAMEL__Search) SetSort(sort []*models.SortParams) *__TABLE_NAME_CAMEL__Search {
	this.Sort = sort
	return this
}

func (this *__TABLE_NAME_CAMEL__Search) FilterIds(ids []int32) *__TABLE_NAME_CAMEL__Search {
	this.query.And(builder.Eq{"id": ids})
	return this
}

func (this *__TABLE_NAME_CAMEL__Search) buildCond() *__TABLE_NAME_CAMEL__Query {
	this.query.Active()

	if this.Id > 0 {
		this.query.And(builder.Eq{"id": this.Id})
	}
	/*
		if len(this.Name) > 0 {
			this.query.And(builder.Like{"name", this.Name})
		}
	*/
	return this.query
}

func (this *__TABLE_NAME_CAMEL__Search) Count() (r int64, err error) {
	return this.buildCond().Count()
}

func (this *__TABLE_NAME_CAMEL__Search) Search() (r []*__TABLE_NAME_CAMEL__, r2 map[int32]*__TABLE_NAME_CAMEL__, err error) {
	this.buildCond().Limit(this.Limit)

	if len(this.Sort) > 0 {
		this.query.OrderBy(models.ParseSortParams(this.Sort)...)
	} else {
		this.query.OrderBy("id DESC")
	}
	if this.Offset > 0 {
		this.query.Limit(this.Limit, this.Offset)
	}
	return this.query.Find()
}
