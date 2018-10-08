package main

import (
	_ "learn-go/routers"
	"github.com/astaxie/beego"
)

func main() {

	beego.AppConfig.String("mysqluser")
	beego.AppConfig.String("mysqlpass")
	beego.AppConfig.String("mysqlurls")
	beego.AppConfig.String("mysqldb")



	beego.Run()
}

