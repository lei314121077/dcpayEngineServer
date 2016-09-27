package routers

import (
	"dc_web/controllers"
	user "dc_web/controllers/user_controller"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &user.UserController{})

}
