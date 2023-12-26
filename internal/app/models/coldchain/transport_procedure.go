package coldchain

import (
	"gorm.io/gorm"
)

type TransportProcedure struct {
	gorm.Model
	TransPID string `gorm:"not null; unique; type:varchar(256)" json:"trans_pid"`
}

type TransportProcedureData struct {
	Temperature string `gorm:"not null; type:varchar(100)" json:"temperature"`
	Source      string `gorm:"not null; type:varchar(100)" json:"source"`
	Destination string `gorm:"type:varchar(100)" json:"destination"`
	Humidity    string `gorm:"not null; type:varchar(100)" json:"humidity"`
}
