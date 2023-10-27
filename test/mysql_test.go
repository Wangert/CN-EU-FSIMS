package test

import (
	"CN-EU-FSIMS/database/mysql"
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/query"
	"context"
	"fmt"
	"github.com/golang/glog"
	"testing"
	"time"
)

const TESTCONFIGPATH = "../conf/config.yaml"

func TestMigrate(t *testing.T) {
	mysql.Init(TESTCONFIGPATH)
	mysql.AutoMigrate()
}

func TestCreateFsimsUser(t *testing.T) {
	mysql.Init(TESTCONFIGPATH)

	newUser := models.FSIMSUser{
		UUID:         "UUID_test",
		Name:         "Name_test",
		Account:      "Account_test",
		PasswordHash: "PasswordHash_test",
		Type:         "Type_test",
		Status:       "Status_test",
		Company:      "Company_test",
		Phone:        "Phone_test",
	}

	err := query.FSIMSUser.WithContext(context.Background()).Create(&newUser)
	if err != nil {
		glog.Errorln("create new user error: %v", err)
	}

	fmt.Println("Create user Successful")

	queryUsers, err := query.FSIMSUser.WithContext(context.Background()).Find()
	if err != nil {
		glog.Errorln(err)
	} else {
		for _, r := range queryUsers {
			fmt.Println(*r)
		}
	}
}

func TestFindFsimsUsers(t *testing.T) {
	mysql.Init(TESTCONFIGPATH)

	ics, err := query.FSIMSUser.WithContext(context.Background()).Find()
	if err != nil {
		panic(err)
	} else {
		for _, ic := range ics {
			fmt.Println(*ic)
		}
	}
}

func TestCreateIndustrialChainWithProcedures(t *testing.T) {
	mysql.Init(TESTCONFIGPATH)

	procedure1 := models.Procedure{
		PID:          "PID-3",
		Name:         "Production",
		SerialNumber: 1,
		PrePID:       "PID-0",
		ICID:         "ICID-11111",
	}

	procedure2 := models.Procedure{
		PID:          "PID-4",
		Name:         "Slaughter",
		SerialNumber: 2,
		PrePID:       "PID-1",
		ICID:         "ICID-11111",
	}

	ic := models.IndustrialChain{
		ICID:               "ICID-22222",
		Type:               "fish",
		State:              2,
		StartTimestamp:     time.Now(),
		CompletedTimestamp: time.Now(),
		Procedures: []models.Procedure{
			procedure1, procedure2,
		},
	}

	err := query.IndustrialChain.WithContext(context.Background()).Create(&ic)
	if err != nil {
		panic(err)
	}

	fmt.Println("create industrial chain successful!")
}

func TestFindIndustrialChains(t *testing.T) {
	mysql.Init(TESTCONFIGPATH)

	ics, err := query.IndustrialChain.WithContext(context.Background()).Find()
	if err != nil {
		panic(err)
	} else {
		for _, ic := range ics {
			fmt.Println(*ic)
		}
	}
}

func TestFindAllProdecuresWithIndustrialChain(t *testing.T) {
	mysql.Init(TESTCONFIGPATH)

	ic, err := query.IndustrialChain.WithContext(context.Background()).First()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("industrial_chain:")
		fmt.Println(*ic)
	}

	fmt.Println("=======================================")
	procedures, err := query.IndustrialChain.Procedures.WithContext(context.Background()).Model(ic).Find()
	if err != nil {
		panic(err)
	} else {
		for _, procedure := range procedures {
			fmt.Println(*procedure)
		}
	}
}

func TestDeleteIndustrialChains(t *testing.T) {
	infoResult, err := query.FSIMSUser.WithContext(context.Background()).Where(query.IndustrialChain.ID.Gt(0)).Delete()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Delete IndustrialChains(ID > 0) Successful")
		fmt.Println(infoResult)
	}
}

func TestDeleteProcedures(t *testing.T) {
	infoResult, err := query.FSIMSUser.WithContext(context.Background()).Where(query.Procedure.ID.Gt(0)).Delete()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Delete Procedures(ID > 0) Successful")
		fmt.Println(infoResult)
	}
}
