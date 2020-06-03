package Test

import (
	"apiGateway/Config"
	"apiGateway/DBModels"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"testing"
	"time"
)

func TestAddPluginInfo(t *testing.T) {
	testPluginInfo := DBModels.PluginInfo{
		PluginName:   "wwww",
		PluginStatus: 0,
		Description:  "sa",
		PluginType:   Config.API,
	}
	isSuc := testPluginInfo.SavePluginInfo()
	fmt.Println(isSuc)
}

func TestDelPluginInfo(t *testing.T) {
	testPluginInfo := DBModels.PluginInfo{
		PluginName: "qweqwe",
	}
	isSuc := testPluginInfo.DeletePluginInfo()
	fmt.Println(isSuc)
}

func TestGetAllPluginInfo(t *testing.T) {
	testPluginInfo := DBModels.PluginInfo{}
	monitorInfoList, err := testPluginInfo.GetPluginInfoList()
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(monitorInfoList)
}

func TestGetPluginInfo(t *testing.T) {
	testPluginInfo := DBModels.PluginInfo{
		PluginName: "qqq",
	}
	err := testPluginInfo.GetPluginInfo()
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(testPluginInfo)
}

func TestCpu(t *testing.T) {
	cpuInfos, _ := cpu.Info()
	res, _ := cpu.Times(false)
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	fmt.Println(((res[0].Total() - res[0].Idle) / res[0].Total()) * 100)

	percent, _ := cpu.Percent(time.Second, false)
	fmt.Println(percent)

}
