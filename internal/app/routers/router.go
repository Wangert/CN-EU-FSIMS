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
	admin := fsims.Group("/admin", middlewares.JwtAuth(), middlewares.CheckPermission(), middlewares.LogMiddleware())
	{
		admin.GET("blockchain/block", handlers.QueryBlockByHeight)
		admin.GET("blockchain/ledgerinfo", handlers.GetLedgerInfo)
		admin.GET("blockchain/latestblock", handlers.GetLastestBlock)

		admin.POST("addpasture", handlers.AddPasture)
		admin.POST("addslaughterhouse", handlers.AddSlaughterHouse)
		admin.POST("addpackhouse", handlers.AddPackageHouse)
		admin.POST("addtransportvehicle", handlers.AddTransportVehicle)

		admin.GET("pastures", handlers.GetPastures)
		admin.GET("searchpas", handlers.SearchPastures)
		admin.GET("slaughterhouses", handlers.GetSlaughterHouses)
		admin.GET("searchsla", handlers.SearchSlaughterHouses)
		admin.GET("packagehouses", handlers.GetPackageHouses)
		admin.GET("searchpac", handlers.SearchPackageHouses)
		admin.GET("transportvehicles", handlers.GetTransportVehicles)
		admin.GET("searchtv", handlers.SearchTransportVehicles)

		admin.POST("addoperator", handlers.AddOperator)
		admin.POST("adduser", handlers.AddUserByAdmin)
		admin.GET("allusers", handlers.GetAllUsers)
		admin.GET("searchusers", handlers.SearchUsers)
		admin.GET("deleteuser", handlers.DeleteUser)
		admin.GET("reset", handlers.ResetPasswordByAdmin)

		admin.GET("viewlog", handlers.ViewLogs)

		// 便于测试
		admin.POST("newfeedingbatch", handlers.NewFeedingBatch)
		admin.POST("addcow", handlers.AddCow)
		admin.GET("getfeedingrecords", handlers.GetFeedingRecords)
		admin.POST("endfeeding", handlers.EndFeeding)
	}

	// industrial chain group
	//ic := fsims.Group("/industrial", middlewares.JwtAuth(), middlewares.CheckPermission())
	//{
	//	ic.GET("all", handlers.AllIndustrialChains)
	//}

	//operator create procedure
	op := fsims.Group("/operator", middlewares.JwtAuth(), middlewares.CheckPermissionAndType())
	{

		op.POST("createproc", handlers.CreateProcedure)
		op.GET("queryproc", handlers.QueryProcedureWithPID)
		op.GET("querychain", handlers.QueryIndustrialChainWithCheckcode)
		op.GET("verify", handlers.VerifyWithCheckcode)
		//op.POST("createproc", handlers.CreateProcedure)
	}

	// pastureoperator router group
	pop := fsims.Group("/pastureoperator", middlewares.JwtAuth(), middlewares.CheckPermission())
	{
		//pop.POST("createproc", handlers.CreateProcedure)

		pop.POST("addcow", handlers.AddCow)
		pop.POST("newfeedingbatch", handlers.NewFeedingBatch)
		pop.GET("getfeedingrecords", handlers.GetFeedingRecords)
		pop.POST("endfeeding", handlers.EndFeeding)
		pop.POST("commitproc", handlers.CommitPastureProcedure)
		pop.POST("inwarehouse", handlers.PastureInWarehouse)
		pop.POST("sendtonext", handlers.SendToSlaughter)
	}

	//slaughteroperator router group
	sop := fsims.Group("/slaughteroperator", middlewares.JwtAuth(), middlewares.CheckPermission())
	{
		sop.POST("commitproc", handlers.CommitSlaughterProcedure)
		sop.POST("receive", handlers.SlaughterReceived)
		sop.POST("inwarehouse", handlers.SlaughterInWarehouse)
		sop.POST("sendtonext", handlers.SendToPack)
	}

	//packoperator router group
	kop := fsims.Group("/packoperator", middlewares.JwtAuth(), middlewares.CheckPermission())
	{
		kop.POST("commitproc", handlers.CommitPackProcedure)
		kop.POST("receive", handlers.PackReceived)
		kop.POST("inwarehouse", handlers.PackInWarehouse)

	}

	//transportoperator router group
	top := fsims.Group("/transportoperator", middlewares.JwtAuth(), middlewares.CheckPermission())
	{
		top.POST("start", handlers.TransportStart)
		top.POST("end", handlers.TransportEnd)
	}
	return e
}
