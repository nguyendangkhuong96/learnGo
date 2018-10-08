package routers

import (
	"github.com/astaxie/beego"
	"github.com/learn-go/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/", &controllers.UserController{}, "get:ListUsers;post:NewUser")
}
