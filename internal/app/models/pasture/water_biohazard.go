package pasture

import "gorm.io/gorm"

type PastureWaterBiohazard struct {
	gorm.Model
	ColiformBacteria uint `json:"coliform_bacteria"` // 总大肠菌群
	PastureWaterID   uint `json:"pasture_water_id"`
}

type PastureWaterBiohazardData struct {
	ColiformBacteria uint `json:"coliform_bacteria"` // 总大肠菌群
}
