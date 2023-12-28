package service

import (
	"CN-EU-FSIMS/fabric"
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/models/coldchain"
	"CN-EU-FSIMS/internal/app/models/product"
	"CN-EU-FSIMS/internal/app/models/query"
	"context"
	"time"
)

const (
	START_STATE_BATCH_TRANS = 1
	END_STATE_BATCH_TRANS   = 2

	MALL_GOOD_PREFIX = "MG"
)

func EndTransport(r *request.ReqEndTransport) (string, error) {
	var err error
	tx := query.Q.Begin()

	defer func() {
		if recover() != nil {
			_ = tx.Rollback()
		}
	}()

	batch, err := tx.TransportBatch.WithContext(context.Background()).
		Where(tx.TransportBatch.BatchNumber.Eq(r.BatchNumber)).
		Preload(tx.TransportBatch.Products).First()
	if err != nil {
		return "", err
	}

	packageProducts := batch.Products

	goods := make([]*product.MallGood, len(packageProducts))

	// 为该运输批次中的package product提交transport procedure数据
	for i, p := range packageProducts {
		m, err := tx.PackageProductAndTransportPIDMap.WithContext(context.Background()).
			Where(tx.PackageProductAndTransportPIDMap.PackageProductNumber.Eq(p.Number)).First()
		if err != nil {
			_ = tx.Rollback()
			return "", err
		}

		// 提交Procedure
		data := coldchain.TransportProcedureData{
			Temperature: r.Temperature,
			Source:      r.Source,
			Destination: r.Destination,
			Humidity:    r.Humidity,
		}

		_, err = tx.PackageProduct.WithContext(context.Background()).
			Where(tx.PackageProduct.Number.Eq(p.Number)).
			Updates(map[string]interface{}{"state": TRANSPORT_END_PACPRO})
		if err != nil {
			_ = tx.Rollback()
			return "", err
		}

		_, err = tx.PackWarehouse.WithContext(context.Background()).
			Where(tx.PackWarehouse.ProductNumber.Eq(p.Number)).
			Updates(map[string]interface{}{"state": TRANSPORT_END_PH})
		if err != nil {
			_ = tx.Rollback()
			return "", err
		}

		checkcode, phash, err := BasicCommitProcedureWithTx(tx, m.TransportPID, data)
		if err != nil {
			_ = tx.Rollback()
			return "", err
		}

		_, err = fabric.UpdateProcedure(m.TransportPID, phash)
		if err != nil {
			_ = tx.Rollback()
			return "", err
		}

		mgNumber := MALL_GOOD_PREFIX + "-" + GenerateNumber(p)

		// 创建mall good
		mg := product.MallGood{
			Number:         mgNumber,
			Type:           p.Type,
			TypeName:       p.TypeName,
			PackMethod:     p.PackMethod,
			PackMethodName: p.PackMethodName,
			ProductionDate: time.Now(),
			ShelfLife:      p.ShelfLife,
			Weight:         p.Weight,
			State:          1,
			MallNumber:     batch.MallNumber,
			FinalPID:       &m.TransportPID,
			Checkcode:      checkcode,
			BuyerIDCard:    nil,
		}

		goods[i] = &mg
	}

	err = tx.MallGood.WithContext(context.Background()).Create(goods...)
	if err != nil {
		_ = tx.Rollback()
		return "", err
	}

	_, err = tx.TransportBatch.WithContext(context.Background()).
		Where(tx.TransportBatch.BatchNumber.Eq(r.BatchNumber)).
		Updates(map[string]interface{}{"state": END_STATE_BATCH_TRANS})
	if err != nil {
		_ = tx.Rollback()
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return "", err
	}

	return batch.BatchNumber, nil
}

func StartTransport(r *request.ReqStartTransport) error {
	var err error
	tx := query.Q.Begin()

	defer func() {
		if recover() != nil {
			_ = tx.Rollback()
		}
	}()

	batch, err := tx.TransportBatch.WithContext(context.Background()).
		Where(tx.TransportBatch.BatchNumber.Eq(r.BatchNumber)).
		Preload(tx.TransportBatch.Products).First()
	if err != nil {
		return err
	}

	packageProducts := batch.Products

	// 为每个package product创建transport procedure
	for _, p := range packageProducts {
		packBatch, err := tx.PackageBatch.WithContext(context.Background()).
			Where(tx.PackageBatch.BatchNumber.Eq(p.BatchNumber)).First()
		if err != nil {
			_ = tx.Rollback()
			return err
		}

		pp := NewProcedureParams{
			Type:     COLDCHAIN_TRANSPORT_TYPE,
			Operator: r.Operator,
			PrePID:   packBatch.PID,
		}
		procedure, err := NewProcedure(&pp)
		if err != nil {
			_ = tx.Rollback()
			return err
		}

		m := coldchain.PackageProductAndTransportPIDMap{
			PackageProductNumber: p.Number,
			TransportPID:         procedure.PID,
		}

		err = tx.PackageProductAndTransportPIDMap.WithContext(context.Background()).Create(&m)
		if err != nil {
			_ = tx.Rollback()
			return err
		}

		err = tx.Procedure.WithContext(context.Background()).Create(&procedure)
		if err != nil {
			_ = tx.Rollback()
			return err
		}

		_, err = tx.PackageProduct.WithContext(context.Background()).
			Where(tx.PackageProduct.Number.Eq(p.Number)).
			Updates(map[string]interface{}{"state": TRANSPORT_PACPRO})
		if err != nil {
			_ = tx.Rollback()
			return err
		}

		_, err = tx.PackWarehouse.WithContext(context.Background()).
			Where(tx.PackWarehouse.ProductNumber.Eq(p.Number)).
			Updates(map[string]interface{}{"state": TRANSPORT_PH})
		if err != nil {
			_ = tx.Rollback()
			return err
		}

		_, err = fabric.UploadProcedure(procedure.PID, procedure.PrePID)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	_, err = tx.TransportBatch.WithContext(context.Background()).
		Where(tx.TransportBatch.BatchNumber.Eq(r.BatchNumber)).
		Updates(map[string]interface{}{"state": START_STATE_BATCH_TRANS})
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

func GetTransportBatches(houseNum string) ([]coldchain.TransportBatchInfo, int64, error) {
	q := query.TransportBatch
	pbs, err := q.WithContext(context.Background()).Where(q.TVNumber.Eq(houseNum)).Preload(q.Products).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(pbs)
	records := make([]coldchain.TransportBatchInfo, count)
	for i, pb := range pbs {
		records[i] = coldchain.ToTransportBatchInfo(pb)
	}

	return records, int64(count), nil
}
