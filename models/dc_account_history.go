package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/**
*资金账户记录
**/
type DcAccountHistory struct{
   id                   string 'orm:"size(50)"'								      //ID 
   create_time          time.Time 'orm:"auto_now_add;type(datetime)"'			  //创建时间
   edit_time            time.Time 'orm:"auto_now;type(datetime)"'			      //更新时间
   version              int64,												      //版本号
   remark               string 'orm:"size(200)"'							      //描述、备注
   account_no           string 'orm:"size(50)"'								      //账户ID
   amount               float64 'orm:"digits(20);decimals(6)"'				      //量级
   balance              float64 'orm:"digits(20);decimals(6)"'				      //平账
   fund_direction       string 'orm:"size(36)"'									  //基金
   is_allow_sett        string 'orm:"size(36)"'									  //是否允许设置
   is_complete_sett     string 'orm:"size(36)"'									  //是否完成设置
   request_no           string 'orm:"size(36)"'									  //请求账户
   bank_trx_no          string 'orm:"size(30)"'									  //银行转账账户
   trx_type             string 'orm:"size(36)"'									  //银行转账类型
   risk_day             int64													  //天数
   user_no              string 'orm:"size(50)"'									  //用户账户
}

func (u *DcAccountHistory) TableName() string {
    return "dc_account_history_tb"   //资金账户记录表
}
