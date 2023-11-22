package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

type Procedure struct {
	PID       string `json:"pid"`
	PrePID    string `json:"pre_pid"`
	CheckCode string `json:"checkcode"`
	PHash     string `json:"p_hash"`
}

// InitLedger adds a base set of assets to the ledger
// new初始化用户数据
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	proc1 := Procedure{
		PID:       "1111",
		PrePID:    "0000",
		CheckCode: "123456",
	}
	proc2 := Procedure{
		PID:       "2222",
		PrePID:    "1111",
		CheckCode: "456789",
	}
	procs := []Procedure{
		proc1, proc2,
	}

	fmt.Println("[InitLedger] UserList:", procs)

	for _, value := range procs {
		userdataJSON, err := json.Marshal(value)
		if err != nil {
			return err
		}
		err = ctx.GetStub().PutState(value.PID, userdataJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// 上传Procedure创建
func (s *SmartContract) UploadProcedure(ctx contractapi.TransactionContextInterface, pid string, pre_pid string) error {
	glog.Info("Blockchain--------UploadProcedure")
	exists, err := s.UserExists(ctx, pid)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", pid)
	}
	proc := Procedure{
		PID:    pid,
		PrePID: pre_pid,
		//CheckCode: checkcode,
	}

	fmt.Println("userdata:  ", proc)
	userdataJSON, err := json.Marshal(proc)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(proc.PID, userdataJSON)
}

// 更新Procedure的checkcode和phash
func (s *SmartContract) UpdateProcedure(ctx contractapi.TransactionContextInterface, pid string, checkcode string, p_hash string) error {
	glog.Info("Blockchain--------UpdateProcedureCheckcode")
	exists, err := s.UserExists(ctx, pid)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the procedure %s does not exist", pid)
	}

	ProcedureDataJSON, err := ctx.GetStub().GetState(pid)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}

	var ProcedureData Procedure
	err = json.Unmarshal(ProcedureDataJSON, &ProcedureData)
	if err != nil {
		fmt.Println("Unmarshal failed")
	}
	//oldProc := UserCarbonData.CarbonTxs
	procdata := Procedure{
		PID:       pid,
		PrePID:    ProcedureData.PrePID,
		CheckCode: checkcode,
		PHash:     p_hash,
	}
	fmt.Println("new Procedure:  ", procdata)

	procdataJSON, err := json.Marshal(procdata)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(pid, procdataJSON)
}

// new查询procedure数据
func (s *SmartContract) QueryProcedureWithPID(ctx contractapi.TransactionContextInterface, pid string) (*Procedure, error) {
	glog.Info("Blockchain--------QueryProcedureWithPID")

	procJSON, err := ctx.GetStub().GetState(pid)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	var proc Procedure
	err = json.Unmarshal(procJSON, &proc)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return nil, err
	} else {
		fmt.Println("proc: ", proc)
	}

	return &proc, nil

}

// UserExists returns true when asset with given ID exists in world state
// old version
// new
func (s *SmartContract) UserExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	fmt.Println("UserExists")

	userJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return userJSON != nil, nil
}
