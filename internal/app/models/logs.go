package models

import (
	"CN-EU-FSIMS/internal/app/handlers/response"
	"time"
)

type Logs struct {
	TimeStamp time.Time `gorm:"not null"`
	UUID      string    `gorm:"not null"`
	Account   string    `gorm:"not null"`
	Type      int       `json:"type"`
	Action    string    `gorm:"not null"`
}

func FsimsLogsToReslogs(logs *Logs) response.ResLogs {
	return response.ResLogs{
		TimeStamp: logs.TimeStamp,
		UUID:      logs.UUID,
		Account:   logs.Account,
		Type:      logs.Type,
		Action:    logs.Action,
	}
}