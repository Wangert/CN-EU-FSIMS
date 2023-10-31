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
