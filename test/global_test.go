package test

import (
	"CN-EU-FSIMS/database/mysql"
	"CN-EU-FSIMS/internal/service"
	"flag"
	"testing"
)

func TestInitDatabaseAndCasbin(t *testing.T) {
	flag.Set("logtostderr", "true")
	mysql.Init(TESTCONFIGPATH)
	mysql.AutoMigrate()

	cs, err := service.NewCasbinService(mysql.DB)
	if err != nil {
		panic(err.Error())
	}
	err = cs.ImportPoliciesIntoDB(CASBIN_ROLE_POLICY_PATH, CASBIN_USER_ROLE_PATH)
	if err != nil {
		panic(err.Error())
	}
}
