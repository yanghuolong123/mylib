package wechat

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	"time"
	"yhl/help"
)

type Share struct {
	Title string
	Desc  string
	Link  string
	Img   string
}

func GetSignPackage(c *context.Context) map[string]interface{} {
	jsapiTicket := GetJsApiTickey()

	urlstr := c.Input.Site() + c.Input.URI()

	timeStamp := time.Now().Unix()
	nonceStr := help.RandStr(32)

	str := "jsapi_ticket=" + jsapiTicket + "&noncestr=" + nonceStr + "&timestamp=" + help.ToStr(timeStamp) + "&url=" + urlstr

	m := map[string]interface{}{}
	m["appId"] = Appid
	m["nonceStr"] = nonceStr
	m["timestamp"] = timeStamp
	m["url"] = urlstr
	m["signature"] = help.Sha1(str)
	m["rawString"] = str

	return m
}

func GetJsApiTickey() (token string) {
	cache := help.Cache
	t := cache.Get("jsapi_ticket_" + Appid)
	if t != nil {
		token = string(t.([]uint8))
		return
	}

	accessToken := GetAccessToken()
	api := ApiUrl + "/cgi-bin/ticket/getticket?type=jsapi&access_token=" + accessToken

	b := httplib.Get(api)
	data := make(map[string]interface{})
	b.ToJSON(&data)
	fmt.Println(data)

	if v, ok := data["ticket"]; ok {
		token = v.(string)
		ttl := time.Duration(7200)
		cache.Put("jsapi_ticket_"+Appid, token, ttl*time.Second)
	}

	return
}
