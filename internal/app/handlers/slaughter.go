package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func UploadSlaughterStaffUniformData(c *gin.Context) {
	var r request.ReqUploadStaffUniformData
	if err := c.ShouldBind(&r); err != nil || r.HouseNumber == "" {
		response.MakeFail(c, http.StatusBadRequest, "upload slaughter staff uniform data parameters error!")
		return
	}

	err := service.UploadSlaughterStaffUniformData(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "upload slaughter staff uniform data error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, "upload slaughter staff uniform data successful!")
	return
}

func UploadSlaughterLightRecord(c *gin.Context) {
	var r request.ReqUploadSlaughterLightRecord
	if err := c.ShouldBind(&r); err != nil || r.HouseNumber == "" {
		response.MakeFail(c, http.StatusBadRequest, "upload slaughter light record parameters error!")
		return
	}

	err := service.UploadSlaughterLightRecord(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "upload slaughter light record error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, "upload slaughter light record successful!")
	return
}

func QuerySlaughterStaffUniformData(c *gin.Context) {
	var r request.ReqSlaughterSensorData
	if err := c.ShouldBind(&r); err != nil || r.HouseNumber == "" {
		response.MakeFail(c, http.StatusBadRequest, "query slaughter staff uniform data parameters error!")
		return
	}

	infos, count, err := service.QuerySlaughterStaffUniformData(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "query slaughter staff uniform data error!")
		return
	}

	res := response.ResSlaughterStaffUniformData{
		Infos: infos,
		Count: count,
	}

	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func QuerySlaughterLightRecord(c *gin.Context) {
	var r request.ReqSlaughterSensorData
	if err := c.ShouldBind(&r); err != nil || r.HouseNumber == "" {
		response.MakeFail(c, http.StatusBadRequest, "query slaughter light record parameters error!")
		return
	}

	infos, count, err := service.QuerySlaughterLightRecord(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "query slaughter light record error!")
		return
	}

	res := response.ResSlaughterLightRecord{
		Infos: infos,
		Count: count,
	}

	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func QuerySlaughterWaterQualityData(c *gin.Context) {
	var r request.ReqSlaughterSensorData
	if err := c.ShouldBind(&r); err != nil || r.HouseNumber == "" {
		response.MakeFail(c, http.StatusBadRequest, "query slaughter water quality data parameters error!")
		return
	}

	infos, count, err := service.QuerySlaughterWaterQualityData(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "query slaughter water quality data error!")
		return
	}

	res := response.ResSlaughterWaterQualityData{
		Infos: infos,
		Count: count,
	}

	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func UploadSlaughterWaterQualityData(c *gin.Context) {
	var r request.ReqUploadSlaughterWaterQualityData
	if err := c.ShouldBind(&r); err != nil || r.HouseNumber == "" {
		response.MakeFail(c, http.StatusBadRequest, "upload slaughter water quality data parameters error!")
		return
	}

	err := service.UploadSlaughterWaterQualityData(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "upload slaughter water quality data error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, "upload slaughter water quality data successful!")
	return
}

func QueryPreCoolShopData(c *gin.Context) {
	var r request.ReqSlaughterSensorData
	if err := c.ShouldBind(&r); err != nil || r.HouseNumber == "" {
		response.MakeFail(c, http.StatusBadRequest, "query precool shop data parameters error!")
		return
	}

	shops, count, err := service.QueryPreCoolShopData(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "query precool shop data error!")
		return
	}

	res := response.ResPreCoolShopData{
		ShopInfos: shops,
		Count:     count,
	}

	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func QuerySlaughterShopData(c *gin.Context) {
	var r request.ReqSlaughterSensorData
	if err := c.ShouldBind(&r); err != nil || r.HouseNumber == "" {
		response.MakeFail(c, http.StatusBadRequest, "query slaughter shop data parameters error!")
		return
	}

	shops, count, err := service.QuerySlaughterShopData(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "query slaughter shop data error!")
		return
	}

	res := response.ResSlaughterShopData{
		ShopInfos: shops,
		Count:     count,
	}

	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func QueryDivisionShopData(c *gin.Context) {
	var r request.ReqSlaughterSensorData
	if err := c.ShouldBind(&r); err != nil || r.HouseNumber == "" {
		response.MakeFail(c, http.StatusBadRequest, "query division shop data parameters error!")
		return
	}

	shops, count, err := service.QueryDivisionShopData(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "query division shop data error!")
		return
	}

	res := response.ResDivisionShopData{
		ShopInfos: shops,
		Count:     count,
	}

	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func QueryAcidShopData(c *gin.Context) {
	var r request.ReqSlaughterSensorData
	if err := c.ShouldBind(&r); err != nil || r.HouseNumber == "" {
		response.MakeFail(c, http.StatusBadRequest, "query acid shop data parameters error!")
		return
	}

	shops, count, err := service.QueryAcidShopData(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "upload acid shop data error!")
		return
	}

	res := response.ResAcidShopData{
		ShopInfos: shops,
		Count:     count,
	}
	response.MakeSuccess(c, http.StatusOK, res)
	return

}

func QueryFrozenShopData(c *gin.Context) {
	var r request.ReqSlaughterSensorData
	if err := c.ShouldBind(&r); err != nil || r.HouseNumber == "" {
		response.MakeFail(c, http.StatusBadRequest, "query frozen shop data parameters error!")
		return
	}

	shops, count, err := service.QueryFrozenShopData(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "upload frozen shop data error!")
		return
	}

	res := response.ResFrozenShopData{
		ShopInfos: shops,
		Count:     count,
	}
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func UploadPreCoolShopData(c *gin.Context) {
	var r request.ReqUploadPreCoolShopData
	if err := c.ShouldBind(&r); err != nil || r.HouseNumber == "" {
		response.MakeFail(c, http.StatusBadRequest, "upload precool shop data parameters error!")
		return
	}

	err := service.UploadPreCoolShopData(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "upload precool shop data error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, "upload precool shop data successful!")
	return
}

func UploadSlaughterShopData(c *gin.Context) {
	var r request.ReqUploadSlaughterShopData
	if err := c.ShouldBind(&r); err != nil || r.HouseNumber == "" {
		response.MakeFail(c, http.StatusBadRequest, "upload slaughter shop data parameters error!")
		return
	}

	err := service.UploadSlaughterShopData(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "upload slaughter shop data error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, "upload slaughter shop data successful!")
	return
}

func UploadDivisionShopData(c *gin.Context) {
	var r request.ReqUploadDivisionShopData
	if err := c.ShouldBind(&r); err != nil || r.HouseNumber == "" {
		response.MakeFail(c, http.StatusBadRequest, "upload division shop parameters error!")
		return
	}

	err := service.UploadDivisionShopData(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "upload division shop data error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, "upload division shop data successful!")
	return
}

func UploadAcidShopData(c *gin.Context) {
	var r request.ReqUploadAcidShopData
	if err := c.ShouldBind(&r); err != nil || r.HouseNumber == "" {
		response.MakeFail(c, http.StatusBadRequest, "upload acid shop parameters error!")
		return
	}

	err := service.UploadAcidShopData(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "upload acid shop data error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, "upload acid shop data successful!")
	return

}

func UploadFrozenShopData(c *gin.Context) {
	var r request.ReqUploadFrozenShopData
	if err := c.ShouldBind(&r); err != nil || r.HouseNumber == "" {
		response.MakeFail(c, http.StatusBadRequest, "upload frozen shop parameters error!")
		return
	}

	err := service.UploadFrozenShopData(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "upload frozen shop data error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, "upload frozen shop data successful!")
	return
}

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
	//glog.Info("test :", r)
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
	glog.Info(r.BatchNumber)
	glog.Info(r.Worker)
	glog.Info(r.HouseNumber)
	glog.Info(r.Type)
	glog.Info(r.Weight)
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
	cowNum := c.PostForm("cow_number")
	glog.Info("cowNum", cowNum)
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

func GetSlaughterData(c *gin.Context) {
	glog.Info("################## FSIMS Get Slaughter Data ##################")

	batchNum := c.Query("batch_number")
	sd, err := service.GetSlaughterData(batchNum)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "get slaughter data error!")
		return
	}
	response.MakeSuccess(c, http.StatusOK, *sd)
	return

}
