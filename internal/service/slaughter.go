package service

import (
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/internal/app/models/slaughter"
	"context"
)

func GetSlaughterInfoByNumber(num string) (slaughter.SlaughterHouseInfo, error) {
	q := query.SlaughterHouse
	sh, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(num)).First()
	if err != nil {
		return slaughter.SlaughterHouseInfo{}, err
	}

	shinfo := slaughter.ToSlaughterHouseInfo(*sh)
	return shinfo, nil
}
