package help

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
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

func Log(filename string, info interface{}) {
	f := "./log/" + filename + "." + time.Now().Format(DateFormat) + ".log"
	m := make(map[string]interface{})
	m["filename"] = f
	m["daily"] = false
	m["rotate"] = false
	jsonStr, _ := json.Marshal(m)

	log := logs.NewLogger(10000)
	log.SetLogger(logs.AdapterFile, string(jsonStr))
	log.SetLogFuncCallDepth(3)
	log.EnableFuncCallDepth(true)
	log.Async()

	msg := fmt.Sprintf("%+v", info)
	log.Info(msg)
}
