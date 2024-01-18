package slaughter

import (
	"time"

	"gorm.io/gorm"
)

type StaUni struct {
	//员工工作服
	gorm.Model
	HouseNumber  string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	TimeRecordAt time.Time `json:"time_record_at"` //记录时间
	StaUni1      float32   `json:"sta_uni_1"`      //紫外灭菌
	StaUni2      float32   `json:"sta_uni_2"`      //臭氧
	StaUni3      float32   `json:"sta_uni_3"`      //臭氧残留
	StaUni4      float32   `json:"sta_uni_4"`      //湿度
	StaUni5      float32   `json:"sta_uni_5"`      //温度
	StaUni6      float32   `json:"sta_uni_6"`      //预冷间消毒记录
	StaUni7      float32   `json:"sta_uni_7"`      //氯含量
	StaUni8      float32   `json:"sta_uni_8"`      //工作服 功率
	StaUni9      float32   `json:"sta_uni_9"`      //工作服 时间
	StaUni10     float32   `json:"sta_uni_10"`     //消毒记录 方式
	StaUni11     float32   `json:"sta_uni_11"`     //消毒记录 浓度
	StaUni12     float32   `json:"sta_uni_12"`     //消毒记录 班次
	StaUni13     float32   `json:"sta_uni_13"`     //消毒记录 器具
	StaUni14     float32   `json:"sta_uni_15"`     //消毒记录 环境
}

type SlaughterLightRecord struct {
	//屠宰环境光照记录
	gorm.Model
	HouseNumber   string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	TimeRecordAt  time.Time `json:"time_record_at"`    //记录时间
	SlaEnvLigRec1 float64   `json:"sla_env_lig_rec_1"` //屠宰间环境
	SlaEnvLigRec2 float64   `json:"sla_env_lig_rec_2"` //车间
	SlaEnvLigRec3 float64   `json:"sla_env_lig_rec_3"` //检疫
	SlaEnvLigRec4 float64   `json:"sla_env_lig_rec_4"` //预冷通道
}

type StaUniInfo struct {
	//员工工作服
	HouseNumber  string  `json:"house_number"`
	TimeRecordAt string  `json:"time_record_at"` //记录时间
	StaUni1      float32 `json:"sta_uni_1"`      //紫外灭菌
	StaUni2      float32 `json:"sta_uni_2"`      //臭氧
	StaUni3      float32 `json:"sta_uni_3"`      //臭氧残留
	StaUni4      float32 `json:"sta_uni_4"`      //湿度
	StaUni5      float32 `json:"sta_uni_5"`      //温度
	StaUni6      float32 `json:"sta_uni_6"`      //预冷间消毒记录
	StaUni7      float32 `json:"sta_uni_7"`      //氯含量
	StaUni8      float32 `json:"sta_uni_8"`      //工作服 功率
	StaUni9      float32 `json:"sta_uni_9"`      //工作服 时间
	StaUni10     float32 `json:"sta_uni_10"`     //消毒记录 方式
	StaUni11     float32 `json:"sta_uni_11"`     //消毒记录 浓度
	StaUni12     float32 `json:"sta_uni_12"`     //消毒记录 班次
	StaUni13     float32 `json:"sta_uni_13"`     //消毒记录 器具
	StaUni14     float32 `json:"sta_uni_15"`     //消毒记录 环境
}

type SlaughterLightRecordInfo struct {
	//屠宰环境光照记录
	HouseNumber   string  `json:"house_number"`
	TimeRecordAt  string  `json:"time_record_at"`    //记录时间
	SlaEnvLigRec1 float64 `json:"sla_env_lig_rec_1"` //屠宰间环境
	SlaEnvLigRec2 float64 `json:"sla_env_lig_rec_2"` //车间
	SlaEnvLigRec3 float64 `json:"sla_env_lig_rec_3"` //检疫
	SlaEnvLigRec4 float64 `json:"sla_env_lig_rec_4"` //预冷通道
}

func ToStaUniInfo(uni *StaUni) StaUniInfo {
	return StaUniInfo{
		HouseNumber:  uni.HouseNumber,
		TimeRecordAt: uni.TimeRecordAt.Format("2006-01-02 15:04:05"),
		StaUni1:      uni.StaUni1,
		StaUni2:      uni.StaUni2,
		StaUni3:      uni.StaUni3,
		StaUni4:      uni.StaUni4,
		StaUni5:      uni.StaUni5,
		StaUni6:      uni.StaUni6,
		StaUni7:      uni.StaUni7,
		StaUni8:      uni.StaUni8,
		StaUni9:      uni.StaUni9,
		StaUni10:     uni.StaUni10,
		StaUni11:     uni.StaUni11,
		StaUni12:     uni.StaUni12,
		StaUni13:     uni.StaUni13,
		StaUni14:     uni.StaUni14,
	}
}

func ToSlaEnvLigRecInfo(rec *SlaughterLightRecord) SlaughterLightRecordInfo {
	return SlaughterLightRecordInfo{
		HouseNumber:   rec.HouseNumber,
		TimeRecordAt:  rec.TimeRecordAt.Format("2006-01-02 15:04:05"),
		SlaEnvLigRec1: rec.SlaEnvLigRec1,
		SlaEnvLigRec2: rec.SlaEnvLigRec2,
		SlaEnvLigRec3: rec.SlaEnvLigRec3,
		SlaEnvLigRec4: rec.SlaEnvLigRec4,
	}
}
