package help

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
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
	m["jsonrpc"] = "2.0"
	m["success"] = true
	if code != 0 {
		m["success"] = false
	}
	m["code"] = code
	m["msg"] = msg
	if data != nil {
		m["data"] = data
	}

	this.Data["json"] = m
	this.ServeJSON()
	this.StopRun()
}
