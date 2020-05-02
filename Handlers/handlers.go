package Handlers

import (
	"github.com/gin-gonic/gin"
)

func GetProdsList(ginCtx *gin.Context) {
	ginCtx.String(200, "prod api" )
}

func GetUser(ginCtx *gin.Context) {
	ginCtx.String(200, "users api")
}
