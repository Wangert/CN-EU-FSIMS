package models

import "time"

type Logs struct {
	TimeStamp time.Time `gorm:"not null"`
	UUID      string    `gorm:"not null"`
	Account   string    `gorm:"not null"`
	Type      int       `json:"type"`
	Action    string    `gorm:"not null"`
}
