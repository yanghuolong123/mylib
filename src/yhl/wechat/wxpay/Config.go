package wxpay

import (
	"fmt"
	"github.com/astaxie/beego"
)

var (
	AppId     string
	MchId     string
	Key       string
	AppSecret string
)

func init() {
	AppId = beego.AppConfig.String("")
	MchId = beego.AppConfig.String("")
	Key = beego.AppConfig.String("")
	AppSecret = beego.AppConfig.String("")
}
