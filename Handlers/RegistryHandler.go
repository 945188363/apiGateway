package Handlers

import (
	"apiGateway/DBModels"
	"apiGateway/Utils/ComponentUtil"
	"apiGateway/Utils/DataUtil"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Registry struct {
	Name         string `form:"Name"`
	RegistryType string `form:"RegistryType"`
	Addr         string `form:"Addr"`
}

func (a Registry) IsEmpty() bool {
	return reflect.DeepEqual(a, Registry{})
}

// Api相关处理
func GetRegistryList(ginCtx *gin.Context) {
	var registryModel DBModels.Registry
	registryList, err := registryModel.GetRegistryList()
	if err != nil {
		ComponentUtil.RuntimeLog().Info("fetch registry list error")
		ginCtx.JSON(502, gin.H{
			"message": "fetch registry list error",
		})
	}
	if len(registryList) == 0 {
		ComponentUtil.RuntimeLog().Info("registry list do not exist")
		ginCtx.JSON(404, gin.H{
			"message": "registry list do not exist",
		})
	}
	ComponentUtil.RuntimeLog().Info("query registry list success")
	ginCtx.JSON(200, gin.H{
		"message": "query registry list success",
		"data":    registryList,
	})
}

func SaveRegistry(ginCtx *gin.Context) {
	var registry Registry
	ginCtx.Bind(&registry)
	var registryModel DBModels.Registry
	DataUtil.CopyFields(&registryModel, registry,
		"Name",
		"RegistryType",
		"Addr")
	ComponentUtil.RuntimeLog().Info("transfer data to Model :", registryModel)
	saveRegistry := registryModel.SaveRegistry()
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

func DeleteRegistry(ginCtx *gin.Context) {
	var registry Registry
	ginCtx.Bind(&registry)
	var registryModel DBModels.Registry
	DataUtil.CopyFields(&registryModel, registry,
		"Name",
		"RegistryType",
		"Addr")
	ComponentUtil.RuntimeLog().Info("transfer data to Model :", registryModel)
	delRegistry := registryModel.DeleteRegistry()
	if delRegistry {
		ComponentUtil.RuntimeLog().Info("delete registry success")
		ginCtx.JSON(200, gin.H{
			"message": "delete registry success",
		})
	} else {
		ComponentUtil.RuntimeLog().Info("internal server error!")
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error!",
		})
	}
}
