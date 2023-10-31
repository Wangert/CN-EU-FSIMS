package pasture

import "gorm.io/gorm"

type PastureProcedure struct {
	gorm.Model
	PasPID         string         `gorm:"not null; unique; index; type:varchar(256)" json:"pas_pid"`
	Water          Water          `gorm:"foreignKey:PasPID; reference:PasPID"`
	Fodder         Fodder         `gorm:"foreignKey:PasPID; reference:PasPID"`
	Soil           Soil           `gorm:"foreignKey:PasPID; reference:PasPID"`
	Air            Air            `gorm:"foreignKey:PasPID; reference:PasPID"`
	FloorBedding   FloorBedding   `gorm:"foreignKey:PasPID; reference:PasPID"`
	WasteDischarge WasteDischarge `gorm:"foreignKey:PasPID; reference:PasPID"`
}
