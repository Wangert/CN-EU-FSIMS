package fabric

/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
	"unsafe"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset/kvrwset"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

var contract *gateway.Contract
var ledgerClient *ledger.Client

// 测试fabric连接
func ConnecttoNetwork() (result bool, channelname string, chaincodename string) {
	result = false
	log.Println("============ ConnecttoNetwork1111 ============")

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "false")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"..",
		"..",
		"fabric",
		"organizations",
		"peerOrganizations",
		"org1",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("channel1")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}
	contract = network.GetContract("Fsims")
	if contract.Name() != "" {
		fmt.Printf("成功连接到fabric网络，已监测到合约： %s\n", contract.Name())
		result = true
		channelname = "channel1"
		chaincodename = contract.Name()
		return
	}
	return
}

func GetLedgerClient() {
	fmt.Println("============ GetLedgerClient ============")
	ccpPath := filepath.Join(
		"..",
		"..",
		"..",
		"fabric",
		"organizations",
		"peerOrganizations",
		"org1",
		"connection-org1.yaml",
	)
	sdk, err := fabsdk.New(config.FromFile(filepath.Clean(ccpPath)))
	if err != nil {
		fmt.Println("fabsdk.New error: ", err)
	}

	channel := sdk.ChannelContext("channel1", fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))

	ledgerClient, err = ledger.New(channel)
	if err != nil {
		fmt.Println("ledger.New err: ", err)
	}
}

// 根据高度查询区块
func QueryBlockByHeight(num int) (Block, error) {
	//fmt.Println("============ QueryBlockByHeight ============")
	//num, _ := strconv.Atoi(c.Query("num"))

	GetLedgerClient()
	rawBlock, err := ledgerClient.QueryBlock(uint64(num))

	txList := []*Transaction{}
	for i := range rawBlock.Data.Data {
		rawEnvelope, err := GetEnvelopeFromBlock(rawBlock.Data.Data[i])
		if err != nil {
			fmt.Printf("QueryBlock return error: %s", err)
		}
		transaction, err := GetTransactionFromEnvelopeDeep(rawEnvelope)
		if err != nil {
			fmt.Printf("QueryBlock return error: %s", err)
		}
		for i := range transaction.TransactionActionList {
			transaction.TransactionActionList[i].BlockNum = rawBlock.Header.Number
		}
		txList = append(txList, transaction)
	}

	block := Block{
		Number:          rawBlock.Header.Number,
		PreviousHash:    rawBlock.Header.PreviousHash,
		DataHash:        rawBlock.Header.DataHash,
		BlockHash:       rawBlock.Header.DataHash,
		TxNum:           len(rawBlock.Data.Data),
		TransactionList: txList,
		CreateTime:      txList[0].TransactionActionList[0].Timestamp,
	}
	return block, err
}

func StringToBytes(data string) []byte {
	return *(*[]byte)(unsafe.Pointer(&data))
}

// 修改路由
// func QueryBlockByHash(hash string)(Block,error) {
// 	//hash := c.PostForm("hash")
// 	//fmt.Println(hash)
// 	GetLedgerClient()

// 	ledgerInfo, err := ledgerClient.QueryInfo()

// 	hashByte := StringToBytes(hash)
// 	fmt.Println("hash", hashByte)
// 	fmt.Println(ledgerInfo.BCI.CurrentBlockHash)

// 	block, err := ledgerClient.QueryBlockByHash(ledgerInfo.BCI.CurrentBlockHash)
// 	// if err != nil {
// 	// 	fmt.Println("querybyhash err: ", err)
// 	// }
// 	// c.JSON(http.StatusOK, gin.H{
// 	// 	"result": true,
// 	// 	"block":  block,
// 	// })

//     return block,err
// }

func GetLastestBlock() ([]Block, error) {
	GetLedgerClient()

	var blocks []Block

	ledgerInfo, err := ledgerClient.QueryInfo()
	if err != nil {
		fmt.Println("err: ")
	}
	height := ledgerInfo.BCI.Height - 1
	for j := height; j > 0 && j > (height-5); j-- {
		rawBlock, _ := ledgerClient.QueryBlock(j)

		txList := []*Transaction{}
		for i := range rawBlock.Data.Data {
			rawEnvelope, err := GetEnvelopeFromBlock(rawBlock.Data.Data[i])
			if err != nil {
				fmt.Printf("QueryBlock return error: %s", err)
			}
			transaction, err := GetTransactionFromEnvelopeDeep(rawEnvelope)
			if err != nil {
				fmt.Printf("QueryBlock return error: %s", err)
			}
			for i := range transaction.TransactionActionList {
				transaction.TransactionActionList[i].BlockNum = rawBlock.Header.Number
			}
			txList = append(txList, transaction)
		}

		block := Block{
			Number:          rawBlock.Header.Number,
			PreviousHash:    rawBlock.Header.PreviousHash,
			DataHash:        rawBlock.Header.DataHash,
			BlockHash:       rawBlock.Header.DataHash,
			TxNum:           len(rawBlock.Data.Data),
			TransactionList: txList,
			CreateTime:      txList[0].TransactionActionList[0].Timestamp,
		}

		blocks = append(blocks, block)
	}
	return blocks, err
	// c.JSON(http.StatusOK, gin.H{
	// 	"result": true,
	// 	"blocks": blocks,
	// })

}

func GetEnvelopeFromBlock(data []byte) (*common.Envelope, error) {
	var err error
	env := &common.Envelope{}
	if err = proto.Unmarshal(data, env); err != nil {
		fmt.Printf("block unmarshal err: %s", err)
	}
	return env, nil
}

func GetTransactionFromEnvelopeDeep(rawEnvelope *common.Envelope) (*Transaction, error) {

	rawPayload := &common.Payload{}
	err := proto.Unmarshal(rawEnvelope.Payload, rawPayload)
	if err != nil {
		fmt.Printf("block unmarshal err: %s", err)
	}

	channelHeader := &common.ChannelHeader{}
	err = proto.Unmarshal(rawPayload.Header.ChannelHeader, channelHeader)
	err = proto.Unmarshal(rawPayload.Header.ChannelHeader, channelHeader)
	if err != nil {
		fmt.Printf("block unmarshal err: %s\n", err)
	}

	transactionObj := &peer.Transaction{}
	err = proto.Unmarshal(rawPayload.Data, transactionObj)
	if err != nil {
		fmt.Printf("block unmarshal err: %s\n", err)
	}
	transactionActionList := []*TransactionAction{}

	for i := range transactionObj.Actions {
		transactionAction, err := GetTransactionActionFromTransactionDeep(transactionObj.Actions[i])
		if err != nil {
			fmt.Printf("block unmarshal err: %s\n", err)
		}
		transactionAction.TxId = channelHeader.TxId
		transactionAction.Type = string(channelHeader.Type)
		transactionAction.Timestamp = time.Unix(channelHeader.Timestamp.Seconds, 0).Format("2006-01-02 15:04:05")
		transactionAction.ChannelId = channelHeader.ChannelId
		transactionActionList = append(transactionActionList, transactionAction)
	}
	transaction := Transaction{transactionActionList}
	return &transaction, nil
}

func GetTransactionActionFromTransactionDeep(transactionAction *peer.TransactionAction) (*TransactionAction, error) {

	ChaincodeActionPayload := &peer.ChaincodeActionPayload{}
	err := proto.Unmarshal(transactionAction.Payload, ChaincodeActionPayload)
	if err != nil {
		fmt.Printf("block unmarshal err: %s\n", err)
	}

	ProposalResponsePayload := &peer.ProposalResponsePayload{}
	ChaincodeAction := &peer.ChaincodeAction{}
	chaincodeId := ""
	NsReadWriteSetList := []*rwset.NsReadWriteSet{}
	ReadWriteSetList := []*kvrwset.KVRWSet{}
	readSetList := []string{}
	writeSetList := []string{}
	if ChaincodeActionPayload.GetAction() != nil {
		err = proto.Unmarshal(ChaincodeActionPayload.Action.ProposalResponsePayload, ProposalResponsePayload)
		if err != nil {
			fmt.Printf("block unmarshal err: %s", err)
		}

		err = proto.Unmarshal(ProposalResponsePayload.Extension, ChaincodeAction)
		if err != nil {
			fmt.Printf("block unmarshal err: %s", err)
		}
		chaincodeId = ChaincodeAction.ChaincodeId.Name

		TxReadWriteSet := &rwset.TxReadWriteSet{}
		err = proto.Unmarshal(ChaincodeAction.Results, TxReadWriteSet)
		if err != nil {
			fmt.Printf("block unmarshal err: %s", err)
		}

		for i := range TxReadWriteSet.NsRwset {
			ReadWriteSet := &kvrwset.KVRWSet{}
			err = proto.Unmarshal(TxReadWriteSet.NsRwset[i].Rwset, ReadWriteSet)
			if err != nil {
				fmt.Printf("block unmarshal err: %s", err)
			}

			for i := range ReadWriteSet.Reads {
				readSetJsonStr, err := json.Marshal(ReadWriteSet.Reads[i])
				if err != nil {
					fmt.Printf("block unmarshal err: %s", err)
				}
				readSetList = append(readSetList, string(readSetJsonStr))
			}

			for i := range ReadWriteSet.Writes {
				writeSetItem := map[string]interface{}{
					"Key":      ReadWriteSet.Writes[i].GetKey(),
					"Value":    string(ReadWriteSet.Writes[i].GetValue()),
					"IsDelete": ReadWriteSet.Writes[i].GetIsDelete(),
				}
				writeSetJsonStr, err := json.Marshal(writeSetItem)
				if err != nil {
					fmt.Printf("block unmarshal err: %s", err)
				}
				writeSetList = append(writeSetList, string(writeSetJsonStr))
			}
			ReadWriteSetList = append(ReadWriteSetList, ReadWriteSet)
			NsReadWriteSetList = append(NsReadWriteSetList, TxReadWriteSet.NsRwset[i])
		}
	} else {
		chaincodeId = "没有交易数据"
	}

	endorsements := []string{}
	if ChaincodeActionPayload.Action.GetEndorsements() != nil {
		for i := range ChaincodeActionPayload.Action.GetEndorsements() {
			endorser := &msp.SerializedIdentity{}
			err = proto.Unmarshal(ChaincodeActionPayload.Action.Endorsements[i].Endorser, endorser)
			if err != nil {
				fmt.Printf("block unmarshal err: %s", err)
			}
			endorsements = append(endorsements, string(endorser.Mspid))
		}
	}
	transactionActionObj := TransactionAction{
		Endorsements: endorsements,
		ChaincodeId:  chaincodeId,
		ReadSetList:  readSetList,
		WriteSetList: writeSetList,
	}

	return &transactionActionObj, nil

}

func GetLedgerInfo() (*fab.BlockchainInfoResponse, error) {

	GetLedgerClient()

	ledgerInfo, err := ledgerClient.QueryInfo()
	return ledgerInfo, err
}

// 获得fabric对应账本的世界状态
// func GetAllFabricResult() []FoodChainData {
// 	log.Println("============ GetAllResult ============")
// 	ConnecttoNetwork()
// 	log.Println(contract.Name())
// 	log.Println("--> Evaluate Transaction: GetAllFoodChainData, function returns all the current assets on the ledger")
// 	result, err := contract.EvaluateTransaction("GetAllFoodChainData")
// 	if err != nil {
// 		log.Fatalf("Failed to evaluate transaction: %v", err)
// 	}
// 	var list []FoodChainData
// 	err = json.Unmarshal(result, &list)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		log.Println(list)
// 	}
// 	log.Println(string(result))
// 	return list
// }

func QueryProcedureWithPID(pid string) Procedure {
	log.Println("============ QueryProcedureWithPID ============")
	ConnecttoNetwork()
	log.Println(contract.Name())
	txn, err := contract.CreateTransaction("QueryProcedureWithPID", gateway.WithEndorsingPeers("peer0.org1", "peer0.org2"))
	if err != nil {
		log.Fatalf("Failed to create Tx: %v", err)
	}
	result, err := txn.Submit(pid)
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	} else {
		fmt.Println("success result: ", string(result))
	}

	var resultJson Procedure
	err = json.Unmarshal(result, &resultJson)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(resultJson)
	}

	return resultJson
}

func UploadProcedure(pid string, pre_id string) ([]byte, error) {
	glog.Info("============ UploadProcedure ============")
	ConnecttoNetwork()
	fmt.Println("contract name: ", contract.Name())
	txn, err := contract.CreateTransaction("UploadProcedure", gateway.WithEndorsingPeers("peer0.org1", "peer0.org2"))
	if err != nil {
		log.Fatalf("Failed to create Tx: %v", err)
	}
	result, err := txn.Submit(pid, pre_id)
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	} else {
		fmt.Println("success result: ", result)
	}
	return result, err
}

func UpdateProcedure(pid string, checkcode string, p_hash string) ([]byte, error) {
	fmt.Println("============ UpdateProcedure ============")
	ConnecttoNetwork()
	fmt.Println("contract name: ", contract.Name())
	txn, err := contract.CreateTransaction("UpdateProcedure", gateway.WithEndorsingPeers("peer0.org1", "peer0.org2"))
	if err != nil {
		log.Fatalf("Failed to create Tx: %v", err)
	}
	result, err := txn.Submit(pid, checkcode, p_hash)
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
		return nil, err
	} else {
		fmt.Println("success result: ", result)
	}
	return result, err

}

func DeleteTest(id string) bool {
	fmt.Println("============ DeleteTest ============")
	ConnecttoNetwork()
	fmt.Println("contract name: ", contract.Name())
	txn, err := contract.CreateTransaction("DeleteAsset", gateway.WithEndorsingPeers("peer0.org1", "peer0.org2"))
	if err != nil {
		log.Fatalf("Failed to create Tx: %v", err)
	}
	result, err := txn.Submit(id)
	res := false
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	} else {
		res = true
		fmt.Println("delete result: ", result)
	}
	return res
}

func populateWallet(wallet *gateway.Wallet) error {
	log.Println("============ Populating wallet ============")
	credPath := filepath.Join(
		"..",
		"..",
		"..",
		"fabric",
		"organizations",
		"peerOrganizations",
		"org1",
		"users",
		"User1@org1",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	return wallet.Put("appUser", identity)
}
