package dao

import (
	"blog_go/database/mysql"
	"log"
)

var (
	Database *mysql.Mysql
)

func InitDB(dbConn string, maxIdle, maxConns int) error {
	log.Println("initdb")
	Database = mysql.NewMysql()
	err := Database.Open(dbConn, maxIdle, maxConns)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
