package slaughter

import "gorm.io/gorm"

type SlaugterProcedure struct {
	gorm.Model
	SlaPID string `gorm:"not null; unique; type:varchar(256)" json:"sla_pid"`
}
