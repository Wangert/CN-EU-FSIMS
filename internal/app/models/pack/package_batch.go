package pack

import "gorm.io/gorm"

type PackageBatch struct {
	gorm.Model
	BatchNumber   string `gorm:"not null; unique; type:varchar(256)" json:"batch_number"`
	HouseNumber   string `gorm:"not null; type:varchar(256)" json:"house_number"`
	State         int    `gorm:"not null" json:"state"`
	PID           string `gorm:"not null; type:varchar(256)" json:"pid"`
	Worker        string `gorm:"not null; type:varchar(100)" json:"worker"`
	ProductNumber string `gorm:"not null; type:varchar(256)" json:"product_number"`
}
