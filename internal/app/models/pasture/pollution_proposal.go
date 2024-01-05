package pasture

import (
	"gorm.io/gorm"
	"time"
)

type TotalWastedWaterPasturePerDay struct {
	//当日牧场污水处理总量
	gorm.Model
	TimeStamp               time.Time `json:"time_stamp"`
	TotalWastedWaterPerDay1 float32   `json:"total_wasted_water_per_day_1"` //正常污水量
	TotalWastedWaterPerDay2 float32   `json:"total_wasted_water_per_day_2"` //正常污水新增
	TotalWastedWaterPerDay3 float32   `json:"total_wasted_water_per_day_3"` //污染超标量
	TotalWastedWaterPerDay4 float32   `json:"total_wasted_water_per_day_4"` //污水新增
}

type TotalWasteResiduePasturePerDay struct {
	//牧场当日废渣处理总量
	gorm.Model
	TimeStamp               time.Time `json:"time_stamp"`
	TotalWastedWaterPerDay1 float32   `json:"total_wasted_water_per_day_1"` //正常废渣量
	TotalWastedWaterPerDay2 float32   `json:"total_wasted_water_per_day_2"` //正常废渣增量
	TotalWastedWaterPerDay3 float32   `json:"total_wasted_water_per_day_3"` //污染废渣超标量
	TotalWastedWaterPerDay4 float32   `json:"total_wasted_water_per_day_4"` //废渣新增
}

type TotalOdorPollutantsPasturePerDay struct {
	//当日牧场恶臭污染物处理总量
	gorm.Model
	TimeStamp                  time.Time `json:"time_stamp"`
	TotalOdorPollutantsPerDay1 float32   `json:"total_odor_pollutants_per_day_1"` //正常恶臭污染物处理总量
	TotalOdorPollutantsPerDay2 float32   `json:"total_odor_pollutants_per_day_2"` //较昨日新增
	TotalOdorPollutantsPerDay3 float32   `json:"total_odor_pollutants_per_day_3"` //超标恶臭污染物处理总量
	TotalOdorPollutantsPerDay4 float32   `json:"total_odor_pollutants_per_day_4"` //超标较昨日新增
}
