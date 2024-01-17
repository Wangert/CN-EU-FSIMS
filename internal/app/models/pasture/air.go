package pasture

import "gorm.io/gorm"

type PastureAir struct {
	gorm.Model
	TotalBacteria   uint    `json:"total_bacteria"`   // 总细菌数
	AmmoniaGas      float64 `json:"ammonia_gas"`      // 氨气
	HydrogenSulfide float64 `json:"hydrogen_sulfide"` // 硫化氢
	CarbonDioxide   float64 `json:"carbon_dioxide"`   // 二氧化碳
	PM10            float64 `json:"pm_10"`            // PM10
	TSP             float64 `json:"tsp"`              // TSP
	Stench          uint    `json:"stench"`           // 恶臭稀释程度
	PasPID          string  `json:"pas_pid"`
}

type PastureAirData struct {
	TotalBacteria   uint    `json:"total_bacteria"`   // 总细菌数
	AmmoniaGas      float64 `json:"ammonia_gas"`      // 氨气
	HydrogenSulfide float64 `json:"hydrogen_sulfide"` // 硫化氢
	CarbonDioxide   float64 `json:"carbon_dioxide"`   // 二氧化碳
	PM10            float64 `json:"pm_10"`            // PM10
	TSP             float64 `json:"tsp"`              // TSP
	Stench          uint    `json:"stench"`           // 恶臭稀释程度
}
