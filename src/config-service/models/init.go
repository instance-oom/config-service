package models

import (
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/samuel/go-zookeeper/zk"
)

var zkClient *zk.Conn
var flags = int32(0)
var acl = zk.WorldACL(zk.PermAll)

//Init : init zookeeper client
func Init() {
	zkaddress := beego.AppConfig.String("zkaddress")
	zk, _, err := zk.Connect(strings.Split(zkaddress, ","), 5000*time.Second)
	if err != nil {
		panic(err)
	}
	zkClient = zk
}
