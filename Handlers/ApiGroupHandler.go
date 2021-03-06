package Handlers

import (
	"apiGateway/DBModels"
	"apiGateway/Utils/ComponentUtil"
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
		ComponentUtil.RuntimeLog().Info("internal server error")
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	if len(apiGroupList) == 0 {
		ComponentUtil.RuntimeLog().Info("api group list do not exist")
		ginCtx.JSON(404, gin.H{
			"message": "api group list do not exist",
		})
	}
	ComponentUtil.RuntimeLog().Info("query api group list success")
	ginCtx.JSON(200, gin.H{
		"message": "query api group list success",
		"data":    apiGroupList,
	})
}

func GetApiGroupListByGroupName(ginCtx *gin.Context) {
	var apiGroup ApiGroup
	ginCtx.ShouldBind(&apiGroup)
	var apiGroupModel DBModels.ApiGroup
	DataUtil.CopyFields(&apiGroupModel, apiGroup,
		"ApiGroupName",
		"Description")
	ComponentUtil.RuntimeLog().Info("transfer data to Model :", apiGroupModel)
	apiGroupList, err := apiGroupModel.GetApiGroupListByGroupName()
	if err != nil {
		ComponentUtil.RuntimeLog().Info("internal server error")
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	if len(apiGroupList) == 0 {
		ComponentUtil.RuntimeLog().Info("api group list do not exist")
		ginCtx.JSON(404, gin.H{
			"message": "api group list do not exist",
		})
	}
	ComponentUtil.RuntimeLog().Info("query api group list success")
	ginCtx.JSON(200, gin.H{
		"message": "query api group list success",
		"data":    apiGroupList,
	})
}

func SaveApiGroup(ginCtx *gin.Context) {
	var apiGroup ApiGroup
	ginCtx.ShouldBind(&apiGroup)
	var apiGroupModel DBModels.ApiGroup
	DataUtil.CopyFields(&apiGroupModel, apiGroup,
		"ApiGroupName",
		"Description")
	ComponentUtil.RuntimeLog().Info("transfer data to Model :", apiGroupModel)
	saveApiGroup := apiGroupModel.SaveApiGroup()
	if saveApiGroup {
		ComponentUtil.RuntimeLog().Info("add api group success")
		ginCtx.JSON(200, gin.H{
			"message": "add api group success",
		})
	} else {
		ComponentUtil.RuntimeLog().Info("api group name is already exist or internal server error!")
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "api group name is already exist or internal server error!",
		})
	}
}

func DeleteApiGroup(ginCtx *gin.Context) {
	var apiGroup ApiGroup
	ginCtx.ShouldBind(&apiGroup)
	var apiGroupModel DBModels.ApiGroup
	DataUtil.CopyFields(&apiGroupModel, apiGroup,
		"ApiGroupName",
		"Description")
	ComponentUtil.RuntimeLog().Info("transfer data to Model :", apiGroupModel)
	delApiGroup := apiGroupModel.DeleteApiGroup()
	if delApiGroup {
		ComponentUtil.RuntimeLog().Info("delete api group success")
		ginCtx.JSON(200, gin.H{
			"message": "delete api group success",
		})
	} else {
		ComponentUtil.RuntimeLog().Info("delete failed or internal server error!")
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete failed or internal server error!",
		})
	}
}
