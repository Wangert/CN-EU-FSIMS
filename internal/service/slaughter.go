package service

import (
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/models/product"
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/internal/app/models/slaughter"
	"CN-EU-FSIMS/internal/app/models/warehouse"
	"context"
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
		Updates(map[string]interface{}{"state": END_STATE_BATCH_SLA})
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

	pp := NewProcedureParams{
		Type:     SLAUGHTER_TYPE,
		Operator: r.Worker,
		PrePID:   r.PrePID,
	}
	procedure, err := NewProcedure(&pp)

	bNum := BATCH_NUMBER_PREFIX + GenerateNumber(r)

	sb := slaughter.SlaughterBatch{
		BatchNumber: bNum,
		HouseNumber: r.HouseNumber,
		State:       INIT_STATE_BATCH_SLA,
		PID:         procedure.PID,
		Worker:      r.Worker,
		CowNumber:   r.CowNumber,
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
