package service

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/query"
	"context"
	"errors"
	"time"
)

type PendingLogs struct {
	TimeStamp time.Time
	UUID      string
	Account   string
	Type      int
	Action    string
}

func AddLog(pl *PendingLogs) error {
	lg := query.Log
	logRecord := models.Log{
		TimeStamp: pl.TimeStamp,
		UUID:      pl.UUID,
		Account:   pl.Account,
		Type:      pl.Type,
		Action:    pl.Action,
	}
	err := lg.WithContext(context.Background()).Create(&logRecord)
	if err != nil {
		return errors.New("Failed to write log!")
	}
	return nil
}

func GetAllLogs() ([]models.LogInfo, int64, error) {
	q := query.Log
	ls, err := q.WithContext(context.Background()).Find()
	if err != nil {
		return []models.LogInfo{}, 0, err
	}
	count := len(ls)
	logs := make([]models.LogInfo, count)
	for i, log := range ls {
		logs[i] = models.ToLogInfo(log)
	}

	return logs, int64(count), nil
}
