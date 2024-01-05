package pasture

import (
	"gorm.io/gorm"
	"time"
)

type PastureWaterRecord struct {
	//饮用水指标
	gorm.Model
	TimeRecordAt time.Time         `json:"time_record_at"`
	HouseNumber  string            `gorm:"not null; type:varchar(256)" json:"house_number"`                   //时间记录
	OapGci       PastureOapGci     `gorm:"foreignKey:PastureWaterRecordID; references:ID" json:"oap_gci"`     //感官性状和一般化学指标
	ToxIndex     PastureToxIndex   `gorm:"foreignKey:PastureWaterRecordID; references:ID" json:"tox_index"`   //毒理性指标
	MicroIndex   PastureMicroIndex `gorm:"foreignKey:PastureWaterRecordID; references:ID" json:"micro_index"` //微生物指标
}

type PastureOapGci struct {
	//感官性状和一般化学指标
	gorm.Model
	PastureWaterRecordID *uint   `json:"pasture_water_record_id"`
	OapGci1              float64 `json:"oap_gci_1"`  //浑浊度
	OapGci2              float64 `json:"oap_gci_2"`  //总硬度（以CaCO3)计
	OapGci3              float64 `json:"oap_gci_3"`  //PH值
	OapGci4              float64 `json:"oap_gci_4"`  //硫酸
	OapGci5              float64 `json:"oap_gci_5"`  //氯化物
	OapGci6              float64 `json:"oap_gci_6"`  //总溶解性固体
	OapGci7              float64 `json:"oap_gci_7"`  //氟化物
	OapGci8              float64 `json:"oap_gci_8"`  //六六六
	OapGci9              float64 `json:"oap_gci_9"`  //滴滴涕
	OapGci10             float64 `json:"oap_gci_10"` //氨氮
}

type PastureToxIndex struct {
	//毒理性指标
	gorm.Model
	PastureWaterRecordID *uint   `json:"pasture_water_record_id"`
	ToxIndex1            float64 `json:"tox_index_1"`  //氰化物
	ToxIndex2            float64 `json:"tox_index_2"`  //总砷
	ToxIndex3            float64 `json:"tox_index_3"`  //总汞
	ToxIndex4            float64 `json:"tox_index_4"`  //总铅
	ToxIndex5            float64 `json:"tox_index_5"`  //铬（六价）
	ToxIndex6            float64 `json:"tox_index_6"`  //总镉
	ToxIndex7            float64 `json:"tox_index_7"`  //硝酸盐
	ToxIndex8            float64 `json:"tox_index_8"`  //砷
	ToxIndex9            float64 `json:"tox_index_9"`  //镉
	ToxIndex10           float64 `json:"tox_index_10"` //氯化物
	ToxIndex11           float64 `json:"tox_index_11"` //硝酸钾
	ToxIndex12           float64 `json:"tox_index_12"` //三氯甲烷
	ToxIndex13           float64 `json:"tox_index_13"` //一氯二溴甲烷
	ToxIndex14           float64 `json:"tox_index_14"` //二氯一溴甲烷
	ToxIndex15           float64 `json:"tox_index_15"` //三溴甲烷
	ToxIndex16           float64 `json:"tox_index_16"` //三卤甲烷（三氯甲烷、一氯二溴甲烷、二氯一溴甲烷、三溴甲烷的总和）
	ToxIndex17           float64 `json:"tox_index_17"` //二氯乙酸
	ToxIndex18           float64 `json:"tox_index_18"` //三氯乙酸
	ToxIndex19           float64 `json:"tox_index_19"` //溴酸盐
	ToxIndex20           float64 `json:"tox_index_20"` //亚氯酸盐
	ToxIndex21           float64 `json:"tox_index_21"` //氯酸盐
	ToxIndex22           float64 `json:"tox_index_22"` //化学耗氧量
}

type PastureMicroIndex struct {
	//微生物指标
	gorm.Model
	PastureWaterRecordID *uint   `json:"pasture_water_record_id"`
	MicroIndex1          float64 `json:"micro_index_1"` //总大肠菌群
	MicroIndex2          float64 `json:"micro_index_2"` //粪大肠菌群
	MicroIndex3          float64 `json:"micro_index_3"` //菌落总数
}

type PastureWaterRecordInfo struct {
	TimeRecordAt string                `json:"time_record_at"`
	HouseNumber  string                `json:"house_number"` //时间记录
	OapGci       PastureOapGciInfo     `json:"oap_gci"`      //感官性状和一般化学指标
	ToxIndex     PastureToxIndexInfo   `json:"tox_index"`    //毒理性指标
	MicroIndex   PastureMicroIndexInfo `json:"micro_index"`  //微生物指标
}

type PastureOapGciInfo struct {
	//感官性状和一般化学指标
	OapGci1  float64 `json:"oap_gci_1"`  //浑浊度
	OapGci2  float64 `json:"oap_gci_2"`  //总硬度（以CaCO3)计
	OapGci3  float64 `json:"oap_gci_3"`  //PH值
	OapGci4  float64 `json:"oap_gci_4"`  //硫酸
	OapGci5  float64 `json:"oap_gci_5"`  //氯化物
	OapGci6  float64 `json:"oap_gci_6"`  //总溶解性固体
	OapGci7  float64 `json:"oap_gci_7"`  //氟化物
	OapGci8  float64 `json:"oap_gci_8"`  //六六六
	OapGci9  float64 `json:"oap_gci_9"`  //滴滴涕
	OapGci10 float64 `json:"oap_gci_10"` //氨氮
}

func ToPastureOapGciInfo(ogi *PastureOapGci) PastureOapGciInfo {
	return PastureOapGciInfo{
		OapGci1:  ogi.OapGci1,
		OapGci2:  ogi.OapGci2,
		OapGci3:  ogi.OapGci3,
		OapGci4:  ogi.OapGci4,
		OapGci5:  ogi.OapGci5,
		OapGci6:  ogi.OapGci6,
		OapGci7:  ogi.OapGci7,
		OapGci8:  ogi.OapGci8,
		OapGci9:  ogi.OapGci9,
		OapGci10: ogi.OapGci10,
	}
}

type PastureToxIndexInfo struct {
	//毒理性指标
	ToxIndex1  float64 `json:"tox_index_1"`  //氰化物
	ToxIndex2  float64 `json:"tox_index_2"`  //总砷
	ToxIndex3  float64 `json:"tox_index_3"`  //总汞
	ToxIndex4  float64 `json:"tox_index_4"`  //总铅
	ToxIndex5  float64 `json:"tox_index_5"`  //铬（六价）
	ToxIndex6  float64 `json:"tox_index_6"`  //总镉
	ToxIndex7  float64 `json:"tox_index_7"`  //硝酸盐
	ToxIndex8  float64 `json:"tox_index_8"`  //砷
	ToxIndex9  float64 `json:"tox_index_9"`  //镉
	ToxIndex10 float64 `json:"tox_index_10"` //氯化物
	ToxIndex11 float64 `json:"tox_index_11"` //硝酸钾
	ToxIndex12 float64 `json:"tox_index_12"` //三氯甲烷
	ToxIndex13 float64 `json:"tox_index_13"` //一氯二溴甲烷
	ToxIndex14 float64 `json:"tox_index_14"` //二氯一溴甲烷
	ToxIndex15 float64 `json:"tox_index_15"` //三溴甲烷
	ToxIndex16 float64 `json:"tox_index_16"` //三卤甲烷（三氯甲烷、一氯二溴甲烷、二氯一溴甲烷、三溴甲烷的总和）
	ToxIndex17 float64 `json:"tox_index_17"` //二氯乙酸
	ToxIndex18 float64 `json:"tox_index_18"` //三氯乙酸
	ToxIndex19 float64 `json:"tox_index_19"` //溴酸盐
	ToxIndex20 float64 `json:"tox_index_20"` //亚氯酸盐
	ToxIndex21 float64 `json:"tox_index_21"` //氯酸盐
	ToxIndex22 float64 `json:"tox_index_22"` //化学耗氧量
}

func ToPastureToxIndexInfo(tii *PastureToxIndex) PastureToxIndexInfo {
	return PastureToxIndexInfo{
		ToxIndex1:  tii.ToxIndex1,
		ToxIndex2:  tii.ToxIndex2,
		ToxIndex3:  tii.ToxIndex3,
		ToxIndex4:  tii.ToxIndex4,
		ToxIndex5:  tii.ToxIndex5,
		ToxIndex6:  tii.ToxIndex6,
		ToxIndex7:  tii.ToxIndex7,
		ToxIndex8:  tii.ToxIndex8,
		ToxIndex9:  tii.ToxIndex9,
		ToxIndex10: tii.ToxIndex10,
		ToxIndex11: tii.ToxIndex11,
		ToxIndex12: tii.ToxIndex12,
		ToxIndex13: tii.ToxIndex13,
		ToxIndex14: tii.ToxIndex14,
		ToxIndex15: tii.ToxIndex15,
		ToxIndex16: tii.ToxIndex16,
		ToxIndex17: tii.ToxIndex17,
		ToxIndex18: tii.ToxIndex18,
		ToxIndex19: tii.ToxIndex19,
		ToxIndex20: tii.ToxIndex20,
		ToxIndex21: tii.ToxIndex21,
		ToxIndex22: tii.ToxIndex22,
	}
}

type PastureMicroIndexInfo struct {
	//微生物指标
	MicroIndex1 float64 `json:"micro_index_1"` //总大肠菌群
	MicroIndex2 float64 `json:"micro_index_2"` //粪大肠菌群
	MicroIndex3 float64 `json:"micro_index_3"` //菌落总数
}

func ToPastureMicroIndexInfo(mii *PastureMicroIndex) PastureMicroIndexInfo {
	return PastureMicroIndexInfo{
		MicroIndex1: mii.MicroIndex1,
		MicroIndex2: mii.MicroIndex2,
		MicroIndex3: mii.MicroIndex3,
	}
}

func ToPastureWaterRecordInfo(wr *PastureWaterRecord) PastureWaterRecordInfo {
	return PastureWaterRecordInfo{
		TimeRecordAt: wr.TimeRecordAt.Format("2006-01-02 15:04:05"),
		HouseNumber:  wr.HouseNumber,
		OapGci:       ToPastureOapGciInfo(&wr.OapGci),
		ToxIndex:     ToPastureToxIndexInfo(&wr.ToxIndex),
		MicroIndex:   ToPastureMicroIndexInfo(&wr.MicroIndex),
	}
}
