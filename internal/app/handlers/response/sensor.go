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
