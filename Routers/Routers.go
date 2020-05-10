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
		gateway.POST("/createApiDetail", Handlers.SaveApi)
		gateway.GET("/queryApiDetailList", Handlers.GetApiLst)
	}
	ginRouter.POST("/users", Handlers.GetUser)

	return ginRouter
}
