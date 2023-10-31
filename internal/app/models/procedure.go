package models

import (
	"gorm.io/gorm"
	"time"
)

const (
	// PASTURE procedure is 0
	PASTURE = iota
	// FATTEN procedure is 1
	FATTEN
	// PREMORTEM_MANAGEMENT procedure is 2
	PREMORTEM_MANAGEMENT
	// SLAUGHTER procedure is 3
	SLAUGHTER
	// PACKAGE procedure is 4
	PACKAGE
	// COLDCHAIN_TRANSPORT procedure is 5
	COLDCHAIN_TRANSPORT
	// SELL procedure is 6
	SELL
)

type Procedure struct {
	gorm.Model
	PID                string         `gorm:"not null; unique; type:varchar(256)" json:"pid"`
	Type               uint           `gorm:"not null" json:"type"`
	Name               string         `gorm:"not null; type:varchar(100)" json:"name"`
	SerialNumber       uint           `gorm:"not null" json:"serial_number"`
	StartTimestamp     time.Time      `gorm:"not null" json:"start_timestamp"`
	CompletedTimestamp time.Time      `json:"completed_timestamp"`
	PrePID             string         `gorm:"not null; type:varchar(256)" json:"pre_pid"`
	ICID               string         `gorm:"not null; type:varchar(256)" json:"icid"`
	SubProcedures      []SubProcedure `gorm:"foreignKey:PID; references:PID" json:"sub_procedures"`
}
