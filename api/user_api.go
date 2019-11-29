package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"peak-exchange/auth"
	. "peak-exchange/model"
	"peak-exchange/service"
	. "peak-exchange/utils"
	"strconv"
	"time"
)

func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user User
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, BuildError(ParamError, "参数错误"))
		} else {

			user.Level = "1"
			user.Avatar = "example.png"
			user.UUID = strconv.Itoa(int(time.Now().Unix()))
			user.KycLevel = "0"
			user.CreateAt = time.Now()
			user.UpdateAt = time.Now()
			user.LastLoginAt = time.Now()
			user.LastLoginIp = ctx.ClientIP()
			//校验用户信息
			err = ValidateStruct(user)
			if err != nil {
				ctx.JSON(http.StatusOK, BuildError(ParamError, err.Error()))
				return
			}
			user.LoginPwd = MD5Pwd(user.LoginPwd)

			//创建用户信息
			userId, err := service.Save(user)
			//创建虚拟账户
			if userId != 0 {
				account := CreateVirtualAccount(userId)
				service.SaveAccount(*account)
			}

			if err != nil {
				ctx.JSON(http.StatusOK, BuildError(OperateError, err.Error()))
			} else {
				user.Id = userId
				token, err := generateToken(user)
				if err == nil {
					ctx.JSON(http.StatusOK, Success(token))
				}
			}
		}
	}
}

// 登录
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user User
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, BuildError(ParamError, "参数错误"))
		} else {
			//TODO 校验用户名密码
			retUser := service.SelectUserByMobile(user.Mobile)
			if (User{}) == retUser {
				ctx.JSON(http.StatusOK, BuildError(UserNotFound, "用户不存在"))
				return
			} else {
				encrypt := MD5Pwd(user.LoginPwd)
				if encrypt != retUser.LoginPwd {
					ctx.JSON(http.StatusOK, BuildError(UserNameOrPwdError, "用户名或密码错误"))
					return
				} else {
					retUser.LastLoginIp = ctx.ClientIP()
					retUser.LastLoginAt = time.Now()
					service.UpdateUser(retUser)
					token, _ := generateToken(retUser)
					ctx.JSON(http.StatusOK, Success(token))
				}
			}
		}
	}
}

// 修改用户信息
func UpdateProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// 退出登录
func Logout() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

// 创建token
func generateToken(user User) (string, error) {
	j := auth.NewJwt()
	claims := auth.Claims{
		Mobile: user.Mobile,
		Id:     user.Id,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), //签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), //过期时间一小时
			Issuer:    "peak_exchange",                 //签名发行者
		},
	}

	// 创建token
	token, err := j.CreateToken(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}
