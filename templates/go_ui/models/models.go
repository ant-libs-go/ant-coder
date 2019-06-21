/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: models.go
# Description:
####################################################################### */

package models

import (
	"errors"
	"fmt"
	"time"

	"__PROJECT_NAME__/libs/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	Orm         *xorm.EngineGroup
	ErrNotFound = errors.New("record is not found")
)

func Init() {
	master, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&autocommit=true",
		config.Get().Db["master"].User,
		config.Get().Db["master"].Pawd,
		config.Get().Db["master"].Host,
		config.Get().Db["master"].Port,
		config.Get().Db["master"].Name))
	if err != nil {
		panic(fmt.Sprintf("master db connect errr: %s", err))
	}
	slave, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&autocommit=true",
		config.Get().Db["slave"].User,
		config.Get().Db["slave"].Pawd,
		config.Get().Db["slave"].Host,
		config.Get().Db["slave"].Port,
		config.Get().Db["slave"].Name))
	if err != nil {
		panic(fmt.Sprintf("slave db connect errr: %s", err))
	}
	Orm, err = xorm.NewEngineGroup(master, []*xorm.Engine{slave})
	Orm.SetConnMaxLifetime(3000)

	Orm.ShowSQL(config.Get().Basic.Debug)
	Orm.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
}

type InfoStatus int

const (
	InfoStatusNormal  InfoStatus = 0
	InfoStatusInvalid InfoStatus = 1
)
