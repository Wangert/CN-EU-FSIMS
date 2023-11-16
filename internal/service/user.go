package service

import (
	"CN-EU-FSIMS/database/mysql"
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/utils/crypto"
	"context"
	"errors"
	"github.com/golang/glog"
)

const (
	UUID_PREFIX = "FSIMSU"

	ADMIN_USER_TYPE    = 1
	CUSTOMER_USER_TYPE = 2
	OPERATOR_USER_TYPE = 3

	ADMIN_USER_NUMBER    = "0000"
	CUSTOMER_USER_NUMBER = "0001"
	OPERATOR_USER_NUMBER = "0002"

	PASSWORD_SALT = "FSIMSPS"
)

func QueryFsimsUserPwdHash(account string) (string, error) {
	u := query.FSIMSUser
	fsimsUser, err := u.WithContext(context.Background()).Where(u.Account.Eq(account)).Select(u.PasswordHash).First()
	if err != nil {
		return "", err
	}
	return fsimsUser.PasswordHash, nil
}

func QueryFsimsUserUuidAndPwdHash(account string) (uuid, password string, err error) {
	u := query.FSIMSUser
	fsimsUser, err := u.WithContext(context.Background()).Where(u.Account.Eq(account)).Select(u.UUID, u.PasswordHash).First()
	if err != nil {
		return "", "", err
	}
	return fsimsUser.UUID, fsimsUser.PasswordHash, nil
}

func AddFsimsUser(user *request.ReqUser) error {
	uuid, err := generateUuid(user.Account, user.Type)
	if err != nil {
		return err
	}

	passwordHash := crypto.CalculateSHA256(user.Password, PASSWORD_SALT)

	fsimsUser := models.FSIMSUser{
		UUID:         uuid,
		Name:         user.Name,
		Account:      user.Account,
		PasswordHash: passwordHash,
		Type:         user.Type,
		Role:         user.Role,
		Status:       1,
		Company:      user.Company,
		Phone:        user.Phone,
	}

	err = query.FSIMSUser.WithContext(context.Background()).Create(&fsimsUser)
	if err != nil {
		glog.Errorln("create new user error: %v", err)
		return errors.New("create new user error")
	}

	err = AddBasicRoleForUser(uuid)
	if err != nil {
		return err
	}
	err = AddRoleForUserWithType(uuid, user.Type)
	if err != nil {
		return err
	}

	return nil
}

func AddBasicRoleForUser(uuid string) error {
	cs, err := NewCasbinService(mysql.DB)
	if err != nil {
		return err
	}

	return cs.AddRoleForUserWithUUID(uuid, "user")
}

func AddRoleForUserWithType(uuid string, ttype int) error {
	var roleName string

	switch ttype {
	case ADMIN_USER_TYPE:
		roleName = "admin"
	case CUSTOMER_USER_TYPE:
		roleName = "customer"
	case OPERATOR_USER_TYPE:
		roleName = "operator"
	default:
		return errors.New("user type is not exist")
	}

	cs, err := NewCasbinService(mysql.DB)
	if err != nil {
		return err
	}

	return cs.AddRoleForUserWithUUID(uuid, roleName)
}

func generateUuid(account string, userType int) (string, error) {
	accountHash := crypto.CalculateSHA256(account, "uuid")

	var uuid string

	switch userType {
	case ADMIN_USER_TYPE:
		uuid = UUID_PREFIX + "-" + ADMIN_USER_NUMBER + "-" + accountHash
	case CUSTOMER_USER_TYPE:
		uuid = UUID_PREFIX + "-" + CUSTOMER_USER_NUMBER + "-" + accountHash
	case OPERATOR_USER_TYPE:
		uuid = UUID_PREFIX + "-" + OPERATOR_USER_NUMBER + "-" + accountHash
	default:
		return "", errors.New("user type is not exist")
	}

	return uuid, nil
}