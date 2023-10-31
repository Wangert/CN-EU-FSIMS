package pasture

import "gorm.io/gorm"

type Water struct {
	gorm.Model
	PhysicalHazard WaterPhysicalHazard `json:"physical_hazard"`
	ChemicalHazard WaterChemicalHazard `json:"chemical_hazard"`
	Biohazard      WaterBiohazard      `json:"biohazard"`
	SensoryTraits  WaterSensoryTraits  `json:"sensory_traits"`
	PasPID         string              `json:"pas_pid"`
}
