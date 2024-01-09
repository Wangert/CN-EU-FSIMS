package slaughter

type SlaughterDisinfectHotWaterTempMoniData struct {
	//消毒热水温度监控

	SlaughterDisinfectHotWaterTempMoni1 float32 `json:"slaughter_disinfect_hot_water_temp_moni_1"` //预剥皮
	SlaughterDisinfectHotWaterTempMoni2 float32 `json:"slaughter_disinfect_hot_water_temp_moni_2"` //检疫台
	SlaughterDisinfectHotWaterTempMoni3 float32 `json:"slaughter_disinfect_hot_water_temp_moni_3"` //去皮后修正
	SlaughterDisinfectHotWaterTempMoni4 float32 `json:"slaughter_disinfect_hot_water_temp_moni_4"` //去脏后
	SlaughterDisinfectHotWaterTempMoni5 float32 `json:"slaughter_disinfect_hot_water_temp_moni_5"` //剔骨台
	SlaughterDisinfectHotWaterTempMoni6 float32 `json:"slaughter_disinfect_hot_water_temp_moni_6"` //洗手台
}

func ToSlaughterDisinfectHotWaterTempMoniData(data *SlaughterDisinfectHotWaterTempMoni) SlaughterDisinfectHotWaterTempMoniData {
	return SlaughterDisinfectHotWaterTempMoniData{
		SlaughterDisinfectHotWaterTempMoni1: data.SlaughterDisinfectHotWaterTempMoni1,
		SlaughterDisinfectHotWaterTempMoni2: data.SlaughterDisinfectHotWaterTempMoni2,
		SlaughterDisinfectHotWaterTempMoni3: data.SlaughterDisinfectHotWaterTempMoni3,
		SlaughterDisinfectHotWaterTempMoni4: data.SlaughterDisinfectHotWaterTempMoni4,
		SlaughterDisinfectHotWaterTempMoni5: data.SlaughterDisinfectHotWaterTempMoni5,
		SlaughterDisinfectHotWaterTempMoni6: data.SlaughterDisinfectHotWaterTempMoni6,
	}
}

type SlaughterStunData struct {
	//击晕参数

	Stun1 float32 `json:"stun_1"` //电压
	Stun2 float32 `json:"stun_2"` //电流
	Stun3 float32 `json:"stun_3"` //作用时间
}

func ToSlaughterStunData(data *SlaughterStun) SlaughterStunData {
	return SlaughterStunData{
		Stun1: data.Stun1,
		Stun2: data.Stun2,
		Stun3: data.Stun3,
	}
}

type BleedElectronicData struct {
	//放血和电刺激参数

	BleedElectronic1 float32 `json:"bleed_electronic_1"` //放血刀热水消毒
	BleedElectronic2 float32 `json:"bleed_electronic_2"` //电流
	BleedElectronic3 float32 `json:"bleed_electronic_3"` //作用时间
	BleedElectronic4 float32 `json:"bleed_electronic_4"` //电刺激参数
	BleedElectronic5 float32 `json:"bleed_electronic_5"` //电刺激时间
}

func ToBleedElectronicData(data *BleedElectronic) BleedElectronicData {
	return BleedElectronicData{
		BleedElectronic1: data.BleedElectronic1,
		BleedElectronic2: data.BleedElectronic2,
		BleedElectronic3: data.BleedElectronic3,
		BleedElectronic4: data.BleedElectronic4,
		BleedElectronic5: data.BleedElectronic5,
	}
}

type AnalMeatPhMoniData struct {
	AnalMeatPhMoni1 float32 `json:"anal_meat_ph_moni_1"` //宰后0 min
	AnalMeatPhMoni2 float32 `json:"anal_meat_ph_moni_2"` //宰后45min
	AnalMeatPhMoni3 float32 `json:"anal_meat_ph_moni_3"` //劈半后胴体
	AnalMeatPhMoni4 float32 `json:"anal_meat_ph_moni_4"` //排酸过程中胴体
	AnalMeatPhMoni5 float32 `json:"anal_meat_ph_moni_5"` //冷鲜肉
}

func ToAnalMeatPhMoniData(data *AnalMeatPhMoni) AnalMeatPhMoniData {
	return AnalMeatPhMoniData{
		AnalMeatPhMoni1: data.AnalMeatPhMoni1,
		AnalMeatPhMoni2: data.AnalMeatPhMoni2,
		AnalMeatPhMoni3: data.AnalMeatPhMoni3,
		AnalMeatPhMoni4: data.AnalMeatPhMoni4,
		AnalMeatPhMoni5: data.AnalMeatPhMoni5,
	}

}

type AnalCutWeightData struct {
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

func ToAnalCutWeightData(data *AnalCutWeight) AnalCutWeightData {
	return AnalCutWeightData{
		AnalCutWeight1:  data.AnalCutWeight1,
		AnalCutWeight2:  data.AnalCutWeight2,
		AnalCutWeight3:  data.AnalCutWeight3,
		AnalCutWeight4:  data.AnalCutWeight4,
		AnalCutWeight5:  data.AnalCutWeight5,
		AnalCutWeight6:  data.AnalCutWeight6,
		AnalCutWeight7:  data.AnalCutWeight7,
		AnalCutWeight8:  data.AnalCutWeight8,
		AnalCutWeight9:  data.AnalCutWeight9,
		AnalCutWeight10: data.AnalCutWeight10,
		AnalCutWeight11: data.AnalCutWeight11,
		AnalCutWeight12: data.AnalCutWeight12,
	}
}

type ToNumGermMonData struct {
	//接触面菌落总数监控

	ToNumGermMon1 float32 `json:"to_num_germ_mon_1"` //喷淋后胴体
	ToNumGermMon2 float32 `json:"to_num_germ_mon_2"` //剔骨案板
	ToNumGermMon3 float32 `json:"to_num_germ_mon_3"` //传送带
	ToNumGermMon4 float32 `json:"to_num_germ_mon_4"` //围裙
	ToNumGermMon5 float32 `json:"to_num_germ_mon_5"` //手套
	ToNumGermMon6 float32 `json:"to_num_germ_mon_6"` //锯骨机
	ToNumGermMon7 float32 `json:"to_num_germ_mon_7"` //刀具
	ToNumGermMon8 float32 `json:"to_num_germ_mon_8"` //地面与墙面
}

func ToToNumGermMonData(data *ToNumGermMon) ToNumGermMonData {
	return ToNumGermMonData{
		ToNumGermMon1: data.ToNumGermMon1,
		ToNumGermMon2: data.ToNumGermMon2,
		ToNumGermMon3: data.ToNumGermMon3,
		ToNumGermMon4: data.ToNumGermMon4,
		ToNumGermMon5: data.ToNumGermMon5,
		ToNumGermMon6: data.ToNumGermMon6,
		ToNumGermMon7: data.ToNumGermMon7,
		ToNumGermMon8: data.ToNumGermMon8,
	}
}

type AirNumGermMonData struct {
	//空气菌落总数监控

	AirNumGermMon1 float32 `json:"air_num_germ_mon_1"` //屠宰车间
	AirNumGermMon2 float32 `json:"air_num_germ_mon_2"` //分割车间
	AirNumGermMon3 float32 `json:"air_num_germ_mon_3"` //排酸车间
}

func ToAirNumGermMonData(data *AirNumGermMon) AirNumGermMonData {
	return AirNumGermMonData{
		AirNumGermMon1: data.AirNumGermMon1,
		AirNumGermMon2: data.AirNumGermMon2,
		AirNumGermMon3: data.AirNumGermMon3,
	}
}

type PreSlaDietManageData struct {
	//宰前饮食管理

	PreSlaDietManage1 string `json:"pre_sla_diet_manage_1"` //牛进场时间
	PreSlaDietManage2 string `json:"pre_sla_diet_manage_2"` //牛停水时间
	PreSlaDietManage3 string `json:"pre_sla_diet_manage_3"` //牛禁食时间
	PreSlaDietManage4 string `json:"pre_sla_diet_manage_4"` //饮水时间
	PreSlaDietManage5 string `json:"pre_sla_diet_manage_5"` //送宰时间
}
