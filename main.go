package main

import (
	_ "blog_go/routers"

	"github.com/astaxie/beego"

	_ "blog_go/config"
	"blog_go/dao"
)

func main() {
	defer dao.Database.Close()
	beego.Run()
}
