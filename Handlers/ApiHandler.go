package Handlers

import (
	"apiGateway/DBModels"
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
		"message": "query api list success",
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
	fmt.Println(api)
	var apiModel DBModels.Api
	DataUtil.CopyFields(&apiModel, api)
	fmt.Println(apiModel)

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
	DataUtil.CopyFields(&apiModel, api)
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
