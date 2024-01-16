package pasture

import "gorm.io/gorm"

type PastureSoilPhysicalHazard struct {
	gorm.Model
	Cadmium       float64 `json:"cadmium"`  // 镉
	Arsenic       float64 `json:"arsenic"`  // 砷
	Copper        float64 `json:"copper"`   // 铜
	Lead          float64 `json:"lead"`     // 铅
	Chromium      float64 `json:"chromium"` // 铬
	Zinc          float64 `json:"zinc"`     // 锌
	PastureSoilID uint    `json:"pasture_soil_id"`
}

type PastureSoilPhysicalHazardData struct {
	Cadmium  float64 `json:"cadmium"`  // 镉
	Arsenic  float64 `json:"arsenic"`  // 砷
	Copper   float64 `json:"copper"`   // 铜
	Lead     float64 `json:"lead"`     // 铅
	Chromium float64 `json:"chromium"` // 铬
	Zinc     float64 `json:"zinc"`     // 锌
}
