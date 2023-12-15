package main

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/coldchain"
	"CN-EU-FSIMS/internal/app/models/fatten"
	"CN-EU-FSIMS/internal/app/models/pasture"
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
const GLOBALCONFIGPATH = "../../conf/config.yaml"

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

<<<<<<< HEAD
	g.ApplyBasic(&models.Logs{})
=======
	g.ApplyBasic(&warehouse.PastureWareHouse{}, &warehouse.SlaughterWareHouse{})
	g.ApplyBasic(&warehouse.SlaughterReceive{})
	g.ApplyBasic(&warehouse.PackReceive{})
	g.ApplyBasic(&warehouse.PackWareHouse{}, coldchain.TransportProcedureData{})
>>>>>>> wangert

	g.ApplyBasic(&models.SubProcedure{},
		&pasture.PastureProcedure{}, &pasture.PastureWater{}, &pasture.PastureFodder{}, &pasture.PastureSoil{}, &pasture.PastureFloorBedding{}, &pasture.PastureAir{},
		&pasture.PastureWasteDischarge{}, &pasture.PastureWaterPhysicalHazard{}, &pasture.PastureWaterChemicalHazard{}, &pasture.PastureWaterSensoryTraits{}, &pasture.PastureWaterBiohazard{},
		&pasture.PastureFodderPhysicalHazard{}, &pasture.PastureFodderBiohazard{}, &pasture.PastureSoilPhysicalHazard{}, &pasture.PastureSoilBiohazard{},
		&pasture.PastureFloorBeddingPhysicalHazard{}, &pasture.PastureFloorBeddingBiohazard{}, &pasture.PastureSiteDisinfectionRecord{},
		&pasture.PastureWorksuitDisinfectionRecord{}, &pasture.PastureTruckDisinfectionRecord{}, &fatten.FattenProcedure{}, &fatten.FattenWater{}, &fatten.FattenSoil{},
		&fatten.FattenWaterPhysicalHazard{}, &fatten.FattenWaterChemicalHazard{}, &fatten.FattenWaterSensoryTraits{}, &fatten.FattenWaterBiohazard{},
		&fatten.FattenSoilPhysicalHazard{}, &fatten.FattenSoilBiohazard{})
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
