package models

import "gorm.io/gorm"

type SubProcedure struct {
	gorm.Model
	SPID         string `gorm:"not null; unique; type:varchar(256)" json:"spid"`
	Name         string `gorm:"not null; type:varchar(100)" json:"name"`
	SerialNumber uint   `gorm:"not null" json:"serial_number"`
	PreSPID      string `gorm:"not null; type:varchar(256)" json:"pre_pid"`
	PID          string `gorm:"not null; type:varchar(256)" json:"pid"`
}
