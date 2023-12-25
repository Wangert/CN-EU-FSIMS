package sell

import (
	"CN-EU-FSIMS/internal/app/models/product"
	"gorm.io/gorm"
)

type Mall struct {
	gorm.Model
	Number      string             `gorm:"not null; unique; type:varchar(256)" json:"house_number"`
	Name        string             `gorm:"not null; type:varchar(100)" json:"name"`
	Address     string             `gorm:"not null; type:varchar(256)" json:"address"`
	State       uint               `gorm:"not null" json:"state"`
	LegalPerson string             `gorm:"not null; type:varchar(100)" json:"legal_person"`
	Goods       []product.MallGood `gorm:"foreignKey:MallNumber; references:Number" json:"goods"`
}

type MallInfo struct {
	Number      string             `json:"house_number"`
	Name        string             `json:"name"`
	Address     string             `json:"address"`
	State       uint               `json:"state"`
	LegalPerson string             `json:"legal_person"`
	Goods       []product.MallGood `json:"goods"`
}

func ToMallInfo(mall *Mall) MallInfo {
	return MallInfo{
		Number:      mall.Number,
		Name:        mall.Name,
		Address:     mall.Address,
		State:       mall.State,
		LegalPerson: mall.LegalPerson,
		Goods:       mall.Goods,
	}
}
