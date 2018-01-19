package search

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

var (
	host string
	port string
)

func init() {
	host = "http://" + beego.AppConfig.String("es.search.host")
	port = beego.AppConfig.String("es.search.port")
}

func Get(uri string) (data map[string]interface{}) {
	url := host + ":" + port + uri
	req := httplib.Get(url)
	req.ToJSON(&data)

	return
}

func Put(uri string, m map[string]interface{}) (data map[string]interface{}) {
	url := host + ":" + port + uri
	req := httplib.Put(url)
	req.JSONBody(m)
	req.ToJSON(&data)

	return
}

func Post(uri string, m map[string]interface{}) (data map[string]interface{}) {
	url := host + ":" + port + uri
	req := httplib.Post(url)
	req.JSONBody(m)
	req.ToJSON(&data)

	return
}

func Delete(uri string) (data map[string]interface{}) {
	url := host + ":" + port + uri
	req := httplib.Delete(url)
	req.ToJSON(&data)

	return
}
