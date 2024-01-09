package service

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/query"
	"context"
	"fmt"
)

func PushNotification(receiver string, event *Event) error {
	err := CheckUserIsExisted(receiver)
	if err != nil {
		return err
	}

	notification := models.Notification{
		Source:    event.Source,
		Content:   event.Content,
		UUID:      receiver,
		EventTime: event.EventTime,
		EventType: event.EventType,
		Affected:  event.AffectedBatchNumber,
		Proposal:  event.Proposal,
		RiskLevel: event.RiskLevel,
		State:     1,
	}

	err = query.Q.Notification.WithContext(context.Background()).Create(&notification)
	if err != nil {
		return err
	}

	return nil
}

//func SafetyResultNotification(result string, reason string, time string, id string, company string, user string, node models.FoodChainNode) bool {
//	fmt.Println("******************* SafetyResultNotification *********************")
//
//	res := mysql.NewNotification(result, reason, time, id, company, user, node)
//	fmt.Println("notification res: ", res)
//
//	return res
//}

func PlaceRiskNotification(result string, place string, productinfo string, riskinfo string) bool {
	fmt.Println("******************* PlaceRiskNotification *********************")

	return true
}
