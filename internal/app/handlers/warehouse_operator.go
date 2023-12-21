package handlers

//
//import (
//	"CN-EU-FSIMS/internal/app/handlers/request"
//	"CN-EU-FSIMS/internal/app/handlers/response"
//	"CN-EU-FSIMS/internal/service/warehouse"
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//	"github.com/golang/glog"
//)
//
//func PastureInWarehouse(c *gin.Context) {
//	var piw request.ReqInPastureWarehouse
//
//	if err := c.ShouldBind(&piw); err != nil {
//		glog.Errorln(err.Error())
//		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
//		return
//	}
//	if err := warehouse.PastureInWarehouse(piw); err != nil {
//		glog.Errorln("in pasture warehouse error")
//		response.MakeFail(c, http.StatusBadRequest, err.Error())
//		return
//	}
//
//	response.MakeSuccess(c, http.StatusOK, "pasture in warehouse successful!")
//	return
//}
//

//
//func SlaughterReceived(c *gin.Context) {
//	var rcp request.SlaughterReceive
//	if err := c.ShouldBind(&rcp); err != nil {
//		glog.Errorln(err.Error())
//		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
//		return
//	}
//
//	if err := warehouse.SlaughterReceived(rcp); err != nil {
//		glog.Errorln("slaughter receive error")
//		response.MakeFail(c, http.StatusBadRequest, err.Error())
//		return
//	}
//
//	response.MakeSuccess(c, http.StatusOK, "sending to slaughter successful!")
//	return
//}
//
//func SlaughterInWarehouse(c *gin.Context) {
//	var pis request.ReqInSlaughterWarehouse
//
//	if err := c.ShouldBind(&pis); err != nil {
//		glog.Errorln(err.Error())
//		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
//		return
//	}
//	if err := warehouse.SlaughterInWarehouse(pis); err != nil {
//		glog.Errorln("in slaughter warehouse error")
//		response.MakeFail(c, http.StatusBadRequest, err.Error())
//		return
//	}
//
//	response.MakeSuccess(c, http.StatusOK, "in slaughter warehouse successful!")
//	return
//}
//
//func SendToPack(c *gin.Context) {
//	var rcp request.ReqSendToPack
//	if err := c.ShouldBind(&rcp); err != nil {
//		glog.Errorln(err.Error())
//		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
//		return
//	}
//
//	if err := warehouse.SendToPack(rcp); err != nil {
//		glog.Errorln("sending to pack error")
//		response.MakeFail(c, http.StatusBadRequest, err.Error())
//		return
//	}
//
//	response.MakeSuccess(c, http.StatusOK, "sending to pack successful!")
//	return
//}
//
//func PackReceived(c *gin.Context) {
//	var rcp request.PackReceive
//	if err := c.ShouldBind(&rcp); err != nil {
//		glog.Errorln(err.Error())
//		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
//		return
//	}
//
//	if err := warehouse.PackReceived(rcp); err != nil {
//		glog.Errorln("pack receive error")
//		response.MakeFail(c, http.StatusBadRequest, err.Error())
//		return
//	}
//
//	response.MakeSuccess(c, http.StatusOK, "sending to pack successful!")
//	return
//}
//
//func PackInWarehouse(c *gin.Context) {
//	var pis request.ReqInPackWarehouse
//
//	if err := c.ShouldBind(&pis); err != nil {
//		glog.Errorln(err.Error())
//		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
//		return
//	}
//	if err := warehouse.PackInWarehouse(pis); err != nil {
//		glog.Errorln("in pack warehouse error")
//		response.MakeFail(c, http.StatusBadRequest, err.Error())
//		return
//	}
//
//	response.MakeSuccess(c, http.StatusOK, "in pack warehouse successful!")
//	return
//}
