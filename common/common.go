package common

import (
	"strconv"
)

var (
	RedisEmailFormatter = "%s_%s_captcha"

	//验证码有效期 5分钟
	RedisEmailExpire = "300"

	RegisterKey       = "register"  //注册key
	ForgetPwdKey      = "forgetPwd" //忘记密码key
	ChangeTradePwdKey = "changePwd" //修改交易密码
	ChangeLoginPwdKey = "changePwd" //修改登录密码

	//记录邮件发送次数
	RedisEmailCountFormatter = "%s_%s_count"
	//每天邮件发送次数上限为10
	EmailSendCount = "10"
)

var (
	ChangeLoginPwdSubject   = "修改登录密码"
	ChangeTradePwdSubject   = "修改交易密码"
	RegisterActivateSubject = "注册激活"
	ForgetPwdSubject        = "忘记密码"
)

var ActiveUrl = "http://localhost:8080/api/dev/v1/user/active?uuid=%s"

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
