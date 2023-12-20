package pasture

import (
	"CN-EU-FSIMS/internal/app/models/product"
	"gorm.io/gorm"
)

type FeedingBatch struct {
	gorm.Model
	BatchNumber string        `gorm:"not null; unique; type:varchar(256)" json:"batch_number"`
	HouseNumber string        `gorm:"not null; type:varchar(256)" json:"house_number"`
	State       int           `gorm:"not null" json:"state"`
	PID         string        `gorm:"not null; type:varchar(256)" json:"pid"`
	Worker      string        `gorm:"not null; type:varchar(100)" json:"worker"`
	Cows        []product.Cow `gorm:"foreignKey:BatchNumber; references:BatchNumber" json:"cows"`
}

type FeedingBatchInfo struct {
	BatchNumber string        `json:"batch_number"`
	HouseNumber string        `json:"house_number"`
	State       int           `json:"state"`
	PID         string        `json:"pid"`
	Worker      string        `json:"worker"`
	Cows        []product.Cow `json:"cows"`
}

func ToFeedingBatchInfo(batch *FeedingBatch) FeedingBatchInfo {
	return FeedingBatchInfo{
		BatchNumber: batch.BatchNumber,
		HouseNumber: batch.HouseNumber,
		State:       batch.State,
		PID:         batch.PID,
		Worker:      batch.Worker,
		Cows:        batch.Cows,
	}
}
