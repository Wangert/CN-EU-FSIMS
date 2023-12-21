package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

func SendToPackage(c *gin.Context) {
	var r request.ReqSendToPackage
	if err := c.ShouldBind(&r); err != nil || !checkSendToPackageParams(&r) {
		glog.Errorln(err.Error())
		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
		return
	}

	err := service.SendToPackage(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "send to package error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, "sending to package successful!")
	return
}

func checkSendToPackageParams(r *request.ReqSendToPackage) bool {
	if r.ProductNumber == "" || r.PackageHouseNumber == "" || r.Operator == "" {
		return false
	}

	return true
}

func EndSlaughter(c *gin.Context) {
	glog.Info("################## FSIMS End Slaughter ##################")

	var r request.ReqEndSlaughter
	if err := c.ShouldBind(&r); err != nil || !checkEndSlaughterParams(&r) {
		response.MakeFail(c, http.StatusBadRequest, "end slaughter parameters error!")
		return
	}

	checkcode, productsNum, err := service.EndSlaughter(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "end slaughter error!")
		return
	}

	res := response.ResEndSlaughter{
		Checkcode:   checkcode,
		ProductsNum: productsNum,
		Count:       int64(len(productsNum)),
	}

	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func checkEndSlaughterParams(r *request.ReqEndSlaughter) bool {
	if r.BatchNumber == "" || r.Worker == "" || r.HouseNumber == "" {
		return false
	}

	return true
}

func NewSlaughterProduct(c *gin.Context) {
	glog.Info("################## FSIMS New Slaughter Product ##################")
	var r request.ReqNewSlaughterProduct
	if err := c.ShouldBind(&r); err != nil || !checkNewSlaughterProductParams(&r) {
		response.MakeFail(c, http.StatusBadRequest, "new slaughter product parameters error!")
		return
	}

	number, err := service.NewSlaughterProduct(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "new slaughter product error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, number)
	return
}

func checkNewSlaughterProductParams(r *request.ReqNewSlaughterProduct) bool {
	if r.BatchNumber == "" || r.Worker == "" || r.HouseNumber == "" || r.Type == 0 || r.Weight == 0 {
		return false
	}

	return true
}

func NewSlaughterBatch(c *gin.Context) {
	glog.Info("################## FSIMS New Slaughter Batch ##################")

	var r request.ReqNewSlaughterBatch
	if err := c.ShouldBind(&r); err != nil || !checkNewSlaughterBatchParams(&r) {
		response.MakeFail(c, http.StatusBadRequest, "new slaughter batch parameters error!")
		return
	}

	batchNum, err := service.NewSlaughterBatch(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "new slaughter batch error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, batchNum)
	return
}

func checkNewSlaughterBatchParams(r *request.ReqNewSlaughterBatch) bool {
	if r.HouseNumber == "" || r.Worker == "" || r.PrePID == "" || r.CowNumber == "" {
		return false
	}

	return true
}

func ConfirmCowFromPasture(c *gin.Context) {
	glog.Info("################## FSIMS Confirm Cow From Pasture ##################")
	cowNum := c.Query("cow_number")
	err := service.ConfirmCowFromPasture(cowNum)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "confirm cow error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, "confirm successful!")
}

func GetSlaughterReceiveRecords(c *gin.Context) {
	glog.Info("################## FSIMS Get Slaughter Receive Records ##################")

	houseNum := c.Query("house_number")
	srs, count, err := service.GetSlaughterReceiveRecords(houseNum)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "get slaughter receive records error!")
		return
	}

	res := response.ResSlaughterReceiveRecords{
		Records: srs,
		Count:   count,
	}
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func GetSlaughterBatches(c *gin.Context) {
	glog.Info("################## FSIMS Get Slaughter Batches ##################")

	houseNum := c.Query("house_number")
	sbs, count, err := service.GetSlaughterBatches(houseNum)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "get slaughter batches error!")
		return
	}

	res := response.ResSlaughterBatches{
		Records: sbs,
		Count:   count,
	}
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func GetSlaughterWarehouseRecords(c *gin.Context) {
	glog.Info("################## FSIMS Get Slaughter Warehouse Records ##################")

	houseNum := c.Query("house_number")
	sws, count, err := service.GetSlaughterWarehouseRecords(houseNum)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "get slaughter warehouse records error!")
		return
	}

	res := response.ResSlaughterWarehouseRecords{
		Records: sws,
		Count:   count,
	}
	response.MakeSuccess(c, http.StatusOK, res)
	return
}
