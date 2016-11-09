package user_controller

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["Website"] = "raymond"
	c.Data["Email"] = "raymondldx@gmail.com"
	c.TplName = "index.tpl"
}
