package response

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/pasture"
	"CN-EU-FSIMS/internal/app/models/slaughter"
)

type ResUsers struct {
	Users []models.UserInfo `json:"users"`
	Count int64             `json:"count"`
}

type ResLogin struct {
	UUID        string `json:"uuid"`
	Token       string `json:"token"`
	UserType    int    `json:"user_type"`
	HouseNumber string `json:"house_number"`
}

type ResSlaAndPasTrashPeyDay struct {
	OdorSlaAndPasTrashPeyDay1    float64 `json:"odor_sla_and_pas_trash_pey_day_1"`
	OdorSlaAndPasTrashPeyDay2    float64 `json:"odor_sla_and_pas_trash_pey_day_2"`
	OdorSlaAndPasTrashPeyDay3    float64 `json:"odor_sla_and_pas_trash_pey_day_3"`
	OdorSlaAndPasTrashPeyDay4    float64 `json:"odor_sla_and_pas_trash_pey_day_4"`
	ResidueSlaAndPasTrashPeyDay1 float64 `json:"residue_sla_and_pas_trash_pey_day_1"`
	ResidueSlaAndPasTrashPeyDay2 float64 `json:"residue_sla_and_pas_trash_pey_day_2"`
	ResidueSlaAndPasTrashPeyDay3 float64 `json:"residue_sla_and_pas_trash_pey_day_3"`
	ResidueSlaAndPasTrashPeyDay4 float64 `json:"residue_sla_and_pas_trash_pey_day_4"`
	WaterSlaAndPasTrashPeyDay1   float64 `json:"water_sla_and_pas_trash_pey_day_1"`
	WaterSlaAndPasTrashPeyDay2   float64 `json:"water_sla_and_pas_trash_pey_day_2"`
	WaterSlaAndPasTrashPeyDay3   float64 `json:"water_sla_and_pas_trash_pey_day_3"`
	WaterSlaAndPasTrashPeyDay4   float64 `json:"water_sla_and_pas_trash_pey_day_4"`
}

type ResPastureTrashDisposalFifteenDays struct {
	Count                  int64                                  `json:"count"`
	PastureDisposalRecords []pasture.AllPasturesTrashDisposalInfo `json:"pasture_disposal_records"`
}

type ResSlaughterDisposalFifteenDays struct {
	Count                    int64                                      `json:"count"`
	SlaughterDisposalRecords []slaughter.AllSlaughtersTrashDisposalInfo `json:"slaughter_disposal_records"`
}
