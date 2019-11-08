package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	. "peak-exchange/model"
)

func GetOrderBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param("platform")
		fmt.Printf("接收参数: %s\n", param)
		ctx.JSON(http.StatusOK, gin.H{"message": "order_book"})
	}
}

// 根据订单编号查询订单详情
func GetOrderByNo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orderNo := ctx.Param("orderNo")
		fmt.Printf("订单号为: %s \n", orderNo)
		ctx.JSON(http.StatusOK, gin.H{"message": "根据订单编号查询订单详情"})
	}
}

// 根据用户ID查询挂单簿
func GetOrderBookByUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "根据用户信息查询挂单簿"})
	}
}

// 挂单
func Maker() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var order Order
		err := ctx.BindJSON(&order)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "挂单"})
		}

	}
}

// 吃单
func Taker() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "吃单"})
	}
}
