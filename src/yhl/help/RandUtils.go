package help

import (
	"github.com/astaxie/beego/utils"
	"math/rand"
	"time"
)

func RandNum(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return 0
	}

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func RandStr(n int) string {
	return string(utils.RandomCreateBytes(n))
}
