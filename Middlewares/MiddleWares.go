package Middlewares

import (
	"apiGateway/Config"
	"apiGateway/DBModels"
	"apiGateway/Utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"runtime/debug"
	"time"
)

// 解决跨域的中间件
func HeadMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Add("Access-Control-Allow-Headers", "Content-Type")
		c.Request.Header.Set("content-type", "application/json")
		c.Next()
	}
}

// 日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessLogInfo := DBModels.LogInfo{}
		accessLogInfo.LogType = Config.AccessLog
		err := accessLogInfo.GetLogInfoByType()
		if err != nil {
			return
		}
		// 若没开启日志，则直接进行下一步
		if accessLogInfo.LogRecordStatus == 0 {
			c.Next()
			return
		}
		logger := Utils.AccessLog(&accessLogInfo)
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求id
		// TODO 链路中间件的request_id，
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUrl := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求ip
		clientIP := c.ClientIP()
		// 重试次数
		retry := ""
		// 主机信息
		host := c.Request.Host
		// 日志格式
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"request_id":   "",
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUrl,
			"retry":        retry,
			"host":         host,
		}).Info()
	}
}

// 全局panic恢复中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				Utils.RuntimeLog().Errorf("unknown panic: [%s], stacktrace: [%s]", err, debug.Stack())
			}
		}()
		c.Next()
	}
}
