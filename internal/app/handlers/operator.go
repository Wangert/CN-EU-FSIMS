package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/response"

	"github.com/gin-gonic/gin"
)

func PastureOperatorUpload(c *gin.Context) {
	response.MakeSuccess(c, 200, "Pasture Operator Upload Data")
}

func SlaughterOperatorUpload(c *gin.Context) {
	response.MakeSuccess(c, 200, "Slaughter Operator Upload Data")
}

func TransportOperatorUpload(c *gin.Context) {
	response.MakeSuccess(c, 200, "Transport Operator Upload Data")
}
