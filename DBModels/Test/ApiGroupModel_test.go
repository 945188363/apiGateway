package Test

import (
	"apiGateway/DBModels"
	"fmt"
	"testing"
)

func TestAddApiGroup(t *testing.T) {
	testApiGroup := DBModels.ApiGroup{
		ApiGroupName:     "666",
	}
	isSuc := testApiGroup.SaveApiGroup()
	fmt.Println(isSuc)
}

func TestDelApiGroup(t *testing.T) {
	testApiGroup := DBModels.ApiGroup{
		ApiGroupName:     "666",
	}
	isSuc := testApiGroup.DeleteApiGroup()
	fmt.Println(isSuc)
}


func TestGetAllApiGroup(t *testing.T) {
	testApiGroup := DBModels.ApiGroup{}
	apiList, err:= testApiGroup.GetApiGroupList()
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(apiList)
}

func TestGetApiGroup(t *testing.T) {
	testApiGroup := DBModels.ApiGroup{
		ApiGroupName:"2aa2222233",
	}
	err:= testApiGroup.GetApiGroup()
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(testApiGroup)
}
