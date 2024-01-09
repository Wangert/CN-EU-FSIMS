package slaughter

import (
	"gorm.io/gorm"
	"time"
)

type SlaughterWaterQualityMon struct {
	//水质监控
	gorm.Model
	HouseNumber              string                   `gorm:"not null; type:varchar(256)" json:"house_number"`
	TimeRecordAt             time.Time                `json:"time_record_at"`                                                                          //记录时间
	SlaughterWaterMicroIndex SlaughterWaterMicroIndex `gorm:"foreignKey:SlaughterWaterQualityMonID; references:ID" json:"slaughter_water_micro_index"` //微生物指标
	OapGciSla                OapGciSla                `gorm:"foreignKey:SlaughterWaterQualityMonID; references:ID" json:"oap_gci_sla"`                 //感官性状和一般化学指标
	ToxinIndexSla            SlaughterToxinIndex      `gorm:"foreignKey:SlaughterWaterQualityMonID; references:ID" json:"toxin_index_sla"`             //屠宰污水毒理指标
}

type SlaughterWaterMicroIndex struct {
	//微生物监控
	gorm.Model
	SlaughterWaterQualityMonID *uint   `json:"slaughter_water_quality_mon_id"`
	SlaughterWaterMicroIndex1  float32 `json:"slaughter_water_micro_index_1"` //总大肠菌群
	SlaughterWaterMicroIndex2  float32 `json:"slaughter_water_micro_index_2"` //大肠埃希氏菌
	SlaughterWaterMicroIndex3  float32 `json:"slaughter_water_micro_index_3"` //菌落总数
}

type OapGciSla struct {
	//感官性状和一般化学指标
	gorm.Model
	SlaughterWaterQualityMonID *uint   `json:"slaughter_water_quality_mon_id"`
	OapGciSla1                 float32 `json:"oap_gci_sla_1"`  //钠
	OapGciSla2                 float32 `json:"oap_gci_sla_2"`  //挥发酚类
	OapGciSla3                 float32 `json:"oap_gci_sla_3"`  //阴离子合成洗涤剂
	OapGciSla4                 float32 `json:"oap_gci_sla_4"`  //2-甲基异莰醇
	OapGciSla5                 float32 `json:"oap_gci_sla_5"`  //土臭素
	OapGciSla6                 float32 `json:"oap_gci_sla_6"`  //色度（铂钴色度单位）/度
	OapGciSla7                 float32 `json:"oap_gci_sla_7"`  //浑浊度（散射浑浊度单位）/NTU
	OapGciSla8                 float32 `json:"oap_gci_sla_8"`  //臭和味
	OapGciSla9                 float32 `json:"oap_gci_sla_9"`  //肉眼可见物
	OapGciSla10                float32 `json:"oap_gci_sla_10"` //pH
	OapGciSla11                float32 `json:"oap_gci_sla_11"` //铝/（mg/L）
	OapGciSla12                float32 `json:"oap_gci_sla_12"` //铁/（mg/L）
	OapGciSla13                float32 `json:"oap_gci_sla_13"` //锰/（mg/L）
	OapGciSla14                float32 `json:"oap_gci_sla_14"` //铜/（mg/L）
	OapGciSla15                float32 `json:"oap_gci_sla_15"` //锌/（mg/L）
	OapGciSla16                float32 `json:"oap_gci_sla_16"` //氯化物/（mg/L）
	OapGciSla17                float32 `json:"oap_gci_sla_17"` //硫酸盐/（mg/L）
	OapGciSla18                float32 `json:"oap_gci_sla_18"` //溶解性总固体/（mg/L）
	OapGciSla19                float32 `json:"oap_gci_sla_19"` //总硬度（以CaCO3计）/（mg/L）
	OapGciSla20                float32 `json:"oap_gci_sla_20"` //高锰酸盐指数（以O2计）/（mg/L）
	OapGciSla21                float32 `json:"oap_gci_sla_21"` //氨（以N计）/（mg/L）
}

type MicroIndexWaterMonSla struct {
	//微生物指标屠宰阶段 和 屠宰污水检测记录
	gorm.Model
	SlaughterWaterQualityMonID *uint   `json:"slaughter_water_quality_mon_id"`
	MicroIndexWaterMonSla1     float32 `json:"micro_index_water_mon_sla_1"` //贾第鞭毛虫/（个/10L)
	MicroIndexWaterMonSla2     float32 `json:"micro_index_water_mon_sla_2"` //隐孢子虫/（个/10L）
	MicroIndexWaterMonSla3     float32 `json:"micro_index_water_mon_sla_3"` //生化需氧
	MicroIndexWaterMonSla4     float32 `json:"micro_index_water_mon_sla_4"` //化学需氧
}

type SlaughterToxinIndex struct {
	//屠宰阶段毒理性指标
	gorm.Model
	SlaughterWaterQualityMonID *uint                    `json:"slaughter_water_quality_mon_id"`
	ToxinIndexSla1             float32                  `json:"toxin_index_sla_1"`                                                                  //锑/（mg/L）
	ToxinIndexSla2             float32                  `json:"toxin_index_sla_2"`                                                                  //钡/（mg/L）
	ToxinIndexSla3             float32                  `json:"toxin_index_sla_3"`                                                                  //铍/（mg/L）
	ToxinIndexSla4             float32                  `json:"toxin_index_sla_4"`                                                                  //硼/（mg/L）
	ToxinIndexSla5             float32                  `json:"toxin_index_sla_5"`                                                                  //钼/（mg/L）
	ToxinIndexSla6             float32                  `json:"toxin_index_sla_6"`                                                                  //镍/（mg/L）
	ToxinIndexSla7             float32                  `json:"toxin_index_sla_7"`                                                                  //银/（mg/L）
	ToxinIndexSla8             float32                  `json:"toxin_index_sla_8"`                                                                  //铊/（mg/L）
	ToxinIndexSla9             float32                  `json:"toxin_index_sla_9"`                                                                  //硒/（mg/L）
	ToxinIndexSla10            float32                  `json:"toxin_index_sla_10"`                                                                 //高氯酸钾/（mg/L）
	ToxinIndexSla11            float32                  `json:"toxin_index_sla_11"`                                                                 //二氯甲烷/（mg/L）
	ToxinIndexSla12            float32                  `json:"toxin_index_sla_12"`                                                                 //1，2-二氯甲烷/（mg/L）
	ToxinIndexSla13            float32                  `json:"toxin_index_sla_13"`                                                                 //四氯化碳/（mg/L）
	ToxinIndexSla14            float32                  `json:"toxin_index_sla_14"`                                                                 //氯乙烯/（mg/L）
	ToxinIndexSla15            float32                  `json:"toxin_index_sla_15"`                                                                 //1，1-二氯乙烯/（mg/L）
	ToxinIndexSla16            float32                  `json:"toxin_index_sla_16"`                                                                 //1，2-二氯乙烯（总量）/（mg/L）
	ToxinIndexSla17            float32                  `json:"toxin_index_sla_17"`                                                                 //三氯乙烯/（mg/L）
	ToxinIndexSla18            float32                  `json:"toxin_index_sla_18"`                                                                 //四氯乙烯/（mg/L）
	ToxinIndexSla19            float32                  `json:"toxin_index_sla_19"`                                                                 //六氯丁二烯/（mg/L）
	ToxinIndexSla20            float32                  `json:"toxin_index_sla_20"`                                                                 //苯/（mg/L）
	ToxinIndexSla21            float32                  `json:"toxin_index_sla_21"`                                                                 //甲苯/（mg/L）
	ToxinIndexSla22            float32                  `json:"toxin_index_sla_22"`                                                                 //二甲苯（总量）/（mg/L）
	ToxinIndexSla23            float32                  `json:"toxin_index_sla_23"`                                                                 //苯乙烯/（mg/L）
	ToxinIndexSla24            float32                  `json:"toxin_index_sla_24"`                                                                 //氯苯/（mg/L）
	ToxinIndexSla25            float32                  `json:"toxin_index_sla_25"`                                                                 //1，4-二氯苯/（mg/L）
	ToxinIndexSla26            float32                  `json:"toxin_index_sla_26"`                                                                 //三氯苯（总量）/（mg/L）
	ToxinIndexSla27            float32                  `json:"toxin_index_sla_27"`                                                                 //六氯苯/（mg/L）
	ToxinIndexSla28            float32                  `json:"toxin_index_sla_28"`                                                                 //七氯/（mg/L）
	ToxinIndexSla29            float32                  `json:"toxin_index_sla_29"`                                                                 //马拉硫磷/（mg/L）
	ToxinIndexSla30            float32                  `json:"toxin_index_sla_30"`                                                                 //乐果/（mg/L）
	ToxinIndexSla31            float32                  `json:"toxin_index_sla_31"`                                                                 //灭草松/（mg/L）
	ToxinIndexSla32            float32                  `json:"toxin_index_sla_32"`                                                                 //百菌清/（mg/L）
	ToxinIndexSla33            float32                  `json:"toxin_index_sla_33"`                                                                 //呋喃丹/（mg/L）
	ToxinIndexSla34            float32                  `json:"toxin_index_sla_34"`                                                                 //毒死蜱/（mg/L）
	ToxinIndexSla35            float32                  `json:"toxin_index_sla_35"`                                                                 //草甘膦/（mg/L）
	ToxinIndexSla36            float32                  `json:"toxin_index_sla_36"`                                                                 //敌敌畏/（mg/L）
	ToxinIndexSla37            float32                  `json:"toxin_index_sla_37"`                                                                 //莠去津/（mg/L）
	ToxinIndexSla38            float32                  `json:"toxin_index_sla_38"`                                                                 //溴氰菊酯/（mg/L）
	ToxinIndexSla39            float32                  `json:"toxin_index_sla_39"`                                                                 //2，4-滴/（mg/L）
	ToxinIndexSla40            float32                  `json:"toxin_index_sla_40"`                                                                 //乙草胺/（mg/L）
	ToxinIndexSla41            float32                  `json:"toxin_index_sla_41"`                                                                 //五氯酚/（mg/L）
	ToxinIndexSla42            float32                  `json:"toxin_index_sla_42"`                                                                 //2，4，6-三氯酚/（mg/L）
	ToxinIndexSla43            float32                  `json:"toxin_index_sla_43"`                                                                 //苯并（a）芘/（mg/L）
	ToxinIndexSla44            float32                  `json:"toxin_index_sla_44"`                                                                 //邻苯二甲酸二（2-乙基己基）脂/（mg/L）
	ToxinIndexSla45            float32                  `json:"toxin_index_sla_45"`                                                                 //丙烯酰胺/（mg/L）
	ToxinIndexSla46            float32                  `json:"toxin_index_sla_46"`                                                                 //环氧氯丙烷/（mg/L）
	ToxinIndexSla47            float32                  `json:"toxin_index_sla_47"`                                                                 //微囊藻毒素-LR（藻类爆发情况发生时）/（mg/L）
	SlaughterWaterToxinIndex   SlaughterWaterToxinIndex `gorm:"foreignKey:SlaughterToxinIndexID; references:ID" json:"slaughter_water_toxin_index"` //其他毒理指标
}

type SlaughterWaterToxinIndex struct {
	//毒理性指标
	gorm.Model
	SlaughterToxinIndexID      *uint   `json:"slaughter_toxin_index_id"`
	SlaughterWaterToxinIndex1  float32 `json:"slaughter_water_toxin_index_1"`  //氰化物
	SlaughterWaterToxinIndex2  float32 `json:"slaughter_water_toxin_index_2"`  //总砷
	SlaughterWaterToxinIndex3  float32 `json:"slaughter_water_toxin_index_3"`  //总汞
	SlaughterWaterToxinIndex4  float32 `json:"slaughter_water_toxin_index_4"`  //总铅
	SlaughterWaterToxinIndex5  float32 `json:"slaughter_water_toxin_index_5"`  //铬（六价）
	SlaughterWaterToxinIndex6  float32 `json:"slaughter_water_toxin_index_6"`  //总镉
	SlaughterWaterToxinIndex7  float32 `json:"slaughter_water_toxin_index_7"`  //硝酸盐
	SlaughterWaterToxinIndex8  float32 `json:"slaughter_water_toxin_index_8"`  //砷
	SlaughterWaterToxinIndex9  float32 `json:"slaughter_water_toxin_index_9"`  //镉
	SlaughterWaterToxinIndex10 float32 `json:"slaughter_water_toxin_index_10"` //氯化物
	SlaughterWaterToxinIndex11 float32 `json:"slaughter_water_toxin_index_11"` //硝酸钾
	SlaughterWaterToxinIndex12 float32 `json:"slaughter_water_toxin_index_12"` //三氯甲烷
	SlaughterWaterToxinIndex13 float32 `json:"slaughter_water_toxin_index_13"` //一氯二溴甲烷
	SlaughterWaterToxinIndex14 float32 `json:"slaughter_water_toxin_index_14"` //二氯一溴甲烷
	SlaughterWaterToxinIndex15 float32 `json:"slaughter_water_toxin_index_15"` //三溴甲烷
	SlaughterWaterToxinIndex16 float32 `json:"slaughter_water_toxin_index_16"` //三卤甲烷（三氯甲烷、一氯二溴甲烷、二氯一溴甲烷、三溴甲烷的总和）
	SlaughterWaterToxinIndex17 float32 `json:"slaughter_water_toxin_index_17"` //二氯乙酸
	SlaughterWaterToxinIndex18 float32 `json:"slaughter_water_toxin_index_18"` //三氯乙酸
	SlaughterWaterToxinIndex19 float32 `json:"slaughter_water_toxin_index_19"` //溴酸盐
	SlaughterWaterToxinIndex20 float32 `json:"slaughter_water_toxin_index_20"` //亚氯酸盐
	SlaughterWaterToxinIndex21 float32 `json:"slaughter_water_toxin_index_21"` //氯酸盐
}

type SlaughterWaterQualityMonInfo struct {
	//水质监控
	HouseNumber              string                       `json:"house_number"`
	TimeRecordAt             string                       `json:"time_record_at"`              //记录时间
	SlaughterWaterMicroIndex SlaughterWaterMicroIndexInfo `json:"slaughter_water_micro_index"` //微生物指标
	OapGciSla                OapGciSlaInfo                `json:"oap_gci_sla"`                 //感官性状和一般化学指标
	ToxinIndexSla            SlaughterToxinIndexInfo      `json:"toxin_index_sla"`             //屠宰污水毒理指标
}

type SlaughterWaterMicroIndexInfo struct {
	//微生物监控
	SlaughterWaterMicroIndex1 float32 `json:"slaughter_water_micro_index_1"` //总大肠菌群
	SlaughterWaterMicroIndex2 float32 `json:"slaughter_water_micro_index_2"` //大肠埃希氏菌
	SlaughterWaterMicroIndex3 float32 `json:"slaughter_water_micro_index_3"` //菌落总数
}

type OapGciSlaInfo struct {
	//感官性状和一般化学指标
	OapGciSla1  float32 `json:"oap_gci_sla_1"`  //钠
	OapGciSla2  float32 `json:"oap_gci_sla_2"`  //挥发酚类
	OapGciSla3  float32 `json:"oap_gci_sla_3"`  //阴离子合成洗涤剂
	OapGciSla4  float32 `json:"oap_gci_sla_4"`  //2-甲基异莰醇
	OapGciSla5  float32 `json:"oap_gci_sla_5"`  //土臭素
	OapGciSla6  float32 `json:"oap_gci_sla_6"`  //色度（铂钴色度单位）/度
	OapGciSla7  float32 `json:"oap_gci_sla_7"`  //浑浊度（散射浑浊度单位）/NTU
	OapGciSla8  float32 `json:"oap_gci_sla_8"`  //臭和味
	OapGciSla9  float32 `json:"oap_gci_sla_9"`  //肉眼可见物
	OapGciSla10 float32 `json:"oap_gci_sla_10"` //pH
	OapGciSla11 float32 `json:"oap_gci_sla_11"` //铝/（mg/L）
	OapGciSla12 float32 `json:"oap_gci_sla_12"` //铁/（mg/L）
	OapGciSla13 float32 `json:"oap_gci_sla_13"` //锰/（mg/L）
	OapGciSla14 float32 `json:"oap_gci_sla_14"` //铜/（mg/L）
	OapGciSla15 float32 `json:"oap_gci_sla_15"` //锌/（mg/L）
	OapGciSla16 float32 `json:"oap_gci_sla_16"` //氯化物/（mg/L）
	OapGciSla17 float32 `json:"oap_gci_sla_17"` //硫酸盐/（mg/L）
	OapGciSla18 float32 `json:"oap_gci_sla_18"` //溶解性总固体/（mg/L）
	OapGciSla19 float32 `json:"oap_gci_sla_19"` //总硬度（以CaCO3计）/（mg/L）
	OapGciSla20 float32 `json:"oap_gci_sla_20"` //高锰酸盐指数（以O2计）/（mg/L）
	OapGciSla21 float32 `json:"oap_gci_sla_21"` //氨（以N计）/（mg/L）
}

type MicroIndexWaterMonSlaInfo struct {
	//微生物指标屠宰阶段 和 屠宰污水检测记录
	MicroIndexWaterMonSla1 float32 `json:"micro_index_water_mon_sla_1"` //贾第鞭毛虫/（个/10L)
	MicroIndexWaterMonSla2 float32 `json:"micro_index_water_mon_sla_2"` //隐孢子虫/（个/10L）
	MicroIndexWaterMonSla3 float32 `json:"micro_index_water_mon_sla_3"` //生化需氧
	MicroIndexWaterMonSla4 float32 `json:"micro_index_water_mon_sla_4"` //化学需氧
}

type SlaughterToxinIndexInfo struct {
	//屠宰阶段毒理性指标
	ToxinIndexSla1           float32                      `json:"toxin_index_sla_1"`                                                                  //锑/（mg/L）
	ToxinIndexSla2           float32                      `json:"toxin_index_sla_2"`                                                                  //钡/（mg/L）
	ToxinIndexSla3           float32                      `json:"toxin_index_sla_3"`                                                                  //铍/（mg/L）
	ToxinIndexSla4           float32                      `json:"toxin_index_sla_4"`                                                                  //硼/（mg/L）
	ToxinIndexSla5           float32                      `json:"toxin_index_sla_5"`                                                                  //钼/（mg/L）
	ToxinIndexSla6           float32                      `json:"toxin_index_sla_6"`                                                                  //镍/（mg/L）
	ToxinIndexSla7           float32                      `json:"toxin_index_sla_7"`                                                                  //银/（mg/L）
	ToxinIndexSla8           float32                      `json:"toxin_index_sla_8"`                                                                  //铊/（mg/L）
	ToxinIndexSla9           float32                      `json:"toxin_index_sla_9"`                                                                  //硒/（mg/L）
	ToxinIndexSla10          float32                      `json:"toxin_index_sla_10"`                                                                 //高氯酸钾/（mg/L）
	ToxinIndexSla11          float32                      `json:"toxin_index_sla_11"`                                                                 //二氯甲烷/（mg/L）
	ToxinIndexSla12          float32                      `json:"toxin_index_sla_12"`                                                                 //1，2-二氯甲烷/（mg/L）
	ToxinIndexSla13          float32                      `json:"toxin_index_sla_13"`                                                                 //四氯化碳/（mg/L）
	ToxinIndexSla14          float32                      `json:"toxin_index_sla_14"`                                                                 //氯乙烯/（mg/L）
	ToxinIndexSla15          float32                      `json:"toxin_index_sla_15"`                                                                 //1，1-二氯乙烯/（mg/L）
	ToxinIndexSla16          float32                      `json:"toxin_index_sla_16"`                                                                 //1，2-二氯乙烯（总量）/（mg/L）
	ToxinIndexSla17          float32                      `json:"toxin_index_sla_17"`                                                                 //三氯乙烯/（mg/L）
	ToxinIndexSla18          float32                      `json:"toxin_index_sla_18"`                                                                 //四氯乙烯/（mg/L）
	ToxinIndexSla19          float32                      `json:"toxin_index_sla_19"`                                                                 //六氯丁二烯/（mg/L）
	ToxinIndexSla20          float32                      `json:"toxin_index_sla_20"`                                                                 //苯/（mg/L）
	ToxinIndexSla21          float32                      `json:"toxin_index_sla_21"`                                                                 //甲苯/（mg/L）
	ToxinIndexSla22          float32                      `json:"toxin_index_sla_22"`                                                                 //二甲苯（总量）/（mg/L）
	ToxinIndexSla23          float32                      `json:"toxin_index_sla_23"`                                                                 //苯乙烯/（mg/L）
	ToxinIndexSla24          float32                      `json:"toxin_index_sla_24"`                                                                 //氯苯/（mg/L）
	ToxinIndexSla25          float32                      `json:"toxin_index_sla_25"`                                                                 //1，4-二氯苯/（mg/L）
	ToxinIndexSla26          float32                      `json:"toxin_index_sla_26"`                                                                 //三氯苯（总量）/（mg/L）
	ToxinIndexSla27          float32                      `json:"toxin_index_sla_27"`                                                                 //六氯苯/（mg/L）
	ToxinIndexSla28          float32                      `json:"toxin_index_sla_28"`                                                                 //七氯/（mg/L）
	ToxinIndexSla29          float32                      `json:"toxin_index_sla_29"`                                                                 //马拉硫磷/（mg/L）
	ToxinIndexSla30          float32                      `json:"toxin_index_sla_30"`                                                                 //乐果/（mg/L）
	ToxinIndexSla31          float32                      `json:"toxin_index_sla_31"`                                                                 //灭草松/（mg/L）
	ToxinIndexSla32          float32                      `json:"toxin_index_sla_32"`                                                                 //百菌清/（mg/L）
	ToxinIndexSla33          float32                      `json:"toxin_index_sla_33"`                                                                 //呋喃丹/（mg/L）
	ToxinIndexSla34          float32                      `json:"toxin_index_sla_34"`                                                                 //毒死蜱/（mg/L）
	ToxinIndexSla35          float32                      `json:"toxin_index_sla_35"`                                                                 //草甘膦/（mg/L）
	ToxinIndexSla36          float32                      `json:"toxin_index_sla_36"`                                                                 //敌敌畏/（mg/L）
	ToxinIndexSla37          float32                      `json:"toxin_index_sla_37"`                                                                 //莠去津/（mg/L）
	ToxinIndexSla38          float32                      `json:"toxin_index_sla_38"`                                                                 //溴氰菊酯/（mg/L）
	ToxinIndexSla39          float32                      `json:"toxin_index_sla_39"`                                                                 //2，4-滴/（mg/L）
	ToxinIndexSla40          float32                      `json:"toxin_index_sla_40"`                                                                 //乙草胺/（mg/L）
	ToxinIndexSla41          float32                      `json:"toxin_index_sla_41"`                                                                 //五氯酚/（mg/L）
	ToxinIndexSla42          float32                      `json:"toxin_index_sla_42"`                                                                 //2，4，6-三氯酚/（mg/L）
	ToxinIndexSla43          float32                      `json:"toxin_index_sla_43"`                                                                 //苯并（a）芘/（mg/L）
	ToxinIndexSla44          float32                      `json:"toxin_index_sla_44"`                                                                 //邻苯二甲酸二（2-乙基己基）脂/（mg/L）
	ToxinIndexSla45          float32                      `json:"toxin_index_sla_45"`                                                                 //丙烯酰胺/（mg/L）
	ToxinIndexSla46          float32                      `json:"toxin_index_sla_46"`                                                                 //环氧氯丙烷/（mg/L）
	ToxinIndexSla47          float32                      `json:"toxin_index_sla_47"`                                                                 //微囊藻毒素-LR（藻类爆发情况发生时）/（mg/L）
	SlaughterWaterToxinIndex SlaughterWaterToxinIndexInfo `gorm:"foreignKey:SlaughterToxinIndexID; references:ID" json:"slaughter_water_toxin_index"` //其他毒理指标
}

type SlaughterWaterToxinIndexInfo struct {
	//毒理性指标
	SlaughterWaterToxinIndex1  float32 `json:"slaughter_water_toxin_index_1"`  //氰化物
	SlaughterWaterToxinIndex2  float32 `json:"slaughter_water_toxin_index_2"`  //总砷
	SlaughterWaterToxinIndex3  float32 `json:"slaughter_water_toxin_index_3"`  //总汞
	SlaughterWaterToxinIndex4  float32 `json:"slaughter_water_toxin_index_4"`  //总铅
	SlaughterWaterToxinIndex5  float32 `json:"slaughter_water_toxin_index_5"`  //铬（六价）
	SlaughterWaterToxinIndex6  float32 `json:"slaughter_water_toxin_index_6"`  //总镉
	SlaughterWaterToxinIndex7  float32 `json:"slaughter_water_toxin_index_7"`  //硝酸盐
	SlaughterWaterToxinIndex8  float32 `json:"slaughter_water_toxin_index_8"`  //砷
	SlaughterWaterToxinIndex9  float32 `json:"slaughter_water_toxin_index_9"`  //镉
	SlaughterWaterToxinIndex10 float32 `json:"slaughter_water_toxin_index_10"` //氯化物
	SlaughterWaterToxinIndex11 float32 `json:"slaughter_water_toxin_index_11"` //硝酸钾
	SlaughterWaterToxinIndex12 float32 `json:"slaughter_water_toxin_index_12"` //三氯甲烷
	SlaughterWaterToxinIndex13 float32 `json:"slaughter_water_toxin_index_13"` //一氯二溴甲烷
	SlaughterWaterToxinIndex14 float32 `json:"slaughter_water_toxin_index_14"` //二氯一溴甲烷
	SlaughterWaterToxinIndex15 float32 `json:"slaughter_water_toxin_index_15"` //三溴甲烷
	SlaughterWaterToxinIndex16 float32 `json:"slaughter_water_toxin_index_16"` //三卤甲烷（三氯甲烷、一氯二溴甲烷、二氯一溴甲烷、三溴甲烷的总和）
	SlaughterWaterToxinIndex17 float32 `json:"slaughter_water_toxin_index_17"` //二氯乙酸
	SlaughterWaterToxinIndex18 float32 `json:"slaughter_water_toxin_index_18"` //三氯乙酸
	SlaughterWaterToxinIndex19 float32 `json:"slaughter_water_toxin_index_19"` //溴酸盐
	SlaughterWaterToxinIndex20 float32 `json:"slaughter_water_toxin_index_20"` //亚氯酸盐
	SlaughterWaterToxinIndex21 float32 `json:"slaughter_water_toxin_index_21"` //氯酸盐
}

func ToSlaughterWaterMicroIndexInfo(index *SlaughterWaterMicroIndex) SlaughterWaterMicroIndexInfo {
	return SlaughterWaterMicroIndexInfo{
		SlaughterWaterMicroIndex1: index.SlaughterWaterMicroIndex1,
		SlaughterWaterMicroIndex2: index.SlaughterWaterMicroIndex2,
		SlaughterWaterMicroIndex3: index.SlaughterWaterMicroIndex3,
	}
}

func ToOapGciSlaInfo(sla *OapGciSla) OapGciSlaInfo {
	return OapGciSlaInfo{
		OapGciSla1:  sla.OapGciSla1,
		OapGciSla2:  sla.OapGciSla2,
		OapGciSla3:  sla.OapGciSla3,
		OapGciSla4:  sla.OapGciSla4,
		OapGciSla5:  sla.OapGciSla5,
		OapGciSla6:  sla.OapGciSla6,
		OapGciSla7:  sla.OapGciSla7,
		OapGciSla8:  sla.OapGciSla8,
		OapGciSla9:  sla.OapGciSla9,
		OapGciSla10: sla.OapGciSla10,
		OapGciSla11: sla.OapGciSla11,
		OapGciSla12: sla.OapGciSla12,
		OapGciSla13: sla.OapGciSla13,
		OapGciSla14: sla.OapGciSla14,
		OapGciSla15: sla.OapGciSla15,
		OapGciSla16: sla.OapGciSla16,
		OapGciSla17: sla.OapGciSla17,
		OapGciSla18: sla.OapGciSla18,
		OapGciSla19: sla.OapGciSla19,
		OapGciSla20: sla.OapGciSla20,
		OapGciSla21: sla.OapGciSla21,
	}
}

func ToMicroIndexWaterMonSlaInfo(sla *MicroIndexWaterMonSla) MicroIndexWaterMonSlaInfo {
	return MicroIndexWaterMonSlaInfo{
		MicroIndexWaterMonSla1: sla.MicroIndexWaterMonSla1,
		MicroIndexWaterMonSla2: sla.MicroIndexWaterMonSla2,
		MicroIndexWaterMonSla3: sla.MicroIndexWaterMonSla3,
		MicroIndexWaterMonSla4: sla.MicroIndexWaterMonSla4,
	}
}

func ToSlaughterToxinIndexInfo(index *SlaughterToxinIndex) SlaughterToxinIndexInfo {
	return SlaughterToxinIndexInfo{
		ToxinIndexSla1:  index.ToxinIndexSla1,
		ToxinIndexSla2:  index.ToxinIndexSla2,
		ToxinIndexSla3:  index.ToxinIndexSla3,
		ToxinIndexSla4:  index.ToxinIndexSla4,
		ToxinIndexSla5:  index.ToxinIndexSla5,
		ToxinIndexSla6:  index.ToxinIndexSla6,
		ToxinIndexSla7:  index.ToxinIndexSla7,
		ToxinIndexSla8:  index.ToxinIndexSla8,
		ToxinIndexSla9:  index.ToxinIndexSla9,
		ToxinIndexSla10: index.ToxinIndexSla10,
		ToxinIndexSla11: index.ToxinIndexSla11,
		ToxinIndexSla12: index.ToxinIndexSla12,
		ToxinIndexSla13: index.ToxinIndexSla13,
		ToxinIndexSla14: index.ToxinIndexSla14,
		ToxinIndexSla15: index.ToxinIndexSla15,
		ToxinIndexSla16: index.ToxinIndexSla16,
		ToxinIndexSla17: index.ToxinIndexSla17,
		ToxinIndexSla18: index.ToxinIndexSla18,
		ToxinIndexSla19: index.ToxinIndexSla19,
		ToxinIndexSla20: index.ToxinIndexSla20,
		ToxinIndexSla21: index.ToxinIndexSla21,
		ToxinIndexSla22: index.ToxinIndexSla22,
		ToxinIndexSla23: index.ToxinIndexSla23,
		ToxinIndexSla24: index.ToxinIndexSla24,
		ToxinIndexSla25: index.ToxinIndexSla25,
		ToxinIndexSla26: index.ToxinIndexSla26,
		ToxinIndexSla27: index.ToxinIndexSla27,
		ToxinIndexSla28: index.ToxinIndexSla28,
		ToxinIndexSla29: index.ToxinIndexSla29,
		ToxinIndexSla30: index.ToxinIndexSla30,
		ToxinIndexSla31: index.ToxinIndexSla31,
		ToxinIndexSla32: index.ToxinIndexSla32,
		ToxinIndexSla33: index.ToxinIndexSla33,
		ToxinIndexSla34: index.ToxinIndexSla34,
		ToxinIndexSla35: index.ToxinIndexSla35,
		ToxinIndexSla36: index.ToxinIndexSla36,
		ToxinIndexSla37: index.ToxinIndexSla37,
		ToxinIndexSla38: index.ToxinIndexSla38,
		ToxinIndexSla39: index.ToxinIndexSla39,
		ToxinIndexSla40: index.ToxinIndexSla40,
		ToxinIndexSla41: index.ToxinIndexSla41,
		ToxinIndexSla42: index.ToxinIndexSla42,
		ToxinIndexSla43: index.ToxinIndexSla43,
		ToxinIndexSla44: index.ToxinIndexSla44,
		ToxinIndexSla45: index.ToxinIndexSla45,
		ToxinIndexSla46: index.ToxinIndexSla46,
		ToxinIndexSla47: index.ToxinIndexSla47,
		SlaughterWaterToxinIndex: SlaughterWaterToxinIndexInfo{
			SlaughterWaterToxinIndex1:  index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex1,
			SlaughterWaterToxinIndex2:  index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex2,
			SlaughterWaterToxinIndex3:  index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex3,
			SlaughterWaterToxinIndex4:  index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex4,
			SlaughterWaterToxinIndex5:  index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex5,
			SlaughterWaterToxinIndex6:  index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex6,
			SlaughterWaterToxinIndex7:  index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex7,
			SlaughterWaterToxinIndex8:  index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex8,
			SlaughterWaterToxinIndex9:  index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex9,
			SlaughterWaterToxinIndex10: index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex10,
			SlaughterWaterToxinIndex11: index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex11,
			SlaughterWaterToxinIndex12: index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex12,
			SlaughterWaterToxinIndex13: index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex13,
			SlaughterWaterToxinIndex14: index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex14,
			SlaughterWaterToxinIndex15: index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex15,
			SlaughterWaterToxinIndex16: index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex16,
			SlaughterWaterToxinIndex17: index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex17,
			SlaughterWaterToxinIndex18: index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex18,
			SlaughterWaterToxinIndex19: index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex19,
			SlaughterWaterToxinIndex20: index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex20,
			SlaughterWaterToxinIndex21: index.SlaughterWaterToxinIndex.SlaughterWaterToxinIndex21,
		},
	}
}

func ToSlaughterWaterQualityMonInfo(mon *SlaughterWaterQualityMon) SlaughterWaterQualityMonInfo {
	return SlaughterWaterQualityMonInfo{
		HouseNumber:              mon.HouseNumber,
		TimeRecordAt:             mon.TimeRecordAt.Format("2006-01-02 15:04:05"),
		SlaughterWaterMicroIndex: ToSlaughterWaterMicroIndexInfo(&mon.SlaughterWaterMicroIndex),
		OapGciSla:                ToOapGciSlaInfo(&mon.OapGciSla),
		ToxinIndexSla:            ToSlaughterToxinIndexInfo(&mon.ToxinIndexSla),
	}
}
