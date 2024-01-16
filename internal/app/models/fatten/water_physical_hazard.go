package fatten

import "gorm.io/gorm"

type FattenWaterPhysicalHazard struct {
	gorm.Model
	Mercury       float64 `json:"mercury"`  // 汞
	Cadmium       float64 `json:"cadmium"`  // 镉
	Lead          float64 `json:"lead"`     // 铅
	Chromium      float64 `json:"chromium"` // 铬
	Arsenic       float64 `json:"arsenic"`  // 砷
	Copper        float64 `json:"copper"`   // 铜
	FattenWaterID uint    `json:"fatten_water_id"`
}
