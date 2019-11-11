package service

import (
	"fmt"
	. "peak-exchange/model"
	"peak-exchange/utils"
	"reflect"
)

func Save(user User) {
	DB := utils.MainDbBegin()
	defer DB.DbCommit()

	result := SelectUserByMobile(user.Mobile)
	if reflect.DeepEqual(result, User{}) {
		DB.Create(&user)
	} else {
		fmt.Println("用户已存在")
	}

}

// 根据手机号查询用户信息
func SelectUserByMobile(mobile string) (user User) {
	DB := utils.MainDbBegin()
	defer DB.DbRollback()
	DB.Where("mobile=?", mobile).Find(&user)
	return user
}
