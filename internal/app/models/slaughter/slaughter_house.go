package slaughter

import (
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/app/models/warehouse"
	"gorm.io/gorm"
)

type SlaughterHouse struct {
	gorm.Model
	HouseNumber string                       `gorm:"not null; unique; type:varchar(256)" json:"house_number"`
	Name        string                       `gorm:"not null; type:varchar(100)" json:"name"`
	Address     string                       `gorm:"not null; type:varchar(256)" json:"address"`
	State       uint                         `gorm:"not null" json:"state"`
	LegalPerson string                       `gorm:"not null; type:varchar(100)" json:"legal_person"`
	SWRecord    warehouse.SlaughterWareHouse `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"sw_record"`
}

func SlaughterHouseToRes(sh *SlaughterHouse) response.ResHouse {
	return response.ResHouse{
		HouseNumber: sh.HouseNumber,
		Name:        sh.Name,
		Address:     sh.Address,
		State:       sh.State,
		LegalPerson: sh.LegalPerson,
	}
}
