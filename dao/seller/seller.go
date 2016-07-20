package seller

import (
	"blog_go/dao"
)

func QuerySellerByMobile(mobile string) (map[string]interface{}, error) {
	sql := "SELECT id,nick_name as nickName,password FROM seller_info WHERE mobile=?"
	return dao.Database.QueryForMap(sql, mobile)
}

//插入用户信息
func InsertSellerInfo(mobile, password string) (int64, error) {
	sql := "INSERT INTO seller_info (nick_name,mobile,password) VALUES (?,?,?)"
	return dao.Database.Insert(sql, mobile, mobile, password)
}
