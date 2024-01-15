package request

type ReqNewSlaughterBatch struct {
	HouseNumber string `json:"house_number" form:"house_number"`
	Worker      string `json:"worker" form:"worker"`
	PrePID      string `json:"pre_pid" form:"pre_pid"`
	CowNumber   string `json:"cow_number" form:"cow_number"`
}

type ReqEndSlaughter struct {
	BatchNumber string `json:"batch_number" form:"batch_number"`
	Worker      string `json:"worker" form:"worker"`
	HouseNumber string `json:"house_number" form:"house_number"`

	SlaughterDisinfectHotWaterTempMoni1 float64 `json:"slaughter_disinfect_hot_water_temp_moni_1" form:"slaughter_disinfect_hot_water_temp_moni_1"` //预剥皮
	SlaughterDisinfectHotWaterTempMoni2 float64 `json:"slaughter_disinfect_hot_water_temp_moni_2" form:"slaughter_disinfect_hot_water_temp_moni_2"` //检疫台
	SlaughterDisinfectHotWaterTempMoni3 float64 `json:"slaughter_disinfect_hot_water_temp_moni_3" form:"slaughter_disinfect_hot_water_temp_moni_3"` //去皮后修正
	SlaughterDisinfectHotWaterTempMoni4 float64 `json:"slaughter_disinfect_hot_water_temp_moni_4" form:"slaughter_disinfect_hot_water_temp_moni_4"` //去脏后
	SlaughterDisinfectHotWaterTempMoni5 float64 `json:"slaughter_disinfect_hot_water_temp_moni_5" form:"slaughter_disinfect_hot_water_temp_moni_5"` //剔骨台
	SlaughterDisinfectHotWaterTempMoni6 float64 `json:"slaughter_disinfect_hot_water_temp_moni_6" form:"slaughter_disinfect_hot_water_temp_moni_6"` //洗手台

	Stun1 float64 `json:"stun_1" form:"stun_1"` //电压
	Stun2 float64 `json:"stun_2" form:"stun_2"` //电流
	Stun3 float64 `json:"stun_3" form:"stun_3"` //作用时间

	BleedElectronic1 float64 `json:"bleed_electronic_1" form:"bleed_electronic_1"` //放血刀热水消毒
	BleedElectronic2 float64 `json:"bleed_electronic_2" form:"bleed_electronic_2"` //电流
	BleedElectronic3 float64 `json:"bleed_electronic_3" form:"bleed_electronic_3"` //作用时间
	BleedElectronic4 float64 `json:"bleed_electronic_4" form:"bleed_electronic_4"` //电刺激参数
	BleedElectronic5 float64 `json:"bleed_electronic_5" form:"bleed_electronic_5"` //电刺激时间

	PreSlaQuanPic1 string `json:"pre_sla_quan_pic_1" form:"pre_sla_quan_pic_1"` //咬肌照片
	PreSlaQuanPic2 string `json:"pre_sla_quan_pic_2" form:"pre_sla_quan_pic_2"` //舌头
	PreSlaQuanPic3 string `json:"pre_sla_quan_pic_3" form:"pre_sla_quan_pic_3"` //额下
	PreSlaQuanPic4 string `json:"pre_sla_quan_pic_4" form:"pre_sla_quan_pic_4"` //咽喉
	PreSlaQuanPic5 string `json:"pre_sla_quan_pic_5" form:"pre_sla_quan_pic_5"` //口腔
	PreSlaQuanPic6 string `json:"pre_sla_quan_pic_6" form:"pre_sla_quan_pic_6"` //肾脏
	PreSlaQuanPic7 string `json:"pre_sla_quan_pic_7" form:"pre_sla_quan_pic_7"` //肝脏
	PreSlaQuanPic8 string `json:"pre_sla_quan_pic_8" form:"pre_sla_quan_pic_8"` //心脏
	PreSlaQuanPic9 string `json:"pre_sla_quan_pic_9" form:"pre_sla_quan_pic_9"` //淋巴结

	SlaughterAnalAfterSlaQuanCar1 string `json:"slaughter_anal_after_sla_quan_car_1" form:"slaughter_anal_after_sla_quan_car_1"` //肛门结扎
	SlaughterAnalAfterSlaQuanCar2 string `json:"slaughter_anal_after_sla_quan_car_2" form:"slaughter_anal_after_sla_quan_car_2"` //胆囊破损情况拍照
	SlaughterAnalAfterSlaQuanCar3 string `json:"slaughter_anal_after_sla_quan_car_3" form:"slaughter_anal_after_sla_quan_car_3"` //腹股沟淋巴
	SlaughterAnalAfterSlaQuanCar4 string `json:"slaughter_anal_after_sla_quan_car_4" form:"slaughter_anal_after_sla_quan_car_4"` //整个胴体

	AnalMeatPhMoni1 float64 `json:"anal_meat_ph_moni_1" form:"anal_meat_ph_moni_1"` //宰后0 min
	AnalMeatPhMoni2 float64 `json:"anal_meat_ph_moni_2" form:"anal_meat_ph_moni_2"` //宰后45min
	AnalMeatPhMoni3 float64 `json:"anal_meat_ph_moni_3" form:"anal_meat_ph_moni_3"` //劈半后胴体
	AnalMeatPhMoni4 float64 `json:"anal_meat_ph_moni_4" form:"anal_meat_ph_moni_4"` //排酸过程中胴体
	AnalMeatPhMoni5 float64 `json:"anal_meat_ph_moni_5" form:"anal_meat_ph_moni_5"` //冷鲜肉

	AnalCutWeight1  float64 `json:"anal_cut_weight_1" form:"anal_cut_weight_1"`   //劈半后胴体 温度
	AnalCutWeight2  float64 `json:"anal_cut_weight_2" form:"anal_cut_weight_2"`   //排酸过程中胴体 温度
	AnalCutWeight3  float64 `json:"anal_cut_weight_3" form:"anal_cut_weight_3"`   //屠宰到入排酸库的时间
	AnalCutWeight4  float64 `json:"anal_cut_weight_4" form:"anal_cut_weight_4"`   //屠宰到入排酸库的记录
	AnalCutWeight5  float64 `json:"anal_cut_weight_5" form:"anal_cut_weight_5"`   //排酸库胴体间隙
	AnalCutWeight6  float64 `json:"anal_cut_weight_6" form:"anal_cut_weight_6"`   //分割前胴体重量记录，称重时间
	AnalCutWeight7  float64 `json:"anal_cut_weight_7" form:"anal_cut_weight_7"`   //分割前胴体重量记录，重量
	AnalCutWeight8  float64 `json:"anal_cut_weight_8" form:"anal_cut_weight_8"`   //分割前胴体重量记录，酮体编号
	AnalCutWeight9  float64 `json:"anal_cut_weight_9" form:"anal_cut_weight_9"`   //分割前胴体温度
	AnalCutWeight10 float64 `json:"anal_cut_weight_10" form:"anal_cut_weight_10"` //分割刀温度
	AnalCutWeight11 float64 `json:"anal_cut_weight_11" form:"anal_cut_weight_11"` //分割刀记录
	AnalCutWeight12 float64 `json:"anal_cut_weight_12" form:"anal_cut_weight_12"` //温湿度监控

	ToNumGermMon1 float64 `json:"to_num_germ_mon_1" form:"to_num_germ_mon_1"` //喷淋后胴体
	ToNumGermMon2 float64 `json:"to_num_germ_mon_2" form:"to_num_germ_mon_2"` //剔骨案板
	ToNumGermMon3 float64 `json:"to_num_germ_mon_3" form:"to_num_germ_mon_3"` //传送带
	ToNumGermMon4 float64 `json:"to_num_germ_mon_4" form:"to_num_germ_mon_4"` //围裙
	ToNumGermMon5 float64 `json:"to_num_germ_mon_5" form:"to_num_germ_mon_5"` //手套
	ToNumGermMon6 float64 `json:"to_num_germ_mon_6" form:"to_num_germ_mon_6"` //锯骨机
	ToNumGermMon7 float64 `json:"to_num_germ_mon_7" form:"to_num_germ_mon_7"` //刀具
	ToNumGermMon8 float64 `json:"to_num_germ_mon_8" form:"to_num_germ_mon_8"` //地面与墙面

	AirNumGermMon1 float64 `json:"air_num_germ_mon_1" form:"air_num_germ_mon_1"` //屠宰车间
	AirNumGermMon2 float64 `json:"air_num_germ_mon_2" form:"air_num_germ_mon_2"` //分割车间
	AirNumGermMon3 float64 `json:"air_num_germ_mon_3" form:"air_num_germ_mon_3"` //排酸车间

	PreSlaDietManage1 string `json:"pre_sla_diet_manage_1" form:"pre_sla_diet_manage_1"` //牛进场时间
	PreSlaDietManage2 string `json:"pre_sla_diet_manage_2" form:"pre_sla_diet_manage_2"` //牛停水时间
	PreSlaDietManage3 string `json:"pre_sla_diet_manage_3" form:"pre_sla_diet_manage_3"` //牛禁食时间
	PreSlaDietManage4 string `json:"pre_sla_diet_manage_4" form:"pre_sla_diet_manage_4"` //饮水时间
	PreSlaDietManage5 string `json:"pre_sla_diet_manage_5" form:"pre_sla_diet_manage_5"` //送宰时间

	// PreSlaInfoRecID uint   `json:"pre_sla_info_rec_id"`
	// PreSlaPicAndEn1 string `json:"pre_sla_pic_and_en_1"` //待宰动物视频
	// PreSlaPicAndEn2 string `json:"pre_sla_pic_and_en_2"` //待宰动物头部
	// PreSlaPicAndEn3 string `json:"pre_sla_pic_and_en_3"` //待宰动物皮毛相片采集
	// PreSlaPicAndEn4 string `json:"pre_sla_pic_and_en_4"` //照度
	// PreSlaPicAndEn5 string `json:"pre_sla_pic_and_en_5"` //噪声
	// PreSlaPicAndEn6 string `json:"pre_sla_pic_and_en_6"` //光照时间

	EnvirTemperature      string `json:"envir_temperature"`
	EnvirLighting         string `json:"envir_lighting"`
	ShockVoltage          string `json:"shock_voltage"`
	BleedingTime          string `json:"bleeding_time"`
	EleShockTime          string `json:"ele_shock_time"`
	KnifeDisinfectionTime string `json:"knife_disinfection_time"`
}

type ReqNewSlaughterProduct struct {
	BatchNumber string  `json:"batch_number" form:"batch_number"`
	Worker      string  `json:"worker" form:"worker"`
	HouseNumber string  `json:"house_number" form:"house_number"`
	Type        int     `json:"type" form:"type"`
	Weight      float64 `json:"weight" form:"weight"`
}

type ReqSendToPackage struct {
	ProductNumber      string `json:"product_number" form:"product_number"`
	Operator           string `json:"operator" form:"operator"`
	PackageHouseNumber string `json:"package_house_number" form:"package_house_number"`
}

type ReqUploadPreCoolShopData struct {
	//预冷间
	HouseNumber       string  `json:"house_number"`
	TimestampRecordAt int64   `json:"timestamp_record_at"` //记录时间戳
	PreCoolShop1      float32 `json:"pre_cool_shop_1"`     //预冷间温度
	PreCoolShop2      float32 `json:"pre_cool_shop_2"`     //预冷间湿度
	PreCoolShop3      float32 `json:"pre_cool_shop_3"`     //副产物温度
}

type ReqUploadSlaughterShopData struct {
	//屠宰车间
	HouseNumber       string  `json:"house_number"`
	TimestampRecordAt int64   `json:"timestamp_record_at"` //记录时间戳
	SlaShop1          float32 `json:"sla_shop_1"`          //紫外灭菌
	SlaShop2          float32 `json:"sla_shop_2"`          //臭氧
	SlaShop3          float32 `json:"sla_shop_3"`          //臭氧残留
	SlaShop4          float32 `json:"sla_shop_4"`          //湿度
	SlaShop5          float32 `json:"sla_shop_5"`          //温度
	SlaShop6          float32 `json:"sla_shop_6"`          //预冷间消毒记录
	SlaShop7          float32 `json:"sla_shop_7"`          //氯含量
	SlaShop8          float32 `json:"sla_shop_8"`          //工作服 功率
	SlaShop9          float32 `json:"sla_shop_9"`          //工作服 时间
	SlaShop10         string  `json:"sla_shop_10"`         //消毒记录 方式
	SlaShop11         string  `json:"sla_shop_11"`         //消毒记录 浓度
	SlaShop12         string  `json:"sla_shop_12"`         //消毒记录 班次
	SlaShop13         string  `json:"sla_shop_13"`         //消毒记录 器具
	SlaShop14         string  `json:"sla_shop_14"`         //消毒记录 环境
}

type ReqUploadDivisionShopData struct {
	//分割车间
	HouseNumber       string  `json:"house_number"`
	TimestampRecordAt int64   `json:"timestamp_record_at"` //记录时间戳
	DivShop1          float32 `json:"div_shop_1"`          //紫外灭菌
	DivShop2          float32 `json:"div_shop_2"`          //臭氧
	DivShop3          float32 `json:"div_shop_3"`          //臭氧残留
	DivShop4          float32 `json:"div_shop_4"`          //湿度
	DivShop5          float32 `json:"div_shop_5"`          //温度
	DivShop6          float32 `json:"div_shop_6"`          //预冷间消毒记录
	DivShop7          float32 `json:"div_shop_7"`          //氯含量
	DivShop8          float32 `json:"div_shop_8"`          //工作服 功率
	DivShop9          float32 `json:"div_shop_9"`          //工作服 时间
	DivShop10         string  `json:"div_shop_10"`         //消毒记录 方式
	DivShop11         string  `json:"div_shop_11"`         //消毒记录 浓度
	DivShop12         string  `json:"div_shop_12"`         //消毒记录 班次
	DivShop13         string  `json:"div_shop_13"`         //消毒记录 器具
	DivShop14         string  `json:"div_shop_14"`         //消毒记录 环境
}

type ReqUploadAcidShopData struct {
	//排酸车间
	HouseNumber       string  `json:"house_number"`
	TimestampRecordAt int64   `json:"timestamp_record_at"` //记录时间戳
	AcidShop1         float32 `json:"acid_shop_1"`         //紫外灭菌
	AcidShop2         float32 `json:"acid_shop_2"`         //臭氧
	AcidShop3         float32 `json:"acid_shop_3"`         //臭氧残留
	AcidShop4         float32 `json:"acid_shop_4"`         //湿度
	AcidShop5         float32 `json:"acid_shop_5"`         //温度
	AcidShop6         float32 `json:"acid_shop_6"`         //预冷间消毒记录
	AcidShop7         float32 `json:"acid_shop_7"`         //氯含量
	AcidShop8         float32 `json:"acid_shop_8"`         //工作服 功率
	AcidShop9         float32 `json:"acid_shop_9"`         //工作服 时间
	AcidShop10        string  `json:"acid_shop_10"`        //消毒记录 方式
	AcidShop11        string  `json:"acid_shop_11"`        //消毒记录 浓度
	AcidShop12        string  `json:"acid_shop_12"`        //消毒记录 班次
	AcidShop13        string  `json:"acid_shop_13"`        //消毒记录 器具
	AcidShop14        string  `json:"acid_shop_14"`        //消毒记录 环境
}

type ReqUploadFrozenShopData struct {
	//冷冻库
	HouseNumber       string  `json:"house_number"`
	TimestampRecordAt int64   `json:"timestamp_record_at"` //记录时间戳
	FroShop1          float32 `json:"fro_shop_1"`          //紫外灭菌
	FroShop2          float32 `json:"fro_shop_2"`          //臭氧
	FroShop3          float32 `json:"fro_shop_3"`          //臭氧残留
	FroShop4          float32 `json:"fro_shop_4"`          //湿度
	FroShop5          float32 `json:"fro_shop_5"`          //温度
	FroShop6          float32 `json:"fro_shop_6"`          //预冷间消毒记录
	FroShop7          float32 `json:"fro_shop_7"`          //氯含量
	FroShop8          float32 `json:"fro_shop_8"`          //工作服 功率
	FroShop9          float32 `json:"fro_shop_9"`          //工作服 时间
	FroShop10         string  `json:"fro_shop_10"`         //消毒记录 方式
	FroShop11         string  `json:"fro_shop_11"`         //消毒记录 浓度
	FroShop12         string  `json:"fro_shop_12"`         //消毒记录 班次
	FroShop13         string  `json:"fro_shop_13"`         //消毒记录 器具
	FroShop14         string  `json:"fro_shop_14"`         //消毒记录 环境
}

type ReqSlaughterSensorData struct {
	HouseNumber    string `json:"house_number" form:"house_number"`
	StartTimestamp int64  `json:"start_timestamp" form:"start_timestamp"`
	EndTimestamp   int64  `json:"end_timestamp" form:"end_timestamp"`
}

type ReqUploadSlaughterWaterQualityData struct {
	HouseNumber              string                      `json:"house_number"`
	TimestampRecordAt        int64                       `json:"timestamp_record_at"` //记录时间戳
	SlaughterWaterMicroIndex ReqSlaughterWaterMicroIndex `json:"slaughter_water_micro_index"`
	OapGciSla                ReqOapGciSla                `json:"oap_gci_sla"`
	SlaughterToxinIndex      ReqSlaughterToxinIndex      `json:"slaughter_toxin_index"`
}

type ReqSlaughterWaterMicroIndex struct {
	//微生物监控
	SlaughterWaterMicroIndex1 float32 `json:"slaughter_water_micro_index_1"` //总大肠菌群
	SlaughterWaterMicroIndex2 float32 `json:"slaughter_water_micro_index_2"` //大肠埃希氏菌
	SlaughterWaterMicroIndex3 float32 `json:"slaughter_water_micro_index_3"` //菌落总数
}

type ReqOapGciSla struct {
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

type ReqMicroIndexWaterMonSla struct {
	//微生物指标屠宰阶段 和 屠宰污水检测记录
	MicroIndexWaterMonSla1 float32 `json:"micro_index_water_mon_sla_1"` //贾第鞭毛虫/（个/10L)
	MicroIndexWaterMonSla2 float32 `json:"micro_index_water_mon_sla_2"` //隐孢子虫/（个/10L）
	MicroIndexWaterMonSla3 float32 `json:"micro_index_water_mon_sla_3"` //生化需氧
	MicroIndexWaterMonSla4 float32 `json:"micro_index_water_mon_sla_4"` //化学需氧
}

type ReqSlaughterToxinIndex struct {
	//屠宰阶段毒理性指标
	ToxinIndexSla1           float32                     `json:"toxin_index_sla_1"`                                                                  //锑/（mg/L）
	ToxinIndexSla2           float32                     `json:"toxin_index_sla_2"`                                                                  //钡/（mg/L）
	ToxinIndexSla3           float32                     `json:"toxin_index_sla_3"`                                                                  //铍/（mg/L）
	ToxinIndexSla4           float32                     `json:"toxin_index_sla_4"`                                                                  //硼/（mg/L）
	ToxinIndexSla5           float32                     `json:"toxin_index_sla_5"`                                                                  //钼/（mg/L）
	ToxinIndexSla6           float32                     `json:"toxin_index_sla_6"`                                                                  //镍/（mg/L）
	ToxinIndexSla7           float32                     `json:"toxin_index_sla_7"`                                                                  //银/（mg/L）
	ToxinIndexSla8           float32                     `json:"toxin_index_sla_8"`                                                                  //铊/（mg/L）
	ToxinIndexSla9           float32                     `json:"toxin_index_sla_9"`                                                                  //硒/（mg/L）
	ToxinIndexSla10          float32                     `json:"toxin_index_sla_10"`                                                                 //高氯酸钾/（mg/L）
	ToxinIndexSla11          float32                     `json:"toxin_index_sla_11"`                                                                 //二氯甲烷/（mg/L）
	ToxinIndexSla12          float32                     `json:"toxin_index_sla_12"`                                                                 //1，2-二氯甲烷/（mg/L）
	ToxinIndexSla13          float32                     `json:"toxin_index_sla_13"`                                                                 //四氯化碳/（mg/L）
	ToxinIndexSla14          float32                     `json:"toxin_index_sla_14"`                                                                 //氯乙烯/（mg/L）
	ToxinIndexSla15          float32                     `json:"toxin_index_sla_15"`                                                                 //1，1-二氯乙烯/（mg/L）
	ToxinIndexSla16          float32                     `json:"toxin_index_sla_16"`                                                                 //1，2-二氯乙烯（总量）/（mg/L）
	ToxinIndexSla17          float32                     `json:"toxin_index_sla_17"`                                                                 //三氯乙烯/（mg/L）
	ToxinIndexSla18          float32                     `json:"toxin_index_sla_18"`                                                                 //四氯乙烯/（mg/L）
	ToxinIndexSla19          float32                     `json:"toxin_index_sla_19"`                                                                 //六氯丁二烯/（mg/L）
	ToxinIndexSla20          float32                     `json:"toxin_index_sla_20"`                                                                 //苯/（mg/L）
	ToxinIndexSla21          float32                     `json:"toxin_index_sla_21"`                                                                 //甲苯/（mg/L）
	ToxinIndexSla22          float32                     `json:"toxin_index_sla_22"`                                                                 //二甲苯（总量）/（mg/L）
	ToxinIndexSla23          float32                     `json:"toxin_index_sla_23"`                                                                 //苯乙烯/（mg/L）
	ToxinIndexSla24          float32                     `json:"toxin_index_sla_24"`                                                                 //氯苯/（mg/L）
	ToxinIndexSla25          float32                     `json:"toxin_index_sla_25"`                                                                 //1，4-二氯苯/（mg/L）
	ToxinIndexSla26          float32                     `json:"toxin_index_sla_26"`                                                                 //三氯苯（总量）/（mg/L）
	ToxinIndexSla27          float32                     `json:"toxin_index_sla_27"`                                                                 //六氯苯/（mg/L）
	ToxinIndexSla28          float32                     `json:"toxin_index_sla_28"`                                                                 //七氯/（mg/L）
	ToxinIndexSla29          float32                     `json:"toxin_index_sla_29"`                                                                 //马拉硫磷/（mg/L）
	ToxinIndexSla30          float32                     `json:"toxin_index_sla_30"`                                                                 //乐果/（mg/L）
	ToxinIndexSla31          float32                     `json:"toxin_index_sla_31"`                                                                 //灭草松/（mg/L）
	ToxinIndexSla32          float32                     `json:"toxin_index_sla_32"`                                                                 //百菌清/（mg/L）
	ToxinIndexSla33          float32                     `json:"toxin_index_sla_33"`                                                                 //呋喃丹/（mg/L）
	ToxinIndexSla34          float32                     `json:"toxin_index_sla_34"`                                                                 //毒死蜱/（mg/L）
	ToxinIndexSla35          float32                     `json:"toxin_index_sla_35"`                                                                 //草甘膦/（mg/L）
	ToxinIndexSla36          float32                     `json:"toxin_index_sla_36"`                                                                 //敌敌畏/（mg/L）
	ToxinIndexSla37          float32                     `json:"toxin_index_sla_37"`                                                                 //莠去津/（mg/L）
	ToxinIndexSla38          float32                     `json:"toxin_index_sla_38"`                                                                 //溴氰菊酯/（mg/L）
	ToxinIndexSla39          float32                     `json:"toxin_index_sla_39"`                                                                 //2，4-滴/（mg/L）
	ToxinIndexSla40          float32                     `json:"toxin_index_sla_40"`                                                                 //乙草胺/（mg/L）
	ToxinIndexSla41          float32                     `json:"toxin_index_sla_41"`                                                                 //五氯酚/（mg/L）
	ToxinIndexSla42          float32                     `json:"toxin_index_sla_42"`                                                                 //2，4，6-三氯酚/（mg/L）
	ToxinIndexSla43          float32                     `json:"toxin_index_sla_43"`                                                                 //苯并（a）芘/（mg/L）
	ToxinIndexSla44          float32                     `json:"toxin_index_sla_44"`                                                                 //邻苯二甲酸二（2-乙基己基）脂/（mg/L）
	ToxinIndexSla45          float32                     `json:"toxin_index_sla_45"`                                                                 //丙烯酰胺/（mg/L）
	ToxinIndexSla46          float32                     `json:"toxin_index_sla_46"`                                                                 //环氧氯丙烷/（mg/L）
	ToxinIndexSla47          float32                     `json:"toxin_index_sla_47"`                                                                 //微囊藻毒素-LR（藻类爆发情况发生时）/（mg/L）
	SlaughterWaterToxinIndex ReqSlaughterWaterToxinIndex `gorm:"foreignKey:SlaughterToxinIndexID; references:ID" json:"slaughter_water_toxin_index"` //其他毒理指标
}

type ReqSlaughterWaterToxinIndex struct {
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

type ReqUploadStaffUniformData struct {
	//员工工作服
	HouseNumber       string  `json:"house_number"`
	TimestampRecordAt int64   `json:"timestamp_record_at"` //记录时间戳
	StaUni1           float32 `json:"sta_uni_1"`           //紫外灭菌
	StaUni2           float32 `json:"sta_uni_2"`           //臭氧
	StaUni3           float32 `json:"sta_uni_3"`           //臭氧残留
	StaUni4           float32 `json:"sta_uni_4"`           //湿度
	StaUni5           float32 `json:"sta_uni_5"`           //温度
	StaUni6           float32 `json:"sta_uni_6"`           //预冷间消毒记录
	StaUni7           float32 `json:"sta_uni_7"`           //氯含量
	StaUni8           float32 `json:"sta_uni_8"`           //工作服 功率
	StaUni9           float32 `json:"sta_uni_9"`           //工作服 时间
	StaUni10          float32 `json:"sta_uni_10"`          //消毒记录 方式
	StaUni11          float32 `json:"sta_uni_11"`          //消毒记录 浓度
	StaUni12          float32 `json:"sta_uni_12"`          //消毒记录 班次
	StaUni13          float32 `json:"sta_uni_13"`          //消毒记录 器具
	StaUni14          float32 `json:"sta_uni_15"`          //消毒记录 环境
}

type ReqUploadSlaughterLightRecord struct {
	//屠宰环境光照记录
	HouseNumber       string  `json:"house_number"`
	TimestampRecordAt int64   `json:"timestamp_record_at"` //记录时间戳
	SlaEnvLigRec1     float32 `json:"sla_env_lig_rec_1"`   //屠宰间环境
	SlaEnvLigRec2     float32 `json:"sla_env_lig_rec_2"`   //车间
	SlaEnvLigRec3     float32 `json:"sla_env_lig_rec_3"`   //检疫
	SlaEnvLigRec4     float32 `json:"sla_env_lig_rec_4"`   //预冷通道
}

type ReqSlaughterWasteWaterPerDay struct {
	TimeStamp                     int64   `json:"time_stamp"`
	HouseNumber                   string  `json:"house_number"`
	ReqSlaughterWasteWaterPerDay1 float32 `json:"req_slaughter_waste_water_per_day_1"`
	ReqSlaughterWasteWaterPerDay2 float32 `json:"req_slaughter_waste_water_per_day_2"`
	ReqSlaughterWasteWaterPerDay3 float32 `json:"req_slaughter_waste_water_per_day_3"`
	ReqSlaughterWasteWaterPerDay4 float32 `json:"req_slaughter_waste_water_per_day_4"`
}

type ReqSlaughterWasteResiduePerDay struct {
	TimeStamp                       int64   `json:"time_stamp"`
	HouseNumber                     string  `json:"house_number"`
	ReqSlaughterWasteResiduePerDay1 float32 `json:"req_slaughter_waste_residue_per_day_1"`
	ReqSlaughterWasteResiduePerDay2 float32 `json:"req_slaughter_waste_residue_per_day_2"`
	ReqSlaughterWasteResiduePerDay3 float32 `json:"req_slaughter_waste_residue_per_day_3"`
	ReqSlaughterWasteResiduePerDay4 float32 `json:"req_slaughter_waste_residue_per_day_4"`
}

type ReqSlaughterOdorPollutantsPerDay struct {
	TimeStamp                         int64   `json:"time_stamp"`
	HouseNumber                       string  `json:"house_number"`
	ReqSlaughterOdorPollutantsPerDay1 float32 `json:"req_slaughter_odor_pollutants_per_day_1"`
	ReqSlaughterOdorPollutantsPerDay2 float32 `json:"req_slaughter_odor_pollutants_per_day_2"`
	ReqSlaughterOdorPollutantsPerDay3 float32 `json:"req_slaughter_odor_pollutants_per_day_3"`
	ReqSlaughterOdorPollutantsPerDay4 float32 `json:"req_slaughter_odor_pollutants_per_day_4"`
}
