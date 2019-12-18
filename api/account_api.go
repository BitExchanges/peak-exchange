package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"peak-exchange/service"
	. "peak-exchange/utils"
)

// 查询账户信息
func GetAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.GetInt("userId")
		accountList := service.SelectUserAccountList(userId)
		ctx.JSON(http.StatusOK, Success(accountList))
	}
}

func GetAccountBalance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.GetInt("userId")

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

// 充值
func Recharge() {

}
