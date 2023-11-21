package service

import (
	"CN-EU-FSIMS/internal/app/handlers/request"
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/pasture"
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/utils"
	"CN-EU-FSIMS/utils/crypto"
	"context"
	"errors"
	"github.com/golang/glog"
	"time"
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

func AddProcedure(rcp request.ReqCreateProcedure) error {
	ts := time.Now()
	pid, err := generatePID(rcp.Type, ts)
	if err != nil {
		return err
	}

	var prepid string
	if rcp.Type == PASTURE_TYPE {
		prepid = "HEADER"
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
		CompletedTimestamp: ts,
		PrePID:             prepid,
		ICID:               "",
		SubProcedures:      nil,
	}

	// 在数据库中新建Procedure
	err = query.Procedure.WithContext(context.Background()).Create(&p)
	if err != nil {
		return err
	}

	// 调用智能合约新建Procedure

	return nil
}

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
		dataBytes := utils.SerializeStructToBytes(data.(pasture.PastureProcedureData))
		pBytes = append(phBytes, dataBytes...)
	case PACKAGE_TYPE:
		dataBytes := utils.SerializeStructToBytes(data.(pasture.PastureProcedureData))
		pBytes = append(phBytes, dataBytes...)
	case COLDCHAIN_TRANSPORT_TYPE:
		dataBytes := utils.SerializeStructToBytes(data.(pasture.PastureProcedureData))
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

func CommitPastureProcedure(cpp request.CommitPastureProcedure) (string, error) {
	// 转换为pasture data
	data := pasture.PastureProcedureData{
		PM10:   cpp.PM10,
		TSP:    cpp.TSP,
		Stench: cpp.Stench,
	}

	// 获取当前Procedure的前Procedure PID
	p, err := QueryProcedureWithPID(cpp.PID)
	if err != nil {
		return "", err
	}

	cts := time.Now()
	pHeader := ProcedureHeader{
		PID:                cpp.PID,
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

	// 计算新的CheckCode
	ncc := crypto.CalculateSHA256(string(pHash+precc), "1111")

	// 更新当前Procedure
	params := map[string]interface{}{"p_hash": pHash, "check_code": ncc, "completed_timestamp": cts}
	err = UpdateProcedureWithPID(cpp.PID, params)
	if err != nil {
		return "", err
	}

	// 更新链上Procedure的pHash

	return ncc, nil
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
	procedure, err := p.WithContext(context.Background()).Where(p.PrePID.Eq(pid)).Select(p.CheckCode).First()
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
