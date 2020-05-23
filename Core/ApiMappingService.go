package Core

import (
	"apiGateway/DBModels"
	"apiGateway/Middlewares"
	"apiGateway/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"sort"
)

// 初始化路由
func InitApiMapping(router *gin.Engine) {
	// 鉴权
	auth := Middlewares.AuthMw{}
	router.Use(auth.BasicAuthMiddleware())
	// 限流
	ratelimit := Middlewares.RateLimiterMw{}
	// 熔断
	breaker := Middlewares.BreakerMw{}
	// 黑白名单
	ipRestriction := Middlewares.IpRestrictionMw{}
	router.Use(ipRestriction.GlobalIpRestrictionMiddleware())
	api := DBModels.Api{}
	apiList, err := api.GetApiList()
	// 根据apiGroup分组
	if err != nil {
		Utils.RuntimeLog().Info("get api list error", err)
	}
	apiListGroup := splitByGroup(apiList)
	var httpInvoker HttpInvoker
	for i := 0; i < len(apiListGroup); i++ {
		for j := 0; j < len(apiListGroup[i]); j++ {
			ratelimit.Api = apiListGroup[i][j]
			breaker.Api = apiListGroup[i][j]
			httpInvoker.Api = apiListGroup[i][j]
			if ratelimit.RateLimiterNum > 0 {
				router.Any(handleUrl(apiListGroup[i][j]), ratelimit.RateLimitMiddleware(), breaker.CircuitBreakerMiddleware(), httpInvoker.Execute)
			} else {
				router.Any(handleUrl(apiListGroup[i][j]), breaker.CircuitBreakerMiddleware(), httpInvoker.Execute)
			}
		}
	}
}

// URL处理
func handleUrl(api DBModels.Api) string {
	var url string
	// 处理URI
	if api.ApiGroup != "" {
		url = fmt.Sprintf("/%s/%s", api.ApiGroup, api.ApiUrl)
	} else {
		url = api.ApiUrl
	}

	return url
}

type sortList []DBModels.Api

func (s sortList) Len() int           { return len(s) }
func (s sortList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortList) Less(i, j int) bool { return s[i].ApiGroup < s[j].ApiGroup }

// 分组聚合
func splitByGroup(apiList []DBModels.Api) [][]DBModels.Api {
	sort.Sort(sortList(apiList))
	returnData := make([][]DBModels.Api, 0)
	i := 0
	var j int
	for {
		if i >= len(apiList) {
			break
		}
		for j = i + 1; j < len(apiList) && apiList[i].ApiGroup == apiList[j].ApiGroup; j++ {
		}
		returnData = append(returnData, apiList[i:j])
		i = j
	}
	return returnData
}
