package help

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

var (
	ClientIp   string
	ClientPort string
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	this.Controller.Init(ctx, controllerName, actionName, app)
	s := strings.Split(ctx.Request.RemoteAddr, ":")
	ClientIp = s[0]
	ClientPort = s[1]
}

func (this *BaseController) SendRes(code int, msg string, data interface{}) {
	m := make(map[string]interface{})
	m["code"] = code
	m["msg"] = msg
	if data != nil {
		m["data"] = data
	}

	this.Data["json"] = m
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) SendResJsonp(code int, msg string, data interface{}) {
	m := make(map[string]interface{})
	m["code"] = code
	m["msg"] = msg
	if data != nil {
		m["data"] = data
	}

	this.Data["jsonp"] = m
	this.ServeJSONP()
	this.StopRun()
}

func (this *BaseController) SendXml(data interface{}) {
	this.Data["xml"] = data
	this.ServeXML()
	this.StopRun()
}
