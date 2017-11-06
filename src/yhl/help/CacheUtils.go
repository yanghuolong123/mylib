package help

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var (
	Cache cache.Cache
)

func init() {
	host := beego.AppConfig.String("redis.host")
	port := beego.AppConfig.String("redis.port")
	conn := `{"conn":"` + host + `:` + port + `", "dbNum":"0"}`
	Cache, _ = cache.NewCache("redis", conn)

	//Cache, _ = cache.NewCache("memory", `{"interval":0}`)
}
