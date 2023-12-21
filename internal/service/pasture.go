package service

import (
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/models/pasture"
	"CN-EU-FSIMS/internal/app/models/product"
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/internal/app/models/warehouse"
	"context"
	"time"

	"github.com/golang/glog"
)

const (
	COW_NUMBER_PREFIX   = "COW-"
	BATCH_NUMBER_PREFIX = "BATCH-"

	// cow state
	INIT_STATE_COW      = 1
	FEEDING_STATE_COW   = 2
	WAREHOUSE_STATE_COW = 3
	SENDING_STATE_COW   = 4
	END_STATE_COW       = 5

	// pasture batch state
	INIT_STATE_BATCH_PAS = 1
	END_STATE_BATCH_PAS  = 2

	// pasture house state
	INIT_STATE_PH    = 1
	SENDING_STATE_PH = 2
	CONFIRM_STATE_PH = 3
)

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
		Updates(map[string]interface{}{"state": SENDING_STATE_PH, "destination": sh.Name})
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
	_, err = tx.FeedingBatch.WithContext(context.Background()).Where(tx.FeedingBatch.BatchNumber.Eq(r.BatchNumber)).Updates(map[string]interface{}{"state": END_STATE_BATCH_PAS})
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

	checkcode, err := BasicCommitProcedureWithTx(tx, pid, data)
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

	pp := NewProcedureParams{
		Type:     PASTURE_TYPE,
		Operator: r.Worker,
		PrePID:   r.PrePID,
	}
	procedure, err := NewProcedure(&pp)

	glog.Infoln("CowNumbers:")
	glog.Infoln(r.CowNumbers)
	cows, err := GetCowsByNumbers(r.CowNumbers)

	glog.Infoln("Cows:")
	glog.Infoln(cows)

	bNum := BATCH_NUMBER_PREFIX + GenerateNumber(r)
	fb := pasture.FeedingBatch{
		BatchNumber: bNum,
		HouseNumber: r.HouseNumber,
		State:       INIT_STATE_BATCH_PAS,
		PID:         procedure.PID,
		Worker:      r.Worker,
		Cows:        cows,
	}

	err = tx.Procedure.WithContext(context.Background()).Create(&procedure)
	if err != nil {
		_ = tx.Rollback()
		return "", err
	}
	err = tx.FeedingBatch.WithContext(context.Background()).Create(&fb)
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
	cow := product.Cow{
		Number:      COW_NUMBER_PREFIX + cowNum,
		Age:         r.Age,
		Weight:      r.Weight,
		State:       INIT_STATE_COW,
		HouseNumber: r.HouseNumber,
		BatchNumber: nil,
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
