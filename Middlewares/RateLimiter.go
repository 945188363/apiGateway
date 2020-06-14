package Middlewares

import (
	"apiGateway/DBModels"
	"apiGateway/Utils/ComponentUtil"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

type RateLimiterMw struct {
	DBModels.Api
}

// 限流中间件
func (mw *RateLimiterMw) RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ComponentUtil.RuntimeLog().Info("start RateLimit MiddleWare...")

		// 获取APi信息
		api, exists := c.Get("ApiInfo")
		if !exists {
			ComponentUtil.RuntimeLog().Info("api info is null.")
			c.Abort()
			return
		}
		mw.Api = api.(DBModels.Api)

		lmt := tollbooth.NewLimiter(float64(mw.RateLimitNum), nil)
		lmt.SetMessage("error,request too many times,you are limited")

		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			c.Abort()
		} else {
			c.Next()
		}
		ComponentUtil.RuntimeLog().Info("end RateLimit MiddleWare...")

	}
}
