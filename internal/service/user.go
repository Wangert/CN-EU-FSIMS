package service

import (
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

	ADMIN_USER_TYPE       = 0
	COMMON_USER_TYPE      = 1
	DATA_SOURCE_USER_TYPE = 2

	ADMIN_USER_NUMBER       = "0000"
	COMMON_USER_NUMBER      = "0001"
	DATA_SOURCE_USER_NUMBER = "0002"

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

	query.FSIMSUser.WithContext(context.Background()).Create(&fsimsUser)
	if err != nil {
		glog.Errorln("create new user error: %v", err)
		return errors.New("create new user error")
	}

	return nil
}

func generateUuid(account string, userType int) (string, error) {
	accountHash := crypto.CalculateSHA256(account, "uuid")

	var uuid string

	switch userType {
	case ADMIN_USER_TYPE:
		uuid = UUID_PREFIX + "-" + ADMIN_USER_NUMBER + "-" + accountHash
	case COMMON_USER_TYPE:
		uuid = UUID_PREFIX + "-" + COMMON_USER_NUMBER + "-" + accountHash
	case DATA_SOURCE_USER_TYPE:
		uuid = UUID_PREFIX + "-" + DATA_SOURCE_USER_NUMBER + "-" + accountHash
	default:
		return "", errors.New("user type is not exist")
	}

	return uuid, nil
}
