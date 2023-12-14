package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

func GetPastures(c *gin.Context) {
	glog.Info("################## Get All pastures ##################")

	res, err := service.GetPastures()
	if err != nil {
		glog.Errorln("query all pastures error!")
		response.MakeFail(c, http.StatusBadRequest, "query all pastures error")
		return
	}

	glog.Info("query all pastures successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func GetSlaughterHouses(c *gin.Context) {
	glog.Info("################## Get All Slaughter Houses ##################")

	res, err := service.GetSlaughterHouses()
	if err != nil {
		glog.Errorln("query all slaughter houses error!")
		response.MakeFail(c, http.StatusBadRequest, "query all slaughter houses error")
		return
	}

	glog.Info("query all slaughter houses successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func GetPackageHouses(c *gin.Context) {
	glog.Info("################## Get All Package Houses ##################")

	res, err := service.GetPackageHouses()
	if err != nil {
		glog.Errorln("query all package houses error!")
		response.MakeFail(c, http.StatusBadRequest, "query all package houses error")
		return
	}

	glog.Info("query all package houses successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func GetTransportVehicles(c *gin.Context) {
	glog.Info("################## Get All Transport Vehicles ##################")

	res, err := service.GetTransportVehicles()
	if err != nil {
		glog.Errorln("query all transport vehicles error!")
		response.MakeFail(c, http.StatusBadRequest, "query all transport vehicles error")
		return
	}

	glog.Info("query all transport vehicles successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func AddOperator(c *gin.Context) {
	glog.Info("################## Add a FSIMS operator by Admin ##################")
	var o request.ReqAddOperator
	if err := c.ShouldBind(&o); err != nil || !checkAddOperatorParams(&o) {
		response.MakeFail(c, http.StatusBadRequest, "Add a user parameters error!")
		return
	}
	// add
	glog.Info("request add  operator parameters:")
	glog.Info(o)
	pwd, err := service.AddFsimsOperator(&o)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "The new user information insert error!")
		return
	}
	glog.Info("fsims operator add successful")

	res := response.ResAddOperator{
		Account:  o.Account,
		Password: pwd,
	}
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func AddPasture(c *gin.Context) {
	glog.Info("################## Add a FSIMS pasture ##################")
	var p request.ReqAddPasture
	if err := c.ShouldBind(&p); err != nil || !checkAddPastureParams(&p) {
		response.MakeFail(c, http.StatusBadRequest, "Add a pasture parameters error!")
		return
	}

	err := service.AddPasture(&p)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "new a pasture error!")
		return
	}

	glog.Info("add pasture successful")

	response.MakeSuccess(c, http.StatusOK, "add pasture successful")
	return
}

func AddSlaughterHouse(c *gin.Context) {
	glog.Info("################## Add a FSIMS slaughter house ##################")
	var sh request.ReqAddSlaughterHouse
	if err := c.ShouldBind(&sh); err != nil || !checkAddSlaughterHouseParams(&sh) {
		response.MakeFail(c, http.StatusBadRequest, "Add a slaughter house parameters error!")
		return
	}

	err := service.AddSlaughterHouse(&sh)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "new a slaughter house error!")
		return
	}

	glog.Info("add slaughter house successful")

	response.MakeSuccess(c, http.StatusOK, "add slaughter house successful")
	return
}

func AddPackageHouse(c *gin.Context) {
	glog.Info("################## Add a FSIMS package house ##################")
	var ph request.ReqAddPackHouse
	if err := c.ShouldBind(&ph); err != nil || !checkAddPackageHouseParams(&ph) {
		response.MakeFail(c, http.StatusBadRequest, "Add a package house parameters error!")
		return
	}

	err := service.AddPackageHouse(&ph)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "new a package house error!")
		return
	}

	glog.Info("add package house successful")

	response.MakeSuccess(c, http.StatusOK, "add package house successful")
	return
}

func AddTransportVehicle(c *gin.Context) {
	glog.Info("################## Add a FSIMS vehicle house ##################")
	var v request.ReqAddVehicle
	if err := c.ShouldBind(&v); err != nil || !checkAddVehicleParams(&v) {
		response.MakeFail(c, http.StatusBadRequest, "Add a vehicle parameters error!")
		return
	}

	err := service.AddVehicle(&v)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "new a vehicle error!")
		return
	}

	glog.Info("add vehicle successful")

	response.MakeSuccess(c, http.StatusOK, "add vehicle successful")
	return
}

func checkAddVehicleParams(v *request.ReqAddVehicle) bool {
	if v.Driver == "" || v.LicenseNumber == "" || v.DriverPhone == "" {
		glog.Errorln("Missing add vehicle parameters")
		return false
	}
	return true
}

func checkAddPackageHouseParams(ph *request.ReqAddPackHouse) bool {
	if ph.Name == "" || ph.Address == "" || ph.LegalPerson == "" {
		glog.Errorln("Missing add package house parameters")
		return false
	}
	return true
}

func checkAddSlaughterHouseParams(sh *request.ReqAddSlaughterHouse) bool {
	if sh.Name == "" || sh.Address == "" || sh.LegalPerson == "" {
		glog.Errorln("Missing add slaughter house parameters")
		return false
	}
	return true
}

func checkAddPastureParams(p *request.ReqAddPasture) bool {
	if p.Name == "" || p.Address == "" || p.LegalPerson == "" {
		glog.Errorln("Missing add pasture parameters")
		return false
	}
	return true
}

func checkAddOperatorParams(reqAddOperator *request.ReqAddOperator) bool {
	if reqAddOperator.Name == "" || reqAddOperator.Account == "" || reqAddOperator.Type == 0 || reqAddOperator.Role == "" {
		glog.Errorln("Missing add operator parameters")
		return false
	}
	return true
}
