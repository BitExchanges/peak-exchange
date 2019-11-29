package service

import (
	"errors"
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
