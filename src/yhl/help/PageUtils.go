package help

import (
	//	"fmt"
	"github.com/astaxie/beego/orm"
	//	"reflect"
	//	"strings"
	"math"
	"yhl/model"
)

/**
 page从0开始
**/
func GetPageList(q model.Query, page, size int) (p model.Page) {
	dataList := q.ReturnModelList
	qs := orm.NewOrm().QueryTable(q.Table)

	if len(q.Condition) > 0 {
		for k, v := range q.Condition {
			qs = qs.Filter(k, v)
		}
	}

	cnt, err := qs.Count()
	Error(err)

	//var maps []orm.Params
	if len(q.OrderBy) > 0 {
		qs = qs.OrderBy(q.OrderBy...)
	}
	if len(q.GroupBy) > 0 {
		qs = qs.GroupBy(q.GroupBy...)
	}
	//qs.Limit(size, page*size).Values(&maps)
	i, err := qs.Limit(size, page*size).All(dataList)
	Error(err)

	p.TotalCount = int(cnt)
	//if p.TotalCount%size == 0 {
	//	p.TotalPage = p.TotalCount / size
	//} else {
	//	p.TotalPage = p.TotalCount/size + 1
	//}
	p.TotalPage = int(math.Ceil(float64(p.TotalCount) / float64(size)))
	p.CurrentPage = page
	p.CurrentSize = int(i) //len(maps)
	p.DataList = dataList  //maps
	p.HasMore = true
	if p.TotalPage <= (p.CurrentPage + 1) {
		p.HasMore = false
	}

	return
}
