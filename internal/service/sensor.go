package service

import (
	"CN-EU-FSIMS/internal/app/handlers/response"
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/sensorquery"
	"context"
	"time"
)

const (
	IPOINTLIMIT = 100
	BARNLIMIT   = 40
)

func GetDewpointsForAllBarnPoints() (response.ResSensorsDewpoint, error) {
	sensors, err := GetAllBarnPoints()
	if err != nil {
		return response.ResSensorsDewpoint{}, err
	}

	var timePoints []string

	sensorsDewpoints := make([]response.SensorDewpoints, len(sensors))
	res := response.ResSensorsDewpoint{
		TimePoints:       timePoints,
		SensorsDewpoints: sensorsDewpoints,
	}

	q := sensorquery.SensorData
	for i, sensor := range sensors {
		ssd, err := q.WithContext(context.Background()).Where(q.DeviceCode.Eq(sensor.DeviceCode)).Order(q.TimeStamp.Desc()).Limit(BARNLIMIT).Find()
		if err != nil {
			return response.ResSensorsDewpoint{}, err
		}

		ds := queryDewpoints(ssd)

		sensorDewpoints := response.SensorDewpoints{
			DeviceCode: sensor.DeviceCode,
			Location:   sensor.Location,
			Dewpoints:  ds,
		}

		if i == 0 {
			res.TimePoints = queryTimePoints(ssd)
		}

		res.SensorsDewpoints[i] = sensorDewpoints
	}

	return res, nil
}

func GetHumiditiesForAllBarnPoints() (response.ResSensorsHumidity, error) {
	sensors, err := GetAllBarnPoints()
	if err != nil {
		return response.ResSensorsHumidity{}, err
	}

	var timePoints []string

	sensorsHumidities := make([]response.SensorHumidities, len(sensors))
	res := response.ResSensorsHumidity{
		TimePoints:        timePoints,
		SensorsHumidities: sensorsHumidities,
	}

	q := sensorquery.SensorData
	for i, sensor := range sensors {
		ssd, err := q.WithContext(context.Background()).Where(q.DeviceCode.Eq(sensor.DeviceCode)).Order(q.TimeStamp.Desc()).Limit(BARNLIMIT).Find()
		if err != nil {
			return response.ResSensorsHumidity{}, err
		}

		hs := queryHumidities(ssd)

		sensorHumidities := response.SensorHumidities{
			DeviceCode: sensor.DeviceCode,
			Location:   sensor.Location,
			Humidities: hs,
		}

		if i == 0 {
			res.TimePoints = queryTimePoints(ssd)
		}

		res.SensorsHumidities[i] = sensorHumidities
	}

	return res, nil
}

func GetCO2sForAllBarnPoints() (response.ResSensorsCO2, error) {
	sensors, err := GetAllBarnPoints()
	if err != nil {
		return response.ResSensorsCO2{}, err
	}

	var timePoints []string

	sensorsCO2s := make([]response.SensorCO2s, len(sensors))
	res := response.ResSensorsCO2{
		TimePoints:  timePoints,
		SensorsCO2s: sensorsCO2s,
	}

	q := sensorquery.SensorData
	for i, sensor := range sensors {
		ssd, err := q.WithContext(context.Background()).Where(q.DeviceCode.Eq(sensor.DeviceCode)).Order(q.TimeStamp.Desc()).Limit(BARNLIMIT).Find()
		if err != nil {
			return response.ResSensorsCO2{}, err
		}

		temps := queryCO2s(ssd)

		sensorCO2s := response.SensorCO2s{
			DeviceCode: sensor.DeviceCode,
			Location:   sensor.Location,
			CO2s:       temps,
		}

		if i == 0 {
			res.TimePoints = queryTimePoints(ssd)
		}

		res.SensorsCO2s[i] = sensorCO2s
	}

	return res, nil
}

func GetTemperaturesForAllBarnPoints() (response.ResSensorsTemperature, error) {
	sensors, err := GetAllBarnPoints()
	if err != nil {
		return response.ResSensorsTemperature{}, err
	}

	var timePoints []string

	sensorsTemps := make([]response.SensorTemps, len(sensors))
	res := response.ResSensorsTemperature{
		TimePoints:   timePoints,
		SensorsTemps: sensorsTemps,
	}

	q := sensorquery.SensorData
	for i, sensor := range sensors {
		ssd, err := q.WithContext(context.Background()).Where(q.DeviceCode.Eq(sensor.DeviceCode)).Order(q.TimeStamp.Desc()).Limit(BARNLIMIT).Find()
		if err != nil {
			return response.ResSensorsTemperature{}, err
		}

		temps := queryTemps(ssd)

		sensorTemps := response.SensorTemps{
			DeviceCode: sensor.DeviceCode,
			Location:   sensor.Location,
			Temps:      temps,
		}

		if i == 0 {
			res.TimePoints = queryTimePoints(ssd)
		}

		res.SensorsTemps[i] = sensorTemps
	}

	return res, nil
}

func GetTemperaturesForAllIndependentPoints() (response.ResSensorsTemperature, error) {
	sensors, err := GetAllIndependentPoints()
	if err != nil {
		return response.ResSensorsTemperature{}, err
	}

	var timePoints []string

	sensorsTemps := make([]response.SensorTemps, len(sensors))
	res := response.ResSensorsTemperature{
		TimePoints:   timePoints,
		SensorsTemps: sensorsTemps,
	}

	q := sensorquery.SensorData
	for i, sensor := range sensors {
		ssd, err := q.WithContext(context.Background()).Where(q.DeviceCode.Eq(sensor.DeviceCode)).Order(q.TimeStamp.Desc()).Limit(IPOINTLIMIT).Find()
		if err != nil {
			return response.ResSensorsTemperature{}, err
		}

		temps := queryTemps(ssd)

		sensorTemps := response.SensorTemps{
			DeviceCode: sensor.DeviceCode,
			Location:   sensor.Location,
			Temps:      temps,
		}

		if i == 0 {
			res.TimePoints = queryTimePoints(ssd)
		}

		res.SensorsTemps[i] = sensorTemps
	}

	return res, nil
}

func GetAllSensors() ([]models.SensorInfo, int64, error) {
	q := sensorquery.Sensor
	ss, err := q.WithContext(context.Background()).Find()
	if err != nil {
		return []models.SensorInfo{}, 0, err
	}

	count, err := q.WithContext(context.Background()).Count()

	sensors := make([]models.SensorInfo, count)

	for i, s := range ss {
		sensors[i] = models.ToSensorInfo(s)
	}

	return sensors, count, nil
}

func GetAllIndependentPoints() ([]models.SensorInfo, error) {
	q := sensorquery.Sensor
	ss, err := q.WithContext(context.Background()).Where(q.Type.Eq("独立采集点")).Find()
	if err != nil {
		return []models.SensorInfo{}, err
	}

	sensors := make([]models.SensorInfo, len(ss))
	for i, s := range ss {
		sensors[i] = models.ToSensorInfo(s)
	}

	return sensors, nil
}

func GetAllBarnPoints() ([]models.SensorInfo, error) {
	q := sensorquery.Sensor
	ss, err := q.WithContext(context.Background()).Where(q.Type.Eq("畜棚监测站")).Find()
	if err != nil {
		return []models.SensorInfo{}, err
	}

	sensors := make([]models.SensorInfo, len(ss))
	for i, s := range ss {
		sensors[i] = models.ToSensorInfo(s)
	}

	return sensors, nil
}

func GetSensorDataByDeviceCode(dc string) ([]models.SensorDataInfo, int64, error) {
	q := sensorquery.SensorData
	ssd, err := q.WithContext(context.Background()).Where(q.DeviceCode.Eq(dc)).Find()
	if err != nil {
		return []models.SensorDataInfo{}, 0, err
	}

	count, err := q.WithContext(context.Background()).Where(q.DeviceCode.Eq(dc)).Count()

	sensorDataArr := make([]models.SensorDataInfo, count)
	for i, sd := range ssd {
		sensorDataArr[i] = models.ToSensorDataInfo(sd)
	}

	return sensorDataArr, count, nil
}

func GetLatestSensorDataByDeviceCode(dc string) (models.SensorDataInfo, error) {
	q := sensorquery.SensorData
	sd, err := q.WithContext(context.Background()).Where(q.DeviceCode.Eq(dc)).Order(q.TimeStamp.Desc()).First()
	if err != nil {
		return models.SensorDataInfo{}, err
	}

	return models.ToSensorDataInfo(sd), nil

}

func queryDewpoints(ssd []*models.SensorData) []string {
	count := len(ssd)
	ds := make([]string, count)
	for i, data := range ssd {
		ds[count-i-1] = data.AirDewpoint
	}

	return ds
}

func queryHumidities(ssd []*models.SensorData) []string {
	count := len(ssd)
	hs := make([]string, count)
	for i, data := range ssd {
		hs[count-i-1] = data.AirHumidity
	}

	return hs
}

func queryCO2s(ssd []*models.SensorData) []string {
	count := len(ssd)
	co2s := make([]string, count)
	for i, data := range ssd {
		co2s[count-i-1] = data.CO2
	}

	return co2s
}

func queryTemps(ssd []*models.SensorData) []string {
	count := len(ssd)
	temps := make([]string, count)
	for i, data := range ssd {
		temps[count-i-1] = data.AirTemperature
	}

	return temps
}

func queryTimePoints(ssd []*models.SensorData) []string {
	count := len(ssd)
	timePoints := make([]string, count)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	for i, data := range ssd {
		timestamp := data.TimeStamp.In(loc)
		timePoint := timestamp.Format("15:04")
		timePoints[count-i-1] = timePoint
	}

	return timePoints
}
