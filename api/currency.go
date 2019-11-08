package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 查询所有币种
func GetCurrencyList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "查询币种"})
	}
}

// 查询币种详情
func GetCurrencyInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "币种详情"})
	}
}
