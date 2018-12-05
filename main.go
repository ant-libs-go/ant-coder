/* ######################################################################
# Author: (zhengfei@fcadx.cn)
# Created Time: 2018-11-14 12:50:43
# File Name: main.go
# Description:
####################################################################### */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"ant-coder/models"
	"ant-coder/models/task"
	//"github.com/go-xorm/builder"
)

// pass through when build project, go build -ldflags "main.__version__ 1.2.1" app
var (
	__version__ string
	pwd         = flag.String("d", "", "work directory")
	verbose     = flag.Bool("v", true, "enable verbose logging (default: false)")
	dsn         = flag.String("n", "", "database connect configured (format: user:pawd@host:port/name)")
	table       = flag.String("t", "", "table name")
)

func init() {
	flag.Parse()

	if len(*pwd) == 0 {
		*pwd, _ = os.Getwd()
	}
	os.Setenv("VERSION", __version__)
	os.Setenv("WORKDIR", *pwd)

	if len(*dsn) == 0 {
		log.Fatalf("you must specify `-n` options")
	}
	if len(*table) == 0 {
		log.Fatalf("you must specify `-t` options")
	}
	models.Init(*dsn, *verbose)
}

func main() {
	//r, err := task.NewDBTaskHandler().SQL("select * from test where id = ?", 2).Get()
	//r, err := task.NewDBTaskHandler().Where(builder.Eq{"id": 3}).Get()
	//r, err := task.NewDBTaskHandler().Where(builder.Eq{"id": 3}).And(builder.Eq{"status": 0}).Get()
	//r, err := task.NewDBTaskHandler().Where(builder.Eq{"id": 3, "status": 0}).Get()
	//r, err := task.NewDBTaskHandler().Cols("id", "wx_appid").Where(builder.Eq{"id": 3}).Get()
	//r, err := task.NewDBTaskHandler().Select("id, wx_appid").Where(builder.Eq{"id": 3}).Get()
	//r, err := task.NewDBTaskHandler().Where(builder.Eq{"id": 3}).OrderBy("id desc", "wx_appid asc").Get()
	//r, err := task.NewDBTaskHandler().Where(builder.Eq{"id": 3}).Limit(1, 1).Get()
	//r, err := task.NewDBTaskHandler().Count()
	//r, err := task.NewDBTaskHandler().Where(builder.Eq{"id": 2}).Exist()
	//r, r2, err := task.NewDBTaskHandler().Limit(2, 1).Active().Find()
	//r, r2, err := task.NewDBTaskHandler().Active().GetByIds([]int32{1, 2})
	//fmt.Println(r)
	//fmt.Println(r2)
	//fmt.Println(err)

	/*
		o := task.NewDBTaskHandler().Where(builder.Eq{"id": 3}).Get()
		o.AccessToken = "ddddaaad"
		err = o.Update(nil)

		o := task.NewDBTaskHandler()
		r, err := o.Where(builder.Eq{"id": 3}).Get()
		err = o.Delete(nil)
	*/

	/*
		search := task.NewDBTaskSearch()
		search.Id = 3
		r, r2, err := search.Search(nil)
	*/
}
