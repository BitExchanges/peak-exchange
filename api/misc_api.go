package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendEmailMsg() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
