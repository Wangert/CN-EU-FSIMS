package fatten

import "gorm.io/gorm"

type FattenSoil struct {
	gorm.Model
	PhysicalHazard FattenSoilPhysicalHazard `json:"physical_hazard"`
	Biohazard      FattenSoilBiohazard      `json:"biohazard"`
	FatPID         string                   `json:"fat_pid"`
}
