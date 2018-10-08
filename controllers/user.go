package controllers

import ("github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/learn-go/models")

type UserController struct {
	beego.Controller
}

func (this *UserController) ListUsers() {
	db := orm.NewOrm()

	var users []models.User
	qs := db.QueryTable("products")

	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")

	qs.Limit(limit, offset).All(&products)
	c.Data["json"] = &products
	c.ServeJson()
}