package wxpay

import (
	"github.com/astaxie/beego"
)

var (
	AppId     string
	MchId     string
	Key       string
	AppSecret string
)

func init() {
	AppId = beego.AppConfig.String("wechat.appid")
	MchId = beego.AppConfig.String("wechat.wxpay.mchid")
	Key = beego.AppConfig.String("wechat.wxpay.key")
	AppSecret = beego.AppConfig.String("wechat.secret")
}
