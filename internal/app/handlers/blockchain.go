package handlers

import (
	"CN-EU-FSIMS/fabric"
	"CN-EU-FSIMS/internal/app/handlers/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func QueryBlockByHeight(c *gin.Context) {
	glog.Info("################## QueryBlockByHeight ##################")
	num, _ := strconv.Atoi(c.Query("num"))

	block, err := fabric.QueryBlockByHeight(num)

	if err != nil {
		glog.Errorln("query all fsims users error!")
		response.MakeFail(c, http.StatusBadRequest, "query all fsims user error")
		return
	}

	glog.Info("query all fsims users successful")
	response.MakeSuccess(c, http.StatusOK, block)
	return
}

func QueryBlockByHash(c *gin.Context) {
	glog.Info("################## QueryBlockByHash ##################")
	hash := c.Query("hash")

	block, err := fabric.QueryBlockByHash(hash)

	if err != nil {
		glog.Errorln("query block error!")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}

	glog.Info("query block successful")
	response.MakeSuccess(c, http.StatusOK, block)
	return
}

func GetLastestBlock(c *gin.Context) {
	glog.Info("################## GetLastestBlock ##################")

	blocks, err := fabric.GetLastestBlock()

	if err != nil {
		glog.Errorln("query lastest block error!")
		response.MakeFail(c, http.StatusBadRequest, "query lastest block error")
		return
	}

	glog.Info("query lastest block successful")
	response.MakeSuccess(c, http.StatusOK, blocks)
	return
}

func GetLedgerInfo(c *gin.Context) {
	glog.Info("################## GetLedgerInfo ##################")

	ledgerInfo, err := fabric.GetLedgerInfo()

	if err != nil {
		glog.Errorln("get ledger info error!")
		response.MakeFail(c, http.StatusBadRequest, "get ledger info error")
		return
	}

	glog.Info("get ledger info successful")
	response.MakeSuccess(c, http.StatusOK, ledgerInfo)
	return
}
