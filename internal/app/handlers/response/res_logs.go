package response

import "CN-EU-FSIMS/internal/app/models"

type ResLogs struct {
	Logs  []models.LogInfo `json:"logs"`
	Count int64
}
