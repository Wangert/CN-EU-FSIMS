package mysql

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/coldchain"
	"CN-EU-FSIMS/internal/app/models/fatten"
	"CN-EU-FSIMS/internal/app/models/pasture"
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/internal/app/models/warehouse"
	"CN-EU-FSIMS/internal/config"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const GLOBALCONFIGPATH = "conf/config.yaml"

var DB *gorm.DB

func ConnectMysqlDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		glog.Errorln("connect mysql db error: %w", err)
		panic("db error")
	}

	return db
}

func Init(cfgPath string) {
	err := config.InitConfig(cfgPath)
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

	DB = ConnectMysqlDB(mysqlDsn)

	query.SetDefault(DB)
}

func AutoMigrate() {
	err := DB.AutoMigrate(&models.FSIMSUser{}, &models.Procedure{}, &models.SubProcedure{},
		&pasture.PastureProcedure{}, &pasture.PastureWater{}, &pasture.PastureFodder{}, &pasture.PastureSoil{}, &pasture.PastureFloorBedding{}, &pasture.PastureAir{},
		&pasture.PastureWasteDischarge{}, &pasture.PastureWaterPhysicalHazard{}, &pasture.PastureWaterChemicalHazard{}, &pasture.PastureWaterSensoryTraits{}, &pasture.PastureWaterBiohazard{},
		&pasture.PastureFodderPhysicalHazard{}, &pasture.PastureFodderBiohazard{}, &pasture.PastureSoilPhysicalHazard{}, &pasture.PastureSoilBiohazard{},
		&pasture.PastureFloorBeddingPhysicalHazard{}, &pasture.PastureFloorBeddingBiohazard{}, &pasture.PastureSiteDisinfectionRecord{},
		&pasture.PastureWorksuitDisinfectionRecord{}, &pasture.PastureTruckDisinfectionRecord{}, &fatten.FattenProcedure{}, &fatten.FattenWater{}, &fatten.FattenSoil{},
		&fatten.FattenWaterPhysicalHazard{}, &fatten.FattenWaterChemicalHazard{}, &fatten.FattenWaterSensoryTraits{}, &fatten.FattenWaterBiohazard{},
		&fatten.FattenSoilPhysicalHazard{}, &fatten.FattenSoilBiohazard{}, &warehouse.PastureWareHouse{}, &warehouse.SlaughterWareHouse{})

	if err != nil {
		fmt.Println(err)
	}

	err = DB.AutoMigrate(&warehouse.SlaughterReceive{})
	if err != nil {
		fmt.Println(err)
	}
	err = DB.AutoMigrate(&warehouse.PackReceive{})
	if err != nil {
		fmt.Println(err)
	}
	err = DB.AutoMigrate(&warehouse.PackWareHouse{}, &coldchain.TransportProcedureData{})
	if err != nil {
		fmt.Println(err)
	}
}

var db *sql.DB

func getmysqldb() {
	fmt.Println("***************** getmysqldb *****************")

	db, _ = sql.Open("mysql", "dbadmin:dqzgslb+1S@(10.176.35.60:3306)/testdb")
	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return
	} else {
		log.Println("connect mysql success")
	}
}

func InsertPlaceRisk(result int, place string, productinfo string, riskinfo string) bool {
	fmt.Println("***************** InsertPlaceRisk *****************")

	getmysqldb()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		log.Println("tx success")
	}
	stmt, err := tx.Prepare("INSERT INTO PlaceRisk (`Place`,`ProductInfo`,`RiskInfo`,`RiskLevel`) VALUES (?,?,?)")
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		log.Println("stmt prepare success")
	}
	res, err := stmt.Exec(place, productinfo, riskinfo, result)
	if err != nil {
		log.Fatal(err)
		return false

	} else {
		log.Println("Exec success")
	}
	tx.Commit()
	log.Println(res.LastInsertId())
	defer db.Close()
	return true
}

func InsertUser(username string, password string, usertype int, company string, telephone string, place string) bool {
	fmt.Println("***************** InsertUser *****************")

	getmysqldb()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		log.Println("tx success")
	}
	stmt, err := tx.Prepare("INSERT INTO UserInfo (`UserName`,`Password`,`UserType`,`Company`,`Telephone`,`WarnNum`,`Place`) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		log.Println("stmt prepare success")
	}
	res, err := stmt.Exec(username, password, usertype, company, telephone, 0, place)
	if err != nil {
		log.Fatal(err)
		return false

	} else {
		log.Println("Exec success")
	}
	tx.Commit()
	log.Println(res.LastInsertId())
	defer db.Close()
	return true

}

//func QueryUser(username string, password string, usertype int) models.User {
//	fmt.Println("***************** QueryUser *****************")
//
//	getmysqldb()
//	var user models.User
//	user.Login = false
//	user.Username = username
//	user.Usertype = usertype
//	sqlstr := "select Password, Company, Telephone, WarnNum, Place from UserInfo where UserName=? and UserType=?"
//	err := db.QueryRow(sqlstr, username, usertype).Scan(&user.Password, &user.Company, &user.Telephone, &user.WarnNum, &user.Place)
//	if err != nil {
//		log.Println(err)
//	}
//	if user.Password == password {
//		log.Println("query check success")
//		user.Login = true
//		fmt.Println("user num: ", user.WarnNum)
//		return user
//	}
//	defer db.Close()
//	return user
//}

//func QueryUserByPlace(place string) []int {
//	fmt.Println("***************** QueryUserByPlace *****************")
//	getmysqldb()
//	result := make([]int, 1)
//
//	sqlStr := "select ID from UserInfo where Place=?"
//	rows, err := db.Query(sqlStr, place)
//	if err != nil {
//		fmt.Printf("query failed, err:%v\n", err)
//		return nil
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var u models.User
//		err := rows.Scan(&u.ID)
//		if err != nil {
//			fmt.Printf("scan failed, err:%v\n", err)
//			return nil
//		}
//		result = append(result, u.ID)
//	}
//	return result[1:]
//}

//func InsertPlaceRiskNotification(result int, place string, productinfo string, riskinfo string, time string) bool {
//	fmt.Println("***************** NewNotification *****************")
//
//	getmysqldb()
//	tx, err := db.Begin()
//	if err != nil {
//		fmt.Println(err)
//		return false
//	} else {
//		log.Println("tx success")
//	}
//	stmt, err := tx.Prepare("INSERT INTO NotificationTable (`SafetyResult`,`ResultReason`,`NotificationTime`,`FoodChainID`,`Company`,`TargetUser`,`Detail`,`Type`) VALUES (?,?,?,?,?,?,?,?)")
//	if err != nil {
//		log.Fatal(err)
//		return false
//	} else {
//		log.Println("stmt prepare success")
//	}
//	res, err := stmt.Exec(result, riskinfo, time, productinfo, place, "", "", 2)
//	tx.Commit()
//	fmt.Println(res.LastInsertId())
//
//	response := false
//	notificationID, err := res.LastInsertId()
//	if err != nil {
//		fmt.Println(err)
//		response = false
//	} else {
//		fmt.Println("notificationID: ", notificationID)
//		response = true
//	}
//
//	TargetUserList := QueryUserByPlace(place)
//	fmt.Println("TargetUserList: ", TargetUserList)
//	for i := range TargetUserList {
//		targetUser := TargetUserList[i]
//		var notList string
//		sqlstr := "select Notification from UserInfo where ID=?"
//		err = db.QueryRow(sqlstr, targetUser).Scan(&notList)
//		if err != nil {
//			fmt.Println(err)
//		} else {
//			fmt.Println("notList: ", notList)
//		}
//
//		list := strings.Split(notList, ",")
//		if list[0] == "" {
//			fmt.Println("len list = 0")
//			list[0] = strconv.Itoa(int(notificationID))
//		} else {
//			list = append(list, strconv.Itoa(int(notificationID)))
//		}
//
//		fmt.Println("new list: ", list)
//
//		var listStr string
//		for i, v := range list {
//			fmt.Println("list item: ", list[i])
//			if i != 0 {
//				listStr += ","
//			}
//			listStr += v
//		}
//		fmt.Println("listStr: ", listStr)
//
//		updateSqlStr := "update UserInfo set Notification=? where ID=?"
//		ret, err := db.Exec(updateSqlStr, listStr, targetUser)
//		if err != nil {
//			fmt.Printf("更新失败, err:%v\n", err)
//			return false
//		}
//		rows, err := ret.RowsAffected()
//		if err != nil {
//			fmt.Printf("更新行失败, err:%v\n", err)
//			return false
//		}
//		fmt.Printf("更新成功, 更新的行数： %d.\n", rows)
//	}
//
//	return response
//
//}

func NewNotification(safetyResult string, reason string, time string, id string, company string, user string, detail models.FoodChainNode) bool {
	fmt.Println("***************** NewNotification *****************")

	getmysqldb()
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		log.Println("tx success")
	}
	stmt, err := tx.Prepare("INSERT INTO NotificationTable (`SafetyResult`,`ResultReason`,`NotificationTime`,`FoodChainID`,`Company`,`TargetUser`,`Detail`,`Type`) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		log.Println("stmt prepare success")
	}

	nodeTest := "{ \"FoodChainProcess\": \"" + detail.FoodChainProcess +
		" \", \"ProductInfo\": \"" + detail.ProductInfo +
		" \", \"Description\": \"" + detail.Description +
		" \", \"UploadCompany\": \"" + detail.UploadCompany +
		" \", \"UploadPerson\": \"" + detail.UploadPerson +
		" \", \"UploadTime\": \"" + detail.UploadTime +
		" \", \"FileHash\": \"" + detail.FileHash +
		" \", \"SafetyResult\": \"" + detail.SafetyResult + "\" }"

	notificationType := 1
	if detail.FoodChainProcess == "零售" {
		notificationType = 3
	}
	res, err := stmt.Exec(safetyResult, reason, time, id, company, user, nodeTest, notificationType)
	tx.Commit()
	fmt.Println(res.LastInsertId())

	result := false
	notificationID, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		result = false
	} else {
		fmt.Println("notificationID: ", notificationID)
		result = true
	}

	var notList string
	sqlstr := "select Notification from UserInfo where UserName=?"
	err = db.QueryRow(sqlstr, user).Scan(&notList)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("notList: ", notList)
	}

	list := strings.Split(notList, ",")
	if list[0] == "" {
		fmt.Println("len list = 0")
		list[0] = strconv.Itoa(int(notificationID))
	} else {
		list = append(list, strconv.Itoa(int(notificationID)))
	}

	fmt.Println("new list: ", list)

	var listStr string
	for i, v := range list {
		if i != 0 {
			listStr += ","
		}
		listStr += v
	}
	fmt.Println("listStr: ", listStr)

	updateSqlStr := "update UserInfo set Notification=? where UserName=?"
	ret, err := db.Exec(updateSqlStr, listStr, user)
	if err != nil {
		fmt.Printf("更新失败, err:%v\n", err)
		return false
	}
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("更新行失败, err:%v\n", err)
		return false
	}
	fmt.Printf("更新成功, 更新的行数： %d.\n", rows)

	return result
}
func GetNotification(user string) []models.Notification {
	fmt.Println("***************** GetNotification *****************")

	getmysqldb()
	var notList string

	fmt.Println("username: ", user)
	sqlstr := "select Notification from UserInfo where UserName=?"
	err := db.QueryRow(sqlstr, user).Scan(&notList)
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		fmt.Println("notList: ", notList)
	}

	list := strings.Split(notList, ",")
	fmt.Println("list: ", list)

	var notificationList []models.Notification

	sqlStr1 := "select SafetyResult, ResultReason, NotificationTime, FoodChainID, Company, Detail, TargetUser, Type from NotificationTable where NotificationID=?"

	for i, v := range list {
		fmt.Println("v: ", list[i])
		var notification models.Notification
		notification.ID = v
		var tempDetail string
		err = db.QueryRow(sqlStr1, v).Scan(&notification.Result, &notification.Reason, &notification.UploadTime, &notification.FoodChainID, &notification.UploadCompany, &tempDetail, &notification.UploadPerson, &notification.Type)
		if err != nil {
			fmt.Println("notification query failed: ", err)
			return nil
		}
		switch notification.Result {
		case "中":
			notification.Advise = "需要对产品所在环境进行一般有毒物质监测并处理，对产品进行测试，如果产品出现异常情况则建议处理整批次产品。"
		case "高":
			notification.Advise = "请调整产品的工作环境温度，并对产品进行质量抽查，温度的提高可能会对微生物、产品营养指标造成影响，请根据对产品质量抽查结果进行处理"
		}
		fmt.Println("tempDetail: ", tempDetail)
		if err = json.Unmarshal([]byte(tempDetail), &notification.Detail); err == nil {
			fmt.Println("================ str 转 struct  ========")
			fmt.Println(notification.Detail)
			fmt.Println(notification.Detail.Description)
		}
		notificationList = append(notificationList, notification)
	}
	fmt.Println("notificationList: ", notificationList)
	return notificationList
}

func ClearWarnNum(user string) {
	fmt.Println("***************** ClearWarnNum *****************")
	getmysqldb()

	sql := "update UserInfo set WarnNum=? where UserName=?"
	ret, err := db.Exec(sql, 0, user)
	if err != nil {
		fmt.Println("update err: ", err)
	}
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("更新行失败, err:%v\n", err)
	}
	fmt.Printf("更新成功, 更新的行数： %d.\n", rows)
}
