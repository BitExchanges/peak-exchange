package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "peak-exchange/utils"
)

// 发送邮件接口
func SendEmailMsg() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var emailType = ctx.Query("type")
		if emailType == "" {
			ctx.JSON(http.StatusOK, BuildError(ParamError, "参数错误"))
			return
		}

		//err := utils.SendEmail("769558579@qq.com", "wangbbhtt@gmail.com", "异地登录")
		//fmt.Println(err)
		ctx.JSON(http.StatusOK, "邮件发送成功")
	}
}

// 测试获取设备类型
func GetDeviceType() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "测试终端类型")
	}
}
