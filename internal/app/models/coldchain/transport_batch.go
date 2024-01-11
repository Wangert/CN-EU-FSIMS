package coldchain

import (
	"CN-EU-FSIMS/internal/app/models/product"
	"time"

	"gorm.io/gorm"
)

type TransportBatch struct {
	gorm.Model
	BatchNumber string                   `gorm:"not null; unique; type:varchar(256)" json:"batch_number"`
	TVNumber    string                   `gorm:"not null; type:varchar(256)" json:"tv_number"`
	State       int                      `gorm:"not null" json:"state"`
	Worker      string                   `gorm:"not null; type:varchar(100)" json:"worker"`
	MallNumber  string                   `gorm:"not null; type:varchar(256)" json:"mall_number"`
	StartTime   *time.Time               `json:"start_time"`
	EndTime     *time.Time               `json:"end_time"`
	Products    []product.PackageProduct `gorm:"foreignKey:TransportBatchNumber; references:BatchNumber" json:"products"`
	//Procedure   models.Procedure         `gorm:"foreignKey:BatchNumber; references:BatchNumber" json:"procedure"`
}

type TransportBatchInfo struct {
	BatchNumber string                   `json:"batch_number"`
	TVNumber    string                   `json:"tv_number"`
	State       int                      `json:"state"`
	Worker      string                   `json:"worker"`
	MallNumber  string                   `json:"mall_number"`
	Products    []product.PackageProduct `json:"products"`
	StartTime   string                   `json:"start_time"`
	EndTime     string                   `json:"end_time"`
}

func ToTransportBatchInfo(batch *TransportBatch) TransportBatchInfo {
	startTime := ""
	if batch.EndTime != nil {
		startTime = batch.StartTime.Format("2006-01-02 15:04:05")
	}
	endTime := ""
	if batch.EndTime != nil {
		endTime = batch.EndTime.Format("2006-01-02 15:04:05")
	}
	return TransportBatchInfo{
		BatchNumber: batch.BatchNumber,
		TVNumber:    batch.TVNumber,
		State:       batch.State,
		Worker:      batch.Worker,
		MallNumber:  batch.MallNumber,
		Products:    batch.Products,
		StartTime:   startTime,
		EndTime:     endTime,
	}
}
