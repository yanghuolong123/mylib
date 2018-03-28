package help

import (
	"fmt"
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
