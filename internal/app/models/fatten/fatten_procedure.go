package fatten

import "gorm.io/gorm"

type FattenProcedure struct {
	gorm.Model
	FatPID string      `gorm:"not null; unique; type:varchar(256)" json:"fat_pid"`
	Water  FattenWater `gorm:"foreignKey:FatPID; references:FatPID"`
	Soil   FattenSoil  `gorm:"foreignKey:FatPID; references:FatPID"`
}
