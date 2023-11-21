package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"github.com/golang/glog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PastureOperatorUpload(c *gin.Context) {
	response.MakeSuccess(c, 200, "Pasture Operator Upload Data")
}

func SlaughterOperatorUpload(c *gin.Context) {
	response.MakeSuccess(c, 200, "Slaughter Operator Upload Data")
}

func TransportOperatorUpload(c *gin.Context) {
	response.MakeSuccess(c, 200, "Transport Operator Upload Data")
}

func CreateProcedure(c *gin.Context) {
	var rcp request.ReqCreateProcedure
	if err := c.ShouldBind(&rcp); err != nil {
		glog.Errorln(err.Error())
		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
		return
	}

	if err := service.AddProcedure(rcp); err != nil {
		glog.Errorln("add procedure error")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}

	response.MakeSuccess(c, http.StatusOK, "create procedure successful!")
	return
}

func checkCreateProcedureParams(rcp *request.ReqCreateProcedure) bool {
	if rcp.Type == 0 || rcp.PrePID == "" || rcp.Operator == "" {
		return false
	}
	return true
}

func CommitPastureProcedure(c *gin.Context) {
	var cpp request.CommitPastureProcedure
	if err := c.ShouldBind(&cpp); err != nil {
		glog.Errorln(err.Error())
		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
		return
	}

	checkcode, err := service.CommitPastureProcedure(cpp)
	if err != nil {
		glog.Errorln("commit pasture procedure error")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}

	response.MakeSuccess(c, http.StatusOK, checkcode)
	return
}
