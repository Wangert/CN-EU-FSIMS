package slaughter

import "gorm.io/gorm"

type SlaughterDisinfectHotWaterTempMoni struct {
	//消毒热水温度监控
	gorm.Model
	SlaughterDisinfectHotWaterTempMoni1 float32 `json:"slaughter_disinfect_hot_water_temp_moni_1"` //预剥皮
	SlaughterDisinfectHotWaterTempMoni2 float32 `json:"slaughter_disinfect_hot_water_temp_moni_2"` //检疫台
	SlaughterDisinfectHotWaterTempMoni3 float32 `json:"slaughter_disinfect_hot_water_temp_moni_3"` //去皮后修正
	SlaughterDisinfectHotWaterTempMoni4 float32 `json:"slaughter_disinfect_hot_water_temp_moni_4"` //去脏后
	SlaughterDisinfectHotWaterTempMoni5 float32 `json:"slaughter_disinfect_hot_water_temp_moni_5"` //剔骨台
	SlaughterDisinfectHotWaterTempMoni6 float32 `json:"slaughter_disinfect_hot_water_temp_moni_6"` //洗手台
}

type PreCoolShop struct {
	//预冷间
	gorm.Model
	PreCoolShop1 float32 `json:"pre_cool_shop_1"` //预冷间温度
	PreCoolShop2 float32 `json:"pre_cool_shop_2"` //预冷间湿度
	PreCoolShop3 float32 `json:"pre_cool_shop_3"` //副产物温度
}

type SlaShop struct {
	//屠宰车间
	gorm.Model
	SlaShop1  float32         `json:"sla_shop_1"`  //紫外灭菌
	SlaShop2  float32         `json:"sla_shop_2"`  //臭氧
	SlaShop3  float32         `json:"sla_shop_3"`  //臭氧残留
	SlaShop4  float32         `json:"sla_shop_4"`  //湿度
	SlaShop5  float32         `json:"sla_shop_5"`  //温度
	SlaShop6  float32         `json:"sla_shop_6"`  //预冷间消毒记录
	SlaShop7  float32         `json:"sla_shop_7"`  //氯含量
	SlaShop8  DisinfectRecord `json:"sla_shop_8"`  //消毒记录
	SlaShop9  float32         `json:"sla_shop_9"`  //工作服 功率
	SlaShop10 float32         `json:"sla_shop_10"` //工作服 时间
}

type DivShop struct {
	//分割车间
	gorm.Model
	DivShop1  float32         `json:"div_shop_1"`  //紫外灭菌
	DivShop2  float32         `json:"div_shop_2"`  //臭氧
	DivShop3  float32         `json:"div_shop_3"`  //臭氧残留
	DivShop4  float32         `json:"div_shop_4"`  //湿度
	DivShop5  float32         `json:"div_shop_5"`  //温度
	DivShop6  float32         `json:"div_shop_6"`  //预冷间消毒记录
	DivShop7  float32         `json:"div_shop_7"`  //氯含量
	DivShop8  DisinfectRecord `json:"div_shop_8"`  //消毒记录
	DivShop9  float32         `json:"div_shop_9"`  //工作服 功率
	DivShop10 float32         `json:"div_shop_10"` //工作服 时间
}

type AcidShop struct {
	//排酸车间
	gorm.Model
	AcidShop1  float32         `json:"acid_shop_1"`  //紫外灭菌
	AcidShop2  float32         `json:"acid_shop_2"`  //臭氧
	AcidShop3  float32         `json:"acid_shop_3"`  //臭氧残留
	AcidShop4  float32         `json:"acid_shop_4"`  //湿度
	AcidShop5  float32         `json:"acid_shop_5"`  //温度
	AcidShop6  float32         `json:"acid_shop_6"`  //预冷间消毒记录
	AcidShop7  float32         `json:"acid_shop_7"`  //氯含量
	AcidShop8  DisinfectRecord `json:"acid_shop_8"`  //消毒记录
	AcidShop9  float32         `json:"acid_shop_9"`  //工作服 功率
	AcidShop10 float32         `json:"acid_shop_10"` //工作服 时间
}

type FroShop struct {
	//冷冻库
	gorm.Model
	FroShop1  float32         `json:"fro_shop_1"`  //紫外灭菌
	FroShop2  float32         `json:"fro_shop_2"`  //臭氧
	FroShop3  float32         `json:"fro_shop_3"`  //臭氧残留
	FroShop4  float32         `json:"fro_shop_4"`  //湿度
	FroShop5  float32         `json:"fro_shop_5"`  //温度
	FroShop6  float32         `json:"fro_shop_6"`  //预冷间消毒记录
	FroShop7  float32         `json:"fro_shop_7"`  //氯含量
	FroShop8  DisinfectRecord `json:"fro_shop_8"`  //消毒记录
	FroShop9  float32         `json:"fro_shop_9"`  //工作服 功率
	FroShop10 float32         `json:"fro_shop_10"` //工作服 时间
}

type PackShop struct {
	//包装车间
	gorm.Model
	PackShop1  float32         `json:"pack_shop_1"`  //紫外灭菌
	PackShop2  float32         `json:"pack_shop_2"`  //臭氧
	PackShop3  float32         `json:"pack_shop_3"`  //臭氧残留
	PackShop4  float32         `json:"pack_shop_4"`  //湿度
	PackShop5  float32         `json:"pack_shop_5"`  //温度
	PackShop6  float32         `json:"pack_shop_6"`  //预冷间消毒记录
	PackShop7  float32         `json:"pack_shop_7"`  //氯含量
	PackShop8  DisinfectRecord `json:"pack_shop_8"`  //消毒记录
	PackShop9  float32         `json:"pack_shop_9"`  //工作服 功率
	PackShop10 float32         `json:"pack_shop_10"` //工作服 时间
}

type StaUni struct {
	//员工工作服
	gorm.Model
	StaUni1  float32         `json:"sta_uni_1"`  //紫外灭菌
	StaUni2  float32         `json:"sta_uni_2"`  //臭氧
	StaUni3  float32         `json:"sta_uni_3"`  //臭氧残留
	StaUni4  float32         `json:"sta_uni_4"`  //湿度
	StaUni5  float32         `json:"sta_uni_5"`  //温度
	StaUni6  float32         `json:"sta_uni_6"`  //预冷间消毒记录
	StaUni7  float32         `json:"sta_uni_7"`  //氯含量
	StaUni8  DisinfectRecord `json:"sta_uni_8"`  //消毒记录
	StaUni9  float32         `json:"sta_uni_9"`  //工作服 功率
	StaUni10 float32         `json:"sta_uni_10"` //工作服 时间
}
type DisinfectRecord struct {
	//消毒记录
	DisRecord1 string  `json:"dis_record_1"` //方式
	DisRecord2 float32 `json:"dis_record_2"` //浓度
	DisRecord3 string  `json:"dis_record_3"` //班次
	DisRecord4 string  `json:"dis_record_4"` //器具
	DisRecord5 string  `json:"dis_record_5"` //环境
}

type SlaEnvLigRec struct {
	//屠宰环境光照记录
	gorm.Model
	SlaInfoMonID  uint    `json:"sla_info_mon_id"`
	SlaEnvLigRec1 float32 `json:"sla_env_lig_rec_1"` //屠宰间环境
	SlaEnvLigRec2 float32 `json:"sla_env_lig_rec_2"` //车间
	SlaEnvLigRec3 float32 `json:"sla_env_lig_rec_3"` //检疫
	SlaEnvLigRec4 float32 `json:"sla_env_lig_rec_4"` //预冷通道
}

type ToNumGermMon struct {
	//接触面菌落总数监控
	gorm.Model
	SlaInfoMonID  uint    `json:"sla_info_mon_id"`
	ToNumGermMon1 float32 `json:"to_num_germ_mon_1"` //喷淋后胴体
	ToNumGermMon2 float32 `json:"to_num_germ_mon_2"` //剔骨案板
	ToNumGermMon3 float32 `json:"to_num_germ_mon_3"` //传送带
	ToNumGermMon4 float32 `json:"to_num_germ_mon_4"` //围裙
	ToNumGermMon5 float32 `json:"to_num_germ_mon_5"` //手套
	ToNumGermMon6 float32 `json:"to_num_germ_mon_6"` //锯骨机
	ToNumGermMon7 float32 `json:"to_num_germ_mon_7"` //刀具
	ToNumGermMon8 float32 `json:"to_num_germ_mon_8"` //地面与墙面
}

type AirNumGermMon struct {
	//空气菌落总数监控
	gorm.Model
	SlaInfoMonID   uint    `json:"sla_info_mon_id"`
	AirNumGermMon1 float32 `json:"air_num_germ_mon_1"` //屠宰车间
	AirNumGermMon2 float32 `json:"air_num_germ_mon_2"` //分割车间
	AirNumGermMon3 float32 `json:"air_num_germ_mon_3"` //排酸车间
}

type SlaughterWaterQualityMon struct {
	//水质监控
	gorm.Model
	SlaughterWaterMicroIndex SlaughterWaterMicroIndex `gorm:"foreignKey:SlaughterWaterQualityMonID; references:ID" json:"slaughter_water_micro_index"` //微生物指标
	SlaughterWaterToxinIndex SlaughterWaterToxinIndex `gorm:"foreignKey:SlaughterWaterQualityMonID; references:ID" json:"slaughter_water_toxin_index"` //毒理性指标
	OapGciSla                OapGciSla                `gorm:"foreignKey:WaterQualityMonID; references:ID" json:"oap_gci_sla"`                          //感官性状和一般化学指标
	MicroIndexWaterMonSla    MicroIndexWaterMonSla    `gorm:"foreignKey:WaterQualityMonID; references:ID" json:"micro_index_water_mon_sla"`            //微生物指标屠宰阶段 和 屠宰污水检测记录
	ToxinIndexSla            SlaughterToxinIndex      `gorm:"foreignKey:WaterQualityMonID; references:ID" json:"toxin_index_sla"`                      //屠宰污水毒理指标
}

type SlaughterWaterMicroIndex struct {
	gorm.Model
	SlaughterWaterQualityMonID uint    `json:"slaughter_water_quality_mon_id"`
	SlaughterWaterMicroIndex1  float32 `json:"slaughter_water_micro_index_1"` //总大肠菌群
	SlaughterWaterMicroIndex2  float32 `json:"slaughter_water_micro_index_2"` //大肠埃希氏菌
	SlaughterWaterMicroIndex3  float32 `json:"slaughter_water_micro_index_3"` //菌落总数
}

type SlaughterWaterToxinIndex struct {
	//毒理性指标
	gorm.Model
	SlaughterWaterQualityMonID uint    `json:"slaughter_water_quality_mon_id"`
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

type OapGciSla struct {
	//感官性状和一般化学指标
	gorm.Model
	SlaughterWaterQualityMonID uint    `json:"slaughter_water_quality_mon_id"`
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
	SlaughterWaterQualityMonID uint    `json:"slaughter_water_quality_mon_id"`
	MicroIndexWaterMonSla1     float32 `json:"micro_index_water_mon_sla_1"` //贾第鞭毛虫/（个/10L)
	MicroIndexWaterMonSla2     float32 `json:"micro_index_water_mon_sla_2"` //隐孢子虫/（个/10L）
	MicroIndexWaterMonSla3     float32 `json:"micro_index_water_mon_sla_3"` //生化需氧
	MicroIndexWaterMonSla4     float32 `json:"micro_index_water_mon_sla_4"` //化学需氧
}

type SlaughterToxinIndex struct {
	//屠宰阶段毒理性指标
	gorm.Model
	SlaughterWaterQualityMonID uint    `json:"slaughter_water_quality_mon_id"`
	ToxinIndexSla1             float32 `json:"toxin_index_sla_1"`  //锑/（mg/L）
	ToxinIndexSla2             float32 `json:"toxin_index_sla_2"`  //钡/（mg/L）
	ToxinIndexSla3             float32 `json:"toxin_index_sla_3"`  //铍/（mg/L）
	ToxinIndexSla4             float32 `json:"toxin_index_sla_4"`  //硼/（mg/L）
	ToxinIndexSla5             float32 `json:"toxin_index_sla_5"`  //钼/（mg/L）
	ToxinIndexSla6             float32 `json:"toxin_index_sla_6"`  //镍/（mg/L）
	ToxinIndexSla7             float32 `json:"toxin_index_sla_7"`  //银/（mg/L）
	ToxinIndexSla8             float32 `json:"toxin_index_sla_8"`  //铊/（mg/L）
	ToxinIndexSla9             float32 `json:"toxin_index_sla_9"`  //硒/（mg/L）
	ToxinIndexSla10            float32 `json:"toxin_index_sla_10"` //高氯酸钾/（mg/L）
	ToxinIndexSla11            float32 `json:"toxin_index_sla_11"` //二氯甲烷/（mg/L）
	ToxinIndexSla12            float32 `json:"toxin_index_sla_12"` //1，2-二氯甲烷/（mg/L）
	ToxinIndexSla13            float32 `json:"toxin_index_sla_13"` //四氯化碳/（mg/L）
	ToxinIndexSla14            float32 `json:"toxin_index_sla_14"` //氯乙烯/（mg/L）
	ToxinIndexSla15            float32 `json:"toxin_index_sla_15"` //1，1-二氯乙烯/（mg/L）
	ToxinIndexSla16            float32 `json:"toxin_index_sla_16"` //1，2-二氯乙烯（总量）/（mg/L）
	ToxinIndexSla17            float32 `json:"toxin_index_sla_17"` //三氯乙烯/（mg/L）
	ToxinIndexSla18            float32 `json:"toxin_index_sla_18"` //四氯乙烯/（mg/L）
	ToxinIndexSla19            float32 `json:"toxin_index_sla_19"` //六氯丁二烯/（mg/L）
	ToxinIndexSla20            float32 `json:"toxin_index_sla_20"` //苯/（mg/L）
	ToxinIndexSla21            float32 `json:"toxin_index_sla_21"` //甲苯/（mg/L）
	ToxinIndexSla22            float32 `json:"toxin_index_sla_22"` //二甲苯（总量）/（mg/L）
	ToxinIndexSla23            float32 `json:"toxin_index_sla_23"` //苯乙烯/（mg/L）
	ToxinIndexSla24            float32 `json:"toxin_index_sla_24"` //氯苯/（mg/L）
	ToxinIndexSla25            float32 `json:"toxin_index_sla_25"` //1，4-二氯苯/（mg/L）
	ToxinIndexSla26            float32 `json:"toxin_index_sla_26"` //三氯苯（总量）/（mg/L）
	ToxinIndexSla27            float32 `json:"toxin_index_sla_27"` //六氯苯/（mg/L）
	ToxinIndexSla28            float32 `json:"toxin_index_sla_28"` //七氯/（mg/L）
	ToxinIndexSla29            float32 `json:"toxin_index_sla_29"` //马拉硫磷/（mg/L）
	ToxinIndexSla30            float32 `json:"toxin_index_sla_30"` //乐果/（mg/L）
	ToxinIndexSla31            float32 `json:"toxin_index_sla_31"` //灭草松/（mg/L）
	ToxinIndexSla32            float32 `json:"toxin_index_sla_32"` //百菌清/（mg/L）
	ToxinIndexSla33            float32 `json:"toxin_index_sla_33"` //呋喃丹/（mg/L）
	ToxinIndexSla34            float32 `json:"toxin_index_sla_34"` //毒死蜱/（mg/L）
	ToxinIndexSla35            float32 `json:"toxin_index_sla_35"` //草甘膦/（mg/L）
	ToxinIndexSla36            float32 `json:"toxin_index_sla_36"` //敌敌畏/（mg/L）
	ToxinIndexSla37            float32 `json:"toxin_index_sla_37"` //莠去津/（mg/L）
	ToxinIndexSla38            float32 `json:"toxin_index_sla_38"` //溴氰菊酯/（mg/L）
	ToxinIndexSla39            float32 `json:"toxin_index_sla_39"` //2，4-滴/（mg/L）
	ToxinIndexSla40            float32 `json:"toxin_index_sla_40"` //乙草胺/（mg/L）
	ToxinIndexSla41            float32 `json:"toxin_index_sla_41"` //五氯酚/（mg/L）
	ToxinIndexSla42            float32 `json:"toxin_index_sla_42"` //2，4，6-三氯酚/（mg/L）
	ToxinIndexSla43            float32 `json:"toxin_index_sla_43"` //苯并（a）芘/（mg/L）
	ToxinIndexSla44            float32 `json:"toxin_index_sla_44"` //邻苯二甲酸二（2-乙基己基）脂/（mg/L）
	ToxinIndexSla45            float32 `json:"toxin_index_sla_45"` //丙烯酰胺/（mg/L）
	ToxinIndexSla46            float32 `json:"toxin_index_sla_46"` //环氧氯丙烷/（mg/L）
	ToxinIndexSla47            float32 `json:"toxin_index_sla_47"` //微囊藻毒素-LR（藻类爆发情况发生时）/（mg/L）
}
