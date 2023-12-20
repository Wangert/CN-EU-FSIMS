package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func Cors(c *gin.Context) {
	//if c.Request.Method != "OPTIONS" {
	//	c.Next()
	//} else {
	//	glog.Info("到达跨域中间件")
	//	c.Header("Access-Control-Allow-Origin", "*")
	//	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	//	c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
	//	c.Header("Access-Control-Max-Age", "2592000")
	//	c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	//	c.Header("Content-Type", "application/json")
	//	c.AbortWithStatus(200)
	//}
	glog.Info("到达跨域中间件")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
	c.Header("Access-Control-Max-Age", "2592000")
	c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Content-Type", "application/json")
	c.Next()
}
