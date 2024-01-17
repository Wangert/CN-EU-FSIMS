package response

import "CN-EU-FSIMS/internal/app/models"

type ResFoodchains struct {
	Foodchains     []models.Foodchain `json:"foodchains"`
	TotalCount     int64              `json:"total_count"`
	CompletedCount int64              `json:"completed_count"`
}

type ResPidInfo struct {
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Address     string `json:"address"`
	HouseNumber string `json:"house_number"`
}
