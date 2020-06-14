package Middlewares

import (
	"apiGateway/DBModels"
	"apiGateway/HandlersCache"
	"apiGateway/Utils/ComponentUtil"
	"apiGateway/Utils/DataUtil"
	"github.com/gin-gonic/gin"
	"strings"
)

type CacheMw struct {
	DBModels.Api
}

func (mw *CacheMw) CacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ComponentUtil.RuntimeLog().Info("start Cache MiddleWare...")

		// 如果是网关操作的接口直接放行
		if url := c.Request.URL.String(); strings.HasPrefix(url, "/gateway") {
			c.Next()
			return
		}
		// 获取截取group后的uri
		url := handleUriGroup(c.Request.URL.Path)
		// 查询缓存是否存在
		catchApi := HandlersCache.GetCatchApi(url)
		// 缓存不存在进行下一步
		if catchApi.IsEmpty() {
			// 从数据库获取当前请求API信息
			apiN := DBModels.Api{ApiUrl: url}
			err := apiN.GetApiByUrl()
			// 数据库也没有就终止访问
			if err != nil {
				c.Abort()
				ComponentUtil.RuntimeLog().Info("this api is NOT registry in apiGateway.")
				return
			}
			mw.Api = apiN
			// 保证缓存和数据库的数据一致性
			HandlersCache.SaveCatchApi(apiN)
		} else {
			_ = DataUtil.CopyFields(&mw.Api, catchApi)
		}
		c.Set("ApiInfo", mw.Api)
		c.Next()
		ComponentUtil.RuntimeLog().Info("end Cache MiddleWare...")
	}
}
