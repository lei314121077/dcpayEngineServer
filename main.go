package main

import (
	_ "dc_web/routers"

	"github.com/astaxie/beego"
)

func main() {

	beego.SetLogger("file", `{"filename":"logs/cons_logs.log"}`)      //日志配置
	beego.Run()
}
