package Handlers

import (
	"github.com/gin-gonic/gin"
)

// 黑白名单相关处理
func GetIpRestrictionLst(ginCtx *gin.Context) {
	ginCtx.String(200, "prod api")
}

func GetIpRestriction(ginCtx *gin.Context) {
	ginCtx.String(200, "prod api")
}

func SaveIpRestriction(ginCtx *gin.Context) {
	ginCtx.String(200, "users api")
}

func DeleteIpRestriction(ginCtx *gin.Context) {
	ginCtx.String(200, "users api")
}
