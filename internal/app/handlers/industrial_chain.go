package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/app/models/query"
	"context"
	"github.com/gin-gonic/gin"
)

func AllIndustrialChains(c *gin.Context) {
	ics, err := query.IndustrialChain.WithContext(context.Background()).Find()
	if err != nil {
		response.MakeFail(c, 300, "query all industrial chains error.")
	}

	response.MakeSuccess(c, 200, ics)
}
