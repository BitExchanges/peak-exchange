package routes

import (
	"github.com/gin-gonic/gin"
	"peak-exchange/api"
	"peak-exchange/auth"
)

func SetInterfaces(e *gin.Engine) {

	// 路由分组
	// 订单组  orderRoute
	// 用户组  userRoute
	// 交易组  tradeRoute
	orderRoute := e.Group("/api/:platform/v1/order")
	orderRoute.Use(auth.Authorize())
	{
		orderRoute.GET("/getOrderBook", api.GetOrderBook())
		orderRoute.GET("/getOrderByNo/:orderNo", api.GetOrderByNo())
		orderRoute.GET("/getOrderByUserId/:userId", api.GetOrderBookByUserId())
		orderRoute.POST("/taker", api.Maker())
	}

	currencyRoute := e.Group("/api/:platform/v1/currency")
	{
		currencyRoute.GET("/currencyList", api.GetCurrencyList())
	}

	userRoute := e.Group("/api/:platform/v1/user")

	{
		userRoute.POST("/register", api.Register())
	}

}
