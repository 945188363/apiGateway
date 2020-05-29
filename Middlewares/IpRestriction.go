package Middlewares

import (
	"apiGateway/Constant/Code"
	"apiGateway/Constant/Message"
	"apiGateway/Core/Domain"
	"apiGateway/DBModels"
	"apiGateway/Utils"
	"apiGateway/Utils/DataUtil"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type IpRestrictionMw struct {
	DBModels.Api
}

// 黑白名单中间件
func (mw *IpRestrictionMw) GlobalIpRestrictionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		remoteIp := Utils.RealIP(c.Request)
		globalIpRestriction := DBModels.IpRestriction{}
		err := globalIpRestriction.GetGlobalIpRestriction()
		if err != nil {
			return
		}
		IpBlackList := strings.Split(globalIpRestriction.IpBlackList, ",")
		IpWhiteList := strings.Split(globalIpRestriction.IpWhiteList, ",")

		// 未设置黑白名单直接放行
		if len(IpWhiteList) == 0 && len(IpBlackList) == 0 {
			c.Next()
			return
		}
		// 设置白名单后，且在白名单里，请求放行
		if DataUtil.Contain(remoteIp, IpWhiteList) && len(IpWhiteList) > 0 {
			c.Next()
			return
		}
		// 设置黑名单后，且在黑名单里，则阻止访问
		if DataUtil.Contain(remoteIp, IpBlackList) && len(IpBlackList) > 0 {
			c.JSON(http.StatusForbidden, Domain.Message{
				Code: Code.IP_FORBIDDEN,
				Msg:  Message.IP_FORBIDDEN,
				Data: nil,
			})
			c.Abort()
		}
	}
}

// 黑白名单中间件
func (mw *IpRestrictionMw) IpRestrictionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		remoteIp := Utils.RealIP(c.Request)
		globalIpRestriction := DBModels.IpRestriction{}
		err := globalIpRestriction.GetIpRestrictionByApi(mw.ApiName)
		if err != nil {
			return
		}
		IpBlackList := strings.Split(globalIpRestriction.IpBlackList, ",")
		IpWhiteList := strings.Split(globalIpRestriction.IpWhiteList, ",")

		// 未设置黑名单直接放行
		if len(IpWhiteList) == 0 && len(IpBlackList) == 0 {
			c.Next()
			return
		}
		// 设置白名单后，且在白名单里，请求放行
		if DataUtil.Contain(remoteIp, IpWhiteList) && len(IpWhiteList) > 0 {
			c.Next()
			return
		}
		// 设置黑名单后，且在黑名单里，则阻止访问
		if DataUtil.Contain(remoteIp, IpBlackList) && len(IpBlackList) > 0 {
			c.JSON(http.StatusForbidden, Domain.Message{
				Code: Code.IP_FORBIDDEN,
				Msg:  Message.IP_FORBIDDEN,
				Data: nil,
			})
			c.Abort()
		}
	}
}
