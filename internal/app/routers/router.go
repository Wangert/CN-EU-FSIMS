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
		admin.POST("addmall", handlers.AddMall)

		admin.GET("pastures", handlers.GetPastures)
		admin.GET("searchpas", handlers.SearchPastures)
		admin.GET("slaughterhouses", handlers.GetSlaughterHouses)
		admin.GET("searchsla", handlers.SearchSlaughterHouses)
		admin.GET("packagehouses", handlers.GetPackageHouses)
		admin.GET("searchpac", handlers.SearchPackageHouses)
		admin.GET("transportvehicles", handlers.GetTransportVehicles)
		admin.GET("searchtv", handlers.SearchTransportVehicles)
		admin.GET("malls", handlers.GetMalls)
		admin.GET("searchmall", handlers.SearchMalls)

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
		admin.GET("warehouse", handlers.GetWarehouseInfos)
		admin.POST("sendtoslaughter", handlers.SendToSlaughter)
		admin.POST("confirmcow", handlers.ConfirmCowFromPasture)
		admin.POST("newslaughterbatch", handlers.NewSlaughterBatch)
		admin.POST("newslaughterproduct", handlers.NewSlaughterProduct)
		admin.POST("endslaughterbatch", handlers.EndSlaughter)
		admin.POST("sendtopackage", handlers.SendToPackage)
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
		pop.GET("warehouse", handlers.GetWarehouseInfos)
		//pop.POST("inwarehouse", handlers.PastureInWarehouse)
		pop.POST("send", handlers.SendToSlaughter)

		//上传传感器数据
		pop.POST("addfeedheavymetal", handlers.AddPastureFeedHeavyMetal)
		pop.POST("addfeedmycotoxins", handlers.UploadPastureFeedMycotoxins)
		pop.POST("addpasturewaterrecord", handlers.UploadPastureWaterRecord)
		pop.POST("addpasturebuffer", handlers.UploadPastureBuffer)
		pop.POST("addpasturearea", handlers.UploadPastureArea)
		pop.POST("addpasturecowhouse", handlers.UploadPastureCowHouse)
		pop.POST("addpasturebasicenvironment", handlers.UploadPastureBasicEnvironment)
		pop.POST("addpasturepaddingrequire", handlers.UploadPasturePaddingRequire)
		pop.POST("addpasturewastedwaterindex", handlers.UploadPastureWastedWaterIndex)
		pop.POST("addpasturedisinfectionrecord", handlers.UploadPastureDisinfectionRecord)

		//查询传感器数据
		pop.GET("/query/sensor/heavymetal", handlers.QueryFeedHeavyMetalData)
		pop.GET("/query/sensor/mycotoxins", handlers.QueryPastureFeedMycotoxinsData)
		pop.GET("/query/sensor/waterrecord", handlers.QueryPastureWaterRecordData)
		pop.GET("/query/sensor/buffer", handlers.QueryPastureBufferData)
		pop.GET("/query/sensor/area", handlers.QueryPastureAreaData)
		pop.GET("/query/sensor/cowhouse", handlers.QueryPastureCowHouseData)
		pop.GET("/query/sensor/basicenvironment", handlers.QueryPastureBasicEnvironmentData)
		pop.GET("/query/sensor/paddingrequire", handlers.QueryPasturePaddingRequireData)
		pop.GET("/query/sensor/wastedwaterindex", handlers.QueryPastureWastedWaterIndexData)
		pop.GET("/query/sensor/disinfectionrecord", handlers.QueryPastureDisinfectionRecordData)

		pop.POST("/upload/pasturewasteresidue", handlers.UploadPastureWasteResiduePerDay)
		pop.POST("/upload/pasturewasteodor", handlers.UploadPastureOdorPollutantsPerDay)
		pop.POST("/upload/pasturewastewater", handlers.UploadPastureWasteWaterPerDay)

		pop.GET("/query/waste", handlers.QueryWasteResidueAndOdor)

	}

	//slaughteroperator router group
	sop := fsims.Group("/slaughteroperator", middlewares.JwtAuth(), middlewares.CheckPermission())
	{
		sop.POST("commitproc", handlers.CommitSlaughterProcedure)

		sop.GET("receiverecords", handlers.GetSlaughterReceiveRecords)
		sop.GET("batches", handlers.GetSlaughterBatches)
		sop.GET("warehouserecords", handlers.GetSlaughterWarehouseRecords)

		sop.POST("receiveconfirm", handlers.ConfirmCowFromPasture)
		sop.POST("newbatch", handlers.NewSlaughterBatch)
		sop.POST("newproduct", handlers.NewSlaughterProduct)
		sop.POST("endbatch", handlers.EndSlaughter)
		sop.POST("send", handlers.SendToPackage)

		sop.POST("/upload/sensor/precoolshop", handlers.UploadPreCoolShopData)
		sop.POST("/upload/sensor/slashop", handlers.UploadSlaughterShopData)
		sop.POST("/upload/sensor/divshop", handlers.UploadDivisionShopData)
		sop.POST("/upload/sensor/acidshop", handlers.UploadAcidShopData)
		sop.POST("/upload/sensor/frozenshop", handlers.UploadFrozenShopData)

		sop.GET("/query/sensor/precoolshop", handlers.QueryPreCoolShopData)
		sop.GET("/query/sensor/slashop", handlers.QuerySlaughterShopData)
		sop.GET("/query/sensor/divshop", handlers.QueryDivisionShopData)
		sop.GET("/query/sensor/acidshop", handlers.QueryAcidShopData)
		sop.GET("/query/sensor/frozenshop", handlers.QueryFrozenShopData)

		sop.POST("/upload/sensor/waterquality", handlers.UploadSlaughterWaterQualityData)
		sop.GET("/query/sensor/waterquality", handlers.QuerySlaughterWaterQualityData)

		sop.POST("/upload/staffuniform", handlers.UploadSlaughterStaffUniformData)
		sop.POST("/upload/light", handlers.UploadSlaughterLightRecord)
		sop.GET("/query/staffuniform", handlers.QuerySlaughterStaffUniformData)
		sop.GET("/query/staffuniform", handlers.QuerySlaughterLightRecord)

		pop.POST("/upload/slaughterwasteresidue", handlers.)
		pop.POST("/upload/slaughterwasteodor")
	}

	//packoperator router group
	kop := fsims.Group("/packoperator", middlewares.JwtAuth(), middlewares.CheckPermission())
	{
		kop.POST("commitproc", handlers.CommitPackProcedure)

		kop.GET("receiverecords", handlers.GetPackageReceiveRecords)
		kop.GET("batches", handlers.GetPackageBatches)
		kop.GET("warehouserecords", handlers.GetPackageWarehouseRecords)

		kop.POST("receiveconfirm", handlers.ConfirmProductFromSlaughter)
		kop.POST("newbatch", handlers.NewPackageBatch)
		kop.POST("newproduct", handlers.NewPackageProduct)
		kop.POST("endbatch", handlers.EndPackage)
		kop.POST("pretransport", handlers.PreTransport)

	}

	//transportoperator router group
	top := fsims.Group("/transportoperator", middlewares.JwtAuth(), middlewares.CheckPermission())
	{
		top.POST("start", handlers.StartTransport)
		top.POST("end", handlers.EndTransport)
	}

	// mall
	mop := fsims.Group("malloperator", middlewares.JwtAuth())
	{
		mop.GET("goods", handlers.GetMallGoods)
	}
	return e
}
