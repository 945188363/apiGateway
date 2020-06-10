package Handlers

import (
	"apiGateway/DBModels"
	"apiGateway/Utils/ComponentUtil"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
)

// Api相关处理
func GetCountList(ginCtx *gin.Context) {
	startTime := ginCtx.Query("startTime")
	endTime := ginCtx.Query("endTime")

	count := DBModels.Count{}
	counts, err := count.GetCountListByData(startTime, endTime)
	if err != nil {
		ComponentUtil.RuntimeLog().Info("fetch count list error")
		ginCtx.JSON(502, gin.H{
			"message": "fetch count list error",
		})
	}
	if len(counts) == 0 {
		ComponentUtil.RuntimeLog().Info("count list do not exist")
		ginCtx.JSON(404, gin.H{
			"message": "count list do not exist",
		})
	}
	ComponentUtil.RuntimeLog().Info("query count list success")
	ginCtx.JSON(200, gin.H{
		"message": "query count list success",
		"data":    counts,
	})
}

func GetCpuInfo(ginCtx *gin.Context) {
	res, err := cpu.Times(false)
	if err != nil {
		ginCtx.JSON(502, gin.H{
			"message": "get cpu info error",
		})
	}
	ginCtx.JSON(200, gin.H{
		"message": "query count list success",
		"data":    ((res[0].Total() - res[0].Idle) / res[0].Total()) * 100,
	})
}
