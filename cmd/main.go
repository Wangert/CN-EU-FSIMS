package main

import (
	"CN-EU-FSIMS/database/mysql"
	"CN-EU-FSIMS/fabric"
	"CN-EU-FSIMS/internal/app/middlewares"
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

func main() {
	fmt.Println("China-Europe Food Safety Intelligent Management and Decision Support System.")

	flag.Set("logtostderr", "true")
	err := config.InitConfig(common.CONFIG_PATH)
	if err != nil {
		panic("init config error: " + err.Error())
	}
	// 初始化mysql
	mysql.Init(common.CONFIG_PATH)

	// 初始化fabric sdk client
	fabric.InitFabricSDKClient()

	// 设置gin运行模式
	gin.SetMode(viper.GetString("runmode"))
	//配置路由和中间件
	//middleware.Cors
	r := routers.Load(gin.New(), middlewares.Cors)

	listenAddr := viper.GetString("server.ip_addr") + ":" + viper.GetString("server.port")
	glog.Infof(http.ListenAndServe(listenAddr, r).Error())
}
