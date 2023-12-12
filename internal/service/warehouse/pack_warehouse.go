package warehouse

import (
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/internal/app/models/warehouse"
	"context"
	"errors"
	"time"

	"github.com/golang/glog"
)

func PackReceived(rcp request.PackReceive) error {
	tx := query.Q.Begin()
	var err error
	defer func() {
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				panic(err)
			}

			glog.Infoln("receive slaughter product successful!")
		}
	}()
	cts := time.Now()

	p := warehouse.PackReceive{
		ProductNumber:    rcp.ProductNumber,
		ProductPID:       rcp.ProductPID,
		ProductType:      rcp.ProductType,
		Operator:         rcp.Operator,
		ReceiveTimestamp: cts,
	}

	params := map[string]interface{}{"out_timestamp": cts, "state": OUT_SUCCESS_STATE}
	err = tx.PackReceive.WithContext(context.Background()).Create(&p)
	info, err := tx.SlaughterWareHouse.WithContext(context.Background()).Where(tx.SlaughterWareHouse.ProductNumber.Eq(rcp.ProductNumber)).Updates(params)
	glog.Infoln(info)
	return nil
}

func PackInWarehouse(rcp request.ReqInPackWarehouse) error {
	ts := time.Now()
	p_number := generateProductNumber(rcp.ProductType, ts, "pack")
	if !isProcedureCompleted(rcp.ProductPID) {
		return errors.New("procedure has not been completed!")
	}
	specificTime := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	p := warehouse.PackWareHouse{
		ProductNumber: p_number,
		ProductPID:    rcp.ProductPID,
		ProductType:   rcp.ProductType,
		State:         ENTRY_STATE,
		InOperator:    rcp.InOperator,
		InTimestamp:   ts,
		OutTimestamp:  specificTime,
	}
	err := query.PackWareHouse.WithContext(context.Background()).Create(&p)
	if err != nil {
		return err
	}
	return nil
}
