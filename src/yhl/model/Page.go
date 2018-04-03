package model

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Page struct {
	TotalCount  int
	TotalPage   int
	CurrentPage int
	CurrentSize int
	DataList    []orm.Params
}

func (this *Page) String() string {
	return fmt.Sprintf("totalCount:%v\n totalPage:%v\n currentPage:%v\n currentSize:%v\n dataList:%v", this.TotalCount, this.TotalPage, this.CurrentPage, this.CurrentSize, this.DataList)
}

type Query struct {
	Model     interface{}
	Condition map[string]interface{}
	Orderby   []string
}
