package Core

import (
	"apiGateway/Config"
	"apiGateway/DBModels"
	"apiGateway/Middlewares"
	"apiGateway/Utils/ComponentUtil"
	"github.com/gin-gonic/gin"
)

func InitAllApi(router *gin.Engine) {
	// 黑白名单
	ipRestriction := Middlewares.IpRestrictionMw{}
	router.Use(ipRestriction.GlobalIpRestrictionMiddleware())
	// 鉴权
	auth := Middlewares.AuthMw{}
	router.Use(auth.BasicAuthMiddleware())
	// 缓存
	cache := Middlewares.CacheMw{}
	router.Use(cache.CacheMiddleware())
	// 参数校验
	paramCheck := Middlewares.ParameterCheckMw{}
	router.Use(paramCheck.ParameterCheckMiddleware())
	// 限流
	rateLimit := Middlewares.RateLimiterMw{}
	// 熔断
	breaker := Middlewares.BreakerMw{}
	// 访问统计
	count := Middlewares.CountMw{}

	router.Any("api/*uri",
		count.CountMiddleware(),
		rateLimit.RateLimitMiddleware(),
		breaker.CircuitBreakerMiddleware(),
		ExecuteAll)

}

func ExecuteAll(ginCtx *gin.Context) {

	// 获取api信息
	value, exists := ginCtx.Get("ApiInfo")
	if !exists {
		ComponentUtil.RuntimeLog().Info("api info is null.")
		return
	}
	ComponentUtil.RuntimeLog().Info("match api protocolType to execute.")
	switch value.(DBModels.Api).ProtocolType {
	case Config.Http:
		httpInvoker := HttpInvoker{}
		httpInvoker.Api = value.(DBModels.Api)
		httpInvoker.Execute(ginCtx)
		break
	case Config.GRPC:
		rpcInvoker := RpcInvoker{}
		rpcInvoker.Api = value.(DBModels.Api)
		rpcInvoker.Execute(ginCtx)
		break
	}

}
