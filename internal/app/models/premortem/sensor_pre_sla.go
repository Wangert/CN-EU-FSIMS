package premortem

import "gorm.io/gorm"

type WaitingSlaughterCircle struct {
	gorm.Model
	WaitingSlaughterCircle1 WaitingSlaughterCircleGerms `gorm:"foreignKey:WaitingSlaughterCircleID; references:ID" json:"waiting_slaughter_circle_1"` //待宰圈致病菌
	WaitingSlaughterCircle2 float64                     `json:"waiting_slaughter_circle_2"`                                                           //待宰圈噪音
	WaitingSlaughterCircle3 float64                     `json:"waiting_slaughter_circle_3"`                                                           //待宰圈照度
	WaitingSlaughterCircle4 float64                     `json:"waiting_slaughter_circle_4"`                                                           //待宰圈光照时间
}

type Gps struct {
	gorm.Model
	PreSlaInfoRecID uint   `json:"pre_sla_info_rec_id"`
	Gps1            string `json:"gps_1"` //gps坐标
	Gps2            string `json:"gps_2"` //时间
}

type WaitingSlaughterCircleGerms struct {
	gorm.Model
	WaitingSlaughterCircleID     uint    `json:"waiting_slaughter_circle_id"`
	WaitingSlaughterCircleGerms1 float64 `json:"waiting_slaughter_circle_germs_1"` //沙门氏
	WaitingSlaughterCircleGerms2 float64 `json:"waiting_slaughter_circle_germs_2"` //STEC
	WaitingSlaughterCircleGerms3 float64 `json:"waiting_slaughter_circle_germs_3"` //单增
}

// 用于返回
type WaitingSlaughterCircleGermsInfo struct {
	WaitingSlaughterCircleGerms1 float64 `json:"waiting_slaughter_circle_germs_1"` //沙门氏
	WaitingSlaughterCircleGerms2 float64 `json:"waiting_slaughter_circle_germs_2"` //STEC
	WaitingSlaughterCircleGerms3 float64 `json:"waiting_slaughter_circle_germs_3"` //单增
}
