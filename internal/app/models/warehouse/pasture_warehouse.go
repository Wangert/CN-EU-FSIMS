package warehouse

import (
	"time"

	"gorm.io/gorm"
)

type PastureWarehouse struct {
	gorm.Model
	CowNumber    string     `gorm:"not null; unique; type:varchar(256)" json:"cow_number"`
	PID          string     `gorm:"not null" json:"pid"`
	Type         string     `gorm:"type:varchar(100)" json:"type"`
	State        uint       `gorm:"not null" json:"state"`
	InOperator   string     `json:"in_operator"`
	OutOperator  string     `json:"out_operator"`
	Destination  string     `json:"destination"`
	InTimestamp  time.Time  `gorm:"not null" json:"in_timestamp"`
	OutTimestamp *time.Time `json:"out_timestamp"`
	HouseNumber  string     `gorm:"not null; type:varchar(256)" json:"house_number"`
}

type PastureWarehouseInfo struct {
	CowNumber    string     `json:"cow_number"`
	PID          string     `son:"pid"`
	Type         string     `json:"type"`
	State        uint       `json:"state"`
	InOperator   string     `json:"in_operator"`
	OutOperator  string     `json:"out_operator"`
	Destination  string     `json:"destination"`
	InTimestamp  time.Time  `gorm:"not null" json:"in_timestamp"`
	OutTimestamp *time.Time `json:"out_timestamp"`
}

func ToPastureWarehouseInfo(warehouse PastureWarehouse) PastureWarehouseInfo {
	return PastureWarehouseInfo{
		CowNumber:    warehouse.CowNumber,
		PID:          warehouse.PID,
		Type:         warehouse.Type,
		State:        warehouse.State,
		InOperator:   warehouse.InOperator,
		OutOperator:  warehouse.OutOperator,
		Destination:  warehouse.Destination,
		InTimestamp:  warehouse.InTimestamp,
		OutTimestamp: warehouse.OutTimestamp,
	}
}
