package test

import (
	"CN-EU-FSIMS/fabric"
	"fmt"
	"testing"
)

func TestFabricConnectNetwork(t *testing.T) {

	re, cn, cd := fabric.ConnecttoNetwork()

	fmt.Println(re)
	fmt.Println(cn)
	fmt.Println(cd)
}

func TestGetLedger(t *testing.T) {
	fabric.GetLedgerClient()
}
