package pasture

import "gorm.io/gorm"

type Soil struct {
	gorm.Model
	PhysicalHazard SoilPhysicalHazard `json:"physical_hazard"`
	Biohazard      SoilBiohazard      `json:"biohazard"`
	PasPID         string             `json:"pas_pid"`
}
