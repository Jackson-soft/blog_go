package controllers

import (
	"github.com/astaxie/beego"
)

//后台管理接口
type ManagerController struct {
	beego.Controller
    managerId uint64

}


//访问接口前的一些判断
func (c *ManagerController) Prepare() {
    
}

//登录
func (c *ManagerController) Login () {
    
}

