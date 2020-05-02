package Handlers

import (
	"github.com/gin-gonic/gin"
)

// Api相关处理
func GetApiLst(ginCtx *gin.Context) {
	ginCtx.String(200, "prod api" )
}

func GetApi(ginCtx *gin.Context) {
	ginCtx.String(200, "prod api" )
}

func SaveApi(ginCtx *gin.Context) {
	ginCtx.String(200, "users api")
}

func DeleteApi(ginCtx *gin.Context) {
	ginCtx.String(200, "users api")
}
