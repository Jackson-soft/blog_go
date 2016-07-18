package controllers

import (
	"blog_go/base/util"
	"blog_go/constant"
	"blog_go/models"
	"errors"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	envelope *models.MsgEnvelope
}

func (c *BaseController) GetPostValue(key string) string {
	return c.Ctx.Request.PostFormValue(key)
}

func (c *BaseController) GetPostUint64(key string, def uint64) (uint64, error) {
	value := c.GetPostValue(key)
	if value != "" {
		return strconv.ParseUint(value, 10, 64)
	} else if value == "" {
		return def, nil
	} else {
		return 0, errors.New("invalid value uint64")
	}
}

func (c *BaseController) GetPostInt64(key string, def int64) (int64, error) {
	value := c.GetPostValue(key)
	if value != "" {
		return strconv.ParseInt(value, 10, 64)
	} else if value == "" {
		return def, nil
	} else {
		return 0, errors.New("invalid value int64")
	}
}

//访问接口前的一些判断
func (c *BaseController) Prepare() {
	c.envelope = &models.MsgEnvelope{
		Data:    util.EmptyMap(),
		Result:  true,
		NowTime: time.Now().Format(constant.FORMAT_DATETIME),
	}
	c.Data["json"] = c.envelope
}
