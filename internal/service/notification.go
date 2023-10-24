package service

import (
	"CN-EU-FSIMS/database/mysql"
	"CN-EU-FSIMS/internal/app/models"
	"fmt"
)

func SafetyResultNotification(result string, reason string, time string, id string, company string, user string, node models.FoodChainNode) bool {
	fmt.Println("******************* SafetyResultNotification *********************")

	res := mysql.NewNotification(result, reason, time, id, company, user, node)
	fmt.Println("notification res: ", res)

	return res
}

func PlaceRiskNotification(result string, place string, productinfo string, riskinfo string) bool {
	fmt.Println("******************* PlaceRiskNotification *********************")

	return true
}
