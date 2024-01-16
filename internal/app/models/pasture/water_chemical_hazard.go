package pasture

import "gorm.io/gorm"

type PastureWaterChemicalHazard struct {
	gorm.Model
	Fluoride        float64 `json:"fluoride"`         // 氟化物
	Cyanide         float64 `json:"cyanide"`          // 氰化物
	Chloride        float64 `json:"chloride"`         // 氰化物
	Nitrate         float64 `json:"nitrate"`          // 硝酸盐
	Sulfate         float64 `json:"sulfate"`          // 硫酸盐
	Sixsixsix       float64 `json:"sixsixsix"`        // 六六六
	DDT             float64 `json:"ddt"`              // 滴滴涕
	AmmoniaNitrogen float64 `json:"ammonia_nitrogen"` // 氨氮
	PastureWaterID  uint    `json:"pasture_water_id"`
}

type PastureWaterChemicalHazardData struct {
	Fluoride        float64 `json:"fluoride"`         // 氟化物
	Cyanide         float64 `json:"cyanide"`          // 氰化物
	Chloride        float64 `json:"chloride"`         // 氰化物
	Nitrate         float64 `json:"nitrate"`          // 硝酸盐
	Sulfate         float64 `json:"sulfate"`          // 硫酸盐
	Sixsixsix       float64 `json:"sixsixsix"`        // 六六六
	DDT             float64 `json:"ddt"`              // 滴滴涕
	AmmoniaNitrogen float64 `json:"ammonia_nitrogen"` // 氨氮
}
