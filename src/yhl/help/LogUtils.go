package help

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

/*
var (
	Log *logs.BeeLogger
)

func init() {
	Log = logs.NewLogger(10000)
	//	Log.SetLogger(logs.AdapterConsole, "")
	//Log.SetLogger(logs.AdapterFile, `{"filename":"/var/log/go-web/go-web.log"}`)
	Log.SetLogger(logs.AdapterFile, `{"filename":"./log/web.log", "daily":true, "rotate":true}`)
	//Log.SetLogger(logs.AdapterFile, `{"filename":"/var/work/work_golang/work/big_wechat/src/webapp/log/go-web.log"}`)
	Log.EnableFuncCallDepth(true)
}
*/

func Log(filename, msg string) {
	f := "./log/" + filename
	m := make(map[string]string)
	m["filename"] = f
	jsonStr, _ := json.Marshal(m)

	log := logs.NewLogger(10000)
	log.SetLogger(logs.AdapterFile, string(jsonStr))
	log.EnableFuncCallDepth(true)

	log.Info(msg)
}
