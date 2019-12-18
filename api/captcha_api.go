package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	. "peak-exchange/model"
	. "peak-exchange/utils"
)

// 校验验证码
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

// 创建验证码
func CreateCaptcha() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var captchaTyp = ctx.Query("captchaType")
		log.Println("打印客户端IP地址: ", ctx.ClientIP())
		if captchaTyp == "" {
			ctx.JSON(http.StatusOK, BuildError(ParamError, "参数错误"))
			return
		}

		id, data, err := GenerateCaptcha(captchaTyp)
		if err != nil {
			ctx.JSON(http.StatusOK, BuildError(OperateError, "操作失败"))
			return
		}
		var captcha Captcha
		captcha.Id = id
		captcha.Value = data
		ctx.JSON(http.StatusOK, Success(captcha))
	}
}
