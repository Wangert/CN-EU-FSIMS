package fatten

import "gorm.io/gorm"

type FattenWaterSensoryTraits struct {
	gorm.Model
	Color          uint   `json:"color"`           // 色
	Turbidity      uint   `json:"turbidity"`       // 浑浊度
	Smell          string `json:"smell"`           // 气味
	PH             uint   `json:"ph"`              // PH值
	Hardness       uint   `json:"hardness"`        // 总硬度
	DissolvedSolid uint   `json:"dissolved_solid"` // 溶解性总固形物
	FattenWaterID  uint   `json:"fatten_water_id"`
}
