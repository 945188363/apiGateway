package Handlers

import (
	"apiGateway/DBModels"
	"apiGateway/Utils/DataUtil"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type LoadBalance struct {
	Name         string `form:"Name"`
	RegistryName string `form:"RegistryName"`
	Strategy     string `form:"Strategy"`
	ServiceName  string `form:"ServiceName"`
}

func (a LoadBalance) IsEmpty() bool {
	return reflect.DeepEqual(a, LoadBalance{})
}

// Api相关处理
func GetLoadBalanceList(ginCtx *gin.Context) {
	var loadBalanceModel DBModels.LoadBalance
	loadBalanceList, err := loadBalanceModel.GetLoadBalanceList()
	if err != nil {
		ginCtx.JSON(502, gin.H{
			"message": "fetch loadBalance list error",
		})
	}
	if len(loadBalanceList) == 0 {
		ginCtx.JSON(404, gin.H{
			"message": "loadBalance list do not exist",
		})
	}
	ginCtx.JSON(200, gin.H{
		"message": "query loadBalance list error",
		"data":    loadBalanceList,
	})
}

func SaveLoadBalance(ginCtx *gin.Context) {
	var loadBalance LoadBalance
	ginCtx.Bind(&loadBalance)
	var loadBalanceModel DBModels.LoadBalance
	DataUtil.CopyFields(&loadBalanceModel, loadBalance,
		"Name",
		"RegistryName",
		"Strategy",
		"ServiceName")
	saveRegistry := loadBalanceModel.SaveLoadBalance()
	if saveRegistry {
		ginCtx.JSON(200, gin.H{
			"message": "save registry success",
		})
	} else {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error!",
		})
	}
}

func DeleteLoadBalance(ginCtx *gin.Context) {
	var loadBalance LoadBalance
	ginCtx.Bind(&loadBalance)
	var loadBalanceModel DBModels.LoadBalance
	DataUtil.CopyFields(&loadBalanceModel, loadBalance,
		"Name",
		"RegistryName",
		"Strategy",
		"ServiceName")
	delLoadBalance := loadBalanceModel.DeleteLoadBalance()
	if delLoadBalance {
		ginCtx.JSON(200, gin.H{
			"message": "delete loadBalance success",
		})
	} else {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error!",
		})
	}
}
