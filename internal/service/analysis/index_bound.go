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
