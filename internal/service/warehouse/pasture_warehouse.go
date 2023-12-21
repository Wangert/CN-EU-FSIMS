package warehouse

//func PastureInWarehouse(rcp request.ReqInPastureWarehouse) error {
//	ts := time.Now()
//	p_number := generateProductNumber(rcp.ProductType, ts, "pasture")
//	glog.Infoln("pid is", rcp.ProductPID)
//	if !isProcedureCompleted(rcp.ProductPID) {
//		return errors.New("procedure has not completed!")
//	}
//	specificTime := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
//	p := warehouse.PastureWareHouse{
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
//	err := query.PastureWareHouse.WithContext(context.Background()).Create(&p)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func SendToSlaughter(rcp request.ReqSendToSlaughter) error {
//	productNumber := rcp.ProductNumber
//	params := map[string]interface{}{"destination": rcp.Destination, "out_operator": rcp.OutOperator, "state": IN_TRANSIT_STATE}
//	err := UpdatePastureWarehouse(productNumber, params)
//	if err != nil {
//		return err
//	}
//	return nil
//}
