package response

import "CN-EU-FSIMS/internal/app/models"

// import "CN-EU-FSIMS/internal/app/models"

type ResNotifications struct {
	Notifications []models.ResNotification `json:"notifications"`
	Count         int64                    `json:"count"`
}
