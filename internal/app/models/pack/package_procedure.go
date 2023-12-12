package pack

import "gorm.io/gorm"

type PackageProcedure struct {
	gorm.Model
	PacPID string `gorm:"not null; unique; type:varchar(256)" json:"pac_pid"`
}

type PackProcedureData struct {
	PackType        string `json:"pack_type"`
	PackTemperature string `json:"pack_temperature"`
}
