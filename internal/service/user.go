package service

import (
	"CN-EU-FSIMS/database/mysql"
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/utils/crypto"
	"context"
	"errors"
	"gorm.io/gorm"

	"github.com/golang/glog"
)

const (
	UUID_PREFIX = "FSIMSU"

	ADMIN_USER_TYPE              = 1
	CUSTOMER_USER_TYPE           = 2
	OPERATOR_USER_TYPE           = 3
	PASTURE_OPERATOR_USER_TYPE   = 4
	SLAUGHTER_OPERATOR_USER_TYPE = 5
	TRANSPORT_OPERATOR_USER_TYPE = 6

	ADMIN_USER_NUMBER              = "0000"
	CUSTOMER_USER_NUMBER           = "0001"
	OPERATOR_USER_NUMBER           = "0002"
	PASTURE_OPERATOR_USER_NUMBER   = "0003"
	SLAUGHTER_OPERATOR_USER_NUMBER = "0004"
	TRANSPORT_OPERATOR_USER_NUMBER = "0005"

	PASSWORD_SALT = "FSIMSPS"
	INIT_PASSWORD = "123456"
)

func CreateFsimUser(user *request.ReqUpdateUser) *models.FSIMSUser {
	var u models.FSIMSUser
	u.Name = user.Name
	u.Phone = user.Phone
	u.Type = user.Type
	u.Role = user.Role
	u.Account = user.Account
	u.Company = user.Company
	u.PasswordHash = crypto.CalculateSHA256(user.Password, PASSWORD_SALT)
	return &u
}

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
	case PASTURE_OPERATOR_USER_TYPE:
		roleName = "pastureoperator"
	case SLAUGHTER_OPERATOR_USER_TYPE:
		roleName = "slaughteroperator"
	case TRANSPORT_OPERATOR_USER_TYPE:
		roleName = "transportoperator"
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
	case PASTURE_OPERATOR_USER_TYPE:
		uuid = UUID_PREFIX + "-" + PASTURE_OPERATOR_USER_NUMBER + "-" + accountHash
	case SLAUGHTER_OPERATOR_USER_TYPE:
		uuid = UUID_PREFIX + "-" + SLAUGHTER_OPERATOR_USER_NUMBER + "-" + accountHash
	case TRANSPORT_OPERATOR_USER_TYPE:
		uuid = UUID_PREFIX + "-" + TRANSPORT_OPERATOR_USER_NUMBER + "-" + accountHash
	default:
		return "", errors.New("user type is not exist")
	}

	return uuid, nil
}

func AddFsimsUserByAdmin(user *request.ReqAddUser) error {
	uuid, err := generateUuid(user.Account, user.Type)
	if err != nil {
		return err
	}

	//密码初始为fsims+123456
	passwordHash := crypto.CalculateSHA256(INIT_PASSWORD, PASSWORD_SALT)

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
		glog.Errorln("create new user error by admin")
		return errors.New("create new user error")
	}
	return nil
}

func ResetFsimsPassWord(account string) error {
	p := crypto.CalculateSHA256(INIT_PASSWORD, PASSWORD_SALT)
	u := query.FSIMSUser
	_, err := u.WithContext(context.Background()).Where(u.Account.Eq(account)).Update(u.PasswordHash, p)
	if err != nil {
		glog.Errorln("Reset the user password error: %v", err)
		return errors.New("Reset the user password error")
	}
	return nil
}

func UpdateFsimsUser(user *request.ReqUpdateUser) error {
	//check whether user exist
	u := query.FSIMSUser
	_, err := u.WithContext(context.Background()).Where(u.Account.Eq(user.Account)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		glog.Errorln("Update a user error: %v", err)
		return errors.New("update a user error: user not existed")
	}
	if err != nil {
		glog.Errorln("Update a user error: %v", err)
		return errors.New("update a user error")
	}
	//update
	var fuu = CreateFsimUser(user)
	_, err = u.WithContext(context.Background()).Where(u.Account.Eq(user.Account)).Updates(&fuu)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFsimUser(account string) error {
	u := query.FSIMSUser
	res, err := u.WithContext(context.Background()).Where(u.Account.Eq(account)).Delete()
	_ = res.RowsAffected
	if errors.Is(err, gorm.ErrRecordNotFound) {
		glog.Errorln("Delete a user error: %v", err)
		return errors.New("delete a user error: user not existed")
	}
	if err != nil {
		glog.Errorln("Update a user error: %v", err)
		return errors.New("delete a user error")
	}
	return nil
}
