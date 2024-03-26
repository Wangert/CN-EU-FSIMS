package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

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
