package Core

import (
	"apiGateway/DBModels"
	"apiGateway/Routers"
	"apiGateway/Utils"
	"sort"
)

// 初始化路由
func InitApiMapping() {
	router := Routers.NewGinRouter()
	api := DBModels.Api{}
	apiList, err := api.GetApiList()
	// 根据apiGroup分组
	if err != nil {
		Utils.RuntimeLog().Info("get api list error", err)
	}
	apiListGroup := splitByGroup(apiList)

	for i := 0; i < len(apiListGroup); i++ {
		for j := 0; j < len(apiListGroup[i]); j++ {
			if apiListGroup[i][j].ApiGroup != "" {
				group := router.Group("/" + apiListGroup[i][j].ApiGroup)
				{
					switch apiListGroup[i][j].ApiMethod {
					case "ALL":
						group.GET(apiListGroup[i][j].ApiUrl)
					case "GET":
						group.GET(apiListGroup[i][j].ApiUrl)
					case "POST":
						group.POST(apiListGroup[i][j].ApiUrl)
					case "PUT":
						group.PUT(apiListGroup[i][j].ApiUrl)
					case "PATCH":
						group.PATCH(apiListGroup[i][j].ApiUrl)
					case "DELETE":
						group.DELETE(apiListGroup[i][j].ApiUrl)
					case "OPTIONS":
						group.OPTIONS(apiListGroup[i][j].ApiUrl)
					case "HEAD":
						group.HEAD(apiListGroup[i][j].ApiUrl)
					default:
						group.GET(apiListGroup[i][j].ApiUrl)
					}
				}
			} else {
				switch apiListGroup[i][j].ApiMethod {
				case "ALL":
					router.GET(apiListGroup[i][j].ApiUrl)
				case "GET":
					router.GET(apiListGroup[i][j].ApiUrl)
				case "POST":
					router.POST(apiListGroup[i][j].ApiUrl)
				case "PUT":
					router.PUT(apiListGroup[i][j].ApiUrl)
				case "PATCH":
					router.PATCH(apiListGroup[i][j].ApiUrl)
				case "DELETE":
					router.DELETE(apiListGroup[i][j].ApiUrl)
				case "OPTIONS":
					router.OPTIONS(apiListGroup[i][j].ApiUrl)
				case "HEAD":
					router.HEAD(apiListGroup[i][j].ApiUrl)
				default:
					router.GET(apiListGroup[i][j].ApiUrl)
				}
			}
		}
	}
}

type sortList []DBModels.Api

func (s sortList) Len() int           { return len(s) }
func (s sortList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortList) Less(i, j int) bool { return s[i].ApiGroup < s[j].ApiGroup }

// 分组聚合
func splitByGroup(apiList []DBModels.Api) [][]DBModels.Api {
	sort.Sort(sortList(apiList))
	returnData := make([][]DBModels.Api, 0)
	i := 0
	var j int
	for {
		if i >= len(apiList) {
			break
		}
		for j = i + 1; j < len(apiList) && apiList[i].ApiGroup == apiList[j].ApiGroup; j++ {
		}
		returnData = append(returnData, apiList[i:j])
		i = j
	}
	return returnData
}
