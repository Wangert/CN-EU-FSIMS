package middlewares

import (
	"CN-EU-FSIMS/database/mysql"
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// CheckPermission
func CheckPermission() func(c *gin.Context) {
	return func(c *gin.Context) {
		glog.Infoln("到达检查权限")
		u, ok := c.Get("uuid")
		uuid := u.(string)
		if !ok {
			response.MakeFail(c, http.StatusForbidden, "uuid not found!")
			c.Abort()
			return
		}
		url := c.Request.URL.Path
		method := c.Request.Method

		glog.Infoln("Url:", url)
		glog.Infoln("Method:", method)

		cs, err := service.NewCasbinService(mysql.DB)
		if err != nil {
			response.MakeFail(c, http.StatusForbidden, err.Error())
			c.Abort()
			return
		}

		pp := service.PendingPolicy{
			UUID:   uuid,
			Url:    url,
			Method: method,
		}
		glog.Info("pp:", pp)
		ok, err = cs.CheckPermission(&pp)
		if err != nil {
			response.MakeFail(c, http.StatusForbidden, err.Error())
			c.Abort()
			return
		}

		if ok {
			glog.Infoln(uuid, " pass!")
			c.Next()
		} else {
			glog.Infoln(uuid, " deny!")
			response.MakeFail(c, http.StatusForbidden, "Not Permission")
			c.Abort()
		}
	}
}

// CheckPermissionAndType
func CheckPermissionAndType() func(c *gin.Context) {
	return func(c *gin.Context) {
		u, ok := c.Get("uuid")
		uuid := u.(string)
		if !ok {
			response.MakeFail(c, http.StatusForbidden, "uuid not found!")
			c.Abort()
			return
		}
		url := c.Request.URL.Path
		parts := strings.Split(uuid, "-")
		var method string
		switch parts[1] {
		case "0003":
			method = "upload_4"
		case "0004":
			method = "upload_5"
		case "0005":
			method = "upload_6"
		case "0006":
			method = "upload_7"
		case "0007":
			method = "upload_8"
		}

		glog.Infoln("Url:", url)
		glog.Infoln("Method:", method)

		cs, err := service.NewCasbinService(mysql.DB)
		if err != nil {
			response.MakeFail(c, http.StatusForbidden, err.Error())
			c.Abort()
			return
		}

		pp := service.PendingPolicy{
			UUID:   uuid,
			Url:    url,
			Method: method,
		}
		glog.Info("pp:", pp)
		ok, err = cs.CheckPermission(&pp)
		if err != nil {
			response.MakeFail(c, http.StatusForbidden, err.Error())
			c.Abort()
			return
		}

		if ok {
			glog.Infoln(uuid, " pass!")
			c.Next()
		} else {
			glog.Infoln(uuid, " deny!")
			response.MakeFail(c, http.StatusForbidden, "Not Permission")
			c.Abort()
		}
	}
}
