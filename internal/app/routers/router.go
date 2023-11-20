package routers

import (
	"CN-EU-FSIMS/internal/app/handlers"
	"CN-EU-FSIMS/internal/app/middlewares"

	"github.com/gin-gonic/gin"
)

func Load(e *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// use middlewares
	e.Use(gin.Recovery(), gin.Logger())
	e.Use(mw...)

	fsims := e.Group("/fsims")

	fsims.GET("home", handlers.Welcome)

	// user router group
	user := fsims.Group("/user")
	{
		user.POST("register", handlers.Register)
		user.POST("login", handlers.Login)
	}

	// admin router group
	admin := fsims.Group("/admin", middlewares.JwtAuth(), middlewares.CheckPermission())
	{
		admin.GET("allusers", handlers.GetAllUsers)
	}

	// industrial chain group
	ic := fsims.Group("/industrial", middlewares.JwtAuth(), middlewares.CheckPermission())
	{
		ic.GET("all", handlers.AllIndustrialChains)
	}

	// pastureoperator router group
	pop := fsims.Group("/pastureoperator", middlewares.JwtAuth(), middlewares.CheckPermission())
	{
		pop.GET("upload", handlers.PastureOperatorUpload)
	}

	// slaughteroperator router group
	sop := fsims.Group("/slaughteroperator", middlewares.JwtAuth(), middlewares.CheckPermission())
	{
		sop.GET("upload", handlers.SlaughterOperatorUpload)
	}

	// transportoperator router group
	top := fsims.Group("/transportoperator", middlewares.JwtAuth(), middlewares.CheckPermission())
	{
		top.GET("upload", handlers.TransportOperatorUpload)
	}
	return e
}
