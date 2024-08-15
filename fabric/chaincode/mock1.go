package chaincode

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/stretchr/testify/mock"
)

// MockTransactionContext 模拟 TransactionContextInterface
type MockTransactionContext struct {
	mock.Mock
	contractapi.TransactionContextInterface
}

// MockChaincodeStub 模拟 ChaincodeStubInterface
type MockChaincodeStub struct {
	mock.Mock
	shim.ChaincodeStubInterface
}

// PutState 模拟 PutState 方法
func (m *MockChaincodeStub) PutState(key string, value []byte) error {
	args := m.Called(key, value)
	return args.Error(0)
}

// GetState 模拟 GetState 方法
func (m *MockChaincodeStub) GetState(key string) ([]byte, error) {
	args := m.Called(key)
	return args.Get(0).([]byte), args.Error(1)
}
