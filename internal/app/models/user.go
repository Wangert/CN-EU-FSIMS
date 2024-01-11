package models

import (
	"gorm.io/gorm"
)

// user table
type FSIMSUser struct {
	gorm.Model
	UUID         string `gorm:"not null" json:"uuid"`
	Name         string `gorm:"not null" json:"name"`
	Account      string `gorm:"not null" json:"account"`
	PasswordHash string `gorm:"not null" json:"password_hash"`
	Type         int    `json:"type"`
	Role         string `json:"role"`
	HouseNumber  string `json:"houseNumber"`
	Status       int    `gorm:"not null" json:"status"`
	Company      string `gorm:"not null" json:"company"`
	Phone        string `gorm:"not null" json:"phone"`
}

type UserInfo struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Account string `json:"account"`
	Type    int    `json:"type"`
	Role    string `json:"role"`
	Status  int    `json:"status"`
	Company string `json:"company"`
	Phone   string `json:"phone"`
}

type TrashDisposalPerDayInfo struct {
	timeStamp                       string  `json:"time_stamp"`
	TrashDisposalPerDayWaterInfo1   float32 `json:"trash_disposal_per_day_water_info_1"`
	TrashDisposalPerDayWaterInfo2   float32 `json:"trash_disposal_per_day_water_info_2"`
	TrashDisposalPerDayWaterInfo3   float32 `json:"trash_disposal_per_day_water_info_3"`
	TrashDisposalPerDayWaterInfo4   float32 `json:"trash_disposal_per_day_water_info_4"`
	TrashDisposalPerDayResidueInfo1 float32 `json:"trash_disposal_per_day_residue_info_1"`
	TrashDisposalPerDayResidueInfo2 float32 `json:"trash_disposal_per_day_residue_info_2"`
	TrashDisposalPerDayResidueInfo3 float32 `json:"trash_disposal_per_day_residue_info_3"`
	TrashDisposalPerDayResidueInfo4 float32 `json:"trash_disposal_per_day_residue_info_4"`
	TrashDisposalPerDayOdorInfo1    float32 `json:"trash_disposal_per_day_odor_info_1"`
	TrashDisposalPerDayOdorInfo2    float32 `json:"trash_disposal_per_day_odor_info_2"`
	TrashDisposalPerDayOdorInfo3    float32 `json:"trash_disposal_per_day_odor_info_3"`
	TrashDisposalPerDayOdorInfo4    float32 `json:"trash_disposal_per_day_odor_info_4"`
}

func FsimsUserToResUser(fsimsUser *FSIMSUser) UserInfo {
	return UserInfo{
		UUID:    fsimsUser.UUID,
		Name:    fsimsUser.Name,
		Account: fsimsUser.Account,
		Type:    fsimsUser.Type,
		Role:    fsimsUser.Role,
		Status:  fsimsUser.Status,
		Company: fsimsUser.Company,
		Phone:   fsimsUser.Phone,
	}
}
