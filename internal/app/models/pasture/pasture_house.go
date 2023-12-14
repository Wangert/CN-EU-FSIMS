package pasture

import (
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/app/models/warehouse"
	"gorm.io/gorm"
)

type PastureHouse struct {
	gorm.Model
	HouseNumber string                     `gorm:"not null; unique; type:varchar(256)" json:"house_number"`
	Name        string                     `gorm:"not null; type:varchar(100)" json:"name"`
	Address     string                     `gorm:"not null; type:varchar(256)" json:"address"`
	State       uint                       `gorm:"not null" json:"state"`
	LegalPerson string                     `gorm:"not null; type:varchar(100)" json:"legal_person"`
	PasHRecord  warehouse.PastureWareHouse `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"pash_record"`
}

func PastureToRes(pas *PastureHouse) response.ResHouse {
	return response.ResHouse{
		HouseNumber: pas.HouseNumber,
		Name:        pas.Name,
		Address:     pas.Address,
		State:       pas.State,
		LegalPerson: pas.LegalPerson,
	}
}
