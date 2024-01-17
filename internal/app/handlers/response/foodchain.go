package response

import "CN-EU-FSIMS/internal/app/models"

type ResFoodchains struct {
	Foodchains []models.Foodchain `json:"foodchains"`
	Count      int64              `json:"count"`
}

type ResPidInfo struct {
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Address     string `json:"address"`
	HouseNumber string `json:"house_number"`
}
