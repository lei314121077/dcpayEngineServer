package routers

import (
	"dc_web/controllers"
	user "dc_web/user_module"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &user.UserController{})

}

func userModule() {
}

func shopModle() {

}
