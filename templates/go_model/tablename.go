/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: __TABLE_NAME__.go
# Description:
####################################################################### */

/**
 * Define constant:
 *
 * type InfoStatus int
 * const (
 *   InfoStatusNormal  InfoStatus = 0
 *   InfoStatusInvalid InfoStatus = 1
 * )
 *
 * type SortType int64
 *
 * const (
 *   SortTypeAsc  SortType = 0
 *   SortTypeDesc SortType = 1
 * )
 *
 * type SortParams struct {
 *   Field string
 *   Type  SortType
 * }
 *
 * type Pager struct {
 *   Offset int `form:"offset"`
 *   Limit  int `form:"limit"`
 * }
 *
 * ErrNotFound = errors.New("record is not found")
 *
 * func ParseSortParams(sorts []*common.SortParams) (r []string) {
 *   r = []string{}
 *   for _, sort := range sorts {
 *     if sort.Type == enums.SortType_Asc {
 *       r = append(r, fmt.Sprintf("%s ASC", sort.Field))
 *     }
 *     if sort.Type == enums.SortType_Desc {
 *       r = append(r, fmt.Sprintf("%s DESC", sort.Field))
 *     }
 *   }
 *   return
 * }
 */

package __TABLE_NAME__

import (
	"strings"

	"__PROJECT_NAME__/models"

	"github.com/ant-libs-go/util"
	"github.com/go-xorm/builder"
	"github.com/go-xorm/xorm"
)

type __TABLE_NAME_CAMEL__ struct {
	__TABLE_META_STRUCT__
}

func (this *__TABLE_NAME_CAMEL__) TableName() string {
	return "__TABLE_NAME__"
}

/*
func (this *__TABLE_NAME_CAMEL__) Format() (r *types.__TABLE_NAME_CAMEL__) {
	r = &types.__TABLE_NAME_CAMEL__{}
	util.Assign(this, r)
	return
}
*/

type __TABLE_NAME_CAMEL__Query struct {
	__TABLE_NAME_CAMEL__ `xorm:"-"`
	session              *xorm.Session `xorm:"-"`
	isNewRecord          bool          `xorm:"-"`
}

func New(session *xorm.Session) *__TABLE_NAME_CAMEL__Query {
	o := &__TABLE_NAME_CAMEL__Query{}
	o.isNewRecord = true
	o.session = session
	return o
}

func (this *__TABLE_NAME_CAMEL__Query) Orm() (r *xorm.EngineGroup) {
	return models.Orm
}

func (this *__TABLE_NAME_CAMEL__Query) Session() (r *xorm.Session) {
	if this.session == nil {
		this.session = this.Orm().NewSession().Table(this)
	}
	return this.session
}

func (this *__TABLE_NAME_CAMEL__Query) Load(inp interface{}, excludes ...string) *__TABLE_NAME_CAMEL__Query {
	util.Assign(inp, this, excludes...)
	if this.Id != 0 {
		this.isNewRecord = false
	}
	return this
}

// r, err := query.SQL("select * from test where id = ?", 2).Get()
func (this *__TABLE_NAME_CAMEL__Query) SQL(query string, params ...interface{}) (r *__TABLE_NAME_CAMEL__Query) {
	this.Session().Sql(query, params...)
	return this
}

/* cond */
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

// r, r2, err := query.OrderBy("type", "status").Find()
func (this *__TABLE_NAME_CAMEL__Query) GroupBy(groups ...string) *__TABLE_NAME_CAMEL__Query {
	this.Session().OrderBy(strings.Join(groups, ", "))
	return this
}

// r, err := query.Where(builder.Eq{"id": 3}).Limit(1, 1).Get()
func (this *__TABLE_NAME_CAMEL__Query) Limit(limit int, offset ...int) *__TABLE_NAME_CAMEL__Query {
	this.Session().Limit(limit, offset...)
	return this
}

/* query */
func (this *__TABLE_NAME_CAMEL__Query) Get() (r *__TABLE_NAME_CAMEL__, err error) {
	r = &__TABLE_NAME_CAMEL__{}
	has, err := this.Session().Get(r)
	if err != nil {
		return nil, err
	}
	if has == false {
		return nil, models.ErrNotFound
	}

	this.Load(r)
	return
}

func (this *__TABLE_NAME_CAMEL__Query) Find() (r []*__TABLE_NAME_CAMEL__, r2 map[int32]*__TABLE_NAME_CAMEL__, err error) {
	err = this.Session().Find(&r)
	r2 = map[int32]*__TABLE_NAME_CAMEL__{}
	for _, m := range r {
		r2[m.Id] = m
	}
	return
}

// r, err := query.Count()
func (this *__TABLE_NAME_CAMEL__Query) Count() (r int64, err error) {
	return this.Session().Count(&__TABLE_NAME_CAMEL__{})
}

// r, err := query.Where(builder.Eq{"id": 2}).Exist()
func (this *__TABLE_NAME_CAMEL__Query) Exist() (r bool, err error) {
	return this.Session().Exist(&__TABLE_NAME_CAMEL__{})
}
