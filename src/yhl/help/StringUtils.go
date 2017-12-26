package help

import (
	"strconv"
)

func IntToString(num int) string {
	return strconv.Itoa(num)
}

func StringToInt(num string) int {
	i, _ := strconv.Atoi(num)
	return i
}
