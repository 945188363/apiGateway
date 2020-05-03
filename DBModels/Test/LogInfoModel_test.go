package Test

import (
	"apiGateway/Config"
	"apiGateway/DBModels"
	"fmt"
	"testing"
)

func TestAddLogInfo(t *testing.T) {
	testLogInfo := DBModels.LogInfo{
		LogType:         Config.RuntimeLog,
		LogName:         "11112123",
		LogRecordStatus: 0,
		LogSavedTime:    90,
		LogPeriod:       Config.Day,
	}
	isSuc := testLogInfo.SaveLogInfo()
	fmt.Println(isSuc)
}

func TestDelInfo(t *testing.T) {
	testLogInfo := DBModels.LogInfo{
		LogType: Config.AccessLog,
		LogName: "test",
	}
	isSuc := testLogInfo.DeleteLogInfo()
	fmt.Println(isSuc)
}

func TestGetAllInfo(t *testing.T) {
	testLogInfo := DBModels.LogInfo{}
	logInfoList, err := testLogInfo.GetLogInfoList()
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(logInfoList)
}

func TestGetInfo(t *testing.T) {
	testLogInfo := DBModels.LogInfo{
		LogType: Config.AccessLog,
		LogName: "test",
	}
	err := testLogInfo.GetLogInfo()
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(testLogInfo)
}
