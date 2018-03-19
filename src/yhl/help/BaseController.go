package help

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
	"time"
	"yhl/model"
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
	ip := ctx.Input.Header("X-Real-IP")
	ClientIp = ip
	s := strings.Split(ctx.Request.RemoteAddr, ":")
	if ip == "" {
		ClientIp = s[0]
	}
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

func (this *BaseController) Prepare() {
	go func(this *BaseController) {
		if MongoDb == nil {
			return
		}

		r := model.TraceRecord{
			Ip:       this.Ctx.Input.IP(),
			Uri:      this.Ctx.Input.URI(),
			Datetime: time.Now().Format(DatetimeFormat),
			Time:     time.Now(),
		}

		err := MongoDb.C("trace_record").Insert(r)
		Error(err)

	}(this)
}
