package main

import (
	_ "blog_go/routers"

	//_ "blog_go/config"
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
	"log"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	constr := beego.AppConfig.String("sqlconstr")
	maxIdle := beego.AppConfig.DefaultInt("maxIdle", 10)
	maxConn := beego.AppConfig.DefaultInt("maxConn", 50)
	log.Println("mysql:", constr, maxIdle, maxConn)
	orm.RegisterDataBase("mydemo", "mysql", constr, maxIdle, maxConn)
}

func main() {
	o := orm.NewOrm()
	o.Using("mydemo")
	beego.Run()
}
