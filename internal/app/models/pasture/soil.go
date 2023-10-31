package pasture

import "gorm.io/gorm"

type PastureSoil struct {
	gorm.Model
	PhysicalHazard PastureSoilPhysicalHazard `json:"physical_hazard"`
	Biohazard      PastureSoilBiohazard      `json:"biohazard"`
	PasPID         string                    `gorm:"not null; type:varchar(256)" json:"pas_pid"`
}
