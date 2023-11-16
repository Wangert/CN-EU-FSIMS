package main

import (
	"CN-EU-FSIMS/database/mysql"
	middleware "CN-EU-FSIMS/internal/app/middlewares"
	"CN-EU-FSIMS/internal/app/routers"
	"CN-EU-FSIMS/internal/common"
	"CN-EU-FSIMS/internal/config"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"net/http"
)

//const MySQLDSN = "root:root@tcp(127.0.0.1:3306)/bft_diagnosis?charset=utf8mb4&parseTime=True"

func main() {
	fmt.Println("China-Europe Food Safety Intelligent Management and Decision Support System.")

	flag.Set("logtostderr", "true")
	err := config.InitConfig(common.CONFIG_PATH)
	if err != nil {
		panic("init config error: " + err.Error())
	}

	mysql.Init(common.CONFIG_PATH)

	// set gin run mode
	gin.SetMode(viper.GetString("runmode"))
	//配置路由和中间件
	r := routers.Load(gin.New(), middleware.Cors)

	listenAddr := viper.GetString("server.ip_addr") + ":" + viper.GetString("server.port")
	glog.Infof(http.ListenAndServe(listenAddr, r).Error())

}
