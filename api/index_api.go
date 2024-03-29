package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"peak-exchange/utils"
	"strings"
)

//登录页面加载
func LoginIndex() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.tmpl", gin.H{"login": "login"})
	}
}

func Index() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, captcha, _ := utils.GenerateCaptcha("digit")

		qrCodeStr := utils.GenerateQRCodeBase64("http://www.baidu.com", 80)

		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"qrcode": qrCodeStr, "captcha": strings.Split(captcha, "data:image/png;base64,")[1]})
	}
}

// 邮箱模板测试
func EmailTemplate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "email.html", gin.H{"test": "test"})
	}
}
