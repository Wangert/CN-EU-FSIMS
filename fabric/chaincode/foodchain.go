package chaincode

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

type CarbonChainNode struct {
	Carboncoin    string `json:"Carboncoin"`
	Carbonquota   string `json:"Carbonquota"`
	UploadCompany string `json:"UploadCompany"`
	FileHash      string `json:"FileHash"`
}

type CarbonChainData struct {
	ID          string            `json:"ID"`
	Signature   string            `json:"Signature"`
	CarbonChain []CarbonChainNode `json:"CarbonChain"`
}

type CarbonTx struct {
	ID          string `json:"ID"`
	SenderID    string `json:"SenderID"`
	ReceiverID  string `json:"ReceiverID"`
	CarbonQuota string `json:"CarbonQuota"`
	UploadTime  string `json:"submittime"`
	State       string `json:"State"`
}

type User struct {
	ID          string              `json:"ID"`
	Name        string              `json:"Name"`
	CarbonQuota string              `json:"CarbonQuota"`
	CarbonCoin  string              `json:"CarbonCoin"`
	CarbonTxs   map[string]CarbonTx `json:"CarbonTxs"`
}

// InitLedger adds a base set of assets to the ledger
//new初始化用户数据
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	// carbonchain1 := []CarbonChainNode{
	// 	{Carboncoin: "100", Carbonquota: "10", UploadCompany: "蒙牛", FileHash: "aaa"},
	// }
	// carbonchain2 := []CarbonChainNode{
	// 	{Carboncoin: "100", Carbonquota: "8", UploadCompany: "蒙牛", FileHash: "aaa"},
	// }

	user1 := User{
		ID:          "user1",
		Name:        "user1",
		CarbonQuota: "100",
		CarbonCoin:  "10",
		CarbonTxs:   map[string]CarbonTx{},
	}

	user2 := User{
		ID:          "user2",
		Name:        "user2",
		CarbonQuota: "100",
		CarbonCoin:  "10",
		CarbonTxs:   map[string]CarbonTx{},
	}

	users := []User{
		user1, user2,
	}

	fmt.Println("[InitLedger] UserList:", users)

	for _, userdata := range users {
		userdataJSON, err := json.Marshal(userdata)
		if err != nil {
			return err
		}
		err = ctx.GetStub().PutState(userdata.ID, userdataJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	// carbonchaindatas := []CarbonChainData{
	// 	{ID: "asset1", Signature: "wyz", CarbonChain: carbonchain1},
	// 	{ID: "asset2", Signature: "wjt", CarbonChain: carbonchain2},
	// }

	// fmt.Println("carbonchaindatas", carbonchaindatas)

	// for _, carbonchaindata := range carbonchaindatas {
	// 	carbonchaindataJSON, err := json.Marshal(carbonchaindata)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	err = ctx.GetStub().PutState(carbonchaindata.ID, carbonchaindataJSON)
	// 	if err != nil {
	// 		return fmt.Errorf("failed to put to world state. %v", err)
	// 	}
	// }

	return nil
}

//new上传用户数据
func (s *SmartContract) UploadUserCarbondata(ctx contractapi.TransactionContextInterface, id string, name string, carbonquota string, carboncoin string) error {
	fmt.Println("UploadUserCarbonData")
	exists, err := s.UserExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	userdata := User{
		ID:          id,
		Name:        name,
		CarbonQuota: carbonquota,
		CarbonCoin:  carboncoin,
		CarbonTxs:   map[string]CarbonTx{},
	}
	fmt.Println("userdata:  ", userdata)
	userdataJSON, err := json.Marshal(userdata)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, userdataJSON)
}

//new更新用户数据
func (s *SmartContract) UpdateUserCarbondata(ctx contractapi.TransactionContextInterface, id string, name string, carbonquota string, carboncoin string) error {
	fmt.Println("UpdateUserCarbondata")
	exists, err := s.UserExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the carbonchain %s does not exist", id)
	}

	UserCarbonDataJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}

	var UserCarbonData User
	err = json.Unmarshal(UserCarbonDataJSON, &UserCarbonData)
	if err != nil {
		fmt.Println("Unmarshal failed")
	}
	oldtx := UserCarbonData.CarbonTxs
	usercarbondata := User{
		ID:          id,
		Name:        name,
		CarbonQuota: carbonquota,
		CarbonCoin:  carboncoin,
		CarbonTxs:   oldtx,
	}
	fmt.Println("new carbonchaindata:  ", usercarbondata)

	usercarbondataJSON, err := json.Marshal(usercarbondata)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, usercarbondataJSON)
}

//new初试化用户交易（交易提交）
func (s *SmartContract) InitiateCarbonTx(ctx contractapi.TransactionContextInterface, txID string, sourceId string, targetId string, carbonquota string, timestamp string) error {

	exists1, err := s.UserExists(ctx, sourceId)
	if err != nil {
		return err
	}
	if !exists1 {
		return fmt.Errorf("Source ID %s does not exist", sourceId)
	}

	exists2, err := s.UserExists(ctx, targetId)
	if err != nil {
		return err
	}
	if !exists2 {
		return fmt.Errorf("Target ID %s does not exist", targetId)
	}

	newTx := CarbonTx{
		ID:          txID,
		SenderID:    sourceId,
		ReceiverID:  targetId,
		CarbonQuota: carbonquota,
		UploadTime:  timestamp,
		State:       "2",
	}

	// update source user
	sourceJSON, err := ctx.GetStub().GetState(sourceId)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}

	var source User
	err = json.Unmarshal(sourceJSON, &source)
	if err != nil {
		fmt.Println("Unmarshal failed")
	}

	if _, ok := source.CarbonTxs[txID]; !ok {
		source.CarbonTxs[txID] = newTx

		newSourceJSON, err := json.Marshal(source)
		if err != nil {
			return err
		}
		err = ctx.GetStub().PutState(source.ID, newSourceJSON)
		if err != nil {
			return err
		}
	}

	// update target user
	targetJSON, err := ctx.GetStub().GetState(targetId)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}

	var target User
	err = json.Unmarshal(targetJSON, &target)
	if err != nil {
		fmt.Println("Unmarshal failed")
	}

	if _, ok := target.CarbonTxs[txID]; !ok {
		target.CarbonTxs[txID] = newTx

		newTargetJSON, err := json.Marshal(target)
		if err != nil {
			return err
		}
		err = ctx.GetStub().PutState(target.ID, newTargetJSON)
		if err != nil {
			return err
		}
	}

	return nil
}

//new初始化实时交易(实时交易提交)
func (s *SmartContract) InitiateOntimeCarbonTx(ctx contractapi.TransactionContextInterface, txID string, sourceId string, carbonquota string, timestamp string) error {

	exists1, err := s.UserExists(ctx, sourceId)
	if err != nil {
		return err
	}
	if !exists1 {
		return fmt.Errorf("Source ID %s does not exist", sourceId)
	}

	newTx := CarbonTx{
		ID:          txID,
		SenderID:    sourceId,
		ReceiverID:  "",
		CarbonQuota: carbonquota,
		UploadTime:  timestamp,
		State:       "2",
	}

	// update source user
	sourceJSON, err := ctx.GetStub().GetState(sourceId)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}

	var source User
	err = json.Unmarshal(sourceJSON, &source)
	if err != nil {
		fmt.Println("Unmarshal failed")
	}

	if _, ok := source.CarbonTxs[txID]; !ok {
		source.CarbonTxs[txID] = newTx

		newSourceJSON, err := json.Marshal(source)
		if err != nil {
			return err
		}
		err = ctx.GetStub().PutState(source.ID, newSourceJSON)
		if err != nil {
			return err
		}
	}
	return nil
}

//new用户交易确认
func (s *SmartContract) ConfirmCarbonTx(ctx contractapi.TransactionContextInterface, sourceId, targetId, txId string) error {
	// target user
	targetJSON, err := ctx.GetStub().GetState(targetId)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}

	var target User
	err = json.Unmarshal(targetJSON, &target)
	if err != nil {
		fmt.Println("Unmarshal failed")
	}

	sourceJSON, err := ctx.GetStub().GetState(sourceId)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}

	var source User
	err = json.Unmarshal(sourceJSON, &source)
	if err != nil {
		fmt.Println("Unmarshal failed")
	}

	var tx CarbonTx
	if v, ok := target.CarbonTxs[txId]; ok {
		tx = v
	} else {
		return fmt.Errorf("tx id is not exist!")
	}

	scq, err := strconv.Atoi(source.CarbonQuota)
	if err != nil {
		return err
	}

	tcq, err := strconv.Atoi(target.CarbonQuota)
	if err != nil {
		return err
	}

	cq, err := strconv.Atoi(tx.CarbonQuota)
	if err != nil {
		return err
	}

	scq -= cq
	tcq += cq

	scqStr := strconv.Itoa(scq)
	tcqStr := strconv.Itoa(tcq)

	source.CarbonQuota = scqStr
	target.CarbonQuota = tcqStr

	tx.State = "3"

	source.CarbonTxs[txId] = tx
	target.CarbonTxs[txId] = tx

	// 补充计算碳币

	newSourceJSON, err := json.Marshal(source)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(source.ID, newSourceJSON)
	if err != nil {
		return err
	}

	newTargetJSON, err := json.Marshal(target)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(target.ID, newTargetJSON)
	if err != nil {
		return err
	}

	return nil
}

//new实时交易确认
func (s *SmartContract) OntimeConfirmCarbonTx(ctx contractapi.TransactionContextInterface, sourceId, targetId, txId string) error {
	// target user
	targetJSON, err := ctx.GetStub().GetState(targetId)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}

	var target User
	err = json.Unmarshal(targetJSON, &target)
	if err != nil {
		fmt.Println("Unmarshal failed")
	}

	sourceJSON, err := ctx.GetStub().GetState(sourceId)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}

	var source User
	err = json.Unmarshal(sourceJSON, &source)
	if err != nil {
		fmt.Println("Unmarshal failed")
	}

	var tx CarbonTx
	if v, ok := source.CarbonTxs[txId]; ok {
		tx = v
	} else {
		return fmt.Errorf("tx id is not exist!")
	}
	//以下内容还未在chaincode中部署测试
	scq, err := strconv.ParseFloat(source.CarbonQuota, 32)
	if err != nil {
		return err
	}
	scc, err := strconv.Atoi(source.CarbonCoin)
	if err != nil {
		return err
	}
	tcq, err := strconv.ParseFloat(target.CarbonQuota, 32)
	if err != nil {
		return err
	}
	tcc, err := strconv.Atoi(target.CarbonCoin)
	if err != nil {
		return err
	}

	cq, err := strconv.ParseFloat(tx.CarbonQuota, 32)
	if err != nil {
		return err
	}

	scq -= cq
	tcq += cq
	scc = scc + 5
	tcc = tcc + 5

	scqStr := strconv.FormatFloat(scq, 'f', 2, 32)
	tcqStr := strconv.FormatFloat(tcq, 'f', 2, 32)
	sccStr := strconv.Itoa(scc)
	tccStr := strconv.Itoa(tcc)

	source.CarbonQuota = scqStr
	target.CarbonQuota = tcqStr
	source.CarbonCoin = sccStr
	target.CarbonCoin = tccStr

	tx.ReceiverID = targetId
	tx.State = "3"

	source.CarbonTxs[txId] = tx
	target.CarbonTxs[txId] = tx

	newSourceJSON, err := json.Marshal(source)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(source.ID, newSourceJSON)
	if err != nil {
		return err
	}

	newTargetJSON, err := json.Marshal(target)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(target.ID, newTargetJSON)
	if err != nil {
		return err
	}

	return nil
}

//new查询用户数据
func (s *SmartContract) QueryUserCarbonData(ctx contractapi.TransactionContextInterface, id string) (*User, error) {
	fmt.Println("QueryUserCarbonData")

	UserCarbonDataJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	var UserCarbonData User
	err = json.Unmarshal(UserCarbonDataJSON, &UserCarbonData)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return nil, err
	} else {
		fmt.Println("UserCarbonData: ", UserCarbonData)
	}

	return &UserCarbonData, nil

}

//new所有数据
func (s *SmartContract) GetAllUserCarbonData(ctx contractapi.TransactionContextInterface) ([]*User, error) {
	fmt.Println("GetAllUserCarbonData")
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var usercarbondatas []*User
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var usercarbondata User
		err = json.Unmarshal(queryResponse.Value, &usercarbondata)
		if err != nil {
			return nil, err
		}
		usercarbondatas = append(usercarbondatas, &usercarbondata)
	}

	return usercarbondatas, nil
}

//old version
func (s *SmartContract) UploadCarbondata(ctx contractapi.TransactionContextInterface, id string, signature string, carboncoin string, carbonquota string, filehash string, uploadcompany string) error {
	fmt.Println("UploadCarbonData")
	exists, err := s.UserExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	carbonchain := []CarbonChainNode{
		{Carboncoin: carboncoin, Carbonquota: carbonquota, UploadCompany: uploadcompany, FileHash: filehash},
	}
	carbonchaindata := CarbonChainData{
		ID:          id,
		Signature:   signature,
		CarbonChain: carbonchain,
	}
	fmt.Println("carbonchaindata:  ", carbonchaindata)
	carbonchaindataJSON, err := json.Marshal(carbonchaindata)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, carbonchaindataJSON)
}

//old version
func (s *SmartContract) UpdateCarbondata(ctx contractapi.TransactionContextInterface, id string, signature string, carboncoin string, carbonquota string, filehash string, uploadcompany string) error {
	fmt.Println("UpdateCarbondata")
	exists, err := s.UserExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the carbonchain %s does not exist", id)
	}

	CarbonDataJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}

	var CarbonData CarbonChainData
	err = json.Unmarshal(CarbonDataJSON, &CarbonData)
	if err != nil {
		fmt.Println("Unmarshal failed")
	}

	carbonchainnode := CarbonChainNode{
		//Foodchainprocess: foodchainprocess,
		Carboncoin:    carboncoin,
		Carbonquota:   carbonquota,
		UploadCompany: uploadcompany,
		FileHash:      filehash,
	}

	oldchain := CarbonData.CarbonChain

	newchain := append(oldchain, carbonchainnode)
	fmt.Println("newchain: ", newchain)
	carbonchaindata := CarbonChainData{
		ID:          id,
		Signature:   signature,
		CarbonChain: newchain,
	}
	fmt.Println("new carbonchaindata:  ", carbonchaindata)

	carbonchaindataJSON, err := json.Marshal(carbonchaindata)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, carbonchaindataJSON)
}

//old version查询数据
func (s *SmartContract) QueryCarbonData(ctx contractapi.TransactionContextInterface, id string) (*CarbonChainData, error) {
	fmt.Println("QueryCarbonData")

	CarbonDataJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	var CarbonData CarbonChainData
	err = json.Unmarshal(CarbonDataJSON, &CarbonData)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return nil, err
	} else {
		fmt.Println("CarbonData: ", CarbonData)
	}

	return &CarbonData, nil

}

// UserExists returns true when asset with given ID exists in world state
//old version
//new
func (s *SmartContract) UserExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	fmt.Println("UserExists")

	userJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return userJSON != nil, nil
}

//old version得到所有数据
func (s *SmartContract) GetAllCarbonChainData(ctx contractapi.TransactionContextInterface) ([]*CarbonChainData, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	fmt.Println("GetAllCarbonChainData")
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var carbonchaindatas []*CarbonChainData
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var carbonchaindata CarbonChainData
		err = json.Unmarshal(queryResponse.Value, &carbonchaindata)
		if err != nil {
			return nil, err
		}
		carbonchaindatas = append(carbonchaindatas, &carbonchaindata)
	}

	return carbonchaindatas, nil
}
