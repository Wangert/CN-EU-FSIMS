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
	lg := query.Logs
	logRecord := models.Logs{
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
