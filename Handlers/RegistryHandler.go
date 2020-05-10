package Handlers

import (
	"apiGateway/DBModels"
	"apiGateway/Utils"
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
		ginCtx.JSON(502, gin.H{
			"message": "fetch registry list error",
		})
	}
	if len(registryList) == 0 {
		ginCtx.JSON(404, gin.H{
			"message": "registry list do not exist",
		})
	}
	ginCtx.JSON(200, gin.H{
		"message": "query registry list error",
		"data":    registryList,
	})
}

func SaveRegistry(ginCtx *gin.Context) {
	var registry Registry
	ginCtx.Bind(&registry)
	var registryModel DBModels.Registry
	Utils.CopyFields(&registryModel, registry,
		"Name",
		"RegistryType",
		"Addr")
	saveRegistry := registryModel.SaveRegistry()
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

func DeleteRegistry(ginCtx *gin.Context) {
	var registry Registry
	ginCtx.Bind(&registry)
	var registryModel DBModels.Registry
	Utils.CopyFields(&registryModel, registry,
		"Name",
		"RegistryType",
		"Addr")
	delRegistry := registryModel.DeleteRegistry()
	if delRegistry {
		ginCtx.JSON(200, gin.H{
			"message": "delete registry success",
		})
	} else {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error!",
		})
	}
}
