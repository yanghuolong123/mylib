package api

import (
	"yhl/help"
)

type TipsController struct {
	help.BaseController
}

func (this *TipsController) Tips() {
	user := this.GetSession("user")
	if user != nil {
		this.Redirect("/", 302)
	}
	msg := this.GetString("msg")
	tpl := this.GetString("tpl")
	layout := this.GetString("layout")
	if tpl == "" {
		tpl = "tips"
	}
	if layout == "" {
		layout = "main"
	}

	this.Data["msg"] = msg
	this.Layout = "layout/" + layout + ".tpl"
	this.TplName = "tips/" + tpl + ".tpl"
}
