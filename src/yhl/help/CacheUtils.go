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
	prefix := beego.AppConfig.String("redis.key.prefix")
	if host == "" || port == "" {
		return
	}

	conn := `{"conn":"` + host + `:` + port + `", "key":"` + prefix + `", "dbNum":"0"}`
	var err error
	Cache, err = cache.NewCache("redis", conn)
	Error(err)

	//Cache, _ = cache.NewCache("memory", `{"interval":0}`)
}
