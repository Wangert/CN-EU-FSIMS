package warehouse

//func SlaughterInWarehouse(rcp request.ReqInSlaughterWarehouse) error {
//	ts := time.Now()
//	p_number := generateProductNumber(rcp.ProductType, ts, "slaughter")
//	if !isProcedureCompleted(rcp.ProductPID) {
//		return errors.New("procedure has not been completed!")
//	}
//	specificTime := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
//	p := warehouse.SlaughterWareHouse{
//		ProductNumber: p_number,
//		ProductPID:    rcp.ProductPID,
//		ProductType:   rcp.ProductType,
//		State:         ENTRY_STATE,
//		InOperator:    rcp.InOperator,
//		OutOperator:   "",
//		Destination:   "",
//		InTimestamp:   ts,
//		OutTimestamp:  specificTime,
//	}
//	err := query.SlaughterWareHouse.WithContext(context.Background()).Create(&p)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func SendToPack(rcp request.ReqSendToPack) error {
//	productNumber := rcp.ProductNumber
//	params := map[string]interface{}{"destination": rcp.Destination, "out_operator": rcp.OutOperator, "state": IN_TRANSIT_STATE}
//	err := UpdateSlaughterWarehouse(productNumber, params)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func SlaughterOutWarehouse(productNumber string) error {
//	cts := time.Now()
//	params := map[string]interface{}{"completed_timestamp": cts, "state": OUT_SUCCESS_STATE}
//	err := UpdateSlaughterWarehouse(productNumber, params)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func SlaughterReceived(rcp request.SlaughterReceive) error {
//	tx := query.Q.Begin()
//	var err error
//	defer func() {
//		if recover() != nil || err != nil {
//			_ = tx.Rollback()
//		} else {
//			err = tx.Commit()
//			if err != nil {
//				panic(err)
//			}
//
//			glog.Infoln("receive pasture product successful!")
//		}
//	}()
//	cts := time.Now()
//
//	p := warehouse.SlaughterReceive{
//		ProductNumber:    rcp.ProductNumber,
//		ProductPID:       rcp.ProductPID,
//		ProductType:      rcp.ProductType,
//		Operator:         rcp.Operator,
//		ReceiveTimestamp: cts,
//	}
//
//	params := map[string]interface{}{"out_timestamp": cts, "state": OUT_SUCCESS_STATE}
//	err = tx.SlaughterReceive.WithContext(context.Background()).Create(&p)
//	info, err := tx.PastureWareHouse.WithContext(context.Background()).Where(tx.PastureWareHouse.ProductNumber.Eq(rcp.ProductNumber)).Updates(params)
//	glog.Infoln(info)
//	return nil
//}
