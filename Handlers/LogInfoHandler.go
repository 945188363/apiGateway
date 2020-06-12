package Handlers

import (
	"apiGateway/DBModels/LogModel"
	"apiGateway/Utils/ComponentUtil"
	"apiGateway/Utils/DataUtil"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

type LogInfo struct {
	LogType         int    `form:"LogType"`
	LogName         string `form:"LogName"`
	LogRecordStatus bool   `form:"LogRecordStatus"`
	LogAddress      string `form:"LogAddress"`
	LogPeriod       string `form:"LogPeriod"`
	LogSavedTime    int    `form:"LogSavedTime"`
	LogRecordField  string `form:"LogRecordField"`
}

func (a LogInfo) IsEmpty() bool {
	return reflect.DeepEqual(a, LogInfo{})
}

// Api相关处理
func GetLogInfoList(ginCtx *gin.Context) {
	var logInfoModel LogModel.LogInfo
	logInfoList, err := logInfoModel.GetLogInfoList()
	if err != nil {
		ComponentUtil.RuntimeLog().Info("fetch log info list error")
		ginCtx.JSON(502, gin.H{
			"message": "fetch log info list error",
		})
	}
	if len(logInfoList) == 0 {
		ComponentUtil.RuntimeLog().Info("log info list do not exist")
		ginCtx.JSON(404, gin.H{
			"message": "log info list do not exist",
		})
	}
	ComponentUtil.RuntimeLog().Info("query log info list error")
	ginCtx.JSON(200, gin.H{
		"message": "query log info list error",
		"data":    logInfoList,
	})
}

func GetLogInfoByType(ginCtx *gin.Context) {
	var logInfoModel LogModel.LogInfo
	logType := ginCtx.Query("LogType")
	if logType == "" {
		ComponentUtil.RuntimeLog().Info("can not find apis whose log type equal empty.")
		ginCtx.JSON(404, gin.H{
			"message": "can not find apis whose log type equal empty.",
		})
		return
	}
	logInfoModel.LogType, _ = strconv.Atoi(logType)
	err := logInfoModel.GetLogInfo()
	if err != nil {
		ComponentUtil.RuntimeLog().Info("fetch api list error")
		ginCtx.JSON(502, gin.H{
			"message": "fetch api list error",
		})
	}
	ComponentUtil.RuntimeLog().Info("query api list success")
	ginCtx.JSON(200, gin.H{
		"message": "query api list success",
		"data":    logInfoModel,
	})
}

func SaveLogInfo(ginCtx *gin.Context) {
	var logInfo LogInfo
	ginCtx.ShouldBind(&logInfo)
	fmt.Println(logInfo)
	var logInfoModel LogModel.LogInfo
	DataUtil.CopyFields(&logInfoModel, logInfo,
		"LogType",
		"LogName",
		"LogAddress",
		"LogPeriod",
		"LogSavedTime",
		"LogRecordField")
	logInfoModel.LogRecordStatus = 2
	if logInfo.LogRecordStatus {
		logInfoModel.LogRecordStatus = 1
	}
	ComponentUtil.RuntimeLog().Info("transfer data to Model :", logInfoModel)
	fmt.Println(logInfoModel)
	saveLogInfo := logInfoModel.SaveLogInfo()
	if saveLogInfo {
		ComponentUtil.RuntimeLog().Info("save log info success")
		ginCtx.JSON(200, gin.H{
			"message": "save log info success",
		})
	} else {
		ComponentUtil.RuntimeLog().Info("internal server error!")
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error!",
		})
	}
}
