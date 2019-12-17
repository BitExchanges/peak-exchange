package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"peak-exchange/auth"
	"peak-exchange/common"
	. "peak-exchange/model"
	"peak-exchange/service"
	. "peak-exchange/utils"
	"time"
)

// 注册
func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestUser RequestUser
		var user User
		err := ctx.BindJSON(&requestUser)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, BuildError(ParamError, "参数错误"))
		} else {
			//校验用户信息
			err = ValidateStruct(requestUser)
			if err != nil {
				ctx.JSON(http.StatusOK, BuildError(ParamError, err.Error()))
				return
			}
			user.Mobile = requestUser.Mobile
			user.Email = requestUser.Email
			//校验验证码
			err = VerifyCaptcha(requestUser.Id, requestUser.CaptchaCode)
			if err != nil {
				ctx.JSON(http.StatusOK, BuildError(CaptchaError, "验证码错误"))
				return
			}

			user.Init()
			user.LastLoginIp = ctx.ClientIP()
			user.LoginPwd = MD5Pwd(user.LoginPwd)
			user.Device = ctx.GetString("device")
			user, err := service.UserRegister(user)

			if err != nil {
				ctx.JSON(http.StatusOK, BuildError(OperateError, err.Error()))
				return
			}

			//TODO 发送激活邮件 需要设置有效期
			var urlStr = fmt.Sprintf(common.ActiveUrl, user.RandomUUID)
			go user.SendEmail1(urlStr)
			ctx.JSON(http.StatusOK, Success("注册成功"))
		}
	}
}

// 激活用户
func ActiveUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := ctx.Query("uuid")

		if uuid == "" {
			ctx.JSON(http.StatusOK, BuildError(ParamError, "参数错误"))
			return
		}
		service.UpdateUserActive(uuid)
		ctx.JSON(http.StatusOK, Success("用户激活成功"))
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

			retUser := service.SelectUserByMobile(user.Mobile)
			if (User{}) == retUser {
				ctx.JSON(http.StatusOK, BuildError(UserNotFound, "用户不存在"))
				return
			} else if retUser.State == 0 {
				ctx.JSON(http.StatusOK, BuildError(NotActive, "请先激活用户"))
				return
			} else {
				encrypt := MD5Pwd(user.LoginPwd)
				if encrypt != retUser.LoginPwd {
					ctx.JSON(http.StatusOK, BuildError(UserNameOrPwdError, "用户名或密码错误"))
					return
				} else {

					//检查登录IP是否在常用地址内
					loginAddress := service.SelectAuthLoginAddressByUserId(retUser.ID, ctx.ClientIP())
					if (AuthLoginAddress{}) == loginAddress {
						if retUser.Email != "" {
							go retUser.SendEmail(0, ctx.ClientIP())
						}
					}
					retUser.LastLoginIp = ctx.ClientIP()
					service.UpdateUser(retUser)
					token, _ := generateToken(retUser)
					retUser.Token = token
					ctx.JSON(http.StatusOK, Success(retUser))
				}
			}
		}
	}
}

// 退出登录
func Logout() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

// 修改用户信息
func UpdateProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

//忘记密码
func ForgetPwd() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user User
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusOK, BuildError(ParamError, "参数错误"))
			return
		}
		if user.Mobile == "" || user.Email == "" {
			ctx.JSON(http.StatusOK, BuildError(ParamError, "参数错误"))
			return
		}
		if user.Mobile != "" {
			fmt.Println("开始发送手机短信")

		}

		if user.Email != "" {
			captchaCode := GenerateCode(4)
			go user.SendEmail1(captchaCode)
			re, err := LimitPool.Get().Do("SET", fmt.Sprintf(common.RedisEmailFormatter, user.Email), captchaCode, "EX", "120")
			if err != nil {
				fmt.Println("redis写入错误: ", err)
			}
			fmt.Println("reply: ", re)
		}
	}
}

// 修改登录密码
func ChangeLoginPwd() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

//修改交易密码
func ChangeTradePwd() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// 创建token
func generateToken(user User) (string, error) {
	j := auth.NewJwt()
	claims := auth.Claims{
		Mobile:   user.Mobile,
		Id:       user.ID,
		LoginPwd: user.LoginPwd,
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
