package slaughter

import (
	"time"

	"gorm.io/gorm"
)

type TotalWasteWaterSlaughterPerDay struct {
	//特定屠宰场当日污水处理总量
	gorm.Model
	TimeStamp                       time.Time `json:"time_stamp"`
	HouseNumber                     string    `json:"house_number"`
	TotalWasteWaterSlaughterPerDay1 float64   `json:"total_waste_water_slaughter_per_day_1"` //正常总
	TotalWasteWaterSlaughterPerDay2 float64   `json:"total_waste_water_slaughter_per_day_2"` //正常较昨日
	TotalWasteWaterSlaughterPerDay3 float64   `json:"total_waste_water_slaughter_per_day_3"` //污染总
	TotalWasteWaterSlaughterPerDay4 float64   `json:"total_waste_water_slaughter_per_day_4"` //污染超昨日
}

type TotalWasteResidueSlaughterPerDay struct {
	//特定屠宰场当日废渣处理总量
	gorm.Model
	TimeStamp                         time.Time `json:"time_stamp"`
	HouseNumber                       string    `json:"house_number"`
	TotalWasteResidueSlaughterPerDay1 float64   `json:"total_waste_residue_slaughter_per_day_1"` //正常废渣量
	TotalWasteResidueSlaughterPerDay2 float64   `json:"total_waste_residue_slaughter_per_day_2"` //正常较阻容
	TotalWasteResidueSlaughterPerDay3 float64   `json:"total_waste_residue_slaughter_per_day_3"` //超标总
	TotalWasteResidueSlaughterPerDay4 float64   `json:"total_waste_residue_slaughter_per_day_4"` //超标较昨日
}

type TotalOdorPollutantsSlaughterPerDay struct {
	//特定屠宰场当日恶臭污染物处理总量
	gorm.Model
	TimeStamp                           time.Time `json:"time_stamp"`
	HouseNumber                         string    `json:"house_number"`
	TotalOdorPollutantsSlaughterPerDay1 float64   `json:"total_odor_pollutants_slaughter_per_day_1"` //正常总
	TotalOdorPollutantsSlaughterPerDay2 float64   `json:"total_odor_pollutants_slaughter_per_day_2"` //正常较昨日
	TotalOdorPollutantsSlaughterPerDay3 float64   `json:"total_odor_pollutants_slaughter_per_day_3"` //超标总
	TotalOdorPollutantsSlaughterPerDay4 float64   `json:"total_odor_pollutants_slaughter_per_day_4"` //超标较昨日
}

type AllSlaughtersTrashDisposal struct {
	//所有屠宰场当日垃圾处理信息
	gorm.Model
	TimeStamp                       time.Time `gorm:"unique; not null" json:"time_stamp"`
	OdorAllSlaughtersTrashDisposal1 float64   `json:"odor_all_slaughters_trash_disposal_1"`
	OdorAllSlaughtersTrashDisposal2 float64   `json:"odor_all_slaughters_trash_disposal_2"`
	OdorAllSlaughtersTrashDisposal3 float64   `json:"odor_all_slaughters_trash_disposal_3"`
	OdorAllSlaughtersTrashDisposal4 float64   `json:"odor_all_slaughters_trash_disposal_4"`
	ResidueSlaughtersTrashDisposal1 float64   `json:"residue_slaughters_trash_disposal_1"`
	ResidueSlaughtersTrashDisposal2 float64   `json:"residue_slaughters_trash_disposal_2"`
	ResidueSlaughtersTrashDisposal3 float64   `json:"residue_slaughters_trash_disposal_3"`
	ResidueSlaughtersTrashDisposal4 float64   `json:"residue_slaughters_trash_disposal_4"`
	WaterSlaughtersTrashDisposal1   float64   `json:"water_slaughters_trash_disposal_1"`
	WaterSlaughtersTrashDisposal2   float64   `json:"water_slaughters_trash_disposal_2"`
	WaterSlaughtersTrashDisposal3   float64   `json:"water_slaughters_trash_disposal_3"`
	WaterSlaughtersTrashDisposal4   float64   `json:"water_slaughters_trash_disposal_4"`
}
