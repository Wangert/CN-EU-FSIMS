package service

import (
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/models/product"
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/internal/app/models/slaughter"
	"CN-EU-FSIMS/internal/app/models/warehouse"
	"context"
	"errors"
	"time"
)

const (
	INIT_STATE_REC_SLA          = 1
	CONFIRM_STATE_REC_SLA       = 2
	SLAUGHTER_STATE_REC_SLA     = 3
	SLAUGHTER_WAREHOUSE_REC_SLA = 4

	INIT_STATE_BATCH_SLA = 1
	END_STATE_BATCH_SLA  = 2

	INIT_SLAPRO      = 1
	WAREHOUSE_SLAPRO = 2
	SENDING_SLAPRO   = 3
	CONFIRM_SLAPRO   = 4

	INIT_SH    = 1
	SENDING_SH = 2
	CONFIRM_SH = 3

	SLICED_BEEF = 1
	BEEF_TENDON = 2
	TRIPE       = 3
	OXTAIL      = 4

	SLAUGHTER_PRODUCT_PREFIX = "SLAPRO-"
)

func UploadSlaughterStaffUniformData(r *request.ReqUploadStaffUniformData) error {
	ok, err := SlaughterHouseIsExisted(r.HouseNumber)
	if !ok {
		return err
	}

	timeRecordAt := time.Unix(r.TimestampRecordAt, 0).UTC().Local()

	staffUni := slaughter.StaUni{
		HouseNumber:  r.HouseNumber,
		TimeRecordAt: timeRecordAt,
		StaUni1:      r.StaUni1,
		StaUni2:      r.StaUni2,
		StaUni3:      r.StaUni3,
		StaUni4:      r.StaUni4,
		StaUni5:      r.StaUni5,
		StaUni6:      r.StaUni6,
		StaUni7:      r.StaUni7,
		StaUni8:      r.StaUni8,
		StaUni9:      r.StaUni9,
		StaUni10:     r.StaUni10,
		StaUni11:     r.StaUni11,
		StaUni12:     r.StaUni12,
		StaUni13:     r.StaUni13,
		StaUni14:     r.StaUni14,
	}

	err = query.Q.StaUni.WithContext(context.Background()).Create(&staffUni)
	if err != nil {
		return err
	}

	return nil
}

func QuerySlaughterStaffUniformData(r *request.ReqSlaughterSensorData) ([]slaughter.StaUniInfo, int64, error) {
	ok, err := SlaughterHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.StaUni
	results, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(results)
	infos := make([]slaughter.StaUniInfo, count)
	for i, result := range results {
		infos[i] = slaughter.ToStaUniInfo(result)
	}

	return infos, int64(count), nil
}

func UploadSlaughterLightRecord(r *request.ReqUploadSlaughterLightRecord) error {
	ok, err := SlaughterHouseIsExisted(r.HouseNumber)
	if !ok {
		return err
	}

	timeRecordAt := time.Unix(r.TimestampRecordAt, 0).UTC().Local()

	light := slaughter.SlaughterLightRecord{
		HouseNumber:   r.HouseNumber,
		TimeRecordAt:  timeRecordAt,
		SlaEnvLigRec1: r.SlaEnvLigRec1,
		SlaEnvLigRec2: r.SlaEnvLigRec2,
		SlaEnvLigRec3: r.SlaEnvLigRec3,
		SlaEnvLigRec4: r.SlaEnvLigRec4,
	}

	err = query.Q.SlaughterLightRecord.WithContext(context.Background()).Create(&light)
	if err != nil {
		return err
	}

	return nil
}

func QuerySlaughterLightRecord(r *request.ReqSlaughterSensorData) ([]slaughter.SlaughterLightRecordInfo, int64, error) {
	ok, err := SlaughterHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.SlaughterLightRecord
	results, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(results)
	infos := make([]slaughter.SlaughterLightRecordInfo, count)
	for i, result := range results {
		infos[i] = slaughter.ToSlaEnvLigRecInfo(result)
	}

	return infos, int64(count), nil
}

func QuerySlaughterWaterQualityData(r *request.ReqSlaughterSensorData) ([]slaughter.SlaughterWaterQualityMonInfo, int64, error) {
	ok, err := SlaughterHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.SlaughterWaterQualityMon
	results, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).
		Preload(q.SlaughterWaterMicroIndex).Preload(q.OapGciSla).
		Preload(q.MicroIndexWaterMonSla).Preload(q.ToxinIndexSla).
		Preload(q.ToxinIndexSla.SlaughterWaterToxinIndex).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(results)
	infos := make([]slaughter.SlaughterWaterQualityMonInfo, count)
	for i, result := range results {
		infos[i] = slaughter.ToSlaughterWaterQualityMonInfo(result)
	}

	return infos, int64(count), nil
}

func UploadSlaughterWaterQualityData(r *request.ReqUploadSlaughterWaterQualityData) error {
	ok, err := SlaughterHouseIsExisted(r.HouseNumber)
	if !ok {
		return err
	}

	timeRecordAt := time.Unix(r.TimestampRecordAt, 0).UTC().Local()

	waterQuality := slaughter.SlaughterWaterQualityMon{
		HouseNumber:  r.HouseNumber,
		TimeRecordAt: timeRecordAt,
		SlaughterWaterMicroIndex: slaughter.SlaughterWaterMicroIndex{
			SlaughterWaterQualityMonID: nil,
			SlaughterWaterMicroIndex1:  r.SlaughterWaterMicroIndex.SlaughterWaterMicroIndex1,
			SlaughterWaterMicroIndex2:  r.SlaughterWaterMicroIndex.SlaughterWaterMicroIndex2,
			SlaughterWaterMicroIndex3:  r.SlaughterWaterMicroIndex.SlaughterWaterMicroIndex3,
		},
		OapGciSla: slaughter.OapGciSla{
			SlaughterWaterQualityMonID: nil,
			OapGciSla1:                 r.OapGciSla.OapGciSla1,
			OapGciSla2:                 r.OapGciSla.OapGciSla2,
			OapGciSla3:                 r.OapGciSla.OapGciSla3,
			OapGciSla4:                 r.OapGciSla.OapGciSla4,
			OapGciSla5:                 r.OapGciSla.OapGciSla5,
			OapGciSla6:                 r.OapGciSla.OapGciSla6,
			OapGciSla7:                 r.OapGciSla.OapGciSla7,
			OapGciSla8:                 r.OapGciSla.OapGciSla8,
			OapGciSla9:                 r.OapGciSla.OapGciSla9,
			OapGciSla10:                r.OapGciSla.OapGciSla10,
			OapGciSla11:                r.OapGciSla.OapGciSla11,
			OapGciSla12:                r.OapGciSla.OapGciSla12,
			OapGciSla13:                r.OapGciSla.OapGciSla13,
			OapGciSla14:                r.OapGciSla.OapGciSla14,
			OapGciSla15:                r.OapGciSla.OapGciSla15,
			OapGciSla16:                r.OapGciSla.OapGciSla16,
			OapGciSla17:                r.OapGciSla.OapGciSla17,
			OapGciSla18:                r.OapGciSla.OapGciSla18,
			OapGciSla19:                r.OapGciSla.OapGciSla19,
			OapGciSla20:                r.OapGciSla.OapGciSla20,
			OapGciSla21:                r.OapGciSla.OapGciSla21,
		},
		ToxinIndexSla: slaughter.SlaughterToxinIndex{
			SlaughterWaterQualityMonID: nil,
			ToxinIndexSla1:             r.SlaughterToxinIndex.ToxinIndexSla1,
			ToxinIndexSla2:             r.SlaughterToxinIndex.ToxinIndexSla2,
			ToxinIndexSla3:             r.SlaughterToxinIndex.ToxinIndexSla3,
			ToxinIndexSla4:             r.SlaughterToxinIndex.ToxinIndexSla4,
			ToxinIndexSla5:             r.SlaughterToxinIndex.ToxinIndexSla5,
			ToxinIndexSla6:             r.SlaughterToxinIndex.ToxinIndexSla6,
			ToxinIndexSla7:             r.SlaughterToxinIndex.ToxinIndexSla7,
			ToxinIndexSla8:             r.SlaughterToxinIndex.ToxinIndexSla8,
			ToxinIndexSla9:             r.SlaughterToxinIndex.ToxinIndexSla9,
			ToxinIndexSla10:            r.SlaughterToxinIndex.ToxinIndexSla10,
			ToxinIndexSla11:            r.SlaughterToxinIndex.ToxinIndexSla11,
			ToxinIndexSla12:            r.SlaughterToxinIndex.ToxinIndexSla12,
			ToxinIndexSla13:            r.SlaughterToxinIndex.ToxinIndexSla13,
			ToxinIndexSla14:            r.SlaughterToxinIndex.ToxinIndexSla14,
			ToxinIndexSla15:            r.SlaughterToxinIndex.ToxinIndexSla15,
			ToxinIndexSla16:            r.SlaughterToxinIndex.ToxinIndexSla16,
			ToxinIndexSla17:            r.SlaughterToxinIndex.ToxinIndexSla17,
			ToxinIndexSla18:            r.SlaughterToxinIndex.ToxinIndexSla18,
			ToxinIndexSla19:            r.SlaughterToxinIndex.ToxinIndexSla19,
			ToxinIndexSla20:            r.SlaughterToxinIndex.ToxinIndexSla20,
			ToxinIndexSla21:            r.SlaughterToxinIndex.ToxinIndexSla21,
			ToxinIndexSla22:            r.SlaughterToxinIndex.ToxinIndexSla22,
			ToxinIndexSla23:            r.SlaughterToxinIndex.ToxinIndexSla23,
			ToxinIndexSla24:            r.SlaughterToxinIndex.ToxinIndexSla24,
			ToxinIndexSla25:            r.SlaughterToxinIndex.ToxinIndexSla25,
			ToxinIndexSla26:            r.SlaughterToxinIndex.ToxinIndexSla26,
			ToxinIndexSla27:            r.SlaughterToxinIndex.ToxinIndexSla27,
			ToxinIndexSla28:            r.SlaughterToxinIndex.ToxinIndexSla28,
			ToxinIndexSla29:            r.SlaughterToxinIndex.ToxinIndexSla29,
			ToxinIndexSla30:            r.SlaughterToxinIndex.ToxinIndexSla30,
			ToxinIndexSla31:            r.SlaughterToxinIndex.ToxinIndexSla31,
			ToxinIndexSla32:            r.SlaughterToxinIndex.ToxinIndexSla32,
			ToxinIndexSla33:            r.SlaughterToxinIndex.ToxinIndexSla33,
			ToxinIndexSla34:            r.SlaughterToxinIndex.ToxinIndexSla34,
			ToxinIndexSla35:            r.SlaughterToxinIndex.ToxinIndexSla35,
			ToxinIndexSla36:            r.SlaughterToxinIndex.ToxinIndexSla36,
			ToxinIndexSla37:            r.SlaughterToxinIndex.ToxinIndexSla37,
			ToxinIndexSla38:            r.SlaughterToxinIndex.ToxinIndexSla38,
			ToxinIndexSla39:            r.SlaughterToxinIndex.ToxinIndexSla39,
			ToxinIndexSla40:            r.SlaughterToxinIndex.ToxinIndexSla40,
			ToxinIndexSla41:            r.SlaughterToxinIndex.ToxinIndexSla41,
			ToxinIndexSla42:            r.SlaughterToxinIndex.ToxinIndexSla42,
			ToxinIndexSla43:            r.SlaughterToxinIndex.ToxinIndexSla43,
			ToxinIndexSla44:            r.SlaughterToxinIndex.ToxinIndexSla44,
			ToxinIndexSla45:            r.SlaughterToxinIndex.ToxinIndexSla45,
			ToxinIndexSla46:            r.SlaughterToxinIndex.ToxinIndexSla46,
			ToxinIndexSla47:            r.SlaughterToxinIndex.ToxinIndexSla47,
			SlaughterWaterToxinIndex: slaughter.SlaughterWaterToxinIndex{
				SlaughterToxinIndexID:      nil,
				SlaughterWaterToxinIndex1:  r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex1,
				SlaughterWaterToxinIndex2:  r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex2,
				SlaughterWaterToxinIndex3:  r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex3,
				SlaughterWaterToxinIndex4:  r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex4,
				SlaughterWaterToxinIndex5:  r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex5,
				SlaughterWaterToxinIndex6:  r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex6,
				SlaughterWaterToxinIndex7:  r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex7,
				SlaughterWaterToxinIndex8:  r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex8,
				SlaughterWaterToxinIndex9:  r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex9,
				SlaughterWaterToxinIndex10: r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex10,
				SlaughterWaterToxinIndex11: r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex11,
				SlaughterWaterToxinIndex12: r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex12,
				SlaughterWaterToxinIndex13: r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex13,
				SlaughterWaterToxinIndex14: r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex14,
				SlaughterWaterToxinIndex15: r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex15,
				SlaughterWaterToxinIndex16: r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex16,
				SlaughterWaterToxinIndex17: r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex17,
				SlaughterWaterToxinIndex18: r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex18,
				SlaughterWaterToxinIndex19: r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex19,
				SlaughterWaterToxinIndex20: r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex20,
				SlaughterWaterToxinIndex21: r.SlaughterToxinIndex.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex21,
			},
		},
	}

	err = query.Q.SlaughterWaterQualityMon.WithContext(context.Background()).Create(&waterQuality)
	if err != nil {
		return err
	}

	return nil
}

func QueryPreCoolShopData(r *request.ReqSlaughterSensorData) ([]slaughter.PreCoolShopInfo, int64, error) {
	ok, err := SlaughterHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.PreCoolShop
	shops, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(shops)
	shopInfos := make([]slaughter.PreCoolShopInfo, count)
	for i, shop := range shops {
		shopInfos[i] = slaughter.ToPreCoolShopInfo(shop)
	}

	return shopInfos, int64(count), nil
}

func QuerySlaughterShopData(r *request.ReqSlaughterSensorData) ([]slaughter.SlaShopInfo, int64, error) {
	ok, err := SlaughterHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.SlaShop
	shops, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(shops)
	shopInfos := make([]slaughter.SlaShopInfo, count)
	for i, shop := range shops {
		shopInfos[i] = slaughter.ToSlaShopInfo(shop)
	}

	return shopInfos, int64(count), nil
}

func QueryDivisionShopData(r *request.ReqSlaughterSensorData) ([]slaughter.DivShopInfo, int64, error) {
	ok, err := SlaughterHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.DivShop
	shops, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(shops)
	shopInfos := make([]slaughter.DivShopInfo, count)
	for i, shop := range shops {
		shopInfos[i] = slaughter.ToDivShopInfo(shop)
	}

	return shopInfos, int64(count), nil
}

func QueryAcidShopData(r *request.ReqSlaughterSensorData) ([]slaughter.AcidShopInfo, int64, error) {
	ok, err := SlaughterHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.AcidShop
	shops, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(shops)
	shopInfos := make([]slaughter.AcidShopInfo, count)
	for i, shop := range shops {
		shopInfos[i] = slaughter.ToAcidShopInfo(shop)
	}

	return shopInfos, int64(count), nil
}

func QueryFrozenShopData(r *request.ReqSlaughterSensorData) ([]slaughter.FroShopInfo, int64, error) {
	ok, err := SlaughterHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.FroShop
	shops, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(shops)
	shopInfos := make([]slaughter.FroShopInfo, count)
	for i, shop := range shops {
		shopInfos[i] = slaughter.ToFroShopInfo(shop)
	}

	return shopInfos, int64(count), nil
}

func UploadPreCoolShopData(r *request.ReqUploadPreCoolShopData) error {
	ok, err := SlaughterHouseIsExisted(r.HouseNumber)
	if !ok {
		return err
	}

	shop := slaughter.PreCoolShop{
		HouseNumber:  r.HouseNumber,
		TimeRecordAt: time.Unix(r.TimestampRecordAt, 0).UTC().Local(),
		PreCoolShop1: r.PreCoolShop1,
		PreCoolShop2: r.PreCoolShop2,
		PreCoolShop3: r.PreCoolShop3,
	}

	err = query.Q.PreCoolShop.WithContext(context.Background()).Create(&shop)
	if err != nil {
		return err
	}

	return nil
}

func UploadSlaughterShopData(r *request.ReqUploadSlaughterShopData) error {
	ok, err := SlaughterHouseIsExisted(r.HouseNumber)
	if !ok {
		return err
	}

	shop := slaughter.SlaShop{
		HouseNumber:  r.HouseNumber,
		TimeRecordAt: time.Unix(r.TimestampRecordAt, 0).UTC().Local(),
		SlaShop1:     r.SlaShop1,
		SlaShop2:     r.SlaShop2,
		SlaShop3:     r.SlaShop3,
		SlaShop4:     r.SlaShop4,
		SlaShop5:     r.SlaShop5,
		SlaShop6:     r.SlaShop6,
		SlaShop7:     r.SlaShop7,
		SlaShop8:     r.SlaShop8,
		SlaShop9:     r.SlaShop9,
		SlaShop10:    r.SlaShop10,
		SlaShop11:    r.SlaShop11,
		SlaShop12:    r.SlaShop12,
		SlaShop13:    r.SlaShop13,
		SlaShop14:    r.SlaShop14,
	}

	err = query.Q.SlaShop.WithContext(context.Background()).Create(&shop)
	return nil
}

func UploadDivisionShopData(r *request.ReqUploadDivisionShopData) error {
	ok, err := SlaughterHouseIsExisted(r.HouseNumber)
	if !ok {
		return err
	}

	shop := slaughter.DivShop{
		HouseNumber:  r.HouseNumber,
		TimeRecordAt: time.Unix(r.TimestampRecordAt, 0).UTC().Local(),
		DivShop1:     r.DivShop1,
		DivShop2:     r.DivShop2,
		DivShop3:     r.DivShop3,
		DivShop4:     r.DivShop4,
		DivShop5:     r.DivShop5,
		DivShop6:     r.DivShop6,
		DivShop7:     r.DivShop7,
		DivShop8:     r.DivShop8,
		DivShop9:     r.DivShop9,
		DivShop10:    r.DivShop10,
		DivShop11:    r.DivShop11,
		DivShop12:    r.DivShop12,
		DivShop13:    r.DivShop13,
		DivShop14:    r.DivShop14,
	}

	err = query.Q.DivShop.WithContext(context.Background()).Create(&shop)
	if err != nil {
		return err
	}
	return nil
}

func UploadAcidShopData(r *request.ReqUploadAcidShopData) error {
	ok, err := SlaughterHouseIsExisted(r.HouseNumber)
	if !ok {
		return err
	}

	shop := slaughter.AcidShop{
		HouseNumber:  r.HouseNumber,
		TimeRecordAt: time.Unix(r.TimestampRecordAt, 0).UTC().Local(),
		AcidShop1:    r.AcidShop1,
		AcidShop2:    r.AcidShop2,
		AcidShop3:    r.AcidShop3,
		AcidShop4:    r.AcidShop4,
		AcidShop5:    r.AcidShop5,
		AcidShop6:    r.AcidShop6,
		AcidShop7:    r.AcidShop7,
		AcidShop8:    r.AcidShop8,
		AcidShop9:    r.AcidShop9,
		AcidShop10:   r.AcidShop10,
		AcidShop11:   r.AcidShop11,
		AcidShop12:   r.AcidShop12,
		AcidShop13:   r.AcidShop13,
		AcidShop14:   r.AcidShop14,
	}

	err = query.Q.AcidShop.WithContext(context.Background()).Create(&shop)
	if err != nil {
		return err
	}

	return nil
}

func UploadFrozenShopData(r *request.ReqUploadFrozenShopData) error {
	ok, err := SlaughterHouseIsExisted(r.HouseNumber)
	if !ok {
		return err
	}

	shop := slaughter.FroShop{
		HouseNumber:  r.HouseNumber,
		TimeRecordAt: time.Unix(r.TimestampRecordAt, 0).UTC().Local(),
		FroShop1:     r.FroShop1,
		FroShop2:     r.FroShop2,
		FroShop3:     r.FroShop3,
		FroShop4:     r.FroShop4,
		FroShop5:     r.FroShop5,
		FroShop6:     r.FroShop6,
		FroShop7:     r.FroShop7,
		FroShop8:     r.FroShop8,
		FroShop9:     r.FroShop9,
		FroShop10:    r.FroShop10,
		FroShop11:    r.FroShop11,
		FroShop12:    r.FroShop12,
		FroShop13:    r.FroShop13,
		FroShop14:    r.FroShop14,
	}

	err = query.Q.FroShop.WithContext(context.Background()).Create(&shop)
	if err != nil {
		return err
	}

	return nil
}

func SlaughterHouseIsExisted(number string) (bool, error) {
	q := query.Q.SlaughterHouse
	shs, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(number)).Find()
	if err != nil {
		return false, err
	}

	if len(shs) == 0 {
		return false, errors.New("slaughter house is not existed")
	}

	return true, nil
}

func SendToPackage(r *request.ReqSendToPackage) error {
	var err error
	tx := query.Q.Begin()
	defer func() {
		if recover() != nil {
			_ = tx.Rollback()
		}
	}()

	pid, err := GetPidBySlaughterProductNumber(r.ProductNumber)
	if err != nil {
		return err
	}

	shinfo, err := GetSlaughterHouseInfoByProductNumber(r.ProductNumber)
	if err != nil {
		return err
	}

	ph, err := GetPackageInfoByNumber(r.PackageHouseNumber)
	if err != nil {
		return err
	}

	// 更新slaughter product 状态
	_, err = tx.Cow.WithContext(context.Background()).
		Where(tx.Cow.Number.Eq(r.ProductNumber)).
		Updates(map[string]interface{}{"state": SENDING_SLAPRO})
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// 更新slaughter warehouse状态
	_, err = tx.SlaughterWarehouse.WithContext(context.Background()).
		Where(tx.SlaughterWarehouse.ProductNumber.Eq(r.ProductNumber)).
		Updates(map[string]interface{}{"state": SENDING_SH, "destination": ph.Name, "out_timestamp": time.Now()})
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// 新增package receive record
	prr := warehouse.PackageReceiveRecord{
		ProductNumber: r.ProductNumber,
		PID:           pid,
		SourceNumber:  shinfo.HouseNumber,
		SourceName:    shinfo.Name,
		Operator:      r.Operator,
		State:         INIT_STATE_REC_PAC,
		ReceiveTime:   time.Now(),
		ConfirmTime:   nil,
		PacNumber:     r.PackageHouseNumber,
	}
	err = tx.PackageReceiveRecord.WithContext(context.Background()).Create(&prr)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func GetSlaughterHouseInfoByProductNumber(num string) (slaughter.SlaughterHouseInfo, error) {
	q1 := query.SlaughterWarehouse
	sw, err := q1.WithContext(context.Background()).Where(q1.ProductNumber.Eq(num)).First()
	if err != nil {
		return slaughter.SlaughterHouseInfo{}, err
	}

	q2 := query.SlaughterHouse
	sh, err := q2.WithContext(context.Background()).Where(q2.HouseNumber.Eq(sw.HouseNumber)).First()
	if err != nil {
		return slaughter.SlaughterHouseInfo{}, err
	}

	shinfo := slaughter.ToSlaughterHouseInfo(sh)

	return shinfo, nil
}

func GetPidBySlaughterProductNumber(num string) (string, error) {
	q := query.SlaughterWarehouse
	sw, err := q.WithContext(context.Background()).Where(q.ProductNumber.Eq(num)).First()
	if err != nil {
		return "", err
	}

	return sw.PID, nil
}

func EndSlaughter(r *request.ReqEndSlaughter) (string, []string, error) {
	var err error
	tx := query.Q.Begin()

	defer func() {
		if recover() != nil {
			_ = tx.Rollback()
		}
	}()

	endTime := time.Now()

	// 读取PID和cow number
	pid, cowNumber, err := GetPidAndCowNumBySlaughterBatchNumber(r.BatchNumber)
	if err != nil {
		return "", nil, err
	}

	// 获取slaughter products
	products, err := GetSlaughterProductsByBatchNumber(r.BatchNumber)
	if err != nil {
		return "", nil, err
	}

	// 更新receive record状态
	_, err = tx.SlaughterReceiveRecord.WithContext(context.Background()).
		Where(tx.SlaughterReceiveRecord.CowNumber.Eq(cowNumber)).
		Updates(map[string]interface{}{"state": SLAUGHTER_WAREHOUSE_REC_SLA})
	if err != nil {
		_ = tx.Rollback()
		return "", nil, err
	}

	// 更新SlaughterBatch状态
	_, err = tx.SlaughterBatch.WithContext(context.Background()).
		Where(tx.SlaughterBatch.BatchNumber.Eq(r.BatchNumber)).
		Updates(map[string]interface{}{"state": END_STATE_BATCH_SLA, "end_time": &endTime})
	if err != nil {
		_ = tx.Rollback()
		return "", nil, err
	}

	// 写入仓库
	count := len(products)
	sws := make([]*warehouse.SlaughterWarehouse, count)
	productNums := make([]string, count)
	for i, product := range products {
		sw := warehouse.SlaughterWarehouse{
			ProductNumber: product.Number,
			PID:           pid,
			Type:          product.TypeName,
			State:         INIT_SH,
			InOperator:    r.Worker,
			OutOperator:   "",
			Destination:   "",
			InTimestamp:   time.Now(),
			OutTimestamp:  nil,
			HouseNumber:   r.HouseNumber,
		}
		sws[i] = &sw
		productNums[i] = product.Number
	}
	err = tx.SlaughterWarehouse.WithContext(context.Background()).Create(sws...)
	if err != nil {
		_ = tx.Rollback()
		return "", nil, err
	}

	// 更新product状态
	_, err = tx.SlaughterProduct.WithContext(context.Background()).
		Where(tx.SlaughterProduct.BatchNumber.Eq(r.BatchNumber)).
		Updates(map[string]interface{}{"state": WAREHOUSE_SLAPRO})
	if err != nil {
		_ = tx.Rollback()
		return "", nil, err
	}

	// 提交Procedure
	data := slaughter.SlaughterProcedureData{
		EnvirTemperature:      r.EnvirTemperature,
		EnvirLighting:         r.EnvirLighting,
		ShockVoltage:          r.ShockVoltage,
		BleedingTime:          r.BleedingTime,
		EleShockTime:          r.EleShockTime,
		KnifeDisinfectionTime: r.KnifeDisinfectionTime,
	}

	checkcode, err := BasicCommitProcedureWithTx(tx, pid, &endTime, data)
	if err != nil {
		_ = tx.Rollback()
		return "", nil, err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return "", nil, err
	}

	return checkcode, productNums, nil
}

func GetSlaughterProductsByBatchNumber(num string) ([]*product.SlaughterProduct, error) {
	q := query.SlaughterProduct
	sps, err := q.WithContext(context.Background()).Where(q.BatchNumber.Eq(num)).Find()
	if err != nil {
		return nil, err
	}

	return sps, nil
}

func NewSlaughterProduct(r *request.ReqNewSlaughterProduct) (string, error) {
	number := SLAUGHTER_PRODUCT_PREFIX + GenerateNumber(r)
	p := product.SlaughterProduct{
		Number:      number,
		Type:        r.Type,
		TypeName:    typeName(r.Type),
		Weight:      r.Weight,
		State:       INIT_SLAPRO,
		HouseNumber: r.HouseNumber,
		BatchNumber: r.BatchNumber,
	}

	q := query.SlaughterProduct
	err := q.WithContext(context.Background()).Create(&p)
	if err != nil {
		return "", err
	}

	return number, nil
}

func typeName(ttype int) string {
	switch ttype {
	case SLICED_BEEF:
		return "牛肉片"
	case BEEF_TENDON:
		return "牛筋"
	case TRIPE:
		return "牛肚"
	case OXTAIL:
		return "牛尾巴"
	default:
		return "未知"
	}
}

func GetPidAndCowNumBySlaughterBatchNumber(num string) (string, string, error) {
	q := query.SlaughterBatch
	fb, err := q.WithContext(context.Background()).Where(q.BatchNumber.Eq(num)).First()
	if err != nil {
		return "", "", err
	}

	return fb.PID, fb.CowNumber, nil
}

func NewSlaughterBatch(r *request.ReqNewSlaughterBatch) (string, error) {
	var err error
	tx := query.Q.Begin()

	defer func() {
		if recover() != nil {
			_ = tx.Rollback()
		}
	}()

	startTime := time.Now()

	bNum := BATCH_NUMBER_PREFIX + GenerateNumber(r)

	pp := NewProcedureParams{
		Type:        SLAUGHTER_TYPE,
		Operator:    r.Worker,
		PrePID:      r.PrePID,
		BatchNumber: bNum,
	}
	procedure, err := NewProcedure(&pp, startTime)

	sb := slaughter.SlaughterBatch{
		BatchNumber: bNum,
		HouseNumber: r.HouseNumber,
		State:       INIT_STATE_BATCH_SLA,
		PID:         procedure.PID,
		Worker:      r.Worker,
		CowNumber:   r.CowNumber,
		StartTime:   &startTime,
	}

	err = tx.Procedure.WithContext(context.Background()).Create(&procedure)
	if err != nil {
		_ = tx.Rollback()
		return "", err
	}
	err = tx.SlaughterBatch.WithContext(context.Background()).Create(&sb)
	if err != nil {
		_ = tx.Rollback()
		return "", err
	}

	// 更新slaughter receive record
	_, err = tx.SlaughterReceiveRecord.WithContext(context.Background()).
		Where(tx.SlaughterReceiveRecord.CowNumber.Eq(r.CowNumber)).
		Updates(map[string]interface{}{"state": SLAUGHTER_STATE_REC_SLA})
	if err != nil {
		_ = tx.Rollback()
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return "", err
	}

	return bNum, nil
}

func ConfirmCowFromPasture(num string) error {
	q := query.SlaughterReceiveRecord
	_, err := q.WithContext(context.Background()).
		Where(q.CowNumber.Eq(num)).
		Updates(map[string]interface{}{"state": CONFIRM_STATE_REC_SLA, "confirm_time": time.Now()})
	if err != nil {
		return err
	}
	return nil
}

func GetSlaughterInfoByNumber(num string) (slaughter.SlaughterHouseInfo, error) {
	q := query.SlaughterHouse
	sh, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(num)).First()
	if err != nil {
		return slaughter.SlaughterHouseInfo{}, err
	}

	shinfo := slaughter.ToSlaughterHouseInfo(sh)
	return shinfo, nil
}

func GetSlaughterReceiveRecords(houseNum string) ([]warehouse.SlaughterReceiveRecordInfo, int64, error) {
	q := query.SlaughterReceiveRecord
	srs, err := q.WithContext(context.Background()).Where(q.SlaughterNumber.Eq(houseNum)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(srs)
	records := make([]warehouse.SlaughterReceiveRecordInfo, count)
	for i, fb := range srs {
		records[i] = warehouse.ToSlaughterReceiveRecordInfo(fb)
	}

	return records, int64(count), nil
}

func GetSlaughterBatches(houseNum string) ([]slaughter.SlaughterBatchInfo, int64, error) {
	q := query.SlaughterBatch
	sbs, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(houseNum)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(sbs)
	records := make([]slaughter.SlaughterBatchInfo, count)
	for i, sb := range sbs {
		records[i] = slaughter.ToSlaughterBatchInfo(sb)
	}

	return records, int64(count), nil
}

func GetSlaughterWarehouseRecords(houseNum string) ([]warehouse.SlaughterWarehouseInfo, int64, error) {
	q := query.SlaughterWarehouse
	sws, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(houseNum)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(sws)
	records := make([]warehouse.SlaughterWarehouseInfo, count)
	for i, sw := range sws {
		records[i] = warehouse.ToSlaughterWarehouseInfo(sw)
	}

	return records, int64(count), nil
}
