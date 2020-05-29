package Middlewares

import (
	"apiGateway/Constant/Code"
	"apiGateway/Constant/Message"
	"apiGateway/Core/Domain"
	"apiGateway/DBModels"
	"apiGateway/Utils/ComponentUtil"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"regexp"
)

type ParameterCheckMw struct {
	DBModels.Api
}

// 限流中间件
func (mw *ParameterCheckMw) ParameterCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			query := c.Request.URL.Query()
			checkParams(query, c)
		}

		if c.Request.Method == http.MethodPost {
			_ = c.Request.ParseForm()
			postForm := c.Request.PostForm
			checkParams(postForm, c)
			// defer c.Request.Body.Close()
			// body, _ := ioutil.ReadAll(c.Request.Body)
		}
	}
}

func checkParams(values url.Values, c *gin.Context) {
	flag := true
	for k, v := range values {
		// 判断参数名称是否为空
		if k == "" {
			flag = false
			break
		}
		// 判断参数值是否为空
		if len(v) == 0 {
			flag = false
			break
		}
		// 判断参数名称是否不符合命名规范 字母开头
		expr := "^[a-z,A-Z]"
		matched, err := regexp.Match(expr, []byte(k))
		if err != nil {
			ComponentUtil.RuntimeLog().Warn("regexp match request parameter error :" + err.Error())
			return
		}
		if !matched {
			flag = false
			break
		}
	}

	if !flag {
		c.JSON(Code.REQUEST_PARAM_ERROR,
			Domain.Message{
				Code: Code.REQUEST_PARAM_ERROR,
				Msg:  Message.REQUEST_PARAM_ERROR,
				Data: nil,
			})
		c.Abort()
		return
	}
	c.Next()
}
