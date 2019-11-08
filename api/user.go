package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"peak-exchange/auth"
	. "peak-exchange/model"
	"peak-exchange/utils"
	"time"
)

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user User
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildError("10001").Error())
		}
	}
}

func generateToken(ctx *gin.Context, user User) {
	j := auth.NewJwt()
	claims := auth.Claims{
		Mobile: user.Mobile,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), //签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), //过期时间一小时
			Issuer:    "peak_exchange",                 //签名发行者
		},
	}

	// 创建token
	token, err := j.CreateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.BuildError("10002").Error())
	}
}
