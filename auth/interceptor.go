package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"log"
	"net/http"
	"peak-exchange/service"
	. "peak-exchange/utils"
)

// 解决跨域问题
func Cross() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}

// 获取设备类型
func GetDevice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取设备类型
		userAgent := ctx.GetHeader("User-Agent")
		agent := user_agent.New(userAgent)
		if agent.Mobile() {
			ctx.Set("device", "mobile")
		} else {
			ctx.Set("device", "web")
		}
		ctx.Next()
	}
}

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
				log.Println("token 非法:", claims)
				ctx.Abort()
				return
			}

			if service.ValidMobileAndPwd(claims.Mobile, claims.LoginPwd) {
				ctx.Set("userId", claims.Id)
				ctx.Next()
			} else {
				ctx.JSON(http.StatusOK, BuildError(AuthorizationFail, "认证失败"))
				ctx.Abort()
				return
			}
		}
	}
}
