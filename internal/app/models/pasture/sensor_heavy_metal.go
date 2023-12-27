package pasture

import (
	"gorm.io/gorm"
	"time"
)

type HeavyMetal struct {
	//牧场饲料重金属
	gorm.Model
	TimeRecordAt time.Time `json:"time_record_at"` //时间记录
	HouseNumber  string    `gorm:"not null; type:varchar(256)" json:"house_number"`
	Ass          As        `gorm:"foreignKey:HeavyMetalID; references:ID" json:"aas"` //砷元素
	Pb           Pb        `gorm:"foreignKey:HeavyMetalID; references:ID" json:"pb"`  //铅元素
	Cd           Cd        `gorm:"foreignKey:HeavyMetalID; references:ID" json:"cd"`  //镉元素
	Cr           Cr        `gorm:"foreignKey:HeavyMetalID; references:ID" json:"cr"`  //铬元素
}

type As struct {
	gorm.Model
	HeavyMetalID uint    `json:"heavy_metal_id"`
	As1          float64 `json:"as_1"` //干草
	As2          float64 `json:"as_2"` //棕榈仁饼
	As3          float64 `json:"as_3"` //藻类及其加工产品
	As4          float64 `json:"as_4"` //甲壳类动物及其副产品（虾油除外）、鱼虾粉、水生软体动物及其副产品(油脂除外)
	As5          float64 `json:"as_5"` //其他矿物质
	As6          float64 `json:"as_6"` //添加剂
	As7          float64 `json:"as_7"` //浓缩
	As8          float64 `json:"as_8"` //精料
}

type Pb struct {
	gorm.Model
	HeavyMetalID uint    `json:"heavy_metal_id"`
	Pb1          float64 `json:"pb_1"` //单细胞蛋白饲料原料
	Pb2          float64 `json:"pb_2"` //矿物质饲料原料
	Pb3          float64 `json:"pb_3"` //饲草、粗饲料及其加工产品
	Pb4          float64 `json:"pb_4"` //浓缩饲料
	Pb5          float64 `json:"pb_5"` //精料补充料
	Pb6          float64 `json:"pb_6"` //配合饲料
	Pb7          float64 `json:"pb_7"` //其他饲料原料
}

type Cd struct {
	gorm.Model
	HeavyMetalID uint    `json:"heavy_metal_id"`
	Cd1          float64 `json:"cd_1"` //藻类及其加工产品
	Cd2          float64 `json:"cd_2"` //其他动物源性饲料原料
	Cd3          float64 `json:"cd_3"` //其他矿物质饲料原料
	Cd4          float64 `json:"cd_4"` //添加剂预混合饲料
	Cd5          float64 `json:"cd_5"` //浓缩饲料
	Cd6          float64 `json:"cd_6"` //犊牛精料补充料
	Cd7          float64 `json:"cd_7"` //其他精料补充料
}

type Cr struct {
	gorm.Model
	HeavyMetalID uint    `json:"heavy_metal_id"`
	Cr1          float64 `json:"cr_1"` //其他添加剂预混合饲料
	Cr2          float64 `json:"cr_2"` //配合饲料
	Cr3          float64 `json:"cr_3"` //其他浓缩饲料
	Cr4          float64 `json:"cr_4"` //配合饲料
}

type HeavyMetalInfo struct {
	//牧场饲料重金属
	TimeRecordAt time.Time `json:"time_record_at"` //时间记录
	HouseNumber  string    `json:"house_number"`
	Ass          AsInfo    `json:"aas"` //砷元素
	Pb           PbInfo    `json:"pb"`  //铅元素
	Cd           CdInfo    `json:"cd"`  //镉元素
	Cr           CrInfo    `json:"cr"`  //铬元素
}

type AsInfo struct {
	As1 float64 `json:"as_1"` //干草
	As2 float64 `json:"as_2"` //棕榈仁饼
	As3 float64 `json:"as_3"` //藻类及其加工产品
	As4 float64 `json:"as_4"` //甲壳类动物及其副产品（虾油除外）、鱼虾粉、水生软体动物及其副产品(油脂除外)
	As5 float64 `json:"as_5"` //其他矿物质
	As6 float64 `json:"as_6"` //添加剂
	As7 float64 `json:"as_7"` //浓缩
	As8 float64 `json:"as_8"` //精料
}

type PbInfo struct {
	Pb1 float64 `json:"pb_1"` //单细胞蛋白饲料原料
	Pb2 float64 `json:"pb_2"` //矿物质饲料原料
	Pb3 float64 `json:"pb_3"` //饲草、粗饲料及其加工产品
	Pb4 float64 `json:"pb_4"` //浓缩饲料
	Pb5 float64 `json:"pb_5"` //精料补充料
	Pb6 float64 `json:"pb_6"` //配合饲料
	Pb7 float64 `json:"pb_7"` //其他饲料原料
}

type CdInfo struct {
	Cd1 float64 `json:"cd_1"` //藻类及其加工产品
	Cd2 float64 `json:"cd_2"` //其他动物源性饲料原料
	Cd3 float64 `json:"cd_3"` //其他矿物质饲料原料
	Cd4 float64 `json:"cd_4"` //添加剂预混合饲料
	Cd5 float64 `json:"cd_5"` //浓缩饲料
	Cd6 float64 `json:"cd_6"` //犊牛精料补充料
	Cd7 float64 `json:"cd_7"` //其他精料补充料
}

type CrInfo struct {
	HeavyMetalID uint    `json:"heavy_metal_id"`
	Cr1          float64 `json:"cr_1"` //其他添加剂预混合饲料
	Cr2          float64 `json:"cr_2"` //配合饲料
	Cr3          float64 `json:"cr_3"` //其他浓缩饲料
	Cr4          float64 `json:"cr_4"` //配合饲料
}
