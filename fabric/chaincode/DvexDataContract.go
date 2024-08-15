package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type DataUploadContract struct {
	contractapi.Contract
}

type DataUploadOperation struct {
	RequestHash []byte `json:"request_hash"`
}

func (duc *DataUploadContract) UploadFileHash(ctx contractapi.TransactionContextInterface, operationID string, requestHash []byte) error {
	//创建实例
	operation := DataUploadOperation{
		RequestHash: requestHash,
	}

	//实例化为json
	operationAsBytes, err := json.Marshal(operation)
	if err != nil {
		return fmt.Errorf("failed to marshal operation: %v", err)
	}

	// 将序列化的数据存储到区块链上，键值为 requestID
	return ctx.GetStub().PutState(operationID, operationAsBytes)
}

// QueryFileHash 根据 requestID 查询文件的哈希值
func (duc *DataUploadContract) QueryFileHash(ctx contractapi.TransactionContextInterface, operationID string) (*DataUploadOperation, error) {
	// 从区块链上获取存储的文件哈希数据
	operationAsBytes, err := ctx.GetStub().GetState(operationID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if operationAsBytes == nil {
		return nil, fmt.Errorf("file hash for request %s does not exist", operationID)
	}

	// 将获取到的 JSON 数据反序列化为 DataUploadOperation 实例
	var operation DataUploadOperation
	err = json.Unmarshal(operationAsBytes, &operation)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation: %v", err)
	}

	return &operation, nil
}
