package service

import (
	"fmt"
	. "peak-exchange/model"
	"peak-exchange/utils"
	"reflect"
)

func Save(user User) {
	//DB := utils.MainDbBegin()
	result := SelectUserByMobile(user.Mobile)
	if reflect.DeepEqual(result, User{}) {
		fmt.Println("可以注册")
	} else {
		fmt.Println("用户已存在")
	}

}

// 根据手机号查询用户信息
func SelectUserByMobile(mobile string) (user User) {
	DB := utils.MainDbBegin()
	DB.Where("mobile=?", mobile).Find(&user)
	return user
}
