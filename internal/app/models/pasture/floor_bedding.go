package pasture

import "gorm.io/gorm"

type PastureFloorBedding struct {
	gorm.Model
	PhysicalHazard PastureFloorBeddingPhysicalHazard `json:"physical_hazard"`
	Biohazard      PastureFloorBeddingBiohazard      `json:"biohazard"`
	PasPID         string                            `gorm:"not null; type:varchar(256)" json:"pas_pid"`
}

type PastureFloorBeddingData struct {
	PhysicalHazard PastureFloorBeddingPhysicalHazardData `json:"physical_hazard"`
	Biohazard      PastureFloorBeddingBiohazardData      `json:"biohazard"`
}
