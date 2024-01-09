package pasture

import (
	"gorm.io/gorm"
	"time"
)

type PastureFeedHeavyMetal struct {
	//牧场饲料重金属
	gorm.Model
	TimeRecordAt  time.Time     `json:"time_record_at"` //时间记录
	HouseNumber   string        `gorm:"not null; type:varchar(256)" json:"house_number"`
	PastureFeedAs PastureFeedAs `gorm:"foreignKey:PastureFeedHeavyMetalID; references:ID" json:"pasture_feed_as"` //砷元素
	PastureFeedPb PastureFeedPb `gorm:"foreignKey:PastureFeedHeavyMetalID; references:ID" json:"pasture_feed_pb"` //铅元素
	PastureFeedCd PastureFeedCd `gorm:"foreignKey:PastureFeedHeavyMetalID; references:ID" json:"pasture_feed_cd"` //镉元素
	PastureFeedCr PastureFeedCr `gorm:"foreignKey:PastureFeedHeavyMetalID; references:ID" json:"pasture_feed_cr"` //铬元素
}

type PastureFeedAs struct {
	gorm.Model
	PastureFeedHeavyMetalID *uint   `json:"pasture_feed_heavy_metal_id"`
	As1                     float64 `json:"as_1"` //干草
	As2                     float64 `json:"as_2"` //棕榈仁饼
	As3                     float64 `json:"as_3"` //藻类及其加工产品
	As4                     float64 `json:"as_4"` //甲壳类动物及其副产品（虾油除外）、鱼虾粉、水生软体动物及其副产品(油脂除外)
	As5                     float64 `json:"as_5"` //其他矿物质
	As6                     float64 `json:"as_6"` //添加剂
	As7                     float64 `json:"as_7"` //浓缩
	As8                     float64 `json:"as_8"` //精料
}

type PastureFeedPb struct {
	gorm.Model
	PastureFeedHeavyMetalID *uint   `json:"pasture_feed_heavy_metal_id"`
	Pb1                     float64 `json:"pb_1"` //单细胞蛋白饲料原料
	Pb2                     float64 `json:"pb_2"` //矿物质饲料原料
	Pb3                     float64 `json:"pb_3"` //饲草、粗饲料及其加工产品
	Pb4                     float64 `json:"pb_4"` //浓缩饲料
	Pb5                     float64 `json:"pb_5"` //精料补充料
	Pb6                     float64 `json:"pb_6"` //配合饲料
	Pb7                     float64 `json:"pb_7"` //其他饲料原料
}

type PastureFeedCd struct {
	gorm.Model
	PastureFeedHeavyMetalID *uint   `json:"pasture_feed_heavy_metal_id"`
	Cd1                     float64 `json:"cd_1"` //藻类及其加工产品
	Cd2                     float64 `json:"cd_2"` //其他动物源性饲料原料
	Cd3                     float64 `json:"cd_3"` //其他矿物质饲料原料
	Cd4                     float64 `json:"cd_4"` //添加剂预混合饲料
	Cd5                     float64 `json:"cd_5"` //浓缩饲料
	Cd6                     float64 `json:"cd_6"` //犊牛精料补充料
	Cd7                     float64 `json:"cd_7"` //其他精料补充料
}

type PastureFeedCr struct {
	gorm.Model
	PastureFeedHeavyMetalID *uint   `json:"pasture_feed_heavy_metal_id"`
	Cr1                     float64 `json:"cr_1"` //其他添加剂预混合饲料
	Cr2                     float64 `json:"cr_2"` //配合饲料
	Cr3                     float64 `json:"cr_3"` //其他浓缩饲料
	Cr4                     float64 `json:"cr_4"` //配合饲料
}

type PastureFeedHeavyMetalInfo struct {
	//牧场饲料重金属
	TimeRecordAt      string            `json:"time_record_at"` //时间记录
	HouseNumber       string            `json:"house_number"`
	PastureFeedAsInfo PastureFeedAsInfo `json:"pasture_feed_as_info"` //砷元素
	PastureFeedPbInfo PastureFeedPbInfo `json:"pasture_feed_pb_info"` //铅元素
	PastureFeedCdInfo PastureFeedCdInfo `json:"pasture_feed_cd_info"` //镉元素
	PastureFeedCrInfo PastureFeedCrInfo `json:"pasture_feed_cr_info"` //铬元素
}

func ToPastureFeedHeavyMetalInfo(hm *PastureFeedHeavyMetal) PastureFeedHeavyMetalInfo {
	return PastureFeedHeavyMetalInfo{
		TimeRecordAt:      hm.TimeRecordAt.Format("2006-01-02 15:04:05"),
		HouseNumber:       hm.HouseNumber,
		PastureFeedAsInfo: ToPastureFeedAsInfo(&hm.PastureFeedAs),
		PastureFeedPbInfo: ToPastureFeedPbInfo(&hm.PastureFeedPb),
		PastureFeedCdInfo: ToPastureFeedCdInfo(&hm.PastureFeedCd),
		PastureFeedCrInfo: ToPastureFeedCrInfo(&hm.PastureFeedCr),
	}
}

type PastureFeedAsInfo struct {
	As1 float64 `json:"as_1"` //干草
	As2 float64 `json:"as_2"` //棕榈仁饼
	As3 float64 `json:"as_3"` //藻类及其加工产品
	As4 float64 `json:"as_4"` //甲壳类动物及其副产品（虾油除外）、鱼虾粉、水生软体动物及其副产品(油脂除外)
	As5 float64 `json:"as_5"` //其他矿物质
	As6 float64 `json:"as_6"` //添加剂
	As7 float64 `json:"as_7"` //浓缩
	As8 float64 `json:"as_8"` //精料
}

func ToPastureFeedAsInfo(as *PastureFeedAs) PastureFeedAsInfo {
	return PastureFeedAsInfo{
		As1: as.As1,
		As2: as.As2,
		As3: as.As3,
		As4: as.As4,
		As5: as.As5,
		As6: as.As6,
		As7: as.As7,
		As8: as.As8,
	}
}

type PastureFeedPbInfo struct {
	Pb1 float64 `json:"pb_1"` //单细胞蛋白饲料原料
	Pb2 float64 `json:"pb_2"` //矿物质饲料原料
	Pb3 float64 `json:"pb_3"` //饲草、粗饲料及其加工产品
	Pb4 float64 `json:"pb_4"` //浓缩饲料
	Pb5 float64 `json:"pb_5"` //精料补充料
	Pb6 float64 `json:"pb_6"` //配合饲料
	Pb7 float64 `json:"pb_7"` //其他饲料原料
}

func ToPastureFeedPbInfo(pb *PastureFeedPb) PastureFeedPbInfo {
	return PastureFeedPbInfo{
		Pb1: pb.Pb1,
		Pb2: pb.Pb2,
		Pb3: pb.Pb3,
		Pb4: pb.Pb4,
		Pb5: pb.Pb5,
		Pb6: pb.Pb6,
		Pb7: pb.Pb7,
	}
}

type PastureFeedCdInfo struct {
	Cd1 float64 `json:"cd_1"` //藻类及其加工产品
	Cd2 float64 `json:"cd_2"` //其他动物源性饲料原料
	Cd3 float64 `json:"cd_3"` //其他矿物质饲料原料
	Cd4 float64 `json:"cd_4"` //添加剂预混合饲料
	Cd5 float64 `json:"cd_5"` //浓缩饲料
	Cd6 float64 `json:"cd_6"` //犊牛精料补充料
	Cd7 float64 `json:"cd_7"` //其他精料补充料
}

func ToPastureFeedCdInfo(fcd *PastureFeedCd) PastureFeedCdInfo {
	return PastureFeedCdInfo{
		Cd1: fcd.Cd1,
		Cd2: fcd.Cd2,
		Cd3: fcd.Cd3,
		Cd4: fcd.Cd4,
		Cd5: fcd.Cd5,
		Cd6: fcd.Cd6,
		Cd7: fcd.Cd7,
	}
}

type PastureFeedCrInfo struct {
	Cr1 float64 `json:"cr_1"` //其他添加剂预混合饲料
	Cr2 float64 `json:"cr_2"` //配合饲料
	Cr3 float64 `json:"cr_3"` //其他浓缩饲料
	Cr4 float64 `json:"cr_4"` //配合饲料
}

func ToPastureFeedCrInfo(fc *PastureFeedCr) PastureFeedCrInfo {
	return PastureFeedCrInfo{
		Cr1: fc.Cr1,
		Cr2: fc.Cr2,
		Cr3: fc.Cr3,
		Cr4: fc.Cr4,
	}
}
