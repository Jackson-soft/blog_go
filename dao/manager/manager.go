package manager

import (
	"blog_go/dao"
)

func QueryManagerByMobile(mobile string) (map[string]interface{}, error) {
	sql := "SELECT id,nick_name as nickName,password FROM manager_info WHERE mobile=?"
	return dao.Database.QueryForMap(sql, mobile)
}

//插入用户信息
func InsertManagerInfo(mobile, password string) (int64, error) {
	sql := "INSERT INTO manager_info (nick_name,mobile,password) VALUES (?,?,?)"
	return dao.Database.Insert(sql, mobile, mobile, password)
}
