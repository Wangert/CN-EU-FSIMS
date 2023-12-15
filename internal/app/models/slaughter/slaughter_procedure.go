package slaughter

import "gorm.io/gorm"

type SlaugterProcedure struct {
	gorm.Model
	SlaPID string `gorm:"not null; unique; type:varchar(256)" json:"sla_pid"`
}

type SlaughterProcedureData struct {
	EnvirTemperature      string `json:"envir_temperature"`
	EnvirLighting         string `json:"envir_lighting"`
	ShockVoltage          string `json:"shock_voltage"`
	BleedingTime          string `json:"bleeding_time"`
	EleShockTime          string `json:"ele_shock_time"`
	KnifeDisinfectionTime string `json:"knife_disinfection_time"`
}
