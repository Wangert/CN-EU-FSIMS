package response

import "CN-EU-FSIMS/internal/app/models"

type ResFoodchains struct {
	Foodchains     []models.Foodchain `json:"foodchains"`
	TotalCount     int64              `json:"total_count"`
	CompletedCount int64              `json:"completed_count"`
}
