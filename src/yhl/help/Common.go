package help

import (
	"fmt"
	"reflect"
	"strings"
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
		default:
			data[strings.ToLower(t.Field(i).Name)] = val
			_ = p
		}
	}

	return data
}
