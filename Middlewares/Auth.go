package Middlewares

import (
	"apiGateway/Constant/Code"
	"apiGateway/Constant/Message"
	"apiGateway/DBModels"
	"apiGateway/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type AuthMw struct {
	DBModels.Api
}

func (mw *AuthMw) JWTAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 如果是网关操作的接口直接放行
		if url := c.Request.URL.String(); strings.HasPrefix(url, "/gateway") {
			c.Next()
			return
		}
		var data interface{}
		msg := Message.SUCCESS
		code := Code.SUCCESS
		token := c.Query("token")
		if token == "" {
			if cookie, err := c.Request.Cookie("token"); err == nil {
				value := cookie.Value
				claims, err := Utils.ParseToken(value)
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{
						"error": "Token error",
					})
				} else if time.Now().Unix() > claims.ExpiresAt {
					c.JSON(http.StatusUnauthorized, gin.H{
						"error": "Token expired",
					})
				}
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()

		} else {
			claims, err := Utils.ParseToken(token)
			if err != nil {
				code = Code.ERROR_AUTH_CHECK_TOKEN_FAIL
				msg = Message.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = Code.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				msg = Message.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != Code.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  msg,
				"data": data,
			})

			c.Abort()
			return
		}
		c.Next()
	}
}

func (mw *AuthMw) AkSkAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (mw *AuthMw) BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果是网关操作的接口直接放行
		if url := c.Request.URL.String(); strings.HasPrefix(url, "/gateway") {
			c.Next()
			return
		}
		token := c.Query("token")
		if token == "" {
			if cookie, err := c.Request.Cookie("token"); err == nil {
				value := cookie.Value
				if value == "tokenValue" {
					c.Next()
					return
				}
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
		} else {
			if token == "tokenValue" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
	}
}
