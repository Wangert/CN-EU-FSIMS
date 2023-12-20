package product

import (
	"gorm.io/gorm"
)

type Cow struct {
	gorm.Model
	Number      string  `gorm:"not null; unique; type:varchar(256)" json:"number"`
	Age         int     `gorm:"not null" json:"age"`
	Weight      float64 `gorm:"not null" json:"weight"`
	State       int     `gorm:"not null" json:"state"`
	HouseNumber string  `gorm:"not null; type:varchar(256)" json:"house_number"`
	BatchNumber *string `gorm:"type:varchar(256)" json:"batch_number"`
}

type CowInfo struct {
	Number      string  `json:"number"`
	Age         int     `json:"age"`
	Weight      float64 `json:"weight"`
	State       int     `json:"state"`
	HouseNumber string  `json:"house_number"`
	BatchNumber *string `json:"batch_number"`
}

func ToCowInfo(cow *Cow) CowInfo {
	return CowInfo{
		Number:      cow.Number,
		Age:         cow.Age,
		Weight:      cow.Weight,
		State:       cow.State,
		HouseNumber: cow.HouseNumber,
		BatchNumber: cow.BatchNumber,
	}
}
