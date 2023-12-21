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
	HouseNumber   string    `gorm:"not null; type:varchar(256)" json:"house_number"`
}

type PackWarehouse struct {
	gorm.Model
	ProductNumber string     `gorm:"not null; unique; type:varchar(256)" json:"product_number"`
	PID           string     `gorm:"not null" json:"pid"`
	Type          string     `gorm:"not null; type:varchar(100)" json:"type"`
	State         uint       `gorm:"not null" json:"state"`
	InOperator    string     `json:"in_operator"`
	InTimestamp   time.Time  `gorm:"not null" json:"in_timestamp"`
	OutTimestamp  *time.Time `json:"out_timestamp"`
	PacNumber     string     `gorm:"not null; type:varchar(256)" json:"pac_number"`
}

type PackWarehouseInfo struct {
	ProductNumber string     `json:"product_number"`
	PID           string     `json:"pid"`
	Type          string     `json:"type"`
	State         uint       `json:"state"`
	InOperator    string     `json:"in_operator"`
	InTimestamp   time.Time  `json:"in_timestamp"`
	OutTimestamp  *time.Time `json:"out_timestamp"`
	PacNumber     string     `json:"pac_number"`
}

func ToPackWarehouseInfo(warehouse PackWarehouse) PackWarehouseInfo {
	return PackWarehouseInfo{
		ProductNumber: warehouse.ProductNumber,
		PID:           warehouse.PID,
		Type:          warehouse.Type,
		State:         warehouse.State,
		InOperator:    warehouse.InOperator,
		InTimestamp:   warehouse.InTimestamp,
		OutTimestamp:  warehouse.OutTimestamp,
		PacNumber:     warehouse.PacNumber,
	}
}
