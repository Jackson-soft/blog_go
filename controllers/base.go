package controllers

import (
	"blog_go/constant"
	"blog_go/models"
	"time"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	envelope *models.MsgEnvelope
}

//访问接口前的一些判断
func (c *BaseController) Prepare() {
	c.envelope = &models.MsgEnvelope{
		Result:  true,
		NowTime: time.Now().Format(constant.FORMAT_DATETIME),
	}
	c.Data["json"] = c.envelope
}
