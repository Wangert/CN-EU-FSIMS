package pasture

import "gorm.io/gorm"

type WaterBiohazard struct {
	gorm.Model
	ColiformBacteria uint `json:"coliform_bacteria"` // 总大肠菌群
	WaterID          uint `json:"water_id"`
}
