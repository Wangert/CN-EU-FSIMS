package slaughter

import (
	"CN-EU-FSIMS/internal/app/models/product"
	"time"

	"gorm.io/gorm"
)

type SlaughterBatch struct {
	gorm.Model
	BatchNumber string                     `gorm:"not null; unique; type:varchar(256)" json:"batch_number"`
	HouseNumber string                     `gorm:"not null; type:varchar(256)" json:"house_number"`
	State       int                        `gorm:"not null" json:"state"`
	PID         string                     `gorm:"not null; unique; type:varchar(256)" json:"pid"`
	Worker      string                     `gorm:"not null; type:varchar(100)" json:"worker"`
	CowNumber   string                     `gorm:"not null; type:varchar(256)" json:"cow_number"`
	StartTime   *time.Time                 `json:"start_time"`
	EndTime     *time.Time                 `json:"end_time"`
	Products    []product.SlaughterProduct `gorm:"foreignKey:BatchNumber; references:BatchNumber" json:"products"`
	//Procedure   models.Procedure           `gorm:"foreignKey:BatchNumber; references:BatchNumber" json:"procedure"`
}

type SlaughterBatchInfo struct {
	BatchNumber string `json:"batch_number"`
	HouseNumber string `json:"house_number"`
	State       int    `json:"state"`
	PID         string `json:"pid"`
	Worker      string `json:"worker"`
	CowNumber   string `json:"cow_number"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
}

func ToSlaughterBatchInfo(batch *SlaughterBatch) SlaughterBatchInfo {
	startTime := ""
	loc, _ := time.LoadLocation("Asia/Shanghai")

	if batch.StartTime != nil {
		stime := batch.StartTime.In(loc)
		startTime = stime.Format("2006-01-02 15:04:05")
	}
	endTime := ""
	if batch.EndTime != nil {
		etime := batch.EndTime.In(loc)
		endTime = etime.Format("2006-01-02 15:04:05")
	}
	return SlaughterBatchInfo{
		BatchNumber: batch.BatchNumber,
		HouseNumber: batch.HouseNumber,
		State:       batch.State,
		PID:         batch.PID,
		Worker:      batch.Worker,
		CowNumber:   batch.CowNumber,
		StartTime:   startTime,
		EndTime:     endTime,
	}
}
