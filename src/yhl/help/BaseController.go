package help

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/utils"
	//	"strings"
	"time"
	"yhl/model"
)

var (
	ClientIp    string
	ClientPort  string
	Version     time.Time
	ClientSite  string
	ClientUri   string
	ClientRoute string
)

func init() {
	Version = time.Now()
}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	this.Controller.Init(ctx, controllerName, actionName, app)
	ClientIp = ctx.Input.IP()
	ClientSite = ctx.Input.Site()
	ClientUri = ctx.Input.URI()
	ClientRoute = ClientSite + ClientUri
	//	ip := ctx.Input.Header("X-Real-IP")
	//	ClientIp = ip
	//	s := strings.Split(ctx.Request.RemoteAddr, ":")
	//	if ip == "" {
	//		ClientIp = s[0]
	//	}
	//	ClientPort = s[1]
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
		if !MongoTrace || MongoDb == nil {
			return
		}
		exclude := []string{"/wechat"}
		if utils.InSlice(this.Ctx.Input.URL(), exclude) {
			return
		}

		r := model.TraceRecord{
			Ip:        this.Ctx.Input.IP(),
			Uri:       this.Ctx.Input.URI(),
			Refer:     this.Ctx.Input.Refer(),
			UserAgent: this.Ctx.Input.Header("User-Agent"), //this.Ctx.Input.UserAgent(),
			Datetime:  time.Now().Format(DatetimeFormat),
			Time:      time.Now().Local(),
		}

		err := MongoDb.C("trace_record").Insert(r)
		Error(err)

	}(this)
}

func (this *BaseController) Tips(msg string) {
	this.Redirect("/tips?msg="+msg, 302)
}
