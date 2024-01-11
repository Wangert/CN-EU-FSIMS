package models

import (
	"time"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	TimeStamp time.Time `gorm:"not null"`
	UUID      string    `gorm:"not null"`
	Account   string    `gorm:"not null"`
	Type      int       `json:"type"`
	Action    string    `gorm:"not null"`
}

type LogInfo struct {
	TimeStamp time.Time `gorm:"not null"`
	UUID      string    `gorm:"not null"`
	Account   string    `gorm:"not null"`
	Type      int       `json:"type"`
	Action    string    `gorm:"not null"`
}

func ToLogInfo(l *Log) LogInfo {
	return LogInfo{
		TimeStamp: l.TimeStamp,
		UUID:      l.UUID,
		Account:   l.Account,
		Type:      l.Type,
		Action:    l.Action,
	}
}
