package Middlewares

import (
	"apiGateway/DBModels"
	"apiGateway/Utils"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BreakerMw struct {
	DBModels.Api
}

func (mw *BreakerMw) CircuitBreakerMiddleware() gin.HandlerFunc {
	cmdName := mw.ApiName + mw.ApiUrl
	cmdConf := hystrix.CommandConfig{
		MaxConcurrentRequests: 3000,
		ErrorPercentThreshold: 20,
		SleepWindow:           10000,
		Timeout:               mw.ApiTimeout,
	}
	hystrix.ConfigureCommand(cmdName, cmdConf)
	return func(c *gin.Context) {
		hystrix.NewStreamHandler()
		_ = hystrix.Do(cmdName, func() error {
			c.Next()
			if c.Err() != nil {
				return c.Err()
			}
			select {
			case <-c.Done():
				return nil
			}
		}, func(err error) error {
			Utils.RuntimeLog().Info(err)
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
	c.JSON(http.StatusOK, gin.H{
		"message": breakerMsg,
	})
}
