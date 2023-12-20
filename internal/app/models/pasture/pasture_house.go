package pasture

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/product"
	"CN-EU-FSIMS/internal/app/models/warehouse"
	"gorm.io/gorm"
)

type PastureHouse struct {
	gorm.Model
	HouseNumber   string                       `gorm:"not null; unique; type:varchar(256)" json:"house_number"`
	Name          string                       `gorm:"not null; type:varchar(100)" json:"name"`
	Address       string                       `gorm:"not null; type:varchar(256)" json:"address"`
	State         uint                         `gorm:"not null" json:"state"`
	LegalPerson   string                       `gorm:"not null; type:varchar(100)" json:"legal_person"`
	Cows          []product.Cow                `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"cows"`
	FeedingRecord []FeedingBatch               `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"feeding_record"`
	PasHRecord    []warehouse.PastureWareHouse `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"pash_record"`
}

func PastureToRes(pas *PastureHouse) models.House {
	return models.House{
		HouseNumber: pas.HouseNumber,
		Name:        pas.Name,
		Address:     pas.Address,
		State:       pas.State,
		LegalPerson: pas.LegalPerson,
	}
}
