package chaincode

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestStoreRequestOperation(t *testing.T) {
	contract := new(DemoContract)
	ctx := new(MockTransactionContext)
	stub := new(MockChaincodeStub)
	ctx.On("GetStub").Return(stub)

	// 模拟 PutState 的行为
	stub.On("PutState", "REQ001", mock.Anything).Return(nil)

	// 调用 StoreRequestOperation 方法
	err := contract.StoreRequestOperation(ctx, "REQ001", []byte("sample_value"))

	// 验证是否调用成功且无错误
	assert.Nil(t, err, "StoreRequestOperation should not return an error")

	// 验证 PutState 方法是否被正确调用
	stub.AssertCalled(t, "PutState", "REQ001", mock.Anything)
}

func TestQueryData(t *testing.T) {
	contract := new(DemoContract)
	ctx := new(MockTransactionContext)
	stub := new(MockChaincodeStub)
	ctx.On("GetStub").Return(stub)

	// 设置模拟的返回数据
	expectedData := DemoData{Value: []byte("sample_value")}
	expectedDataAsBytes, _ := json.Marshal(expectedData)

	// 模拟 GetState 的行为
	stub.On("GetState", "REQ001").Return(expectedDataAsBytes, nil)

	// 调用 QueryData 方法
	result, err := contract.QueryData(ctx, "REQ001")

	// 验证无错误返回
	assert.Nil(t, err, "QueryData should not return an error")

	// 验证返回的数据是否正确
	assert.Equal(t, expectedData, *result, "QueryData should return the correct data")
}
