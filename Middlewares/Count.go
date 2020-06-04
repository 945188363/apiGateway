package Middlewares

import (
	"apiGateway/DBModels"
	"github.com/gin-gonic/gin"
	"strings"
)

type CountMw struct {
	DBModels.Api
}

func (mw *CountMw) CountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取当前API信息
		apiN := DBModels.Api{ApiUrl: handleUriGroup(c.Request.URL.Path)}
		_ = apiN.GetApiByUrl()
		mw.Api = apiN
		c.Set("ApiInfo", mw.Api)

		// 保存访问信息
		count := DBModels.Count{
			ApiName: mw.ApiName,
		}
		count.SaveCount()

		// 下一步
		c.Next()
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
