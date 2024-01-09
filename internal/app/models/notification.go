package models

import (
	"gorm.io/gorm"
	"time"
)

type Notification struct {
	gorm.Model
	Source    string    `gorm:"not null" json:"source"`
	Content   string    `gorm:"not null" json:"content"`
	UUID      string    `gorm:"not null" json:"UUID"`
	EventTime time.Time `gorm:"not null" json:"event_time"`
	EventType int       `gorm:"not null" json:"event_type"`
	Affected  string    `gorm:"not null" json:"affected"`
	Proposal  string    `gorm:"not null" json:"proposal"`
	RiskLevel int       `gorm:"not null" json:"risk_level"`
	State     int       `gorm:"not null" json:"state"`
}
