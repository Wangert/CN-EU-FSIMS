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

const REMOTE_DEFAULT_DSN = "sdau:sdau@tcp(182.92.99.82:3306)/img?charset=utf8mb4&parseTime=True"
const REMOTE_GLOBAL_CONFIG_PATH = "conf/remote.yaml"

func generateModelAndQueryWithRemoteDB(dsn, dbName string) {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("Connect mysql fail: %w", err))
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "./cmd/" + dbName + "/remote_query",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db)

	g.ApplyBasic(&models.Spectra{})

	g.ApplyBasic(&models.UploadImgs{})

	g.Execute()
}

func main() {
	err := config.InitConfig(REMOTE_GLOBAL_CONFIG_PATH)
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
	fmt.Println("mysqlDsn", mysqlDsn)
	if database == "" {
		database = "default_dal"
	} else {
		database = database + "_dal"
	}

	if mysqlDsn != "" {
		fmt.Println("11111111111111")
		generateModelAndQueryWithRemoteDB(mysqlDsn, database)
	} else {
		generateModelAndQueryWithRemoteDB(REMOTE_DEFAULT_DSN, database)
	}
}
