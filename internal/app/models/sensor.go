package models

import "gorm.io/gorm"

type Sensor struct {
	gorm.Model
	DeviceCode string `gorm:"not null; unique; index" json:"device_code"` // 设备识别码
	Location   string `json:"location"`
	State      string `json:"state"`
	Type       string `json:"type"`
}

type SensorInfo struct {
	DeviceCode string `json:"device_code"` // 设备识别码
	Location   string `json:"location"`
	State      string `json:"state"`
	Type       string `json:"type"`
}

func ToSensorInfo(s *Sensor) SensorInfo {
	return SensorInfo{
		DeviceCode: s.DeviceCode,
		Location:   s.Location,
		State:      s.State,
		Type:       s.Type,
	}
}
