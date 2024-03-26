package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

func GetImgAndClass(c *gin.Context) {
	glog.Info("################## GetImgAndClass ##################")

	imgs, count, err := service.GetAllImgs()

	if err != nil {
		glog.Errorln("get all imgs error!")
		response.MakeFail(c, http.StatusBadRequest, "get all imgs error")
		return
	}

	res := response.ResImgs{
		Imgs:  imgs,
		Count: count,
	}

	glog.Info("get all imgs successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}
