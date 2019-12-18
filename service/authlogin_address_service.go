package service

import (
	"errors"
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

// 根据用户ID 查询授权地址列表
func SelectAuthLoginAddressPage(userId, limit, page int) ([]AuthLoginAddress, int, error) {
	var authAddressList []AuthLoginAddress
	var count int
	db := utils.MainDbBegin()
	defer db.CommonDB()
	db.Table("auth_login_address").Count(&count)
	if db.Offset(limit*(page-1)).Limit(limit).Where("user_id=?", userId).Find(&authAddressList).RecordNotFound() {
		return nil, 0, errors.New("暂无数据")
	}
	return authAddressList, count, nil
}

// 新增授权地址
func InsertAuthLoginAddress(authLoginAddress AuthLoginAddress) {
	db := utils.MainDbBegin()
	defer db.CommonDB()
	db.Create(&authLoginAddress)
}

// 删除授权地址
func DeleteAuthLoginAddress(userId, id int) {
	db := utils.MainDbBegin()
	defer db.CommonDB()
	db.Delete(&AuthLoginAddress{}, "id=? AND user_id=?", id, userId)
}
