package service

import (
	"errors"
	"peak-exchange/erc20"
	. "peak-exchange/model"
	"peak-exchange/utils"
	"reflect"
	"time"
)

func Save(user User) (int, error) {
	DB := utils.MainDbBegin()
	defer DB.DbCommit()

	result := SelectUserByMobile(user.Mobile)
	if reflect.DeepEqual(result, User{}) {
		DB.Create(&user)
		return user.Id, nil
	} else {
		DB.DbRollback()
		return 0, errors.New("用户已存在")
	}

}

// 根据手机号查询用户信息
func SelectUserByMobile(mobile string) (user User) {
	DB := utils.MainDbBegin()
	defer DB.DbRollback()
	DB.Select([]string{
		"id",
		"uuid",
		"nick_name",
		"avatar",
		"mobile",
		"login_pwd",
		"email",
		"level",
		"kyc_level",
		"identity_card",
		"card_type",
		"last_login_at",
		"last_login_ip",
	}).Where("mobile=?", mobile).Find(&user)
	return user
}

//更新用户信息
func UpdateUser(user User) {
	db := utils.MainDbBegin()
	defer db.DbCommit()
	db.Model(&user).Updates(map[string]interface{}{"last_login_at": time.Now(), "last_login_ip": user.LastLoginIp})
}

// 更新用户最后登录时间 以及登录IP
func UpdateUserLogin(userId int, lastLoginAt time.Time, lastLoginIp string) {
	db := utils.MainDbBegin()
	defer db.DbCommit()
	db.Exec("UPDATE user SET last_login_at=?,last_login_ip=? WHERE user_id=?", lastLoginAt, lastLoginIp, userId)
}

// 保存用户钱包地址
func SaveWalletAddress(userId int) {
	db := utils.MainDbBegin()
	defer db.DbCommit()
	privateKey, Address := erc20.GenerateUserWallet()
	wallet := NewWallet(userId, privateKey, Address)
	db.Create(&wallet)
}
