package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

func GetMallGoods(c *gin.Context) {
	glog.Info("################## FSIMS Get Mall Goods ##################")

	mallNumber := c.Query("mall_number")
	goods, count, err := service.GetMallGoods(mallNumber)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, "get mall goods error!")
		return
	}

	res := response.ResMallGoods{
		Records: goods,
		Count:   count,
	}

	response.MakeSuccess(c, http.StatusOK, res)
	return
}
