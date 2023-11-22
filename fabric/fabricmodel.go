package fabric

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
)

type InitInfo struct {
	ChannelID      string
	ChannelConfig  string
	OrgAdmin       string
	OrgName        string
	OrdererOrgName string
	OrgResMgmt     *resmgmt.Client
}

type Block struct {
	Number          uint64         `json:"number"`          //
	PreviousHash    []byte         `json:"previousHash"`    //
	DataHash        []byte         `json:"dataHash"`        //
	BlockHash       []byte         `json:"blockHash"`       //
	TxNum           int            `json:"txNum"`           //
	TransactionList []*Transaction `json:"transactionList"` //
	CreateTime      string         `json:"createTime"`      //
}
type Transaction struct {
	TransactionActionList []*TransactionAction `json:"transactionActionList"` //
}
type TransactionAction struct {
	TxId         string   `json:"txId"`         //
	BlockNum     uint64   `json:"blockNum"`     //
	Type         string   `json:"type"`         //
	Timestamp    string   `json:"timestamp"`    //
	ChannelId    string   `json:"channelId"`    //
	Endorsements []string `json:"endorsements"` //
	ChaincodeId  string   `json:"chaincodeId"`  //
	ReadSetList  []string `json:"readSetList"`  //
	WriteSetList []string `json:"writeSetList"` //
}

type Procedure struct {
	PID       string `json:"pid"`
	PrePID    string `json:"pre_pid"`
	CheckCode string `json:"checkcode"`
}
