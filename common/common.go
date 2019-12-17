package common

import (
	"strconv"
)

var (
	RedisEmailForgetPwd       = "%s_forget_captcha"
	RedisEmailForgetPwdExpire = "300"
)

// 计算分页参数
func LimitAndPage(limit, page string) (int, int) {
	limitTemp, pageTemp := 20, 1
	if limit != "" {
		limit, _ := strconv.Atoi(limit)
		if limit > 100 {
			limitTemp = 20
		} else {
			limitTemp = limit
		}
	}

	if page != "" {
		page, _ := strconv.Atoi(page)
		if page < 1 {
			pageTemp = 1
		} else {
			pageTemp = page
		}
	}
	return limitTemp, pageTemp
}
