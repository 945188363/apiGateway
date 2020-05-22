package Middlewares

import (
	"apiGateway/DBModels"
	"context"
	"errors"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type BreakerMw struct {
	DBModels.Api
}

func (mw *BreakerMw) CircuitBreakerMiddleware() gin.HandlerFunc {
	cmdName := mw.ApiName + mw.ApiUrl
	cmdConf := hystrix.CommandConfig{
		Timeout:                mw.ApiTimeout,
		MaxConcurrentRequests:  mw.RateLimiterNum,
		RequestVolumeThreshold: 5,
		ErrorPercentThreshold:  20,
		SleepWindow:            10000,
	}
	hystrix.ConfigureCommand(cmdName, cmdConf)
	return func(c *gin.Context) {
		_ = hystrix.Do(cmdName, func() error {
			ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(mw.ApiTimeout)*time.Millisecond)
			var err error
			defer func() {
				// 检查是否超时
				if ctx.Err() == context.DeadlineExceeded {
					// 返回信息并且终止请求
					err = errors.New("timeout error")
					c.Abort()
				} else {
					err = nil
				}
				// 完成后清空资源
				cancel()
			}()

			// 包装上下文，增加Timeout限制
			c.Request = c.Request.WithContext(ctx)
			c.Next()
			return err
		}, func(err error) error {
			// Utils.RuntimeLog().Info(err)
			breakerResponse(c, mw.ApiReturnContent)
			return nil
		})
	}
}

var defaultCircuitBreakerMsg = "CircuitBreaker"

// 降级方法
func breakerResponse(c *gin.Context, apiReturnContent string) {
	breakerMsg := defaultCircuitBreakerMsg
	if apiReturnContent != "" {
		breakerMsg = apiReturnContent
	}
	breakerMsg = "ass"
	c.JSON(http.StatusOK, gin.H{
		"message": breakerMsg,
	})
	c.Abort()
}
