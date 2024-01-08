package service

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/query"
	"CN-EU-FSIMS/internal/service/analysis"
	"CN-EU-FSIMS/utils"
	"context"
	"github.com/robfig/cron/v3"
	"time"
)

const (
	PASTURE_FEED_HEAVY_METAL  = "牧场饲料重金属"
	PASTURE_FEED_MYCOTOXINS   = "牧场饲料真菌"
	PASTURE_WATER             = "牧场水质"
	PASTURE_BASIC_ENVIRONMENT = "牧场基本环境"
	PASTURE_BUFFER            = "牧场缓冲区"
	PASTURE_AREA              = "牧场场区"
	PASTURE_COWHOUSE          = "牧场牛舍"
	PASTURE_PADDING           = "牧场垫料"

	SLAUGHTER_WATER          = "屠宰场水质"
	SLAUGHTER_PRECOOL_SHOP   = "屠宰场预冷车间"
	SLAUGHTER_SLAUGHTER_SHOP = "屠宰场屠宰车间"
	SLAUGHTER_DIVISION_SHOP  = "屠宰场分割车间"
	SLAUGHTER_ACID_SHOP      = "屠宰场排酸车间"
	SLAUGHTER_FROZEN_SHOP    = "屠宰场冷冻车间"
)

var monitoringNames = [...]string{PASTURE_FEED_HEAVY_METAL, PASTURE_FEED_MYCOTOXINS, PASTURE_WATER, PASTURE_BASIC_ENVIRONMENT,
	PASTURE_BUFFER, PASTURE_AREA, PASTURE_COWHOUSE, PASTURE_PADDING, SLAUGHTER_WATER, SLAUGHTER_PRECOOL_SHOP,
	SLAUGHTER_SLAUGHTER_SHOP, SLAUGHTER_DIVISION_SHOP, SLAUGHTER_ACID_SHOP, SLAUGHTER_FROZEN_SHOP}

type TimedTask struct {
	Spec string
	Fc   func()
}

func TimedTasksStart(c *cron.Cron, tts []TimedTask) {

	//tt := TimedTask{
	//	Spec: "",
	//	Fc: func() {
	//		currentTime := time.Now()
	//		err := PastureFeedHeavyMetalMonitoring(currentTime)
	//		if err != nil {
	//			glog.Errorln(err)
	//		}
	//	},
	//}

	for _, tt := range tts {
		_, _ = c.AddFunc(tt.Spec, tt.Fc)
	}

	c.Start()
}

// 牧场饲料真菌毒素监测
func PastureFeedMycotoxinsMonitoring(currentTime time.Time) error {
	m := query.Q.MonitoringTimeRecord
	mtr, err := m.WithContext(context.Background()).Where(m.IndexName.Eq(PASTURE_FEED_MYCOTOXINS)).First()
	if err != nil {
		return err
	}

	preTime := *mtr.LastTime

	f := query.Q.PastureFeedMycotoxins
	pasFeedMycotoxinsRecords, err := f.WithContext(context.Background()).Where(f.TimeRecordAt.Between(preTime, currentTime)).Find()
	if err != nil {
		return err
	}

	for _, record := range pasFeedMycotoxinsRecords {
		// 识别危害指标
		abnormalList, abnormalCount, err := analysis.JudgeHarmForPastureFeedFungus(record)
		if err != nil {
			return err
		}

		riskLevel := analysis.RiskLevel(abnormalCount)

		// 获取影响的FeedBatch
		b := query.Q.FeedingBatch
		p := query.Q.Procedure
		bs, err := b.WithContext(context.Background()).Where(b.HouseNumber.Eq(record.HouseNumber)).
			Preload(b.Procedure.Where(p.StartTimestamp.Lte(record.TimeRecordAt)).
				Where(p.CompletedTimestamp.Gte(record.TimeRecordAt))).Find()
		if err != nil {
			return err
		}

		// 受影响的cow
		cowsNumber := []string{}
		for _, batch := range bs {
			for _, cow := range batch.Cows {
				cowsNumber = append(cowsNumber, cow.Number)
			}
		}

		// 获取接收人
		u := query.Q.FSIMSUser
		users, err := u.WithContext(context.Background()).Where(u.Type.Neq(CUSTOMER_USER_TYPE)).Find()
		if err != nil {
			return err
		}

		// 创建事件
		event := Event{
			Source:                 record.HouseNumber,
			Content:                PASTURE_ABNORMAL_FEED_FUNGUS_INDEX_CONTENT,
			EventTime:              record.TimeRecordAt,
			EventType:              1,
			AffectedProductsNumber: utils.StrArrToStr(cowsNumber),
			Proposal:               utils.StrArrToStr(abnormalList),
			RiskLevel:              riskLevel,
		}

		// 发送通知
		for _, user := range users {
			err = PushNotification(user.UUID, &event)
			if err != nil {
				return err
			}
		}
	}

	_, err = m.WithContext(context.Background()).Where(m.IndexName.Eq(PASTURE_FEED_HEAVY_METAL)).
		Updates(map[string]interface{}{"last_time": currentTime})
	if err != nil {
		return err
	}

	return nil
}

// 牧场饲料重金属监测
func PastureFeedHeavyMetalMonitoring(currentTime time.Time) error {
	m := query.Q.MonitoringTimeRecord
	mtr, err := m.WithContext(context.Background()).Where(m.IndexName.Eq(PASTURE_FEED_HEAVY_METAL)).First()
	if err != nil {
		return err
	}

	preTime := *mtr.LastTime

	f := query.Q.PastureFeedHeavyMetal
	pasFeedHeavyMetalRecords, err := f.WithContext(context.Background()).Where(f.TimeRecordAt.Between(preTime, currentTime)).Find()
	if err != nil {
		return err
	}

	for _, record := range pasFeedHeavyMetalRecords {
		// 识别危害指标
		abnormalList, abnormalCount, err := analysis.JudgeHarmForPastureFeedHeavyMetal(record)
		if err != nil {
			return err
		}

		riskLevel := analysis.RiskLevel(abnormalCount)

		// 获取影响的FeedBatch
		b := query.Q.FeedingBatch
		p := query.Q.Procedure
		bs, err := b.WithContext(context.Background()).Where(b.HouseNumber.Eq(record.HouseNumber)).
			Preload(b.Procedure.Where(p.StartTimestamp.Lte(record.TimeRecordAt)).
				Where(p.CompletedTimestamp.Gte(record.TimeRecordAt))).Find()
		if err != nil {
			return err
		}

		// 受影响的cow
		cowsNumber := []string{}
		for _, batch := range bs {
			for _, cow := range batch.Cows {
				cowsNumber = append(cowsNumber, cow.Number)
			}
		}

		// 获取接收人
		u := query.Q.FSIMSUser
		users, err := u.WithContext(context.Background()).Where(u.Type.Neq(CUSTOMER_USER_TYPE)).Find()
		if err != nil {
			return err
		}

		// 创建事件
		event := Event{
			Source:                 record.HouseNumber,
			Content:                PASTURE_ABNORMAL_FEED_HEAVY_METAL_INDEX_CONTENT,
			EventTime:              record.TimeRecordAt,
			EventType:              1,
			AffectedProductsNumber: utils.StrArrToStr(cowsNumber),
			Proposal:               utils.StrArrToStr(abnormalList),
			RiskLevel:              riskLevel,
		}

		// 发送通知
		for _, user := range users {
			err = PushNotification(user.UUID, &event)
			if err != nil {
				return err
			}
		}
	}

	_, err = m.WithContext(context.Background()).Where(m.IndexName.Eq(PASTURE_FEED_HEAVY_METAL)).
		Updates(map[string]interface{}{"last_time": currentTime})
	if err != nil {
		return err
	}

	return nil
}

// 初始化各监控记录的监测历史表
func InitMonitoringTimeRecords() error {
	var err error
	tx := query.Q.Begin()
	defer func() {
		if recover() != nil {
			_ = tx.Rollback()
		}
	}()

	for _, name := range monitoringNames {
		record := models.MonitoringTimeRecord{
			IndexName: name,
			LastTime:  nil,
		}

		err = query.Q.MonitoringTimeRecord.WithContext(context.Background()).Create(&record)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return nil
}