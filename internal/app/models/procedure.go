package models

import (
	"gorm.io/gorm"
	"time"
)

type Procedure struct {
	gorm.Model
	PID                string         `gorm:"not null; unique; type:varchar(256)" json:"pid"`
	Type               uint           `gorm:"not null" json:"type"`
	Name               string         `gorm:"not null; type:varchar(100)" json:"name"`
	state              uint           `gorm:"not null" json:"state"`
	PHash              string         `gorm:"unique; type:varchar(256)" json:"p_hash"`
	CheckCode          string         `gorm:"type:varchar(256)" json:"check_code"`
	SerialNumber       uint           `gorm:"not null" json:"serial_number"`
	Operator           string         `json:"operator"`
	StartTimestamp     time.Time      `gorm:"not null" json:"start_timestamp"`
	CompletedTimestamp time.Time      `json:"completed_timestamp"`
	PrePID             string         `gorm:"not null; type:varchar(256)" json:"pre_pid"`
	ICID               string         `gorm:"not null; type:varchar(256)" json:"icid"`
	SubProcedures      []SubProcedure `gorm:"foreignKey:PID; references:PID" json:"sub_procedures"`
}
