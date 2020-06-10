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

type Monitor struct {
	MonitorType   string `form:"MonitorType"`
	MonitorStatus bool   `form:"MonitorStatus"`
	MonitorConfig string `form:"MonitorConfig"`
}

func (a Monitor) IsEmpty() bool {
	return reflect.DeepEqual(a, Monitor{})
}

// Api相关处理
func GetMonitors(ginCtx *gin.Context) {
	var monitorModel DBModels.MonitorInfo
	monitorList, err := monitorModel.GetMonitorInfoList()
	if err != nil {
		ComponentUtil.RuntimeLog().Info("fetch monitor list error")
		ginCtx.JSON(502, gin.H{
			"message": "fetch monitor list error",
		})
	}
	if len(monitorList) == 0 {
		ComponentUtil.RuntimeLog().Info("monitor list do not exist")
		ginCtx.JSON(404, gin.H{
			"message": "monitor list do not exist",
		})
	}
	ComponentUtil.RuntimeLog().Info("query monitor list success")
	ginCtx.JSON(200, gin.H{
		"message": "query monitor list success",
		"data":    monitorList,
	})
}

func SaveMonitors(ginCtx *gin.Context) {
	var monitor Monitor
	ginCtx.Bind(&monitor)
	fmt.Println(monitor)
	var monitorModel DBModels.MonitorInfo
	DataUtil.CopyFields(&monitorModel, monitor,

		"MonitorType",
		"MonitorConfig")
	monitorModel.MonitorStatus = 2
	if monitor.MonitorStatus {
		monitorModel.MonitorStatus = 1
	}
	ComponentUtil.RuntimeLog().Info("transfer data to Model :", monitorModel)
	fmt.Println(monitorModel)
	saveMonitor := monitorModel.SaveMonitorInfo()
	if saveMonitor {
		ComponentUtil.RuntimeLog().Info("save monitor success")
		ginCtx.JSON(200, gin.H{
			"message": "save monitor success",
		})
	} else {
		ComponentUtil.RuntimeLog().Info("internal server error!")
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error!",
		})
	}
}
