package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"peak-exchange/common"
	"peak-exchange/erc20"
	"peak-exchange/model"
	"peak-exchange/service"
	"peak-exchange/utils"
	"strconv"
	"time"
)

// 查询用户钱包列表
func GetWallet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//userId := ctx.GetInt("userId")
		//limit, page := common.LimitAndPage(ctx.Query("limit"), ctx.Query("page"))

	}
}

//批量生成地址
func BatchGenerateAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		countStr := ctx.Query("count")
		count, _ := strconv.Atoi(countStr)

		for i := 0; i < count; i++ {
			key, add := erc20.GenerateUserWallet()
			var wallet model.Wallet
			wallet.Address = add
			wallet.PrivateKey = key
			wallet.CreateAt = time.Now()
			wallet.UpdateAt = time.Now()
			wallet.Type = "erc20"
			wallet.Currency = "usd"
			service.BatchSaveAddress(wallet)
		}

		ctx.JSON(http.StatusOK, utils.Success("批处理地址生成完毕"))
	}
}

//查询地址列表
func GetAddressList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limit, page := common.LimitAndPage(ctx.Query("limit"), ctx.Query("page"))
		wallets, count, err := service.SelectAddressByPage(limit, page)
		if err != nil {
			ctx.JSON(http.StatusOK, utils.BuildError(utils.NotFound, "暂无数据"))
			return
		}
		var pageResponse utils.PageResponse
		ctx.JSON(http.StatusOK, pageResponse.PageInit(wallets, page, count, limit))
	}
}
