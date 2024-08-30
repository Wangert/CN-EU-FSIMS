package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type DemoData struct {
	RequestID string `json:"request_id"`
	Value     string `json:"value"`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	dd1 := DemoData{
		RequestID: "0000",
		Value:     "Hello World",
	}

	dd2 := DemoData{
		RequestID: "1111",
		Value:     "Hello World Aloong",
	}

	dds := []DemoData{dd1, dd2}

	fmt.Println("[InitLedger] DemoData: ", dds)
	for _, value := range dds {
		userdataJSON, err := json.Marshal(value)
		if err != nil {
			return err
		}
		err = ctx.GetStub().PutState(value.RequestID, userdataJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}
	return nil
}

func (dd *SmartContract) StoreRequestOperation(ctx contractapi.TransactionContextInterface, requestID string, value string) error {
	glog.Info("Blockchain--------Upload------Data-------DAVEX")
	exists, err := dd.RequestDataExists(ctx, requestID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", requestID)
	}
	operation := DemoData{
		RequestID: requestID,
		Value:     value,
	}

	// 将 RequestOperation 实例序列化为 JSON
	operationAsBytes, err := json.Marshal(operation)
	if err != nil {
		return fmt.Errorf("failed to marshal operation: %v", err)
	}

	// 将序列化的数据存储到区块链上，键值为 requestID
	return ctx.GetStub().PutState(requestID, operationAsBytes)
}

func (dd *SmartContract) QueryData(ctx contractapi.TransactionContextInterface, requestID string) (*DemoData, error) {
	// 从区块链上获取存储的请求操作数据
	glog.Info("Blockchain--------QUERY------Data-------DAVEX")
	operationAsBytes, err := ctx.GetStub().GetState(requestID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if operationAsBytes == nil {
		return nil, fmt.Errorf("request operation %s does not exist", requestID)
	}

	// 将获取到的 JSON 数据反序列化为 RequestOperation 实例
	var operation DemoData
	err = json.Unmarshal(operationAsBytes, &operation)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation: %v", err)
	} else {
		fmt.Println("Concrete Data: ", operation)
	}

	return &operation, nil
}

func (dd *SmartContract) RequestDataExists(ctx contractapi.TransactionContextInterface, requestID string) (bool, error) {
	fmt.Println("RequestID exists")
	ddJSON, err := ctx.GetStub().GetState(requestID)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return ddJSON != nil, nil
}

func (dd *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*DemoData, error) {
	var assets []*DemoData
	// 获取一个迭代器来遍历世界状态中的所有键值对
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to get assets: %v", err)
	}
	defer resultsIterator.Close()

	// 遍历迭代器中的每一个键值对
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		// 解码资产信息
		var asset DemoData
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}
	return assets, nil
}

func main() {
	ddChainCode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %v\n", err)
	}
	if err := ddChainCode.Start(); err != nil {
		fmt.Printf("Error starting new Smart Contract: %v\n", err)
	}
}
