package pasture

import (
	"gorm.io/gorm"
	"time"
)

type TotalWastedWaterPasturePerDay struct {
	//当日牧场污水处理总量
	gorm.Model
	TimeStamp               time.Time `json:"time_stamp"`
	HouseNumber             string    `json:"house_number"`
	TotalWastedWaterPerDay1 float64   `json:"total_wasted_water_per_day_1"` //正常污水量
	TotalWastedWaterPerDay2 float64   `json:"total_wasted_water_per_day_2"` //正常污水新增
	TotalWastedWaterPerDay3 float64   `json:"total_wasted_water_per_day_3"` //污染超标量
	TotalWastedWaterPerDay4 float64   `json:"total_wasted_water_per_day_4"` //污水新增
}

type TotalWasteResiduePasturePerDay struct {
	//牧场当日废渣处理总量
	gorm.Model
	TimeStamp               time.Time `json:"time_stamp"`
	HouseNumber             string    `json:"house_number"`
	TotalWastedWaterPerDay1 float64   `json:"total_wasted_water_per_day_1"` //正常废渣量
	TotalWastedWaterPerDay2 float64   `json:"total_wasted_water_per_day_2"` //正常废渣增量
	TotalWastedWaterPerDay3 float64   `json:"total_wasted_water_per_day_3"` //污染废渣超标量
	TotalWastedWaterPerDay4 float64   `json:"total_wasted_water_per_day_4"` //废渣新增
}

type TotalOdorPollutantsPasturePerDay struct {
	//当日牧场恶臭污染物处理总量
	gorm.Model
	TimeStamp                  time.Time `json:"time_stamp"`
	HouseNumber                string    `json:"house_number"`
	TotalOdorPollutantsPerDay1 float64   `json:"total_odor_pollutants_per_day_1"` //正常恶臭污染物处理总量
	TotalOdorPollutantsPerDay2 float64   `json:"total_odor_pollutants_per_day_2"` //较昨日新增
	TotalOdorPollutantsPerDay3 float64   `json:"total_odor_pollutants_per_day_3"` //超标恶臭污染物处理总量
	TotalOdorPollutantsPerDay4 float64   `json:"total_odor_pollutants_per_day_4"` //超标较昨日新增
}

type AllPasturesTrashDisposal struct {
	//所有牧场当日垃圾处理信息
	gorm.Model
	TimeStamp                     time.Time `gorm:"unique; not null" json:"time_stamp"`
	OdorPasturesTrashDisposal1    float64   `json:"odor_pastures_trash_disposal_1"` //正常总
	OdorPasturesTrashDisposal2    float64   `json:"odor_pastures_trash_disposal_2"` //正常较昨日新增
	OdorPasturesTrashDisposal3    float64   `json:"odor_pastures_trash_disposal_3"` //超标总
	OdorPasturesTrashDisposal4    float64   `json:"odor_pastures_trash_disposal_4"` //超标较昨日新增
	ResiduePasturesTrashDisposal1 float64   `json:"residue_pastures_trash_disposal_1"`
	ResiduePasturesTrashDisposal2 float64   `json:"residue_pastures_trash_disposal_2"`
	ResiduePasturesTrashDisposal3 float64   `json:"residue_pastures_trash_disposal_3"`
	ResiduePasturesTrashDisposal4 float64   `json:"residue_pastures_trash_disposal_4"`
	WaterPasturesTrashDisposal1   float64   `json:"water_pastures_trash_disposal_1"`
	WaterPasturesTrashDisposal2   float64   `json:"water_pastures_trash_disposal_2"`
	WaterPasturesTrashDisposal3   float64   `json:"water_pastures_trash_disposal_3"`
	WaterPasturesTrashDisposal4   float64   `json:"water_pastures_trash_disposal_4"`
}

type AllPasturesTrashDisposalInfo struct {
	TimeStamp                     string  `json:"time_stamp"`
	OdorPasturesTrashDisposal1    float64 `json:"odor_pastures_trash_disposal_1"` //正常总
	OdorPasturesTrashDisposal2    float64 `json:"odor_pastures_trash_disposal_2"` //正常较昨日新增
	OdorPasturesTrashDisposal3    float64 `json:"odor_pastures_trash_disposal_3"` //超标总
	OdorPasturesTrashDisposal4    float64 `json:"odor_pastures_trash_disposal_4"` //超标较昨日新增
	ResiduePasturesTrashDisposal1 float64 `json:"residue_pastures_trash_disposal_1"`
	ResiduePasturesTrashDisposal2 float64 `json:"residue_pastures_trash_disposal_2"`
	ResiduePasturesTrashDisposal3 float64 `json:"residue_pastures_trash_disposal_3"`
	ResiduePasturesTrashDisposal4 float64 `json:"residue_pastures_trash_disposal_4"`
	WaterPasturesTrashDisposal1   float64 `json:"water_pastures_trash_disposal_1"`
	WaterPasturesTrashDisposal2   float64 `json:"water_pastures_trash_disposal_2"`
	WaterPasturesTrashDisposal3   float64 `json:"water_pastures_trash_disposal_3"`
	WaterPasturesTrashDisposal4   float64 `json:"water_pastures_trash_disposal_4"`
}

func ToAllPasturesTrashDisposalInfo(pa *AllPasturesTrashDisposal) AllPasturesTrashDisposalInfo {
	return AllPasturesTrashDisposalInfo{
		TimeStamp:                     pa.TimeStamp.Format("2006-01-02 15:04:05"),
		OdorPasturesTrashDisposal1:    pa.OdorPasturesTrashDisposal1,
		OdorPasturesTrashDisposal2:    pa.OdorPasturesTrashDisposal2,
		OdorPasturesTrashDisposal3:    pa.OdorPasturesTrashDisposal3,
		OdorPasturesTrashDisposal4:    pa.OdorPasturesTrashDisposal4,
		ResiduePasturesTrashDisposal1: pa.ResiduePasturesTrashDisposal1,
		ResiduePasturesTrashDisposal2: pa.ResiduePasturesTrashDisposal2,
		ResiduePasturesTrashDisposal3: pa.ResiduePasturesTrashDisposal3,
		ResiduePasturesTrashDisposal4: pa.ResiduePasturesTrashDisposal4,
		WaterPasturesTrashDisposal1:   pa.WaterPasturesTrashDisposal1,
		WaterPasturesTrashDisposal2:   pa.WaterPasturesTrashDisposal2,
		WaterPasturesTrashDisposal3:   pa.WaterPasturesTrashDisposal3,
		WaterPasturesTrashDisposal4:   pa.WaterPasturesTrashDisposal4,
	}
}
