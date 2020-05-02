package Handlers

import (
	"github.com/gin-gonic/gin"
)

// ApiGroup相关处理
func GetApiGroupLst(ginCtx *gin.Context) {
	ginCtx.String(200, "prod api" )
}

func GetApiGroup(ginCtx *gin.Context) {
	ginCtx.String(200, "prod api" )
}

func SaveApiGroup(ginCtx *gin.Context) {
	ginCtx.String(200, "users api")
}

func DeleteApiGroup(ginCtx *gin.Context) {
	ginCtx.String(200, "users api")
}

