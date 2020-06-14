package Test

import (
	"apiGateway/DBModels"
	"fmt"
	"testing"
)

func TestAddApi(t *testing.T) {
	testApi := DBModels.Api{
		ApiName: "2aa2222233",
		ApiUrl:  "333",
	}
	isSuc := testApi.SaveApi()
	fmt.Println(isSuc)
}

func TestDelApi(t *testing.T) {
	testApi := DBModels.Api{
		ApiName: "22233",
	}
	isSuc := testApi.DeleteApi()
	fmt.Println(isSuc)
}

func TestGetAllApi(t *testing.T) {
	testApi := DBModels.Api{}
	apiList, err := testApi.GetApiList()
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(apiList)
}

func TestGetApi(t *testing.T) {
	testApi := DBModels.Api{
		ApiName: "111121312",
	}
	err := testApi.GetApi()
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(testApi)
}
