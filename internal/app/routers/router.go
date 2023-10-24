package routers

import (
	"CN-EU-FSIMS/internal/app/handlers"
	"github.com/gin-gonic/gin"
)

func Load(e *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// use middlewares
	e.Use(gin.Recovery(), gin.Logger())
	e.Use(mw...)

	fsims := e.Group("/fsims")

	// user router group
	user := fsims.Group("/user")
	{
		user.POST("login", handlers.Login)
	}

	return e
}
