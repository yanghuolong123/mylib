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
	db := beego.AppConfig.String("mongo.db")
	if host == "" || db == "" {
		return
	}

	MongoConn, err := mgo.Dial(host)

	Error(err)
	if err == nil {
		MongoDb = MongoConn.DB(db)
	}

}
