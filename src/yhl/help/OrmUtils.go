package help

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbhost := beego.AppConfig.String("mysql.host")
	dbport := beego.AppConfig.String("mysql.port")
	dbuser := beego.AppConfig.String("mysql.user")
	dbpasswd := beego.AppConfig.String("mysql.pass")
	dbname := beego.AppConfig.String("mysql.dbname")

	conn := dbuser + ":" + dbpasswd + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Local"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", conn)
}
