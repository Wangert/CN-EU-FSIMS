package service

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/sensorquery"
	"context"
)

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
