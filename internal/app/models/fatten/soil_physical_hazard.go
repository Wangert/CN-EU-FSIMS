package fatten

import "gorm.io/gorm"

type FattenSoilPhysicalHazard struct {
	gorm.Model
	Cadmium      float32 `json:"cadmium"`  // 镉
	Arsenic      float32 `json:"arsenic"`  // 砷
	Copper       float32 `json:"copper"`   // 铜
	Lead         float32 `json:"lead"`     // 铅
	Chromium     float32 `json:"chromium"` // 铬
	Zinc         float32 `json:"zinc"`     // 锌
	FattenSoilID uint    `json:"fatten_soil_id"`
}
