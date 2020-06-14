package Middlewares

import (
	"apiGateway/DBModels"
	"apiGateway/Utils/ComponentUtil"
	"github.com/gin-gonic/gin"
	"strings"
)

type CountMw struct {
	DBModels.Api
}

func (mw *CountMw) CountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ComponentUtil.RuntimeLog().Info("start Count MiddleWare...")
		// 获取APi信息
		api, exists := c.Get("ApiInfo")
		if !exists {
			ComponentUtil.RuntimeLog().Info("api info is null.")
			c.Abort()
			return
		}

		mw.Api = api.(DBModels.Api)
		// 保存访问信息
		count := DBModels.Count{
			ApiName: mw.ApiName,
		}
		count.SaveCount()

		// 下一步
		c.Next()
		ComponentUtil.RuntimeLog().Info("end Count MiddleWare...")
	}
}

// 处理带组的Uri
func handleUriGroup(uri string) string {
	group := DBModels.ApiGroup{}
	groupList, _ := group.GetApiGroupList()
	for _, group := range groupList {
		if strings.Contains(uri, group.ApiGroupName) {
			return strings.ReplaceAll(uri, "/"+group.ApiGroupName+"/", "")
		}
	}
	return strings.Replace(uri, "/", "", 1)
}
