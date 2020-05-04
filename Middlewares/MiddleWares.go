package Middlewares

import (
	"apiGateway/DBModels"
	"apiGateway/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 解决跨域的中间件
func HeadMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Add("Access-Control-Allow-Headers", "Content-Type")
		c.Request.Header.Set("content-type", "application/json")
	}
}

// 黑白名单中间件
func IpRestrictionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		remoteIp := Utils.RealIP(c.Request)
		globalIpRestriction := DBModels.IpRestriction{}
		IpBlackList := strings.Split(globalIpRestriction.IpBlackList, ",")
		IpWhiteList := strings.Split(globalIpRestriction.IpWhiteList, ",")
		if Utils.Contain(remoteIp, IpWhiteList) {
			c.Next()
		} else if Utils.Contain(remoteIp, IpBlackList) {
			c.Abort()
			c.JSON(http.StatusForbidden, gin.H{
				"message": "your ip is in ipBlackList.",
			})
		} else {
			c.Next()
		}
	}
}
