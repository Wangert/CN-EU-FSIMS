package pasture

import (
	"time"

	"gorm.io/gorm"
)

type PastureBuffer struct {
	//缓冲区各参数
	gorm.Model
	TimeRecordAt time.Time `json:"time_record_at"`
	HouseNumber  string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	Buffer1      float64   `json:"buffer_1"`  //NH3氨气
	Buffer2      float64   `json:"buffer_2"`  //H2S硫化氢
	Buffer3      float64   `json:"buffer_3"`  //CO2二氧化碳
	Buffer4      float64   `json:"buffer_4"`  //PM10
	Buffer5      float64   `json:"buffer_5"`  //TSP
	Buffer6      float64   `json:"buffer_6"`  //ODO 恶臭
	Buffer7      float64   `json:"buffer_7"`  //Cd(镉)
	Buffer8      float64   `json:"buffer_8"`  //As（砷）
	Buffer9      float64   `json:"buffer_9"`  //Cu（铜）
	Buffer10     float64   `json:"buffer_10"` //Pb（铅）
	Buffer11     float64   `json:"buffer_11"` //Cr(铬)
	Buffer12     float64   `json:"buffer_12"` //Zn(锌）
}

type PastureBufferInfo struct {
	TimeRecordAt string  `json:"time_record_at"`
	HouseNumber  string  `gorm:"not null; type:varchar(256)" json:"house_number"`
	Buffer1      float64 `json:"buffer_1"`  //NH3氨气
	Buffer2      float64 `json:"buffer_2"`  //H2S硫化氢
	Buffer3      float64 `json:"buffer_3"`  //CO2二氧化碳
	Buffer4      float64 `json:"buffer_4"`  //PM10
	Buffer5      float64 `json:"buffer_5"`  //TSP
	Buffer6      float64 `json:"buffer_6"`  //ODO 恶臭
	Buffer7      float64 `json:"buffer_7"`  //Cd(镉)
	Buffer8      float64 `json:"buffer_8"`  //As（砷）
	Buffer9      float64 `json:"buffer_9"`  //Cu（铜）
	Buffer10     float64 `json:"buffer_10"` //Pb（铅）
	Buffer11     float64 `json:"buffer_11"` //Cr(铬)
	Buffer12     float64 `json:"buffer_12"` //Zn(锌）
}

func ToPastureBufferInfo(bu *PastureBuffer) PastureBufferInfo {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time := bu.TimeRecordAt.In(loc)
	return PastureBufferInfo{
		TimeRecordAt: time.Format("2006-01-02 15:04:05"),
		HouseNumber:  bu.HouseNumber,
		Buffer1:      bu.Buffer1,
		Buffer2:      bu.Buffer2,
		Buffer3:      bu.Buffer3,
		Buffer4:      bu.Buffer4,
		Buffer5:      bu.Buffer5,
		Buffer6:      bu.Buffer6,
		Buffer7:      bu.Buffer7,
		Buffer8:      bu.Buffer8,
		Buffer9:      bu.Buffer9,
		Buffer10:     bu.Buffer10,
		Buffer11:     bu.Buffer11,
		Buffer12:     bu.Buffer12,
	}
}

type PastureArea struct {
	//场区
	gorm.Model
	TimeRecordAt time.Time `json:"time_record_at"`
	HouseNumber  string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	CattleFarm1  float64   `json:"cattle_farm_1"`  //NH3氨气
	CattleFarm2  float64   `json:"cattle_farm_2"`  //H2S硫化氢
	CattleFarm3  float64   `json:"cattle_farm_3"`  //CO2二氧化碳
	CattleFarm4  float64   `json:"cattle_farm_4"`  //PM10
	CattleFarm5  float64   `json:"cattle_farm_5"`  //TSP
	CattleFarm6  float64   `json:"cattle_farm_6"`  //ODO 恶臭
	CattleFarm7  float64   `json:"cattle_farm_7"`  //Cd(镉)
	CattleFarm8  float64   `json:"cattle_farm_8"`  //As（砷）
	CattleFarm9  float64   `json:"cattle_farm_9"`  //Cu（铜）
	CattleFarm10 float64   `json:"cattle_farm_10"` //Pb（铅）
	CattleFarm11 float64   `json:"cattle_farm_11"` //Cr(铬)
	CattleFarm12 float64   `json:"cattle_farm_12"` //Zn(锌）
}

type PastureAreaInfo struct {
	TimeRecordAt string  `json:"time_record_at"`
	HouseNumber  string  `gorm:"not null; type:varchar(256)" json:"house_number"`
	CattleFarm1  float64 `json:"cattle_farm_1"`  //NH3氨气
	CattleFarm2  float64 `json:"cattle_farm_2"`  //H2S硫化氢
	CattleFarm3  float64 `json:"cattle_farm_3"`  //CO2二氧化碳
	CattleFarm4  float64 `json:"cattle_farm_4"`  //PM10
	CattleFarm5  float64 `json:"cattle_farm_5"`  //TSP
	CattleFarm6  float64 `json:"cattle_farm_6"`  //ODO 恶臭
	CattleFarm7  float64 `json:"cattle_farm_7"`  //Cd(镉)
	CattleFarm8  float64 `json:"cattle_farm_8"`  //As（砷）
	CattleFarm9  float64 `json:"cattle_farm_9"`  //Cu（铜）
	CattleFarm10 float64 `json:"cattle_farm_10"` //Pb（铅）
	CattleFarm11 float64 `json:"cattle_farm_11"` //Cr(铬)
	CattleFarm12 float64 `json:"cattle_farm_12"` //Zn(锌）
}

func ToPastureAreaInfo(ar *PastureArea) PastureAreaInfo {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time := ar.TimeRecordAt.In(loc)
	return PastureAreaInfo{
		TimeRecordAt: time.Format("2006-01-02 15:04:05"),
		HouseNumber:  ar.HouseNumber,
		CattleFarm1:  ar.CattleFarm1,
		CattleFarm2:  ar.CattleFarm2,
		CattleFarm3:  ar.CattleFarm3,
		CattleFarm4:  ar.CattleFarm4,
		CattleFarm5:  ar.CattleFarm5,
		CattleFarm6:  ar.CattleFarm6,
		CattleFarm7:  ar.CattleFarm7,
		CattleFarm8:  ar.CattleFarm8,
		CattleFarm9:  ar.CattleFarm9,
		CattleFarm10: ar.CattleFarm10,
		CattleFarm11: ar.CattleFarm11,
		CattleFarm12: ar.CattleFarm12,
	}
}

type CowHouse struct {
	//牛舍
	gorm.Model
	TimeRecordAt time.Time `json:"time_record_at"`
	HouseNumber  string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	CowHouse1    float64   `json:"cow_house_1"`  //NH3氨气
	CowHouse2    float64   `json:"cow_house_2"`  //H2S硫化氢
	CowHouse3    float64   `json:"cow_house_3"`  //CO2二氧化碳
	CowHouse4    float64   `json:"cow_house_4"`  //PM10
	CowHouse5    float64   `json:"cow_house_5"`  //TSP
	CowHouse6    float64   `json:"cow_house_6"`  //ODO 恶臭
	CowHouse7    float64   `json:"cow_house_7"`  //Cd(镉)
	CowHouse8    float64   `json:"cow_house_8"`  //As（砷）
	CowHouse9    float64   `json:"cow_house_9"`  //Cu（铜）
	CowHouse10   float64   `json:"cow_house_10"` //Pb（铅）
	CowHouse11   float64   `json:"cow_house_11"` //Cr(铬)
	CowHouse12   float64   `json:"cow_house_12"` //Zn(锌）
}

type CowHouseInfo struct {
	TimeRecordAt string  `json:"time_record_at"`
	HouseNumber  string  `gorm:"not null; type:varchar(256)" json:"house_number"`
	CowHouse1    float64 `json:"cow_house_1"`  //NH3氨气
	CowHouse2    float64 `json:"cow_house_2"`  //H2S硫化氢
	CowHouse3    float64 `json:"cow_house_3"`  //CO2二氧化碳
	CowHouse4    float64 `json:"cow_house_4"`  //PM10
	CowHouse5    float64 `json:"cow_house_5"`  //TSP
	CowHouse6    float64 `json:"cow_house_6"`  //ODO 恶臭
	CowHouse7    float64 `json:"cow_house_7"`  //Cd(镉)
	CowHouse8    float64 `json:"cow_house_8"`  //As（砷）
	CowHouse9    float64 `json:"cow_house_9"`  //Cu（铜）
	CowHouse10   float64 `json:"cow_house_10"` //Pb（铅）
	CowHouse11   float64 `json:"cow_house_11"` //Cr(铬)
	CowHouse12   float64 `json:"cow_house_12"` //Zn(锌）
}

func ToCowHouseInfo(ch *CowHouse) CowHouseInfo {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time := ch.TimeRecordAt.In(loc)
	return CowHouseInfo{
		TimeRecordAt: time.Format("2006-01-02 15:04:05"),
		HouseNumber:  ch.HouseNumber,
		CowHouse1:    ch.CowHouse1,
		CowHouse2:    ch.CowHouse2,
		CowHouse3:    ch.CowHouse3,
		CowHouse4:    ch.CowHouse4,
		CowHouse5:    ch.CowHouse5,
		CowHouse6:    ch.CowHouse6,
		CowHouse7:    ch.CowHouse7,
		CowHouse8:    ch.CowHouse8,
		CowHouse9:    ch.CowHouse9,
		CowHouse10:   ch.CowHouse10,
		CowHouse11:   ch.CowHouse11,
		CowHouse12:   ch.CowHouse12,
	}
}

type PastureBasicEnvironment struct {
	//常见环境因素
	gorm.Model
	TimeRecordAt time.Time `json:"time_record_at"`
	HouseNumber  string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	Environment1 float64   `json:"environment_1"` //温度
	Environment2 float64   `json:"environment_2"` //相对湿度
	Environment3 float64   `json:"environment_3"` //风速
	Environment4 float64   `json:"environment_4"` //照度
	Environment5 float64   `json:"environment_5"` //噪声
	Environment6 float64   `json:"environment_6"` //光照时间
}

type PastureBasicEnvironmentInfo struct {
	TimeRecordAt string  `json:"time_record_at"`
	HouseNumber  string  `gorm:"not null; type:varchar(256)" json:"house_number"`
	Environment1 float64 `json:"environment_1"` //温度
	Environment2 float64 `json:"environment_2"` //相对湿度
	Environment3 float64 `json:"environment_3"` //风速
	Environment4 float64 `json:"environment_4"` //照度
	Environment5 float64 `json:"environment_5"` //噪声
	Environment6 float64 `json:"environment_6"` //光照时间
}

func ToPastureBasicEnvironmentInfo(be *PastureBasicEnvironment) PastureBasicEnvironmentInfo {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time := be.TimeRecordAt.In(loc)
	return PastureBasicEnvironmentInfo{
		TimeRecordAt: time.Format("2006-01-02 15:04:05"),
		HouseNumber:  be.HouseNumber,
		Environment1: be.Environment1,
		Environment2: be.Environment2,
		Environment3: be.Environment3,
		Environment4: be.Environment4,
		Environment5: be.Environment5,
		Environment6: be.Environment6,
	}
}

type PasturePaddingRequire struct {
	//牧场垫料要求
	gorm.Model
	TimeRecordAt    time.Time `json:"time_record_at"`
	HouseNumber     string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	PaddingRequire1 float64   `json:"padding_require_1"` //T-Hg（总汞）
	PaddingRequire2 float64   `json:"padding_require_2"` //Pb（铅）
	PaddingRequire3 float64   `json:"padding_require_3"` //Cr（铬）
	PaddingRequire4 float64   `json:"padding_require_4"` //Cd（镉）
	PaddingRequire5 float64   `json:"padding_require_5"` //TTC(总大肠菌群数)
	PaddingRequire6 float64   `json:"padding_require_6"` //TBC(细菌总数)
	PaddingRequire7 float64   `json:"padding_require_7"` //AFB1（黄曲霉素B1）
	PaddingRequire8 float64   `json:"padding_require_8"` //STC(沙门菌数)
}

type PasturePaddingRequireInfo struct {
	TimeRecordAt    string  `json:"time_record_at"`
	HouseNumber     string  `gorm:"not null; type:varchar(256)" json:"house_number"`
	PaddingRequire1 float64 `json:"padding_require_1"` //T-Hg（总汞）
	PaddingRequire2 float64 `json:"padding_require_2"` //Pb（铅）
	PaddingRequire3 float64 `json:"padding_require_3"` //Cr（铬）
	PaddingRequire4 float64 `json:"padding_require_4"` //Cd（镉）
	PaddingRequire5 float64 `json:"padding_require_5"` //TTC(总大肠菌群数)
	PaddingRequire6 float64 `json:"padding_require_6"` //TBC(细菌总数)
	PaddingRequire7 float64 `json:"padding_require_7"` //AFB1（黄曲霉素B1）
	PaddingRequire8 float64 `json:"padding_require_8"` //STC(沙门菌数)
}

func ToPasturePaddingRequireInfo(pr *PasturePaddingRequire) PasturePaddingRequireInfo {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time := pr.TimeRecordAt.In(loc)
	return PasturePaddingRequireInfo{
		TimeRecordAt:    time.Format("2006-01-02 15:04:05"),
		HouseNumber:     pr.HouseNumber,
		PaddingRequire1: pr.PaddingRequire1,
		PaddingRequire2: pr.PaddingRequire2,
		PaddingRequire3: pr.PaddingRequire3,
		PaddingRequire4: pr.PaddingRequire4,
		PaddingRequire5: pr.PaddingRequire5,
		PaddingRequire6: pr.PaddingRequire6,
		PaddingRequire7: pr.PaddingRequire7,
		PaddingRequire8: pr.PaddingRequire8,
	}
}

type PastureWastedWaterIndex struct {
	//牧场废水指标
	gorm.Model
	TimeRecordAt      time.Time `json:"time_record_at"`
	HouseNumber       string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	WastedWaterIndex1 float64   `json:"wasted_water_index_1"` //BOD5(五日生化需氧量)
	WastedWaterIndex2 float64   `json:"wasted_water_index_2"` //COD(化学需氧量)
	WastedWaterIndex3 float64   `json:"wasted_water_index_3"` //NH3-N(氨氮)
	WastedWaterIndex4 float64   `json:"wasted_water_index_4"` //TP(总磷)
	WastedWaterIndex5 float64   `json:"wasted_water_index_5"` //TSS(悬浮物）
	WastedWaterIndex6 float64   `json:"wasted_water_index_6"` //FCC（粪大肠菌群个数）
	WastedWaterIndex7 float64   `json:"wasted_water_index_7"` //蛔虫卵个数
	WastedWaterIndex8 float64   `json:"wasted_water_index_8"` //PH值
	WastedWaterIndex9 float64   `json:"wasted_water_index_9"` //流量
}

type PastureWastedWaterIndexInfo struct {
	TimeRecordAt      string  `json:"time_record_at"`
	HouseNumber       string  `gorm:"not null; type:varchar(256)" json:"house_number"`
	WastedWaterIndex1 float64 `json:"wasted_water_index_1"` //BOD(五日生化需氧量)
	WastedWaterIndex2 float64 `json:"wasted_water_index_2"` //COD(化学需氧量)
	WastedWaterIndex3 float64 `json:"wasted_water_index_3"` //NH3-N(氨氮)
	WastedWaterIndex4 float64 `json:"wasted_water_index_4"` //TP(总磷)
	WastedWaterIndex5 float64 `json:"wasted_water_index_5"` //TSS(悬浮物）
	WastedWaterIndex6 float64 `json:"wasted_water_index_6"` //FCC（粪大肠菌群个数）
	WastedWaterIndex7 float64 `json:"wasted_water_index_7"` //蛔虫卵个数
	WastedWaterIndex8 float64 `json:"wasted_water_index_8"` //PH值
	WastedWaterIndex9 float64 `json:"wasted_water_index_9"` //流量
}

func ToPastureWastedWaterIndexInfo(ww *PastureWastedWaterIndex) PastureWastedWaterIndexInfo {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time := ww.TimeRecordAt.In(loc)
	return PastureWastedWaterIndexInfo{
		TimeRecordAt:      time.Format("2006-01-02 15:04:05"),
		HouseNumber:       ww.HouseNumber,
		WastedWaterIndex1: ww.WastedWaterIndex1,
		WastedWaterIndex2: ww.WastedWaterIndex2,
		WastedWaterIndex3: ww.WastedWaterIndex3,
		WastedWaterIndex4: ww.WastedWaterIndex4,
		WastedWaterIndex5: ww.WastedWaterIndex5,
		WastedWaterIndex6: ww.WastedWaterIndex6,
		WastedWaterIndex7: ww.WastedWaterIndex7,
		WastedWaterIndex8: ww.WastedWaterIndex8,
		WastedWaterIndex9: ww.WastedWaterIndex9,
	}
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

type PastureDisinfectionRecordInfo struct {
	TimeRecordAt   string `json:"time_record_at"`
	HouseNumber    string `gorm:"not null; type:varchar(256)" json:"house_number"`
	FarmDisRecord1 string `json:"farm_dis_record_1"` //牧场消毒时间
	FarmDisRecord2 string `json:"farm_dis_record_2"` //消毒剂种类
	FarmDisRecord3 string `json:"farm_dis_record_3"` //消毒剂浓度
}

func ToPastureDisinfectionRecordInfo(dr *PastureDisinfectionRecord) PastureDisinfectionRecordInfo {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time := dr.TimeRecordAt.In(loc)
	return PastureDisinfectionRecordInfo{
		TimeRecordAt:   time.Format("2006-01-02 15:04:05"),
		HouseNumber:    dr.HouseNumber,
		FarmDisRecord1: dr.FarmDisRecord1,
		FarmDisRecord2: dr.FarmDisRecord2,
		FarmDisRecord3: dr.FarmDisRecord3,
	}
}
