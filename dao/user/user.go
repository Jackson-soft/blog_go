package user

import (
	"blog_go/dao"
)

func QueryUserByMobile(mobile string) (map[string]interface{}, error) {
	sql := "SELECT id,nick_name as nickName,password FROM user_info WHERE mobile=?"
	return dao.Database.QueryForMap(sql, mobile)
}

//插入用户信息
func InsertUserInfo(mobile, password string) (int64, error) {
	sql := "INSERT INTO user_info (nick_name,mobile,password) VALUES (?,?,?)"
	return dao.Database.Insert(sql, mobile, mobile, password)
}
