package Routers

import (
	"apiGateway/Handlers"
	"apiGateway/Middlewares"
	"github.com/gin-gonic/gin"
)

func NewGinRouter() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(Middlewares.HeadMiddleware())
	ginRouter.Use(gin.Recovery())
	v1GinGroup := ginRouter.Group("/v1")
	{
		v1GinGroup.GET("/prods", Handlers.GetProdsList)
	}
	gateway := ginRouter.Group("/gateway")
	{
		// api 相关路由
		gateway.POST("/createApiDetail", Handlers.SaveApi)
		gateway.GET("/queryApiDetailList", Handlers.GetApiLst)
		gateway.POST("/updateApiDetail", Handlers.SaveApi)
		gateway.DELETE("deleteApiDetail", Handlers.DeleteApi)
		// api group 相关路由
		gateway.POST("/createApiGroupDetail", Handlers.SaveApiGroup)
		gateway.GET("/queryApiGroupList", Handlers.GetApiGroupList)
		gateway.POST("/updateApiGroupDetail", Handlers.SaveApiGroup)
		gateway.DELETE("/deleteApiGroupDetail", Handlers.DeleteApiGroup)
	}
	ginRouter.POST("/users", Handlers.GetUser)

	return ginRouter
}
