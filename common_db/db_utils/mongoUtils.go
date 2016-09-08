package db_utils

//mongodb工具类

import (
	"github.com/astaxie/beego"
)

func init() {
	var mg_user string := beego.AppConfig.String("mongo_user")
	var mg_host string := beego.AppConfig.String("mongo_host")
	var mg_prot string := beego.AppConfig.String("mongo_prot")
	var mg_name string := beego.AppConfig.String("mongo_name")
	var mg_pswd string := beego.AppConfig.String("mongo_pswd")
	

}






