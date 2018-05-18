package api

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
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

func (this *UploadController) WebUpload() {
	method := this.Ctx.Input.Method()
	if method == "OPTIONS" {
		this.StopRun()
	}
	filename := this.Ctx.Input.Query("name")
	chunks := this.Ctx.Input.Query("chunks")
	chunk := this.Ctx.Input.Query("chunk")
	if chunk == "" {
		chunks = "1"
		chunk = "0"
	}
	f, h, err := this.GetFile("file")
	if err != nil {
		help.Error(err)
		this.SendResJsonp(101, "fail", err.Error())
	}
	defer f.Close()
	//	fmt.Println("==== chunks:", chunks)
	//	fmt.Println("==== chunk:", chunk)
	ext := filepath.Ext(h.Filename)
	filename = help.Md5(filename) + ext
	prefix := "tmp/"
	part := prefix + filename + "_" + chunk + ".part"
	this.SaveToFile("file", part)
	count, err := strconv.Atoi(chunks)
	cache := help.Cache
	cache.Incr(filename)
	num, err := strconv.Atoi(string(cache.Get(filename).([]uint8)))
	dir := "uploads/"
	y, m, d := help.Date()
	dir = dir + fmt.Sprintf("%d/%d/%d/", y, m, d)
	//	fmt.Println("==== num:", num)
	//	fmt.Println("==== count:", count)
	//	fmt.Println("========== filename:", filename)
	outfile := dir + fmt.Sprintf("%s%d", time.Now().Format(help.DatetimeNumFormat), help.RandNum(10000, 99999)) + ext
	if num == count {
		go func(prefix, filename, outDir, outfile string) {
			cache.Delete(filename)
			if !help.PathExist(outDir) {
				os.MkdirAll(outDir, os.ModePerm)
			}
			out, _ := os.OpenFile(outfile, os.O_CREATE|os.O_WRONLY, 0666)
			bWriter := bufio.NewWriter(out)
			for i := 0; i < count; i++ {
				infile := prefix + filename + "_" + strconv.Itoa(i) + ".part"
				in, _ := os.Open(infile)
				bReader := bufio.NewReader(in)
				for {
					buffer := make([]byte, 1024)
					readCount, err := bReader.Read(buffer)
					if err == io.EOF {
						os.Remove(infile)
						break
					} else {
						bWriter.Write(buffer[:readCount])
					}
				}
			}
			bWriter.Flush()

		}(prefix, filename, dir, outfile)

		cache.Delete(filename)
	} else if num > count {
		cache.Delete(filename)
		this.SendResJsonp(101, "fail", "uplad fail, please upload again!")
	}

	this.SendResJsonp(0, "ok", "/"+outfile)
}

func (this *UploadController) SendResJsonp(code int, msg string, data interface{}) {
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
