package models

import (
	"time"

	"gorm.io/gorm"
)

type SensorData struct {
	gorm.Model
	TimeStamp           time.Time `gorm:"not null"`
	AirTemperature      string    `json:"air_temperature"`      //空气温度
	Lllumination        string    `json:"lllumination"`         // 光照强度
	DeviceCode          string    `json:"device_code"`          // 设备识别码
	CO2                 string    `json:"co2"`                  // CO2
	AirHumidity         string    `json:"air_humidity"`         // 空气湿度
	NH3                 string    `json:"nh3"`                  // NH3
	H2S                 string    `json:"h2s"`                  // H2S
	AtmosphericPressure string    `json:"atmospheric_pressure"` // 大气压强
	WindSpeed           string    `json:"wind_speed"`           // 风速
	CO                  string    `json:"co"`                   // CO
	AirDewpoint         string    `json:"air_dewpoint"`         // 空气露点
}

type SensorDataInfo struct {
	TimeStamp           time.Time `gorm:"not null"`
	AirTemperature      string    `json:"air_temperature"`      //空气温度
	Lllumination        string    `json:"lllumination"`         // 光照强度
	DeviceCode          string    `json:"device_code"`          // 设备识别码
	CO2                 string    `json:"co2"`                  // CO2
	AirHumidity         string    `json:"air_humidity"`         // 空气湿度
	NH3                 string    `json:"nh3"`                  // NH3
	H2S                 string    `json:"h2s"`                  // H2S
	AtmosphericPressure string    `json:"atmospheric_pressure"` // 大气压强
	WindSpeed           string    `json:"wind_speed"`           // 风速
	CO                  string    `json:"co"`                   // CO
	AirDewpoint         string    `json:"air_dewpoint"`         // 空气露点
}

func ToSensorDataInfo(b *SensorData) SensorDataInfo {
	return SensorDataInfo{
		TimeStamp:           b.TimeStamp,
		AirTemperature:      b.AirTemperature,
		Lllumination:        b.Lllumination,
		DeviceCode:          b.DeviceCode,
		CO2:                 b.CO2,
		AirHumidity:         b.AirHumidity,
		NH3:                 b.NH3,
		H2S:                 b.H2S,
		AtmosphericPressure: b.AtmosphericPressure,
		WindSpeed:           b.WindSpeed,
		CO:                  b.CO,
		AirDewpoint:         b.AirDewpoint,
	}
}
