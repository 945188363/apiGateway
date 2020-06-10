package Middlewares

import (
	"apiGateway/Constant/Code"
	"apiGateway/Core/Domain"
	"apiGateway/DBModels"
	"apiGateway/Utils/ComponentUtil"
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
		MaxConcurrentRequests:  mw.RateLimitNum,
		RequestVolumeThreshold: 5,
		ErrorPercentThreshold:  20,
		SleepWindow:            10000,
	}
	hystrix.ConfigureCommand(cmdName, cmdConf)
	return func(c *gin.Context) {
		ComponentUtil.RuntimeLog().Info("start Count MiddleWare...")
		// 获取api信息
		api, exists := c.Get("ApiInfo")
		if exists {
			mw.Api = api.(DBModels.Api)
		}

		// 执行hystrix策略
		_ = hystrix.Do(cmdName, func() error {
			ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(mw.ApiTimeout+100)*time.Millisecond)
			var err error
			// 检查是否超时
			defer func() {
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
			// 服务调用链是否调用成功
			if err = c.Err(); err != nil {
				return err
			}
			return err
		}, func(err error) error {
			// Utils.RuntimeLog().Info(err)
			breakerResponse(c, mw.ApiReturnContent)
			return nil
		})
		ComponentUtil.RuntimeLog().Info("end Count MiddleWare...")
	}
}

var defaultCircuitBreakerMsg = "CircuitBreaker"

// 降级方法
func breakerResponse(c *gin.Context, apiReturnContent string) {
	breakerMsg := defaultCircuitBreakerMsg
	if apiReturnContent != "" {
		breakerMsg = apiReturnContent
	}
	c.JSON(http.StatusOK, Domain.Message{
		Code: Code.SUCCESS,
		Msg:  breakerMsg,
		Data: nil,
	})
	c.Abort()
}
