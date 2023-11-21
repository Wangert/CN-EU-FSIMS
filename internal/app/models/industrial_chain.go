package models

import (
	"time"
)

type IndustrialChainTest struct {
	//gorm.Model
	ICID               string    `gorm:"not null; unique; index; type:varchar(256)" json:"icid"`
	Type               string    `json:"type"`
	State              uint      `gorm:"not null" json:"state"`
	StartTimestamp     time.Time `gorm:"not null" json:"start_timestamp"`
	CompletedTimestamp time.Time `json:"completed_timestamp"`
}
