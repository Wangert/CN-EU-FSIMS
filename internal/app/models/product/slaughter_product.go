package product

import "gorm.io/gorm"

type SlaughterProduct struct {
	gorm.Model
	Number      string  `gorm:"not null; unique; type:varchar(256)" json:"number"`
	Type        int     `gorm:"not null" json:"type"`
	TypeName    string  `gorm:"not null" json:"type_name"`
	Weight      float64 `gorm:"not null" json:"weight"`
	State       int     `gorm:"not null" json:"state"`
	HouseNumber string  `gorm:"not null; type:varchar(256)" json:"house_number"`
	BatchNumber string  `gorm:"not null; type:varchar(256)" json:"batch_number"`
}

type SlaughterProductInfo struct {
	Number      string  `json:"number"`
	Type        int     `json:"type"`
	TypeName    string  `json:"type_name"`
	Weight      float64 `json:"weight"`
	State       int     `json:"state"`
	HouseNumber string  `json:"house_number"`
	BatchNumber string  `json:"batch_number"`
}

func ToSlaughterProductInfo(product *SlaughterProduct) SlaughterProductInfo {
	return SlaughterProductInfo{
		Number:      product.Number,
		Type:        product.Type,
		TypeName:    product.TypeName,
		Weight:      product.Weight,
		State:       product.State,
		HouseNumber: product.HouseNumber,
		BatchNumber: product.BatchNumber,
	}
}
