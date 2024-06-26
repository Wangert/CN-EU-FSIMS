package warehouse

import (
	"time"

	"gorm.io/gorm"
)

//type SlaughterWareHouse struct {
//	gorm.Model
//	ProductNumber string    `gorm:"not null; unique; type:varchar(256)" json:"product_number"`
//	ProductPID    string    `gorm:"not null" json:"product_pid"`
//	ProductType   string    `gorm:"not null; type:varchar(100)" json:"product_type"`
//	State         uint      `gorm:"not null" json:"state"`
//	InOperator    string    `json:"in_operator"`
//	OutOperator   string    `json:"out_operator"`
//	Destination   string    `json:"destination"`
//	InTimestamp   time.Time `gorm:"not null" json:"in_timestamp"`
//	OutTimestamp  time.Time `json:"out_timestamp"`
//	HouseNumber   string    `gorm:"not null; type:varchar(256)" json:"house_number"`
//}

type SlaughterWarehouse struct {
	gorm.Model
	ProductNumber string     `gorm:"not null; unique; type:varchar(256)" json:"product_number"`
	PID           string     `gorm:"not null" json:"pid"`
	Type          string     `gorm:"type:varchar(100)" json:"type"`
	State         uint       `gorm:"not null" json:"state"`
	InOperator    string     `json:"in_operator"`
	OutOperator   string     `json:"out_operator"`
	Destination   string     `json:"destination"`
	InTimestamp   time.Time  `gorm:"not null" json:"in_timestamp"`
	OutTimestamp  *time.Time `json:"out_timestamp"`
	HouseNumber   string     `gorm:"not null; type:varchar(256)" json:"house_number"`
}

type SlaughterWarehouseInfo struct {
	ProductNumber string     `json:"product_number"`
	PID           string     `json:"pid"`
	Type          string     `json:"type"`
	State         uint       `json:"state"`
	InOperator    string     `json:"in_operator"`
	OutOperator   string     `json:"out_operator"`
	Destination   string     `json:"destination"`
	InTimestamp   time.Time  `json:"in_timestamp"`
	OutTimestamp  *time.Time `json:"out_timestamp"`
	HouseNumber   string     `json:"house_number"`
}

func ToSlaughterWarehouseInfo(warehouse *SlaughterWarehouse) SlaughterWarehouseInfo {
	return SlaughterWarehouseInfo{
		ProductNumber: warehouse.ProductNumber,
		PID:           warehouse.PID,
		Type:          warehouse.Type,
		State:         warehouse.State,
		InOperator:    warehouse.InOperator,
		OutOperator:   warehouse.OutOperator,
		Destination:   warehouse.Destination,
		InTimestamp:   warehouse.InTimestamp,
		OutTimestamp:  warehouse.OutTimestamp,
		HouseNumber:   warehouse.HouseNumber,
	}
}
