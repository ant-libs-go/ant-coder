/* ######################################################################
# Author: (zhengfei@fcadx.cn)
# Created Time: 2018-12-03 18:02:53
# File Name: User.go
# Description:
####################################################################### */

package User

import (
	"strings"

	"ant-coder/models"

	"github.com/ant-libs-go/util"
	"github.com/go-xorm/builder"
	"github.com/go-xorm/xorm"
	"gitlab.com/feichi/fcad_thrift/libs/go/enums"
	types "gitlab.com/feichi/fcad_thrift/libs/go/fcmp_component_types"
)

type User struct {
	Id          int32 `xorm:"pk autoincr"`
	WxAppid     string
	AccessToken string
	ExpiresTime int32
	Status      enums.InfoStatus
	CreatedAt   int64 `xorm:"created"`
	UpdatedAt   int64 `xorm:"updated"`
}

func (this *User) TableName() string {
	return "test"
}

func (this *User) Thrift() (r *types.User) {
	r = &types.User{}
	util.Assign(this, r)
	return
}

type UserQuery struct {
	//ReferrerCode    string `xorm:"-"`
	User        `xorm:"-"`
	session     *xorm.Session `xorm:"-"`
	isNewRecord bool          `xorm:"-"`
}

func New() *UserQuery {
	o := &UserQuery{}
	o.isNewRecord = true
	return o
}

func (this *UserQuery) Orm() (r *xorm.EngineGroup) {
	return models.Orm
}

func (this *UserQuery) Session() (r *xorm.Session) {
	if this.session == nil {
		this.session = this.Orm().NewSession()
	}
	return this.session
}

func (this *UserQuery) Load(inp interface{}, excludes ...string) *UserQuery {
	util.Assign(inp, this, excludes...)
	return this
}

func (this *UserQuery) SQL(query string, params ...interface{}) (r *UserQuery) {
	this.Session().Sql(query, params...)
	return this
}

/* cond */
func (this *UserQuery) Where(cond builder.Cond) *UserQuery {
	this.Session().Where(cond)
	return this
}

func (this *UserQuery) And(cond builder.Cond) *UserQuery {
	this.Session().And(cond)
	return this
}

func (this *UserQuery) Or(cond builder.Cond) *UserQuery {
	this.Session().Or(cond)
	return this
}

/* misc */
func (this *UserQuery) Cols(cols ...string) *UserQuery {
	this.Session().Cols(cols...)
	return this
}

func (this *UserQuery) Select(str string) *UserQuery {
	this.Session().Select(str)
	return this
}

func (this *UserQuery) OrderBy(orders ...string) *UserQuery {
	this.Session().OrderBy(strings.Join(orders, ", "))
	return this
}

func (this *UserQuery) GroupBy(groups ...string) *UserQuery {
	this.Session().OrderBy(strings.Join(groups, ", "))
	return this
}

func (this *UserQuery) Limit(limit int, offset ...int) *UserQuery {
	this.Session().Limit(limit, offset...)
	return this
}

/* query */
func (this *UserQuery) Get() (r *User, err error) {
	r = &User{}
	if has, err := this.Session().Get(r); has == false || err != nil {
		return nil, err
	}
	this.isNewRecord = false
	return
}

func (this *UserQuery) Find() (r []*User, r2 map[int32]*User, err error) {
	err = this.Session().Find(&r)
	r2 = map[int32]*User{}
	for _, m := range r {
		r2[m.Id] = m
	}
	return
}

func (this *UserQuery) Count() (r int64, err error) {
	return this.Session().Count(&User{})
}

func (this *UserQuery) Exist() (r bool, err error) {
	return this.Session().Exist(&User{})
}
