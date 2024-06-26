package service

import (
	"CN-EU-FSIMS/database/mysql"
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/pasture"
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/internal/app/models/slaughter"
	"CN-EU-FSIMS/utils/crypto"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"time"

	"github.com/golang/glog"
)

const (
	UUID_PREFIX = "FSIMSU"

	ADMIN_USER_TYPE              = 1
	CUSTOMER_USER_TYPE           = 2
	OPERATOR_USER_TYPE           = 3
	PASTURE_OPERATOR_USER_TYPE   = 4
	SLAUGHTER_OPERATOR_USER_TYPE = 5
	PACK_OPERATOR_USER_TYPE      = 6
	TRANSPORT_OPERATOR_USER_TYPE = 7
	BUYER_OPERATOR_USER_TYPE     = 8

	ADMIN_USER_NUMBER              = "0000"
	CUSTOMER_USER_NUMBER           = "0001"
	OPERATOR_USER_NUMBER           = "0002"
	PASTURE_OPERATOR_USER_NUMBER   = "0003"
	SLAUGHTER_OPERATOR_USER_NUMBER = "0004"
	PACK_OPERATOR_USER_NUMBER      = "0005"
	TRANSPORT_OPERATOR_USER_NUMBER = "0006"
	BUYER_OPERATOR_USER_NUMBER     = "0007"

	PASSWORD_SALT = "FSIMSPS"
	INIT_PASSWORD = "Fsims123456!"
)

func QueryEndFeedCow() (int64, error) {
	q := query.Q.Cow
	count, err := q.WithContext(context.Background()).Where(q.State.Neq(1)).Where(q.State.Neq(2)).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func QueryEndSlaCow() (int64, error) {
	q := query.Q.Cow
	count, err := q.WithContext(context.Background()).Where(q.State.Neq(1)).Where(q.State.Neq(2)).Where(q.State.Neq(3)).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func QueryBatchesEnd() (int64, error) {
	q := query.Q.TransportBatch
	count, err := q.WithContext(context.Background()).Where(q.State.Eq(END_STATE_BATCH_TRANS)).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func QuerySlaughterTrashFifteenDays(r *request.ReqTrashFifteenDays) ([]slaughter.AllSlaughtersTrashDisposalInfo, int64, error) {
	startTime := time.Unix(r.StartTimeStamp, 0).Truncate(24 * time.Hour).UTC()
	endTime := time.Unix(r.EndTimeStamp, 0).Truncate(24 * time.Hour).UTC()
	q1 := query.Q.AllSlaughtersTrashDisposal
	res1, err := q1.WithContext(context.Background()).Where(q1.TimeStamp.Between(startTime, endTime)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(res1)
	infos := make([]slaughter.AllSlaughtersTrashDisposalInfo, count)
	for i, res := range res1 {
		infos[i] = slaughter.ToAllSlaughtersTrashDisposalInfo(res)
	}
	return infos, int64(count), nil
}

func QueryPastureTrashFifteenDays(r *request.ReqTrashFifteenDays) ([]pasture.AllPasturesTrashDisposalInfo, int64, error) {
	startTime := time.Unix(r.StartTimeStamp, 0).Truncate(24 * time.Hour).UTC().Local()
	endTime := time.Unix(r.EndTimeStamp, 0).Truncate(24 * time.Hour).UTC().Local()

	q2 := query.Q.AllPasturesTrashDisposal
	res2, err := q2.WithContext(context.Background()).Where(q2.TimeStamp.Between(startTime, endTime)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(res2)
	infos := make([]pasture.AllPasturesTrashDisposalInfo, count)
	for i, res := range res2 {
		infos[i] = pasture.ToAllPasturesTrashDisposalInfo(res)
	}
	return infos, int64(count), nil
}
func QueryTrashPerDay(r *request.ReqTrashPerDay) (models.TrashDisposalPerDayInfo, error) {
	t := time.Unix(r.TimeStamp, 0).Truncate(24 * time.Hour)
	fmt.Println("查询的时间戳是：", t)
	//检查slaughter和pasture是否存在
	p, err := CheckPastureTrashIsExisted(t)
	if err != nil {
		return models.TrashDisposalPerDayInfo{}, err
	}
	s, err := CheckSlaughterTrashIsExisted(t)

	if err != nil {
		return models.TrashDisposalPerDayInfo{}, err
	}

	if p && s {
		ps := query.Q.AllPasturesTrashDisposal
		pas, err1 := ps.WithContext(context.Background()).Where(ps.TimeStamp.Eq(t)).First()
		if err1 != nil {
			return models.TrashDisposalPerDayInfo{}, err1
		}

		ss := query.Q.AllSlaughtersTrashDisposal
		sla, err2 := ss.WithContext(context.Background()).Where(ss.TimeStamp.Eq(t)).First()
		if err2 != nil {
			return models.TrashDisposalPerDayInfo{}, err2
		}

		infos := models.TrashDisposalPerDayInfo{
			TrashDisposalPerDayWaterInfo1:   pas.WaterPasturesTrashDisposal1 + sla.WaterSlaughtersTrashDisposal1,
			TrashDisposalPerDayWaterInfo2:   pas.WaterPasturesTrashDisposal2 + sla.WaterSlaughtersTrashDisposal2,
			TrashDisposalPerDayWaterInfo3:   pas.WaterPasturesTrashDisposal3 + sla.WaterSlaughtersTrashDisposal3,
			TrashDisposalPerDayWaterInfo4:   pas.WaterPasturesTrashDisposal4 + sla.WaterSlaughtersTrashDisposal4,
			TrashDisposalPerDayResidueInfo1: pas.ResiduePasturesTrashDisposal1 + sla.ResidueSlaughtersTrashDisposal1,
			TrashDisposalPerDayResidueInfo2: pas.ResiduePasturesTrashDisposal2 + sla.ResidueSlaughtersTrashDisposal2,
			TrashDisposalPerDayResidueInfo3: pas.ResiduePasturesTrashDisposal3 + sla.ResidueSlaughtersTrashDisposal3,
			TrashDisposalPerDayResidueInfo4: pas.ResiduePasturesTrashDisposal4 + sla.ResidueSlaughtersTrashDisposal4,
			TrashDisposalPerDayOdorInfo1:    pas.OdorPasturesTrashDisposal1 + sla.OdorAllSlaughtersTrashDisposal1,
			TrashDisposalPerDayOdorInfo2:    pas.OdorPasturesTrashDisposal2 + sla.OdorAllSlaughtersTrashDisposal2,
			TrashDisposalPerDayOdorInfo3:    pas.OdorPasturesTrashDisposal3 + sla.OdorAllSlaughtersTrashDisposal3,
			TrashDisposalPerDayOdorInfo4:    pas.OdorPasturesTrashDisposal4 + sla.OdorAllSlaughtersTrashDisposal4,
		}
		return infos, nil
	} else if p && !s {
		ps := query.Q.AllPasturesTrashDisposal
		pas, err3 := ps.WithContext(context.Background()).Where(ps.TimeStamp.Eq(t)).First()
		if err3 != nil {
			return models.TrashDisposalPerDayInfo{}, err3
		}

		infos := models.TrashDisposalPerDayInfo{
			TrashDisposalPerDayWaterInfo1:   pas.WaterPasturesTrashDisposal1,
			TrashDisposalPerDayWaterInfo2:   pas.WaterPasturesTrashDisposal2,
			TrashDisposalPerDayWaterInfo3:   pas.WaterPasturesTrashDisposal3,
			TrashDisposalPerDayWaterInfo4:   pas.WaterPasturesTrashDisposal4,
			TrashDisposalPerDayResidueInfo1: pas.ResiduePasturesTrashDisposal1,
			TrashDisposalPerDayResidueInfo2: pas.ResiduePasturesTrashDisposal2,
			TrashDisposalPerDayResidueInfo3: pas.ResiduePasturesTrashDisposal3,
			TrashDisposalPerDayResidueInfo4: pas.ResiduePasturesTrashDisposal4,
			TrashDisposalPerDayOdorInfo1:    pas.OdorPasturesTrashDisposal1,
			TrashDisposalPerDayOdorInfo2:    pas.OdorPasturesTrashDisposal2,
			TrashDisposalPerDayOdorInfo3:    pas.OdorPasturesTrashDisposal3,
			TrashDisposalPerDayOdorInfo4:    pas.OdorPasturesTrashDisposal4,
		}
		return infos, nil
	} else if !p && s {
		ss := query.Q.AllSlaughtersTrashDisposal
		sla, err4 := ss.WithContext(context.Background()).Where(ss.TimeStamp.Eq(t)).First()
		if err4 != nil {
			return models.TrashDisposalPerDayInfo{}, err4
		}

		infos := models.TrashDisposalPerDayInfo{
			TrashDisposalPerDayWaterInfo1:   sla.WaterSlaughtersTrashDisposal1,
			TrashDisposalPerDayWaterInfo2:   sla.WaterSlaughtersTrashDisposal2,
			TrashDisposalPerDayWaterInfo3:   sla.WaterSlaughtersTrashDisposal3,
			TrashDisposalPerDayWaterInfo4:   sla.WaterSlaughtersTrashDisposal4,
			TrashDisposalPerDayResidueInfo1: sla.ResidueSlaughtersTrashDisposal1,
			TrashDisposalPerDayResidueInfo2: sla.ResidueSlaughtersTrashDisposal2,
			TrashDisposalPerDayResidueInfo3: sla.ResidueSlaughtersTrashDisposal3,
			TrashDisposalPerDayResidueInfo4: sla.ResidueSlaughtersTrashDisposal4,
			TrashDisposalPerDayOdorInfo1:    sla.OdorAllSlaughtersTrashDisposal1,
			TrashDisposalPerDayOdorInfo2:    sla.OdorAllSlaughtersTrashDisposal2,
			TrashDisposalPerDayOdorInfo3:    sla.OdorAllSlaughtersTrashDisposal3,
			TrashDisposalPerDayOdorInfo4:    sla.OdorAllSlaughtersTrashDisposal4,
		}
		return infos, nil
	} else {
		return models.TrashDisposalPerDayInfo{}, nil
	}
	//q1 := query.Q.AllPasturesTrashDisposal
	//q2 := query.Q.AllSlaughtersTrashDisposal
	//pas, err := q1.WithContext(context.Background()).Where(q1.TimeStamp.Eq(t)).First()
	//if err != nil {
	//	return models.TrashDisposalPerDayInfo{}, err
	//}
	//sla, err := q2.WithContext(context.Background()).Where(q2.TimeStamp.Eq(t)).First()
	//if err != nil {
	//	return models.TrashDisposalPerDayInfo{}, err
	//}
	//

	//return info, nil
}

func GetUsersByCondition(condition map[string]interface{}) ([]models.UserInfo, int64, error) {
	var n, r, c, h string
	var t int

	if val, ok := condition["name"]; ok {
		n = val.(string)
	}
	if val, ok := condition["role"]; ok {
		r = val.(string)
	}
	if val, ok := condition["company"]; ok {
		c = val.(string)
	}
	if val, ok := condition["house_number"]; ok {
		h = val.(string)
	}

	u := query.FSIMSUser
	usDo := u.WithContext(context.Background()).Where(u.Name.Like("%" + n + "%")).
		Where(u.Role.Like("%" + r + "%")).
		Where(u.Company.Like("%" + c + "%")).
		Where(u.HouseNumber.Like("%" + h + "%"))

	var us []*models.FSIMSUser
	if val, ok := condition["type"]; ok {
		t = val.(int)
		results, err := usDo.Where(u.Type.Eq(t)).Find()
		if err != nil {
			return []models.UserInfo{}, 0, err
		}

		us = results
	} else {
		results, err := usDo.Find()
		if err != nil {
			return []models.UserInfo{}, 0, err
		}

		us = results
	}

	count := len(us)
	users := make([]models.UserInfo, count)
	for i, user := range us {
		users[i] = models.FsimsUserToResUser(user)
	}

	return users, int64(count), nil
}

func GetAllUsers() ([]models.UserInfo, int64, error) {
	u := query.FSIMSUser
	us, err := u.WithContext(context.Background()).Find()
	if err != nil {
		return []models.UserInfo{}, 0, err
	}

	count, err := u.WithContext(context.Background()).Count()

	users := make([]models.UserInfo, count)
	for i, user := range us {
		users[i] = models.FsimsUserToResUser(user)
	}

	return users, count, nil
}

func QueryFsimsUserPwdHash(account string) (string, error) {
	u := query.FSIMSUser
	fsimsUser, err := u.WithContext(context.Background()).Where(u.Account.Eq(account)).Select(u.PasswordHash).First()
	if err != nil {
		return "", err
	}
	return fsimsUser.PasswordHash, nil
}

func QueryFsimsUserUuidAndPwdHash(account string) (uuid, password, housenumber string, usertype int, err error) {
	u := query.FSIMSUser
	fsimsUser, err := u.WithContext(context.Background()).Where(u.Account.Eq(account)).Select(u.UUID, u.PasswordHash, u.Type, u.HouseNumber).First()
	if err != nil {
		return "", "", "", 0, err
	}
	return fsimsUser.UUID, fsimsUser.PasswordHash, fsimsUser.HouseNumber, fsimsUser.Type, nil
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

func RemoveBasicRoleForUser(uuid string) error {
	cs, err := NewCasbinService(mysql.DB)
	if err != nil {
		return err
	}
	return cs.RemoveRoleForUserWithUUID(uuid, "user")
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
	case PACK_OPERATOR_USER_TYPE:
		roleName = "packoperator"
	case TRANSPORT_OPERATOR_USER_TYPE:
		roleName = "transportoperator"
	case BUYER_OPERATOR_USER_TYPE:
		roleName = "buyeroperator"
	default:
		return errors.New("user type is not exist")
	}

	cs, err := NewCasbinService(mysql.DB)
	if err != nil {
		return err
	}

	return cs.AddRoleForUserWithUUID(uuid, roleName)
}

func RemoveRoleForUserWithType(uuid string, ttype int) error {
	return nil
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
	case PACK_OPERATOR_USER_TYPE:
		uuid = UUID_PREFIX + "-" + PACK_OPERATOR_USER_NUMBER + "-" + accountHash
	case TRANSPORT_OPERATOR_USER_TYPE:
		uuid = UUID_PREFIX + "-" + TRANSPORT_OPERATOR_USER_NUMBER + "-" + accountHash
	case BUYER_OPERATOR_USER_TYPE:
		uuid = UUID_PREFIX + "-" + BUYER_OPERATOR_USER_NUMBER + "-" + accountHash
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

func ResetFsimsPassWord(account *request.ReqAccount) error {
	p := crypto.CalculateSHA256(INIT_PASSWORD, PASSWORD_SALT)
	u := query.FSIMSUser
	uuid, err := generateUuid(account.Account, account.Type)
	if err != nil {
		glog.Errorln("Reset the user password error: %v", err)
		return errors.New("Reset the user password error")
	}
	_, err = u.WithContext(context.Background()).Where(u.UUID.Eq(uuid)).Update(u.PasswordHash, p)
	if err != nil {
		glog.Errorln("Reset the user password error: %v", err)
		return errors.New("Reset the user password error")
	}
	return nil
}

//func UpdateFsimsUser(user *request.ReqUpdateUser) error {
//	//check whether user exist
//	u := query.FSIMSUser
//	_, err := u.WithContext(context.Background()).Where(u.Account.Eq(user.Account)).First()
//	//First make a note of its uuid
//	if errors.Is(err, gorm.ErrRecordNotFound) {
//		glog.Errorln("Update a user error: %v", err)
//		return errors.New("update a user error: user not existed")
//	}
//	if err != nil {
//		glog.Errorln("Update a user error: %v", err)
//		return errors.New("update a user error")
//	}
//	//update
//	var fuu = CreateFsimUser(user)
//	_, err = u.WithContext(context.Background()).Where(u.Account.Eq(user.Account)).Updates(&fuu)
//	if err != nil {
//		return err
//	}
//
//	//casbin_rule_update
//
//	return nil
//}

func DeleteFsimUser(account *request.ReqAccount) error {
	u := query.FSIMSUser
	uuid, err := generateUuid(account.Account, account.Type)
	if err != nil {
		glog.Errorln("delete a user error: %v", err)
		return errors.New("delete a user error")
	}
	//Delete the corresponding records in the casbin table first

	res, err := u.WithContext(context.Background()).Where(u.UUID.Eq(uuid)).Delete()
	_ = res.RowsAffected

	if errors.Is(err, gorm.ErrRecordNotFound) {
		glog.Errorln("Delete a user error: %v", err)
		return errors.New("delete a user error: user not existed")
	}
	if err != nil {
		glog.Errorln("delete a user error: %v", err)
		return errors.New("delete a user error")
	}
	return nil
}

func CheckUserIsExisted(uuid string) error {
	//检查用户uuid是否已经存在
	u := query.FSIMSUser
	_, err := u.WithContext(context.Background()).Where(u.UUID.Eq(uuid)).First()
	return err
}

func CheckSlaughterTrashIsExisted(t time.Time) (bool, error) {
	time := t.Truncate(24 * time.Hour)
	fmt.Println(time)
	q1 := query.Q.AllSlaughtersTrashDisposal
	infos, err := q1.WithContext(context.Background()).Where(q1.TimeStamp.Eq(time)).Find()
	if err != nil {
		return false, err
	}
	if len(infos) == 0 {
		return false, nil
	}
	return true, nil
}

func CheckPastureTrashIsExisted(t time.Time) (bool, error) {
	q := query.Q.AllPasturesTrashDisposal
	infos, err := q.WithContext(context.Background()).Where(q.TimeStamp.Eq(t)).Find()
	if err != nil {
		return false, err
	}
	if len(infos) == 0 {
		return false, nil
	}
	return true, nil
}
