package slaughter

import "gorm.io/gorm"

/*
本文件结构是以牛为对象
*/

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
type SlaughterStun struct {
	//击晕参数
	gorm.Model
	Stun1 float32 `json:"stun_1"` //电压
	Stun2 float32 `json:"stun_2"` //电流
	Stun3 float32 `json:"stun_3"` //作用时间
}

type BleedElectronic struct {
	//放血和电刺激参数
	gorm.Model
<<<<<<< HEAD
	BleedElectronic1 float32 `json:"bleed_electronic_1"` //放血刀热水消毒
	BleedElectronic2 float32 `json:"bleed_electronic_2"` //电流
	BleedElectronic3 float32 `json:"bleed_electronic_3"` //作用时间
	BleedElectronic4 float32 `json:"bleed_electronic_4"` //电刺激参数
	BleedElectronic5 float32 `json:"bleed_electronic_5"` //电刺激时间
=======
	SlaughterProcedureMonitoringDataID *string `json:"slaughter_procedure_monitoring_data_id"`
	BleedElectronic1                   float64 `json:"bleed_electronic_1"` //放血刀热水消毒
	BleedElectronic2                   float64 `json:"bleed_electronic_2"` //电流
	BleedElectronic3                   float64 `json:"bleed_electronic_3"` //作用时间
	BleedElectronic4                   float64 `json:"bleed_electronic_4"` //电刺激参数
	BleedElectronic5                   float64 `json:"bleed_electronic_5"` //电刺激时间
>>>>>>> parent of 46e6215 (modify slaughter shops)
}

type PreSlaQuanPic struct {
	//宰前检疫工位照片
	//照片类型未定，先记录
	gorm.Model
	PreSlaQuanPic1 string `json:"pre_sla_quan_pic_1"` //咬肌照片
	PreSlaQuanPic2 string `json:"pre_sla_quan_pic_2"` //舌头
	PreSlaQuanPic3 string `json:"pre_sla_quan_pic_3"` //额下
	PreSlaQuanPic4 string `json:"pre_sla_quan_pic_4"` //咽喉
	PreSlaQuanPic5 string `json:"pre_sla_quan_pic_5"` //口腔
	PreSlaQuanPic6 string `json:"pre_sla_quan_pic_6"` //肾脏
	PreSlaQuanPic7 string `json:"pre_sla_quan_pic_7"` //肝脏
	PreSlaQuanPic8 string `json:"pre_sla_quan_pic_8"` //心脏
	PreSlaQuanPic9 string `json:"pre_sla_quan_pic_9"` //淋巴结
}

type SlaughterAnalAfterSlaQuanCar struct {
	//肛门结扎，宰后检疫，胴体图像采集
	//照片先只做记录
	gorm.Model
	SlaughterAnalAfterSlaQuanCar1 string `json:"slaughter_anal_after_sla_quan_car_1"` //肛门结扎
	SlaughterAnalAfterSlaQuanCar2 string `json:"slaughter_anal_after_sla_quan_car_2"` //胆囊破损情况拍照
	SlaughterAnalAfterSlaQuanCar3 string `json:"slaughter_anal_after_sla_quan_car_3"` //腹股沟淋巴
	SlaughterAnalAfterSlaQuanCar4 string `json:"slaughter_anal_after_sla_quan_car_4"` //整个胴体
}

type AnalMeatPhMoni struct {
	gorm.Model
	//胴体、肉pH监控
	AnalMeatPhMoni1 float32 `json:"anal_meat_ph_moni_1"` //宰后0 min
	AnalMeatPhMoni2 float32 `json:"anal_meat_ph_moni_2"` //宰后45min
	AnalMeatPhMoni3 float32 `json:"anal_meat_ph_moni_3"` //劈半后胴体
	AnalMeatPhMoni4 float32 `json:"anal_meat_ph_moni_4"` //排酸过程中胴体
	AnalMeatPhMoni5 float32 `json:"anal_meat_ph_moni_5"` //冷鲜肉
}

type AnalCutWeight struct {
	/*
		胴体、冷鲜肉温度监控
		屠宰到入排酸库的时间，记录
		排酸库胴体间隙
		分割前胴体重量记录
		分割前胴体温度，分割刀温度，记录，分割后肉称重
	*/
	gorm.Model
	AnalCutWeight1  float32 `json:"anal_cut_weight_1"`  //劈半后胴体 温度
	AnalCutWeight2  float32 `json:"anal_cut_weight_2"`  //排酸过程中胴体 温度
	AnalCutWeight3  float32 `json:"anal_cut_weight_3"`  //屠宰到入排酸库的时间
	AnalCutWeight4  float32 `json:"anal_cut_weight_4"`  //屠宰到入排酸库的记录
	AnalCutWeight5  float32 `json:"anal_cut_weight_5"`  //排酸库胴体间隙
	AnalCutWeight6  float32 `json:"anal_cut_weight_6"`  //分割前胴体重量记录，称重时间
	AnalCutWeight7  float32 `json:"anal_cut_weight_7"`  //分割前胴体重量记录，重量
	AnalCutWeight8  float32 `json:"anal_cut_weight_8"`  //分割前胴体重量记录，酮体编号
	AnalCutWeight9  float32 `json:"anal_cut_weight_9"`  //分割前胴体温度
	AnalCutWeight10 float32 `json:"anal_cut_weight_10"` //分割刀温度
	AnalCutWeight11 float32 `json:"anal_cut_weight_11"` //分割刀记录
	AnalCutWeight12 float32 `json:"anal_cut_weight_12"` //温湿度监控
}

type ToNumGermMon struct {
	//接触面菌落总数监控
	gorm.Model
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
	AirNumGermMon1 float32 `json:"air_num_germ_mon_1"` //屠宰车间
	AirNumGermMon2 float32 `json:"air_num_germ_mon_2"` //分割车间
	AirNumGermMon3 float32 `json:"air_num_germ_mon_3"` //排酸车间
}

type PreSlaDietManage struct {
	//宰前饮食管理
	gorm.Model
	PreSlaDietManage1 string `json:"pre_sla_diet_manage_1"` //牛进场时间
	PreSlaDietManage2 string `json:"pre_sla_diet_manage_2"` //牛停水时间
	PreSlaDietManage3 string `json:"pre_sla_diet_manage_3"` //牛禁食时间
	PreSlaDietManage4 string `json:"pre_sla_diet_manage_4"` //饮水时间
	PreSlaDietManage5 string `json:"pre_sla_diet_manage_5"` //送宰时间
}

type PreSlaPicAndEn struct {
	//类型后续调整，先做标记
	//宰前视频，图片，环境
	gorm.Model
	PreSlaInfoRecID uint   `json:"pre_sla_info_rec_id"`
	PreSlaPicAndEn1 string `json:"pre_sla_pic_and_en_1"` //待宰动物视频
	PreSlaPicAndEn2 string `json:"pre_sla_pic_and_en_2"` //待宰动物头部
	PreSlaPicAndEn3 string `json:"pre_sla_pic_and_en_3"` //待宰动物皮毛相片采集
	PreSlaPicAndEn4 string `json:"pre_sla_pic_and_en_4"` //照度
	PreSlaPicAndEn5 string `json:"pre_sla_pic_and_en_5"` //噪声
	PreSlaPicAndEn6 string `json:"pre_sla_pic_and_en_6"` //光照时间
}
