package model

import (
	"fmt"
)

type Page struct {
	TotalNum    int
	TotalPage   int
	DataList    []interface{}
	CurrentPage int
}

func (this *Page) String() string {
	return fmt.Sprintf("totalNum:%v\n totalPage:%v\n currentPage:%v\n dataList:%v", this.TotalNum, this.TotalPage, this.CurrentPage, this.DataList)
}

type Query struct {
	Model     interface{}
	Condition map[string]interface{}
	Orderby   []string
}
