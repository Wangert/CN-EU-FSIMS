package product

import (
	"gorm.io/gorm"
	"time"
)

type MallGood struct {
	gorm.Model
	Number         string    `gorm:"not null; unique; type:varchar(256)" json:"number"`
	Type           int       `gorm:"not null" json:"type"`
	TypeName       string    `gorm:"not null" json:"type_name"`
	PackMethod     int       `gorm:"not null" json:"pack_method"`
	PackMethodName string    `gorm:"not null" json:"pack_method_name"`
	ProductionDate time.Time `gorm:"not null" json:"production_date"`
	ShelfLife      string    `gorm:"not null" json:"shelf_life"`
	Weight         float64   `gorm:"not null" json:"weight"`
	State          int       `gorm:"not null" json:"state"`
	MallNumber     string    `gorm:"not null; type:varchar(256)" json:"mall_number"`
	FinalPID       *string   `gorm:"type:varchar(256)" json:"final_pid"`
	Checkcode      string    `gorm:"not null; type:varchar(256)" json:"checkcode"`
	BuyerIDCard    *string   `gorm:"type:varchar(256)" json:"buyer_id_card"`
}

type MallGoodInfo struct {
	Number         string    `json:"number"`
	TypeName       string    `json:"type_name"`
	PackMethodName string    `json:"pack_method_name"`
	ProductionDate time.Time `json:"production_date"`
	ShelfLife      string    `json:"shelf_life"`
	Weight         float64   `json:"weight"`
	State          int       `json:"state"`
	MallNumber     string    `json:"mall_number"`
	FinalPID       *string   `json:"final_pid"`
	Checkcode      string    `json:"checkcode"`
	BuyerIDCard    *string   `json:"buyer_id_card"`
}

func ToMallGoodInfo(good *MallGood) MallGoodInfo {
	return MallGoodInfo{
		Number:         good.Number,
		TypeName:       good.TypeName,
		PackMethodName: good.PackMethodName,
		ProductionDate: good.ProductionDate,
		ShelfLife:      good.ShelfLife,
		Weight:         good.Weight,
		State:          good.State,
		MallNumber:     good.MallNumber,
		FinalPID:       good.FinalPID,
		Checkcode:      good.Checkcode,
		BuyerIDCard:    good.BuyerIDCard,
	}
}
