package models

import (
	"gorm.io/gorm"
	"time"
)

type MonitoringTimeRecord struct {
	gorm.Model
	IndexName string     `gorm:"not null; unique;" json:"index_name"`
	LastTime  *time.Time `json:"last_time"`
}
