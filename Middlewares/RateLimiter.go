package Middlewares

import (
	"apiGateway/DBModels"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

type RateLimiterMw struct {
	DBModels.Api
	RateLimiterNum int
}

// 限流中间件
func (mw *RateLimiterMw) BasicAuthMiddleware() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(float64(mw.RateLimiterNum), nil)
	lmt.SetMessage("error,request too many times,you are limited")
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
