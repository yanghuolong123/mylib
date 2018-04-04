package help

import (
	//	"fmt"
	"github.com/astaxie/beego/orm"
	//	"reflect"
	//	"strings"
	"yhl/model"
)

/**
 page从0开始
**/
func GetPageList(q model.Query, page, size int) (p model.Page) {
	//	m := q.Model
	//	t := reflect.TypeOf(m)
	//v := reflect.ValueOf(m)

	//tableName := "tbl_" + strings.ToLower(t.Name())
	qs := orm.NewOrm().QueryTable(q.Table)

	if len(q.Condition) > 0 {
		for k, v := range q.Condition {
			qs = qs.Filter(k, v)
		}
	}

	cnt, err := qs.Count()
	Error(err)

	var maps []orm.Params
	if len(q.OrderBy) > 0 {
		qs = qs.OrderBy(q.OrderBy...)
	}
	if len(q.GroupBy) > 0 {
		qs = qs.GroupBy(q.GroupBy...)
	}
	qs.Limit(size, page*size).Values(&maps)

	p.TotalCount = int(cnt)
	if p.TotalCount%size == 0 {
		p.TotalPage = p.TotalCount / size
	} else {
		p.TotalPage = p.TotalCount/size + 1
	}
	p.CurrentPage = page
	p.CurrentSize = len(maps)
	p.DataList = maps
	p.HasMore = true
	if p.TotalPage <= (p.CurrentPage + 1) {
		p.HasMore = false
	}

	return
}
