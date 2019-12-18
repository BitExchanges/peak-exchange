package routes

import (
	"github.com/gin-gonic/gin"
	"peak-exchange/api"
	"peak-exchange/auth"
)

func SetInterfaces(e *gin.Engine) {

	//全局解决跨域问题
	e.Use(auth.Cross())
	// 订单组
	orderRoute := e.Group("/api/:platform/v1/order")
	orderRoute.Use(auth.Authorize())
	{
		orderRoute.GET("/getOrderBook", api.GetOrderBook())             //查询所有订单簿
		orderRoute.GET("/getOrderByNo/:orderNo", api.GetOrderByNo())    //根据订单号查询详细
		orderRoute.GET("/getOrderByUserId", api.GetOrderBookByUserId()) //查询用户订单簿
		orderRoute.POST("/maker", api.Maker())                          //挂单
		orderRoute.POST("/cancelOrder/:orderNo")                        //撤单

	}

	//账户组
	accountRoute := e.Group("/api/:platform/v1/account")
	{
		accountRoute.GET("/balance/:type/:currency/") //查询账户余额
		accountRoute.POST("/recharge/:currency")      //账户充值
		accountRoute.POST("/transfer/:")              //转账
	}

	//币组
	currencyRoute := e.Group("/api/:platform/v1/currency")
	currencyRoute.Use(auth.GetDevice())
	{
		currencyRoute.GET("/currencyList", api.GetCurrencyList()) //查询当前交易对
	}

	//用户组
	userRoute := e.Group("/api/:platform/v1/user")
	userRoute.Use(auth.GetDevice())
	userRoute.Use(auth.CheckLoginIp())
	{
		userRoute.POST("/register", api.Register())        //注册
		userRoute.POST("/login", api.Login())              //登录
		userRoute.POST("/logout", api.Logout())            //退出登录  TODO 暂时没有想好如何处理jwt问题
		userRoute.POST("/forgetLoginPwd", api.ForgetPwd()) //忘记登录密码
		userRoute.GET("/active", api.ActiveUser())         //激活
		userRoute.Use(auth.Authorize())
		userRoute.POST("/updateProfile", api.UpdateProfile())   //更新个人资料
		userRoute.POST("/changeLoginPwd", api.ChangeLoginPwd()) //修改登录密码
		userRoute.POST("/changeTradePwd", api.ChangeTradePwd()) //修改交易密码

	}
	//杂项组
	miscRoute := e.Group("/api/:platform/v1/misc")
	miscRoute.Use(auth.GetDevice())
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

	//验证码
	captchaRoute := e.Group("api/:platform/v1/captcha")
	captchaRoute.Use(auth.GetDevice())
	{
		captchaRoute.GET("/generateCaptcha", api.CreateCaptcha()) //生成验证码
		captchaRoute.POST("/verify", api.Verify())                //校验验证码
	}

	//模板测试组
	templateRoute := e.Group("/template")
	{
		templateRoute.GET("/login", api.LoginIndex())
		templateRoute.GET("/index", api.Index())
		templateRoute.GET("/email", api.EmailTemplate())

	}

}
