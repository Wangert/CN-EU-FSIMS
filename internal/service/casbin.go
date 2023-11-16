package service

import (
	"CN-EU-FSIMS/utils"
	"errors"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/golang/glog"
	"gorm.io/gorm"
)

const (
	CASBIN_MODEL_PATH       = "conf/casbin/fsims_rbac_model.conf"
	CASBIN_ROLE_POLICY_PATH = "conf/casbin/fsims_rbac_role_policy.csv"
	CASBIN_USER_ROLE_PATH   = "conf/casbin/fsims_rbac_user_role.csv"
)

type CasbinService struct {
	enforcer *casbin.Enforcer
	adapter  *gormadapter.Adapter
}

type PendingPolicy struct {
	UUID   string
	Url    string
	Method string
}

func NewCasbinService(db *gorm.DB) (*CasbinService, error) {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		glog.Errorln("new casbin adapter error! ")
		return nil, err
	}

	enforcer, err := casbin.NewEnforcer(CASBIN_MODEL_PATH, adapter)
	if err != nil {
		glog.Errorln("new casbin enforcer error!")
		return nil, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, err
	}

	return &CasbinService{
		enforcer,
		adapter,
	}, nil
}

func (cs *CasbinService) CheckPermission(pp *PendingPolicy) (bool, error) {
	return cs.enforcer.Enforce(pp.UUID, pp.Url, pp.Method)
}

func (cs *CasbinService) AddRoleForUserWithUUID(uuid, roleName string) error {
	ok, err := cs.enforcer.AddGroupingPolicy(uuid, roleName)
	if err != nil || !ok {
		return errors.New("add grouping policy error")
	}

	return nil
}

func (cs *CasbinService) GetAllRoles() []string {
	return cs.enforcer.GetAllRoles()
}

func (cs *CasbinService) GetAllObjects() []string {
	return cs.enforcer.GetAllObjects()
}

func (cs *CasbinService) ImportPoliciesIntoDB(policyPath, userPath string) error {
	policies, err := utils.ReadCSVFile(policyPath)
	if err != nil {
		return err
	}
	glog.Infoln("policies:")
	glog.Infoln(policies)

	_, err = cs.enforcer.AddPolicies(policies)
	if err != nil {
		return err
	}

	users, err := utils.ReadCSVFile(userPath)
	if err != nil {
		return err
	}

	glog.Infoln("users:")
	glog.Infoln(users)

	_, err = cs.enforcer.AddGroupingPolicies(users)
	if err != nil {
		return err
	}

	err = cs.enforcer.SavePolicy()
	if err != nil {
		return err
	}

	//ok, _ := cs.enforcer.Enforce("youqu", "/admin/hh", "POST")
	//if ok {
	//	glog.Infoln("pass!")
	//	glog.Infoln(cs.enforcer.GetGroupingPolicy())
	//	glog.Infoln(cs.enforcer.GetFilteredGroupingPolicy(0, "youqu"))
	//	glog.Infoln(cs.enforcer.GetFilteredGroupingPolicy(0, "admin"))
	//} else {
	//	glog.Infoln("deny!")
	//	glog.Infoln(cs.enforcer.GetGroupingPolicy())
	//	glog.Infoln(cs.enforcer.GetFilteredGroupingPolicy(0, "youqu"))
	//	glog.Infoln(cs.enforcer.GetFilteredGroupingPolicy(0, "admin"))
	//}

	return nil
}
