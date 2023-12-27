package service

import (
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/models/coldchain"
	"CN-EU-FSIMS/internal/app/models/pack"
	"CN-EU-FSIMS/internal/app/models/product"
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/internal/app/models/warehouse"
	"context"
	"time"

	"github.com/golang/glog"
)

const (
	INIT_STATE_REC_PAC    = 1
	CONFIRM_STATE_REC_PAC = 2
	PACKAGE_STATE_REC_PAC = 3
	END_STATE_REC_PAC     = 4

	INIT_STATE_BATCH_PAC = 1
	END_STATE_BATCH_PAC  = 2

	INIT_PACPRO          = 1
	WAREHOUSE_PACPRO     = 2
	PRE_TRANSPORT_PACPRO = 3
	TRANSPORT_PACPRO     = 4
	TRANSPORT_END_PACPRO = 5

	INIT_PH          = 1
	PRE_TRANSPORT_PH = 2
	TRANSPORT_PH     = 3
	TRANSPORT_END_PH = 4

	PACKAGE_PRODUCT_PREFIX = "PACPRO-"

	BAG = 1
	BOX = 2
)

func PreTransport(r *request.ReqPreTransport) error {
	var err error
	tx := query.Q.Begin()

	defer func() {
		if recover() != nil {
			_ = tx.Rollback()
		}
	}()
	bNum := BATCH_NUMBER_PREFIX + GenerateNumber(r)
	// 获取package products
	packProducts := make([]product.PackageProduct, len(r.PackageProductNumbers))

	// 更新Package product状态（待运输确认）
	// 更新Package house状态（待运输确认）
	for i, number := range r.PackageProductNumbers {
		_, err = tx.PackageProduct.WithContext(context.Background()).
			Where(tx.PackageProduct.Number.Eq(number)).
			Updates(map[string]interface{}{"state": PRE_TRANSPORT_PACPRO})
		if err != nil {
			_ = tx.Rollback()
			glog.Info("1")
			return err
		}

		_, err = tx.PackWarehouse.WithContext(context.Background()).
			Where(tx.PackWarehouse.ProductNumber.Eq(number)).
			Updates(map[string]interface{}{"state": PRE_TRANSPORT_PH})
		if err != nil {
			_ = tx.Rollback()
			return err
		}

		p, err := tx.PackageProduct.WithContext(context.Background()).
			Where(tx.PackageProduct.Number.Eq(number)).First()
		if err != nil {
			_ = tx.Rollback()
			return err
		}

		packProducts[i] = *p
	}

	// 新增transport batch，设置初始状态还未确认运输
	tBatch := coldchain.TransportBatch{
		BatchNumber: bNum,
		TVNumber:    r.TVNumber,
		State:       0,
		Worker:      r.Worker,
		MallNumber:  r.MallNumber,
		Products:    packProducts,
	}

	err = tx.TransportBatch.WithContext(context.Background()).Create(&tBatch)
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

func EndPackage(r *request.ReqEndPackage) (string, []string, error) {
	var err error
	tx := query.Q.Begin()

	defer func() {
		if recover() != nil {
			_ = tx.Rollback()
		}
	}()

	// 读取PID和product number
	pid, productNumber, err := GetPidAndProductNumByPackageBatchNumber(r.BatchNumber)
	if err != nil {
		return "", nil, err
	}

	// 获取package products
	products, err := GetPackageProductsByBatchNumber(r.BatchNumber)
	if err != nil {
		return "", nil, err
	}

	// 更新receive record状态
	_, err = tx.PackageReceiveRecord.WithContext(context.Background()).
		Where(tx.PackageReceiveRecord.ProductNumber.Eq(productNumber)).
		Updates(map[string]interface{}{"state": END_STATE_REC_PAC})
	if err != nil {
		_ = tx.Rollback()
		return "", nil, err
	}

	// 更新PackageBatch状态
	_, err = tx.PackageBatch.WithContext(context.Background()).
		Where(tx.PackageBatch.BatchNumber.Eq(r.BatchNumber)).
		Updates(map[string]interface{}{"state": END_STATE_BATCH_PAC})
	if err != nil {
		_ = tx.Rollback()
		return "", nil, err
	}

	// 写入仓库
	count := len(products)
	pws := make([]*warehouse.PackWarehouse, count)
	productNums := make([]string, count)
	for i, p := range products {
		pw := warehouse.PackWarehouse{
			ProductNumber: p.Number,
			PID:           pid,
			Type:          p.TypeName,
			State:         INIT_PH,
			InOperator:    r.Worker,
			OutOperator:   "",
			VehicleNumber: nil,
			InTimestamp:   time.Now(),
			OutTimestamp:  nil,
			PacNumber:     r.HouseNumber,
		}
		pws[i] = &pw
		productNums[i] = p.Number
	}
	err = tx.PackWarehouse.WithContext(context.Background()).Create(pws...)
	if err != nil {
		_ = tx.Rollback()
		return "", nil, err
	}

	// 更新product状态
	_, err = tx.PackageProduct.WithContext(context.Background()).
		Where(tx.PackageProduct.BatchNumber.Eq(r.BatchNumber)).
		Updates(map[string]interface{}{"state": WAREHOUSE_PACPRO})
	if err != nil {
		_ = tx.Rollback()
		return "", nil, err
	}

	// 提交Procedure
	data := pack.PackProcedureData{
		PackType:        r.PackType,
		PackTemperature: r.PackTemperature,
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

func GetPackageProductsByBatchNumber(num string) ([]*product.PackageProduct, error) {
	q := query.PackageProduct
	sps, err := q.WithContext(context.Background()).Where(q.BatchNumber.Eq(num)).Find()
	if err != nil {
		return nil, err
	}

	return sps, nil
}

func GetPidAndProductNumByPackageBatchNumber(num string) (string, string, error) {
	q := query.PackageBatch
	fb, err := q.WithContext(context.Background()).Where(q.BatchNumber.Eq(num)).First()
	if err != nil {
		return "", "", err
	}

	return fb.PID, fb.ProductNumber, nil
}

func GetPidByPackageProductNumber(num string) (string, error) {
	q := query.SlaughterWarehouse
	sw, err := q.WithContext(context.Background()).Where(q.ProductNumber.Eq(num)).First()
	if err != nil {
		return "", err
	}

	return sw.PID, nil
}

func NewPackageProduct(r *request.ReqNewPackageProduct) (string, error) {
	number := PACKAGE_PRODUCT_PREFIX + GenerateNumber(r)
	p := product.PackageProduct{
		Number:         number,
		Type:           r.Type,
		TypeName:       typeName(r.Type),
		PackMethod:     r.PackMethod,
		PackMethodName: packMethodName(r.PackMethod),
		ShelfLife:      r.ShelfLife,
		Weight:         r.Weight,
		State:          INIT_PACPRO,
		HouseNumber:    r.HouseNumber,
		BatchNumber:    r.BatchNumber,
	}

	q := query.PackageProduct
	err := q.WithContext(context.Background()).Create(&p)
	if err != nil {
		return "", err
	}

	return number, nil
}

func packMethodName(method int) string {
	switch method {
	case BAG:
		return "袋装"
	case BOX:
		return "盒装"
	default:
		return "未知"
	}
}

func NewPackageBatch(r *request.ReqNewPackageBatch) (string, error) {
	var err error
	tx := query.Q.Begin()

	defer func() {
		if recover() != nil {
			_ = tx.Rollback()
		}
	}()

	pp := NewProcedureParams{
		Type:     PACKAGE_TYPE,
		Operator: r.Worker,
		PrePID:   r.PrePID,
	}
	procedure, err := NewProcedure(&pp)

	bNum := BATCH_NUMBER_PREFIX + GenerateNumber(r)

	pb := pack.PackageBatch{
		BatchNumber:   bNum,
		HouseNumber:   r.HouseNumber,
		State:         INIT_STATE_BATCH_PAC,
		PID:           procedure.PID,
		Worker:        r.Worker,
		ProductNumber: r.ProductNumber,
	}

	err = tx.Procedure.WithContext(context.Background()).Create(&procedure)
	if err != nil {
		_ = tx.Rollback()
		return "", err
	}
	err = tx.PackageBatch.WithContext(context.Background()).Create(&pb)
	if err != nil {
		_ = tx.Rollback()
		return "", err
	}

	// 更新package receive record
	_, err = tx.PackageReceiveRecord.WithContext(context.Background()).
		Where(tx.PackageReceiveRecord.ProductNumber.Eq(r.ProductNumber)).
		Updates(map[string]interface{}{"state": PACKAGE_STATE_REC_PAC})
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

func ConfirmProductFromSlaughter(num string) error {
	var err error
	tx := query.Q.Begin()
	glog.Info("num", num)
	defer func() {
		if recover() != nil {
			_ = tx.Rollback()
		}
	}()

	q1 := tx.PackageReceiveRecord
	_, err = q1.WithContext(context.Background()).
		Where(q1.ProductNumber.Eq(num)).
		Updates(map[string]interface{}{"state": CONFIRM_STATE_REC_PAC, "confirm_time": time.Now()})
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	q2 := tx.SlaughterWarehouse
	_, err = q2.WithContext(context.Background()).
		Where(q2.ProductNumber.Eq(num)).
		Updates(map[string]interface{}{"state": CONFIRM_SH})
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	q3 := tx.SlaughterProduct
	_, err = q3.WithContext(context.Background()).
		Where(q3.Number.Eq(num)).
		Updates(map[string]interface{}{"state": CONFIRM_SLAPRO})
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

func GetPackageReceiveRecords(houseNum string) ([]warehouse.PackageReceiveRecordInfo, int64, error) {
	q := query.PackageReceiveRecord
	prs, err := q.WithContext(context.Background()).Where(q.PacNumber.Eq(houseNum)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(prs)
	records := make([]warehouse.PackageReceiveRecordInfo, count)
	for i, pr := range prs {
		records[i] = warehouse.ToPackageReceiveRecordInfo(pr)
	}

	return records, int64(count), nil
}

func GetPackageBatches(houseNum string) ([]pack.PackageBatchInfo, int64, error) {
	q := query.PackageBatch
	pbs, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(houseNum)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(pbs)
	records := make([]pack.PackageBatchInfo, count)
	for i, pb := range pbs {
		records[i] = pack.ToPackageBatchInfo(pb)
	}

	return records, int64(count), nil
}

func GetPackageWarehouseRecords(houseNum string) ([]warehouse.PackWarehouseInfo, int64, error) {
	q := query.PackWarehouse
	pws, err := q.WithContext(context.Background()).Where(q.PacNumber.Eq(houseNum)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(pws)
	records := make([]warehouse.PackWarehouseInfo, count)
	for i, sw := range pws {
		records[i] = warehouse.ToPackWarehouseInfo(sw)
	}

	return records, int64(count), nil
}

func GetPackageInfoByNumber(num string) (pack.PackageHouseInfo, error) {
	q := query.PackageHouse
	ph, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(num)).First()
	if err != nil {
		return pack.PackageHouseInfo{}, err
	}

	phinfo := pack.ToPackageHouseInfo(ph)
	return phinfo, nil
}

func GetPackageProducts(num string) ([]product.PackageProductInfo, int64, error) {
	q := query.PackageProduct
	products, err := q.WithContext(context.Background()).Where(q.HouseNumber.Eq(num)).Find()
	if err != nil {
		return nil, 0, err
	}
	count := len(products)
	records := make([]product.PackageProductInfo, count)
	for i, sw := range products {
		records[i] = product.ToPackageProductInfo(sw)
	}

	return records, int64(count), nil
}
