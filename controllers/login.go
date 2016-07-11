package controllers

import "github.com/astaxie/beego"

// 这里处理所有的登录
type UserLoginController struct {
	beego.Controller
}

//
func (c *UserLoginController) Login() {
	c.SessionRegenerateID()
}

//卖家登录
type SellerLoginController struct {
	beego.Controller
}

func (c *SellerLoginController) Login() {
	c.SessionRegenerateID()
}

//后台人员登录
type ManagerLoginController struct {
	beego.Controller
}

func (c *ManagerLoginController) Login() {
	id := 1
	c.SessionRegenerateID()
	c.SetSession("managerID", id)
}
