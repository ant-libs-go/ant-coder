/* ######################################################################
# Author: (zhengfei@fcadx.cn)
# Created Time: 2018-12-05 15:24:26
# File Name: task_query.go
# Description:
####################################################################### */

package task

import "github.com/go-xorm/builder"

func (this *DBTaskQuery) Active() *DBTaskQuery {
	return this.And(builder.Eq{"status": 0})
}

func (this *DBTaskQuery) GetByIds(ids []int32) (r []*DBTask, r2 map[int32]*DBTask, err error) {
	return this.And(builder.Eq{"id": ids}).Find()
}
