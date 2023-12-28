package pasture

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/product"
	"CN-EU-FSIMS/internal/app/models/warehouse"
	"gorm.io/gorm"
)

type PastureHouse struct {
	gorm.Model
	HouseNumber             string                       `gorm:"not null; unique; type:varchar(256)" json:"house_number"`
	Name                    string                       `gorm:"not null; type:varchar(100)" json:"name"`
	Address                 string                       `gorm:"not null; type:varchar(256)" json:"address"`
	State                   uint                         `gorm:"not null" json:"state"`
	LegalPerson             string                       `gorm:"not null; type:varchar(100)" json:"legal_person"`
	Cows                    []product.Cow                `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"cows"`
	FeedingRecord           []FeedingBatch               `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"feeding_record"`
	PasHRecord              []warehouse.PastureWarehouse `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"pas_h_record"`
	HeavyMetalRecords       []PastureFeedHeavyMetal      `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"heavy_metal_records"`
	PastureAreaRecords      []PastureArea                `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"pasture_area_records"`
	WaterRecords            []PastureWaterRecord         `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"water_records"`
	BufferRecords           []PastureBuffer              `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"buffer_records"`
	CowHouseRecords         []CowHouse                   `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"cow_house_records"`
	BasicEnvironment        []PastureBasicEnvironment    `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"basic_environment"`
	PaddingRecords          []PasturePaddingRequire      `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"padding_records"`
	WastedWaterIndexRecords []PastureWastedWaterIndex    `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"wasted_water_index_records"`
	DisinfectionRecords     []PastureDisinfectionRecord  `gorm:"foreignKey:HouseNumber; references:HouseNumber" json:"disinfection_records"`
}

type PastureHouseInfo struct {
	HouseNumber string `json:"house_number"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	State       uint   `json:"state"`
	LegalPerson string `json:"legal_person"`
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

func ToPastureHouseInfo(house PastureHouse) PastureHouseInfo {
	return PastureHouseInfo{
		HouseNumber: house.HouseNumber,
		Name:        house.Name,
		Address:     house.Address,
		State:       house.State,
		LegalPerson: house.LegalPerson,
	}
}
