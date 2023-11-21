package pasture

import "gorm.io/gorm"

type PastureSoilBiohazard struct {
	gorm.Model
	ColiformBacteria uint `json:"coliform_bacteria"` // 总大肠菌群
	TotalBacteria    uint `json:"total_bacteria"`    // 总细菌数
	PastureSoilID    uint `json:"pasture_soil_id"`
}

type PastureSoilBiohazardData struct {
	ColiformBacteria uint `json:"coliform_bacteria"` // 总大肠菌群
	TotalBacteria    uint `json:"total_bacteria"`    // 总细菌数
}
