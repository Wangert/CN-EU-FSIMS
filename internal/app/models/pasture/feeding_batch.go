package pasture

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/product"
	"gorm.io/gorm"
	"time"
)

type FeedingBatch struct {
	gorm.Model
	BatchNumber string           `gorm:"not null; unique; type:varchar(256)" json:"batch_number"`
	HouseNumber string           `gorm:"not null; type:varchar(256)" json:"house_number"`
	State       int              `gorm:"not null" json:"state"`
	PID         string           `gorm:"not null; type:varchar(256)" json:"pid"`
	Worker      string           `gorm:"not null; type:varchar(100)" json:"worker"`
	StartTime   *time.Time       `json:"start_time"`
	EndTime     *time.Time       `json:"end_time"`
	Cows        []product.Cow    `gorm:"foreignKey:BatchNumber; references:BatchNumber" json:"cows"`
	Procedure   models.Procedure `gorm:"foreignKey:BatchNumber; references:BatchNumber" json:"procedure"`
}

type FeedingBatchInfo struct {
	BatchNumber string        `json:"batch_number"`
	HouseNumber string        `json:"house_number"`
	State       int           `json:"state"`
	PID         string        `json:"pid"`
	Worker      string        `json:"worker"`
	Cows        []product.Cow `json:"cows"`
	StartTime   string        `json:"start_time"`
	EndTime     string        `json:"end_time"`
}

func ToFeedingBatchInfo(batch *FeedingBatch) FeedingBatchInfo {
	return FeedingBatchInfo{
		BatchNumber: batch.BatchNumber,
		HouseNumber: batch.HouseNumber,
		State:       batch.State,
		PID:         batch.PID,
		Worker:      batch.Worker,
		Cows:        batch.Cows,
		StartTime:   batch.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:     batch.EndTime.Format("2006-01-02 15:04:05"),
	}
}
