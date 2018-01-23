package help

import (
	"github.com/astaxie/beego/utils"
	"math/rand"
)

func RandNum(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return 0
	}

	return rand.Intn(max-min) + min
}

func RandStr(n int) string {
	return string(utils.RandomCreateBytes(n))
}
