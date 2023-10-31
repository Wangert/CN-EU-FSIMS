package fatten

import "gorm.io/gorm"

type FattenSoilBiohazard struct {
	gorm.Model
	ColiformBacteria uint `json:"coliform_bacteria"` // 总大肠菌群
	TotalBacteria    uint `json:"total_bacteria"`    // 总细菌数
	FattenSoilID     uint `json:"fatten_soil_id"`
}
