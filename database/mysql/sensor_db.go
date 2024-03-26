package mysql

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/sensorquery"
	"CN-EU-FSIMS/internal/config"
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SensorDB *gorm.DB

func ConnectSensorDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		glog.Errorln("connect mysql db error: %w", err)
		panic("db error")
	}

	return db
}

func InitSensorDB(cfgPath string) {
	err := config.InitConfig(cfgPath)
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

	SensorDB = ConnectSensorDB(mysqlDsn)

	sensorquery.SetDefault(SensorDB)
}

func AutoMigrateSensorDB() {
	err := DB.AutoMigrate(&models.Sensor{}, &models.SensorData{})
	if err != nil {
		fmt.Println(err)
	}
}
