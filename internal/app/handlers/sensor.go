package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

func GetDewpointsForAllBarnPoints(c *gin.Context) {
	glog.Info("################## GetDewpointsForAllBarnPoints ##################")

	res, err := service.GetDewpointsForAllBarnPoints()

	if err != nil {
		glog.Errorln("get dewpoints for all barn points error!")
		response.MakeFail(c, http.StatusBadRequest, "get dewpoints for all barn points error")
		return
	}

	glog.Info("get dewpoints for all barn points successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func GetHumiditiesForAllBarnPoints(c *gin.Context) {
	glog.Info("################## GetHumiditiesForAllBarnPoints ##################")

	res, err := service.GetHumiditiesForAllBarnPoints()

	if err != nil {
		glog.Errorln("get humidities for all barn points error!")
		response.MakeFail(c, http.StatusBadRequest, "get humidities for all barn points error")
		return
	}

	glog.Info("get humidities for all barn points successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func GetCO2sForAllBarnPoints(c *gin.Context) {
	glog.Info("################## GetCO2sForAllBarnPoints ##################")

	res, err := service.GetCO2sForAllBarnPoints()

	if err != nil {
		glog.Errorln("get co2s for all barn points error!")
		response.MakeFail(c, http.StatusBadRequest, "get co2s for all barn points error")
		return
	}

	glog.Info("get co2s for all barn points successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func GetTemperaturesForAllBarnPoints(c *gin.Context) {
	glog.Info("################## GetTemperaturesForAllBarnPoints ##################")

	res, err := service.GetTemperaturesForAllBarnPoints()

	if err != nil {
		glog.Errorln("get temperatures for all barn points error!")
		response.MakeFail(c, http.StatusBadRequest, "get temperatures for all barn points error")
		return
	}

	glog.Info("get temperatures for all barn points successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func GetTemperaturesForAllIndependentPoints(c *gin.Context) {
	glog.Info("################## GetTemperaturesForAllIndependentPoints ##################")

	res, err := service.GetTemperaturesForAllIndependentPoints()

	if err != nil {
		glog.Errorln("get temperatures for all independent points error!")
		response.MakeFail(c, http.StatusBadRequest, "get temperatures for all independent points error")
		return
	}

	glog.Info("get temperatures for all independent points successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func GetAllSensors(c *gin.Context) {
	glog.Info("################## GetSensors ##################")

	sensors, count, err := service.GetAllSensors()

	if err != nil {
		glog.Errorln("get all sensors error!")
		response.MakeFail(c, http.StatusBadRequest, "get all sensors error")
		return
	}

	res := response.ResSensors{
		Sensors: sensors,
		Count:   count,
	}

	glog.Info("get all sensors successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func GetSensorDataByDeviceCode(c *gin.Context) {
	glog.Info("################## GetSensorDataByDeviceCode ##################")
	dc := c.Query("device_code")

	ssd, count, err := service.GetSensorDataByDeviceCode(dc)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "query sensor data by device code error!")
		return
	}

	res := response.ResSensorDataArr{
		SensorDataArr: ssd,
		Count:         count,
	}

	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func GetLatestSensorDataByDeviceCode(c *gin.Context) {
	glog.Info("################## GetLatestSensorDataByDeviceCode ##################")
	dc := c.Query("device_code")

	sd, err := service.GetLatestSensorDataByDeviceCode(dc)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "query latest sensor data by device code error!")
		return
	}

	response.MakeSuccess(c, http.StatusOK, sd)
	return
}
