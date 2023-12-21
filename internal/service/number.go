package service

import (
	"CN-EU-FSIMS/utils"
	"CN-EU-FSIMS/utils/crypto"
	"time"
)

func GenerateNumber(i interface{}) string {
	t := time.Now().String()
	s := utils.SerializeStructToString(i)
	h := crypto.CalculateSHA256(s+t, "number")

	return h
}
