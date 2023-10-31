package pasture

import "gorm.io/gorm"

type SoilBiohazard struct {
	gorm.Model
	ColiformBacteria uint `json:"coliform_bacteria"` // 总大肠菌群
	TotalBacteria    uint `json:"total_bacteria"`    // 总细菌数
	SoilID           uint `json:"soil_id"`
}
