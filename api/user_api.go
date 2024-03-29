package api

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
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

			if requestUser.Id == "" || requestUser.CaptchaCode == "" {
				ctx.JSON(http.StatusOK, BuildError(CaptchaError, "验证码错误"))
				return
			}
			//校验验证码
			err = VerifyCaptcha(requestUser.Id, requestUser.CaptchaCode)
			if err != nil {
				ctx.JSON(http.StatusOK, BuildError(CaptchaError, "验证码错误"))
				return
			}

			user.Init()
			user.LastLoginIp = ctx.ClientIP()
			user.LoginPwd = MD5Pwd(requestUser.LoginPwd)
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
		var requestUser RequestUser
		var retUser User
		err := ctx.BindJSON(&requestUser)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, BuildError(ParamError, "参数错误"))
		} else {

			//TODO 登录暂时不校验 验证码
			//if requestUser.Id != "" && requestUser.CaptchaCode != "" {
			//	err := VerifyCaptcha(requestUser.Id, requestUser.CaptchaCode)
			//	if err != nil {
			//		ctx.JSON(http.StatusOK, BuildError(CaptchaError, "验证码错误"))
			//		return
			//	}
			//} else if requestUser.Id == "" || requestUser.CaptchaCode == "" {
			//	ctx.JSON(http.StatusOK,BuildError(ParamError,"参数错误"))
			//	return
			//}
			if requestUser.LoginType == "mobile" {
				retUser = service.SelectUserByMobile(requestUser.Mobile)
			} else if requestUser.LoginType == "email" {
				retUser = service.SelectUserByEmail(requestUser.Email)
			} else {
				ctx.JSON(http.StatusOK, BuildError(ParamError, "参数错误"))
				return
			}

			if (User{}) == retUser {
				ctx.JSON(http.StatusOK, BuildError(UserNotFound, "用户不存在"))
				return
			} else if retUser.State == 0 {
				ctx.JSON(http.StatusOK, BuildError(NotActive, "请先激活用户"))
				return
			} else {
				encrypt := MD5Pwd(requestUser.LoginPwd)
				if encrypt != retUser.LoginPwd {
					ctx.JSON(http.StatusOK, BuildError(UserNameOrPwdError, "用户名或密码错误"))
					return
				} else {

					//检查登录IP是否在常用地址内
					loginAddress := service.SelectAuthLoginAddressByUserId(retUser.ID, ctx.ClientIP())
					if (AuthLoginAddress{}) == loginAddress {

						//如果登录地IP不在授权范围内  需要发送确认邮件进行授权
						//在未授权之前无法操作资金账户 修改个人资料及其他
						if retUser.Email != "" {
							log.Println("登陆地IP不在授权地址列表，需要发送邮件确认")
							//TODO 后续需要打开邮件发发送开关
							//go retUser.SendEmail(0, ctx.ClientIP())
						}
					}
					retUser.LastLoginIp = ctx.ClientIP()
					service.UpdateUser(retUser)
					token, _ := generateToken(retUser)
					retUser.Token = token
					var responseUser ResponseUser
					userBytes, _ := json.Marshal(retUser)
					json.Unmarshal(userBytes, &responseUser)
					ctx.JSON(http.StatusOK, Success(responseUser))
				}
			}
		}
	}
}

// 退出登录
func Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwt := auth.NewJwt()
		token := ctx.GetHeader("token")
		jwt.RefreshToken(token)
		ctx.JSON(http.StatusOK, Success("退出登录成功"))
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
		redisCon := LimitPool.Get()
		var requestUser RequestUser
		err := ctx.BindJSON(&requestUser)
		if err != nil {
			ctx.JSON(http.StatusOK, BuildError(ParamError, "参数错误"))
			return
		}
		if requestUser.Email == "" {
			ctx.JSON(http.StatusOK, BuildError(ParamError, "参数错误"))
			return
		}
		if requestUser.LoginPwd != requestUser.ConfirmLoginPwd {
			ctx.JSON(http.StatusOK, BuildError(PasswordDisagreement, "密码不一致"))
			return
		}

		user := service.SelectUserByEmail(requestUser.Email)
		if (User{}) == user {
			ctx.JSON(http.StatusOK, BuildError(UserNotFound, "用户不存在"))
			return
		}

		reply, err := redisCon.Do("GET", fmt.Sprintf(common.RedisEmailFormatter, requestUser.Email, common.ForgetPwdKey))
		if err != nil {
			ctx.JSON(http.StatusOK, BuildError(SystemError, "系统错误"))
			return
		}

		if reply.(string) != requestUser.CaptchaCode {
			ctx.JSON(http.StatusOK, BuildError(CaptchaError, "验证码错误"))
			return
		}

		encrptPwd := MD5Pwd(requestUser.LoginPwd)
		service.UpdateLoginPwdByEmail(requestUser.Email, encrptPwd)
		ctx.JSON(http.StatusOK, Success("修改登录密码成功"))
	}
}

// 修改登录密码
func ChangeLoginPwd() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userId = ctx.GetInt("userId")
		var requestUser RequestUser
		err := ctx.BindJSON(&requestUser)
		if err != nil {
			ctx.JSON(http.StatusOK, BuildError(ParamError, "参数错误"))
			return
		}

		//TODO 需要手机验证码  或邮箱验证码
		service.UpdateUserLoginPwd(userId, requestUser.TradePwd)
	}
}

//修改交易密码
func ChangeTradePwd() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userId = ctx.GetInt("userId")
		var requestUser RequestUser
		err := ctx.BindJSON(&requestUser)
		if err != nil {
			ctx.JSON(http.StatusOK, BuildError(ParamError, "参数错误"))
			return
		}
		//TODO 需要手机验证码  或邮箱验证码
		service.UpdateUserTradePwd(userId, requestUser.TradePwd)
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
