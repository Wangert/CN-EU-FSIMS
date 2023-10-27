package handlers

import (
	"CN-EU-FSIMS/internal/app/handlers/response"
	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	response.MakeSuccess(c, 200, "Welcome to FSIMS")
}
