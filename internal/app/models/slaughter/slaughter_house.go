package slaughter

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/warehouse"
	"gorm.io/gorm"
)

type SlaughterHouse struct {
	gorm.Model
	HouseNumber      string                             `gorm:"not null; unique; type:varchar(256)" json:"house_number"`
	Name             string                             `gorm:"not null; type:varchar(100)" json:"name"`
	Address          string                             `gorm:"not null; type:varchar(256)" json:"address"`
	State            uint                               `gorm:"not null" json:"state"`
	LegalPerson      string                             `gorm:"not null; type:varchar(100)" json:"legal_person"`
	ReceiveRecords   []warehouse.SlaughterReceiveRecord `gorm:"foreignKey:SlaughterNumber; references:HouseNumber" json:"receive_records"`
	SlaughterRecords []SlaughterBatch                   `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"slaughter_records"`
	SWRecords        []warehouse.SlaughterWarehouse     `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"sw_records"`
}

type SlaughterHouseInfo struct {
	HouseNumber string `json:"house_number"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	State       uint   `json:"state"`
	LegalPerson string `json:"legal_person"`
}

func ToSlaughterHouseInfo(house *SlaughterHouse) SlaughterHouseInfo {
	return SlaughterHouseInfo{
		HouseNumber: house.HouseNumber,
		Name:        house.Name,
		Address:     house.Address,
		State:       house.State,
		LegalPerson: house.LegalPerson,
	}
}

func SlaughterHouseToRes(sh *SlaughterHouse) models.House {
	return models.House{
		HouseNumber: sh.HouseNumber,
		Name:        sh.Name,
		Address:     sh.Address,
		State:       sh.State,
		LegalPerson: sh.LegalPerson,
	}
}
