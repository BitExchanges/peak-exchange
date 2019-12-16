package routes

import (
	"github.com/gin-gonic/gin"
	"peak-exchange/api"
	"peak-exchange/auth"
)

func SetInterfaces(e *gin.Engine) {

	// 订单组
	orderRoute := e.Group("/api/:platform/v1/order")
	//orderRoute.Use(auth.Authorize())
	{
		orderRoute.GET("/getOrderBook", api.GetOrderBook())
		orderRoute.GET("/getOrderByNo/:orderNo", api.GetOrderByNo())
		orderRoute.GET("/getOrderByUserId", api.GetOrderBookByUserId())
		orderRoute.POST("/taker", api.Maker())
	}

	//币组
	currencyRoute := e.Group("/api/:platform/v1/currency")
	{
		currencyRoute.GET("/currencyList", api.GetCurrencyList()) //查询当前交易对
	}

	//用户组
	userRoute := e.Group("/api/:platform/v1/user")
	{
		userRoute.POST("/register", api.Register())             //注册
		userRoute.POST("/login", api.Login())                   //登录
		userRoute.POST("/updateProfile", api.UpdateProfile())   //更新个人资料
		userRoute.POST("/forgetLoginPwd", api.ForgetPwd())      //忘记登录密码
		userRoute.POST("/changeLoginPwd", api.ChangeLoginPwd()) //修改登录密码
		userRoute.POST("/changeTradePwd", api.ChangeTradePwd()) //修改交易密码
		userRoute.POST("/logout", api.Logout())                 //退出登录
	}
	//杂项组
	miscRoute := e.Group("/api/:platform/v1/misc")
	miscRoute.Use(auth.Authorize())
	{
		miscRoute.POST("/sendEmail", api.SendEmailMsg()) //发送邮件
		miscRoute.GET("/device", api.GetDeviceType())    //查看设备类型
	}

	//区块
	blockRoute := e.Group("/api/:platform/v1/block")
	{
		blockRoute.GET("/getBlockHead", api.GetCurrentBlockHead())
	}

	//钱包
	walletRoute := e.Group("/api/:platform/v1/wallet")
	{
		walletRoute.GET("/getWallet")
		walletRoute.GET("/addressList", api.GetAddressList())
		walletRoute.GET("/batchAddress", api.BatchGenerateAddress())
	}

	templateRoute := e.Group("/template")
	{
		templateRoute.GET("/login", api.LoginIndex())
		templateRoute.GET("/index", api.Index())
	}

	//验证码
	captchaRoute := e.Group("/captcha")
	{
		captchaRoute.GET("/generateCaptcha")       //生成验证码
		captchaRoute.POST("/verify", api.Verify()) //校验验证码
	}

}
