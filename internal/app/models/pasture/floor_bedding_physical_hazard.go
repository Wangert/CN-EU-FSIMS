package pasture

import "gorm.io/gorm"

type PastureFloorBeddingPhysicalHazard struct {
	gorm.Model
	Mercury               float64 `json:"mercury"`  // 汞
	Cadmium               float64 `json:"cadmium"`  // 镉
	Lead                  float64 `json:"lead"`     // 铅
	Chromium              float64 `json:"chromium"` // 铬
	PastureFloorBeddingID uint    `json:"pasture_floor_bedding_id"`
}

type PastureFloorBeddingPhysicalHazardData struct {
	Mercury               float64 `json:"mercury"`  // 汞
	Cadmium               float64 `json:"cadmium"`  // 镉
	Lead                  float64 `json:"lead"`     // 铅
	Chromium              float64 `json:"chromium"` // 铬
	PastureFloorBeddingID uint    `json:"pasture_floor_bedding_id"`
}
