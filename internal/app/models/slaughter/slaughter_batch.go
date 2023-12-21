package slaughter

import (
	"CN-EU-FSIMS/internal/app/models/product"
	"gorm.io/gorm"
)

type SlaughterBatch struct {
	gorm.Model
	BatchNumber string                     `gorm:"not null; unique; type:varchar(256)" json:"batch_number"`
	HouseNumber string                     `gorm:"not null; type:varchar(256)" json:"house_number"`
	State       int                        `gorm:"not null" json:"state"`
	PID         string                     `gorm:"not null; type:varchar(256)" json:"pid"`
	Worker      string                     `gorm:"not null; type:varchar(100)" json:"worker"`
	CowNumber   string                     `gorm:"not null; type:varchar(256)" json:"cow_number"`
	Products    []product.SlaughterProduct `gorm:"foreignKey:BatchNumber; references:BatchNumber" json:"products"`
}

type SlaughterBatchInfo struct {
	BatchNumber string `json:"batch_number"`
	HouseNumber string `json:"house_number"`
	State       int    `json:"state"`
	PID         string `json:"pid"`
	Worker      string `json:"worker"`
	CowNumber   string `json:"cow_number"`
}

func ToSlaughterBatchInfo(batch *SlaughterBatch) SlaughterBatchInfo {
	return SlaughterBatchInfo{
		BatchNumber: batch.BatchNumber,
		HouseNumber: batch.HouseNumber,
		State:       batch.State,
		PID:         batch.PID,
		Worker:      batch.Worker,
		CowNumber:   batch.CowNumber,
	}
}
