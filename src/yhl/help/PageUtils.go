package help

import (
	//	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"yhl/model"
)

func GetPageList(q model.Query, page, size int) (p model.Page) {
	m := q.Model
	t := reflect.TypeOf(m)
	//v := reflect.ValueOf(m)

	tableName := "tbl_" + strings.ToLower(t.Name())
	qs := orm.NewOrm().QueryTable(tableName)

	cnt, err := qs.Count()
	Error(err)

	var maps []orm.Params
	qs = qs.OrderBy("-create_time").Limit(size, page*size)
	qs.Values(&maps)

	//	fmt.Println("======== maps:", maps)

	p.TotalCount = int(cnt)
	p.TotalPage = int(p.TotalCount/size) + 1
	p.CurrentPage = page
	p.CurrentSize = len(maps)
	p.DataList = maps

	return
}
