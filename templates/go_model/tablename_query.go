/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: __TABLE_NAME___query.go
# Description:
####################################################################### */

package __TABLE_NAME__

import (
	"errors"
	"fmt"

	"github.com/go-xorm/builder"
	"github.com/go-xorm/xorm"
	"gitlab.com/feichi/fcad_thrift/libs/go/enums"
)

// r, r2, err := query.Active().Find()
func (this *__TABLE_NAME_CAMEL__Query) Active() *__TABLE_NAME_CAMEL__Query {
	return this.And(builder.Eq{"status": 0})
}

/* query */
// r, err := query.GetById(2)
func (this *__TABLE_NAME_CAMEL__Query) GetById(id int32) (r *__TABLE_NAME_CAMEL__, err error) {
	if id == 0 {
		return
	}
	return this.Where(builder.Eq{"id": id}).Get()
}

// r, r2, err := query.GetByIds([]int32{1, 2})
func (this *__TABLE_NAME_CAMEL__Query) GetByIds(ids []int32) (r []*__TABLE_NAME_CAMEL__, r2 map[int32]*__TABLE_NAME_CAMEL__, err error) {
	if len(ids) == 0 {
		return
	}
	return this.Where(builder.Eq{"id": ids}).Find()
}

/* create */
// https://www.kancloud.cn/xormplus/xorm/167094
// &__TABLE_NAME_CAMEL__{}
func (this *__TABLE_NAME_CAMEL__Query) Insert(inp interface{}, session *xorm.Session) (err error) {
	this.Load(inp)
	this.Status = enums.InfoStatus_Normal

	if session != nil {
		this.session = session
	}
	_, err = this.Session().Insert(&this.__TABLE_NAME_CAMEL__)
	return
}

// []*__TABLE_NAME_CAMEL__{}
func (this *__TABLE_NAME_CAMEL__Query) InsertAll(inp []*__TABLE_NAME_CAMEL__, session *xorm.Session) (err error) {
	for _, v := range inp {
		v.Status = enums.InfoStatus_Normal
	}

	if session != nil {
		this.session = session
	}
	_, err = this.Session().Insert(&inp)
	return
}

/* update */
func (this *__TABLE_NAME_CAMEL__Query) Update(inp interface{}, session *xorm.Session) (err error) {
	if this.Id == 0 {
		return errors.New("id not set")
	}
	if this.isNewRecord == true {
		__TABLE_NAME__, err := this.GetById(this.Id)
		if err != nil {
			return fmt.Errorf("id#%d not found", this.Id)
		}
		this.Load(__TABLE_NAME__)
	}
	this.Load(inp)

	if session != nil {
		this.session = session
	}
	_, err = this.Where(builder.Eq{"id": this.Id}).Session().Update(&this.__TABLE_NAME_CAMEL__)
	return
}

/* delete */
func (this *__TABLE_NAME_CAMEL__Query) Delete(session *xorm.Session) (err error) {
	if session != nil {
		this.session = session
	}
	return this.Cols("status").Update(&struct {
		Status enums.InfoStatus
	}{
		Status: enums.InfoStatus_Deleted,
	}, nil)
}

/* update all */
/* delete all */

/* insert or update */
// notice: unsafe, unrealized
func (this *__TABLE_NAME_CAMEL__Query) InsertOrUpdate(inp interface{}, keys builder.Cond, session *xorm.Session) (err error) {
	return
}

/* get or insert */
// notice: unsafe, unrealized
func (this *__TABLE_NAME_CAMEL__Query) GetOrInsert(cond builder.Cond, inp interface{}, session *xorm.Session) (err error) {
	return
}
