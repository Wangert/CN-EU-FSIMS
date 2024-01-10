package pack

import (
	"time"

	"gorm.io/gorm"
)

type PackageBatch struct {
	gorm.Model
	BatchNumber   string     `gorm:"not null; unique; type:varchar(256)" json:"batch_number"`
	HouseNumber   string     `gorm:"not null; type:varchar(256)" json:"house_number"`
	State         int        `gorm:"not null" json:"state"`
	PID           string     `gorm:"not null; type:varchar(256)" json:"pid"`
	Worker        string     `gorm:"not null; type:varchar(100)" json:"worker"`
	StartTime     *time.Time `json:"start_time"`
	EndTime       *time.Time `json:"end_time"`
	ProductNumber string     `gorm:"not null; type:varchar(256)" json:"product_number"`
	//Procedure     models.Procedure `gorm:"foreignKey:BatchNumber; references:BatchNumber" json:"procedure"`
}

type PackageBatchInfo struct {
	BatchNumber   string `json:"batch_number"`
	HouseNumber   string `json:"house_number"`
	State         int    `json:"state"`
	PID           string `json:"pid"`
	Worker        string `json:"worker"`
	ProductNumber string `json:"product_number"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
}

func ToPackageBatchInfo(batch *PackageBatch) PackageBatchInfo {
	startTime := ""
	if batch.EndTime != nil {
		startTime = batch.StartTime.Format("2006-01-02 15:04:05")
	}
	endTime := ""
	if batch.EndTime != nil {
		endTime = batch.EndTime.Format("2006-01-02 15:04:05")
	}
	return PackageBatchInfo{
		BatchNumber:   batch.BatchNumber,
		HouseNumber:   batch.HouseNumber,
		State:         batch.State,
		PID:           batch.PID,
		Worker:        batch.Worker,
		ProductNumber: batch.ProductNumber,
		StartTime:     startTime,
		EndTime:       endTime,
	}
}
