package config

import (
	"blog_go/dao"
	"log"
	"os"

	"github.com/astaxie/beego"
)

func init() {
	log.Println("init db")
	constr := beego.AppConfig.String("connstr")
	//设置最大空闲连接
	maxIdle := beego.AppConfig.DefaultInt("maxIdle", 10)
	//设置最大数据库连接
	maxConn := beego.AppConfig.DefaultInt("maxConn", 100)
	log.Println("constr:", constr)
	err := dao.InitDB(constr, maxIdle, maxConn)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
}
