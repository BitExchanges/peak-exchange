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
		return user.ID, nil
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
		"state",
		"random_uuid",
	}).Where("mobile=?", mobile).Find(&user)
	return user
}

// 根据邮箱查询用户信息
func SelectUserByEmail(email string) (user User) {
	db := utils.MainDbBegin()
	defer db.DbCommit()
	db.Select([]string{
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
		"state",
		"random_uuid",
	}).Where("email=?", email).Find(&user)
	return user
}

// 根据邮箱修改登录密码
func UpdateLoginPwdByEmail(email, pwd string) {
	db := utils.MainDbBegin()
	defer db.CommonDB()
	db.Model(&User{}).Where("email=?", email).Update("login_pwd", pwd)
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

// 用户注册
func UserRegister(user User) (User, error) {

	db := utils.MainDbBegin()
	//添加用户信息
	var count int
	db.Model(&User{}).Where("mobile=?", user.Mobile).Count(&count)
	if count > 0 {
		return User{}, errors.New("用户已存在")
	}

	if db.Create(&user).Error != nil {
		return User{}, errors.New("创建用户失败")
	}

	//创建虚拟账户信息
	virtualAccount := CreateVirtualAccount(user.ID)
	if db.Create(&virtualAccount).Error != nil {
		db.DbRollback()
		return User{}, errors.New("创建用户账户信息失败")
	}

	//创建真实账户
	realAccount := CreateRealAccount(user.ID)
	if db.Create(&realAccount).Error != nil {
		db.DbRollback()
		return User{}, errors.New("创建账户信息失败")
	}

	//创建钱包地址
	privateKey, address := erc20.GenerateUserWallet()
	wallet := NewWallet(user.ID, privateKey, address)

	if db.Create(&wallet).Error != nil {
		db.DbRollback()
		return User{}, errors.New("钱包地址创建失败")
	}

	db.DbCommit()
	return user, nil
}

// 校验用户手机号和密码
func ValidMobileAndPwd(mobile, pwd string) bool {
	db := utils.MainDbBegin()
	var count int
	db.Model(&User{}).Where("mobile=? and login_pwd=?", mobile, pwd).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

// 修改用户状态信息
func UpdateUserActive(uuid string) {
	db := utils.MainDbBegin()
	db.Exec("UPDATE user SET state=1 WHERE random_uuid=?", uuid)
	db.DbCommit()
}
