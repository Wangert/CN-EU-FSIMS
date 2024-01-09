package service

import (
	"CN-EU-FSIMS/fabric"
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/coldchain"
	"CN-EU-FSIMS/internal/app/models/pack"
	"CN-EU-FSIMS/internal/app/models/pasture"
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/internal/app/models/slaughter"
	"CN-EU-FSIMS/utils"
	"CN-EU-FSIMS/utils/crypto"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang/glog"
)

const (
	PID_PREFIX = "FSIMSP"

	// PASTURE procedure is 0
	PASTURE_TYPE = iota
	// FATTEN procedure is 1
	FATTEN_TYPE
	// PREMORTEM_MANAGEMENT procedure is 2
	PREMORTEM_MANAGEMENT_TYPE
	// SLAUGHTER procedure is 3
	SLAUGHTER_TYPE
	// PACKAGE procedure is 4
	PACKAGE_TYPE
	// COLDCHAIN_TRANSPORT procedure is 5
	COLDCHAIN_TRANSPORT_TYPE
	// SELL procedure is 6
	SELL_TYPE

	PASTURE_NUMBER              = "0100"
	FATTEN_NUMBER               = "0101"
	PREMORTEM_MANAGEMENT_NUMBER = "0102"
	SLAUGHTER_NUMBER            = "0103"
	PACKAGE_NUMBER              = "0104"
	COLDCHAIN_TRANSPORT_NUMBER  = "0105"
	SELL_NUMBER                 = "0106"
)

type ProcedureHeader struct {
	PID                string
	Type               uint
	StartTimestamp     time.Time
	CompletedTimestamp time.Time
	PrePID             string
}

type NewProcedureParams struct {
	Type        uint
	Operator    string
	PrePID      string
	BatchNumber string
}

func NewProcedure(params *NewProcedureParams, startTime time.Time) (models.Procedure, error) {
	pid, err := generatePID(params.Type, time.Now())
	if err != nil {
		return models.Procedure{}, err
	}

	var prepid string
	if params.Type == PASTURE_TYPE {
		prepid = "HEADER"
	} else {
		prepid = params.PrePID
	}

	p := models.Procedure{
		PID:                pid,
		Type:               params.Type,
		Name:               "",
		PHash:              "",
		CheckCode:          "",
		SerialNumber:       time.Now().UnixNano(),
		Operator:           params.Operator,
		StartTimestamp:     startTime,
		CompletedTimestamp: nil,
		PrePID:             prepid,
		ICID:               "",
		BatchNumber:        &params.BatchNumber,
	}

	return p, nil
}

func AddProcedure(rcp request.ReqCreateProcedure) error {
	ts := time.Now()
	pid, err := generatePID(rcp.Type, ts)
	if err != nil {
		return err
	}

	var prepid string
	if rcp.Type == PASTURE_TYPE {
		prepid = "HEADER"
	} else {
		prepid = rcp.PrePID
	}

	p := models.Procedure{
		PID:                pid,
		Type:               rcp.Type,
		Name:               "",
		PHash:              "",
		CheckCode:          "",
		SerialNumber:       0,
		Operator:           rcp.Operator,
		StartTimestamp:     ts,
		CompletedTimestamp: nil,
		PrePID:             prepid,
		ICID:               "",
	}

	// 在数据库中新建Procedure
	err = query.Procedure.WithContext(context.Background()).Create(&p)
	if err != nil {
		return err
	}

	// 调用智能合约新建Procedure
	_, err = fabric.UploadProcedure(pid, prepid)
	if err != nil {
		return err
	}
	return nil
}

//func AddTransportProcedure(rcp request.TransportStart) error {
//	ts := time.Now()
//	pid, err := generatePID(rcp.Type, ts)
//	if err != nil {
//		return err
//	}
//
//	p := models.Procedure{
//		PID:                pid,
//		Type:               rcp.Type,
//		Name:               "",
//		PHash:              "",
//		CheckCode:          "",
//		SerialNumber:       0,
//		Operator:           rcp.Operator,
//		StartTimestamp:     ts,
//		CompletedTimestamp: ts,
//		PrePID:             rcp.PrePID,
//		ICID:               "",
//	}
//	d := coldchain.TransportProcedureData{
//		TID:                pid,
//		ProductNumber:      rcp.ProductNumber,
//		TVNumber:           rcp.CarNumber,
//		Operator:           rcp.Operator,
//		Temperature:        rcp.Temperature,
//		Source:             rcp.Source,
//		Destination:        rcp.Destination,
//		Humidity:           rcp.Humidity,
//		LoadingTime:        rcp.LoadingTime,
//		UnloadingTime:      "",
//		StartTimestamp:     ts,
//		CompletedTimestamp: ts,
//	}
//	tx := query.Q.Begin()
//	defer func() {
//		if recover() != nil || err != nil {
//			_ = tx.Rollback()
//		} else {
//			err = tx.Commit()
//			if err != nil {
//				panic(err)
//			}
//
//			glog.Infoln("add transport procedure successful!")
//		}
//	}()
//	cts := time.Now()
//	// 在数据库中新建Procedure
//	err = tx.Procedure.WithContext(context.Background()).Create(&p)
//	err = tx.TransportProcedureData.WithContext(context.Background()).Create(&d)
//	params := map[string]interface{}{"out_timestamp": cts, "state": 2}
//	info, err := tx.PackWarehouse.WithContext(context.Background()).Where(tx.PackWarehouse.ProductNumber.Eq(rcp.ProductNumber)).Updates(params)
//	glog.Info("info", info)
//	// 调用智能合约新建Procedure
//	_, err = fabric.UploadProcedure(pid, rcp.PrePID)
//	if err != nil {
//		return err
//	}
//	return nil
//
//}

// 该函数未完成，仅仅支持对Pasture过程的PHash进行计算!!!!!
func generatePHash(ph *ProcedureHeader, data interface{}) (string, error) {

	var pBytes []byte
	phBytes := utils.SerializeStructToBytes(*ph)

	switch ph.Type {
	case PASTURE_TYPE:
		dataBytes := utils.SerializeStructToBytes(data.(pasture.PastureProcedureData))
		pBytes = append(phBytes, dataBytes...)
	case FATTEN_TYPE:
		dataBytes := utils.SerializeStructToBytes(data.(pasture.PastureProcedureData))
		pBytes = append(phBytes, dataBytes...)
	case PREMORTEM_MANAGEMENT_TYPE:
		dataBytes := utils.SerializeStructToBytes(data.(pasture.PastureProcedureData))
		pBytes = append(phBytes, dataBytes...)
	case SLAUGHTER_TYPE:
		dataBytes := utils.SerializeStructToBytes(data.(slaughter.SlaughterProcedureData))
		pBytes = append(phBytes, dataBytes...)
	case PACKAGE_TYPE:
		dataBytes := utils.SerializeStructToBytes(data.(pack.PackProcedureData))
		pBytes = append(phBytes, dataBytes...)
	case COLDCHAIN_TRANSPORT_TYPE:
		dataBytes := utils.SerializeStructToBytes(data.(coldchain.TransportProcedureData))
		pBytes = append(phBytes, dataBytes...)
	case SELL_TYPE:
		dataBytes := utils.SerializeStructToBytes(data.(pasture.PastureProcedureData))
		pBytes = append(phBytes, dataBytes...)
	default:
		return "", errors.New("type is not exist")
	}

	pHash := crypto.CalculateSHA256(string(pBytes), "1000")
	return pHash, nil
}

// PastureProcedure提交
func CommitPastureProcedure(cpp request.CommitPastureProcedure) (string, error) {
	// 转换为pasture data
	data := pasture.PastureProcedureData{
		PM10:   cpp.PM10,
		TSP:    cpp.TSP,
		Stench: cpp.Stench,
	}

	//后续需要存储data到数据库

	pid := cpp.PID
	ncc, err := BasicCommitProcedure(pid, data)
	return ncc, err
}

// SlaughterProcedure提交
func CommitSlaughterProcedure(csp request.CommitSlaughterProcedure) (string, error) {
	// 转换为pasture data
	data := slaughter.SlaughterProcedureData{
		EnvirTemperature:      csp.EnvirTemperature,
		EnvirLighting:         csp.EnvirLighting,
		ShockVoltage:          csp.ShockVoltage,
		BleedingTime:          csp.BleedingTime,
		EleShockTime:          csp.EleShockTime,
		KnifeDisinfectionTime: csp.KnifeDisinfectionTime,
	}

	//后续需要存储data到数据库

	pid := csp.PID
	ncc, err := BasicCommitProcedure(pid, data)
	return ncc, err
}

// PackProcedure提交
func CommitPackProcedure(cpp request.CommitPackProcedure) (string, error) {
	data := pack.PackProcedureData{
		PackType:        cpp.PackType,
		PackTemperature: cpp.PackTemperature,
	}

	pid := cpp.PID
	ncc, err := BasicCommitProcedure(pid, data)
	return ncc, err
}

// TransportProcedure提交
//func TransportEnd(rte request.TransportEnd) (string, error) {
//	// 转换为pasture data
//	cts := time.Now()
//	pid := rte.PID
//	p := query.TransportProcedureData
//	params := map[string]interface{}{"unloading_time": rte.UnloadingTime, "completed_timestamp": cts}
//	info, err := p.WithContext(context.Background()).Where(p.TID.Eq(pid)).Updates(params)
//	if err != nil {
//		glog.Errorln(info)
//		return "", err
//	}
//	q := query.TransportProcedureData
//	data, err := q.WithContext(context.Background()).Where(q.TID.Eq(rte.PID)).First()
//	ncc, err := BasicCommitProcedure(pid, *data)
//	return ncc, err
//}

// 计算checkcode并更新procedure

func BasicCommitProcedureWithTx(tx *query.QueryTx, pid string, endTime *time.Time, data interface{}) (string, string, error) {
	// 获取当前Procedure的前Procedure PID
	p, err := QueryProcedureWithPID(pid)
	if err != nil {
		return "", "", err
	}

	pHeader := ProcedureHeader{
		PID:                pid,
		Type:               p.Type,
		StartTimestamp:     p.StartTimestamp,
		CompletedTimestamp: *endTime,
		PrePID:             p.PrePID,
	}

	// 计算PHash
	pHash, err := generatePHash(&pHeader, data)
	if err != nil {
		return "", "", err
	}

	// 计算CheckCode
	// 获取PrePID的CheckCode
	var precc string
	if p.PrePID == "HEADER" {
		precc = "HSGAFSGSGAHGS8HS65GSAJ"
	} else {
		precc, err = QueryProcedureCheckcode(p.PrePID)
		if err != nil {
			return "", "", err
		}
	}
	fmt.Println("pHash:", pHash)
	fmt.Println("precc", precc)
	// 计算新的CheckCode
	ncc := crypto.CalculateSHA256(string(pHash+precc), "1111")

	// 更新当前Procedure
	params := map[string]interface{}{"p_hash": pHash, "check_code": ncc, "completed_timestamp": endTime}
	err = UpdateProcedureWithPIDWithTx(tx, pid, params)
	if err != nil {
		return "", "", err
	}

	//// 更新链上Procedure的pHash
	//_, err = fabric.UpdateProcedure(pid, pHash)
	return ncc, pHash, nil
}

// 计算checkcode并更新procedure
func BasicCommitProcedure(pid string, data interface{}) (string, error) {
	// 获取当前Procedure的前Procedure PID
	p, err := QueryProcedureWithPID(pid)
	if err != nil {
		return "", err
	}

	cts := time.Now()
	pHeader := ProcedureHeader{
		PID:                pid,
		Type:               p.Type,
		StartTimestamp:     p.StartTimestamp,
		CompletedTimestamp: cts,
		PrePID:             p.PrePID,
	}

	// 计算PHash
	pHash, err := generatePHash(&pHeader, data)
	if err != nil {
		return "", err
	}

	// 计算CheckCode
	// 获取PrePID的CheckCode
	var precc string
	if p.PrePID == "HEADER" {
		precc = "HSGAFSGSGAHGS8HS65GSAJ"
	} else {
		precc, err = QueryProcedureCheckcode(p.PrePID)
		if err != nil {
			return "", err
		}
	}
	fmt.Println("pHash:", pHash)
	fmt.Println("precc", precc)
	// 计算新的CheckCode
	ncc := crypto.CalculateSHA256(string(pHash+precc), "1111")

	// 更新当前Procedure
	params := map[string]interface{}{"p_hash": pHash, "check_code": ncc, "completed_timestamp": cts}
	err = UpdateProcedureWithPID(pid, params)
	if err != nil {
		return "", err
	}

	// 更新链上Procedure的pHash
	_, err = fabric.UpdateProcedure(pid, pHash)
	return ncc, nil
}

func UpdateProcedureWithPIDWithTx(tx *query.QueryTx, pid string, params map[string]interface{}) error {
	p := tx.Procedure
	info, err := p.WithContext(context.Background()).Where(p.PID.Eq(pid)).Updates(params)
	if err != nil {
		glog.Errorln(info)
		return err
	}
	glog.Infoln(info)
	return nil
}

func UpdateProcedureWithPID(pid string, params map[string]interface{}) error {
	p := query.Procedure
	info, err := p.WithContext(context.Background()).Where(p.PID.Eq(pid)).Updates(params)
	if err != nil {
		glog.Errorln(info)
		return err
	}
	glog.Infoln(info)
	return nil
}

func QueryProcedureCheckcode(pid string) (string, error) {
	p := query.Procedure
	procedure, err := p.WithContext(context.Background()).Where(p.PID.Eq(pid)).Select(p.CheckCode).First()
	if err != nil {
		return "", err
	}
	return procedure.CheckCode, nil
}

func QueryProcedureWithPID(pid string) (*models.Procedure, error) {
	p := query.Procedure
	procedures, err := p.WithContext(context.Background()).Where(p.PID.Eq(pid)).First()
	if err != nil {
		return nil, err
	}
	return procedures, nil
}

func generatePID(ptype uint, startTimestamp time.Time) (string, error) {
	tsHash := crypto.CalculateSHA256(startTimestamp.String(), "pid")

	pid := PID_PREFIX

	switch ptype {
	case PASTURE_TYPE:
		pid += PASTURE_NUMBER
	case FATTEN_TYPE:
		pid += FATTEN_NUMBER
	case PREMORTEM_MANAGEMENT_TYPE:
		pid += PREMORTEM_MANAGEMENT_NUMBER
	case SLAUGHTER_TYPE:
		pid += SLAUGHTER_NUMBER
	case PACKAGE_TYPE:
		pid += PACKAGE_NUMBER
	case COLDCHAIN_TRANSPORT_TYPE:
		pid += COLDCHAIN_TRANSPORT_NUMBER
	case SELL_TYPE:
		pid += SELL_NUMBER
	default:
		return "", errors.New("generate pid error")
	}

	return pid + tsHash, nil
}

func QueryIndustrialChain(checkcode string) ([]*models.Procedure, error) {

	var procedure *models.Procedure
	var err error
	var procedures []*models.Procedure
	for {
		p := query.Procedure
		procedure, err = p.WithContext(context.Background()).Where(p.CheckCode.Eq(checkcode)).First()
		if err != nil {
			return nil, err
		}
		if procedure == nil {
			return nil, errors.New("No checkcode!")
		}
		procedures = append(procedures, procedure)
		if procedure.PrePID == "HEADER" {
			break
		}
		pre_pid := procedure.PrePID
		pre_procedure, err := p.WithContext(context.Background()).Where(p.PID.Eq(pre_pid)).First()
		if err != nil {
			return nil, err
		}
		checkcode = pre_procedure.CheckCode
	}
	return procedures, nil
}

func VerifyWithCheckcode(checkcode string) (string, error) {
	p := query.Procedure
	procedure, err := p.WithContext(context.Background()).Where(p.CheckCode.Eq(checkcode)).First()
	if err != nil {
		return "", err
	}
	glog.Info("pid:", procedure.PID)

	res, err := fabric.VerifyOnChainWithCheckcode(procedure.PID, checkcode)

	return res, err
}
