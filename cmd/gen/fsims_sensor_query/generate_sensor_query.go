package main

import (
	"CN-EU-FSIMS/internal/app/models"
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

func generateSensorModelAndQueryWithDB(dsn string, dbName string) {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("Connect mysql fail: %w", err))
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "cmd/" + dbName + "/sensorquery",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db)

	g.ApplyBasic(&models.Sensor{}, &models.SensorData{})
	g.Execute()
}

func main() {
	err := config.InitConfig(GLOBALCONFIGPATH)
	if err != nil {
		glog.Error(err)
		panic(err)
	}

	username := viper.GetString("sensor_db.username")
	password := viper.GetString("sensor_db.password")
	ipAddr := viper.GetString("sensor_db.ip_addr")
	port := viper.GetString("sensor_db.port")
	database := viper.GetString("sensor_db.database")

	mysqlDsn := username + ":" + password + "@tcp(" + ipAddr + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True"

	if database == "" {
		database = "fsims_sensor_default_dal"
	} else {
		database = database + "_dal"
	}

	if mysqlDsn != "" {
		generateSensorModelAndQueryWithDB(mysqlDsn, database)
	} else {
		generateSensorModelAndQueryWithDB(DEFAULTDSN, database)
	}
}
