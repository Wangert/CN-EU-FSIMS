package request

import (
	"CN-EU-FSIMS/internal/app/models/pasture"
)

type ReqAddCow struct {
	Age         int     `json:"age" form:"age"`
	Weight      float64 `json:"weight" form:"weight"`
	HouseNumber string  `json:"house_number" form:"house_number"`
}

type ReqNewFeedingBatch struct {
	HouseNumber string   `json:"house_number" form:"house_number"`
	Worker      string   `json:"worker" form:"worker"`
	PrePID      string   `json:"pre_pid" form:"pre_pid"`
	CowNumbers  []string `json:"cow_numbers" form:"cow_numbers"`
}

type ReqEndFeeding struct {
	BatchNumber string  `json:"batch_number" form:"batch_number"`
	Worker      string  `json:"worker" form:"worker"`
	HouseNumber string  `json:"house_number" form:"house_number"`
	PM10        float32 `json:"pm_10" form:"pm_10"`   // PM10
	TSP         float32 `json:"tsp" form:"tsp"`       // TSP
	Stench      uint    `json:"stench" form:"stench"` // 恶臭稀释程度
}

type ReqSendToSlaughter struct {
	CowNumber            string `json:"cow_number" form:"cow_number"`
	Operator             string `json:"operator" form:"operator"`
	SlaughterHouseNumber string `json:"slaughter_house_number" form:"slaughter_house_number"`
}

type ReqAddPastureFeedHeavyMetal struct {
	HouseNumber string                    `json:"house_number"`
	RecordAt    int64                     `json:"record_at"`
	As          pasture.PastureFeedAsInfo `json:"as"` //砷元素
	Pb          pasture.PastureFeedPbInfo `json:"pb"` //铅元素
	Cd          pasture.PastureFeedCdInfo `json:"cd"` //镉元素
	Cr          pasture.PastureFeedCrInfo `json:"cr"` //铬元素
}

type ReqAddPastureFeedCass struct {
	////饲料中真菌毒素、农兽药残留
	HouseNumber string               `json:"house_number"`
	RecordAt    int64                `json:"record_at"`
	Afb1        pasture.Afb1Info     `json:"afb_1"`       //黄曲霉毒素B1
	Don         pasture.DonInfo      `json:"don"`         //玉米赤霉烯酮
	T2toxin     pasture.T2toxinInfo  `json:"t2Toxin"`     //脱氧雪腐镰刀菌烯醇（呕吐毒素）
	T2VomZea    pasture.T2VomZeaInfo `json:"t_2_vom_zea"` //T-2毒素 伏马毒素 赭曲霉毒素A
}

type ReqAddPastureWaterRecord struct {
	HouseNumber string                        `json:"house_number"` //时间记录
	RecordAt    int64                         `json:"recordAt"`
	OapGci      pasture.PastureOapGciInfo     `json:"oap_gci"`     //感官性状和一般化学指标
	ToxIndex    pasture.PastureToxIndexInfo   `json:"tox_index"`   //毒理性指标
	MicroIndex  pasture.PastureMicroIndexInfo `json:"micro_index"` //微生物指标
}

type ReqAddPastureBuffer struct {
	HouseNumber string  `json:"houseNumber"`
	TimeStamp   int64   `json:"time_stamp"`
	Buffer1     float32 `json:"buffer_1"`  //NH3氨气
	Buffer2     float32 `json:"buffer_2"`  //H2S硫化氢
	Buffer3     float32 `json:"buffer_3"`  //CO2二氧化碳
	Buffer4     float32 `json:"buffer_4"`  //PM10
	Buffer5     float32 `json:"buffer_5"`  //TSP
	Buffer6     float32 `json:"buffer_6"`  //ODO 恶臭
	Buffer7     float32 `json:"buffer_7"`  //Cd(镉)
	Buffer8     float32 `json:"buffer_8"`  //As（砷）
	Buffer9     float32 `json:"buffer_9"`  //Cu（铜）
	Buffer10    float32 `json:"buffer_10"` //Pb（铅）
	Buffer11    float32 `json:"buffer_11"` //Cr(铬)
	Buffer12    float32 `json:"buffer_12"` //Zn(锌）
}

type ReqAddPastureArea struct {
	TimeStamp    int64   `json:"time_stamp"`
	HouseNumber  string  `gorm:"not null; type:varchar(256)" json:"house_number"`
	CattleFarm1  float32 `json:"cattle_farm_1"`  //NH3氨气
	CattleFarm2  float32 `json:"cattle_farm_2"`  //H2S硫化氢
	CattleFarm3  float32 `json:"cattle_farm_3"`  //CO2二氧化碳
	CattleFarm4  float32 `json:"cattle_farm_4"`  //PM10
	CattleFarm5  float32 `json:"cattle_farm_5"`  //TSP
	CattleFarm6  float32 `json:"cattle_farm_6"`  //ODO 恶臭
	CattleFarm7  float32 `json:"cattle_farm_7"`  //Cd(镉)
	CattleFarm8  float32 `json:"cattle_farm_8"`  //As（砷）
	CattleFarm9  float32 `json:"cattle_farm_9"`  //Cu（铜）
	CattleFarm10 float32 `json:"cattle_farm_10"` //Pb（铅）
	CattleFarm11 float32 `json:"cattle_farm_11"` //Cr(铬)
	CattleFarm12 float32 `json:"cattle_farm_12"` //Zn(锌）
}

type ReqAddCowHouse struct {
	TimeStamp   int64   `json:"time_stamp"`
	HouseNumber string  `gorm:"not null; type:varchar(256)" json:"house_number"`
	CowHouse1   float32 `json:"cow_house_1"`  //NH3氨气
	CowHouse2   float32 `json:"cow_house_2"`  //H2S硫化氢
	CowHouse3   float32 `json:"cow_house_3"`  //CO2二氧化碳
	CowHouse4   float32 `json:"cow_house_4"`  //PM10
	CowHouse5   float32 `json:"cow_house_5"`  //TSP
	CowHouse6   float32 `json:"cow_house_6"`  //ODO 恶臭
	CowHouse7   float32 `json:"cow_house_7"`  //Cd(镉)
	CowHouse8   float32 `json:"cow_house_8"`  //As（砷）
	CowHouse9   float32 `json:"cow_house_9"`  //Cu（铜）
	CowHouse10  float32 `json:"cow_house_10"` //Pb（铅）
	CowHouse11  float32 `json:"cow_house_11"` //Cr(铬)
	CowHouse12  float32 `json:"cow_house_12"` //Zn(锌）
}

type ReqAddPastureBasicEnvironment struct {
	TimeStamp    int64   `json:"time_stamp"`
	HouseNumber  string  `gorm:"not null; type:varchar(256)" json:"house_number"`
	Environment1 float32 `json:"environment_1"` //温度
	Environment2 float32 `json:"environment_2"` //相对湿度
	Environment3 float32 `json:"environment_3"` //风速
	Environment4 float32 `json:"environment_4"` //照度
	Environment5 float32 `json:"environment_5"` //噪声
	Environment6 float32 `json:"environment_6"` //光照时间
}

type ReqAddPasturePaddingRequire struct {
	TimeStamp       int64   `json:"time_stamp"`
	HouseNumber     string  `gorm:"not null; type:varchar(256)" json:"house_number"`
	PaddingRequire1 float32 `json:"padding_require_1"` //Hg（总汞）
	PaddingRequire2 float32 `json:"padding_require_2"` //Pb（铅）
	PaddingRequire3 float32 `json:"padding_require_3"` //Cr（铬）
	PaddingRequire4 float32 `json:"padding_require_4"` //Cd（镉）
	PaddingRequire5 float32 `json:"padding_require_5"` //TTC(总大肠菌群数)
	PaddingRequire6 float32 `json:"padding_require_6"` //TBC(细菌总数)
	PaddingRequire7 float32 `json:"padding_require_7"` //AFB1（黄曲霉素B1）
	PaddingRequire8 float32 `json:"padding_require_8"` //STC(沙门菌数)
}

type ReqAddPastureWastedWaterIndex struct {
	TimeStamp         int64   `json:"time_stamp"`
	HouseNumber       string  `gorm:"not null; type:varchar(256)" json:"house_number"`
	WastedWaterIndex1 float32 `json:"wasted_water_index_1"` //BOD(五日生化需氧量)
	WastedWaterIndex2 float32 `json:"wasted_water_index_2"` //COD(化学需氧量)
	WastedWaterIndex3 float32 `json:"wasted_water_index_3"` //NH3-N(氨氮)
	WastedWaterIndex4 float32 `json:"wasted_water_index_4"` //TP(总磷)
	WastedWaterIndex5 float32 `json:"wasted_water_index_5"` //TSS(悬浮物）
	WastedWaterIndex6 float32 `json:"wasted_water_index_6"` //FCC（粪大肠菌群个数）
	WastedWaterIndex7 float32 `json:"wasted_water_index_7"` //蛔虫卵个数
	WastedWaterIndex8 float32 `json:"wasted_water_index_8"` //PH值
	WastedWaterIndex9 float32 `json:"wasted_water_index_9"` //流量
}

type ReqAddPastureDisinfectionRecord struct {
	TimeStamp      int64  `json:"time_stamp"`
	HouseNumber    string `gorm:"not null; type:varchar(256)" json:"house_number"`
	FarmDisRecord1 string `json:"farm_dis_record_1"` //牧场消毒时间
	FarmDisRecord2 string `json:"farm_dis_record_2"` //消毒剂种类
	FarmDisRecord3 string `json:"farm_dis_record_3"` //消毒剂浓度
}

type ReqPastureSensorData struct {
	HouseNumber    string `json:"house_number"`
	StartTimestamp int64  `json:"start_timestamp"`
	EndTimestamp   int64  `json:"end_timestamp"`
}

type ReqPastureWasteWaterPerDay struct {
	TimeStamp                   int64   `json:"time_stamp"`
	HouseNumber                 string  `json:"house_number"`
	ReqPastureWasteWaterPerDay1 float32 `json:"req_pasture_waste_water_per_day_1"`
	ReqPastureWasteWaterPerDay2 float32 `json:"req_pasture_waste_water_per_day_2"`
	ReqPastureWasteWaterPerDay3 float32 `json:"req_pasture_waste_water_per_day_3"`
	ReqPastureWasteWaterPerDay4 float32 `json:"req_pasture_waste_water_per_day_4"`
}

type ReqPastureWasteResiduePerDay struct {
	TimeStamp                  int64   `json:"time_stamp"`
	HouseNumber                string  `gorm:"not null; type:varchar(256)" json:"house_number"`
	PastureWasteResiduePerDay1 float32 `json:"pasture_waste_residue_per_day_1"`
	PastureWasteResiduePerDay2 float32 `json:"pasture_waste_residue_per_day_2"`
	PastureWasteResiduePerDay3 float32 `json:"pasture_waste_residue_per_day_3"`
	PastureWasteResiduePerDay4 float32 `json:"pasture_waste_residue_per_day_4"`
}

type ReqPastureOdorPollutantsPerDay struct {
	TimeStamp                    int64   `json:"time_stamp"`
	HouseNumber                  string  `gorm:"not null; type:varchar(256)" json:"house_number"`
	PastureOdorPollutantsPerDay1 float32 `json:"pasture_odor_pollutants_per_day_1"`
	PastureOdorPollutantsPerDay2 float32 `json:"pasture_odor_pollutants_per_day_2"`
	PastureOdorPollutantsPerDay3 float32 `json:"pasture_odor_pollutants_per_day_3"`
	PastureOdorPollutantsPerDay4 float32 `json:"pasture_odor_pollutants_per_day_4"`
}

type ReqWasteResidueOdor struct {
	TimeStamp int64 `json:"time_stamp"`
}
