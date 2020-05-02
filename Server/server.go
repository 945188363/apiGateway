package Server

import (
	"apiGateway/Config"
	"apiGateway/Routers"
	"github.com/gin-gonic/gin"
)

// 配置并启动服务
func Run() {
	// gin 运行时 release debug test
	gin.SetMode(Config.Env)

	// 注册路由
	router := Routers.NewGinRouter()

	addr := Config.Host + ":" + Config.Port
	// 启动服务
	err := router.Run(addr)

	if nil != err {
		panic("server run error: " + err.Error())
	}
}