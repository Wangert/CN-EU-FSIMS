package pasture

import "gorm.io/gorm"

type Cass struct {
	////饲料中真菌毒素、农兽药残留
	gorm.Model
	TimeRecordAt string   `json:"time_record_at"`                                      //数据记录时间
	Afb1         Afb1     `gorm:"foreignKey:CassID; references:ID" json:"afb_1"`       //黄曲霉毒素B1
	Don          Don      `gorm:"foreignKey:CassID; references:ID" json:"don"`         //玉米赤霉烯酮
	T2toxin      T2toxin  `gorm:"foreignKey:CassID; references:ID" json:"t2Toxin"`     //脱氧雪腐镰刀菌烯醇（呕吐毒素）
	T2VomZea     T2VomZea `gorm:"foreignKey:CassID; references:ID" json:"t_2_vom_zea"` //T-2毒素 伏马毒素 赭曲霉毒素A
}

type Afb1 struct {
	gorm.Model
	CassID uint    `json:"cass_id"`
	Afb11  float64 `json:"afb_11"` //玉米加工产品、花生饼(粕)
	Afb12  float64 `json:"afb_12"` //植物油脂(玉米油花生油除外)
	Afb13  float64 `json:"afb_13"` //玉米油 、花生油
	Afb14  float64 `json:"afb_14"` //其他植物性饲料原料
	Afb15  float64 `json:"afb_15"` //其他浓缩饲料
	Afb16  float64 `json:"afb_16"` //犊牛、羔羊精料补充料
	Afb17  float64 `json:"afb_17"` //其他精料补充料
}

type Don struct {
	gorm.Model
	CassID uint    `json:"cass_id"`
	Don1   float64 `json:"don_1"` //玉米及其加工产品 (玉米皮、喷浆玉米皮、玉米浆干粉除外)
	Don2   float64 `json:"don_2"` //玉米皮、喷浆玉米皮.玉米浆干粉、玉米酒糟
	Don3   float64 `json:"don_3"` //其他植物性饲料原料
	Don4   float64 `json:"don_4"` //其他配合饲料
}

type T2toxin struct {
	gorm.Model
	CassID   uint    `json:"cass_id"`
	T2toxin1 float64 `json:"t_2_toxin_1"` //植物性饲料原料
	T2toxin2 float64 `json:"t_2_toxin_2"` //其他精料补充料
	T2toxin3 float64 `json:"t_2_toxin_3"` //其他配合饲料
}

type T2VomZea struct {
	gorm.Model
	CassID     uint    `json:"cass_id"`
	T2AVomZea1 float64 `json:"t_2_a_vom_zea_1"` //Zea 配合饲料
	T2AVomZea2 float64 `json:"t_2_a_vom_zea_2"` //Vom 其他反刍动物精料补充料
	T2AVomZea3 float64 `json:"t_2_a_vom_zea_3"` //T2	植物性饲料原料
}
