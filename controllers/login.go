package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Login() {
	beego.Info("wqlldl")
	this.Data["json"] = "fdfid"
	this.ServeJSON()
}