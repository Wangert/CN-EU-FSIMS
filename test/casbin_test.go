package test

import (
	"CN-EU-FSIMS/database/mysql"
	"CN-EU-FSIMS/internal/service"
	"flag"
	"fmt"
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/golang/glog"
)

const (
	CASBIN_MODEL_PATH       = "../conf/casbin/fsims_rbac_model.conf"
	CASBIN_POLICY_PATH      = "../conf/casbin/fsims_rbac_policy.csv"
	CASBIN_ROLE_POLICY_PATH = "../conf/casbin/fsims_rbac_role_policy.csv"
	CASBIN_USER_ROLE_PATH   = "../conf/casbin/fsims_rbac_user_role.csv"
)

func TestPolicy(t *testing.T) {
	e, err := casbin.NewEnforcer(CASBIN_MODEL_PATH, CASBIN_POLICY_PATH)
	if err != nil {
		panic("new casbin enforcer error!")
	}

	sub := "wangert"
	obj := "/user/register"
	act := "POST"

	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		panic("enforce error")
	}

	if ok {
		fmt.Println("pass")
	} else {
		fmt.Println("deny")
	}
}

func TestPolicyByDB(t *testing.T) {
	flag.Set("logtostderr", "true")

	mysql.Init(TESTCONFIGPATH)
	cs, err := service.NewCasbinService(mysql.DB)
	if err != nil {
		panic(err.Error())
	}

	//roles := cs.GetAllRoles()
	//glog.Infoln("Roles:")
	//glog.Infoln(roles)
	//
	//objects := cs.GetAllObjects()
	//glog.Infoln("Objects:")
	//glog.Infoln(objects)

	pp := service.PendingPolicy{
		UUID:   "wangert",
		Url:    "/fsims/user/register",
		Method: "POST",
	}

	ok, err := cs.CheckPermission(&pp)
	if err != nil {
		glog.Errorln("check permission error!")
	}

	if ok {
		glog.Infoln("pass!")
	} else {
		glog.Infoln("deny!")
	}

}

func TestCasbinAndGorm(t *testing.T) {
	flag.Set("logtostderr", "true")

	//err := config.InitConfig(common.CONFIG_PATH)
	//if err != nil {
	//	panic("init config error: " + err.Error())
	//}

	mysql.Init(TESTCONFIGPATH)
	cs, err := service.NewCasbinService(mysql.DB)
	if err != nil {
		panic(err.Error())
	}
	err = cs.ImportPoliciesIntoDB(CASBIN_ROLE_POLICY_PATH, CASBIN_USER_ROLE_PATH)
	if err != nil {
		panic(err.Error())
	}
}
