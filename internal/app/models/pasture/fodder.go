package pasture

import "gorm.io/gorm"

type PastureFodder struct {
	gorm.Model
	PhysicalHazard PastureFodderPhysicalHazard `json:"physical_hazard"`
	Biohazard      PastureFodderBiohazard      `json:"biohazard"`
	PasPID         string                      `gorm:"not null; type:varchar(256)" json:"pas_pid"`
}

type PastureFodderData struct {
	PhysicalHazard PastureFodderPhysicalHazardData `json:"physical_hazard"`
	Biohazard      PastureFodderBiohazardData      `json:"biohazard"`
}
