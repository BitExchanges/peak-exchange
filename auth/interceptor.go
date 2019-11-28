package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	. "peak-exchange/utils"
)

// 认证处理
func Authorize() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctx.JSON(http.StatusOK, BuildError(AccessDenied, "暂无权限"))
			ctx.Abort()
			return
		} else {
			//解析token
			j := NewJwt()
			claims, err := j.ParseToken(token)
			if err != nil {
				ctx.JSON(http.StatusOK, BuildError(IllegalToken, "非法token"))
				fmt.Println("token 非法:", claims)
				return
			}
			ctx.Set("userId", claims.Id)
			ctx.Next()
		}
	}
}
