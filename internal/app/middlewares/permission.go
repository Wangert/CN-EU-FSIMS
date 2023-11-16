package middlewares

import (
	"CN-EU-FSIMS/database/mysql"
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

// CheckPermission
func CheckPermission() func(c *gin.Context) {
	return func(c *gin.Context) {
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
