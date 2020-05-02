package Routers

import (
	"apiGateway/Handlers"
	"github.com/gin-gonic/gin"
)

func NewGinRouter() *gin.Engine {
	ginRouter := gin.Default()
	v1GinGroup := ginRouter.Group("/v1")
	{
		v1GinGroup.GET("/prods", Handlers.GetProdsList)
	}
	ginRouter.POST( "/users", Handlers.GetUser)

	return ginRouter
}
