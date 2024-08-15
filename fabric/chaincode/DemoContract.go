package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type DemoContract struct {
	contractapi.Contract
}

type DemoData struct {
	Value []byte `json:"value"`
}

func (dd *DemoContract) StoreRequestOperation(ctx contractapi.TransactionContextInterface, requestID string, value []byte) error {
	glog.Info("Blockchain--------Upload------Data-------DAVEX")
	operation := DemoData{Value: value}

	// 将 RequestOperation 实例序列化为 JSON
	operationAsBytes, err := json.Marshal(operation)
	if err != nil {
		return fmt.Errorf("failed to marshal operation: %v", err)
	}

	// 将序列化的数据存储到区块链上，键值为 requestID
	return ctx.GetStub().PutState(requestID, operationAsBytes)
}

func (dd *DemoContract) QueryData(ctx contractapi.TransactionContextInterface, requestID string) (*DemoData, error) {
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
	}

	return &operation, nil
}
