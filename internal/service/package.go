package service

import (
	"CN-EU-FSIMS/internal/app/models/pack"
	"CN-EU-FSIMS/internal/app/models/query"
	"context"
)

const (
	INIT_STATE_REC_PAC = 1
)

func GetPackageInfoByNumber(num string) (pack.PackageHouseInfo, error) {
	q := query.PackageHouse
	ph, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(num)).First()
	if err != nil {
		return pack.PackageHouseInfo{}, err
	}

	phinfo := pack.ToPackageHouseInfo(ph)
	return phinfo, nil
}
