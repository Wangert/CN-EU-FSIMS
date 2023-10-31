package pasture

import "gorm.io/gorm"

type PastureWaterPhysicalHazard struct {
	gorm.Model
	Mercury        float32 `json:"mercury"`  // 汞
	Cadmium        float32 `json:"cadmium"`  // 镉
	Lead           float32 `json:"lead"`     // 铅
	Chromium       float32 `json:"chromium"` // 铬
	Arsenic        float32 `json:"arsenic"`  // 砷
	Copper         float32 `json:"copper"`   // 铜
	PastureWaterID uint    `json:"pasture_water_id"`
}
