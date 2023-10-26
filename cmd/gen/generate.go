package main

import (
	"CN-EU-FSIMS/internal/config"
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const DEFAULTDSN = "root:root@tcp(127.0.0.1:3306)/FSIMS?charset=utf8mb4&parseTime=True"
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

	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}

func main() {
	err := config.InitConfig(GLOBALCONFIGPATH)
	if err != nil {
		glog.Error(err)
		panic(err)
	}

	mysqlDsn := viper.GetString("mysql_dsn")
	//println("MysqlDsn: ", mysqlDsn)
	dbName := viper.GetString("mysql_database")

	if dbName == "" {
		dbName = "default_dal"
	} else {
		dbName = dbName + "_dal"
	}

	if mysqlDsn != "" {
		generateModelAndQueryWithDB(mysqlDsn, dbName)
	} else {
		generateModelAndQueryWithDB(DEFAULTDSN, dbName)
	}
}
