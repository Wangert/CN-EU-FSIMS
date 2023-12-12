package coldchain

import (
	"time"

	"gorm.io/gorm"
)

type TransportProcedureData struct {
	gorm.Model
	TID                string    `gorm:"not null; unique; type:varchar(256)" json:"t_id"`
	ProductNumber      string    `gorm:"not null; type:varchar(256)" json:"product_number"`
	CarNumber          string    `gorm:"not null; type:varchar(100)" json:"car_number"`
	Operator           string    `gorm:"not null; type:varchar(100)" json:"operator"`
	Temperature        string    `gorm:"not null; type:varchar(100)" json:"temperature"`
	Source             string    `gorm:"not null; type:varchar(100)" json:"source"`
	Destination        string    `gorm:"type:varchar(100)" json:"destination"`
	Humidity           string    `gorm:"not null; type:varchar(100)" json:"humidity"`
	LoadingTime        string    `gorm:"not null; type:varchar(100)" json:"loading_time"`
	UnloadingTime      string    `gorm:"type:varchar(100)" json:"unloading_time"`
	StartTimestamp     time.Time `gorm:"not null" json:"start_timestamp"`
	CompletedTimestamp time.Time `json:"completed_timestamp"`
}
