package response

import "CN-EU-FSIMS/internal/app/models"

type ResSensors struct {
	Sensors []models.SensorInfo `json:"sensors"`
	Count   int64               `json:"count"`
}

type ResSensorDataArr struct {
	SensorDataArr []models.SensorDataInfo `json:"sensor_data_arr"`
	Count         int64                   `json:"count"`
}

type ResSensorsTemperature struct {
	TimePoints   []string      `json:"time_points"`
	SensorsTemps []SensorTemps `json:"sensor_temps"`
}

type SensorTemps struct {
	DeviceCode string   `json:"device_code"`
	Location   string   `json:"location"`
	Temps      []string `json:"temps"`
}

type ResSensorsCO2 struct {
	TimePoints  []string     `json:"time_points"`
	SensorsCO2s []SensorCO2s `json:"sensors_co2s"`
}

type SensorCO2s struct {
	DeviceCode string   `json:"device_code"`
	Location   string   `json:"location"`
	CO2s       []string `json:"co2s"`
}

type ResSensorsHumidity struct {
	TimePoints        []string           `json:"time_points"`
	SensorsHumidities []SensorHumidities `json:"sensors_humidities"`
}

type SensorHumidities struct {
	DeviceCode string   `json:"device_code"`
	Location   string   `json:"location"`
	Humidities []string `json:"humidities"`
}

type ResSensorsDewpoint struct {
	TimePoints       []string          `json:"time_points"`
	SensorsDewpoints []SensorDewpoints `json:"sensors_dewpoints"`
}

type SensorDewpoints struct {
	DeviceCode string   `json:"device_code"`
	Location   string   `json:"location"`
	Dewpoints  []string `json:"dewpoints"`
}
