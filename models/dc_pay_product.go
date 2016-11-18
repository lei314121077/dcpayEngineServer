package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/**
**支付产品
**/
type DcPayProduct struct {
   id                   string 'orm:"size(50)"'	
   create_time          time.Time 'orm:"auto_now_add;type(datetime)"'				//创建时间
   edit_time            time.Time 'orm:"auto_now_add;type(datetime)"'				//修改时间
   version              int64														//版本编号
   status               string 'orm:"size(36)"'										//状态码
   product_code         string 'orm:"size(50)"'										//支付产品编号
   product_name         string 'orm:"size(200)"'									//支付产品名称
   audit_status         string 'orm:"size(45)"'										//审计状态码
}

func (u *DcPayProduct) TableName() string {
    return "dc_pay_product_tb"   //支付产品表
}