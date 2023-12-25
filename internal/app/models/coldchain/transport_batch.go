package coldchain

import (
	"CN-EU-FSIMS/internal/app/models/product"
	"gorm.io/gorm"
)

type TransportBatch struct {
	gorm.Model
	BatchNumber string                   `gorm:"not null; unique; type:varchar(256)" json:"batch_number"`
	TVNumber    string                   `gorm:"not null; type:varchar(256)" json:"tv_number"`
	State       int                      `gorm:"not null" json:"state"`
	PID         string                   `gorm:"not null; type:varchar(256)" json:"pid"`
	Worker      string                   `gorm:"not null; type:varchar(100)" json:"worker"`
	Products    []product.PackageProduct `gorm:"foreignKey:TransportBatchNumber; references:BatchNumber" json:"products"`
}

type TransportBatchInfo struct {
	BatchNumber string                   `json:"batch_number"`
	TVNumber    string                   `json:"tv_number"`
	State       int                      `json:"state"`
	PID         string                   `json:"pid"`
	Worker      string                   `json:"worker"`
	Products    []product.PackageProduct `json:"products"`
}

func ToTransportBatchInfo(batch *TransportBatch) TransportBatchInfo {
	return TransportBatchInfo{
		BatchNumber: batch.BatchNumber,
		TVNumber:    batch.TVNumber,
		State:       batch.State,
		PID:         batch.PID,
		Worker:      batch.Worker,
		Products:    batch.Products,
	}
}
