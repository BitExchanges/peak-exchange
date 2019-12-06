package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"peak-exchange/utils"
)

//登录页面加载
func LoginIndex() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.tmpl", gin.H{"login": "login"})
	}
}

func Index() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		qrCodeStr := utils.GenerateQRCodeBase64("http://www.baidu.com", 256)
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"qrcode": qrCodeStr})
	}
}
