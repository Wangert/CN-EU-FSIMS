package slaughter

import "gorm.io/gorm"
import "CN-EU-FSIMS/internal/app/models/premortem"

type ChiledFreshIndex struct {
	//冷鲜肉相关指标
	gorm.Model
	TimeRecordAt string        `json:"time_record_at"`                                                  //记录时间
	PaGerm       PaGerm        `gorm:"foreignKey:ChiledFreshIndexID; references:ID" json:"pa_germ"`     ////致病菌监测
	OtherIndex   OtherIndex    `gorm:"foreignKey:ChiledFreshIndexID; references:ID" json:"other_index"` //其他指标
	DrugsResi    DrugsResi     `gorm:"foreignKey:ChiledFreshIndexID; references:ID" json:"drugs_resi"`  //农兽药残留
	Gps          premortem.Gps `gorm:"foreignKey:ChiledFreshIndexID; references:ID" json:"gps"`         //gps信息
	TranTemp     float32       `json:"tran_temp"`                                                       //运输温度
}
type PaGerm struct {
	//致病菌监测
	gorm.Model
	ChiledFreshIndexID uint    `json:"chiled_fresh_index_id"`
	PaGerm1            float32 `json:"pa_germ_1"` //光谱监测沙门
	PaGerm2            float32 `json:"pa_germ_2"` //STEC
	PaGerm3            float32 `json:"pa_germ_3"` //李斯特
	PaGerm4            float32 `json:"pa_germ_4"` //阪崎杆菌
	PaGerm5            float32 `json:"pa_germ_5"` //金葡
}

type OtherIndex struct {
	//其他指标
	gorm.Model
	ChiledFreshIndexID uint    `json:"chiled_fresh_index_id"`
	OtherIndex1        float32 `json:"other_index_1"` //菌落总数（光谱）
	OtherIndex2        float32 `json:"other_index_2"` //挥发性盐基氮（光谱）
	OtherIndex3        float32 `json:"other_index_3"` //大肠菌群（光谱）
	OtherIndex4        float32 `json:"other_index_4"` //假单胞
	OtherIndex5        float32 `json:"other_index_5"` //乳酸菌（光谱）
	OtherIndex6        float32 `json:"other_index_6"` //铅
	OtherIndex7        float32 `json:"other_index_7"` //无机砷
	OtherIndex8        float32 `json:"other_index_8"` //牛肉水分
	OtherIndex9        float32 `json:"other_index_9"` //牛肉蛋白质
}

type DrugsResi struct {
	//农兽药残留
	gorm.Model
	ChiledFreshIndexID uint    `json:"chiled_fresh_index_id"`
	DrugsResi1         float32 `json:"drugs_resi_1"` //恩诺沙星
	DrugsResi2         float32 `json:"drugs_resi_2"` //庆大霉素
	DrugsResi3         float32 `json:"drugs_resi_3"` //青霉素
	DrugsResi4         float32 `json:"drugs_resi_4"` //氯氢吡啶
}
