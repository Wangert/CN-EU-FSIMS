package warehouse

import (
	"time"

	"gorm.io/gorm"
)

type PackWareHouse struct {
	gorm.Model
	ProductNumber string    `gorm:"not null; unique; type:varchar(256)" json:"product_number"`
	ProductPID    string    `gorm:"not null" json:"product_pid"`
	ProductType   string    `gorm:"not null; type:varchar(100)" json:"product_type"`
	State         uint      `gorm:"not null" json:"state"`
	InOperator    string    `json:"in_operator"`
	InTimestamp   time.Time `gorm:"not null" json:"in_timestamp"`
	OutTimestamp  time.Time `json:"out_timestamp"`
}
