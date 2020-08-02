/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2019-02-20 14:28:16
# File Name: go_model_coder.go
# Description:
####################################################################### */

package coder

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"ant-coder/models"

	"github.com/ant-libs-go/util"
	"xorm.io/xorm/schemas"
)

type GoModelCoder struct {
	name string
	opts []*Option
	tpls []*Tpl
}

func NewGoModelCoder() *GoModelCoder {
	o := &GoModelCoder{}
	o.name = "go_model"
	o.opts = []*Option{
		&Option{
			Name:  "dsn",
			Usage: "Input database connect configured (format: user:pawd@host:port/name):",
			Cache: true},
		&Option{
			Name:  "table",
			Usage: "Input table name:"},
		&Option{
			Name:  "author",
			Usage: "Input author name:",
			Cache: true},
		&Option{
			Name:  "project_name",
			Usage: "Input project name:",
			Cache: true},
	}
	o.tpls = []*Tpl{
		&Tpl{Src: "templates/go_model/tablename.go", Dst: "tablename.go"},
		&Tpl{Src: "templates/go_model/tablename_query.go", Dst: "tablename_query.go"},
		&Tpl{Src: "templates/go_model/tablename_search.go", Dst: "tablename_search.go"},
	}
	return o
}

func (this *GoModelCoder) GetName() (r string) {
	return this.name
}

func (this *GoModelCoder) GetOptions() (r []*Option) {
	return this.opts
}

func (this *GoModelCoder) GetTpls() (r []*Tpl) {
	return this.tpls
}

func (this *GoModelCoder) Init() (err error) {
	err = models.Init(GetOptionValueByKey(this.opts, "dsn"))
	if err != nil {
		fmt.Println("....... model\t[no]")
		return
	}
	fmt.Println("....... model\t[ok]")
	return
}

func (this *GoModelCoder) GetBaseDirName() (r string) {
	return GetOptionValueByKey(this.opts, "table")
}

func (this *GoModelCoder) GetMacros() (fileNameMacros []*Macro, fileContMacros []*Macro, err error) {
	fileNameMacros = []*Macro{
		&Macro{Key: "tablename", Val: GetOptionValueByKey(this.opts, "table")},
	}
	fileContMacros = []*Macro{}
	var macroVal string
	if macroVal, err = this.getBusinessFuncs(); err != nil {
		return
	}
	fileContMacros = append(fileContMacros, &Macro{Key: "__BUSINESS_FUNCTIONS__", Val: macroVal})

	fileContMacros = append(fileContMacros, &Macro{Key: "__TABLE_NAME__", Val: GetOptionValueByKey(this.opts, "table")})
	fileContMacros = append(fileContMacros, &Macro{Key: "__TABLE_NAME_CAMEL__", Val: util.CamelString(GetOptionValueByKey(this.opts, "table"))})
	fileContMacros = append(fileContMacros, &Macro{Key: "__AUTHOR__", Val: GetOptionValueByKey(this.opts, "author")})
	fileContMacros = append(fileContMacros, &Macro{Key: "__CREATE_DATETIME__", Val: time.Now().Format("2006-01-02 15:04:05")})
	fileContMacros = append(fileContMacros, &Macro{Key: "__PROJECT_NAME__", Val: GetOptionValueByKey(this.opts, "project_name")})
	fileContMacros = append(fileContMacros, &Macro{Key: "__PROJECT_NAME_CAMEL__", Val: util.CamelString(GetOptionValueByKey(this.opts, "project_name"))})

	if macroVal, err = this.getTableMetaStruct(); err != nil {
		return
	}
	fileContMacros = append(fileContMacros, &Macro{Key: "__TABLE_META_STRUCT__", Val: macroVal})

	return
}

// &{Name:tag Type:<nil> columnsSeq:[id name type is_must is_good created_at updated_at deleted_at] columnsMap:map[created_at:[0xc000132ea0] deleted_at:[0xc000133040] id:[0xc0001329c0] is_good:[0xc000132dd0] is_must:[0xc000132d00] name:[0xc000132b60] type:[0xc000132c30] updated_at:[0xc000132f70]] columns:[0xc0001329c0 0xc000132b60 0xc000132c30 0xc000132d00 0xc000132dd0 0xc000132ea0 0xc000132f70 0xc000133040] Indexes:map[idx_is:0xc000089a80 uniq_idx_name_type:0xc000089a40] PrimaryKeys:[id] AutoIncrement:id Created:map[] Updated: Deleted: Version: StoreEngine:InnoDB Charset: Comment:标签表}
func (this *GoModelCoder) getBusinessFuncs() (r string, err error) {
	var table *schemas.Table
	tables, _ := models.Orm.DBMetas()
	for _, tab := range tables {
		if tab.Name != GetOptionValueByKey(this.opts, "table") {
			continue
		}
		table = tab
	}
	if table == nil {
		err = fmt.Errorf("table:%s not found", GetOptionValueByKey(this.opts, "table"))
		return
	}

	type index struct {
		Name string
		Type string
		Cols []string
	}
	idxs := []*index{}
	if len(table.PrimaryKeys) > 0 {
		idxs = append(idxs, &index{Name: "pk", Type: "pk", Cols: table.PrimaryKeys})
	}
	for _, idx := range table.Indexes {
		switch idx.Type {
		case 1:
			idxs = append(idxs, &index{Name: idx.Name, Type: "idx", Cols: idx.Cols})
		case 2:
			idxs = append(idxs, &index{Name: idx.Name, Type: "uk", Cols: idx.Cols})
		}
	}

	funcOncePk := `func (this *__TABLE_NAME_CAMEL__Query) GetBy__ONCE_FUNC_NAME__(__ONCE_FUNC_PARAMS__) (r *__TABLE_NAME_CAMEL__, err error) {
	return this.Where(builder.Eq{__ONCE_FUNC_BUILDER__}).Get()
}`
	funcMultiPk := `func (this *__TABLE_NAME_CAMEL__Query) GetBy__MULTI_FUNC_NAME__(__MULTI_FUNC_PARAMS__) (r []*__TABLE_NAME_CAMEL__, r2 map[int32]*__TABLE_NAME_CAMEL__, err error) {
	return this.Where(builder.Eq{__MULTI_FUNC_BUILDER__}).Find()
}`
	funcOnceUk := `func (this *__TABLE_NAME_CAMEL__Query) GetBy__ONCE_FUNC_NAME__(__ONCE_FUNC_PARAMS__) (r *__TABLE_NAME_CAMEL__, err error) {
	return this.Where(builder.Eq{__ONCE_FUNC_BUILDER__}).Get()
}`
	funcMultiUk := `func (this *__TABLE_NAME_CAMEL__Query) GetBy__MULTI_FUNC_NAME__(__MULTI_FUNC_PARAMS__) (r []*__TABLE_NAME_CAMEL__, r2 map[int32]*__TABLE_NAME_CAMEL__, err error) {
	return this.Where(builder.Eq{__MULTI_FUNC_BUILDER__}).Find()
}`
	funcMultiIdx := `func (this *__TABLE_NAME_CAMEL__Query) FindBy__MULTI_FUNC_NAME__(__MULTI_FUNC_PARAMS__) (r []*__TABLE_NAME_CAMEL__, r2 map[int32]*__TABLE_NAME_CAMEL__, err error) {
	return this.Where(builder.Eq{__MULTI_FUNC_BUILDER__}).Find()
}`

	funcs := []string{}
	for _, idx := range idxs {
		onceFuncNames := []string{}
		onceFuncParams := []string{}
		onceFuncBuilders := []string{}
		multiFuncNames := []string{}
		multiFuncParams := []string{}
		multiFuncBuilders := []string{}
		for _, col := range idx.Cols {
			colVarName := util.CamelString(col)
			if colVarName == "type" {
				colVarName = "tye"
			}
			colType, _ := this.getColType(table.GetColumn(col))
			onceFuncNames = append(onceFuncNames, col)
			onceFuncParams = append(onceFuncParams, fmt.Sprintf("%s %s", colVarName, colType))
			onceFuncBuilders = append(onceFuncBuilders, fmt.Sprintf(`"%s": %s`, col, colVarName))
			multiFuncNames = append(multiFuncNames, fmt.Sprintf("%ss", col))
			multiFuncParams = append(multiFuncParams, fmt.Sprintf("%ss []%s", colVarName, colType))
			multiFuncBuilders = append(multiFuncBuilders, fmt.Sprintf(`"%s": %ss`, col, colVarName))
		}
		macros := []*Macro{
			&Macro{Key: "__ONCE_FUNC_NAME__", Val: util.CamelString(strings.Join(onceFuncNames, "_and_"))},
			&Macro{Key: "__ONCE_FUNC_PARAMS__", Val: strings.Join(onceFuncParams, ", ")},
			&Macro{Key: "__ONCE_FUNC_BUILDER__", Val: strings.Join(onceFuncBuilders, ", ")},
			&Macro{Key: "__MULTI_FUNC_NAME__", Val: util.CamelString(strings.Join(multiFuncNames, "_and_"))},
			&Macro{Key: "__MULTI_FUNC_PARAMS__", Val: strings.Join(multiFuncParams, ", ")},
			&Macro{Key: "__MULTI_FUNC_BUILDER__", Val: strings.Join(multiFuncBuilders, ", ")},
		}

		switch idx.Type {
		case "pk":
			funcs = append(funcs, MacroReplace(funcOncePk, macros))
			if len(idx.Cols) == 1 {
				funcs = append(funcs, MacroReplace(funcMultiPk, macros))
			}
		case "uk":
			funcs = append(funcs, MacroReplace(funcOnceUk, macros))
			if len(idx.Cols) == 1 {
				funcs = append(funcs, MacroReplace(funcMultiUk, macros))
			}
		case "idx":
			funcs = append(funcs, MacroReplace(funcMultiIdx, macros))
		}
	}
	r = strings.Join(funcs, "\n\n")
	return
}

func (this *GoModelCoder) getTableMetaStruct() (r string, err error) {
	fields := []string{}
	var table *schemas.Table

	tables, _ := models.Orm.DBMetas()
	for _, tab := range tables {
		if tab.Name != GetOptionValueByKey(this.opts, "table") {
			continue
		}
		table = tab
	}
	if table == nil {
		err = fmt.Errorf("table:%s not found", GetOptionValueByKey(this.opts, "table"))
		return
	}
	for _, col := range table.Columns() {
		var tn, tp, tg string
		tn = util.CamelString(col.Name)
		if tp, err = this.getColType(col); err != nil {
			return
		}
		if tg, err = this.getColTag(col, table.Indexes); err != nil {
			return
		}
		fields = append(fields, fmt.Sprintf("\t%s\t\t%s\t`%s`", tn, tp, tg))
	}
	r = strings.Join(fields, "\n")
	return
}

func (this *GoModelCoder) getColType(col *schemas.Column) (r string, err error) {
	switch strings.ToUpper(col.SQLType.Name) {
	case "TINYINT", "SMALLINT", "MEDIUMINT", "INT", "SERIAL", "BIT":
		r = reflect.TypeOf(int32(1)).String()
	case "BIGINT", "BIGSERIAL":
		r = reflect.TypeOf(int64(1)).String()
	case "FLOAT", "REAL":
		r = reflect.TypeOf(float32(1)).String()
	case "DOUBLE":
		r = reflect.TypeOf(float64(1)).String()
	case "CHAR", "VARCHAR", "TINYTEXT", "MEDIUMTEXT", "TEXT", "LONGTEXT", "ENUM", "SET", "DECIMAL", "NUMERIC":
		r = reflect.TypeOf("").String()
	case "TINYBLOB", "MEDIUMBLOB", "BLOB", "LONGBLOB", "BINARY", "VARBINARY":
		r = reflect.TypeOf([]byte{}).String()
	case "BOOL":
		r = reflect.TypeOf(true).String()
	case "DATETIME", "DATE", "TIME", "TIMESTAMP", "YEAR":
		r = reflect.TypeOf(time.Now()).String()
	default:
		err = fmt.Errorf("Column type#%s not support", col.SQLType.Name)
	}
	if r == "[]uint8" {
		r = "[]byte"
	}
	return
}

func (this *GoModelCoder) getColTag(col *schemas.Column, indexes map[string]*schemas.Index) (r string, err error) {
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
	if col.Name == "deleted_at" {
		tags = append(tags, "deleted")
	}
	if col.Nullable == false {
		tags = append(tags, "not null")
	}
	for name := range col.Indexes {
		index := indexes[name]
		switch index.Type {
		case schemas.UniqueType:
			tags = append(tags, fmt.Sprintf("unique(%s)", index.Name))
		case schemas.IndexType:
			tags = append(tags, fmt.Sprintf("index(%s)", index.Name))
		}
	}
	if len(col.Default) > 0 {
		tags = append(tags, fmt.Sprintf("default(%s)", col.Default))
	}
	if len(col.Comment) > 0 {
		tags = append(tags, fmt.Sprintf("comment('%s')", col.Comment))
	}
	r = fmt.Sprintf(`xorm:"%s"`, strings.Join(tags, " "))
	return
}
