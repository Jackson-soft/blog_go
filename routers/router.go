package routers

import (
	"blog_go/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})

	beego.AutoRouter(&controllers.UserLoginController{})
	beego.AutoRouter(&controllers.SellerLoginController{})
	beego.AutoRouter(&controllers.ManagerLoginController{})

	beego.AutoRouter(&controllers.SellerController{})
	beego.AutoRouter(&controllers.ManagerController{})
}
