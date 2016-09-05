package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.TplName = "404.tpl"
	c.Data["content"] = "404 Not Found 请求失败，您请求所的资源未在服务器上发现。"
}

func (this *ErrorController) ErrorDb() {
	this.TplName = "Db.tpl"
	this.Data["content"] = "db error 请求失败，您所请求的资源未在数据库找到。"

}
