package pasture

import (
	"gorm.io/gorm"
	"time"
)

type PastureBuffer struct {
	//缓冲区各参数
	gorm.Model
	TimeRecordAt time.Time `json:"time_record_at"`
	HouseNumber  string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	Buffer1      float32   `json:"buffer_1"`  //NH3氨气
	Buffer2      float32   `json:"buffer_2"`  //H2S硫化氢
	Buffer3      float32   `json:"buffer_3"`  //CO2二氧化碳
	Buffer4      float32   `json:"buffer_4"`  //PM10
	Buffer5      float32   `json:"buffer_5"`  //TSP
	Buffer6      float32   `json:"buffer_6"`  //ODO 恶臭
	Buffer7      float32   `json:"buffer_7"`  //Cd(镉)
	Buffer8      float32   `json:"buffer_8"`  //As（砷）
	Buffer9      float32   `json:"buffer_9"`  //Cu（铜）
	Buffer10     float32   `json:"buffer_10"` //Pb（铅）
	Buffer11     float32   `json:"buffer_11"` //Cr(铬)
	Buffer12     float32   `json:"buffer_12"` //Zn(锌）
}

type PastureArea struct {
	//场区
	gorm.Model
	TimeRecordAt time.Time `json:"time_record_at"`
	HouseNumber  string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	CattleFarm1  float32   `json:"cattle_farm_1"`  //NH3氨气
	CattleFarm2  float32   `json:"cattle_farm_2"`  //H2S硫化氢
	CattleFarm3  float32   `json:"cattle_farm_3"`  //CO2二氧化碳
	CattleFarm4  float32   `json:"cattle_farm_4"`  //PM10
	CattleFarm5  float32   `json:"cattle_farm_5"`  //TSP
	CattleFarm6  float32   `json:"cattle_farm_6"`  //ODO 恶臭
	CattleFarm7  float32   `json:"cattle_farm_7"`  //Cd(镉)
	CattleFarm8  float32   `json:"cattle_farm_8"`  //As（砷）
	CattleFarm9  float32   `json:"cattle_farm_9"`  //Cu（铜）
	CattleFarm10 float32   `json:"cattle_farm_10"` //Pb（铅）
	CattleFarm11 float32   `json:"cattle_farm_11"` //Cr(铬)
	CattleFarm12 float32   `json:"cattle_farm_12"` //Zn(锌）
}

type CowHouse struct {
	//牛舍
	gorm.Model
	TimeRecordAt time.Time `json:"time_record_at"`
	HouseNumber  string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	CowHouse1    float32   `json:"cow_house_1"`  //NH3氨气
	CowHouse2    float32   `json:"cow_house_2"`  //H2S硫化氢
	CowHouse3    float32   `json:"cow_house_3"`  //CO2二氧化碳
	CowHouse4    float32   `json:"cow_house_4"`  //PM10
	CowHouse5    float32   `json:"cow_house_5"`  //TSP
	CowHouse6    float32   `json:"cow_house_6"`  //ODO 恶臭
	CowHouse7    float32   `json:"cow_house_7"`  //Cd(镉)
	CowHouse8    float32   `json:"cow_house_8"`  //As（砷）
	CowHouse9    float32   `json:"cow_house_9"`  //Cu（铜）
	CowHouse10   float32   `json:"cow_house_10"` //Pb（铅）
	CowHouse11   float32   `json:"cow_house_11"` //Cr(铬)
	CowHouse12   float32   `json:"cow_house_12"` //Zn(锌）
}

type PastureBasicEnvironment struct {
	//常见环境因素
	gorm.Model
	TimeRecordAt time.Time `json:"time_record_at"`
	HouseNumber  string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	Environment1 float32   `json:"environment_1"` //温度
	Environment2 float32   `json:"environment_2"` //相对湿度
	Environment3 float32   `json:"environment_3"` //风速
	Environment4 float32   `json:"environment_4"` //照度
	Environment5 float32   `json:"environment_5"` //噪声
	Environment6 float32   `json:"environment_6"` //光照时间
}

type PasturePaddingRequire struct {
	//牧场垫料要求
	gorm.Model
	TimeRecordAt    time.Time `json:"time_record_at"`
	HouseNumber     string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	PaddingRequire1 float32   `json:"padding_require_1"` //T-Hg（总汞）
	PaddingRequire2 float32   `json:"padding_require_2"` //Pb（铅）
	PaddingRequire3 float32   `json:"padding_require_3"` //Cr（铬）
	PaddingRequire4 float32   `json:"padding_require_4"` //Cd（镉）
	PaddingRequire5 float32   `json:"padding_require_5"` //TTC(总大肠菌群数)
	PaddingRequire6 float32   `json:"padding_require_6"` //TBC(细菌总数)
	PaddingRequire7 float32   `json:"padding_require_7"` //AFB1（黄曲霉素B1）
	PaddingRequire8 float32   `json:"padding_require_8"` //STC(沙门菌数)
}

type PastureWastedWaterIndex struct {
	//牧场废水指标
	gorm.Model
	TimeRecordAt      time.Time `json:"time_record_at"`
	HouseNumber       string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	WastedWaterIndex1 float32   `json:"wasted_water_index_1"` //BOD(五日生化需氧量)
	WastedWaterIndex2 float32   `json:"wasted_water_index_2"` //COD(化学需氧量)
	WastedWaterIndex3 float32   `json:"wasted_water_index_3"` //NH3-N(氨氮)
	WastedWaterIndex4 float32   `json:"wasted_water_index_4"` //TP(总磷)
	WastedWaterIndex5 float32   `json:"wasted_water_index_5"` //TSS(悬浮物）
	WastedWaterIndex6 float32   `json:"wasted_water_index_6"` //FCC（粪大肠菌群个数）
	WastedWaterIndex7 float32   `json:"wasted_water_index_7"` //蛔虫卵个数
	WastedWaterIndex8 float32   `json:"wasted_water_index_8"` //PH值
	WastedWaterIndex9 float32   `json:"wasted_water_index_9"` //流量
}

type PastureDisinfectionRecord struct {
	//牧场消毒记录
	gorm.Model
	TimeRecordAt   time.Time `json:"time_record_at"`
	HouseNumber    string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	FarmDisRecord1 string    `json:"farm_dis_record_1"` //牧场消毒时间
	FarmDisRecord2 string    `json:"farm_dis_record_2"` //消毒剂种类
	FarmDisRecord3 string    `json:"farm_dis_record_3"` //消毒剂浓度
}
