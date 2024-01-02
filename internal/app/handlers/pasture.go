package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

func AddPastureFeedHeavyMetal(c *gin.Context) {
	var r request.ReqAddPastureFeedHeavyMetal
	fmt.Println(r.HouseNumber)
	if err := c.ShouldBind(&r); err != nil || r.HouseNumber == "" {
		glog.Errorln("add pasture feed heavy metal params error!")
		response.MakeFail(c, http.StatusNotAcceptable, "add pasture feed heavy metal params error!")
		return
	}

	err := service.AddPastureFeedHeavyMetal(&r)
	if err != nil {
		return
	}
	//err = service
	response.MakeSuccess(c, http.StatusOK, "add pasture feed heavy metal successful!")
	return
}

func SendToSlaughter(c *gin.Context) {
	var r request.ReqSendToSlaughter
	if err := c.ShouldBind(&r); err != nil || !checkSendToSlaughterParams(&r) {
		glog.Errorln("send to slaughter params error!")
		response.MakeFail(c, http.StatusNotAcceptable, "send to slaughter params error!")
		return
	}

	err := service.SendToSlaughter(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "send to slaughter error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, "sending to slaughter successful!")
	return
}

func checkSendToSlaughterParams(r *request.ReqSendToSlaughter) bool {
	if r.CowNumber == "" || r.SlaughterHouseNumber == "" || r.Operator == "" {
		return false
	}

	return true
}

func GetWarehouseInfos(c *gin.Context) {
	glog.Info("################## FSIMS Get Warehouse Infos ##################")

	houseNum := c.Query("house_number")

	pws, count, err := service.GetWarehouseInfosByPastureNumber(houseNum)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "get warehouse infos error!")
		return
	}

	res := response.ResWarehouseInfos{
		PastureWarehouses: pws,
		Count:             count,
	}

	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func EndFeeding(c *gin.Context) {
	glog.Info("################## FSIMS End Feeding ##################")

	var r request.ReqEndFeeding
	if err := c.ShouldBind(&r); err != nil || !checkEndFeedingParams(&r) {
		response.MakeFail(c, http.StatusBadRequest, "end feeding parameters error!")
		return
	}

	checkcode, cowsNum, err := service.EndFeeding(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "end feeding error!")
		return
	}

	res := response.ResEndFeeding{
		Checkcode: checkcode,
		CowsNum:   cowsNum,
		Count:     int64(len(cowsNum)),
	}

	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func checkEndFeedingParams(r *request.ReqEndFeeding) bool {
	if r.BatchNumber == "" || r.Worker == "" || r.HouseNumber == "" {
		return false
	}

	return true
}

func GetFeedingRecords(c *gin.Context) {
	glog.Info("################## FSIMS Get Feeding Records ##################")

	houseNum := c.Query("house_number")
	bs, count, err := service.GetFeedingRecords(houseNum)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "get feeding records error!")
		return
	}

	res := response.ResFeedingRecords{
		FeedingBatches: bs,
		Count:          count,
	}
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func NewFeedingBatch(c *gin.Context) {
	glog.Info("################## FSIMS New Feeding Batch ##################")

	var r request.ReqNewFeedingBatch
	if err := c.ShouldBind(&r); err != nil || !checkNewFeedingBatchParams(&r) {
		response.MakeFail(c, http.StatusBadRequest, "new feeding batch parameters error!")
		return
	}

	batchNum, err := service.NewFeedingBatch(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "new feeding batch error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, batchNum)
	return
}

func checkNewFeedingBatchParams(r *request.ReqNewFeedingBatch) bool {
	if r.HouseNumber == "" {
		return false
	}
	if len(r.CowNumbers) == 0 {
		return false
	}
	return true
}

func AddCow(c *gin.Context) {
	glog.Info("################## FSIMS Add Cow ##################")

	var r request.ReqAddCow
	if err := c.ShouldBind(&r); err != nil || !checkAddCowParams(&r) {
		response.MakeFail(c, http.StatusBadRequest, "add cow parameters error!")
		return
	}

	b, err := service.PastureIsExist(r.HouseNumber)
	if err != nil {
		glog.Errorln(err.Error())
		response.MakeFail(c, http.StatusBadRequest, "query pasture house error!")
		return
	}
	if !b {
		response.MakeFail(c, http.StatusBadRequest, "pasture is not exist!")
		return
	}

	cowInfo, err := service.AddCow(&r)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "add cow error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, cowInfo)
	return
}

func checkAddCowParams(r *request.ReqAddCow) bool {
	if r.Age == 0 || r.Weight == 0 || r.HouseNumber == "" {
		return false
	}
	return true
}
