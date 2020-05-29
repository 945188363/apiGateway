package Handlers

import (
	"apiGateway/DBModels"
	"apiGateway/Utils/DataUtil"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type ApiGroup struct {
	ApiGroupName string `form:"ApiGroupName"`
	Description  string `form:"Description"`
}

func (a ApiGroup) IsEmpty() bool {
	return reflect.DeepEqual(a, ApiGroup{})
}

// ApiGroup相关处理
func GetApiGroupList(ginCtx *gin.Context) {
	var apiGroupModel DBModels.ApiGroup
	apiGroupList, err := apiGroupModel.GetApiGroupList()
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	if len(apiGroupList) == 0 {
		ginCtx.JSON(404, gin.H{
			"message": "api group list do not exist",
		})
	}
	ginCtx.JSON(200, gin.H{
		"message": "query api group list success",
		"data":    apiGroupList,
	})
}

func GetApiGroup(ginCtx *gin.Context) {
	ginCtx.String(200, "prod api")
}

func SaveApiGroup(ginCtx *gin.Context) {
	var apiGroup ApiGroup
	ginCtx.Bind(&apiGroup)
	var apiGroupModel DBModels.ApiGroup
	DataUtil.CopyFields(&apiGroupModel, apiGroup,
		"ApiGroupName",
		"Description")
	saveApiGroup := apiGroupModel.SaveApiGroup()
	if saveApiGroup {
		ginCtx.JSON(200, gin.H{
			"message": "add api group success",
		})
	} else {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "api group name is already exist or internal server error!",
		})
	}
}

func DeleteApiGroup(ginCtx *gin.Context) {
	var apiGroup ApiGroup
	ginCtx.Bind(&apiGroup)
	var apiGroupModel DBModels.ApiGroup
	DataUtil.CopyFields(&apiGroupModel, apiGroup,
		"ApiGroupName",
		"Description")
	delApiGroup := apiGroupModel.DeleteApiGroup()
	if delApiGroup {
		ginCtx.JSON(200, gin.H{
			"message": "delete api group success",
		})
	} else {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete failed or internal server error!",
		})
	}
}
