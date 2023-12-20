package middlewares

import (
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

// JwtAuth 基于JWT的认证中间件
func JwtAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		glog.Infoln("到达JWT中间件")
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// Token放在Header的Authorization中
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.MakeFail(c, http.StatusBadRequest, "Authroization is empty")
			c.Abort()
			return
		}

		glog.Infoln("token:", token)

		userClaims, err := service.ParseToken(token)
		if err != nil {
			response.MakeFail(c, http.StatusBadRequest, "invalid token")
			c.Abort()
			return
		}
		// 将当前请求的信息保存到请求的上下文gin.context中
		c.Set("uuid", userClaims.UUID)
		c.Next() // 后续的处理函数可以用过c.Get("userAccount")来获取当前请求的用户信息
	}
}
