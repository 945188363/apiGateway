package Handlers

import (
	"apiGateway/DBModels"
	"apiGateway/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

type Api struct {
	ApiName          string `form:"ApiName"`
	ApiUrl           string `form:"ApiUrl"`
	BackendUrl       string `form:"BackendUrl"`
	ApiMethod        string `form:"ApiMethod"`
	ApiTimeout       string `form:"ApiTimeout"`
	ApiRetry         string `form:"ApiRetry"`
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
		ginCtx.JSON(502, gin.H{
			"message": "fetch api list error",
		})
	}
	if len(apiList) == 0 {
		ginCtx.JSON(404, gin.H{
			"message": "api list do not exist",
		})
	}

	ginCtx.JSON(200, gin.H{
		"message": "query api list error",
		"data":    apiList,
	})
}

func GetApi(ginCtx *gin.Context) {
	ginCtx.String(200, "prod api")
}

func GetApiByGroup(ginCtx *gin.Context) {
	var apiModel DBModels.Api
	apiGroupName := ginCtx.Query("ApiGroupName")
	if apiGroupName == "" {
		ginCtx.JSON(404, gin.H{
			"message": "can not find apis whose api group name equal empty.",
		})
		return
	}

	apiList, err := apiModel.GetApiByGroup(apiGroupName)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	if len(apiList) > 0 {
		ginCtx.JSON(200, gin.H{
			"message": "query api list by api group success",
			"data":    apiList,
		})
		return
	} else {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "this api group's api is null",
		})
	}
}

func SaveApi(ginCtx *gin.Context) {
	var api Api
	ginCtx.Bind(&api)
	var apiModel DBModels.Api
	Utils.CopyFields(&apiModel, api,
		"ApiName",
		"ApiUrl",
		"BackendUrl",
		"ApiMethod",
		"ApiReturnContent",
		"ApiGroup")
	apiModel.ApiRetry, _ = strconv.Atoi(api.ApiRetry)
	apiModel.ApiTimeout, _ = strconv.Atoi(api.ApiTimeout)
	saveApi := apiModel.SaveApi()
	if saveApi {
		ginCtx.JSON(200, gin.H{
			"message": "add api success",
		})
	} else {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "api name is already exist or internal server error!",
		})
	}
}

func DeleteApi(ginCtx *gin.Context) {
	var api Api
	ginCtx.Bind(&api)
	var apiModel DBModels.Api
	Utils.CopyFields(&apiModel, api,
		"ApiName",
		"ApiUrl",
		"BackendUrl",
		"ApiMethod",
		"ApiReturnContent",
		"ApiGroup")
	apiModel.ApiRetry, _ = strconv.Atoi(api.ApiRetry)
	apiModel.ApiTimeout, _ = strconv.Atoi(api.ApiTimeout)
	delApi := apiModel.DeleteApi()
	if delApi {
		ginCtx.JSON(200, gin.H{
			"message": "delete api success",
		})
	} else {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete failed or internal server error!",
		})
	}
}
