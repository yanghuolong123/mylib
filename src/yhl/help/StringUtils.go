package help

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/axgle/mahonia"
	"strconv"
)

func IntToStr(num int) string {
	return strconv.Itoa(num)
}

func StrToInt(num string) int {
	i, err := strconv.Atoi(num)
	Error(err)

	return i
}

func ToStr(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func SubStr(str string, start, length int) string {
	return beego.Substr(str, start, length)
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
