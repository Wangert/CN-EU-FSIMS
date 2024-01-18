package service

import (
	"CN-EU-FSIMS/fabric"
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/models/pasture"
	"CN-EU-FSIMS/internal/app/models/product"
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/internal/app/models/warehouse"
	"context"
	"errors"
	"time"

	"github.com/golang/glog"
)

const (
	COW_NUMBER_PREFIX   = "COW-"
	BATCH_NUMBER_PREFIX = "BATCH-"

	// cow state
	INIT_STATE_COW          = 1
	FEEDING_STATE_COW       = 2
	WAREHOUSE_STATE_COW     = 3
	SENDING_STATE_COW       = 4
	CONFIRM_STATE_COW       = 5
	SLAUGHTER_END_STATE_COW = 6

	// pasture batch state
	INIT_STATE_BATCH_PAS = 1
	END_STATE_BATCH_PAS  = 2

	// pasture house state
	INIT_STATE_PH    = 1
	SENDING_STATE_PH = 2
	CONFIRM_STATE_PH = 3
)

func QueryPastureBuffer(r *request.ReqPastureSensorData) ([]pasture.PastureBufferInfo, int64, error) {
	ok, err := PastureHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.PastureBuffer
	results, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(results)
	infos := make([]pasture.PastureBufferInfo, count)
	for i, result := range results {
		infos[i] = pasture.ToPastureBufferInfo(result)
	}
	return infos, int64(count), nil
}

func QueryPastureCowHouse(r *request.ReqPastureSensorData) ([]pasture.CowHouseInfo, int64, error) {
	ok, err := PastureHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.CowHouse
	results, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(results)
	infos := make([]pasture.CowHouseInfo, count)
	for i, result := range results {
		infos[i] = pasture.ToCowHouseInfo(result)
	}
	return infos, int64(count), nil
}

func QueryPastureArea(r *request.ReqPastureSensorData) ([]pasture.PastureAreaInfo, int64, error) {
	ok, err := PastureHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.PastureArea
	results, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(results)
	infos := make([]pasture.PastureAreaInfo, count)
	for i, result := range results {
		infos[i] = pasture.ToPastureAreaInfo(result)
	}
	return infos, int64(count), nil
}

func QueryPastureBasicEnvironment(r *request.ReqPastureSensorData) ([]pasture.PastureBasicEnvironmentInfo, int64, error) {
	ok, err := PastureHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.PastureBasicEnvironment
	results, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(results)
	infos := make([]pasture.PastureBasicEnvironmentInfo, count)
	for i, result := range results {
		infos[i] = pasture.ToPastureBasicEnvironmentInfo(result)
	}
	return infos, int64(count), nil
}

func QueryPasturePaddingRequire(r *request.ReqPastureSensorData) ([]pasture.PasturePaddingRequireInfo, int64, error) {
	ok, err := PastureHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.PasturePaddingRequire
	results, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(results)
	infos := make([]pasture.PasturePaddingRequireInfo, count)
	for i, result := range results {
		infos[i] = pasture.ToPasturePaddingRequireInfo(result)
	}
	return infos, int64(count), nil
}

func QueryPastureWastedWaterIndex(r *request.ReqPastureSensorData) ([]pasture.PastureWastedWaterIndexInfo, int64, error) {
	ok, err := PastureHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.PastureWastedWaterIndex
	results, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(results)
	infos := make([]pasture.PastureWastedWaterIndexInfo, count)
	for i, result := range results {
		infos[i] = pasture.ToPastureWastedWaterIndexInfo(result)
	}
	return infos, int64(count), nil
}

func QueryPastureDisinfectionRecord(r *request.ReqPastureSensorData) ([]pasture.PastureDisinfectionRecordInfo, int64, error) {
	ok, err := PastureHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.PastureDisinfectionRecord
	results, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(results)
	infos := make([]pasture.PastureDisinfectionRecordInfo, count)
	for i, result := range results {
		infos[i] = pasture.ToPastureDisinfectionRecordInfo(result)
	}
	return infos, int64(count), nil
}

func QueryPastureFeedHeavyMetal(r *request.ReqPastureSensorData) ([]pasture.PastureFeedHeavyMetalInfo, int64, error) {
	ok, err := PastureHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.PastureFeedHeavyMetal
	results, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).Preload(q.PastureFeedAs).
		Preload(q.PastureFeedPb).Preload(q.PastureFeedCr).Preload(q.PastureFeedCd).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(results)
	infos := make([]pasture.PastureFeedHeavyMetalInfo, count)
	for i, result := range results {
		infos[i] = pasture.ToPastureFeedHeavyMetalInfo(result)
	}
	return infos, int64(count), nil
}

func QueryPastureFeedMycotoxins(r *request.ReqPastureSensorData) ([]pasture.PastureFeedMycotoxinsInfo, int64, error) {
	ok, err := PastureHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.PastureFeedMycotoxins
	results, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).Preload(q.Don).Preload(q.Afb1).Preload(q.T2VomZea).
		Preload(q.T2toxin).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(results)
	infos := make([]pasture.PastureFeedMycotoxinsInfo, count)
	for i, result := range results {
		infos[i] = pasture.ToPastureFeedMycotoxinsInfo(result)
	}
	return infos, int64(count), nil
}

func QueryPastureWaterRecord(r *request.ReqPastureSensorData) ([]pasture.PastureWaterRecordInfo, int64, error) {
	ok, err := PastureHouseIsExisted(r.HouseNumber)
	if !ok {
		return nil, 0, err
	}

	startTime := time.Unix(r.StartTimestamp, 0).UTC().Local()
	endTime := time.Unix(r.EndTimestamp, 0).UTC().Local()

	q := query.Q.PastureWaterRecord
	results, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(r.HouseNumber)).
		Where(q.TimeRecordAt.Between(startTime, endTime)).Preload(q.OapGci).Preload(q.ToxIndex).Preload(q.MicroIndex).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(results)
	infos := make([]pasture.PastureWaterRecordInfo, count)
	for i, result := range results {
		infos[i] = pasture.ToPastureWaterRecordInfo(result)
	}
	return infos, int64(count), nil
}

func UploadPastureBuffer(r *request.ReqAddPastureBuffer) error {
	var err error
	pb := pasture.PastureBuffer{
		TimeRecordAt: time.Unix(r.TimeStamp, 0),
		HouseNumber:  r.HouseNumber,
		Buffer1:      r.Buffer1,
		Buffer2:      r.Buffer2,
		Buffer3:      r.Buffer3,
		Buffer4:      r.Buffer4,
		Buffer5:      r.Buffer5,
		Buffer6:      r.Buffer6,
		Buffer7:      r.Buffer7,
		Buffer8:      r.Buffer8,
		Buffer9:      r.Buffer9,
		Buffer10:     r.Buffer10,
		Buffer11:     r.Buffer11,
		Buffer12:     r.Buffer12,
	}

	u := query.PastureBuffer
	err = u.WithContext(context.Background()).Create(&pb)
	if err != nil {
		return err
	}
	return nil
}

func UploadCowHouse(r *request.ReqAddCowHouse) error {
	var err error
	cw := pasture.CowHouse{
		TimeRecordAt: time.Unix(r.TimeStamp, 0),
		HouseNumber:  r.HouseNumber,
		CowHouse1:    r.CowHouse1,
		CowHouse2:    r.CowHouse2,
		CowHouse3:    r.CowHouse3,
		CowHouse4:    r.CowHouse4,
		CowHouse5:    r.CowHouse5,
		CowHouse6:    r.CowHouse6,
		CowHouse7:    r.CowHouse7,
		CowHouse8:    r.CowHouse8,
		CowHouse9:    r.CowHouse9,
		CowHouse10:   r.CowHouse10,
		CowHouse11:   r.CowHouse11,
		CowHouse12:   r.CowHouse12,
	}

	q := query.CowHouse
	err = q.WithContext(context.Background()).Create(&cw)
	if err != nil {
		return err
	}
	return nil
}
func UploadPastureArea(r *request.ReqAddPastureArea) error {
	var err error
	pa := pasture.PastureArea{
		TimeRecordAt: time.Unix(r.TimeStamp, 0),
		HouseNumber:  r.HouseNumber,
		CattleFarm1:  r.CattleFarm1,
		CattleFarm2:  r.CattleFarm2,
		CattleFarm3:  r.CattleFarm3,
		CattleFarm4:  r.CattleFarm4,
		CattleFarm5:  r.CattleFarm5,
		CattleFarm6:  r.CattleFarm6,
		CattleFarm7:  r.CattleFarm7,
		CattleFarm8:  r.CattleFarm8,
		CattleFarm9:  r.CattleFarm9,
		CattleFarm10: r.CattleFarm10,
		CattleFarm11: r.CattleFarm11,
		CattleFarm12: r.CattleFarm12,
	}

	q := query.PastureArea
	err = q.WithContext(context.Background()).Create(&pa)
	if err != nil {
		return err
	}
	return nil
}

func UploadPastureBasicEnvironment(r *request.ReqAddPastureBasicEnvironment) error {
	var err error
	q := query.PastureBasicEnvironment

	be := pasture.PastureBasicEnvironment{
		TimeRecordAt: time.Unix(r.TimeStamp, 0),
		HouseNumber:  r.HouseNumber,
		Environment1: r.Environment1,
		Environment2: r.Environment2,
		Environment3: r.Environment3,
		Environment4: r.Environment4,
		Environment5: r.Environment5,
		Environment6: r.Environment6,
	}
	err = q.WithContext(context.Background()).Create(&be)

	if err != nil {
		return err
	}

	return nil
}

func UploadPasturePaddingRequire(r *request.ReqAddPasturePaddingRequire) error {
	var err error
	q := query.PasturePaddingRequire

	pr := pasture.PasturePaddingRequire{
		TimeRecordAt:    time.Unix(r.TimeStamp, 0),
		HouseNumber:     r.HouseNumber,
		PaddingRequire1: r.PaddingRequire1,
		PaddingRequire2: r.PaddingRequire2,
		PaddingRequire3: r.PaddingRequire3,
		PaddingRequire4: r.PaddingRequire4,
		PaddingRequire5: r.PaddingRequire5,
		PaddingRequire6: r.PaddingRequire6,
		PaddingRequire7: r.PaddingRequire7,
		PaddingRequire8: r.PaddingRequire8,
	}
	err = q.WithContext(context.Background()).Create(&pr)
	if err != nil {
		return err
	}
	return nil
}

func UploadPastureWastedWaterIndex(r *request.ReqAddPastureWastedWaterIndex) error {
	var err error
	q := query.PastureWastedWaterIndex

	ww := pasture.PastureWastedWaterIndex{
		TimeRecordAt:      time.Unix(r.TimeStamp, 0),
		HouseNumber:       r.HouseNumber,
		WastedWaterIndex1: r.WastedWaterIndex1,
		WastedWaterIndex2: r.WastedWaterIndex2,
		WastedWaterIndex3: r.WastedWaterIndex3,
		WastedWaterIndex4: r.WastedWaterIndex4,
		WastedWaterIndex5: r.WastedWaterIndex5,
		WastedWaterIndex6: r.WastedWaterIndex6,
		WastedWaterIndex7: r.WastedWaterIndex7,
		WastedWaterIndex8: r.WastedWaterIndex8,
		WastedWaterIndex9: r.WastedWaterIndex9,
	}
	err = q.WithContext(context.Background()).Create(&ww)

	if err != nil {
		return err
	}

	return nil
}

func UploadPastureDisinfectionRecord(r *request.ReqAddPastureDisinfectionRecord) error {
	var err error
	q := query.PastureDisinfectionRecord

	dr := pasture.PastureDisinfectionRecord{
		TimeRecordAt:   time.Unix(r.TimeStamp, 0),
		HouseNumber:    r.HouseNumber,
		FarmDisRecord1: r.FarmDisRecord1,
		FarmDisRecord2: r.FarmDisRecord2,
		FarmDisRecord3: r.FarmDisRecord3,
	}
	err = q.WithContext(context.Background()).Create(&dr)

	if err != nil {
		return err
	}

	return nil
}

func UploadPastureOdorPollutantsPerDay(r *request.ReqPastureOdorPollutantsPerDay) error {
	q := query.TotalOdorPollutantsPasturePerDay
	qall := query.Q.AllPasturesTrashDisposal
	//根据t去搜索总表，若已创建表，则直接更新，否则创建
	t := time.Unix(r.TimeStamp, 0).Truncate(24 * time.Hour)
	ok, err := PastureAllTrashTimeExisted(t)
	if !ok {
		aptp := pasture.AllPasturesTrashDisposal{
			TimeStamp:                     t,
			OdorPasturesTrashDisposal1:    r.PastureOdorPollutantsPerDay1,
			OdorPasturesTrashDisposal2:    r.PastureOdorPollutantsPerDay2,
			OdorPasturesTrashDisposal3:    r.PastureOdorPollutantsPerDay3,
			OdorPasturesTrashDisposal4:    r.PastureOdorPollutantsPerDay4,
			ResiduePasturesTrashDisposal1: 0,
			ResiduePasturesTrashDisposal2: 0,
			ResiduePasturesTrashDisposal3: 0,
			ResiduePasturesTrashDisposal4: 0,
			WaterPasturesTrashDisposal1:   0,
			WaterPasturesTrashDisposal2:   0,
			WaterPasturesTrashDisposal3:   0,
			WaterPasturesTrashDisposal4:   0,
		}
		err = qall.WithContext(context.Background()).Create(&aptp)
		if err != nil {
			return err
		}
	} else {
		result, err := qall.WithContext(context.Background()).Where(qall.TimeStamp.Eq(t)).UpdateSimple(qall.OdorPasturesTrashDisposal1.Add(r.PastureOdorPollutantsPerDay1), qall.OdorPasturesTrashDisposal2.Add(r.PastureOdorPollutantsPerDay2),
			qall.OdorPasturesTrashDisposal3.Add(r.PastureOdorPollutantsPerDay3), qall.OdorPasturesTrashDisposal4.Add(r.PastureOdorPollutantsPerDay4))
		if err != nil {
			return err
		}
		if result.RowsAffected == 0 {
			return errors.New("No match was found!")
		}
	}
	op := pasture.TotalOdorPollutantsPasturePerDay{
		TimeStamp:                  time.Unix(r.TimeStamp, 0).Truncate(24 * time.Hour),
		HouseNumber:                r.HouseNumber,
		TotalOdorPollutantsPerDay1: r.PastureOdorPollutantsPerDay1,
		TotalOdorPollutantsPerDay2: r.PastureOdorPollutantsPerDay2,
		TotalOdorPollutantsPerDay3: r.PastureOdorPollutantsPerDay3,
		TotalOdorPollutantsPerDay4: r.PastureOdorPollutantsPerDay4,
	}
	err = q.WithContext(context.Background()).Create(&op)
	if err != nil {
		return err
	}

	return nil
}

//	func QueryTotalWasteResidueOdorPerDay(r *request.ReqWasteResidueOdor) (pasture.TotalWasteResiduePasturePerDay, pasture.TotalOdorPollutantsPasturePerDay, error) {

func UploadPastureWasteWaterPerDay(r *request.ReqPastureWasteWaterPerDay) error {
	q := query.TotalWastedWaterPasturePerDay
	qall := query.AllPasturesTrashDisposal
	//根据t去搜索总表，若已创建表，则直接更新，否则创建
	t := time.Unix(r.TimeStamp, 0).Truncate(24 * time.Hour)
	ok, err := PastureAllTrashTimeExisted(t)
	if !ok {
		aptp := pasture.AllPasturesTrashDisposal{
			TimeStamp:                     t,
			OdorPasturesTrashDisposal1:    0,
			OdorPasturesTrashDisposal2:    0,
			OdorPasturesTrashDisposal3:    0,
			OdorPasturesTrashDisposal4:    0,
			ResiduePasturesTrashDisposal1: 0,
			ResiduePasturesTrashDisposal2: 0,
			ResiduePasturesTrashDisposal3: 0,
			ResiduePasturesTrashDisposal4: 0,
			WaterPasturesTrashDisposal1:   r.ReqPastureWasteWaterPerDay1,
			WaterPasturesTrashDisposal2:   r.ReqPastureWasteWaterPerDay2,
			WaterPasturesTrashDisposal3:   r.ReqPastureWasteWaterPerDay3,
			WaterPasturesTrashDisposal4:   r.ReqPastureWasteWaterPerDay4,
		}
		err = qall.WithContext(context.Background()).Create(&aptp)
		if err != nil {
			return err
		}
	} else {
		//更新总表信息
		result, _ := qall.WithContext(context.Background()).Where(qall.TimeStamp.Eq(t)).UpdateSimple(qall.WaterPasturesTrashDisposal1.Add(r.ReqPastureWasteWaterPerDay1),
			qall.WaterPasturesTrashDisposal2.Add(r.ReqPastureWasteWaterPerDay2), qall.WaterPasturesTrashDisposal3.Add(r.ReqPastureWasteWaterPerDay3))
		if result.RowsAffected == 0 {
			return errors.New("No match was found!")
		}
	}
	ww := pasture.TotalWastedWaterPasturePerDay{
		TimeStamp:               time.Unix(r.TimeStamp, 0).Truncate(24 * time.Hour),
		HouseNumber:             r.HouseNumber,
		TotalWastedWaterPerDay1: r.ReqPastureWasteWaterPerDay1,
		TotalWastedWaterPerDay2: r.ReqPastureWasteWaterPerDay2,
		TotalWastedWaterPerDay3: r.ReqPastureWasteWaterPerDay3,
		TotalWastedWaterPerDay4: r.ReqPastureWasteWaterPerDay4,
	}
	err = q.WithContext(context.Background()).Create(&ww)
	if err != nil {
		return err
	}

	return nil
}

func UploadPastureWasteResidue(r *request.ReqPastureWasteResiduePerDay) error {
	q := query.TotalWasteResiduePasturePerDay
	qall := query.Q.AllPasturesTrashDisposal
	//根据t去搜索总表，若已创建表，则直接更新，否则创建
	t := time.Unix(r.TimeStamp, 0).Truncate(24 * time.Hour)
	ok, err := PastureAllTrashTimeExisted(t)
	if !ok {
		aptp := pasture.AllPasturesTrashDisposal{
			TimeStamp:                     t,
			OdorPasturesTrashDisposal1:    0,
			OdorPasturesTrashDisposal2:    0,
			OdorPasturesTrashDisposal3:    0,
			OdorPasturesTrashDisposal4:    0,
			ResiduePasturesTrashDisposal1: r.PastureWasteResiduePerDay1,
			ResiduePasturesTrashDisposal2: r.PastureWasteResiduePerDay2,
			ResiduePasturesTrashDisposal3: r.PastureWasteResiduePerDay3,
			ResiduePasturesTrashDisposal4: r.PastureWasteResiduePerDay4,
			WaterPasturesTrashDisposal1:   0,
			WaterPasturesTrashDisposal2:   0,
			WaterPasturesTrashDisposal3:   0,
			WaterPasturesTrashDisposal4:   0,
		}
		err = qall.WithContext(context.Background()).Create(&aptp)
		if err != nil {
			return err
		}
	} else {
		//更新总表信息
		result, _ := qall.WithContext(context.Background()).Where(qall.TimeStamp.Eq(t)).UpdateSimple(qall.ResiduePasturesTrashDisposal1.Add(r.PastureWasteResiduePerDay1),
			qall.ResiduePasturesTrashDisposal2.Add(r.PastureWasteResiduePerDay2), qall.ResiduePasturesTrashDisposal3.Add(r.PastureWasteResiduePerDay3),
			qall.ResiduePasturesTrashDisposal4.Add(r.PastureWasteResiduePerDay4))
		if result.RowsAffected == 0 {
			return errors.New("No match was found!")
		}
	}

	wr := pasture.TotalWasteResiduePasturePerDay{
		TimeStamp:               time.Unix(r.TimeStamp, 0).Truncate(24 * time.Hour),
		HouseNumber:             r.HouseNumber,
		TotalWastedWaterPerDay1: r.PastureWasteResiduePerDay1,
		TotalWastedWaterPerDay2: r.PastureWasteResiduePerDay2,
		TotalWastedWaterPerDay3: r.PastureWasteResiduePerDay3,
		TotalWastedWaterPerDay4: r.PastureWasteResiduePerDay4,
	}
	err = q.WithContext(context.Background()).Create(&wr)
	if err != nil {
		return err
	}
	return nil
}

func AddPastureFeedHeavyMetal(r *request.ReqAddPastureFeedHeavyMetal) error {
	var err error
	tx := query.Q.Begin()
	defer func() {
		if recover() != nil {
			_ = tx.Rollback()
		}
	}()

	as := pasture.PastureFeedAs{
		PastureFeedHeavyMetalID: nil,
		As1:                     r.As.As1,
		As2:                     r.As.As2,
		As3:                     r.As.As3,
		As4:                     r.As.As4,
		As5:                     r.As.As5,
		As6:                     r.As.As6,
		As7:                     r.As.As7,
		As8:                     r.As.As8,
	}

	pb := pasture.PastureFeedPb{
		PastureFeedHeavyMetalID: nil,
		Pb1:                     r.Pb.Pb1,
		Pb2:                     r.Pb.Pb2,
		Pb3:                     r.Pb.Pb3,
		Pb4:                     r.Pb.Pb4,
		Pb5:                     r.Pb.Pb5,
		Pb6:                     r.Pb.Pb6,
		Pb7:                     r.Pb.Pb7,
	}

	cd := pasture.PastureFeedCd{
		PastureFeedHeavyMetalID: nil,
		Cd1:                     r.Cd.Cd1,
		Cd2:                     r.Cd.Cd2,
		Cd3:                     r.Cd.Cd3,
		Cd4:                     r.Cd.Cd4,
		Cd5:                     r.Cd.Cd5,
		Cd6:                     r.Cd.Cd6,
		Cd7:                     r.Cd.Cd7,
	}

	cr := pasture.PastureFeedCr{
		PastureFeedHeavyMetalID: nil,
		Cr1:                     r.Cr.Cr1,
		Cr2:                     r.Cr.Cr2,
		Cr3:                     r.Cr.Cr3,
		Cr4:                     r.Cr.Cr4,
	}

	hm := pasture.PastureFeedHeavyMetal{
		TimeRecordAt:  time.Unix(r.RecordAt, 0),
		HouseNumber:   r.HouseNumber,
		PastureFeedAs: as,
		PastureFeedPb: pb,
		PastureFeedCd: cd,
		PastureFeedCr: cr,
	}

	err = tx.PastureFeedHeavyMetal.WithContext(context.Background()).Create(&hm)
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
	//h, err := tx.HeavyMetal.WithContext(context.Background()).Where(tx.HeavyMetal.TimeRecordAt.)
}

func UploadPastureFeedMycotoxins(r *request.ReqAddPastureFeedCass) error {
	var err error
	tx := query.Q.Begin()
	defer func() {
		if recover() != nil {
			_ = tx.Rollback()
		}
	}()

	afb1 := pasture.Afb1{
		PastureFeedMycotoxinsID: nil,
		Afb11:                   r.Afb1.Afb11,
		Afb12:                   r.Afb1.Afb12,
		Afb13:                   r.Afb1.Afb13,
		Afb14:                   r.Afb1.Afb14,
		Afb15:                   r.Afb1.Afb15,
		Afb16:                   r.Afb1.Afb16,
		Afb17:                   r.Afb1.Afb17,
	}

	don := pasture.Don{
		PastureFeedMycotoxinsID: nil,
		Don1:                    r.Don.Don1,
		Don2:                    r.Don.Don2,
		Don3:                    r.Don.Don3,
		Don4:                    r.Don.Don4,
	}

	t2toxin := pasture.T2toxin{
		PastureFeedMycotoxinsID: nil,
		T2toxin1:                r.T2toxin.T2toxin1,
		T2toxin2:                r.T2toxin.T2toxin2,
		T2toxin3:                r.T2toxin.T2toxin3,
	}

	t2vomzea := pasture.T2VomZea{
		PastureFeedMycotoxinsID: nil,
		T2AVomZea1:              r.T2VomZea.T2AVomZea1,
		T2AVomZea2:              r.T2VomZea.T2AVomZea2,
		T2AVomZea3:              r.T2VomZea.T2AVomZea3,
	}

	pfm := pasture.PastureFeedMycotoxins{
		TimeRecordAt: time.Unix(r.RecordAt, 0),
		HouseNumber:  r.HouseNumber,
		Afb1:         afb1,
		Don:          don,
		T2toxin:      t2toxin,
		T2VomZea:     t2vomzea,
	}
	err = tx.PastureFeedMycotoxins.WithContext(context.Background()).Create(&pfm)
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

func UploadPastureWaterRecord(r *request.ReqAddPastureWaterRecord) error {
	var err error
	tx := query.Q.Begin()
	defer func() {
		if recover() != nil {
			_ = tx.Rollback()
		}
	}()

	oapgci := pasture.PastureOapGci{
		PastureWaterRecordID: nil,
		OapGci1:              r.OapGci.OapGci1,
		OapGci2:              r.OapGci.OapGci2,
		OapGci3:              r.OapGci.OapGci3,
		OapGci4:              r.OapGci.OapGci4,
		OapGci5:              r.OapGci.OapGci5,
		OapGci6:              r.OapGci.OapGci6,
		OapGci7:              r.OapGci.OapGci7,
		OapGci8:              r.OapGci.OapGci8,
		OapGci9:              r.OapGci.OapGci9,
		OapGci10:             r.OapGci.OapGci10,
	}

	toxindex := pasture.PastureToxIndex{
		PastureWaterRecordID: nil,
		ToxIndex1:            r.ToxIndex.ToxIndex1,
		ToxIndex2:            r.ToxIndex.ToxIndex2,
		ToxIndex3:            r.ToxIndex.ToxIndex3,
		ToxIndex4:            r.ToxIndex.ToxIndex4,
		ToxIndex5:            r.ToxIndex.ToxIndex5,
		ToxIndex6:            r.ToxIndex.ToxIndex6,
		ToxIndex7:            r.ToxIndex.ToxIndex7,
		ToxIndex8:            r.ToxIndex.ToxIndex8,
		ToxIndex9:            r.ToxIndex.ToxIndex9,
		ToxIndex10:           r.ToxIndex.ToxIndex10,
		ToxIndex11:           r.ToxIndex.ToxIndex11,
		ToxIndex12:           r.ToxIndex.ToxIndex12,
		ToxIndex13:           r.ToxIndex.ToxIndex13,
		ToxIndex14:           r.ToxIndex.ToxIndex14,
		ToxIndex15:           r.ToxIndex.ToxIndex15,
		ToxIndex16:           r.ToxIndex.ToxIndex16,
		ToxIndex17:           r.ToxIndex.ToxIndex17,
		ToxIndex18:           r.ToxIndex.ToxIndex18,
		ToxIndex19:           r.ToxIndex.ToxIndex19,
		ToxIndex20:           r.ToxIndex.ToxIndex20,
		ToxIndex21:           r.ToxIndex.ToxIndex21,
		ToxIndex22:           r.ToxIndex.ToxIndex22,
	}

	microindex := pasture.PastureMicroIndex{
		PastureWaterRecordID: nil,
		MicroIndex1:          r.MicroIndex.MicroIndex1,
		MicroIndex2:          r.MicroIndex.MicroIndex2,
		MicroIndex3:          r.MicroIndex.MicroIndex3,
	}

	pwr := pasture.PastureWaterRecord{
		TimeRecordAt: time.Unix(r.RecordAt, 0),
		HouseNumber:  r.HouseNumber,
		OapGci:       oapgci,
		ToxIndex:     toxindex,
		MicroIndex:   microindex,
	}

	err = tx.PastureWaterRecord.WithContext(context.Background()).Create(&pwr)
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

func SendToSlaughter(r *request.ReqSendToSlaughter) error {
	var err error
	tx := query.Q.Begin()
	defer func() {
		if recover() != nil {
			_ = tx.Rollback()
		}
	}()

	pid, err := GetPidByCowNumber(r.CowNumber)
	if err != nil {
		return err
	}

	phinfo, err := GetPastureHouseInfoByCowNumber(r.CowNumber)
	if err != nil {
		return err
	}

	sh, err := GetSlaughterInfoByNumber(r.SlaughterHouseNumber)
	if err != nil {
		return err
	}

	// 更新cow状态
	_, err = tx.Cow.WithContext(context.Background()).
		Where(tx.Cow.Number.Eq(r.CowNumber)).
		Updates(map[string]interface{}{"state": SENDING_STATE_COW})
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// 更新pasture warehouse状态
	_, err = tx.PastureWarehouse.WithContext(context.Background()).
		Where(tx.PastureWarehouse.CowNumber.Eq(r.CowNumber)).
		Updates(map[string]interface{}{"state": SENDING_STATE_PH, "destination": sh.Name, "out_timestamp": time.Now()})
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// 新增slaughter receive record
	srr := warehouse.SlaughterReceiveRecord{
		CowNumber:       r.CowNumber,
		PID:             pid,
		SourceNumber:    phinfo.HouseNumber,
		SourceName:      phinfo.Name,
		State:           INIT_STATE_REC_SLA,
		Operator:        r.Operator,
		ReceiveTime:     time.Now(),
		ConfirmTime:     nil,
		SlaughterNumber: r.SlaughterHouseNumber,
	}
	err = tx.SlaughterReceiveRecord.WithContext(context.Background()).Create(&srr)
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

func GetPastureHouseInfoByCowNumber(num string) (pasture.PastureHouseInfo, error) {
	q1 := query.PastureWarehouse
	pw, err := q1.WithContext(context.Background()).Where(q1.CowNumber.Eq(num)).First()
	if err != nil {
		return pasture.PastureHouseInfo{}, err
	}

	q2 := query.PastureHouse
	ph, err := q2.WithContext(context.Background()).Where(q2.HouseNumber.Eq(pw.HouseNumber)).First()
	if err != nil {
		return pasture.PastureHouseInfo{}, err
	}

	phinfo := pasture.ToPastureHouseInfo(*ph)

	return phinfo, nil
}

func GetPidByCowNumber(num string) (string, error) {
	q := query.PastureWarehouse
	pw, err := q.WithContext(context.Background()).Where(q.CowNumber.Eq(num)).First()
	if err != nil {
		return "", err
	}

	return pw.PID, nil
}

func GetWarehouseInfosByPastureNumber(num string) ([]warehouse.PastureWarehouseInfo, int64, error) {
	q := query.PastureWarehouse
	pws, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(num)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(pws)
	records := make([]warehouse.PastureWarehouseInfo, count)
	for i, pw := range pws {
		records[i] = warehouse.ToPastureWarehouseInfo(*pw)
	}

	return records, int64(count), nil
}

func EndFeeding(r *request.ReqEndFeeding) (string, []string, error) {
	var err error
	tx := query.Q.Begin()

	defer func() {
		if recover() != nil {
			_ = tx.Rollback()
		}
	}()

	endTime := time.Now()

	// 读取PID
	pid, err := GetPidByFeedingBatchNumber(r.BatchNumber)
	if err != nil {
		return "", nil, err
	}
	cowsNum, err := GetCowsNumberByFeedingBatchNumber(r.BatchNumber)
	if err != nil {
		return "", nil, err
	}

	// 更新Cow状态
	_, err = tx.Cow.WithContext(context.Background()).Where(tx.Cow.BatchNumber.Eq(r.BatchNumber)).Updates(map[string]interface{}{"state": WAREHOUSE_STATE_COW})
	if err != nil {
		_ = tx.Rollback()
		return "", nil, err
	}
	// 更新FeedingBatch状态
	_, err = tx.FeedingBatch.WithContext(context.Background()).Where(tx.FeedingBatch.BatchNumber.Eq(r.BatchNumber)).Updates(map[string]interface{}{"state": END_STATE_BATCH_PAS, "end_time": endTime})
	if err != nil {
		_ = tx.Rollback()
		return "", nil, err
	}
	// 写入仓库
	pws := make([]*warehouse.PastureWarehouse, len(cowsNum))
	for i, num := range cowsNum {
		pw := warehouse.PastureWarehouse{
			CowNumber:    num,
			PID:          pid,
			Type:         "",
			State:        INIT_STATE_PH,
			InOperator:   r.Worker,
			OutOperator:  "",
			Destination:  "",
			InTimestamp:  time.Now(),
			OutTimestamp: nil,
			HouseNumber:  r.HouseNumber,
		}
		pws[i] = &pw
	}
	err = tx.PastureWarehouse.WithContext(context.Background()).Create(pws...)
	if err != nil {
		_ = tx.Rollback()
		return "", nil, err
	}

	// 提交Procedure
	data := pasture.PastureProcedureData{
		PM10:   r.PM10,
		TSP:    r.TSP,
		Stench: r.Stench,
	}

	checkcode, pHash, err := BasicCommitProcedureWithTx(tx, pid, &endTime, data)
	if err != nil {
		_ = tx.Rollback()
		return "", nil, err
	}
	_, err = fabric.UpdateProcedure(pid, pHash)
	if err != nil {
		_ = tx.Rollback()
		return "", nil, err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return "", nil, err
	}

	return checkcode, cowsNum, nil
}

func GetCowsNumberByFeedingBatchNumber(num string) ([]string, error) {
	q := query.Cow
	cows, err := q.WithContext(context.Background()).Where(q.BatchNumber.Eq(num)).Find()
	if err != nil {
		return nil, err
	}

	nums := make([]string, len(cows))
	for i, cow := range cows {
		nums[i] = cow.Number
	}

	return nums, nil
}

func GetPidByFeedingBatchNumber(num string) (string, error) {
	q := query.FeedingBatch
	fb, err := q.WithContext(context.Background()).Where(q.BatchNumber.Eq(num)).First()
	if err != nil {
		return "", err
	}

	return fb.PID, nil
}

func GetFeedingRecords(houseNum string) ([]pasture.FeedingBatchInfo, int64, error) {
	q := query.FeedingBatch
	fbs, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(houseNum)).Preload(q.Cows).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(fbs)
	records := make([]pasture.FeedingBatchInfo, count)
	for i, fb := range fbs {
		records[i] = pasture.ToFeedingBatchInfo(fb)
	}

	return records, int64(count), nil
}

func NewFeedingBatch(r *request.ReqNewFeedingBatch) (string, error) {
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
		Type:        PASTURE_TYPE,
		Operator:    r.Worker,
		PrePID:      r.PrePID,
		BatchNumber: bNum,
	}

	procedure, err := NewProcedure(&pp, startTime)

	glog.Infoln("CowNumbers:")
	glog.Infoln(r.CowNumbers)
	cows, err := GetCowsByNumbers(r.CowNumbers)

	glog.Infoln("Cows:")
	glog.Infoln(cows)

	fb := pasture.FeedingBatch{
		BatchNumber: bNum,
		HouseNumber: r.HouseNumber,
		State:       INIT_STATE_BATCH_PAS,
		PID:         procedure.PID,
		Worker:      r.Worker,
		Cows:        cows,
		StartTime:   &startTime,
	}
	err = tx.FeedingBatch.WithContext(context.Background()).Create(&fb)

	if err != nil {
		_ = tx.Rollback()
		return "", err
	}
	err = tx.Procedure.WithContext(context.Background()).Create(&procedure)
	if err != nil {
		_ = tx.Rollback()
		return "", err
	}

	for _, number := range r.CowNumbers {
		_, err = tx.Cow.WithContext(context.Background()).Where(tx.Cow.Number.Eq(number)).Updates(map[string]interface{}{"state": FEEDING_STATE_COW})
		if err != nil {
			_ = tx.Rollback()
			return "", err
		}
	}
	//上链
	_, err = fabric.UploadProcedure(procedure.PID, procedure.PrePID)
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

func GetCowList(house_number string) ([]product.Cow, error) {
	q := query.Cow
	cows, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(house_number)).Find()
	if err != nil {
		return nil, err
	}
	var cowsRes []product.Cow
	for i, _ := range cows {
		cowsRes = append(cowsRes, *cows[i])
	}
	return cowsRes, nil
}

func GetCowsByNumbers(numbers []string) ([]product.Cow, error) {
	cows := make([]product.Cow, len(numbers))
	q := query.Cow
	for i, number := range numbers {
		c, err := q.WithContext(context.Background()).Where(q.Number.Eq(number)).First()
		if err != nil {
			return nil, err
		}

		cows[i] = *c
	}

	return cows, nil
}

func PastureIsExist(number string) (bool, error) {
	q := query.PastureHouse
	count, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(number)).Count()
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}

	return true, nil
}

func AddCow(r *request.ReqAddCow) (product.CowInfo, error) {
	cowNum := GenerateNumber(r)

	et := time.Unix(r.EntryTime, 0).UTC().Local()

	cow := product.Cow{
		Number:               COW_NUMBER_PREFIX + cowNum,
		Age:                  r.Age,
		Weight:               r.Weight,
		QuarantineCertNumber: &r.QuarantineCertNumber,
		OwnerIDCard:          r.OwnerIDCard,
		OwnerAddress:         r.OwnerAddress,
		EntryTime:            &et,
		State:                INIT_STATE_COW,
		HouseNumber:          r.HouseNumber,
		BatchNumber:          nil,
	}

	q := query.Cow
	err := q.WithContext(context.Background()).Create(&cow)
	if err != nil {
		return product.CowInfo{}, err
	}

	cowInfo := product.ToCowInfo(&cow)
	return cowInfo, nil
}

func createPastureProcedure(pasPID string) pasture.PastureProcedure {
	pastureProcedure := pasture.PastureProcedure{
		PasPID: pasPID,
		Water: pasture.PastureWater{
			PhysicalHazard: pasture.PastureWaterPhysicalHazard{},
			ChemicalHazard: pasture.PastureWaterChemicalHazard{},
			Biohazard:      pasture.PastureWaterBiohazard{},
			SensoryTraits:  pasture.PastureWaterSensoryTraits{},
			PasPID:         pasPID,
		},
		Fodder: pasture.PastureFodder{
			PhysicalHazard: pasture.PastureFodderPhysicalHazard{},
			Biohazard:      pasture.PastureFodderBiohazard{},
			PasPID:         pasPID,
		},
		Soil: pasture.PastureSoil{
			PhysicalHazard: pasture.PastureSoilPhysicalHazard{},
			Biohazard:      pasture.PastureSoilBiohazard{},
			PasPID:         pasPID,
		},
		Air: pasture.PastureAir{
			TotalBacteria:   0,
			AmmoniaGas:      0,
			HydrogenSulfide: 0,
			CarbonDioxide:   0,
			PM10:            0,
			TSP:             0,
			Stench:          0,
			PasPID:          pasPID,
		},
		FloorBedding: pasture.PastureFloorBedding{
			PhysicalHazard: pasture.PastureFloorBeddingPhysicalHazard{},
			Biohazard:      pasture.PastureFloorBeddingBiohazard{},
			PasPID:         pasPID,
		},
		WasteDischarge: pasture.PastureWasteDischarge{
			WaterFivedayBOD:          0,
			WaterChemicalOxygen:      0,
			WaterAmmoniaNitrogen:     0,
			WaterPhosphorus:          0,
			WaterSuspendedMatter:     0,
			WaterFecalColiform:       0,
			WaterAO:                  0,
			WasteSlagFecalColiform:   0,
			WasteSlagAOMortalityRate: 0,
			O3Concentration:          0,
			PasPID:                   pasPID,
		},
	}

	return pastureProcedure
}

func PastureHouseIsExisted(number string) (bool, error) {
	q := query.Q.PastureHouse
	pas, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(number)).Find()
	if err != nil {
		return false, err
	}

	if len(pas) == 0 {
		return false, errors.New("pasture house is not existed")
	}

	return true, nil
}

func PastureAllTrashTimeExisted(t time.Time) (bool, error) {
	//先对时间进行处理
	truncated := t.Truncate(24 * time.Hour)
	q := query.Q.AllPasturesTrashDisposal
	info, err := q.WithContext(context.Background()).Where(q.TimeStamp.Eq(truncated)).Find()
	if err != nil {
		return false, err
	}

	if len(info) == 0 {
		return false, errors.New("Garbage disposal information has not been recorded for this time in pasture")
	}

	return true, nil
}
