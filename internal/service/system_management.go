package service

import (
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/coldchain"
	"CN-EU-FSIMS/internal/app/models/pack"
	"CN-EU-FSIMS/internal/app/models/pasture"
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/internal/app/models/slaughter"
	"CN-EU-FSIMS/utils/crypto"
	"context"
	"errors"
	"github.com/golang/glog"
	"time"
)

const DEFAULT_PASSWORD = "FsimsOperator123456!"
const (
	PASTURE_PREFIX         = "PASP"
	SLAUGHTER_HOUSE_PREFIX = "SLAHP"
	PACKAGE_HOUSE_PREFIX   = "PACHP"
	VEHICLE_PREFIX         = "VEH"
)

func GetPastures() ([]response.ResHouse, error) {
	pastures, err := query.PastureHouse.WithContext(context.Background()).Find()
	if err != nil {
		return []response.ResHouse{}, err
	}

	res := make([]response.ResHouse, len(pastures))
	for i, p := range pastures {
		res[i] = pasture.PastureToRes(p)
	}

	return res, nil
}

func GetSlaughterHouses() ([]response.ResHouse, error) {
	shs, err := query.SlaughterHouse.WithContext(context.Background()).Find()
	if err != nil {
		return []response.ResHouse{}, err
	}

	res := make([]response.ResHouse, len(shs))
	for i, sh := range shs {
		res[i] = slaughter.SlaughterHouseToRes(sh)
	}

	return res, nil
}

func GetPackageHouses() ([]response.ResHouse, error) {
	phs, err := query.PackageHouse.WithContext(context.Background()).Find()
	if err != nil {
		return []response.ResHouse{}, err
	}

	res := make([]response.ResHouse, len(phs))
	for i, ph := range phs {
		res[i] = pack.PackageHouseToRes(ph)
	}

	return res, nil
}

func GetTransportVehicles() ([]response.ResTransportVehicle, error) {
	tvs, err := query.TransportVehicle.WithContext(context.Background()).Find()
	if err != nil {
		return []response.ResTransportVehicle{}, err
	}

	res := make([]response.ResTransportVehicle, len(tvs))
	for i, tv := range tvs {
		res[i] = coldchain.TransportVehicleToRes(tv)
	}

	return res, nil
}

func AddFsimsOperator(o *request.ReqAddOperator) (string, error) {
	uuid, err := generateUuid(o.Account, o.Type)
	if err != nil {
		return "", err
	}
	passwordHash := crypto.CalculateSHA256(DEFAULT_PASSWORD, PASSWORD_SALT)

	fsimsUser := models.FSIMSUser{
		UUID:         uuid,
		Name:         o.Name,
		Account:      o.Account,
		PasswordHash: passwordHash,
		Type:         o.Type,
		Role:         o.Role,
		Status:       1,
		Company:      o.Company,
		Phone:        o.Phone,
		HouseNumber:  o.HouseNumber,
	}
	err = query.FSIMSUser.WithContext(context.Background()).Create(&fsimsUser)
	if err != nil {
		glog.Errorln("create new user error by admin")
		return "", errors.New("create new user error")
	}

	err = AddBasicRoleForUser(uuid)
	if err != nil {
		return "", err
	}
	err = AddRoleForUserWithType(uuid, o.Type)
	if err != nil {
		return "", err
	}

	return DEFAULT_PASSWORD, nil
}

func AddPasture(p *request.ReqAddPasture) error {
	hNum := generateHouseNumber(PASTURE_PREFIX, p.Address)

	pasHouse := pasture.PastureHouse{
		HouseNumber: hNum,
		Name:        p.Name,
		Address:     p.Address,
		State:       1,
		LegalPerson: p.LegalPerson,
	}

	err := query.PastureHouse.WithContext(context.Background()).Create(&pasHouse)
	if err != nil {
		return err
	}

	return nil
}

func AddSlaughterHouse(s *request.ReqAddSlaughterHouse) error {
	hNum := generateHouseNumber(SLAUGHTER_HOUSE_PREFIX, s.Address)

	slaHouse := slaughter.SlaughterHouse{
		HouseNumber: hNum,
		Name:        s.Name,
		Address:     s.Address,
		State:       1,
		LegalPerson: s.LegalPerson,
	}
	err := query.SlaughterHouse.WithContext(context.Background()).Create(&slaHouse)
	if err != nil {
		return err
	}

	return nil
}

func AddPackageHouse(p *request.ReqAddPackHouse) error {
	hNum := generateHouseNumber(PACKAGE_HOUSE_PREFIX, p.Address)

	pacHouse := pack.PackageHouse{
		HouseNumber: hNum,
		Name:        p.Name,
		Address:     p.Address,
		State:       1,
		LegalPerson: p.LegalPerson,
	}
	err := query.PackageHouse.WithContext(context.Background()).Create(&pacHouse)
	if err != nil {
		return err
	}

	return nil
}

func AddVehicle(v *request.ReqAddVehicle) error {
	tvNum := generateVehicleNumber(v.LicenseNumber)

	tv := coldchain.TransportVehicle{
		TVNumber:      tvNum,
		LicenseNumber: v.LicenseNumber,
		Driver:        v.Driver,
		DriverPhone:   v.DriverPhone,
		State:         1,
	}

	err := query.TransportVehicle.WithContext(context.Background()).Create(&tv)
	if err != nil {
		return err
	}

	return nil
}

func generateHouseNumber(prefix, address string) string {
	t := time.Now().String()
	addressHash := crypto.CalculateSHA256(address+t, "address")

	return prefix + addressHash
}

func generateVehicleNumber(license string) string {
	t := time.Now().String()
	addressHash := crypto.CalculateSHA256(license+t, "license")

	return VEHICLE_PREFIX + addressHash
}
