package model

import (
	"fmt"
	//	"github.com/astaxie/beego/orm"
)

type Page struct {
	TotalCount  int
	TotalPage   int
	CurrentPage int
	CurrentSize int
	DataList    interface{} //[]orm.Params
	HasMore     bool
}

func (this *Page) String() string {
	return fmt.Sprintf("totalCount:%v\n totalPage:%v\n currentPage:%v\n currentSize:%v\n hasMore:%v\n dataList:%v", this.TotalCount, this.TotalPage, this.CurrentPage, this.CurrentSize, this.HasMore, this.DataList)
}

type Query struct {
	Table           string
	Condition       map[string]interface{}
	OrderBy         []string
	GroupBy         []string
	ReturnModelList interface{}
}
