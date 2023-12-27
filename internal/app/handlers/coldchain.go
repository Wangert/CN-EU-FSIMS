package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

func EndTransport(c *gin.Context) {
	glog.Info("################## FSIMS End Transport ##################")
	var r request.ReqEndTransport
	if err := c.ShouldBind(&r); err != nil || !checkEndTransportParams(&r) {
		response.MakeFail(c, http.StatusBadRequest, "transport start params error!")
		return
	}

	batchNumber, err := service.EndTransport(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "end transport error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, "["+batchNumber+"]:end transport successful!")
	return
}

func checkEndTransportParams(r *request.ReqEndTransport) bool {
	if r.BatchNumber == "" || r.Worker == "" {
		return false
	}

	return true
}

func StartTransport(c *gin.Context) {
	glog.Info("################## FSIMS Transport Start ##################")
	var r request.ReqStartTransport
	if err := c.ShouldBind(&r); err != nil || !checkStartTransportParams(&r) {
		response.MakeFail(c, http.StatusBadRequest, "transport start params error!")
		return
	}

	err := service.StartTransport(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "transport start error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, "transport start successful!")

	return
}

func checkStartTransportParams(r *request.ReqStartTransport) bool {
	if r.BatchNumber == "" || r.Operator == "" {
		return false
	}

	return true
}
