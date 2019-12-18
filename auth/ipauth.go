package auth

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	. "peak-exchange/model"
	"peak-exchange/service"
	. "peak-exchange/utils"
	"strings"
)

var (
	AllowMethod = []string{"/user/login",
		"/user/register",
		"/user/logout"}
)

// 监测用户IP是否存在异常行为
// 如果存在异常行为 POST PUT DELETE 等操作不予通过
func CheckLoginIp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		allow := true
		method := strings.ToUpper(ctx.Request.Method)
		requestUrl := ctx.Request.RequestURI
		userId := ctx.GetInt("userId")
		ip := ctx.ClientIP()
		log.Println("请求URL: ", requestUrl)
		address := service.SelectAuthLoginAddressByUserId(userId, ip)
		//判断用户操作IP是否在授权表中
		if (AuthLoginAddress{}) == address {
			if method == http.MethodPost ||
				method == http.MethodPut ||
				method == http.MethodDelete {

				//TODO 将来修改为二分查找
				for _, item := range AllowMethod {
					//如果访问地址在授权URL 范围内 则通行
					if item == strings.Split(requestUrl, "v1")[1] {
						allow = true
						break
					} else {
						allow = false
					}
				}
				if !allow {
					log.Println("规则不满足")
					ctx.JSON(http.StatusOK, BuildError(AccessDenied, "暂无权限"))
					ctx.Abort()
					return
				}
			}
		}
		ctx.Next()
	}
}
