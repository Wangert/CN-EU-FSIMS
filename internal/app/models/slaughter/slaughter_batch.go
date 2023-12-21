package slaughter

import (
	"gorm.io/gorm"
)

type SlaughterBatch struct {
	gorm.Model
	BatchNumber string `gorm:"not null; unique; type:varchar(256)" json:"batch_number"`
	HouseNumber string `gorm:"not null; type:varchar(256)" json:"house_number"`
	State       int    `gorm:"not null" json:"state"`
	PID         string `gorm:"not null; type:varchar(256)" json:"pid"`
	Worker      string `gorm:"not null; type:varchar(100)" json:"worker"`
	CowNumber   string `gorm:"not null; type:varchar(256)" json:"cow_number"`
}
