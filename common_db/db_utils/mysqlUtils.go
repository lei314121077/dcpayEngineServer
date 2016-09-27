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
 
	db, err = sql.Open("mysql", "%s:%s@tcp(%s:%s)/%s?charset=utf8" %(my_user, my_pswd, my_host, my_prot, my_name)) 

	if err != nil {
        log.Fatal(err)
    }
    
    //设置数据库连接池最大的链接数
    db.SetMaxOpenConns(2000)  
    db.SetMaxIdleConns(1000)
    
    //检查是否可能连接到数据库
    err_ping = db.Ping() 
    if err_ping != nil{
    	log.Fatal(err)
    }

    //关闭连接池
    defer db.Close()
}

// http://dmdgeeker.com/goBook/docs/ch05/overview.html
//TODO 新增一条记录  
func insertProcess (tabname string, column make(map[string]string))  (res int){
	/**
	*param tabname 数据库表名
	*param column dict_type{属性名:属性值}
	*return 执行结果
	**/
	
	stmt, err := db.Prepare('INSERT %s (%s) values (%s)')
	res, err := stmt.Exec


	return result
}

//TODO 查询记录
func queryProcess(tabname string, column [...]string, conditions make(map[string]string))(res string){
	/**
	*param tabname 数据库表名
	*param cloumn  array_type{属性名1，属性名2}
	*return res 
	**/


}

//TODO 单个查询
func queryOne(tabname string, cloumn [...]string, conditions make(map[string]string))(res string){
	/**
	*param tabname 数据库表名
	*param cloumn  array_type{属性名1，属性名2}
	*return res 	
	**/

}













