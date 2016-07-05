package config

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func init() {
	constr := beego.AppConfig.String("sqlconstr")
	log.Println("constr:", constr)
	db, err := sql.Open("mysql", constr)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err)
	}
}
