package warehouse

// import (
// 	"CN-EU-FSIMS/internal/app/models/query"
// 	"CN-EU-FSIMS/utils/crypto"
// 	"context"
// 	"time"

// 	"github.com/golang/glog"
// )

// const (
// 	//入库状态为0
// 	ENTRY_STATE = iota
// 	//运输中状态为1
// 	IN_TRANSIT_STATE
// 	//已出库状态为2
// 	OUT_SUCCESS_STATE

// 	PID_PREFIX = "FSIMSPN"
// )

// func generateProductNumber(ptype string, startTimestamp time.Time, procedure string) string {
// 	time_str := startTimestamp.String()
// 	pnhash := crypto.CalculateSHA256(string(time_str+ptype), procedure)

// 	p_number := PID_PREFIX

// 	return p_number + procedure + pnhash

// }

// func isProcedureCompleted(pid string) bool {
// 	p := query.Procedure
// 	glog.Infoln("pid is ", pid)
// 	proc, err := p.WithContext(context.Background()).Where(p.PID.Eq(pid)).First()
// 	if err != nil {
// 		return false
// 	}
// 	timezero := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
// 	if proc.CompletedTimestamp != timezero {
// 		return true
// 	} else {
// 		return false
// 	}
// }

//func UpdatePastureWarehouse(productNumber string, params map[string]interface{}) error {
//	p := query.PastureWareHouse
//	info, err := p.WithContext(context.Background()).Where(p.ProductNumber.Eq(productNumber)).Updates(params)
//	if err != nil {
//		glog.Errorln(info)
//		return err
//	}
//	glog.Infoln(info)
//	return nil
//}
//func UpdateSlaughterWarehouse(productNumber string, params map[string]interface{}) error {
//	p := query.SlaughterWareHouse
//	info, err := p.WithContext(context.Background()).Where(p.ProductNumber.Eq(productNumber)).Updates(params)
//	if err != nil {
//		glog.Errorln(info)
//		return err
//	}
//	glog.Infoln(info)
//	return nil
//}
