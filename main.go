package main

import (
	_ "blog_go/routers"

	_ "blog_go/config"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
