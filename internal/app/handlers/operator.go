package handlers

import (
	"CN-EU-FSIMS/fabric"
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"net/http"

	"github.com/golang/glog"

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

func QueryProcedureWithPID(c *gin.Context) {
	pid := c.Query("pid")
	glog.Info("pid:", pid)
	proc, err := fabric.QueryProcedureWithPID(pid)

	if err != nil {
		glog.Errorln("query procedure error")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}

	glog.Info("query all procedure successful")
	response.MakeSuccess(c, http.StatusOK, proc)
	return
}

func QueryIndustrialChainWithCheckcode(c *gin.Context) {
	checkcode := c.Query("checkcode")
	glog.Info("checkcode:", checkcode)

	procs, err := service.QueryIndustrialChain(checkcode)

	if err != nil {
		glog.Errorln("query industrail chain by checkcode error")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}

	glog.Info("query industrail chain by checkcode successful")
	response.MakeSuccess(c, http.StatusOK, procs)
	return
}

func VerifyWithCheckcode(c *gin.Context) {
	checkcode := c.Query("checkcode")
	glog.Info("checkcode:", checkcode)

	res, err := service.VerifyWithCheckcode(checkcode)

	if err != nil {
		glog.Errorln("verify with checkcode error")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}

	glog.Info("verify with checkcode successful")
	response.MakeSuccess(c, http.StatusOK, res)
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

func CommitSlaughterProcedure(c *gin.Context) {
	var csp request.CommitSlaughterProcedure
	if err := c.ShouldBind(&csp); err != nil {
		glog.Errorln(err.Error())
		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
		return
	}

	checkcode, err := service.CommitSlaughterProcedure(csp)
	if err != nil {
		glog.Errorln("commit slaughter procedure error")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}

	response.MakeSuccess(c, http.StatusOK, checkcode)
	return
}

func CommitPackProcedure(c *gin.Context) {
	var cpp request.CommitPackProcedure
	if err := c.ShouldBind(&cpp); err != nil {
		glog.Errorln(err.Error())
		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
		return
	}

	checkcode, err := service.CommitPackProcedure(cpp)
	if err != nil {
		glog.Errorln("commit pack procedure error")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}

	response.MakeSuccess(c, http.StatusOK, checkcode)
	return
}

func TransportStart(c *gin.Context) {
	var ctp request.TransportStart
	if err := c.ShouldBind(&ctp); err != nil {
		glog.Errorln(err.Error())
		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
		return
	}

	if err := service.AddTransportProcedure(ctp); err != nil {
		glog.Errorln("commit transport procedure error")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}

	response.MakeSuccess(c, http.StatusOK, "create transport start successful")
	return
}

func TransportEnd(c *gin.Context) {
	var te request.TransportEnd
	if err := c.ShouldBind(&te); err != nil {
		glog.Errorln(err.Error())
		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
		return
	}

	checkcode, err := service.TransportEnd(te)
	if err != nil {
		glog.Errorln("commit transport end error")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}

	response.MakeSuccess(c, http.StatusOK, checkcode)
	return
}
