package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"peak-exchange/utils"
)

// 认证处理
func Authorize() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctx.JSON(http.StatusOK, utils.Response{Head: map[string]string{"code": "10000", "msg": "token认证失败"}})
			ctx.Abort()
			return
		} else {
			//解析token
			j := NewJwt()
			claims, err := j.ParseToken(token)
			if err != nil {
				fmt.Println("token解析校验失败:", claims)
			}
			ctx.Set("userId", claims.Id)
			ctx.Next()
		}
	}
}
