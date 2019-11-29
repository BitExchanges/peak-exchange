package service

import (
	. "peak-exchange/model"
	"peak-exchange/utils"
)

// 根据用户ID,ip地址 查询授权常用登录地
func SelectAuthLoginAddressByUserId(userId int, ipAddress string) (authAddress AuthLoginAddress) {
	db := utils.MainDbBegin()
	defer db.DbCommit()
	db.Where("user_id=? AND ip_address=? AND state=1", userId, ipAddress).Find(&authAddress)
	return authAddress
}

// 修改常用地址状态
func UpdateAuthLoginAddressState(userId, id int) {
	db := utils.MainDbBegin()
	defer db.DbCommit()
	db.Exec("UPDATE auth_login_address SET state=1 WHERE user_id=? AND id=?", userId, id)

}
