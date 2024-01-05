package slaughter

import (
	"gorm.io/gorm"
	"time"
)

type TotalWasteWaterSlaughterPerDay struct {
	//当日屠宰场污水处理总量
	gorm.Model
	TimeStamp                       time.Time `json:"time_stamp"`
	TotalWasteWaterSlaughterPerDay1 float32   `json:"total_waste_water_slaughter_per_day_1"`
	TotalWasteWaterSlaughterPerDay2 float32   `json:"total_waste_water_slaughter_per_day_2"`
	TotalWasteWaterSlaughterPerDay3 float32   `json:"total_waste_water_slaughter_per_day_3"`
	TotalWasteWaterSlaughterPerDay4 float32   `json:"total_waste_water_slaughter_per_day_4"`
}

type TotalWasteResidueSlaughterPerDay struct {
	//屠宰场当日废渣处理总量
	gorm.Model
	TimeStamp                         time.Time `json:"time_stamp"`
	TotalWasteResidueSlaughterPerDay1 float32   `json:"total_waste_residue_slaughter_per_day_1"` //正常废渣量
	TotalWasteResidueSlaughterPerDay2 float32   `json:"total_waste_residue_slaughter_per_day_2"`
	TotalWasteResidueSlaughterPerDay3 float32   `json:"total_waste_residue_slaughter_per_day_3"`
	TotalWasteResidueSlaughterPerDay4 float32   `json:"total_waste_residue_slaughter_per_day_4"`
}

type TotalOdorPollutantsSlaughterPerDay struct {
	//屠宰场当日恶臭污染物处理总量
	gorm.Model
	TimeStamp                           time.Time `json:"time_stamp"`
	TotalOdorPollutantsSlaughterPerDay1 float32   `json:"total_odor_pollutants_slaughter_per_day_1"`
	TotalOdorPollutantsSlaughterPerDay2 float32   `json:"total_odor_pollutants_slaughter_per_day_2"`
	TotalOdorPollutantsSlaughterPerDay3 float32   `json:"total_odor_pollutants_slaughter_per_day_3"`
	TotalOdorPollutantsSlaughterPerDay4 float32   `json:"total_odor_pollutants_slaughter_per_day_4"`
}
