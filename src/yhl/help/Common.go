package help

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		val := v.Field(i).Interface()
		switch p := val.(type) {
		case int:
			data[strings.ToLower(t.Field(i).Name)] = fmt.Sprintf("%d", val)
		case *int:
			data[strings.ToLower(t.Field(i).Name)] = fmt.Sprintf("%d", val)
		case xml.Name:
		default:
			data[strings.ToLower(t.Field(i).Name)] = val
			_ = p
		}
	}

	return data
}

func Error(err error) {
	if err != nil {
		Log("error", err.Error())
	}
}

func GenOrderNo() string {
	return time.Now().Format(DatetimeNumFormat) + fmt.Sprintf("%d", RandNum(10000, 99999))
}

func GetAPPRootPath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	p, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	return filepath.Dir(p)
}

func HtmlToStr(html string) string {
	return beego.HTML2str(html)
}
