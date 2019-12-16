package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "peak-exchange/model"
	. "peak-exchange/utils"
)

func Verify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var captcha Captcha

		err := ctx.BindJSON(&captcha)
		if err != nil {
			ctx.JSON(http.StatusOK, BuildError(ParamError, "参数错误"))
			return
		}
		err = VerifyCaptcha(captcha.Id, captcha.Value)
		if err != nil {
			ctx.JSON(http.StatusOK, BuildError(CaptchaError, "验证码错误"))
			return
		}
		ctx.JSON(http.StatusOK, Success(""))
	}
}
