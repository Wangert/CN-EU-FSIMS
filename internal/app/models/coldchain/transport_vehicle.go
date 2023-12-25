package coldchain

import (
	"gorm.io/gorm"
)

type TransportVehicle struct {
	gorm.Model
	TVNumber      string `gorm:"not null; unique; type:varchar(256)" json:"tv_number"`
	LicenseNumber string `gorm:"not null; type:varchar(256)" json:"license_number"`
	Driver        string `gorm:"not null; type:varchar(100)" json:"driver"`
	DriverPhone   string `gorm:"not null; type:varchar(100)" json:"driver_phone"`
	State         uint   `gorm:"not null" json:"state"`
}

type TransportVehicleInfo struct {
	TVNumber      string `json:"tv_number"`
	LicenseNumber string `json:"license_number"`
	Driver        string `json:"driver"`
	DriverPhone   string `json:"driver_phone"`
	State         uint   `json:"state"`
}

func ToTransportVehicleInfo(tv *TransportVehicle) TransportVehicleInfo {
	return TransportVehicleInfo{
		TVNumber:      tv.TVNumber,
		LicenseNumber: tv.LicenseNumber,
		Driver:        tv.Driver,
		DriverPhone:   tv.DriverPhone,
		State:         tv.State,
	}
}
