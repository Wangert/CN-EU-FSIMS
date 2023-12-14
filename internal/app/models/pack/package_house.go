package pack

import (
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/app/models/warehouse"
	"gorm.io/gorm"
)

type PackageHouse struct {
	gorm.Model
	HouseNumber string                  `gorm:"not null; unique; type:varchar(256)" json:"house_number"`
	Name        string                  `gorm:"not null; type:varchar(100)" json:"name"`
	Address     string                  `gorm:"not null; type:varchar(256)" json:"address"`
	State       uint                    `gorm:"not null" json:"state"`
	LegalPerson string                  `gorm:"not null; type:varchar(100)" json:"legal_person"`
	PWRecord    warehouse.PackWareHouse `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"pw_record"`
}

func PackageHouseToRes(ph *PackageHouse) response.ResHouse {
	return response.ResHouse{
		HouseNumber: ph.HouseNumber,
		Name:        ph.Name,
		Address:     ph.Address,
		State:       ph.State,
		LegalPerson: ph.LegalPerson,
	}
}
