package controllers

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
