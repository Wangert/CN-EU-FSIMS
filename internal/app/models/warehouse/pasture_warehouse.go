package warehouse

import (
	"time"

	"gorm.io/gorm"
)

type PastureWareHouse struct {
	gorm.Model
	ProductNumber string    `gorm:"not null; unique; type:varchar(256)" json:"product_number"`
	ProductPID    string    `gorm:"not null" json:"product_pid"`
	ProductType   string    `gorm:"not null; type:varchar(100)" json:"product_type"`
	State         uint      `gorm:"not null" json:"state"`
	InOperator    string    `json:"in_operator"`
	OutOperator   string    `json:"out_operator"`
	Destination   string    `json:"destination"`
	InTimestamp   time.Time `gorm:"not null" json:"in_timestamp"`
	OutTimestamp  time.Time `json:"out_timestamp"`
	HouseNumber   string    `gorm:"not null; type:varchar(256)" json:"house_number"`
}
