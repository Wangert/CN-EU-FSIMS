package fatten

import "gorm.io/gorm"

type FattenWaterChemicalHazard struct {
	gorm.Model
	Fluoride        float32 `json:"fluoride"`         // 氟化物
	Cyanide         float32 `json:"cyanide"`          // 氰化物
	Chloride        float32 `json:"chloride"`         // 氰化物
	Nitrate         float32 `json:"nitrate"`          // 硝酸盐
	Sulfate         float32 `json:"sulfate"`          // 硫酸盐
	Sixsixsix       float32 `json:"sixsixsix"`        // 六六六
	DDT             float32 `json:"ddt"`              // 滴滴涕
	AmmoniaNitrogen float32 `json:"ammonia_nitrogen"` // 氨氮
	FattenWaterID   uint    `json:"fatten_water_id"`
}
