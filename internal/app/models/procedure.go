package models

import (
	"time"

	"gorm.io/gorm"
)

type Procedure struct {
	gorm.Model
	PID                string     `gorm:"not null; unique; type:varchar(256)" json:"pid"`
	Type               uint       `gorm:"not null" json:"type"`
	Name               string     `gorm:"not null; type:varchar(100)" json:"name"`
	State              uint       `gorm:"not null" json:"state"`
	PHash              string     `gorm:"type:varchar(256)" json:"p_hash"`
	CheckCode          string     `gorm:"type:varchar(256)" json:"check_code"`
	SerialNumber       int64      `gorm:"not null" json:"serial_number"`
	Operator           string     `json:"operator"`
	StartTimestamp     time.Time  `gorm:"not null" json:"start_timestamp"`
	CompletedTimestamp *time.Time `json:"completed_timestamp"`
	PrePID             string     `gorm:"not null; type:varchar(256)" json:"pre_pid"`
	ICID               string     `gorm:"not null; type:varchar(256)" json:"icid"`
	BatchNumber        *string    `gorm:"type:varchar(256)" json:"batch_number"`
}
