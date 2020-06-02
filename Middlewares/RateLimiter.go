package Middlewares

import (
	"apiGateway/DBModels"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

type RateLimiterMw struct {
	DBModels.Api
}

// 限流中间件
func (mw *RateLimiterMw) RateLimitMiddleware() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(float64(mw.RateLimitNum), nil)
	lmt.SetMessage("error,request too many times,you are limited")
	return func(c *gin.Context) {
		// 获取APi信息
		api, exists := c.Get("ApiInfo")
		if exists {
			mw.Api = api.(DBModels.Api)
		}

		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
