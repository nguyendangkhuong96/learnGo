package routers

import (
	"learn-go/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //beego.Router("/user", &controllers.UserController{})
	beego.Get("/user",func(ctx *context.Context){
		ctx.Output.Body([]byte("hello world"))
	})

}
