package models

import "gorm.io/gorm"

type Procedure struct {
	gorm.Model
	PID          string `gorm:"not null; unique; type:varchar(256)" json:"pid"`
	Name         string `gorm:"not null; type:varchar(100)" json:"name"`
	SerialNumber uint   `gorm:"not null" json:"serial_number"`
	PrePID       string `gorm:"not null; type:varchar(256)" json:"pre_pid"`
	ICID         string `gorm:"not null; type:varchar(256)" json:"icid"`
}
