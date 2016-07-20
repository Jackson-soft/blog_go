package controllers

import (
	"blog_go/base/util"
	"blog_go/dao/user"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserLoginController struct {
	BaseController
}

//用户注册
func (c *UserLoginController) Register() {
	envelope := c.envelope
	defer c.ServeJSON()
	sMobile := c.GetPostValue("mobile")
	sPassword := c.GetPostValue("password")
	log.Println("values:", sMobile, sPassword)
	if sMobile == "" || sPassword == "" {
		envelope.Result = false
		envelope.ErrMsg = "参数不全"
		return
	}

	mUser, err := user.QueryUserByMobile(sMobile)
	if err != nil {
		log.Println(err)
		envelope.Result = false
		envelope.ErrMsg = "数据库错误"
		return
	}
	if mUser != nil {
		envelope.Result = false
		envelope.ErrMsg = "该手机号已注册"
		return
	}
	hPassword, err := bcrypt.GenerateFromPassword([]byte(sPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		envelope.ErrMsg = "密码加密失败"
		envelope.Result = false
		return
	}
	log.Println("hPassword:", string(hPassword))
	nId, err := user.InsertUserInfo(sMobile, string(hPassword))
	if err != nil {
		log.Println(err)
		envelope.Result = false
		envelope.ErrMsg = "数据库错误"
		return
	}
	mResult := util.EmptyMap()
	mResult["id"] = nId
	mResult["nickName"] = sMobile
	envelope.Data = mResult
}

//用户登录
func (c *UserLoginController) Login() {
	envelope := c.envelope
	defer c.ServeJSON()

	sMobile := c.GetPostValue("mobile")
	sPassword := c.GetPostValue("password")
	if sMobile == "" || sPassword == "" {
		envelope.Result = false
		envelope.ErrMsg = "参数不全"
		return
	}

	mUser, err := user.QueryUserByMobile(sMobile)
	if err != nil {
		log.Println(err)
		envelope.Result = false
		envelope.ErrMsg = "数据库错误"
		return
	}

	if mUser == nil {
		envelope.ErrMsg = "没有此用户"
		envelope.Result = false
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(mUser["password"].(string)), []byte(sPassword))
	if err != nil {
		log.Println(err)
		envelope.Result = false
		envelope.ErrMsg = "密码校验失败"
		return
	}

	c.SetSession("userMobile", sMobile)

	mResult := util.EmptyMap()
	mResult["id"] = mUser["id"]
	mResult["nickName"] = mUser["nickName"]
	envelope.Data = mResult
}

//用户的接口
type UserController struct {
	BaseController
}

//订单号
func (c *UserController) DealNum() {
	envelope := c.envelope
	defer c.ServeJSON()
	sDealNum := c.GetPostValue("dealNum")
	if sDealNum == "" {
		envelope.ErrMsg = "交易号不能为空"
		envelope.Result = false
		return
	}
}
