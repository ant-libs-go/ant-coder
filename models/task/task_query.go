/* ######################################################################
# Author: (zhengfei@fcadx.cn)
# Created Time: 2018-12-05 15:24:26
# File Name: User_query.go
# Description:
####################################################################### */

package User

import (
	"errors"

	"github.com/go-xorm/builder"
	"github.com/go-xorm/xorm"
	"gitlab.com/feichi/fcad_thrift/libs/go/enums"
	types "gitlab.com/feichi/fcad_thrift/libs/go/fcmp_component_types"
)

func (this *UserQuery) Active() *UserQuery {
	return this.And(builder.Eq{"status": 0})
}

/* query */
func (this *UserQuery) GetById(id int32) (r *User, err error) {
	if id == 0 {
		return
	}
	return this.Where(builder.Eq{"id": id}).Get()
}

func (this *UserQuery) GetByIds(ids []int32) (r []*User, r2 map[int32]*User, err error) {
	if len(ids) == 0 {
		return
	}
	return this.Where(builder.Eq{"id": ids}).Find()
}

/* create */
func (this *UserQuery) Create(inp *types.User, session *xorm.Session) (err error) {
	this.Load(inp)
	this.Status = enums.InfoStatus_Normal

	if session != nil {
		this.session = session
	}
	_, err = this.Session().Insert(&this.User)
	return
}

/* update */
func (this *UserQuery) Update(inp *types.User, session *xorm.Session) (err error) {
	if this.Id == 0 {
		return errors.New("id not set")
	}
	this.Load(inp)

	if session != nil {
		this.session = session
	}
	_, err = this.Where(builder.Eq{"id": this.Id}).Session().Update(&this.User)
	return
}

/* delete */
func (this *UserQuery) Delete(session *xorm.Session) (err error) {
	if this.Id == 0 {
		return errors.New("id not set")
	}
	this.Status = enums.InfoStatus_Deleted

	if session != nil {
		this.session = session
	}
	_, err = this.Where(builder.Eq{"id": this.Id}).Cols("status").Session().Update(&this.User)
	return
}
