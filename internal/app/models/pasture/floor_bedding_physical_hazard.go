package pasture

import "gorm.io/gorm"

type FloorBeddingPhysicalHazard struct {
	gorm.Model
	Mercury        float32 `json:"mercury"`  // 汞
	Cadmium        float32 `json:"cadmium"`  // 镉
	Lead           float32 `json:"lead"`     // 铅
	Chromium       float32 `json:"chromium"` // 铬
	FloorBeddingID uint    `json:"floor_bedding_id"`
}
