package warehouse

import (
	"time"

	"gorm.io/gorm"
)

type PackReceive struct {
	gorm.Model
	ProductNumber    string    `gorm:"not null; unique; type:varchar(256)" json:"product_number"`
	ProductPID       string    `gorm:"not null" json:"product_pid"`
	ProductType      string    `gorm:"not null; type:varchar(100)" json:"product_type"`
	Operator         string    `gorm:"not null; type:varchar(100)" json:"operator"`
	ReceiveTimestamp time.Time `gorm:"not null" json:"receive_timestamp"`
}
