package api

import (
	"github.com/gin-gonic/gin"
)

// 根据用户ID 获取授权登录地址
func GetAuthLoginAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//userId := ctx.GetInt("userId")
		//limit, page := common.LimitAndPage(ctx.Query("limit"), ctx.Query("page"))

	}
}
