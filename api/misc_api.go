package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"peak-exchange/common"
	. "peak-exchange/utils"
)

// 发送邮件接口
func SendEmailMsg() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var emailType = ctx.Query("type")
		var email = ctx.Query("email")
		if emailType == "" || email == "" {
			ctx.JSON(http.StatusOK, BuildError(ParamError, "参数错误"))
			return
		}
		redisCon := LimitPool.Get()
		var reply interface{}
		var err error
		var subject string
		var formatter string

		//TODO 暂时先不记录发送邮件次数
		//var counter interface{}

		switch emailType {
		case common.RegisterKey:
			subject = common.RegisterActivateSubject
			formatter = fmt.Sprintf(common.RedisEmailFormatter, email, common.RegisterKey)
			reply, err = redisCon.Do("GET", formatter)
		case common.ForgetPwdKey:
			subject = common.ForgetPwdSubject
			formatter = fmt.Sprintf(common.RedisEmailFormatter, common.ForgetPwdKey)
			reply, err = redisCon.Do("GET", formatter)
		case common.ChangeLoginPwdKey:
			subject = common.ChangeLoginPwdSubject
			formatter = fmt.Sprintf(common.RedisEmailFormatter, common.ChangeLoginPwdKey)
			reply, err = redisCon.Do("GET", formatter)
		case common.ChangeTradePwdKey:
			subject = common.ChangeTradePwdSubject
			formatter = fmt.Sprintf(common.RedisEmailFormatter, common.ChangeTradePwdKey)
			reply, err = redisCon.Do("GET", formatter)
		}

		if err != nil {
			log.Println("查询redis出错: ", err)
			ctx.JSON(http.StatusOK, BuildError(EmailModuleError, "邮件发送服务异常"))
			return
		}

		// 如果发送过邮件，未到期则不允许再次发送
		if reply != nil {
			ctx.JSON(http.StatusOK, BuildError(EmailFrequently, "邮件发送频繁,请5分钟后尝试再次发送"))
			return
		}

		captchaCode := GenerateCode(4)
		go SendEmails(email, subject, subject, subject, subject, captchaCode, "")
		reply, err = redisCon.Do("SET", formatter, captchaCode, "EX", common.RedisEmailExpire)

		if err != nil {
			log.Println("redis写入错误: ", err)
			ctx.JSON(http.StatusOK, BuildError(EmailModuleError, "邮件发送服务异常"))
			return
		}
		ctx.JSON(http.StatusOK, "邮件发送成功")
	}
}

// 测试获取设备类型
func GetDeviceType() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "测试终端类型")
	}
}
