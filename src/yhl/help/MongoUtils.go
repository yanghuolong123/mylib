package help

import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo"
)

var (
	MongoConn *mgo.Session
	MongoDb   *mgo.Database
)

func init() {
	host := beego.AppConfig.String("mongo.host")
	MongoConn, _ := mgo.Dial(host)

	db := beego.AppConfig.String("mongo.db")
	MongoDb = MongoConn.DB(db)
}
