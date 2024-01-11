package analysis

const (
	// 牧场饲料重金属
	// 砷含量
	As1_MAX = 4  //干草上界
	As2_MAX = 4  //棕榈仁饼上界
	As3_MAX = 40 //藻类及其加工产品上界
	As4_MAX = 15 //甲壳类动物及其副产品（虾油除外）、鱼虾粉、水生软体动物及其副产品(油脂除外)上界
	As5_MAX = 10 //其他矿物质上界
	As6_MAX = 10 //添加剂上界
	As7_MAX = 4  //浓缩上界
	As8_MAX = 4  //精料上界
	// 铅含量
	Pb1_MAX = 5   //单细胞蛋白饲料原料上界
	Pb2_MAX = 15  //矿物质饲料原料上界
	Pb3_MAX = 30  //饲草、粗饲料及其加工产品上界
	Pb4_MAX = 10  //浓缩饲料上界
	Pb5_MAX = 8   //精料补充料上界
	Pb6_MAX = 5   //配合饲料上界
	Pb7_MAX = 0.1 //其他饲料原料上界
	// 镉含量
	Cd1_MAX = 2    //藻类及其加工产品
	Cd2_MAX = 2    //其他动物源性饲料原料
	Cd3_MAX = 2    //其他矿物质饲料原料
	Cd4_MAX = 5    //添加剂预混合饲料
	Cd5_MAX = 1.25 //浓缩饲料
	Cd6_MAX = 0.5  //犊牛精料补充料
	Cd7_MAX = 1    //其他精料补充料
	// 铬含量
	Cr1_MAX = 5 //其他添加剂预混合饲料
	Cr2_MAX = 5 //配合饲料
	Cr3_MAX = 5 //其他浓缩饲料
	Cr4_MAX = 5 //配合饲料

	// 牧场饲料真菌
	// 黄曲霉毒素B1
	Afb11_MAX = 50 //玉米加工产品、花生饼(粕)
	Afb12_MAX = 10 //植物油脂(玉米油花生油除外)
	Afb13_MAX = 20 //玉米油 、花生油
	Afb14_MAX = 30 //其他植物性饲料原料
	Afb15_MAX = 20 //其他浓缩饲料
	Afb16_MAX = 20 //犊牛、羔羊精料补充料
	Afb17_MAX = 30 //其他精料补充料
	// 玉米赤霉烯酮
	Don1_MAX = 0.5 //玉米及其加工产品 (玉米皮、喷浆玉米皮、玉米浆干粉除外)
	Don2_MAX = 1.5 //玉米皮、喷浆玉米皮.玉米浆干粉、玉米酒糟
	Don3_MAX = 1.0 //其他植物性饲料原料
	Don4_MAX = 0.5 //其他配合饲料
	// 脱氧雪腐镰刀菌烯醇(呕吐毒素）
	T2toxin1_MAX = 5 //植物性饲料原料
	T2toxin2_MAX = 3 //其他精料补充料
	T2toxin3_MAX = 3 //其他配合饲料
	// 赭曲霉毒素A  T-2毒素   伏马毒素（B1＋B2）
	T2AVomZea1_MAX = 100 //Zea 配合饲料
	T2AVomZea2_MAX = 50  //Vom 其他反刍动物精料补充料
	T2AVomZea3_MAX = 0.5 //T2	植物性饲料原料

	// 牧场水质
	// 感官性状和一般化学指标
	OapGci1_MAX  = 10    //浑浊度
	OapGci2_MAX  = 550   //总硬度（以CaCO3)计
	OapGci3_MAX  = 9.0   //PH值上界
	OapGci3_MIN  = 5.5   //PH值下界
	OapGci4_MAX  = 300   //硫酸
	OapGci5_MAX  = 300   //氯化物
	OapGci6_MAX  = 1500  //总溶解性固体
	OapGci7_MAX  = 1.2   //氟化物
	OapGci8_MAX  = 0.001 //六六六
	OapGci9_MAX  = 0.005 //滴滴涕
	OapGci10_MAX = 0.55  //氨氮
	// 毒理学指标
	ToxIndex1_MAX  = 0.05  //氰化物
	ToxIndex2_MAX  = 0.05  //总砷
	ToxIndex3_MAX  = 0.001 //总汞
	ToxIndex4_MAX  = 0.05  //总铅
	ToxIndex5_MAX  = 0.05  //铬（六价）
	ToxIndex6_MAX  = 0.01  //总镉
	ToxIndex7_MAX  = 20    //硝酸盐
	ToxIndex8_MAX  = 0.01  //砷
	ToxIndex9_MAX  = 0.005 //镉
	ToxIndex10_MAX = 1     //氯化物
	ToxIndex11_MAX = 10    //硝酸钾
	ToxIndex12_MAX = 0.06  //三氯甲烷
	ToxIndex13_MAX = 0.1   //一氯二溴甲烷
	ToxIndex14_MAX = 0.06  //二氯一溴甲烷
	ToxIndex15_MAX = 0.1   //三溴甲烷
	ToxIndex16_MAX = 1     //三卤甲烷（三氯甲烷、一氯二溴甲烷、二氯一溴甲烷、三溴甲烷的总和）
	ToxIndex17_MAX = 0.05  //二氯乙酸
	ToxIndex18_MAX = 0.1   //三氯乙酸
	ToxIndex19_MAX = 0.01  //溴酸盐
	ToxIndex20_MAX = 0.7   //亚氯酸盐
	ToxIndex21_MAX = 0.7   //氯酸盐
	ToxIndex22_MAX = 4.0   //化学耗氧量
	// 微生物指标
	MicroIndex1_MAX = 10  //总大肠菌群
	MicroIndex2_MAX = 10  //粪大肠菌群
	MicroIndex3_MAX = 100 //菌落总数

	// 牧场缓冲区
	Buffer1_MAX  = 2   //NH3氨气
	Buffer2_MAX  = 1   //H2S硫化氢
	Buffer3_MAX  = 380 //CO2二氧化碳
	Buffer4_MAX  = 0.5 //PM10
	Buffer5_MAX  = 1   //TSP
	Buffer6_MAX  = 40  //ODO 恶臭
	Buffer7_MAX  = 0.3 //Cd(镉)
	Buffer8_MAX  = 30  //As（砷）
	Buffer9_MAX  = 50  //Cu（铜）
	Buffer10_MAX = 250 //Pb（铅）
	Buffer11_MAX = 250 //Cr(铬)
	Buffer12_MAX = 200 //Zn(锌）

	// 牧场场区
	CattleFarm1_MAX  = 5   //NH3氨气
	CattleFarm2_MAX  = 2   //H2S硫化氢
	CattleFarm3_MAX  = 750 //CO2二氧化碳
	CattleFarm4_MAX  = 1   //PM10
	CattleFarm5_MAX  = 1   //TSP
	CattleFarm6_MAX  = 50  //ODO 恶臭
	CattleFarm7_MAX  = 0.3 //Cd(镉)
	CattleFarm8_MAX  = 30  //As（砷）
	CattleFarm9_MAX  = 100 //Cu（铜）
	CattleFarm10_MAX = 300 //Pb（铅）
	CattleFarm11_MAX = 300 //Cr(铬)
	CattleFarm12_MAX = 250 //Zn(锌）

	// 牧场牛舍
	CowHouse1_MAX  = 20   //NH3氨气
	CowHouse2_MAX  = 8    //H2S硫化氢
	CowHouse3_MAX  = 1500 //CO2二氧化碳
	CowHouse4_MAX  = 2    //PM10
	CowHouse5_MAX  = 4    //TSP
	CowHouse6_MAX  = 70   //ODO 恶臭
	CowHouse7_MAX  = 0.6  //Cd(镉)
	CowHouse8_MAX  = 20   //As（砷）
	CowHouse9_MAX  = 100  //Cu（铜）
	CowHouse10_MAX = 350  //Pb（铅）
	CowHouse11_MAX = 350  //Cr(铬)
	CowHouse12_MAX = 300  //Zn(锌）

	// 牧场基本环境指数
	Environment1_MAX = 15  //温度
	Environment1_MIN = 10  //温度
	Environment2_MAX = 80  //相对湿度
	Environment3_MAX = 1.0 //风速
	Environment4_MAX = 501 //照度
	Environment5_MAX = 75  //噪声
	Environment6_MAX = 7   //光照时间

	// 牧场垫料
	PaddingRequire1_MAX = 2.0     //T-Hg（总汞）
	PaddingRequire2_MAX = 30      //Pb（铅）
	PaddingRequire3_MAX = 1.0     //Cr（铬）
	PaddingRequire4_MAX = 0.5     //Cd（镉）
	PaddingRequire5_MAX = 100000  //TTC(总大肠菌群数)
	PaddingRequire6_MAX = 5000000 //TBC(细菌总数)
	PaddingRequire7_MAX = 0.006   //AFB1（黄曲霉素B1）
	PaddingRequire8_MAX = 0       //STC(沙门菌数)

	// 屠宰场水质
	// 微生物
	SlaughterWaterMicroIndex1_MAX = 0   //总大肠菌群
	SlaughterWaterMicroIndex2_MAX = 0   //大肠埃希氏菌
	SlaughterWaterMicroIndex3_MAX = 100 //菌落总数
	// 感官性状和一般化学指标
	OapGciSla1_MAX  = 2000    //钠
	OapGciSla2_MAX  = 0.002   //挥发酚类
	OapGciSla3_MAX  = 0.3     //阴离子合成洗涤剂
	OapGciSla4_MAX  = 0.00001 //2-甲基异莰醇
	OapGciSla5_MAX  = 0.00001 //土臭素
	OapGciSla6_MAX  = 15      //色度（铂钴色度单位）/度
	OapGciSla7_MAX  = 1       //浑浊度（散射浑浊度单位）/NTU
	OapGciSla8_MAX  = 0       //臭和味
	OapGciSla9_MAX  = 0       //肉眼可见物
	OapGciSla10_MAX = 8.5     //pH
	OapGciSla10_MIN = 6.5     //pH
	OapGciSla11_MAX = 0.2     //铝/（mg/L）
	OapGciSla12_MAX = 0.3     //铁/（mg/L）
	OapGciSla13_MAX = 0.1     //锰/（mg/L）
	OapGciSla14_MAX = 1       //铜/（mg/L）
	OapGciSla15_MAX = 1       //锌/（mg/L）
	OapGciSla16_MAX = 250     //氯化物/（mg/L）
	OapGciSla17_MAX = 250     //硫酸盐/（mg/L）
	OapGciSla18_MAX = 1000    //溶解性总固体/（mg/L）
	OapGciSla19_MAX = 450     //总硬度（以CaCO3计）/（mg/L）
	OapGciSla20_MAX = 3       //高锰酸盐指数（以O2计）/（mg/L）
	OapGciSla21_MAX = 0.5     //氨（以N计）/（mg/L）
	// 毒理指标
	ToxinIndexSla1_MAX             = 0.005   //锑/（mg/L）
	ToxinIndexSla2_MAX             = 0.7     //钡/（mg/L）
	ToxinIndexSla3_MAX             = 0.002   //铍/（mg/L）
	ToxinIndexSla4_MAX             = 1       //硼/（mg/L）
	ToxinIndexSla5_MAX             = 0.07    //钼/（mg/L）
	ToxinIndexSla6_MAX             = 0.02    //镍/（mg/L）
	ToxinIndexSla7_MAX             = 0.05    //银/（mg/L）
	ToxinIndexSla8_MAX             = 0.0001  //铊/（mg/L）
	ToxinIndexSla9_MAX             = 0.01    //硒/（mg/L）
	ToxinIndexSla10_MAX            = 0.07    //高氯酸钾/（mg/L）
	ToxinIndexSla11_MAX            = 0.02    //二氯甲烷/（mg/L）
	ToxinIndexSla12_MAX            = 0.03    //1，2-二氯甲烷/（mg/L）
	ToxinIndexSla13_MAX            = 0.002   //四氯化碳/（mg/L）
	ToxinIndexSla14_MAX            = 0.001   //氯乙烯/（mg/L）
	ToxinIndexSla15_MAX            = 0.03    //1，1-二氯乙烯/（mg/L）
	ToxinIndexSla16_MAX            = 0.05    //1，2-二氯乙烯（总量）/（mg/L）
	ToxinIndexSla17_MAX            = 0.02    //三氯乙烯/（mg/L）
	ToxinIndexSla18_MAX            = 0.04    //四氯乙烯/（mg/L）
	ToxinIndexSla19_MAX            = 0.0006  //六氯丁二烯/（mg/L）
	ToxinIndexSla20_MAX            = 0.01    //苯/（mg/L）
	ToxinIndexSla21_MAX            = 0.7     //甲苯/（mg/L）
	ToxinIndexSla22_MAX            = 0.5     //二甲苯（总量）/（mg/L）
	ToxinIndexSla23_MAX            = 0.02    //苯乙烯/（mg/L）
	ToxinIndexSla24_MAX            = 0.3     //氯苯/（mg/L）
	ToxinIndexSla25_MAX            = 0.3     //1，4-二氯苯/（mg/L）
	ToxinIndexSla26_MAX            = 0.02    //三氯苯（总量）/（mg/L）
	ToxinIndexSla27_MAX            = 0.001   //六氯苯/（mg/L）
	ToxinIndexSla28_MAX            = 0.0004  //七氯/（mg/L）
	ToxinIndexSla29_MAX            = 0.25    //马拉硫磷/（mg/L）
	ToxinIndexSla30_MAX            = 0.006   //乐果/（mg/L）
	ToxinIndexSla31_MAX            = 0.3     //灭草松/（mg/L）
	ToxinIndexSla32_MAX            = 0.01    //百菌清/（mg/L）
	ToxinIndexSla33_MAX            = 0.007   //呋喃丹/（mg/L）
	ToxinIndexSla34_MAX            = 0.03    //毒死蜱/（mg/L）
	ToxinIndexSla35_MAX            = 0.7     //草甘膦/（mg/L）
	ToxinIndexSla36_MAX            = 0.001   //敌敌畏/（mg/L）
	ToxinIndexSla37_MAX            = 0.002   //莠去津/（mg/L）
	ToxinIndexSla38_MAX            = 0.02    //溴氰菊酯/（mg/L）
	ToxinIndexSla39_MAX            = 0.03    //2，4-滴/（mg/L）
	ToxinIndexSla40_MAX            = 0.02    //乙草胺/（mg/L）
	ToxinIndexSla41_MAX            = 0.009   //五氯酚/（mg/L）
	ToxinIndexSla42_MAX            = 0.2     //2，4，6-三氯酚/（mg/L）
	ToxinIndexSla43_MAX            = 0.00001 //苯并（a）芘/（mg/L）
	ToxinIndexSla44_MAX            = 0.008   //邻苯二甲酸二（2-乙基己基）脂/（mg/L）
	ToxinIndexSla45_MAX            = 0.0005  //丙烯酰胺/（mg/L）
	ToxinIndexSla46_MAX            = 0.0004  //环氧氯丙烷/（mg/L）
	ToxinIndexSla47_MAX            = 0.001
	SlaughterWaterToxinIndex1_MAX  = 1     //氰化物
	SlaughterWaterToxinIndex2_MAX  = 0.01  //总砷
	SlaughterWaterToxinIndex3_MAX  = 0.001 //总汞
	SlaughterWaterToxinIndex4_MAX  = 0.01  //总铅
	SlaughterWaterToxinIndex5_MAX  = 0.05  //铬（六价）
	SlaughterWaterToxinIndex6_MAX  = 0.005 //总镉
	SlaughterWaterToxinIndex7_MAX  = 20    //硝酸盐
	SlaughterWaterToxinIndex8_MAX  = 0.1   //砷
	SlaughterWaterToxinIndex9_MAX  = 0.005 //镉
	SlaughterWaterToxinIndex10_MAX = 1     //氯化物
	SlaughterWaterToxinIndex11_MAX = 10    //硝酸钾
	SlaughterWaterToxinIndex12_MAX = 0.06  //三氯甲烷
	SlaughterWaterToxinIndex13_MAX = 0.1   //一氯二溴甲烷
	SlaughterWaterToxinIndex14_MAX = 0.06  //二氯一溴甲烷
	SlaughterWaterToxinIndex15_MAX = 0.1   //三溴甲烷
	SlaughterWaterToxinIndex16_MAX = 1     //三卤甲烷（三氯甲烷、一氯二溴甲烷、二氯一溴甲烷、三溴甲烷的总和）
	SlaughterWaterToxinIndex17_MAX = 0.05  //二氯乙酸
	SlaughterWaterToxinIndex18_MAX = 0.1   //三氯乙酸
	SlaughterWaterToxinIndex19_MAX = 0.01  //溴酸盐
	SlaughterWaterToxinIndex20_MAX = 0.7   //亚氯酸盐
	SlaughterWaterToxinIndex21_MAX = 0.7   //氯酸盐

	// 屠宰过程中参数
	Stun1_MAX            = 200 //电压
	Stun2_MAX            = 1.5 //电流上界
	Stun2_MIN            = 1   //电流下界
	Stun3_MIN            = 7   //作用时间
	BleedElectronic1_MIN = 82  //放血刀热水消毒
	BleedElectronic2_MIN = 6   //沥血时间
	BleedElectronic3_MAX = 15  //致昏到宰杀放血时间
	BleedElectronic4_MAX = 42  //电刺激参数
	BleedElectronic5_MIN = 15  //电刺激时间
	AnalMeatPhMoni1_MAX  = 6   //宰后0 min PH上界
	AnalMeatPhMoni1_MIN  = 5.4 //宰后0 min PH上界
	AnalMeatPhMoni2_MAX  = 6   //宰后45min PH上界
	AnalMeatPhMoni2_MIN  = 5.4 //宰后45min PH下界
	AnalMeatPhMoni3_MAX  = 6   //劈半后胴体 PH上界
	AnalMeatPhMoni3_MIN  = 5.4 //劈半后胴体 PH下界
	AnalMeatPhMoni4_MAX  = 6   //排酸过程中胴体 PH上界
	AnalMeatPhMoni4_MIN  = 5.4 //排酸过程中胴体 PH下界
	AnalMeatPhMoni5_MAX  = 6   //冷鲜肉 PH上界
	AnalMeatPhMoni5_MIN  = 5.4 //冷鲜肉 PH下界

	// 屠宰场预冷间指标
	PreCoolShop1_MAX = 4  //预冷间温度
	PreCoolShop1_MIN = 0  //预冷间温度
	PreCoolShop2_MAX = 90 //预冷间湿度
	PreCoolShop2_MIN = 85 //预冷间湿度
	PreCoolShop3_MAX = 3  //副产物温度

)

var PastureFeedHeavyMetalBounds = map[string]float64{"as_1": As1_MAX, "as_2": As2_MAX, "as_3": As3_MAX,
	"as_4": As4_MAX, "as_5": As5_MAX, "as_6": As6_MAX, "as_7": As7_MAX, "as_8": As8_MAX, "pb_1": Pb1_MAX,
	"pb_2": Pb2_MAX, "pb_3": Pb3_MAX, "pb_4": Pb4_MAX, "pb_5": Pb5_MAX, "pb_6": Pb6_MAX, "pb_7": Pb7_MAX,
	"cd_1": Cd1_MAX, "cd_2": Cd2_MAX, "cd_3": Cd3_MAX, "cd_4": Cd4_MAX, "cd_5": Cd5_MAX, "cd_6": Cd6_MAX,
	"cd_7": Cd7_MAX, "cr_1": Cr1_MAX, "cr_2": Cr2_MAX, "cr_3": Cr3_MAX, "cr_4": Cr4_MAX}

var PastureFeedMycotoxinsBounds = map[string]float64{"afb_11": Afb11_MAX, "afb_12": Afb12_MAX, "afb_13": Afb13_MAX,
	"afb_14": Afb14_MAX, "afb_15": Afb15_MAX, "afb_16": Afb16_MAX, "afb_17": Afb17_MAX, "don_1": Don1_MAX, "don_2": Don2_MAX,
	"don_3": Don3_MAX, "don_4": Don4_MAX, "t_2_toxin_1": T2toxin1_MAX, "t_2_toxin_2": T2toxin2_MAX,
	"t_2_toxin_3": T2toxin3_MAX, "t_2_a_vom_zea_1": T2AVomZea1_MAX, "t_2_a_vom_zea_2": T2AVomZea2_MAX,
	"t_2_a_vom_zea_3": T2AVomZea3_MAX}

var PastureWaterQualityUpperBounds = map[string]float64{"oap_gci_1": OapGci1_MAX, "oap_gci_2": OapGci2_MAX, "oap_gci_3": OapGci3_MAX,
	"oap_gci_4": OapGci4_MAX, "oap_gci_5": OapGci5_MAX, "oap_gci_6": OapGci6_MAX, "oap_gci_7": OapGci7_MAX, "oap_gci_8": OapGci8_MAX,
	"oap_gci_9": OapGci9_MAX, "oap_gci_10": OapGci10_MAX, "tox_index_1": ToxIndex1_MAX, "tox_index_2": ToxIndex2_MAX, "tox_index_3": ToxIndex3_MAX,
	"tox_index_4": ToxIndex4_MAX, "tox_index_5": ToxIndex5_MAX, "tox_index_6": ToxIndex6_MAX, "tox_index_7": ToxIndex7_MAX, "tox_index_8": ToxIndex8_MAX,
	"tox_index_9": ToxIndex9_MAX, "tox_index_10": ToxIndex10_MAX, "tox_index_11": ToxIndex11_MAX, "tox_index_12": ToxIndex12_MAX, "tox_index_13": ToxIndex13_MAX,
	"tox_index_14": ToxIndex14_MAX, "tox_index_15": ToxIndex15_MAX, "tox_index_16": ToxIndex16_MAX, "tox_index_17": ToxIndex17_MAX, "tox_index_18": ToxIndex18_MAX,
	"tox_index_19": ToxIndex19_MAX, "tox_index_20": ToxIndex20_MAX, "tox_index_21": ToxIndex21_MAX, "tox_index_22": ToxIndex22_MAX, "micro_index_1": MicroIndex1_MAX,
	"micro_index_2": MicroIndex2_MAX, "micro_index_3": MicroIndex3_MAX}
var PastureWaterQualityLowerBounds = map[string]float64{"oap_gci_3": OapGci3_MIN}

var PastureBufferUpperBounds = map[string]float64{"buffer_1": Buffer1_MAX, "buffer_2": Buffer2_MAX, "buffer_3": Buffer3_MAX, "buffer_4": Buffer4_MAX,
	"buffer_5": Buffer5_MAX, "buffer_6": Buffer6_MAX, "buffer_7": Buffer7_MAX, "buffer_8": Buffer8_MAX, "buffer_9": Buffer9_MAX, "buffer_10": Buffer10_MAX,
	"buffer_11": Buffer11_MAX, "buffer_12": Buffer12_MAX}

var PastureAreaUpperBounds = map[string]float64{"cattle_farm_1": CattleFarm1_MAX, "cattle_farm_2": CattleFarm2_MAX, "cattle_farm_3": CattleFarm3_MAX, "cattle_farm_4": CattleFarm4_MAX,
	"cattle_farm_5": CattleFarm5_MAX, "cattle_farm_6": CattleFarm6_MAX, "cattle_farm_7": CattleFarm7_MAX, "cattle_farm_8": CattleFarm8_MAX, "cattle_farm_9": CattleFarm9_MAX,
	"cattle_farm_10": CattleFarm10_MAX, "cattle_farm_11": CattleFarm11_MAX, "cattle_farm_12": CattleFarm12_MAX}

var PastureCowHouseUpperBounds = map[string]float64{"cow_house_1": CowHouse1_MAX, "cow_house_2": CowHouse2_MAX, "cow_house_3": CowHouse3_MAX, "cow_house_4": CowHouse4_MAX, "cow_house_5": CowHouse5_MAX,
	"cow_house_6": CowHouse6_MAX, "cow_house_7": CowHouse7_MAX, "cow_house_8": CowHouse8_MAX, "cow_house_9": CowHouse9_MAX, "cow_house_10": CowHouse10_MAX, "cow_house_11": CowHouse11_MAX, "cow_house_12": CowHouse12_MAX}

var PastureBasicEnvironmentUpperBounds = map[string]float64{"environment_1": Environment1_MAX, "environment_2": Environment2_MAX, "environment_3": Environment3_MAX, "environment_4": Environment4_MAX,
	"environment_5": Environment5_MAX, "environment_6": Environment6_MAX}
var PastureBasicEnvironmentLowerBounds = map[string]float64{"environment_1": Environment1_MIN}

var PasturePaddingUpperBounds = map[string]float64{"padding_require_1": PaddingRequire1_MAX, "padding_require_2": PaddingRequire2_MAX, "padding_require_3": PaddingRequire3_MAX, "padding_require_4": PaddingRequire4_MAX,
	"padding_require_5": PaddingRequire5_MAX, "padding_require_6": PaddingRequire6_MAX, "padding_require_7": PaddingRequire7_MAX, "padding_require_8": PaddingRequire8_MAX}

var SlaughterWaterQualityUpperBounds = map[string]float64{"slaughter_water_micro_index_1": SlaughterWaterMicroIndex1_MAX, "slaughter_water_micro_index_2": SlaughterWaterMicroIndex2_MAX, "slaughter_water_micro_index_3": SlaughterWaterMicroIndex3_MAX,
	"oap_gci_sla_1": OapGciSla1_MAX, "oap_gci_sla_2": OapGciSla2_MAX, "oap_gci_sla_3": OapGciSla3_MAX, "oap_gci_sla_4": OapGciSla4_MAX, "oap_gci_sla_5": OapGciSla5_MAX, "oap_gci_sla_6": OapGciSla6_MAX, "oap_gci_sla_7": OapGciSla7_MAX,
	"oap_gci_sla_8": OapGciSla8_MAX, "oap_gci_sla_9": OapGciSla9_MAX, "oap_gci_sla_10": OapGciSla10_MAX, "oap_gci_sla_11": OapGciSla11_MAX, "oap_gci_sla_12": OapGciSla12_MAX, "oap_gci_sla_13": OapGciSla13_MAX, "oap_gci_sla_14": OapGciSla14_MAX,
	"oap_gci_sla_15": OapGciSla15_MAX, "oap_gci_sla_16": OapGciSla16_MAX, "oap_gci_sla_17": OapGciSla17_MAX, "oap_gci_sla_18": OapGciSla18_MAX, "oap_gci_sla_19": OapGciSla19_MAX, "oap_gci_sla_20": OapGciSla20_MAX, "oap_gci_sla_21": OapGciSla21_MAX,
	"toxin_index_sla_1": ToxinIndexSla1_MAX, "toxin_index_sla_2": ToxinIndexSla2_MAX, "toxin_index_sla_3": ToxinIndexSla3_MAX, "toxin_index_sla_4": ToxinIndexSla4_MAX, "toxin_index_sla_5": ToxinIndexSla5_MAX, "toxin_index_sla_6": ToxinIndexSla6_MAX,
	"toxin_index_sla_7": ToxinIndexSla7_MAX, "toxin_index_sla_8": ToxinIndexSla8_MAX, "toxin_index_sla_9": ToxinIndexSla9_MAX, "toxin_index_sla_10": ToxinIndexSla10_MAX, "toxin_index_sla_11": ToxinIndexSla11_MAX, "toxin_index_sla_12": ToxinIndexSla12_MAX,
	"toxin_index_sla_13": ToxinIndexSla13_MAX, "toxin_index_sla_14": ToxinIndexSla14_MAX, "toxin_index_sla_15": ToxinIndexSla15_MAX, "toxin_index_sla_16": ToxinIndexSla16_MAX, "toxin_index_sla_17": ToxinIndexSla17_MAX, "toxin_index_sla_18": ToxinIndexSla18_MAX,
	"toxin_index_sla_19": ToxinIndexSla19_MAX, "toxin_index_sla_20": ToxinIndexSla20_MAX, "toxin_index_sla_21": ToxinIndexSla21_MAX, "toxin_index_sla_22": ToxinIndexSla22_MAX, "toxin_index_sla_23": ToxinIndexSla23_MAX, "toxin_index_sla_24": ToxinIndexSla24_MAX,
	"toxin_index_sla_25": ToxinIndexSla25_MAX, "toxin_index_sla_26": ToxinIndexSla26_MAX, "toxin_index_sla_27": ToxinIndexSla27_MAX, "toxin_index_sla_28": ToxinIndexSla28_MAX, "toxin_index_sla_29": ToxinIndexSla29_MAX, "toxin_index_sla_30": ToxinIndexSla30_MAX,
	"toxin_index_sla_31": ToxinIndexSla31_MAX, "toxin_index_sla_32": ToxinIndexSla32_MAX, "toxin_index_sla_33": ToxinIndexSla33_MAX, "toxin_index_sla_34": ToxinIndexSla34_MAX, "toxin_index_sla_35": ToxinIndexSla35_MAX, "toxin_index_sla_36": ToxinIndexSla36_MAX,
	"toxin_index_sla_37": ToxinIndexSla37_MAX, "toxin_index_sla_38": ToxinIndexSla38_MAX, "toxin_index_sla_39": ToxinIndexSla39_MAX, "toxin_index_sla_40": ToxinIndexSla40_MAX, "toxin_index_sla_41": ToxinIndexSla41_MAX, "toxin_index_sla_42": ToxinIndexSla42_MAX,
	"toxin_index_sla_43": ToxinIndexSla43_MAX, "toxin_index_sla_44": ToxinIndexSla44_MAX, "toxin_index_sla_45": ToxinIndexSla45_MAX, "toxin_index_sla_46": ToxinIndexSla46_MAX, "toxin_index_sla_47": ToxinIndexSla47_MAX, "slaughter_water_toxin_index_1": SlaughterWaterToxinIndex1_MAX,
	"slaughter_water_toxin_index_2": SlaughterWaterToxinIndex2_MAX, "slaughter_water_toxin_index_3": SlaughterWaterToxinIndex3_MAX, "slaughter_water_toxin_index_4": SlaughterWaterToxinIndex4_MAX, "slaughter_water_toxin_index_5": SlaughterWaterToxinIndex5_MAX,
	"slaughter_water_toxin_index_6": SlaughterWaterToxinIndex6_MAX, "slaughter_water_toxin_index_7": SlaughterWaterToxinIndex7_MAX, "slaughter_water_toxin_index_8": SlaughterWaterToxinIndex8_MAX, "slaughter_water_toxin_index_9": SlaughterWaterToxinIndex9_MAX,
	"slaughter_water_toxin_index_10": SlaughterWaterToxinIndex10_MAX, "slaughter_water_toxin_index_11": SlaughterWaterToxinIndex11_MAX, "slaughter_water_toxin_index_12": SlaughterWaterToxinIndex12_MAX, "slaughter_water_toxin_index_13": SlaughterWaterToxinIndex13_MAX,
	"slaughter_water_toxin_index_14": SlaughterWaterToxinIndex14_MAX, "slaughter_water_toxin_index_15": SlaughterWaterToxinIndex15_MAX, "slaughter_water_toxin_index_16": SlaughterWaterToxinIndex16_MAX, "slaughter_water_toxin_index_17": SlaughterWaterToxinIndex17_MAX,
	"slaughter_water_toxin_index_18": SlaughterWaterToxinIndex18_MAX, "slaughter_water_toxin_index_19": SlaughterWaterToxinIndex19_MAX, "slaughter_water_toxin_index_20": SlaughterWaterToxinIndex20_MAX, "slaughter_water_toxin_index_21": SlaughterWaterToxinIndex21_MAX}
var SlaughterWaterQualityLowerBounds = map[string]float64{"oap_gci_sla_10": OapGciSla10_MIN}

var SlaughterProcedureSensorDataUpperBounds = map[string]float64{"stun_1": Stun1_MAX, "stun_2": Stun2_MAX, "bleed_electronic_3": BleedElectronic3_MAX, "bleed_electronic_4": BleedElectronic4_MAX, "anal_meat_ph_moni_1": AnalMeatPhMoni1_MAX, "anal_meat_ph_moni_2": AnalMeatPhMoni2_MAX,
	"anal_meat_ph_moni_3": AnalMeatPhMoni3_MAX, "anal_meat_ph_moni_4": AnalMeatPhMoni4_MAX, "anal_meat_ph_moni_5": AnalMeatPhMoni5_MAX}
var SlaughterProcedureSensorDataLowerBounds = map[string]float64{"stun_2": Stun2_MIN, "stun_3": Stun3_MIN, "bleed_electronic_1": BleedElectronic1_MIN, "bleed_electronic_2": BleedElectronic2_MIN, "bleed_electronic_5": BleedElectronic5_MIN, "anal_meat_ph_moni_1": AnalMeatPhMoni1_MIN,
	"anal_meat_ph_moni_2": AnalMeatPhMoni2_MIN, "anal_meat_ph_moni_3": AnalMeatPhMoni3_MIN, "anal_meat_ph_moni_4": AnalMeatPhMoni4_MIN, "anal_meat_ph_moni_5": AnalMeatPhMoni5_MIN}
