package main

import (
	_ "dc_web/routers"

	"github.com/astaxie/beego"
)

func main() {

	beego.BConfig.WebConfig.Session.SessionName = "paySessionID" //session name
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)      //日志配置
	beego.Run()
}
