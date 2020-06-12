package Handlers

import (
	"apiGateway/DBModels"
	"apiGateway/Utils/ComponentUtil"
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
		ComponentUtil.RuntimeLog().Info("fetch loadBalance list error")
		ginCtx.JSON(502, gin.H{
			"message": "fetch loadBalance list error",
		})
	}
	if len(loadBalanceList) == 0 {
		ComponentUtil.RuntimeLog().Info("loadBalance list do not exist")
		ginCtx.JSON(404, gin.H{
			"message": "loadBalance list do not exist",
		})
	}
	ComponentUtil.RuntimeLog().Info("query loadBalance list success")
	ginCtx.JSON(200, gin.H{
		"message": "query loadBalance list success",
		"data":    loadBalanceList,
	})
}

// Api相关处理
func GetLoadBalanceListByName(ginCtx *gin.Context) {
	var loadBalance LoadBalance
	ginCtx.ShouldBind(&loadBalance)
	var loadBalanceModel DBModels.LoadBalance
	DataUtil.CopyFields(&loadBalanceModel, loadBalance,
		"Name",
		"RegistryName",
		"Strategy",
		"ServiceName")
	ComponentUtil.RuntimeLog().Info("transfer data to Model :", loadBalanceModel)
	loadBalanceList, err := loadBalanceModel.GetLoadBalanceListByName()
	if err != nil {
		ComponentUtil.RuntimeLog().Info("fetch loadBalance list error")
		ginCtx.JSON(502, gin.H{
			"message": "fetch loadBalance list error",
		})
	}
	if len(loadBalanceList) == 0 {
		ComponentUtil.RuntimeLog().Info("loadBalance list do not exist")
		ginCtx.JSON(404, gin.H{
			"message": "loadBalance list do not exist",
		})
	}
	ComponentUtil.RuntimeLog().Info("query loadBalance list success")
	ginCtx.JSON(200, gin.H{
		"message": "query loadBalance list success",
		"data":    loadBalanceList,
	})
}

func SaveLoadBalance(ginCtx *gin.Context) {
	var loadBalance LoadBalance
	ginCtx.ShouldBind(&loadBalance)
	var loadBalanceModel DBModels.LoadBalance
	DataUtil.CopyFields(&loadBalanceModel, loadBalance,
		"Name",
		"RegistryName",
		"Strategy",
		"ServiceName")
	ComponentUtil.RuntimeLog().Info("transfer data to Model :", loadBalanceModel)
	saveRegistry := loadBalanceModel.SaveLoadBalance()
	if saveRegistry {
		ComponentUtil.RuntimeLog().Info("save registry success")
		ginCtx.JSON(200, gin.H{
			"message": "save registry success",
		})
	} else {
		ComponentUtil.RuntimeLog().Info("internal server error!")
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error!",
		})
	}
}

func DeleteLoadBalance(ginCtx *gin.Context) {
	var loadBalance LoadBalance
	ginCtx.ShouldBind(&loadBalance)
	var loadBalanceModel DBModels.LoadBalance
	DataUtil.CopyFields(&loadBalanceModel, loadBalance,
		"Name",
		"RegistryName",
		"Strategy",
		"ServiceName")
	ComponentUtil.RuntimeLog().Info("transfer data to Model :", loadBalanceModel)
	delLoadBalance := loadBalanceModel.DeleteLoadBalance()
	if delLoadBalance {
		ComponentUtil.RuntimeLog().Info("delete loadBalance success")
		ginCtx.JSON(200, gin.H{
			"message": "delete loadBalance success",
		})
	} else {
		ComponentUtil.RuntimeLog().Info("internal server error!")
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error!",
		})
	}
}
