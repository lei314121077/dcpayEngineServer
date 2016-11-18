package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/**
**支付方式
**/
type DcPayWay struct{
   id                   string 'orm:"size(50)"'	
   version              int64
   create_time          time.Time 'orm:"auto_now_add;type(datetime)"'			  //创建时间
   edit_time            time.Time 'orm:"auto_now_add;type(datetime)"'			  //修改时间
   pay_way_code         string 'orm:"size(50)"'									  //支付方式编号
   pay_way_name         string 'orm:"size(100)"'    							  //支付方式名称
   pay_type_code        string 'orm:"size(50)"'									  //支付类型编号
   pay_type_name        string 'orm:"size(100)"'								  //支付类型名称
   pay_product_code     string 'orm:"size(50)"'									  //支付产品编号
   status               string 'orm:"size(36)"' 								  //状态(100:正常状态,101非正常状态)
   sorts                int64 default 1000 										  //排序(倒序排序,默认值1000)
   pay_rate             float64 												  //商户支付费率
}

func (u *DcPayWay) TableName() string {
    return "dc_pay_way_tb"   //支付方式表
}