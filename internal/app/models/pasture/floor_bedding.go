package pasture

import "gorm.io/gorm"

type FloorBedding struct {
	gorm.Model
	PhysicalHazard FloorBeddingPhysicalHazard `json:"physical_hazard"`
	Biohazard      FloorBeddingBiohazard      `json:"biohazard"`
	PasPID         string                     `json:"pas_pid"`
}
