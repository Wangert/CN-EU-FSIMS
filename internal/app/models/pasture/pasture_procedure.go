package pasture

import "gorm.io/gorm"

type PastureProcedure struct {
	gorm.Model
	PasPID         string                `gorm:"not null; unique; index; type:varchar(256)" json:"pas_pid"`
	Water          PastureWater          `gorm:"foreignKey:PasPID; references:PasPID" json:"water"`
	Fodder         PastureFodder         `gorm:"foreignKey:PasPID; references:PasPID" json:"fodder"`
	Soil           PastureSoil           `gorm:"foreignKey:PasPID; references:PasPID" json:"soil"`
	Air            PastureAir            `gorm:"foreignKey:PasPID; references:PasPID" json:"air"`
	FloorBedding   PastureFloorBedding   `gorm:"foreignKey:PasPID; references:PasPID" json:"floor_bedding"`
	WasteDischarge PastureWasteDischarge `gorm:"foreignKey:PasPID; references:PasPID" json:"waste_discharge"`
}

type PastureProcedureData struct {
	Water          PastureWaterData
	Fodder         PastureFodderData
	Soil           PastureSoilData
	Air            PastureAirData
	FloorBedding   PastureFloorBeddingData
	WasteDischarge PastureWasteDischargeData
	PM10           float64 `json:"pm_10"`  // PM10
	TSP            float64 `json:"tsp"`    // TSP
	Stench         uint    `json:"stench"` // 恶臭稀释程度
}
