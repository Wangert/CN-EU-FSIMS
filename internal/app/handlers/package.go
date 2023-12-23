package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func EndPackage(c *gin.Context) {
	glog.Info("################## FSIMS End Package ##################")

	var r request.ReqEndPackage
	if err := c.ShouldBind(&r); err != nil || !checkEndPackageParams(&r) {
		response.MakeFail(c, http.StatusBadRequest, "end package parameters error!")
		return
	}

	checkcode, productsNum, err := service.EndPackage(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "end package error!")
		return
	}

	res := response.ResEndPackage{
		Checkcode:   checkcode,
		ProductsNum: productsNum,
		Count:       int64(len(productsNum)),
	}

	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func checkEndPackageParams(r *request.ReqEndPackage) bool {
	if r.BatchNumber == "" || r.Worker == "" || r.HouseNumber == "" {
		return false
	}

	return true
}

func NewPackageProduct(c *gin.Context) {
	glog.Info("################## FSIMS New Package Product ##################")
	var r request.ReqNewPackageProduct
	if err := c.ShouldBind(&r); err != nil || !checkNewPackageProductParams(&r) {
		response.MakeFail(c, http.StatusBadRequest, "new package product parameters error!")
		return
	}

	number, err := service.NewPackageProduct(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "new package product error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, number)
	return
}

func checkNewPackageProductParams(r *request.ReqNewPackageProduct) bool {
	if r.BatchNumber == "" || r.Worker == "" || r.HouseNumber == "" || r.Type == 0 || r.Weight == 0 || r.ShelfLife == "" || r.PackMethod == 0 {
		return false
	}

	return true
}

func NewPackageBatch(c *gin.Context) {
	glog.Info("################## FSIMS New Package Batch ##################")

	var r request.ReqNewPackageBatch
	if err := c.ShouldBind(&r); err != nil || !checkNewPackageBatchParams(&r) {
		response.MakeFail(c, http.StatusBadRequest, "new package batch parameters error!")
		return
	}

	batchNum, err := service.NewPackageBatch(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "new package batch error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, batchNum)
	return
}

func checkNewPackageBatchParams(r *request.ReqNewPackageBatch) bool {
	if r.HouseNumber == "" || r.Worker == "" || r.PrePID == "" || r.ProductNumber == "" {
		return false
	}

	return true
}

func ConfirmProductFromSlaughter(c *gin.Context) {
	glog.Info("################## FSIMS Confirm Product From Slaughter ##################")
	productNum := c.PostForm("product_number")
	glog.Info("productNum", productNum)
	err := service.ConfirmProductFromSlaughter(productNum)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "confirm slaughter product error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, "confirm successful!")
}

func GetPackageReceiveRecords(c *gin.Context) {
	glog.Info("################## FSIMS Get Package Receive Records ##################")

	houseNum := c.Query("house_number")
	prs, count, err := service.GetPackageReceiveRecords(houseNum)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "get package receive records error!")
		return
	}

	res := response.ResPackageReceiveRecords{
		Records: prs,
		Count:   count,
	}
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func GetPackageBatches(c *gin.Context) {
	glog.Info("################## FSIMS Get Package Batches ##################")

	houseNum := c.Query("house_number")
	pbs, count, err := service.GetPackageBatches(houseNum)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "get package batches error!")
		return
	}

	res := response.ResPackageBatches{
		Records: pbs,
		Count:   count,
	}
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func GetPackageWarehouseRecords(c *gin.Context) {
	glog.Info("################## FSIMS Get Slaughter Warehouse Records ##################")

	houseNum := c.Query("house_number")
	pws, count, err := service.GetPackageWarehouseRecords(houseNum)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "get package warehouse records error!")
		return
	}

	res := response.ResPackageWarehouseRecords{
		Records: pws,
		Count:   count,
	}
	response.MakeSuccess(c, http.StatusOK, res)
	return
}
