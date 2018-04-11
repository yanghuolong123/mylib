package api

import (
	"fmt"
	"path/filepath"
	"time"
	"yhl/help"
)

type UploadController struct {
	help.BaseController
}

func (this *UploadController) Uploadfile() {
	f, h, err := this.GetFile("file")
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}
	defer f.Close()

	ext := filepath.Ext(h.Filename)
	filename := time.Now().Format(help.DatetimeNumFormat) + help.IntToStr(help.RandNum(1000, 9999)) + ext
	dir := "uploads/"
	y, m, d := help.Date()
	dir += fmt.Sprintf("%d/%d/%d/", y, m, d)
	help.MkDirPath(dir)
	fileRoute := dir + filename
	this.SaveToFile("file", fileRoute)

	this.SendRes(0, "success", "/"+fileRoute)
}

func (this *UploadController) CkUploadfile() {
	callback := this.GetString("CKEditorFuncNum")
	f, h, err := this.GetFile("upload")
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}
	defer f.Close()

	ext := filepath.Ext(h.Filename)
	filename := time.Now().Format(help.DatetimeNumFormat) + help.IntToStr(help.RandNum(1000, 9999)) + ext
	dir := "uploads/"
	y, m, d := help.Date()
	dir += fmt.Sprintf("%d/%d/%d/", y, m, d)
	help.MkDirPath(dir)
	fileRoute := dir + filename
	this.SaveToFile("upload", fileRoute)
	fpath := "/" + fileRoute

	out := "<script type=\"text/javascript\">"
	out += "window.parent.CKEDITOR.tools.callFunction(" + callback + ",\"" + fpath + "\",\"\")"
	out += "</script>"

	this.Ctx.WriteString(out)
}
