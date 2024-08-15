package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type DirectorySmartContract struct {
	contractapi.Contract
}

// 请求操作上链
type RequestOperation struct {
	RequestHash    []byte `json:"request_hash"`
	TimeStamp      string `json:"time_stamp"`
	AgentID        string `json:"agent_id"`
	CenterID       string `json:"center_id"`
	RequestType    string `json:"request_type"`
	ResponseStatus string `json:"response_status"`
	ResponseHash   []byte `json:"response_hash"`
}

// StoreRequestOperation 将请求操作存储到区块链上
func (dsc *DirectorySmartContract) StoreRequestOperation(ctx contractapi.TransactionContextInterface, requestID string, requestHash []byte, timestamp string, agentID string, centerID string, requestType string, responseStatus string, responseHash []byte) error {
	// 创建 RequestOperation 实例
	operation := RequestOperation{
		RequestHash:    requestHash,
		TimeStamp:      timestamp,
		AgentID:        agentID,
		CenterID:       centerID,
		RequestType:    requestType,
		ResponseStatus: responseStatus,
		ResponseHash:   responseHash,
	}

	// 将 RequestOperation 实例序列化为 JSON
	operationAsBytes, err := json.Marshal(operation)
	if err != nil {
		return fmt.Errorf("failed to marshal operation: %v", err)
	}

	// 将序列化的数据存储到区块链上，键值为 requestID
	return ctx.GetStub().PutState(requestID, operationAsBytes)
}

// QueryRequestOperation 根据 requestID 查询请求操作记录
func (dsc *DirectorySmartContract) QueryRequestOperation(ctx contractapi.TransactionContextInterface, requestID string) (*RequestOperation, error) {
	// 从区块链上获取存储的请求操作数据
	operationAsBytes, err := ctx.GetStub().GetState(requestID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if operationAsBytes == nil {
		return nil, fmt.Errorf("request operation %s does not exist", requestID)
	}

	// 将获取到的 JSON 数据反序列化为 RequestOperation 实例
	var operation RequestOperation
	err = json.Unmarshal(operationAsBytes, &operation)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation: %v", err)
	}

	return &operation, nil
}
