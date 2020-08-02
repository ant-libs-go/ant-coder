/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: __TABLE_NAME___query.go
# Description:
####################################################################### */

package __TABLE_NAME__

import (
	"fmt"
	"strings"

	"__PROJECT_NAME__/models"

	"xorm.io/builder"
	"xorm.io/xorm"
)

type __TABLE_NAME_CAMEL__Query struct {
	isAutoClose bool
	session     *xorm.Session
}

func (this *__TABLE_NAME_CAMEL__Query) Session() (r *xorm.Session) {
	return this.session
}

/* basic query */

func (this *__TABLE_NAME_CAMEL__Query) Get() (r *__TABLE_NAME_CAMEL__, err error) {
	if this.isAutoClose == true {
		defer this.Session().Close()
	}
	r = &__TABLE_NAME_CAMEL__{}
	has, err := this.Session().Get(r)
	if err != nil {
		return nil, err
	}
	if has == false {
		return nil, models.ErrNotFound
	}
	return
}

func (this *__TABLE_NAME_CAMEL__Query) Find() (r []*__TABLE_NAME_CAMEL__, r2 map[int32]*__TABLE_NAME_CAMEL__, err error) {
	if this.isAutoClose == true {
		defer this.Session().Close()
	}
	err = this.Session().Find(&r)
	r2 = make(map[int32]*__TABLE_NAME_CAMEL__, len(r))
	for _, m := range r {
		r2[m.Id] = m
	}
	return
}

func (this *__TABLE_NAME_CAMEL__Query) Count() (r int64, err error) {
	if this.isAutoClose == true {
		defer this.Session().Close()
	}
	return this.Session().Count(&__TABLE_NAME_CAMEL__{})
}

func (this *__TABLE_NAME_CAMEL__Query) SumFloat(column string) (r float64, err error) {
	if this.isAutoClose == true {
		defer this.Session().Close()
	}
	return this.Session().Sum(&__TABLE_NAME_CAMEL__{}, column)
}

func (this *__TABLE_NAME_CAMEL__Query) SumInt(column string) (r int64, err error) {
	if this.isAutoClose == true {
		defer this.Session().Close()
	}
	return this.Session().SumInt(&__TABLE_NAME_CAMEL__{}, column)
}

func (this *__TABLE_NAME_CAMEL__Query) SumsFloat(columns ...string) (r []float64, err error) {
	if this.isAutoClose == true {
		defer this.Session().Close()
	}
	return this.Session().Sums(&__TABLE_NAME_CAMEL__{}, columns...)
}

func (this *__TABLE_NAME_CAMEL__Query) SumsInt(columns ...string) (r []int64, err error) {
	if this.isAutoClose == true {
		defer this.Session().Close()
	}
	return this.Session().SumsInt(&__TABLE_NAME_CAMEL__{}, columns...)
}

func (this *__TABLE_NAME_CAMEL__Query) Exist() (r bool, err error) {
	if this.isAutoClose == true {
		defer this.Session().Close()
	}
	return this.Session().Exist(&__TABLE_NAME_CAMEL__{})
}

/*
	rows, err := engine.Rows(&user{})
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(user)
	}
*/
func (this *__TABLE_NAME_CAMEL__Query) Rows() (r *xorm.Rows, err error) {
	if this.isAutoClose == true {
		defer this.Session().Close()
	}
	return this.Session().Rows(&__TABLE_NAME_CAMEL__{})
}

/* cond */

// r, err := query.SQL("select * from test where id = ?", 2).Get()
func (this *__TABLE_NAME_CAMEL__Query) SQL(query string, params ...interface{}) (r *__TABLE_NAME_CAMEL__Query) {
	this.Session().SQL(query, params...)
	return this
}

// r, err := query.Where(builder.Eq{"id": 3}).Get()
// r, err := query.Where(builder.Eq{"id": 3, "status": 0}).Get()
func (this *__TABLE_NAME_CAMEL__Query) Where(cond builder.Cond) *__TABLE_NAME_CAMEL__Query {
	this.Session().Where(cond)
	return this
}

// r, err := query.Where(builder.Eq{"id": 3}).And(builder.Eq{"status": 0}).Get()
func (this *__TABLE_NAME_CAMEL__Query) And(cond builder.Cond) *__TABLE_NAME_CAMEL__Query {
	this.Session().And(cond)
	return this
}

// r, err := query.Where(builder.Eq{"id": 3}).Or(builder.Eq{"status": 0}).Get()
func (this *__TABLE_NAME_CAMEL__Query) Or(cond builder.Cond) *__TABLE_NAME_CAMEL__Query {
	this.Session().Or(cond)
	return this
}

/* misc */

// r, err := query.Cols("id", "name").Where(builder.Eq{"id": 3}).Get()
func (this *__TABLE_NAME_CAMEL__Query) Cols(cols ...string) *__TABLE_NAME_CAMEL__Query {
	this.Session().Cols(cols...)
	return this
}

// 查询或更新所有字段，常与Update配合使用，因Update默认只更新非0非bool字段
// r, err := query.AllCols().Where(builder.Eq{"id": 3}).Get()
func (this *__TABLE_NAME_CAMEL__Query) AllCols() *__TABLE_NAME_CAMEL__Query {
	this.Session().AllCols()
	return this
}

// 与Cols相反
// r, err := query.Omit("id", "name").Where(builder.Eq{"id": 3}).Update(&User)
func (this *__TABLE_NAME_CAMEL__Query) Omit(cols ...string) *__TABLE_NAME_CAMEL__Query {
	this.Session().Omit(cols...)
	return this
}

// r, err := query.Select("id, name").Where(builder.Eq{"id": 3}).Get()
func (this *__TABLE_NAME_CAMEL__Query) Select(str string) *__TABLE_NAME_CAMEL__Query {
	this.Session().Select(str)
	return this
}

// r, r2, err := query.Where(builder.Eq{"id": 3}).OrderBy("id desc", "tid asc").Find()
func (this *__TABLE_NAME_CAMEL__Query) OrderBy(orders ...string) *__TABLE_NAME_CAMEL__Query {
	this.Session().OrderBy(strings.Join(orders, ", "))
	return this
}

// r, r2, err := query.GroupBy("type", "status").Find()
func (this *__TABLE_NAME_CAMEL__Query) GroupBy(groups ...string) *__TABLE_NAME_CAMEL__Query {
	this.Session().GroupBy(strings.Join(groups, ", "))
	return this
}

// r, err := query.Where(builder.Eq{"id": 3}).Limit(1, 1).Get()
func (this *__TABLE_NAME_CAMEL__Query) Limit(limit int, offset ...int) *__TABLE_NAME_CAMEL__Query {
	this.Session().Limit(limit, offset...)
	return this
}

/* business */

// &Struct{}
func (this *__TABLE_NAME_CAMEL__Query) Insert(inp interface{}) (err error) {
	m := &__TABLE_NAME_CAMEL__{}
	_, err = this.Session().Insert(m.Load(inp))
	return
}

// []*Struct{}
func (this *__TABLE_NAME_CAMEL__Query) InsertAll(inps ...interface{}) (err error) {
	ms := []*__TABLE_NAME_CAMEL__{}
	for _, inp := range inps {
		m := &__TABLE_NAME_CAMEL__{}
		ms = append(ms, m.Load(inp))
	}
	_, err = this.Session().Insert(ms)
	return
}

/* insert or update */
// notice: unsafe, unrealized
func (this *__TABLE_NAME_CAMEL__Query) InsertOrUpdate(inp interface{}, keys builder.Cond) (err error) {
	return
}

/* get or insert */
// notice: unsafe, unrealized
func (this *__TABLE_NAME_CAMEL__Query) GetOrInsert(cond builder.Cond, inp interface{}) (err error) {
	return
}

__BUSINESS_FUNCTIONS__

/* user custom */

// ....

// vim: set noexpandtab ts=4 sts=4 sw=4 :
