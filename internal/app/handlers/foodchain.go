package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func GetAllFoodchains(c *gin.Context) {
	fcs, totalCount, completedCount, err := service.QueryAllFoodchains()
	if err != nil {
		glog.Errorln("query all foodchains error!")
		response.MakeFail(c, http.StatusBadRequest, "query all foodchains error")
		return
	}
	glog.Info("query all foodchains successful")

	res := response.ResFoodchains{
		Foodchains:     fcs,
		TotalCount:     totalCount,
		CompletedCount: completedCount,
	}
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func GetPidInfo(c *gin.Context) {
	pid := c.Query("pid")
	glog.Infoln("pid:", pid)
	starttime, endtime, address, house_number, err := service.QueryPidInfo(pid)
	if err != nil {
		glog.Errorln("query pid info error!")
		response.MakeFail(c, http.StatusBadRequest, "querypid info error")
		return
	}

	res := response.ResPidInfo{
		StartTime:   starttime,
		EndTime:     endtime,
		Address:     address,
		HouseNumber: house_number,
	}
	response.MakeSuccess(c, http.StatusOK, res)
	return

}

func GetProductsByPid(c *gin.Context) {
	pid := c.Query("pid")
	ptypeStr := c.Query("type")

	ptype, err := strconv.Atoi(ptypeStr)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "type is error!")
		return
	}

	res, err := service.QueryProductsByPid(pid, ptype)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "query products by pid error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, res)
	return
}
