package utils

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// 获取客户端真实IP
func GetRealIP(context gin.Context) (ip string) {
	ips := context.ClientIP()
	realIps := strings.Split(ips, ",")
	ip = realIps[0]
	return
}
