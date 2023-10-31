package pasture

import "gorm.io/gorm"

type Fodder struct {
	gorm.Model
	PhysicalHazard FodderPhysicalHazard `json:"physical_hazard"`
	Biohazard      FodderBiohazard      `json:"biohazard"`
	PasPID         string               `json:"pas_pid"`
}
