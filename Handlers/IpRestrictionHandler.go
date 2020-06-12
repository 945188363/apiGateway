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

type IpRestriction struct {
	Name         string `form:"MonitorType"`
	Global       bool   `form:"Global"`
	Status       bool   `form:"Status"`
	IpWhiteList  string `form:"IpWhiteList"`
	IpBlackList  string `form:"IpBlackList"`
	ApiList      string `form:"ApiList"`
	ApiGroupList string `form:"ApiGroupList"`
}

func (a IpRestriction) IsEmpty() bool {
	return reflect.DeepEqual(a, IpRestriction{})
}

// 黑白名单相关处理
func GetIpRestrictionList(ginCtx *gin.Context) {
	var ipRestrictionModel DBModels.IpRestriction
	ipRestrictionList, err := ipRestrictionModel.GetIpRestrictionList()
	if err != nil {
		ComponentUtil.RuntimeLog().Info("fetch ipRestriction list error")
		ginCtx.JSON(502, gin.H{
			"message": "fetch ipRestriction list error",
		})
	}
	if len(ipRestrictionList) == 0 {
		ComponentUtil.RuntimeLog().Info("ipRestriction list do not exist")
		ginCtx.JSON(404, gin.H{
			"message": "ipRestriction list do not exist",
		})
	}
	ComponentUtil.RuntimeLog().Info("query ipRestriction list success")
	ginCtx.JSON(200, gin.H{
		"message": "query ipRestriction list success",
		"data":    ipRestrictionList,
	})
}

func GetIpRestriction(ginCtx *gin.Context) {
	ginCtx.String(200, "prod api")
}

func SaveIpRestriction(ginCtx *gin.Context) {
	var ipRestriction IpRestriction
	ginCtx.ShouldBind(&ipRestriction)
	fmt.Println(ipRestriction)
	var ipRestrictionModel DBModels.IpRestriction
	DataUtil.CopyFields(&ipRestrictionModel, ipRestriction,
		"Name",
		"IpWhiteList",
		"IpBlackList",
		"ApiList",
		"ApiGroupList")
	ipRestrictionModel.Status = 2
	if ipRestriction.Status {
		ipRestrictionModel.Status = 1
	}
	ipRestrictionModel.Global = 0
	if ipRestriction.Global {
		ipRestrictionModel.Global = 1
	}
	ComponentUtil.RuntimeLog().Info("transfer data to Model :", ipRestrictionModel)
	fmt.Println(ipRestrictionModel)
	saveIpRestriction := ipRestrictionModel.SaveIpRestriction()
	if saveIpRestriction {
		ComponentUtil.RuntimeLog().Info("save ipRestriction success")
		ginCtx.JSON(200, gin.H{
			"message": "save ipRestriction success",
		})
	} else {
		ComponentUtil.RuntimeLog().Info("internal server error!")
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error!",
		})
	}
}
