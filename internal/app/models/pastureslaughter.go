package models

import (
	"gorm.io/gorm"
	"time"
)

type PastureSlaughterTrashDisposal struct {
	gorm.Model
	TimeStamp                             time.Time `json:"time_stamp"`
	WaterPastureSlaughterTrashDisposal1   float32   `json:"water_pasture_slaughter_trash_disposal_1"`
	WaterPastureSlaughterTrashDisposal2   float32   `json:"water_pasture_slaughter_trash_disposal_2"`
	WaterPastureSlaughterTrashDisposal3   float32   `json:"water_pasture_slaughter_trash_disposal_3"`
	WaterPastureSlaughterTrashDisposal4   float32   `json:"water_pasture_slaughter_trash_disposal_4"`
	ResiduePastureSlaughterTrashDisposal1 float32   `json:"residue_pasture_slaughter_trash_disposal_1"`
	ResiduePastureSlaughterTrashDisposal2 float32   `json:"residue_pasture_slaughter_trash_disposal_2"`
	ResiduePastureSlaughterTrashDisposal3 float32   `json:"residue_pasture_slaughter_trash_disposal_3"`
	ResiduePastureSlaughterTrashDisposal4 float32   `json:"residue_pasture_slaughter_trash_disposal_4"`
	OdorPastureSlaughterTrashDisposal1    float32   `json:"odor_pasture_slaughter_trash_disposal_1"`
	OdorPastureSlaughterTrashDisposal2    float32   `json:"odor_pasture_slaughter_trash_disposal_2"`
	OdorPastureSlaughterTrashDisposal3    float32   `json:"odor_pasture_slaughter_trash_disposal_3"`
	OdorPastureSlaughterTrashDisposal4    float32   `json:"odor_pasture_slaughter_trash_disposal_4"`
}
