package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*
 *资金账户
*/
type DcAccount struct{
   id                	string 'orm:"size(50)"'								      //ID 
   create_time          time.Time 'orm:"auto_now_add;type(datetime)"'		//创建时间
   edit_time            time.Time 'orm:"auto_now;type(datetime)"'			   //更新时间
   version              int64,												         //版本号
   remark               string 'orm:"size(200)"'							      //描述、备注
   account_no           string 'orm:"size(50)"'								      //账户ID
   balance              float64 'orm:"digits(20);decimals(6)"'				   //平账
   unbalance            float64 'orm:"digits(20);decimals(6)"'				   //不平账
   security_money       float64 'orm:"digits(20);decimals(6)"'				   //安全资金
   status               string 'orm:"size(36)"'								      //状态
   total_income         dfloat64 'orm:"digits(20);decimals(6)"'				//总收益
   total_expend         float64 'orm:"digits(20);decimals(6)"'				   //总花费
   today_income         float64 'orm:"digits(20);decimals(6)"'				   //当日收益
   today_expend         float64 'orm:"digits(20);decimals(6)"'				   //当日支出
   account_type         string 'orm:"size(50)"'								      //账户类型
   sett_amount          float64 'orm:"digits(20);decimals(6)"'				   //
   user_no              string 'orm:"size(50)"'								      //用户ID
   primary key (id)                                                        //主键


}

func (u *DcAccount) TableName() string {
    return "dc_account_tb"   //资金账户表
}



