package Handlers

import (
	"apiGateway/DBModels"
	"github.com/gin-gonic/gin"
)

// Api相关处理
func GetCountList(ginCtx *gin.Context) {
	startTime := ginCtx.Query("startTime")
	endTime := ginCtx.Query("endTime")

	count := DBModels.Count{}
	counts, err := count.GetCountListByData(startTime, endTime)
	if err != nil {
		ginCtx.JSON(502, gin.H{
			"message": "fetch count list error",
		})
	}
	if len(counts) == 0 {
		ginCtx.JSON(404, gin.H{
			"message": "count list do not exist",
		})
	}
	ginCtx.JSON(200, gin.H{
		"message": "query count list error",
		"data":    counts,
	})
}
