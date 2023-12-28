package coldchain

import (
	"gorm.io/gorm"
)

type TransportProcedure struct {
	gorm.Model
	TransPID string `gorm:"not null; unique; type:varchar(256)" json:"trans_pid"`
}

type TransportProcedureData struct {
	Temperature string `json:"temperature"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Humidity    string `json:"humidity"`
}
