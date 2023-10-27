package main

import (
	"CN-EU-FSIMS/database/mysql"
	"fmt"
)

//const MySQLDSN = "root:root@tcp(127.0.0.1:3306)/bft_diagnosis?charset=utf8mb4&parseTime=True"

func main() {
	fmt.Println("China-Europe Food Safety Intelligent Management and Decision Support System.")

	mysql.Init(mysql.GLOBALCONFIGPATH)
	//mysql.CreateTest()
	//mysql.DeleteTest()
	mysql.AutoMigrate()
}
