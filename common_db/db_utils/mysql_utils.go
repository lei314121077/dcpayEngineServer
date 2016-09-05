package db_utils

//mysql工具类

import (
	"github.com/astaxie/beego"
	
	
)

func init() {
	var my_user string := beego.AppConfig.String("mysql_user")
	var my_host string := beego.AppConfig.String("mysql_host")
	var my_prot string := beego.AppConfig.String("mysql_port")
	var my_name string := beego.AppConfig.String("mysql_name")
	var my_pswd string := beego.AppConfig.String("mysql_pswd")
}

//新增一条记录
func insert_one(n int){
	
}
