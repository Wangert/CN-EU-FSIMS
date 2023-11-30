package middlewares

import (
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"time"
)

func LogMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		u, ok := c.Get("uuid")
		uuid := u.(string)
		if !ok {
			response.MakeFail(c, http.StatusForbidden, "uuid not found!")
			c.Abort()
			return
		}

		//a, ok := c.Get("account")
		//account := a.(string)
		//if !ok {
		//	response.MakeFail(c, http.StatusForbidden, "uuid not found!")
		//	c.Abort()
		//	return
		//}
		//
		//ty, ok := c.Get("type")
		//typ := ty.(int)
		//if !ok {
		//	response.MakeFail(c, http.StatusForbidden, "uuid not found!")
		//	c.Abort()
		//	return
		//}

		//simulate action
		url := c.Request.URL.Path
		method := c.Request.Method
		action := method + " " + url

		pl := service.PendingLogs{
			TimeStamp: time.Now(),
			UUID:      uuid,
			Account:   "",
			Type:      0,
			Action:    action,
		}
		glog.Info("pl", pl)
		err := service.AddLog(&pl)
		if err != nil {
			response.MakeFail(c, http.StatusForbidden, err.Error())
			c.Abort()
			return
		}

		glog.Info("Write log successfully!")
		c.Next()
	}
}
