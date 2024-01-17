package pasture

import "gorm.io/gorm"

type PastureWaterPhysicalHazard struct {
	gorm.Model
	Mercury        float64 `json:"mercury"`  // 汞
	Cadmium        float64 `json:"cadmium"`  // 镉
	Lead           float64 `json:"lead"`     // 铅
	Chromium       float64 `json:"chromium"` // 铬
	Arsenic        float64 `json:"arsenic"`  // 砷
	Copper         float64 `json:"copper"`   // 铜
	PastureWaterID uint    `json:"pasture_water_id"`
}

type PastureWaterPhysicalHazardData struct {
	Mercury  float64 `json:"mercury"`  // 汞
	Cadmium  float64 `json:"cadmium"`  // 镉
	Lead     float64 `json:"lead"`     // 铅
	Chromium float64 `json:"chromium"` // 铬
	Arsenic  float64 `json:"arsenic"`  // 砷
	Copper   float64 `json:"copper"`   // 铜
}
