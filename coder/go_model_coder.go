/* ######################################################################
# Author: (zhengfei@fcadx.cn)
# Created Time: 2019-02-20 14:28:16
# File Name: go_model_coder.go
# Description:
####################################################################### */

package coder

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"ant-coder/models"

	"github.com/ant-libs-go/util"
	"github.com/go-xorm/core"
)

type GoModelCoder struct {
	opts  []*Option
	tpls  []*Tpl
	table string
}

func NewGoModelCoder() *GoModelCoder {
	o := &GoModelCoder{}
	o.opts = []*Option{
		&Option{
			Name:  "dsn",
			Usage: "Input database connect configured (format: user:pawd@host:port/name):"},
		&Option{
			Name:  "table",
			Usage: "Input table name:"},
		&Option{
			Name:  "author",
			Usage: "Input author name:"},
		&Option{
			Name:  "project_name",
			Usage: "Input project name:"},
	}
	o.tpls = []*Tpl{
		&Tpl{Src: "templates/go_model/tablename.go"},
		&Tpl{Src: "templates/go_model/tablename_query.go"},
		&Tpl{Src: "templates/go_model/tablename_search.go"},
	}
	return o
}

func (this *GoModelCoder) GetOptions() (r []*Option) {
	return this.opts
}

func (this *GoModelCoder) Init() (err error) {
	this.table = GetOptionValueByKey(this.opts, "table")
	err = models.Init(GetOptionValueByKey(this.opts, "dsn"))
	if err != nil {
		fmt.Println("....... model\t[no]")
		return
	}
	fmt.Println("....... model\t[ok]")
	return
}

func (this *GoModelCoder) Generate() (err error) {
	// mkdir dir
	dir := fmt.Sprintf("%s/%s", os.Getenv("WORKDIR"), this.table)
	if err = Mkdir(dir); err != nil {
		fmt.Println(fmt.Sprintf("....... directory#%s mkdir\t[no]", dir))
		return
	}
	fmt.Println("....... directory mkdir\t[ok]")

	// render tpl
	fileNameMacros, fileContMacros, err := this.getMacros()
	if err != nil {
		return err
	}
	for _, tpl := range this.tpls {
		tpl.Dst = fmt.Sprintf("%s/%s", dir, MacroReplace(strings.TrimPrefix(tpl.Src, "templates/go_model/"), fileNameMacros))

		if err = Mkdir(path.Dir(tpl.Dst)); err != nil {
			fmt.Println(fmt.Sprintf("....... directory#%s mkdir\t[no]", path.Dir(tpl.Dst)))
			return
		}

		if err = RenderTpl(tpl, fileContMacros); err != nil {
			return
		}
	}
	fmt.Println("....... render template\t[ok]")

	return
}

func (this *GoModelCoder) getMacros() (fileNameMacros []*Macro, fileContMacros []*Macro, err error) {
	fileNameMacros = []*Macro{
		&Macro{Key: "tablename", Val: this.table},
	}
	fileContMacros = []*Macro{
		&Macro{Key: "__TABLE_NAME__", Val: this.table},
		&Macro{Key: "__TABLE_NAME_CAMEL__", Val: util.CamelString(this.table)},
		&Macro{Key: "__AUTHOR__", Val: GetOptionValueByKey(this.opts, "author")},
		&Macro{Key: "__CREATE_DATETIME__", Val: time.Now().Format("2006-01-02 15:04:05")},
		&Macro{Key: "__PROJECT_NAME__", Val: GetOptionValueByKey(this.opts, "project_name")},
		&Macro{Key: "__PROJECT_NAME_CAMEL__", Val: util.CamelString(GetOptionValueByKey(this.opts, "project_name"))},
	}

	var desc string
	desc, err = this.getTableDesc()
	if err != nil {
		return
	}
	fileContMacros = append(fileContMacros, &Macro{Key: "__TABLE_META_STRUCT__", Val: desc})

	return
}

func (this *GoModelCoder) getTableDesc() (r string, err error) {
	fields := []string{}
	tables, _ := models.Orm.DBMetas()
	for _, table := range tables {
		if table.Name != this.table {
			continue
		}
		for _, col := range table.Columns() {
			var tn, tp, tg string
			tn = util.CamelString(col.Name)
			if tp, err = this.getColType(col); err != nil {
				return
			}
			if tg, err = this.getColTag(col); err != nil {
				return
			}
			fields = append(fields, fmt.Sprintf("\t%s\t\t%s\t`%s`", tn, tp, tg))
		}
	}
	r = strings.Join(fields, "\n")
	return
}

func (this *GoModelCoder) getColType(col *core.Column) (r string, err error) {
	if col.Name == "status" {
		r = "models.InfoStatus"
	}
	if len(r) > 0 {
		return
	}

	switch col.SQLType.Name {
	case "TINYINT", "INT":
		r = "int32"
	case "BIGINT":
		r = "int64"
	case "CHAR", "VARCHAR", "TEXT":
		r = "string"
	default:
		err = fmt.Errorf("Column type#%s not support", col.SQLType.Name)
	}
	return
}

func (this *GoModelCoder) getColTag(col *core.Column) (r string, err error) {
	tags := []string{}
	if col.IsPrimaryKey == true {
		tags = append(tags, "pk")
	}
	if col.IsAutoIncrement == true {
		tags = append(tags, "autoincr")
	}
	if col.Name == "created_at" {
		tags = append(tags, "created")
	}
	if col.Name == "updated_at" {
		tags = append(tags, "updated")
	}

	r = fmt.Sprintf(`xorm:"%s"`, strings.Join(tags, " "))
	return
}

/*
	table = &{
		Name:ad
		Type:<nil>
		columnsSeq:[id media_id icon title desc button_action play_sleep_time earn type status created_at updated_at]
		columnsMap:map[
			title:[0xc000144340]
			button_action:[0xc0001444e0]
			earn:[0xc000144680]
			id:[0xc000144000]
			icon:[0xc000144270]
			desc:[0xc000144410]
			play_sleep_time:[0xc0001445b0]
			type:[0xc000144750]
			status:[0xc000144820]
			created_at:[0xc0001448f0]
			updated_at:[0xc0001449c0]
			media_id:[0xc0001441a0]
		]
		columns:[
			0xc000144000
			0xc0001441a0
			0xc000144270
			0xc000144340
			0xc000144410
			0xc0001444e0
			0xc0001445b0
			0xc000144680
			0xc000144750
			0xc000144820
			0xc0001448f0
			0xc0001449c0
		]
		PrimaryKeys:[id]
		AutoIncrement:id
		StoreEngine:InnoDB
		Indexes:map[]
		Created:map[]
		Updated:
		Deleted:
		Version:
		Cacher:<nil>
		Charset:
		Comment:
	}

	column = &{
		Name:id
		TableName:
		FieldName:
		SQLType:{
			Name:INT
			DefaultLength:11
			DefaultLength2:0
		}
		IsJSON:false
		Length:11
		Length2:0
		Nullable:false
		Default:
		Indexes:map[]
		IsPrimaryKey:true
		IsAutoIncrement:true
		MapType:0
		IsCreated:false
		IsUpdated:false
		IsDeleted:false
		IsCascade:false
		IsVersion:false
		DefaultIsEmpty:false
		EnumOptions:map[]
		SetOptions:map[]
		DisableTimeZone:false
		TimeZone:UTC
		Comment:
	}
*/
