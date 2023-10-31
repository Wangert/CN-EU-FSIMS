package fatten

import "gorm.io/gorm"

type FattenWaterBiohazard struct {
	gorm.Model
	ColiformBacteria uint `json:"coliform_bacteria"` // 总大肠菌群
	FattenWaterID    uint `json:"fatten_water_id"`
}
