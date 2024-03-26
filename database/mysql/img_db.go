package mysql

import (
	"CN-EU-FSIMS/internal/app/models/imgquery"
	"CN-EU-FSIMS/internal/config"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ImgDB *gorm.DB

func ConnectImgDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		glog.Errorln("connect mysql db error: %w", err)
		panic("db error")
	}

	return db
}

func InitImgDB(cfgPath string) {
	err := config.InitConfig(cfgPath)
	if err != nil {
		glog.Error(err)
		panic(err)
	}

	username := viper.GetString("img_db.username")
	password := viper.GetString("img_db.password")
	ipAddr := viper.GetString("img_db.ip_addr")
	port := viper.GetString("img_db.port")
	database := viper.GetString("img_db.database")

	mysqlDsn := username + ":" + password + "@tcp(" + ipAddr + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True"

	ImgDB = ConnectImgDB(mysqlDsn)

	imgquery.SetDefault(ImgDB)
}

func AutoMigrateImgDB() {
}
