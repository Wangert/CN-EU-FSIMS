package pasture

import (
	"CN-EU-FSIMS/internal/app/models/product"
	"time"

	"gorm.io/gorm"
)

type FeedingBatch struct {
	gorm.Model
	BatchNumber string        `gorm:"not null; unique; type:varchar(256)" json:"batch_number"`
	HouseNumber string        `gorm:"not null; type:varchar(256)" json:"house_number"`
	State       int           `gorm:"not null" json:"state"`
	PID         string        `gorm:"not null; unique; type:varchar(256)" json:"pid"`
	Worker      string        `gorm:"not null; type:varchar(100)" json:"worker"`
	StartTime   *time.Time    `json:"start_time"`
	EndTime     *time.Time    `json:"end_time"`
	Cows        []product.Cow `gorm:"foreignKey:BatchNumber; references:BatchNumber" json:"cows"`
	//Procedure   models.Procedure `gorm:"foreignKey:BatchNumber; references:BatchNumber" json:"procedure"`
}

type FeedingBatchInfo struct {
	BatchNumber string            `json:"batch_number"`
	HouseNumber string            `json:"house_number"`
	State       int               `json:"state"`
	PID         string            `json:"pid"`
	Worker      string            `json:"worker"`
	Cows        []product.CowInfo `json:"cows"`
	StartTime   string            `json:"start_time"`
	EndTime     string            `json:"end_time"`
}

func ToFeedingBatchInfo(batch *FeedingBatch) FeedingBatchInfo {
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
	var cows []product.CowInfo
	for i, _ := range batch.Cows {
		// records[i] = pasture.ToFeedingBatchInfo(fb)
		cows = append(cows, product.ToCowInfo(&batch.Cows[i]))
	}

	return FeedingBatchInfo{
		BatchNumber: batch.BatchNumber,
		HouseNumber: batch.HouseNumber,
		State:       batch.State,
		PID:         batch.PID,
		Worker:      batch.Worker,
		Cows:        cows,
		StartTime:   startTime,
		EndTime:     endTime,
	}
}
