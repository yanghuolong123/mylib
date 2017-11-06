package help

import (
	"github.com/astaxie/beego/logs"
)

var (
	Log *logs.BeeLogger
)

func init() {
	Log = logs.NewLogger(10000)
	Log.SetLogger(logs.AdapterConsole, "")
	//Log.SetLogger(logs.AdapterFile, `{"filename":"/var/log/go-web/go-web.log"}`)
	Log.EnableFuncCallDepth(true)
}
