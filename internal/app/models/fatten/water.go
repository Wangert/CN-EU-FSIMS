package fatten

import "gorm.io/gorm"

type FattenWater struct {
	gorm.Model
	PhysicalHazard FattenWaterPhysicalHazard `json:"physical_hazard"`
	ChemicalHazard FattenWaterChemicalHazard `json:"chemical_hazard"`
	Biohazard      FattenWaterBiohazard      `json:"biohazard"`
	SensoryTraits  FattenWaterSensoryTraits  `json:"sensory_traits"`
	FatPID         string                    `json:"fat_pid"`
}
