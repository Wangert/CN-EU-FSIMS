package product

import (
	"time"

	"gorm.io/gorm"
)

type Cow struct {
	gorm.Model
	Number               string     `gorm:"not null; unique; type:varchar(256)" json:"number"`
	Age                  int        `gorm:"not null" json:"age"`
	Weight               float64    `gorm:"not null" json:"weight"`
	QuarantineCertNumber *string    `json:"quarantine_cert_number"`
	OwnerIDCard          string     `gorm:"not null" json:"owner_id_card"`
	OwnerAddress         string     `json:"owner_address"`
	EntryTime            *time.Time `json:"entry_time"`
	State                int        `gorm:"not null" json:"state"`
	HouseNumber          string     `gorm:"not null; type:varchar(256)" json:"house_number"`
	BatchNumber          *string    `gorm:"type:varchar(256)" json:"batch_number"`
}

type CowInfo struct {
	Number               string  `json:"number"`
	Age                  int     `json:"age"`
	Weight               float64 `json:"weight"`
	QuarantineCertNumber string  `json:"quarantine_cert_number"`
	OwnerIDCard          string  `json:"owner_id_card"`
	OwnerAddress         string  `json:"owner_address"`
	EntryTime            string  `json:"entry_time"`
	State                int     `json:"state"`
	HouseNumber          string  `json:"house_number"`
	BatchNumber          *string `json:"batch_number"`
}

func ToCowInfo(cow *Cow) CowInfo {
	qcn := ""
	if cow.QuarantineCertNumber != nil {
		qcn = *cow.QuarantineCertNumber
	}

	et := ""
	loc, _ := time.LoadLocation("Asia/Shanghai")

	if cow.EntryTime != nil {
		time := cow.EntryTime.In(loc)
		et = time.Format("2006-01-02 15:04:05")
	}

	return CowInfo{
		Number:               cow.Number,
		Age:                  cow.Age,
		Weight:               cow.Weight,
		QuarantineCertNumber: qcn,
		OwnerIDCard:          cow.OwnerIDCard,
		OwnerAddress:         cow.OwnerAddress,
		EntryTime:            et,
		State:                cow.State,
		HouseNumber:          cow.HouseNumber,
		BatchNumber:          cow.BatchNumber,
	}
}
