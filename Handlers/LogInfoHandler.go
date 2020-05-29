package Handlers

import (
	"apiGateway/DBModels"
	"apiGateway/Utils/DataUtil"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

type LogInfo struct {
	LogType         int    `form:"LogType"`
	LogName         string `form:"LogName"`
	LogRecordStatus int    `form:"LogRecordStatus"`
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
	var logInfoModel DBModels.LogInfo
	logInfoList, err := logInfoModel.GetLogInfoList()
	if err != nil {
		ginCtx.JSON(502, gin.H{
			"message": "fetch log info list error",
		})
	}
	if len(logInfoList) == 0 {
		ginCtx.JSON(404, gin.H{
			"message": "log info list do not exist",
		})
	}
	ginCtx.JSON(200, gin.H{
		"message": "query log info list error",
		"data":    logInfoList,
	})
}

func GetLogInfoByType(ginCtx *gin.Context) {
	var logInfoModel DBModels.LogInfo
	logType := ginCtx.Query("LogType")
	if logType == "" {
		ginCtx.JSON(404, gin.H{
			"message": "can not find apis whose log type equal empty.",
		})
		return
	}
	logInfoModel.LogType, _ = strconv.Atoi(logType)
	err := logInfoModel.GetLogInfo()
	if err != nil {
		ginCtx.JSON(502, gin.H{
			"message": "fetch api list error",
		})
	}
	ginCtx.JSON(200, gin.H{
		"message": "query api list error",
		"data":    logInfoModel,
	})
}

func SaveLogInfo(ginCtx *gin.Context) {
	var logInfo LogInfo
	ginCtx.Bind(&logInfo)
	var logInfoModel DBModels.LogInfo
	DataUtil.CopyFields(&logInfoModel, logInfo,
		"LogType",
		"LogName",
		"LogRecordStatus",
		"LogAddress",
		"LogPeriod",
		"LogSavedTime",
		"LogRecordField")
	saveLogInfo := logInfoModel.SaveLogInfo()
	if saveLogInfo {
		ginCtx.JSON(200, gin.H{
			"message": "save log info success",
		})
	} else {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error!",
		})
	}
}
