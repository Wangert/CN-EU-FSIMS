package pasture

import "gorm.io/gorm"

type PastureWater struct {
	gorm.Model
	PhysicalHazard PastureWaterPhysicalHazard `gorm:"foreignKey:PastureWaterID" json:"physical_hazard"`
	ChemicalHazard PastureWaterChemicalHazard `gorm:"foreignKey:PastureWaterID" json:"chemical_hazard"`
	Biohazard      PastureWaterBiohazard      `gorm:"foreignKey:PastureWaterID" json:"biohazard"`
	SensoryTraits  PastureWaterSensoryTraits  `gorm:"foreignKey:PastureWaterID" json:"sensory_traits"`
	PasPID         string                     `gorm:"type:varchar(256)" json:"pas_pid"`
}
