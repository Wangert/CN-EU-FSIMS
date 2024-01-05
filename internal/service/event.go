package service

import (
	"time"
)

const (
	PASTURE_ABNORMAL_FEED_HEAVY_METAL_INDEX_CONTENT = "牧场饲料重金属指数异常"
	PASTURE_ABNORMAL_FEED_FUNGUS_INDEX_CONTENT      = "牧场饲料真菌指数异常"
	PASTURE_ABNORMAL_WATER_INDEX_CONTENT            = "牧场水质指数异常"
	PASTURE_ABNORMAL_BASIC_ENVIRONMENT_INDEX        = "牧场基本环境指数异常"
	PASTURE_ABNORMAL_BUFFER_INDEX                   = "牧场缓冲区环境指数异常"
	PASTURE_ABNORMAL_AREA_INDEX                     = "牧场场区环境指数异常"
	PASTURE_ABNORMAL_COWHOUSE_INDEX                 = "牧场牛舍环境指数异常"
	PASTURE_ABNORMAL_PADDING_INDEX                  = "牧场垫料指数异常"

	SLAUGHTER_ABNORMAL_WATER_INDEX_CONTENT          = "屠宰场水质指数异常"
	SLAUGHTER_ABNORMAL_PRECOOL_SHOP_INDEX_CONTENT   = "屠宰场预冷车间指数异常"
	SLAUGHTER_ABNORMAL_SLAUGHTER_SHOP_INDEX_CONTENT = "屠宰场屠宰车间指数异常"
	SLAUGHTER_ABNORMAL_DIVISION_SHOP_INDEX_CONTENT  = "屠宰场分割车间指数异常"
	SLAUGHTER_ABNORMAL_ACID_SHOP_INDEX_CONTENT      = "屠宰场排酸车间指数异常"
	SLAUGHTER_ABNORMAL_FROZEN_SHOP_INDEX_CONTENT    = "屠宰场冷冻车间指数异常"
)

type Event struct {
	Source                 string    `gorm:"not null" json:"source"`
	Content                string    `gorm:"not null" json:"content"`
	EventTime              time.Time `gorm:"not null" json:"event_time"`
	EventType              int       `gorm:"not null" json:"event_type"`
	Proposal               string    `gorm:"not null" json:"proposal"`
	AffectedProductsNumber string    `json:"affected_products_number"`
	RiskLevel              int       `gorm:"not null" json:"risk_level"`
}
