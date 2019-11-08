package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"peak-exchange/utils"
)

// 认证处理
func Authorize() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			ctx.JSON(http.StatusOK, utils.Response{Head: map[string]string{"code": "10000", "msg": "token认证失败"}})
			ctx.Abort()
			return
		} else {
			ctx.Next()
		}
	}
}
