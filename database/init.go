package database

import (
	"CN-EU-FSIMS/database/mysql"
	"CN-EU-FSIMS/internal/common"
)

func InitDB() {
	mysql.Init(common.CONFIG_PATH)
	mysql.InitSensorDB(common.CONFIG_PATH)
	mysql.InitImgDB(common.CONFIG_PATH)
}
