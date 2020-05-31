package Test

import (
	"apiGateway/Config"
	"apiGateway/DBModels"
	"fmt"
	"testing"
)

func TestAddMonitorInfo(t *testing.T) {
	testMonitorInfo := DBModels.MonitorInfo{
		MonitorType:   Config.ELK,
		MonitorStatus: 0,
		MonitorConfig: "1111122222",
	}
	isSuc := testMonitorInfo.SaveMonitorInfo()
	fmt.Println(isSuc)
}

func TestDelMonitorInfo(t *testing.T) {
	testMonitorInfo := DBModels.MonitorInfo{
		MonitorType: Config.ELK,
	}
	isSuc := testMonitorInfo.DeleteMonitorInfo()
	fmt.Println(isSuc)
}

func TestGetAllMonitorInfo(t *testing.T) {
	testMonitorInfo := DBModels.MonitorInfo{}
	monitorInfoList, err := testMonitorInfo.GetMonitorInfoList()
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(monitorInfoList)
}

func TestGetMonitorInfo(t *testing.T) {
	testMonitorInfo := DBModels.MonitorInfo{
		MonitorType: Config.ELK,
	}
	err := testMonitorInfo.GetMonitorInfo()
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(testMonitorInfo)
}
