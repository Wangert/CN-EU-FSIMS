package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

func GetAllFoodchains(c *gin.Context) {
	fcs, count, err := service.QueryAllFoodchains()
	if err != nil {
		glog.Errorln("query all foodchains error!")
		response.MakeFail(c, http.StatusBadRequest, "query all foodchains error")
		return
	}
	glog.Info("query all foodchains successful")

	res := response.ResFoodchains{
		Foodchains: fcs,
		Count:      count,
	}
	response.MakeSuccess(c, http.StatusOK, res)
	return
}
