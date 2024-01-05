package main

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/coldchain"
	"CN-EU-FSIMS/internal/app/models/pack"
	"CN-EU-FSIMS/internal/app/models/pasture"
	"CN-EU-FSIMS/internal/app/models/premortem"
	"CN-EU-FSIMS/internal/app/models/product"
	"CN-EU-FSIMS/internal/app/models/sell"
	"CN-EU-FSIMS/internal/app/models/slaughter"
	"CN-EU-FSIMS/internal/app/models/warehouse"
	"CN-EU-FSIMS/internal/config"
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const DEFAULTDSN = "root:root@tcp(127.0.0.1:3306)/fsims?charset=utf8mb4&parseTime=True"
const GLOBALCONFIGPATH = "conf/config.yaml"

func generateModelAndQueryWithDB(dsn string, dbName string) {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("Connect mysql fail: %w", err))
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "./cmd/" + dbName + "/query",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db)

	g.ApplyBasic(&models.Procedure{})

	g.ApplyBasic(&models.FSIMSUser{})

	g.ApplyBasic(&models.Log{})

	//g.ApplyBasic(coldchain.TransportProcedureData{})

	g.ApplyBasic(&pasture.PastureHouse{}, &slaughter.SlaughterHouse{}, &pack.PackageHouse{}, &coldchain.TransportVehicle{}, &sell.Mall{})
	g.ApplyBasic(&product.Cow{}, &pasture.FeedingBatch{}, &warehouse.PastureWarehouse{},
		&warehouse.SlaughterReceiveRecord{}, &slaughter.SlaughterBatch{}, &product.SlaughterProduct{}, &warehouse.SlaughterWarehouse{},
		&warehouse.PackageReceiveRecord{}, &pack.PackageBatch{}, &product.PackageProduct{}, &warehouse.PackWarehouse{}, &coldchain.TransportBatch{}, &coldchain.PackageProductAndTransportPIDMap{}, &product.MallGood{})

	// 牧场传感器数据
	g.ApplyBasic(&pasture.PastureFeedMycotoxins{}, &pasture.Afb1{}, &pasture.Don{}, &pasture.T2toxin{}, &pasture.T2VomZea{},
		&pasture.PastureBuffer{}, &pasture.PastureArea{}, &pasture.PastureBasicEnvironment{}, &pasture.PasturePaddingRequire{},
		&pasture.PastureFeedHeavyMetal{}, &pasture.PastureFeedAs{}, &pasture.PastureFeedPb{}, &pasture.PastureFeedCd{}, &pasture.PastureFeedCr{},
		&pasture.PastureWastedWaterIndex{}, &pasture.PastureDisinfectionRecord{}, &pasture.PastureWaterRecord{}, &pasture.PastureOapGci{},
		&pasture.PastureToxIndex{}, &pasture.PastureMicroIndex{}, &pasture.CowHouse{}, &pasture.TotalOdorPollutantsPasturePerDay{}, &pasture.TotalWasteResiduePasturePerDay{}, &pasture.TotalWastedWaterPasturePerDay{})

	//宰前传感器数据
	g.ApplyBasic(&premortem.WaitingSlaughterCircle{}, &premortem.Gps{}, &premortem.WaitingSlaughterCircleGerms{})

	// 屠宰场传感器数据
	g.ApplyBasic(&slaughter.PreCoolShop{}, &slaughter.SlaShop{}, &slaughter.DivShop{}, &slaughter.AcidShop{}, &slaughter.FroShop{}, &slaughter.PackShop{}, &slaughter.StaUni{}, &slaughter.SlaughterLightRecord{},
		&slaughter.SlaughterWaterQualityMon{}, &slaughter.SlaughterWaterMicroIndex{}, &slaughter.OapGciSla{}, &slaughter.MicroIndexWaterMonSla{}, &slaughter.SlaughterToxinIndex{}, &slaughter.SlaughterWaterToxinIndex{},
		&slaughter.TotalOdorPollutantsSlaughterPerDay{}, &slaughter.TotalWasteResidueSlaughterPerDay{})

	//以牛为对象的数据，暂时不用
	g.ApplyBasic(&slaughter.SlaughterDisinfectHotWaterTempMoni{}, &slaughter.SlaughterStun{}, &slaughter.BleedElectronic{}, &slaughter.PreSlaQuanPic{}, &slaughter.SlaughterAnalAfterSlaQuanCar{}, &slaughter.AnalMeatPhMoni{}, &slaughter.AnalCutWeight{},
		&slaughter.ToNumGermMon{}, &slaughter.AirNumGermMon{})

	g.ApplyBasic(
		&pasture.PastureProcedure{}, &pasture.PastureWater{}, &pasture.PastureFodder{}, &pasture.PastureSoil{}, &pasture.PastureFloorBedding{}, &pasture.PastureAir{},
		&pasture.PastureWasteDischarge{}, &pasture.PastureWaterPhysicalHazard{}, &pasture.PastureWaterChemicalHazard{}, &pasture.PastureWaterSensoryTraits{}, &pasture.PastureWaterBiohazard{},
		&pasture.PastureFodderPhysicalHazard{}, &pasture.PastureFodderBiohazard{}, &pasture.PastureSoilPhysicalHazard{}, &pasture.PastureSoilBiohazard{},
		&pasture.PastureFloorBeddingPhysicalHazard{}, &pasture.PastureFloorBeddingBiohazard{}, &pasture.PastureSiteDisinfectionRecord{},
		&pasture.PastureWorksuitDisinfectionRecord{}, &pasture.PastureTruckDisinfectionRecord{})

	//g.ApplyBasic( &fatten.FattenProcedure{}, &fatten.FattenWater{}, &fatten.FattenSoil{},
	//	&fatten.FattenWaterPhysicalHazard{}, &fatten.FattenWaterChemicalHazard{}, &fatten.FattenWaterSensoryTraits{}, &fatten.FattenWaterBiohazard{},
	//	&fatten.FattenSoilPhysicalHazard{}, &fatten.FattenSoilBiohazard{})

	g.ApplyBasic(&models.Notification{}, &models.MonitoringTimeRecord{})
	g.Execute()
}

func main() {
	err := config.InitConfig(GLOBALCONFIGPATH)
	if err != nil {
		glog.Error(err)
		panic(err)
	}

	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	ipAddr := viper.GetString("mysql.ip_addr")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")

	mysqlDsn := username + ":" + password + "@tcp(" + ipAddr + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True"

	if database == "" {
		database = "default_dal"
	} else {
		database = database + "_dal"
	}

	if mysqlDsn != "" {
		generateModelAndQueryWithDB(mysqlDsn, database)
	} else {
		generateModelAndQueryWithDB(DEFAULTDSN, database)
	}
}
