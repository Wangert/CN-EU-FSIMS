package response

import "CN-EU-FSIMS/internal/app/models"

type ResNotification struct {
	Notifications []models.NotificationInfo `json:"notifications"`
	Count         int64                     `json:"count"`
}
