package api

import "github.com/gin-gonic/gin"

// 查询账户信息
func GetAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// 创建账户
//func CreateAccount() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		token := ctx.GetHeader("token")
//		//TODO token解析
//
//	}
//}
