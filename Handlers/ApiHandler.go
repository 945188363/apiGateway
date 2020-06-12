package Handlers

import (
	"apiGateway/DBModels"
	"apiGateway/Utils/ComponentUtil"
	"apiGateway/Utils/DataUtil"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Api struct {
	ApiName          string `form:"ApiName"`
	ApiUrl           string `form:"ApiUrl"`
	ProtocolType     string `form:"ProtocolType"`
	BackendUrl       string `form:"BackendUrl"`
	ApiMethod        string `form:"ApiMethod"`
	RateLimitNum     int    `form:"RateLimitNum"`
	ApiTimeout       int    `form:"ApiTimeout"`
	ApiRetry         int    `form:"ApiRetry"`
	ApiReturnType    string `form:"ApiReturnType"`
	ApiReturnContent string `form:"ApiReturnContent"`
	ApiGroup         string `form:"ApiGroup"`
}

func (a Api) IsEmpty() bool {
	return reflect.DeepEqual(a, Api{})
}

// Api相关处理
func GetApiLst(ginCtx *gin.Context) {
	var apiModel DBModels.Api
	apiList, err := apiModel.GetApiList()
	if err != nil {
		ComponentUtil.RuntimeLog().Info("fetch api list error")
		ginCtx.JSON(502, gin.H{
			"message": "fetch api list error",
		})
	}
	if len(apiList) == 0 {
		ComponentUtil.RuntimeLog().Info("api list do not exist")
		ginCtx.JSON(404, gin.H{
			"message": "api list do not exist",
		})
	}
	ComponentUtil.RuntimeLog().Info("query api list success")
	ginCtx.JSON(200, gin.H{
		"message": "query api list success",
		"data":    apiList,
	})
}

func GetApiListByApiNameAndGroupName(ginCtx *gin.Context) {
	var api Api
	ginCtx.ShouldBind(&api)
	var apiModel DBModels.Api
	DataUtil.CopyFields(&apiModel, api)
	fmt.Println(apiModel)
	ComponentUtil.RuntimeLog().Info("transfer data to Model :", apiModel)
	apiList, err := apiModel.GetApiListByApiNameAndGroupName()
	if err != nil {
		ComponentUtil.RuntimeLog().Info("fetch api list error")
		ginCtx.JSON(502, gin.H{
			"message": "fetch api list error",
		})
	}
	if len(apiList) == 0 {
		ComponentUtil.RuntimeLog().Info("api list do not exist")
		ginCtx.JSON(404, gin.H{
			"message": "api list do not exist",
		})
	}
	ComponentUtil.RuntimeLog().Info("query api list success")
	ginCtx.JSON(200, gin.H{
		"message": "query api list success",
		"data":    apiList,
	})
}

func GetApiByGroup(ginCtx *gin.Context) {
	var apiModel DBModels.Api
	apiGroupName := ginCtx.Query("ApiGroupName")
	if apiGroupName == "" {
		ComponentUtil.RuntimeLog().Info("can not find apis whose api group name equal empty.")
		ginCtx.JSON(404, gin.H{
			"message": "can not find apis whose api group name equal empty.",
		})
		return
	}

	apiList, err := apiModel.GetApiByGroup(apiGroupName)
	if err != nil {
		ComponentUtil.RuntimeLog().Info("internal server error")
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	if len(apiList) > 0 {
		ComponentUtil.RuntimeLog().Info("query api list by api group success")
		ginCtx.JSON(200, gin.H{
			"message": "query api list by api group success",
			"data":    apiList,
		})
		return
	} else {
		ComponentUtil.RuntimeLog().Info("this api group's api is null")
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "this api group's api is null",
		})
	}
}

func SaveApi(ginCtx *gin.Context) {
	var api Api
	ginCtx.ShouldBind(&api)
	fmt.Println(api)
	var apiModel DBModels.Api
	DataUtil.CopyFields(&apiModel, api)
	fmt.Println(apiModel)
	ComponentUtil.RuntimeLog().Info("transfer data to Model :", apiModel)
	saveApi := apiModel.SaveApi()
	if saveApi {
		ComponentUtil.RuntimeLog().Info("add api success")
		ginCtx.JSON(200, gin.H{
			"message": "add api success",
		})
	} else {
		ComponentUtil.RuntimeLog().Info("api name is already exist or internal server error!")
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "api name is already exist or internal server error!",
		})
	}
}

func DeleteApi(ginCtx *gin.Context) {
	var api Api
	ginCtx.ShouldBind(&api)
	var apiModel DBModels.Api
	DataUtil.CopyFields(&apiModel, api)
	ComponentUtil.RuntimeLog().Info("transfer data to Model :", apiModel)
	delApi := apiModel.DeleteApi()
	if delApi {
		ComponentUtil.RuntimeLog().Info("delete api success")
		ginCtx.JSON(200, gin.H{
			"message": "delete api success",
		})
	} else {
		ComponentUtil.RuntimeLog().Info("delete failed or internal server error!")
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete failed or internal server error!",
		})
	}
}
