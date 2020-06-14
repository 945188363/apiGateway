package Routers

import (
	"apiGateway/Core"
	"apiGateway/Handlers"
	"apiGateway/Middlewares"
	"github.com/gin-gonic/gin"
)

func NewGinRouter() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(Middlewares.HeadMiddleware())
	ginRouter.Use(gin.Recovery())
	ginRouter.Use(Middlewares.LoggerMiddleware())
	v1GinGroup := ginRouter.Group("/v1")
	{
		v1GinGroup.GET("/prods", Handlers.GetProdsList)
	}
	gateway := ginRouter.Group("/gateway")
	{
		// api 相关路由
		gateway.POST("/createApiDetail", Handlers.SaveApi)
		gateway.GET("/queryApiDetailList", Handlers.GetApiLst)
		gateway.POST("/queryApiListByApiNameAndGroupName", Handlers.GetApiListByApiNameAndGroupName)
		gateway.POST("/updateApiDetail", Handlers.SaveApi)
		gateway.POST("deleteApiDetail", Handlers.DeleteApi)
		// api group 相关路由
		gateway.POST("/createApiGroupDetail", Handlers.SaveApiGroup)
		gateway.GET("/queryApiGroupList", Handlers.GetApiGroupList)
		gateway.POST("/queryApiGroupListByGroupName", Handlers.GetApiGroupListByGroupName)
		gateway.POST("/updateApiGroupDetail", Handlers.SaveApiGroup)
		gateway.POST("/deleteApiGroupDetail", Handlers.DeleteApiGroup)
		// registry 相关路由
		gateway.POST("/createRegistry", Handlers.SaveRegistry)
		gateway.GET("/queryRegistry", Handlers.GetRegistryList)
		gateway.POST("/queryRegistryByName", Handlers.GetRegistryListByName)
		gateway.POST("/updateRegistry", Handlers.SaveRegistry)
		gateway.POST("/deleteRegistry", Handlers.DeleteRegistry)
		// loadBalance 相关路由
		gateway.POST("/createLoadBalance", Handlers.SaveLoadBalance)
		gateway.GET("/queryLoadBalance", Handlers.GetLoadBalanceList)
		gateway.POST("/queryLoadBalanceByName", Handlers.GetLoadBalanceListByName)
		gateway.POST("/updateLoadBalance", Handlers.SaveLoadBalance)
		gateway.POST("/deleteLoadBalance", Handlers.DeleteLoadBalance)
		// log 相关路由
		gateway.POST("/saveRuntimeLogSetting", Handlers.SaveLogInfo)
		gateway.POST("/saveAccessLogSetting", Handlers.SaveLogInfo)
		gateway.GET("/queryLogSetting", Handlers.GetLogInfoList)
		// monitor 相关路由
		gateway.POST("/saveELKUrl", Handlers.SaveMonitors)
		gateway.POST("/savePrometheus", Handlers.SaveMonitors)
		gateway.GET("/queryMonitors", Handlers.GetMonitors)
		// api 访问统计路由
		gateway.GET("/queryCount", Handlers.GetCountList)
		gateway.GET("/queryCpu", Handlers.GetCpuInfo)
		// api 访问统计路由
		gateway.GET("/queryIpRestriction", Handlers.GetIpRestrictionList)
		gateway.POST("/saveIpRestriction", Handlers.SaveIpRestriction)

	}
	ginRouter.POST("/users", Handlers.GetUser)
	// Core.InitApiMapping(ginRouter)
	// 缓存服务
	Core.InitAllApi(ginRouter)
	return ginRouter
}
