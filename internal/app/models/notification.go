package models

import (
	"time"

	"gorm.io/gorm"
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

type NotificationInfo struct {
	ID         int    `json:"id"`
	Source     string `gorm:"not null" json:"source"`
	Content    string `gorm:"not null" json:"content"`
	UUID       string `gorm:"not null" json:"UUID"`
	EventTime  string `gorm:"not null" json:"event_time"`
	NoticeTime string `json:"notice_time"`
	EventType  int    `gorm:"not null" json:"event_type"`
	Affected   string `gorm:"not null" json:"affected"`
	Proposal   string `gorm:"not null" json:"proposal"`
	RiskLevel  int    `gorm:"not null" json:"risk_level"`
	State      int    `gorm:"not null" json:"state"`
}

func ToNotificationInfo(data *Notification) NotificationInfo {
	return NotificationInfo{
		ID:         int(data.ID),
		Source:     data.Source,
		Content:    data.Content,
		UUID:       data.UUID,
		EventTime:  data.EventTime.Format("2006-01-02 15:04:05"),
		NoticeTime: data.CreatedAt.Format("2006-01-02 15:04:05"),
		EventType:  data.EventType,
		Affected:   data.Affected,
		Proposal:   data.Proposal,
		RiskLevel:  data.RiskLevel,
		State:      data.State,
	}
}

type ResNotification struct {
	// Notifications []models.NotificationInfo `json:"notifications"`
	// Count         int64                     `json:"count"`
	ID         int     `json:"id"`
	SourceName string  `json:"source_name"`
	Content    string  `json:"content"`
	UUID       string  `json:"UUID"`
	EventTime  string  `json:"event_time"`
	NoticeTime string  `json:"notice_time"`
	EventType  int     `json:"event_type"`
	Affected   []Batch `json:"affected"`
	Proposal   string  `json:"proposal"`
	RiskLevel  int     `json:"risk_level"`
	State      int     `json:"state"`
}

type Batch struct {
	BatchNumber string `json:"batch_number"`
	State       int    `json:"state"`
}
