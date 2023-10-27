package models

import (
	"gorm.io/gorm"
	"time"
)

type IndustrialChain struct {
	gorm.Model
	ICID               string      `gorm:"not null; unique; index; type:varchar(256)" json:"icid"`
	Type               string      `json:"type"`
	State              uint        `gorm:"not null" json:"state"`
	StartTimestamp     time.Time   `gorm:"not null" json:"start_timestamp"`
	CompletedTimestamp time.Time   `json:"completed_timestamp"`
	Procedures         []Procedure `gorm:"foreignKey:ICID; references:ICID" json:"procedures"`
}
