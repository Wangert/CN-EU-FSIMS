package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

func ShelfLifeForecast(c *gin.Context) {
	var r request.ReqShelfLifeForecast
	if err := c.ShouldBind(&r); err != nil {
		glog.Errorln("Shelf life forecast error")
		response.MakeFail(c, http.StatusBadRequest, "Shelf life forecast failure")
		return
	}

	/*data, err := service.GetOneSpectralData()
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "Shelf life forecast failure")
		return
	}*/

	shelfLife, err := service.ShelfLifeForecast(r.Data)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "Shelf life forecast failure")
		return
	}

	response.MakeSuccess(c, http.StatusOK, shelfLife)
	return
}

func GetAllSpectralData(c *gin.Context) {
	glog.Info("################## GetAllSpectralData ##################")

	sd, count, err := service.GetAllSpectralData()

	if err != nil {
		glog.Errorln("get all spectral data error!")
		response.MakeFail(c, http.StatusBadRequest, "get all spectral data error")
		return
	}

	res := response.ResSpectralData{
		Spectras: sd,
		Count:    count,
	}

	glog.Info("get all spectral data successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}
