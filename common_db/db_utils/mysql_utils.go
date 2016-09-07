package db_utils

//mysql工具类

import (
	"github.com/astaxie/beego"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	
)

var db *sql.DB

func init() {
	var my_user string := beego.AppConfig.String("mysql_user")
	var my_host string := beego.AppConfig.String("mysql_host")
	var my_prot string := beego.AppConfig.String("mysql_port")
	var my_name string := beego.AppConfig.String("mysql_name")
	var my_pswd string := beego.AppConfig.String("mysql_pswd")
	
	//链接配置如下 db, _ = sql.Open("mysql", "root:passwd@tcp(127.0.0.1:3306)/testdb?charset=utf8")
 
	db, _ = sql.Open("mysql", "%s:%s@tcp(%s:%s)/%s?charset=utf8" %(my_user, my_pswd, my_host, my_prot, my_name)) 
	db.SetMaxOpenConns(2000)  
    db.SetMaxIdleConns(1000)
    db.Ping()

}

//新增一条记录
func insert_process (tabname string, column make(map[string]string))  (res int){
	/**
	*param tabname 数据库表名
	*param column {属性名:属性值}
	*return 执行结果
	**/
	var result := 0
	db.Prepare('INSERT ') 
	return result
}
