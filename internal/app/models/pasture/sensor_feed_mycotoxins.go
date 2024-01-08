package pasture

import (
	"gorm.io/gorm"
	"time"
)

type PastureFeedMycotoxins struct {
	////饲料中真菌毒素、农兽药残留
	gorm.Model
	TimeRecordAt time.Time `json:"time_record_at"` //数据记录时间
	HouseNumber  string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	Afb1         Afb1      `gorm:"foreignKey:PastureFeedMycotoxinsID; references:ID" json:"afb_1"`       //黄曲霉毒素B1
	Don          Don       `gorm:"foreignKey:PastureFeedMycotoxinsID; references:ID" json:"don"`         //玉米赤霉烯酮
	T2toxin      T2toxin   `gorm:"foreignKey:PastureFeedMycotoxinsID; references:ID" json:"t2Toxin"`     //脱氧雪腐镰刀菌烯醇（呕吐毒素）
	T2VomZea     T2VomZea  `gorm:"foreignKey:PastureFeedMycotoxinsID; references:ID" json:"t_2_vom_zea"` //T-2毒素 伏马毒素 赭曲霉毒素A
}

type Afb1 struct {
	gorm.Model
	PastureFeedMycotoxinsID *uint   `json:"pasture_feed_mycotoxins_id"`
	Afb11                   float64 `json:"afb_11"` //玉米加工产品、花生饼(粕)
	Afb12                   float64 `json:"afb_12"` //植物油脂(玉米油花生油除外)
	Afb13                   float64 `json:"afb_13"` //玉米油 、花生油
	Afb14                   float64 `json:"afb_14"` //其他植物性饲料原料
	Afb15                   float64 `json:"afb_15"` //其他浓缩饲料
	Afb16                   float64 `json:"afb_16"` //犊牛、羔羊精料补充料
	Afb17                   float64 `json:"afb_17"` //其他精料补充料
}

type Don struct {
	gorm.Model
	PastureFeedMycotoxinsID *uint   `json:"pasture_feed_mycotoxins_id"`
	Don1                    float64 `json:"don_1"` //玉米及其加工产品 (玉米皮、喷浆玉米皮、玉米浆干粉除外)
	Don2                    float64 `json:"don_2"` //玉米皮、喷浆玉米皮.玉米浆干粉、玉米酒糟
	Don3                    float64 `json:"don_3"` //其他植物性饲料原料
	Don4                    float64 `json:"don_4"` //其他配合饲料
}

type T2toxin struct {
	gorm.Model
	PastureFeedMycotoxinsID *uint   `json:"pasture_feed_mycotoxins_id"`
	T2toxin1                float64 `json:"t_2_toxin_1"` //植物性饲料原料
	T2toxin2                float64 `json:"t_2_toxin_2"` //其他精料补充料
	T2toxin3                float64 `json:"t_2_toxin_3"` //其他配合饲料
}

type T2VomZea struct {
	gorm.Model
	PastureFeedMycotoxinsID *uint   `json:"pasture_feed_mycotoxins_id"`
	T2AVomZea1              float64 `json:"t_2_a_vom_zea_1"` //Zea 配合饲料
	T2AVomZea2              float64 `json:"t_2_a_vom_zea_2"` //Vom 其他反刍动物精料补充料
	T2AVomZea3              float64 `json:"t_2_a_vom_zea_3"` //T2	植物性饲料原料
}

type PastureFeedMycotoxinsInfo struct {
	////饲料中真菌毒素、农兽药残留
	TimeRecordAt string       `json:"time_record_at"` //数据记录时间
	HouseNumber  string       `json:"house_number"`
	Afb1         Afb1Info     `json:"afb_1"`       //黄曲霉毒素B1
	Don          DonInfo      `json:"don"`         //玉米赤霉烯酮
	T2toxin      T2toxinInfo  `json:"t2Toxin"`     //脱氧雪腐镰刀菌烯醇（呕吐毒素）
	T2VomZea     T2VomZeaInfo `json:"t_2_vom_zea"` //T-2毒素 伏马毒素 赭曲霉毒素A
}

func ToPastureFeedMycotoxinsInfo(mc *PastureFeedMycotoxins) PastureFeedMycotoxinsInfo {
	return PastureFeedMycotoxinsInfo{
		TimeRecordAt: mc.TimeRecordAt.Format("2006-01-02 15:04:05"),
		HouseNumber:  mc.HouseNumber,
		Afb1:         ToAfb1Info(&mc.Afb1),
		Don:          ToDonInfo(&mc.Don),
		T2toxin:      ToT2toxinInfo(&mc.T2toxin),
		T2VomZea:     ToT2VomZeaInfo(&mc.T2VomZea),
	}
}

type Afb1Info struct {
	Afb11 float64 `json:"afb_11"` //玉米加工产品、花生饼(粕)
	Afb12 float64 `json:"afb_12"` //植物油脂(玉米油花生油除外)
	Afb13 float64 `json:"afb_13"` //玉米油 、花生油
	Afb14 float64 `json:"afb_14"` //其他植物性饲料原料
	Afb15 float64 `json:"afb_15"` //其他浓缩饲料
	Afb16 float64 `json:"afb_16"` //犊牛、羔羊精料补充料
	Afb17 float64 `json:"afb_17"` //其他精料补充料
}

func ToAfb1Info(af *Afb1) Afb1Info {
	return Afb1Info{
		Afb11: af.Afb11,
		Afb12: af.Afb12,
		Afb13: af.Afb13,
		Afb14: af.Afb14,
		Afb15: af.Afb15,
		Afb16: af.Afb16,
		Afb17: af.Afb17,
	}
}

type DonInfo struct {
	Don1 float64 `json:"don_1"` //玉米及其加工产品 (玉米皮、喷浆玉米皮、玉米浆干粉除外)
	Don2 float64 `json:"don_2"` //玉米皮、喷浆玉米皮.玉米浆干粉、玉米酒糟
	Don3 float64 `json:"don_3"` //其他植物性饲料原料
	Don4 float64 `json:"don_4"` //其他配合饲料
}

func ToDonInfo(d *Don) DonInfo {
	return DonInfo{
		Don1: d.Don1,
		Don2: d.Don2,
		Don3: d.Don3,
		Don4: d.Don4,
	}
}

type T2toxinInfo struct {
	T2toxin1 float64 `json:"t_2_toxin_1"` //植物性饲料原料
	T2toxin2 float64 `json:"t_2_toxin_2"` //其他精料补充料
	T2toxin3 float64 `json:"t_2_toxin_3"` //其他配合饲料
}

func ToT2toxinInfo(tti *T2toxin) T2toxinInfo {
	return T2toxinInfo{
		T2toxin1: tti.T2toxin1,
		T2toxin2: tti.T2toxin3,
		T2toxin3: tti.T2toxin3,
	}
}

type T2VomZeaInfo struct {
	T2AVomZea1 float64 `json:"t_2_a_vom_zea_1"` //Zea 配合饲料
	T2AVomZea2 float64 `json:"t_2_a_vom_zea_2"` //Vom 其他反刍动物精料补充料
	T2AVomZea3 float64 `json:"t_2_a_vom_zea_3"` //T2	植物性饲料原料
}

func ToT2VomZeaInfo(tzi *T2VomZea) T2VomZeaInfo {
	return T2VomZeaInfo{
		T2AVomZea1: tzi.T2AVomZea1,
		T2AVomZea2: tzi.T2AVomZea2,
		T2AVomZea3: tzi.T2AVomZea3,
	}
}