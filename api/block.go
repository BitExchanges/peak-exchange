package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	. "peak-exchange/utils"
)

// 查询区块头
func GetCurrentBlockHead() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header, err := EthClient.HeaderByNumber(context.Background(), nil)
		if err != nil {
			ctx.JSON(http.StatusOK, BuildError(BlockError, "区块错误"))
			return
		}
		ctx.JSON(http.StatusOK, Success(header))
	}
}
