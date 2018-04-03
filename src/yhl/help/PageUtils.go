package help

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"yhl/model"
)

func GetPageList(q model.Query, page, size int) (p model.Page) {
	m := q.Model
	t := reflect.TypeOf(m)
	v := reflect.ValueOf(m)
	//fmt.Println(t.Name())
	tableName := "tbl_" + strings.ToLower(t.Name())
	fmt.Println("======== t:", t)
	fmt.Println("========== tableName:", tableName)
	qs := orm.NewOrm().QueryTable(tableName)

	ty := v.Interface()
	fmt.Println("==== interface:", ty)
	//dataList := []ty{}
	//s := []m.(type){}
	//dataList := make([]m)
	//qs.All(&dataList)
	//fmt.Println(dataList)
	_ = qs

	return
}
