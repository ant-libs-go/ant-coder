/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: __TABLE_NAME__.go
# Description:
####################################################################### */

/**
 * Define constant:
 *
 * ErrNotFound = errors.New("record is not found")
 *
 * type SortType int64
 *
 * const (
 * 	 SortTypeAsc  SortType = 0
 * 	 SortTypeDesc SortType = 1
 * )
 *
 * type SortParams struct {
 * 	 Field string
 * 	 Type  SortType
 * }
 *
 * type Pager struct {
 * 	 Offset int
 * 	 Limit  int
 * }
 *
 * func ParseSortParams(sorts []*SortParams) (r []string) {
 * 	 r = []string{}
 * 	 for _, sort := range sorts {
 *     if sort.Type == SortTypeAsc {
 * 	     r = append(r, fmt.Sprintf("%s ASC", sort.Field))
 * 	   }
 * 	   if sort.Type == SortTypeDesc {
 * 	     r = append(r, fmt.Sprintf("%s DESC", sort.Field))
 *     }
 * 	 }
 *   return
 * }
 */

package __TABLE_NAME__

import (
	"fmt"

	"__PROJECT_NAME__/models"

	"github.com/ant-libs-go/util"
	"xorm.io/xorm"
)

type __TABLE_NAME_CAMEL__ struct {
	isNewRecord bool          `xorm:"-"`
	isAutoClose bool          `xorm:"-"`
	session     *xorm.Session `xorm:"-"`
	__TABLE_META_STRUCT__
}

func New(session *xorm.Session) *__TABLE_NAME_CAMEL__ {
	o := &__TABLE_NAME_CAMEL__{isNewRecord: true, session: session}
	if session != nil {
		o.isAutoClose = true
	}
	return o
}

func (this *__TABLE_NAME_CAMEL__) Orm() (r *xorm.EngineGroup) {
	return models.Orm
}

func (this *__TABLE_NAME_CAMEL__) Session() (r *xorm.Session) {
	if this.session == nil {
		this.session = this.Orm().NewSession().Table(this)
	}
	return this.session
}

func (this *__TABLE_NAME_CAMEL__) TableName() string {
	return "__TABLE_NAME__"
}

func (this *__TABLE_NAME_CAMEL__) Load(inp interface{}, excludes ...string) *__TABLE_NAME_CAMEL__ {
	if err := util.Assign(inp, this, excludes...); err != nil {
		fmt.Println(err)
		return nil
	}
	if this.Id != 0 {
		this.isNewRecord = false
	}
	return this
}

func (this *__TABLE_NAME_CAMEL__) BeforeValidate()                       {}
func (this *__TABLE_NAME_CAMEL__) AfterValidate()                        {}
func (this *__TABLE_NAME_CAMEL__) BeforeSave()                           {}
func (this *__TABLE_NAME_CAMEL__) AfterSave()                            {}
func (this *__TABLE_NAME_CAMEL__) BeforeInsert()                         {}
func (this *__TABLE_NAME_CAMEL__) AfterInsert()                          {}
func (this *__TABLE_NAME_CAMEL__) BeforeUpdate()                         {}
func (this *__TABLE_NAME_CAMEL__) AfterUpdate()                          {}
func (this *__TABLE_NAME_CAMEL__) BeforeDelete()                         {}
func (this *__TABLE_NAME_CAMEL__) AfterDelete()                          {}
func (this *__TABLE_NAME_CAMEL__) BeforeSet(name string, cell xorm.Cell) {}
func (this *__TABLE_NAME_CAMEL__) AfterSet(name string, cell xorm.Cell)  { this.isNewRecord = false }

func (this *__TABLE_NAME_CAMEL__) Validate() (err error) {
	this.BeforeValidate()
	defer this.AfterValidate()

	return
}

func (this *__TABLE_NAME_CAMEL__) Save(runValidation bool) (err error) {
	this.BeforeSave()
	defer this.AfterSave()

	if this.isNewRecord == true {
		return this.Insert(runValidation)
	}
	return this.Update(runValidation)
}

func (this *__TABLE_NAME_CAMEL__) Insert(runValidation bool) (err error) {
	if this.isAutoClose == true {
		defer this.Session().Close()
	}
	if this.isNewRecord == false {
		return fmt.Errorf("record already exists")
	}
	if err = this.Validate(); err != nil {
		return
	}
	var id int64
	id, err = this.Session().Insert(this)
	this.Id = int32(id)
	return
}

func (this *__TABLE_NAME_CAMEL__) Update(runValidation bool, columns ...string) (err error) {
	if this.isAutoClose == true {
		defer this.Session().Close()
	}
	if this.isNewRecord == true {
		return fmt.Errorf("record don't exist")
	}
	if err = this.Validate(); err != nil {
		return
	}
	sess := this.Session().ID(this.Id)
	if len(columns) > 0 {
		sess = sess.Cols(columns...)
	}
	var affected int64
	if affected, err = sess.Update(this); err != nil {
		return
	}
	if affected == 0 {
		return fmt.Errorf("record is not found")
	}
	return
}

func (this *__TABLE_NAME_CAMEL__) Delete() (err error) {
	if this.isAutoClose == true {
		defer this.Session().Close()
	}
	if this.isNewRecord == true {
		return fmt.Errorf("record don't exist")
	}
	var affected int64
	if affected, err = this.Session().ID(this.Id).Delete(&__TABLE_NAME_CAMEL__{}); err != nil {
		return
	}
	if affected == 0 {
		return fmt.Errorf("record is not found")
	}
	return
}

func (this *__TABLE_NAME_CAMEL__) Query() *__TABLE_NAME_CAMEL__Query {
	return &__TABLE_NAME_CAMEL__Query{isAutoClose: this.isAutoClose, session: this.Session()}
}

func (this *__TABLE_NAME_CAMEL__) Search() *__TABLE_NAME_CAMEL__Search {
	return &__TABLE_NAME_CAMEL__Search{query: this.Query(), Pager: models.Pager{Limit: 100}}
}

// vim: set noexpandtab ts=4 sts=4 sw=4 :
