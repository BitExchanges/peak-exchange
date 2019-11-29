package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"peak-exchange/utils"
)

func SendEmailMsg() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := utils.SendEmail("769558579@qq.com", "wangbbhtt@gmail.com", "异地登录")
		fmt.Println(err)
		ctx.JSON(http.StatusOK, "邮件发送成功")
	}
}
