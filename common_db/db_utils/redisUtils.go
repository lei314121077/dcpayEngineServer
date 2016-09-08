package db_utils

import (
	"github.com/astaxie/beego"
)


func init(){
	var red_user string := beego.AppConfig.String("redis_user")
	var red_host string := beego.AppConfig.String("redis_host")
	var red_host string := beego.AppConfig.String("redis_port")
	// := beego.AppConfig.String("")
	// := beego.AppConfig.String("")
	
	
}