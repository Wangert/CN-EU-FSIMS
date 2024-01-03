package pasture

import (
	"gorm.io/gorm"
	"time"
)

type PastureWaterRecord struct {
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
	gorm.Model
	TimeRecordAt time.Time         `json:"time_record_at"`
	HouseNumber  string            `json:"house_number"` //时间记录
	OapGci       PastureOapGci     `json:"oap_gci"`      //感官性状和一般化学指标
	ToxIndex     PastureToxIndex   `json:"tox_index"`    //毒理性指标
	MicroIndex   PastureMicroIndex `json:"micro_index"`  //微生物指标
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

type PastureMicroIndexInfo struct {
	//微生物指标
	MicroIndex1 float64 `json:"micro_index_1"` //总大肠菌群
	MicroIndex2 float64 `json:"micro_index_2"` //粪大肠菌群
	MicroIndex3 float64 `json:"micro_index_3"` //菌落总数
}
