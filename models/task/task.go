/* ######################################################################
# Author: (zhengfei@fcadx.cn)
# Created Time: 2018-12-03 18:02:53
# File Name: task.go
# Description:
####################################################################### */

package task

import (
	"errors"
	"strings"

	"ant-coder/models"
	"fcad_thrift/enums"

	"github.com/ant-libs-go/util"
	"github.com/go-xorm/builder"
	"github.com/go-xorm/xorm"
)

type DBTask struct {
	Id          int32 `xorm:"pk autoincr"`
	WxAppid     string
	WxSecret    string
	Name        string
	AccessToken string
	ExpiresTime int32
	Status      enums.InfoStatus
	CreatedAt   int64 `xorm:"created"`
	UpdatedAt   int64 `xorm:"updated"`
}

func (this *DBTask) TableName() string {
	return "test"
}

func (this *DBTask) Thrift() interface{} {
	// input your codes
	return nil
}

type DBTaskQuery struct {
	DBTask      `xorm:"-"`
	session     *xorm.Session `xorm:"-"`
	isNewRecord bool          `xorm:"-"`
}

func NewDBTaskQuery() *DBTaskQuery {
	o := &DBTaskQuery{}
	o.isNewRecord = true
	return o
}

func (this *DBTaskQuery) Orm() (r *xorm.Engine) {
	return models.Orm
}

func (this *DBTaskQuery) Session() (r *xorm.Session) {
	if this.session == nil {
		this.session = this.Orm().NewSession()
	}
	return this.session
}

func (this *DBTaskQuery) Load(inp interface{}, excludes ...string) *DBTaskQuery {
	util.Assign(inp, this, excludes...)
	return this
}

func (this *DBTaskQuery) SQL(query string, params ...interface{}) (r *DBTaskQuery) {
	this.Session().Sql(query, params...)
	return this
}

/* cond */
func (this *DBTaskQuery) Where(cond builder.Cond) *DBTaskQuery {
	this.Session().Where(cond)
	return this
}

func (this *DBTaskQuery) And(cond builder.Cond) *DBTaskQuery {
	this.Session().And(cond)
	return this
}

func (this *DBTaskQuery) Or(cond builder.Cond) *DBTaskQuery {
	this.Session().Or(cond)
	return this
}

/* misc */
func (this *DBTaskQuery) Cols(cols ...string) *DBTaskQuery {
	this.Session().Cols(cols...)
	return this
}

func (this *DBTaskQuery) Select(str string) *DBTaskQuery {
	this.Session().Select(str)
	return this
}

func (this *DBTaskQuery) OrderBy(orders ...string) *DBTaskQuery {
	this.Session().OrderBy(strings.Join(orders, ", "))
	return this
}

func (this *DBTaskQuery) GroupBy(groups ...string) *DBTaskQuery {
	this.Session().OrderBy(strings.Join(groups, ", "))
	return this
}

func (this *DBTaskQuery) Limit(limit int, offset ...int) *DBTaskQuery {
	this.Session().Limit(limit, offset...)
	return this
}

/* query */
func (this *DBTaskQuery) Get() (r *DBTask, err error) {
	r = &DBTask{}
	if has, err := this.Session().Get(r); has == false || err != nil {
		return nil, err
	}
	this.isNewRecord = false
	util.Assign(r, this)
	return
}

func (this *DBTaskQuery) Find() (r []*DBTask, r2 map[int32]*DBTask, err error) {
	err = this.Session().Find(&r)
	r2 = map[int32]*DBTask{}
	for _, m := range r {
		r2[m.Id] = m
	}
	return
}

func (this *DBTaskQuery) Count() (r int64, err error) {
	return this.Session().Count(&DBTask{})
}

func (this *DBTaskQuery) Exist() (r bool, err error) {
	return this.Session().Exist(&DBTask{})
}

/* update */
func (this *DBTaskQuery) Update(engine interface{}) (err error) {
	if this.Id == 0 {
		return errors.New("id not set")
	}

	if engine == nil {
		_, err = this.Where(builder.Eq{"id": this.Id}).Session().Update(&this.DBTask)
	}
	if eng, ok := engine.(*xorm.Session); ok {
		_, err = eng.Where("id = ?", this.Id).Update(&this.DBTask)
	}
	return
}

/* delete */
func (this *DBTaskQuery) Delete(engine interface{}) (err error) {
	if this.Id == 0 {
		return errors.New("id not set")
	}
	this.Status = enums.InfoStatus_Deleted
	this.Cols("status").Update(engine)
	return
}
