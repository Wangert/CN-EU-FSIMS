package response

import "CN-EU-FSIMS/internal/app/models"

type ResFoodchains struct {
	Foodchains []models.Foodchain `json:"foodchains"`
	Count      int64              `json:"count"`
}
