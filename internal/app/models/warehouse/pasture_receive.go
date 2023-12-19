package warehouse

import (
	"gorm.io/gorm"
	"time"
)

type PastureReceive struct {
	gorm.Model
	CowNumber        string    `gorm:"not null; unique; type:varchar(256)" json:"product_number"`
	PID              string    `gorm:"not null" json:"pid"`
	ProductType      string    `gorm:"not null; type:varchar(100)" json:"product_type"`
	Operator         string    `gorm:"not null; type:varchar(100)" json:"operator"`
	ReceiveTimestamp time.Time `gorm:"not null" json:"receive_timestamp"`
}
